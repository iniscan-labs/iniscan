package models

import "github.com/mss-boot-io/mss-boot/pkg/response/actions"

type TokenListStat struct {
	actions.ModelGorm
	ID                  int     `json:"id" gorm:"primaryKey;autoIncrement"`
	Price               float64 `json:"price" gorm:"type:numeric(80,0);not null"`
	Holders             int     `json:"holders"`
	TransferCountHour   int     `json:"transfer_count_hour"`
	TransferCount1Day   int     `json:"transfer_count_1_day"`
	TransferCount2Day   int     `json:"transfer_count_2_day"`
	TransferCount7Day   int     `json:"transfer_count_7_day"`
	TransferCount30Day  int     `json:"transfer_count_30_day"`
	TransferAmountHour  float64 `json:"transfer_amount_hour" gorm:"type:numeric(80,0)"`
	TransferAmount1Day  float64 `json:"transfer_amount_1_day" gorm:"type:numeric(80,0)"`
	TransferAmount2Day  float64 `json:"transfer_amount_2_day" gorm:"type:numeric(80,0)"`
	TransferAmount7Day  float64 `json:"transfer_amount_7_day" gorm:"type:numeric(80,0)"`
	TransferAmount30Day float64 `json:"transfer_amount_30_day" gorm:"type:numeric(80,0)"`
	UniqueSenderHour    int     `json:"unique_sender_hour"`
	UniqueSender1Day    int     `json:"unique_sender_1_day"`
	UniqueSender2Day    int     `json:"unique_sender_2_day"`
	UniqueSender7Day    int     `json:"unique_sender_7_day"`
	UniqueSender30Day   int     `json:"unique_sender_30_day"`
	TotalSupply         int     `json:"total_supply"`
	UniqueReceiverHour  int     `json:"unique_receiver_hour"`
	UniqueReceiver1Day  int     `json:"unique_receiver_1_day"`
	UniqueReceiver2Day  int     `json:"unique_receiver_2_day"`
	UniqueReceiver7Day  int     `json:"unique_receiver_7_day"`
	UniqueReceiver30Day int     `json:"unique_receiver_30_day"`
	Contract            string  `json:"contract" gorm:"size:80;not null"`
}

func (TokenListStat) TableName() string {
	return "token_list_stat"
}
