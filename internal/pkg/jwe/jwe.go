package jwe

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwe"
	"github.com/lestrrat-go/jwx/jwt"
	"strings"
)

type JWT struct {
	//iss  【issuer】發布者的url地址
	IssuerKey interface{}
	//sub 【subject】該JWT所面向的用戶，用於處理特定應用，不是常用的字段
	SubjectKey interface{}
	//aud 【audience】接受者的url地址
	AudienceKey interface{}
	//exp 【expiration】 該jwt銷毀的時間；unix時間戳
	ExpirationKey interface{}
	//nbf  【not before】 該jwt的使用時間不能早於該時間；unix時間戳
	NotBeforeKey interface{}
	//iat   【issued at】 該jwt的發佈時間；unix 時間戳
	IssuedAtKey interface{}
	//jti    【JWT ID】 該jwt的唯一ID編號
	JwtIDKey interface{}
	//其他設定
	Other map[string]interface{}
	//令牌
	Token string
	//公開金鑰RSA_256
	PublicKey string
	//私人鑰匙RSA_256
	PrivateKey string
}

func (j *JWT) Created() (*JWT, error) {
	if len(j.PublicKey) <= 0 {
		return nil, errors.New("public key is empty")
	}

	key, _ := pem.Decode([]byte(j.PublicKey))
	public, _ := x509.ParsePKIXPublicKey(key.Bytes)
	publicKey := public.(*rsa.PublicKey)
	token := jwt.New()
	_ = token.Set(jwt.IssuerKey, j.IssuerKey)
	_ = token.Set(jwt.SubjectKey, j.SubjectKey)
	_ = token.Set(jwt.AudienceKey, j.AudienceKey)
	_ = token.Set(jwt.ExpirationKey, j.ExpirationKey)
	_ = token.Set(jwt.NotBeforeKey, j.NotBeforeKey)
	_ = token.Set(jwt.IssuedAtKey, j.IssuedAtKey)
	_ = token.Set(jwt.JwtIDKey, j.JwtIDKey)
	for k, v := range j.Other {
		_ = token.Set(k, v)
	}

	payload, err := json.MarshalIndent(token, "", "  ")
	if err != nil {
		return nil, err
	}

	protected := jwe.NewHeaders()
	_ = protected.Set(`type`, `JWE`)
	encrypted, err := jwe.Encrypt(payload, jwa.RSA_OAEP_256, publicKey,
		jwa.A256CBC_HS512, jwa.Deflate, jwe.WithProtectedHeaders(protected))
	if err != nil {
		return nil, err
	}

	j.Token = string(encrypted)
	return j, nil
}

func (j *JWT) Verify() (*JWT, error) {
	if len(j.PrivateKey) <= 0 {
		return nil, errors.New("public key is empty")
	}

	j.Token = strings.Replace(j.Token, "BASE ", "", -1)
	j.Token = strings.Replace(j.Token, "BEARER ", "", -1)
	j.Token = strings.Replace(j.Token, "base ", "", -1)
	j.Token = strings.Replace(j.Token, "bearer ", "", -1)
	j.Token = strings.Replace(j.Token, "Base ", "", -1)
	j.Token = strings.Replace(j.Token, "Bearer ", "", -1)
	key, _ := pem.Decode([]byte(j.PrivateKey))
	private, _ := x509.ParsePKCS8PrivateKey(key.Bytes)
	privateKey := private.(*rsa.PrivateKey)
	token, err := jwt.Parse([]byte(j.Token), jwt.WithValidate(true),
		jwt.WithDecrypt(jwa.RSA_OAEP_256, privateKey))
	if err != nil {
		return nil, err
	}

	if token.Expiration().IsZero() {
		return nil, errors.New("decrypt error")
	}

	j.IssuerKey, _ = token.Get(jwt.IssuerKey)
	j.SubjectKey, _ = token.Get(jwt.SubjectKey)
	j.AudienceKey, _ = token.Get(jwt.AudienceKey)
	j.ExpirationKey, _ = token.Get(jwt.ExpirationKey)
	j.NotBeforeKey, _ = token.Get(jwt.NotBeforeKey)
	j.IssuedAtKey, _ = token.Get(jwt.IssuedAtKey)
	j.JwtIDKey, _ = token.Get(jwt.JwtIDKey)
	j.Other = token.PrivateClaims()
	return j, nil
}
