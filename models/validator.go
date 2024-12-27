package models

import "github.com/mss-boot-io/mss-boot/pkg/response/actions"

type Validator struct {
	actions.ModelGorm
	Address     string  `json:"address" gorm:"size:80;primaryKey"`
	VotePool    string  `json:"vote_pool" gorm:"size:80"`
	State       int     `json:"state" gorm:"default:0"`
	Active      int     `json:"active" gorm:"default:0"`
	TotalVote   float64 `json:"total_vote" gorm:"type:numeric(80,0);default:0"`
	Blockheight int64   `json:"blockheight"`
	TotalReward float64 `json:"total_reward" gorm:"type:numeric(80,0)"`
	VoteTime    int64   `json:"vote_time"`
}

func (Validator) TableName() string {
	return "validator"
}
