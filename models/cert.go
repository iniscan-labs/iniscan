package models

import "github.com/mss-boot-io/mss-boot/pkg/response/actions"

type Cert struct {
	actions.ModelGorm
	ID        int    `json:"id" gorm:"primaryKey;autoIncrement"`
	SeriNo    string `json:"seri_no" gorm:"size:80"`
	BlockTime int64  `json:"block_time"`
	Expire    int64  `json:"expire"`
	Owner     string `json:"owner" gorm:"size:255"`
	CertIndex int    `json:"cert_index"`
	Raw       string `json:"raw"`
	State     int16  `json:"state"`
}

func (Cert) TableName() string {
	return "cert"
}
