package models

import "github.com/mss-boot-io/mss-boot/pkg/response/actions"

type TransactionLogs struct {
	actions.ModelGorm
	TransactionHash  string `json:"transaction_hash" gorm:"size:80"`
	LogIndex         int64  `json:"log_index"`
	Address          string `json:"address" gorm:"size:80"`
	BlockHash        string `json:"block_hash" gorm:"size:80;index"`
	BlockNumber      int64  `json:"block_number"`
	Data             string `json:"data"`
	Removed          bool   `json:"removed"`
	Topics           string `json:"topics"`
	TransactionIndex int64  `json:"transaction_index"`
}

func (TransactionLogs) TableName() string {
	return "transaction_logs"
}
