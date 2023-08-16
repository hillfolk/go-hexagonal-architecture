package web

import (
	gin "github.com/gin-gonic/gin"
	"github.com/hillfolk/go-hexagonal-architecture/infra/api"
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
	var sendMoneyRequest request.SendMoneyRequest
	if err := ctx.Bind(&sendMoneyRequest); err != nil {
		api.RenderFail(ctx, http.StatusBadRequest, err)
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
		api.RenderFail(ctx, http.StatusInternalServerError, err)
	}
	api.RenderSucc(ctx, http.StatusOK, "success")
}
