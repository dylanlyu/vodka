package account

import (
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"vodka.app/internal/pkg/code"
	"vodka.app/internal/pkg/log"
	"vodka.app/internal/pkg/util"
	accountModel "vodka.app/internal/v1/structure/accounts"
	companyModel "vodka.app/internal/v1/structure/companies"
)

func (r *resolver) Created(trx *gorm.DB, input *accountModel.Created) interface{} {
	defer trx.Rollback()
	//Todo 角色名稱
	_, err := r.Company.WithTrx(trx).GetByID(&companyModel.Field{CompanyID: input.CompanyID,
		IsDeleted: util.PointerBool(false)})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	account, err := r.Account.WithTrx(trx).Created(input)
	if err != nil {
		if err.Error() == "account already exists" {
			return code.GetCodeMessage(code.BadRequest, err.Error())
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, account.AccountID)
}

func (r *resolver) List(input *accountModel.Fields) interface{} {
	input.IsDeleted = util.PointerBool(false)
	output := &accountModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, accounts, err := r.Account.List(input)
	accountsByte, err := json.Marshal(accounts)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(accountsByte, &output.Accounts)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *accountModel.Field) interface{} {
	input.IsDeleted = util.PointerBool(false)
	account, err := r.Account.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &accountModel.Single{}
	accountByte, _ := json.Marshal(account)
	err = json.Unmarshal(accountByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) Deleted(input *accountModel.Updated) interface{} {
	_, err := r.Account.GetByID(&accountModel.Field{AccountID: input.AccountID,
		IsDeleted: util.PointerBool(false)})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.Account.Deleted(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *accountModel.Updated) interface{} {
	account, err := r.Account.GetByID(&accountModel.Field{AccountID: input.AccountID,
		IsDeleted: util.PointerBool(false)})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.Account.Updated(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, account.AccountID)
}
