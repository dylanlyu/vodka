package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

func AesEncryptECB(origData []byte, key []byte) (encrypted []byte, err error) {
	signal, err := aes.NewCipher(generateKey(key))
	if err != nil {
		return nil, err
	}

	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}

	encrypted = make([]byte, len(plain))
	for bs, be := 0, signal.BlockSize(); bs <= len(origData); bs, be = bs+signal.BlockSize(), be+signal.BlockSize() {
		signal.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted, nil
}

func AesDecryptECB(encrypted []byte, key []byte) (decrypted []byte, err error) {
	signal, err := aes.NewCipher(generateKey(key))
	if err != nil {
		return nil, err
	}

	decrypted = make([]byte, len(encrypted))
	for bs, be := 0, signal.BlockSize(); bs < len(encrypted); bs, be = bs+signal.BlockSize(), be+signal.BlockSize() {
		signal.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim], nil
}

func AesEncryptCBC(origData []byte, key []byte) (encrypted []byte, err error) {
	signal, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := signal.BlockSize()
	origData = pkcs7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(signal, key[:blockSize])
	encrypted = make([]byte, len(origData))
	blockMode.CryptBlocks(encrypted, origData)
	return encrypted, nil
}

func AesDecryptCBC(encrypted []byte, key []byte) (decrypted []byte, err error) {
	signal, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := signal.BlockSize()
	blockMode := cipher.NewCBCDecrypter(signal, key[:blockSize])
	decrypted = make([]byte, len(encrypted))
	blockMode.CryptBlocks(decrypted, encrypted)
	decrypted = unPadding(decrypted)
	return decrypted, err
}

func AesCryptCTR(plainText []byte, key []byte) (dst []byte, err error) {
	signal, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	iv := bytes.Repeat([]byte("1"), signal.BlockSize())
	stream := cipher.NewCFBEncrypter(signal, iv)
	dst = make([]byte, len(plainText))
	stream.XORKeyStream(dst, plainText)
	return dst, nil
}

func AesEncryptCFB(origData []byte, key []byte) (encrypted []byte, err error) {
	signal, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	encrypted = make([]byte, aes.BlockSize+len(origData))
	iv := encrypted[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(signal, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], origData)
	return encrypted, err
}

func AesDecryptCFB(encrypted []byte, key []byte) (decrypted []byte, err error) {
	signal, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(encrypted) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(signal, iv)
	stream.XORKeyStream(encrypted, encrypted)
	return encrypted, nil
}

func AesEncryptOFB(origData []byte, key []byte) (encrypted []byte, err error) {
	signal, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	origData = pkcs7Padding(origData, aes.BlockSize)
	encrypted = make([]byte, aes.BlockSize+len(origData))
	iv := encrypted[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewOFB(signal, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], origData)
	return encrypted, err
}

func AesDecryptOFB(encrypted []byte, key []byte) (decrypted []byte, err error) {
	signal, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	iv := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]
	if len(encrypted)%aes.BlockSize != 0 {
		return nil, errors.New("ciphertext too short")
	}

	decrypted = make([]byte, len(encrypted))
	mode := cipher.NewOFB(signal, iv)
	mode.XORKeyStream(decrypted, encrypted)
	decrypted = unPadding(decrypted)
	return decrypted, nil
}

func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

func pkcs5Padding(ciphertext []byte) []byte {
	return pkcs7Padding(ciphertext, 8)
}

func pkcs7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

func unPadding(origData []byte) []byte {
	length := len(origData)
	padding := int(origData[length-1])
	return origData[:(length - padding)]
}
