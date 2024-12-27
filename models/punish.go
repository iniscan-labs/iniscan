package models

import "github.com/mss-boot-io/mss-boot/pkg/response/actions"

type Punish struct {
	actions.ModelGorm
	PunishID     int     `json:"punish_id" gorm:"primaryKey;autoIncrement"`
	Address      string  `json:"address" gorm:"size:80"`
	PunishAmount float64 `json:"punish_amount" gorm:"type:numeric(80,0)"`
	LeftAmount   float64 `json:"left_amount" gorm:"type:numeric(80,0)"`
	PunishTime   int64   `json:"punish_time"`
	Type         string  `json:"type" gorm:"size:255;comment:'0 provider 1 validator'"`
}

func (Punish) TableName() string {
	return "punish"
}
