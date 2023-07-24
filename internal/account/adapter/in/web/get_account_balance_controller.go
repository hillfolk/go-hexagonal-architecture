package web

import (
	"github.com/gin-gonic/gin"
	"github.com/hillfolk/go-hexagonal-architecture/internal/account/adapter/in/web/request"
	"github.com/hillfolk/go-hexagonal-architecture/internal/account/application/port/in"
	"github.com/hillfolk/go-hexagonal-architecture/internal/account/application/service"
	"github.com/hillfolk/go-hexagonal-architecture/internal/account/domain"
	"net/http"
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
	var accountBalanceRequest request.GetAccountBalanceRequest
	if err := ctx.Bind(&accountBalanceRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	money, err := a.service.GetAccountBalance(domain.NewAccountId(accountBalanceRequest.AccountId))
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
