package models

import "github.com/mss-boot-io/mss-boot/pkg/response/actions"

type EventInfo struct {
	actions.ModelGorm
	EventName     string `json:"event_name" gorm:"size:255"`
	EventAbi      string `json:"event_abi"`
	EventTopic    string `json:"event_topic" gorm:"size:255;primaryKey"`
	EventFullName string `json:"event_full_name"`
	UpdateTime    int64  `json:"update_time" gorm:"index"`
}

func (EventInfo) TableName() string {
	return "event_info"
}
