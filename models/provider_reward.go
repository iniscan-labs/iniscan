package models

import "github.com/mss-boot-io/mss-boot/pkg/response/actions"

type ProviderReward struct {
	actions.ModelGorm
	BlockNumber  int64  `json:"block_number" gorm:"index"`
	BlockTime    int64  `json:"block_time"`
	TotalReward  int64  `json:"total_reward"`
	Reward       int64  `json:"reward"`
	RemainReward int64  `json:"remain_reward"`
	Address      string `json:"address" gorm:"size:80"`
	Days         int64  `json:"days"`
}

func (ProviderReward) TableName() string {
	return "provider_reward"
}
