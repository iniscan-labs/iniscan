package models

import "github.com/mss-boot-io/mss-boot/pkg/response/actions"

type FunctionInfo struct {
	actions.ModelGorm
	MethodID          string `json:"method_id" gorm:"size:255;unique"`
	FunctionName      string `json:"function_name"`
	FunctionSignature string `json:"function_signature" gorm:"size:255"`
	FunctionAbi       string `json:"function_abi"`
	UpdateTime        int64  `json:"update_time" gorm:"index"`
	MethodName        string `json:"method_name" gorm:"size:255"`
}

func (FunctionInfo) TableName() string {
	return "function_info"
}
