package login

import (
	"errors"
	"os"

	"gorm.io/gorm"
	"vodka.app/internal/pkg/code"
	"vodka.app/internal/pkg/jwe"
	"vodka.app/internal/pkg/log"
	"vodka.app/internal/pkg/util"
	accountsModel "vodka.app/internal/v1/structure/accounts"
	jweModel "vodka.app/internal/v1/structure/jwe"
	loginsModel "vodka.app/internal/v1/structure/logins"
)

func (r *resolver) Web(input *loginsModel.Login) interface{} {
	acknowledge, accounts, err := r.Account.AcknowledgeAccount(&accountsModel.Field{
		Account:   util.PointerString(input.Account),
		Password:  util.PointerString(input.Password),
		CompanyID: util.PointerString(input.CompanyID),
	})
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	if acknowledge == false {
		return code.GetCodeMessage(code.PermissionDenied, "Incorrect account password")
	}

	token, err := r.JWE.Created(&jweModel.JWE{
		AccountID: accounts[0].AccountID,
		CompanyID: input.CompanyID,
		Name:      accounts[0].Name,
	})
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, token)

}

func (r *resolver) Refresh(input *jweModel.Refresh) interface{} {
	j := &jwe.JWT{
		PrivateKey: os.Getenv("privateKey"),
		Token:      input.RefreshToken,
	}

	if len(j.Token) == 0 {
		return code.GetCodeMessage(code.JWTRejected, "refresh_token is error")
	}

	j, err := j.Verify()
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.JWTRejected, "refresh_token is error")
	}

	account, err := r.Account.GetByID(&accountsModel.Field{
		AccountID: j.Other["account_id"].(string),
		IsDeleted: util.PointerBool(false),
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.JWTRejected, "refresh_token is error")
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	token, err := r.JWE.Created(&jweModel.JWE{
		AccountID: account.AccountID,
		CompanyID: account.CompanyID,
		Name:      account.Name,
	})
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	token.RefreshToken = input.RefreshToken
	return code.GetCodeMessage(code.Successful, token)
}
