package models

import "github.com/mss-boot-io/mss-boot/pkg/response/actions"

type Provider struct {
	actions.ModelGorm
	ID           int     `json:"id" gorm:"primaryKey;autoIncrement"`
	Address      string  `json:"address" gorm:"size:255;unique"`
	Owner        string  `json:"owner" gorm:"size:255"`
	Uri          string  `json:"uri" gorm:"size:255"`
	ActiveLease  int     `json:"active_lease"`
	Cpu          int64   `json:"cpu"`
	MaxCpu       int64   `json:"max_cpu"`
	Memory       int64   `json:"memory"`
	MaxMemory    int64   `json:"max_memory"`
	Storage      int64   `json:"storage"`
	MaxStorage   int64   `json:"max_storage"`
	State        int16   `json:"state"`
	StakedAmount float64 `json:"staked_amount" gorm:"type:numeric(80,0)"`
	Country      string  `json:"country" gorm:"size:50"`
}

func (Provider) TableName() string {
	return "provider"
}
