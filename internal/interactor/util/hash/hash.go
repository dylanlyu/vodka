package hash

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"io/ioutil"
)

// Base64StdEncode encode string with base64 encoding
func Base64StdEncode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

// Base64BydEncode encode byte with base64 encoding
func Base64BydEncode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Base64StdDecode decode a base64 encoded string
func Base64StdDecode(data string) []byte {
	b, _ := base64.StdEncoding.DecodeString(data)
	return b
}

// Md5String return the md5 value of string
func Md5String(data string) string {
	hash := md5.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

// Md5File return the md5 value of file
func Md5File(filename string) (string, error) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	hash := md5.New()
	hash.Write(f)
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// HmacMd5 return the hmac hash of string use md5
func HmacMd5(data, key string) string {
	hash := hmac.New(md5.New, []byte(key))
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum([]byte("")))
}

// HmacSha1 return the hmac hash of string use sha1
func HmacSha1(data, key string) string {
	hash := hmac.New(sha1.New, []byte(key))
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum([]byte("")))
}

// HmacSha256 return the hmac hash of string use sha256
func HmacSha256(data, key string) string {
	hash := hmac.New(sha256.New, []byte(key))
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum([]byte("")))
}

// HmacSha512 return the hmac hash of string use sha512
func HmacSha512(data, key string) string {
	hash := hmac.New(sha512.New, []byte(key))
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum([]byte("")))
}

// Sha1 return the sha1 value (SHA-1 hash algorithm) of string
func Sha1(data string) string {
	hash := sha1.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum([]byte("")))
}

// Sha256 return the sha256 value (SHA256 hash algorithm) of string
func Sha256(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum([]byte("")))
}

// Sha512 return the sha512 value (SHA512 hash algorithm) of string
func Sha512(data string) string {
	hash := sha512.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum([]byte("")))
}
