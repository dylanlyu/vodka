package tunnel

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"net"
	"runtime"
	"time"
)

var err error

type Endpoint struct {
	Host string
	Port int32
}

func (endpoint *Endpoint) getLocation() string {
	return fmt.Sprintf("%s:%d", endpoint.Host, endpoint.Port)
}

func (endpoint *Endpoint) getNetwork() string {
	return fmt.Sprintf("%s", "tcp")
}

type SSHConfig struct {
	User         string
	Password     string
	AuthKey      string
	AuthPassword string
	Server       Endpoint
	Local        []Endpoint
	Remote       []Endpoint
	Timeout      time.Duration
	Debug        bool

	IsLongConnection bool
	clientConfig     *ssh.ClientConfig
	client           *ssh.Client
	errCh            chan error
	localListen      []net.Listener
}

func (config *SSHConfig) Run() {
	if config.Debug {
		log.Println("debug model.")
		go func() {
			for true {
				log.Printf("goroutine times per second: %d", runtime.NumGoroutine()-1)
				time.Sleep(time.Second * 1)
			}
		}()
	}

	if !config.Debug {
		log.Println("prod model.")
	}

	log.Println("ssh tunnel is run.")
	config.errCh = make(chan error)
	config.clientConfig, err = config.getServerConfig()
	if err != nil {
		config.errCh <- fmt.Errorf("get ssh client config error :%s", err)
	}

	if config.IsLongConnection == true {
		if config.Debug {
			log.Printf("long connection to %s ", config.Server.getLocation())
		}

		if !config.Debug {
			log.Println("long connection model.")
		}

		config.client, err = ssh.Dial(config.Server.getNetwork(), config.Server.getLocation(), config.clientConfig)
		if err != nil {
			config.errCh <- fmt.Errorf("conncet client error :%s", err)
		}

		defer func(client *ssh.Client) {
			err = client.Close()
			if err != nil {
				config.errCh <- fmt.Errorf("close client error :%s", err)
			}
		}(config.client)
	}

	for i := range config.Local {
		config.createLocal(i)
	}

	select {
	case err := <-config.errCh:
		log.Printf("%s", err)
		err = config.client.Close()
		if err != nil {
			log.Printf("close client error :%s", err)
		}

		for _, listener := range config.localListen {
			err = listener.Close()
			if err != nil {
				log.Printf("close local listener error :%s", err)
			}
		}

		return
	}

}

func (config *SSHConfig) createLocal(count int) {
	sshConfig := &SSHConfig{}
	marshal, err := json.Marshal(config)
	if err != nil {
		config.errCh <- fmt.Errorf("json marshal error: %s", err.Error())
		return
	}

	err = json.Unmarshal(marshal, sshConfig)
	if err != nil {
		config.errCh <- fmt.Errorf("json unmarshal error: %s", err.Error())
		return
	}

	if config.clientConfig != nil {
		sshConfig.clientConfig = config.clientConfig
	}

	if config.client != nil {
		sshConfig.client = config.client
	}

	log.Printf("start listen local: %s", sshConfig.Local[count].getLocation())
	localListen, err := net.Listen(sshConfig.Local[count].getNetwork(), sshConfig.Local[count].getLocation())
	if err != nil {
		config.errCh <- fmt.Errorf("local listen on %s failed: %s", sshConfig.Local[count].getLocation(), err.Error())
		return
	}

	config.localListen = append(config.localListen, localListen)
	go func() {
		for true {
			if config.Debug {
				log.Printf("accepted connection from %s", sshConfig.Local[count].getLocation())
			}

			localConn, err := localListen.Accept()
			if err != nil {
				config.errCh <- fmt.Errorf("local accept on %s failed: %s", sshConfig.Local[count].getLocation(), err.Error())
				return
			}

			go config.forward(localConn, sshConfig, count)
		}
	}()
}

func (config *SSHConfig) forward(local net.Conn, sshConfig *SSHConfig, count int) {
	defer func(local net.Conn) {
		err = local.Close()
		if err != nil {
			config.errCh <- fmt.Errorf("close loacl accept error :%s", err)
			return
		}
	}(local)

	if sshConfig.IsLongConnection == false {
		if config.Debug {
			log.Printf("short connection to %s ", config.Server.getLocation())
		}

		if !config.Debug {
			log.Println("short connection model.")
		}

		sshConfig.client, err = ssh.Dial(config.Server.getNetwork(), config.Server.getLocation(), config.clientConfig)
		if err != nil {
			config.errCh <- fmt.Errorf("conncet client error :%s", err)
			return
		}

		defer func(client *ssh.Client) {
			err = client.Close()
			if err != nil {
				config.errCh <- fmt.Errorf("close client error :%s", err)
				return
			}
		}(sshConfig.client)
	}

	remote, err := sshConfig.client.Dial(sshConfig.Remote[count].getNetwork(), sshConfig.Remote[count].getLocation())
	if err != nil {
		config.errCh <- fmt.Errorf("remote dial to %s failed: %s", sshConfig.Local[count].getLocation(), err.Error())
		return
	}
	defer func(remote net.Conn) {
		err := remote.Close()
		if err != nil {
			config.errCh <- fmt.Errorf("close remote error :%s", err)
			return
		}
	}(remote)

	connStr := fmt.Sprintf("%s(tcp) <-> %s(ssh) <-> %s(tcp) <-> %s(tcp)",
		sshConfig.Local[count].getLocation(), sshConfig.Server.getLocation(), local.RemoteAddr().String(),
		sshConfig.Remote[count].getLocation())
	log.Printf("ssh tunnel open: %s", connStr)
	wait := make(chan bool)
	go func() {
		config.ioCopy(remote, local)
	}()

	go func() {
		config.ioCopy(local, remote)
		wait <- true
	}()

	<-wait
	log.Printf("ssl tunnel close: %s", connStr)
}

func (config *SSHConfig) ioCopy(source net.Conn, target net.Conn) {
	_, err = io.Copy(source, target)
	if err != nil {
		log.Printf("error on io.Copy: %s", err)
	}
}

func (config *SSHConfig) getServerConfig() (clientConfig *ssh.ClientConfig, err error) {
	clientConfig = &ssh.ClientConfig{
		User:            config.User,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         config.Timeout,
	}

	authMethod, err := config.getAuthMethod()
	if err != nil {
		return nil, err
	}

	clientConfig.Auth = []ssh.AuthMethod{authMethod}
	return clientConfig, nil
}

func (config *SSHConfig) getAuthMethod() (authMethod ssh.AuthMethod, err error) {
	if config.Password != "" {
		authMethod = ssh.Password(config.Password)
	}

	if config.AuthKey != "" {
		key, err := config.parsePrivateKey()
		if err != nil {
			return nil, err
		}

		authMethod = ssh.PublicKeys(key)
	}

	return authMethod, nil
}

func (config *SSHConfig) parsePrivateKey() (key ssh.Signer, err error) {
	if config.AuthPassword != "" {
		key, err = ssh.ParsePrivateKeyWithPassphrase([]byte(config.AuthKey), []byte(config.AuthPassword))
	} else {
		key, err = ssh.ParsePrivateKey([]byte(config.AuthKey))
	}

	if err != nil {
		return nil, err
	}

	return key, nil
}
