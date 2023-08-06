package web

import (
	"github.com/gin-gonic/gin"
	"github.com/hillfolk/go-hexagonal-architecture/internal/account/application/port/in"
	"github.com/hillfolk/go-hexagonal-architecture/internal/account/application/service"
	"github.com/hillfolk/go-hexagonal-architecture/internal/account/domain"
	"net/http"
	"strconv"
)

type GetAccountBalanceController struct {
	service in.GetAccountBalanceQuery
}

func NewGetAccountBalanceController() *GetAccountBalanceController {

	return &GetAccountBalanceController{
		service: service.NewGetAccountBalanceService(),
	}
}

func (a *GetAccountBalanceController) GetAccountBalance(ctx *gin.Context) {
	id, _ := ctx.Params.Get("accountId")
	accountId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "accountId must be a number",
		})
	}

	money, err := a.service.GetAccountBalance(domain.NewAccountId(accountId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"balance": money.Amount,
		"status":  "ok",
	})
}
