package main

import (
	"app.inherited.magic/config"
	"app.inherited.magic/internal/interactor/util/tunnel"
	"strconv"
	"strings"
	"time"
)

func main() {
	var local, remote []tunnel.Endpoint
	localForward := strings.Split(config.SSHLocalForward, ",")
	for _, forward := range localForward {
		desPatch := strings.Split(forward, ":")
		port, _ := strconv.Atoi(desPatch[0])
		desPatchLocal := tunnel.Endpoint{
			Host: "127.0.0.1",
			Port: int32(port),
		}

		port, _ = strconv.Atoi(desPatch[2])
		desPatchRemote := tunnel.Endpoint{
			Host: desPatch[1],
			Port: int32(port),
		}

		local = append(local, desPatchLocal)
		remote = append(remote, desPatchRemote)
	}

	ssh := tunnel.SSHConfig{
		User:         config.SSHUser,
		AuthKey:      config.SSHAuthKey,
		Password:     config.SSHPassword,
		AuthPassword: config.SSHAuthPassword,
		Server: tunnel.Endpoint{
			Host: config.SSHAddress,
			Port: config.SSHPort,
		},
		Local:            local,
		Remote:           remote,
		Timeout:          5 * time.Second,
		Debug:            false,
		IsLongConnection: true,
	}

	ssh.Run()
}
