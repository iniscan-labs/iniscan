package models

import "github.com/mss-boot-io/mss-boot/pkg/response/actions"

type Contract struct {
	actions.ModelGorm
	ContractAddress     string `json:"contract_address" gorm:"size:80;not null;unique"`
	CreateTxHash        string `json:"create_tx_hash" gorm:"size:80"`
	Symbol              string `json:"symbol" gorm:"size:80"`
	ContractType        int64  `json:"contract_type"`
	Creator             string `json:"creator" gorm:"size:80"`
	Decimals            int64  `json:"decimals"`
	TotalSupply         string `json:"total_supply" gorm:"type:numeric(80,0)"`
	Desc                string `json:"desc"`
	ByteCode            string `json:"byte_code"`
	ContractABI         string `json:"contract_abi"`
	Name                string `json:"name" gorm:"size:80"`
	StandardJson        string `json:"standard_json"`
	Compiler            string `json:"compiler" gorm:"size:200"`
	Status              int64  `json:"status"`
	License             string `json:"license" gorm:"size:255"`
	ByteCodeStatus      bool   `json:"byte_code_status" gorm:"default:false"`
	ContractName        string `json:"contract_name" gorm:"size:255"`
	OptimizationEnabled string `json:"optimization_enabled" gorm:"size:255"`
	EvmVersion          string `json:"evm_version" gorm:"size:255"`
}

func (Contract) TableName() string {
	return "contract"
}
