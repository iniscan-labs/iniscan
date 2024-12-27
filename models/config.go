package models

import "github.com/mss-boot-io/mss-boot/pkg/response/actions"

type Config struct {
	actions.ModelGorm
	Key   string `json:"key" gorm:"size:255;unique"`
	Value string `json:"value"`
}

func (Config) TableName() string {
	return "config"
}
