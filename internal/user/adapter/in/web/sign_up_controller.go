package web

import (
	"github.com/gin-gonic/gin"
	"github.com/hillfolk/go-hexagonal-architecture/internal/account/adapter/in/web/request"
	"github.com/hillfolk/go-hexagonal-architecture/internal/user/application/port/in"
	"net/http"
)

type SignUpController struct {
	signUpUseCase in.SignUpCommand
}

func NewSignUpController() *SignUpController {
	return &SignUpController{}
}

func (c *SignUpController) SignUp(ctx *gin.Context) {
	var request request.SendMoneyRequest
	if err := ctx.Bind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
