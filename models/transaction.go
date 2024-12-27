package models

import (
	"github.com/gin-gonic/gin"
	"github.com/iniscan-labs/iniscan/dto"
	"github.com/mss-boot-io/mss-boot/pkg/config/gormdb"
	"github.com/mss-boot-io/mss-boot/pkg/response/actions"
)

type Transaction struct {
	actions.ModelGorm
	Hash             string `json:"hash" gorm:"size:80;index"`
	Status           int64  `json:"status"`
	ErrorInfo        string `json:"error_info"`
	BlockNumber      int64  `json:"block_number" gorm:"index"`
	BlockTime        int64  `json:"block_time" gorm:"index"`
	From             string `json:"from" gorm:"size:80;index"`
	To               string `json:"to" gorm:"size:80;index"`
	Value            string `json:"value" gorm:"type:numeric(80,0)"`
	Fee              string `json:"fee" gorm:"type:numeric(32,0)"`
	GasLimit         int64  `json:"gas_limit"`
	GasUsed          int64  `json:"gas_used"`
	GasPrice         int64  `json:"gas_price"`
	TransactionType  int64  `json:"transaction_type"`
	MaxPriority      int64  `json:"max_priority"`
	MaxFee           int64  `json:"max_fee"`
	Nonce            int64  `json:"nonce"`
	InputData        string `json:"input_data"`
	BlockHash        string `json:"block_hash" gorm:"size:80;index"`
	TransactionIndex int64  `json:"transaction_index" gorm:"index"`
}

func (Transaction) TableName() string {
	return "transaction"
}

// 获取交易信息
func GetTransaction(ctx *gin.Context, hash string) (Transaction, error) {
	transaction := Transaction{}
	db := gormdb.DB.Model(transaction)
	db = db.Where("hash = ?", hash)
	db = db.First(transaction)
	return transaction, db.Error
}

// 获取交易信息列表
func GetTransactionList(ctx *gin.Context, search *dto.TransactionSearch) ([]Transaction, error) {
	transactions := []Transaction{}
	db := gormdb.DB.Model(Transaction{})
	if search.Hash != "" {
		db = db.Where("hash = ?", search.Hash)
	}
	if search.From != "" {
		db = db.Where("from = ?", search.From)
	}
	if search.To != "" {
		db = db.Where("'to' = ?", search.To)
	}
	if search.BlockNumber != 0 {
		db = db.Where("block_number = ?", search.BlockNumber)
	}
	if search.BlockTime != 0 {
		db = db.Where("block_time = ?", search.BlockTime)
	}
	db = db.Order("block_time DESC")
	db = db.Find(&transactions)

	return transactions, db.Error
}
