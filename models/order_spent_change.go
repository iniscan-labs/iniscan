package models

import "github.com/mss-boot-io/mss-boot/pkg/response/actions"

type OrderSpentChange struct {
	actions.ModelGorm
	ID      int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Address string `json:"address" gorm:"size:80;unique"`
}

func (OrderSpentChange) TableName() string {
	return "order_spent_change"
}
