package models

import "github.com/mss-boot-io/mss-boot/pkg/response/actions"

type BalanceChange struct {
	actions.ModelGorm
	Address      string `json:"address" gorm:"size:80"`
	Contract     string `json:"contract" gorm:"size:80"`
	ContractType int64  `json:"contract_type"`
	TokenID      string `json:"token_id" gorm:"size:80"`
}

func (BalanceChange) TableName() string {
	return "balance_change"
}
