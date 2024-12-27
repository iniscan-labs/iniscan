package dto

import (
	"github.com/mss-boot-io/mss-boot/pkg/response/actions"
)

type AccountSearch struct {
	actions.Pagination `search:"inline"`
	// address
	Address string `json:"address" form:"address" search:"type:contains;column:address"`
	// contract
	Contract string `json:"contract" form:"contract" search:"type:contains;column:contract"`
}

type AccountSummaryResponse struct {
	Balance string `json:"balance"`
	Count   int64  `json:"count"`
}
