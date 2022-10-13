package jwx

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"strings"
)

type JWT struct {
	//iss  【issuer】發布者的url地址
	IssuerKey any
	//sub 【subject】該JWT所面向的用戶，用於處理特定應用，不是常用的字段
	SubjectKey any
	//aud 【audience】接受者的url地址
	AudienceKey any
	//exp 【expiration】 該jwt銷毀的時間；unix時間戳
	ExpirationKey any
	//nbf  【not before】 該jwt的使用時間不能早於該時間；unix時間戳
	NotBeforeKey any
	//iat   【issued at】 該jwt的發佈時間；unix 時間戳
	IssuedAtKey any
	//jti    【JWT ID】 該jwt的唯一ID編號
	JwtIDKey any
	//其他設定
	Other map[string]any
	//令牌
	Token string
	//公開金鑰RSA_256
	PublicKey string
	//私人鑰匙RSA_256
	PrivateKey string
}

func (j *JWT) Create() (*JWT, error) {
	if len(j.PrivateKey) <= 0 {
		return nil, errors.New("public key is empty")
	}

	key, _ := pem.Decode([]byte(j.PrivateKey))
	private, err := x509.ParsePKCS8PrivateKey(key.Bytes)
	if err != nil {
		return nil, err
	}

	privateKey := private.(*rsa.PrivateKey)
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

	encrypted, err := jwt.Sign(token, jwt.WithKey(jwa.RS512, privateKey))
	if err != nil {
		return nil, err
	}

	j.Token = string(encrypted)
	return j, nil
}

func (j *JWT) Verify() (*JWT, error) {
	if len(j.PublicKey) <= 0 {
		return nil, errors.New("public key is empty")
	}

	j.Token = strings.Replace(j.Token, "BASE ", "", -1)
	j.Token = strings.Replace(j.Token, "BEARER ", "", -1)
	j.Token = strings.Replace(j.Token, "base ", "", -1)
	j.Token = strings.Replace(j.Token, "bearer ", "", -1)
	j.Token = strings.Replace(j.Token, "Base ", "", -1)
	j.Token = strings.Replace(j.Token, "Bearer ", "", -1)
	key, _ := pem.Decode([]byte(j.PublicKey))
	public, _ := x509.ParsePKIXPublicKey(key.Bytes)
	publicKey := public.(*rsa.PublicKey)
	token, err := jwt.Parse([]byte(j.Token), jwt.WithValidate(true),
		jwt.WithKey(jwa.RS512, publicKey))
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
