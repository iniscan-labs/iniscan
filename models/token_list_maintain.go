package models

import "github.com/mss-boot-io/mss-boot/pkg/response/actions"

type TokenListMaintain struct {
	actions.ModelGorm
	ID              int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Description     string `json:"description" gorm:"not null"`
	ContractAddress string `json:"contract_address" gorm:"size:80;not null;index"`
	PriceMethod     string `json:"price_method" gorm:"not null"`
	ListPriority    int    `json:"list_priority" gorm:"not null"`
	Tag             string `json:"tag" gorm:"not null"`
	LogoPath        string `json:"logo_path" gorm:"size:128;not null"`
	Symbol          string `json:"symbol" gorm:"size:40;not null"`
}

func (TokenListMaintain) TableName() string {
	return "token_list_maintain"
}
