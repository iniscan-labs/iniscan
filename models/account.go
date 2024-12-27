package models

import "github.com/mss-boot-io/mss-boot/pkg/response/actions"

type Account struct {
	actions.ModelGorm
	ID         int64  `json:"id" gorm:"primaryKey"`
	Address    string `json:"address" gorm:"size:80;unique"`
	Tag        string `json:"tag" gorm:"size:80"`
	TxnCount   int64  `json:"txn_count"`
	UpdateTime int64  `json:"update_time" gorm:"not null;default:0"`
	Status     int16  `json:"status" gorm:"not null;default:0"`
}

func (*Account) TableName() string {
	return "account"
}
