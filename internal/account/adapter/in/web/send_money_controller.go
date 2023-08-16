package web

import (
	gin "github.com/gin-gonic/gin"
	"github.com/hillfolk/go-hexagonal-architecture/internal/account/adapter/in/web/request"
	"github.com/hillfolk/go-hexagonal-architecture/internal/account/application/port/in"
	"github.com/hillfolk/go-hexagonal-architecture/internal/account/application/service"
	account "github.com/hillfolk/go-hexagonal-architecture/internal/account/domain"
	"github.com/hillfolk/go-hexagonal-architecture/utils"
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
	var sendMoneyRequest request.SendMoneyRequest
	if err := ctx.Bind(&sendMoneyRequest); err != nil {
		utils.RenderFail(ctx, http.StatusBadRequest, err)
		return
	}
	// TODO: add validation

	if err := a.sendMoneyUseCase.SendMoney(
		in.SendMoneyCommand{
			SourceAccountId: account.AccountId(sendMoneyRequest.SourceAccountId),
			TargetAccountId: account.AccountId(sendMoneyRequest.TargetAccountId),
			Amount: account.Money{
				Amount: sendMoneyRequest.Amount,
			},
		},
	); err != nil {
		utils.RenderFail(ctx, http.StatusInternalServerError, err)
	}
	utils.RenderSucc(ctx, http.StatusOK, "success")
}
