package models

import "github.com/mss-boot-io/mss-boot/pkg/response/actions"

type ContractVerifyJob struct {
	actions.ModelGorm
	ContractAddress string `json:"contract_address" gorm:"size:200"`
	Compiler        string `json:"compiler" gorm:"size:200"`
	StandardJson    string `json:"standard_json"`
	Status          int64  `json:"status" gorm:"comment:'0 unhandle 1 verified -1 failed'"`
	License         string `json:"license" gorm:"size:255"`
}

func (ContractVerifyJob) TableName() string {
	return "contract_verify_job"
}
