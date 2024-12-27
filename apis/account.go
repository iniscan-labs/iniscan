package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iniscan-labs/iniscan/dto"
	"github.com/iniscan-labs/iniscan/models"
	"github.com/mss-boot-io/mss-boot/pkg/response"
	"github.com/mss-boot-io/mss-boot/pkg/response/actions"
	"github.com/mss-boot-io/mss-boot/pkg/response/controller"
)

func init() {
	e := &Account{
		Simple: controller.NewSimple(
			controller.WithAuth(true),
			controller.WithModel(new(models.Account)),
			controller.WithSearch(new(dto.AccountSearch)),
			controller.WithModelProvider(actions.ModelProviderGorm),
		),
	}
	response.AppendController(e)
}

type Account struct {
	*controller.Simple
}

func (e *Account) Other(r *gin.RouterGroup) {
	r.POST("/address/address_summary", e.AddressSummary)
}

// AddressSummary  account summary
// @Summary  account summary
// @Description  account summary
// @Tags account
// @Accept  application/json
// @Product application/json
// @param data body dto.AccountSearch true "data"
// @Success 200 {object} dto.AccountSummaryResponse
// @Router /explorer/api/address/address_summary [post]
// @Security Bearer
func (e *Account) AddressSummary(ctx *gin.Context) {
	api := response.Make(ctx)
	req := &dto.AccountSearch{}
	if api.Bind(req).Error != nil {
		api.Err(http.StatusUnprocessableEntity)
		return
	}
	resp := &dto.AccountSummaryResponse{
		Balance: "0",
		Count:   0,
	}
	// get balance
	accountBalance, err := models.GetAccountBalance(ctx, req)
	if err != nil {
		api.Err(http.StatusInternalServerError)
		return
	}
	resp.Balance = accountBalance.Value
	// get transactions
	transactions, err := models.GetTransactionList(ctx, &dto.TransactionSearch{
		To: req.Address,
	})
	if err != nil {
		api.Err(http.StatusInternalServerError)
		return
	}
	for _, transaction := range transactions {
		if transaction.From == req.Address {
			resp.Count++
		}
	}

	api.OK(resp)
}
