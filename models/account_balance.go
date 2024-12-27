package models

import (
	"github.com/gin-gonic/gin"
	"github.com/iniscan-labs/iniscan/dto"
	"github.com/mss-boot-io/mss-boot/pkg/config/gormdb"
	"github.com/mss-boot-io/mss-boot/pkg/response/actions"
)

type AccountBalance struct {
	actions.ModelGorm
	Address     string `json:"address" gorm:"size:80;uniqueIndex:uidx_address_contract_tokenid"`
	Contract    string `json:"contract" gorm:"size:80;uniqueIndex:uidx_address_contract_tokenid"`
	TokenID     string `json:"token_id" gorm:"size:80;uniqueIndex:uidx_address_contract_tokenid"`
	Value       string `json:"value" gorm:"type:numeric(80,0)"`
	LockBalance string `json:"lock_balance" gorm:"type:numeric(80,0)"`
}

func (AccountBalance) TableName() string {
	return "account_balance"
}

func GetAccountBalance(ctx *gin.Context, search *dto.AccountSearch) (*AccountBalance, error) {
	accountBalance := &AccountBalance{}
	db := gormdb.DB.Model(accountBalance)
	if search.Address != "" {
		db = db.Where("address = ?", search.Address)
	}
	if search.Contract != "" {
		db = db.Where("contract = ?", search.Contract)
	}
	db = db.First(accountBalance)
	return accountBalance, db.Error
}
