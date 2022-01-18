package jwe

import (
	"os"
	"time"

	"vodka.app/internal/pkg/jwe"
	"vodka.app/internal/pkg/log"
	"vodka.app/internal/pkg/util"
	model "vodka.app/internal/v1/structure/jwe"
)

func (s service) Created(input *model.JWE) (output *model.Token, err error) {
	// TODO implement me
	other := map[string]interface{}{
		"company_id": input.CompanyID,
		"account_id": input.AccountID,
		"name":       input.Name,
	}

	accessExpiration := util.NowToUTC().Add(time.Minute * 5).Unix()
	j := &jwe.JWT{
		PublicKey:     os.Getenv("publicKey"),
		Other:         other,
		ExpirationKey: accessExpiration,
	}

	j, err = j.Created()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	accessToken := j.Token
	refreshTokenExpiration := util.NowToUTC().Add(time.Hour * 8).Unix()
	j.ExpirationKey = refreshTokenExpiration
	j, err = j.Created()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	refreshToken := j.Token
	output = &model.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return output, nil
}
