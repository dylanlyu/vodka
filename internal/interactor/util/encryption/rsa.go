package encryption

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

func GetRSAKey(bits int) (privateKey, publicKey string, err error) {
	key, err := GenRSAPrivateKey(bits)
	if err != nil {
		return "", "", err
	}

	private, err := GetRSAPrivateKey(key)
	if err != nil {
		return "", "", err
	}

	public, err := GetRSAPublicKey(key)
	if err != nil {
		return "", "", err
	}

	return string(private), string(public), nil
}

func GenRSAPrivateKey(bits int) (key *rsa.PrivateKey, err error) {
	key, err = rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func GetRSAPrivateKey(key *rsa.PrivateKey) (privateKey []byte, err error) {
	derStream, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		return nil, err
	}

	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}

	return pem.EncodeToMemory(block), nil
}

func GetRSAPublicKey(key *rsa.PrivateKey) (publicKey []byte, err error) {
	public := &key.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(public)
	if err != nil {
		panic(err)
	}

	block := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}

	return pem.EncodeToMemory(block), nil
}
