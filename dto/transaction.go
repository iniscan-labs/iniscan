package dto

type TransactionSearch struct {
	Hash        string `json:"hash" form:"hash"`
	From        string `json:"from" form:"from"`
	To          string `json:"to" form:"to"`
	BlockNumber int64  `json:"block_number" form:"block_number"`
	BlockTime   int64  `json:"block_time" form:"block_time"`
}
