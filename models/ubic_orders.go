package models

import "github.com/mss-boot-io/mss-boot/pkg/response/actions"

type UbicOrders struct {
	actions.ModelGorm
	State           int     `json:"state"`
	FinalPrice      string  `json:"final_price" gorm:"type:numeric(255,0)"`
	Balance         string  `json:"balance" gorm:"type:numeric(255,0)"`
	BlockTime       int64   `json:"block_time"`
	Spent           string  `json:"spent" gorm:"type:numeric(255,0)"`
	Provider        string  `json:"provider" gorm:"size:80"`
	Cpu             string  `gorm:"type:numeric(255,0)"`
	Memory          string  `json:"memory" gorm:"type:numeric(255,0)"`
	Storage         string  `json:"storage" gorm:"type:numeric(255,0)"`
	Owner           string  `json:"owner" gorm:"size:80"`
	Cert            float64 `json:"cert" gorm:"type:numeric(255,0)"`
	SdlTrxHash      string  `json:"sdl_trx_hash" gorm:"size:255"`
	Address         string  `json:"address" gorm:"size:80"`
	OrderID         int64   `json:"order_id" gorm:"primaryKey"`
	BlockNumber     string  `json:"block_number" gorm:"size:64"`
	TransactionHash string  `json:"transaction_hash" gorm:"size:255"`
	BlockHash       string  `json:"block_hash" gorm:"size:255"`
}

func (UbicOrders) TableName() string {
	return "ubic_orders"
}
