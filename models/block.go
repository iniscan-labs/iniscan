package models

import "github.com/mss-boot-io/mss-boot/pkg/response/actions"

type Block struct {
	actions.ModelGorm
	BlockHash                string  `json:"block_hash" gorm:"size:80;unique"`
	BlockNumber              int64   `json:"block_number" gorm:"index"`
	TransactionCount         int64   `json:"transaction_count"`
	InternalTransactionCount int64   `json:"internal_transaction_count"`
	Validator                string  `json:"validator" gorm:"size:80"`
	BlockReward              float64 `json:"block_reward" gorm:"type:numeric(80,0)"`
	Difficult                int64   `json:"difficult"`
	TotalDifficult           int64   `json:"total_difficult"`
	BlockSize                int64   `json:"block_size"`
	GasUsed                  int64   `json:"gas_used"`
	GasLimit                 int64   `json:"gas_limit"`
	ExtraData                string  `json:"extra_data"`
	ParentHash               string  `json:"parent_hash" gorm:"size:80"`
	Sha3Uncle                string  `json:"sha3_uncle" gorm:"size:80"`
	Nonce                    string  `json:"nonce" gorm:"size:80"`
	Uncles                   string  `json:"uncles"`
	BlockTime                int64   `json:"block_time"`
	Provider                 string  `json:"provider" gorm:"size:80"`
	Team                     string  `json:"team" gorm:"size:80"`
	ValidatorRate            int64   `json:"validator_rate"`
	TeamRate                 int64   `json:"team_rate"`
}

func (Block) TableName() string {
	return "block"
}
