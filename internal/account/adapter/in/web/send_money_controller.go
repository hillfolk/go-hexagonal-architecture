package web

import (
	gin "github.com/gin-gonic/gin"
	"github.com/hillfolk/go-hexagonal-architecture/internal/account/adapter/in/web/request"
	"github.com/hillfolk/go-hexagonal-architecture/internal/account/application/port/in"
	"github.com/hillfolk/go-hexagonal-architecture/internal/account/application/service"
	account "github.com/hillfolk/go-hexagonal-architecture/internal/account/domain"
	"net/http"
)

type SendMoneyController struct {
	sendMoneyUseCase in.SendMoneyUseCase
}

func NewSendMoneyController() *SendMoneyController {
	return &SendMoneyController{
		sendMoneyUseCase: service.NewSendMoneyService(),
	}
}

func (a *SendMoneyController) SendMoney(ctx *gin.Context) {
	var request request.SendMoneyRequest
	if err := ctx.Bind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	// TODO: add validation

	if err := a.sendMoneyUseCase.SendMoney(
		in.SendMoneyCommand{
			SourceAccountId: account.AccountId(request.SourceAccountId),
			TargetAccountId: account.AccountId(request.TargetAccountId),
			Amount: account.Money{
				Amount: request.Amount,
			},
		},
	); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
