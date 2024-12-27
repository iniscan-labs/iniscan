package models

import "github.com/mss-boot-io/mss-boot/pkg/response/actions"

type TokenPriceHistory struct {
	actions.ModelGorm
	ID              int     `json:"id" gorm:"primaryKey;autoIncrement"`
	ContractAddress string  `json:"contract_address" gorm:"size:80;not null"`
	Price           float64 `json:"price" gorm:"type:numeric(80,0);default:NULL"`
	Timestamp       int     `json:"timestamp"`
}

func (TokenPriceHistory) TableName() string {
	return "token_price_history"
}
