package models

import "github.com/mss-boot-io/mss-boot/pkg/response/actions"

type VoteDetail struct {
	actions.ModelGorm
	UserAddress     string  `json:"user_address" gorm:"size:80"`
	Contract        string  `json:"contract" gorm:"size:80"`
	Value           float64 `json:"value" gorm:"type:numeric(80,0)"`
	TransactionHash string  `json:"transaction_hash" gorm:"size:80"`
	LogIndex        int     `json:"log_index"`
}

func (VoteDetail) TableName() string {
	return "vote_detail"
}
