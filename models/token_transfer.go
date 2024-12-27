package models

import "github.com/mss-boot-io/mss-boot/pkg/response/actions"

type TokenTransfer struct {
	actions.ModelGorm
	TransactionHash string `json:"transaction_hash" gorm:"size:80"`
	LogIndex        int64  `json:"log_index"`
	Contract        string `json:"contract" gorm:"size:80;index"`
	TokenType       int64  `json:"token_type" gorm:"index;comment:'1:erc20 2:erc721 3: erc1155'"`
	Value           string `json:"value" gorm:"type:numeric(80,0)"`
	From            string `json:"from" gorm:"size:80;index"`
	To              string `json:"to" gorm:"size:80;index"`
	MethodID        string `json:"method_id" gorm:"size:80"`
	BlockHash       string `json:"block_hash" gorm:"size:80;index"`
	BlockTime       int64  `json:"block_time" gorm:"index"`
	TokenID         string `json:"token_id" gorm:"size:80"`
}

func (TokenTransfer) TableName() string {
	return "token_transfer"
}
