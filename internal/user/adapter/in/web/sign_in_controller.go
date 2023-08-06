package web

import (
	"github.com/gin-gonic/gin"
	"github.com/hillfolk/go-hexagonal-architecture/internal/user/adapter/in/web/request"
	"github.com/hillfolk/go-hexagonal-architecture/internal/user/application/port/in"
	"net/http"
)

type SignInController struct {
	signInUseCase in.SignInCommand
}

func NewSignInController() *SignInController {
	return &SignInController{}
}

func (c *SignInController) SignIn(ctx *gin.Context) {
	var signInRequest request.SignInRequest
	if err := ctx.Bind(&signInRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
