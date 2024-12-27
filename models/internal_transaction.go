package models

import "github.com/mss-boot-io/mss-boot/pkg/response/actions"

type InternalTransaction struct {
	actions.ModelGorm
	BlockHash             string `json:"block_hash" gorm:"size:80;index"`
	BlockNumber           int64  `json:"block_number" gorm:"index"`
	ParentTransactionHash string `json:"parent_transaction_hash" gorm:"size:80"`
	From                  string `json:"from" gorm:"size:80;index"`
	To                    string `json:"to" gorm:"size:80;index"`
	Value                 string `json:"value" gorm:"type:numeric(80,0)"`
	TypeTraceAddress      string `json:"type_trace_address" gorm:"size:80"`
	GasLimit              int64  `json:"gas_limit"`
	Op                    string `json:"op" gorm:"size:40"`
}

func (InternalTransaction) TableName() string {
	return "internal_transaction"
}
