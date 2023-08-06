package web

import (
	"github.com/gin-gonic/gin"
	"github.com/hillfolk/go-hexagonal-architecture/internal/user/application/port/in"
	"github.com/hillfolk/go-hexagonal-architecture/internal/user/application/service"
	"github.com/pkg/errors"
	"net/http"
)

type GetUserDetailController struct {
	getUserDetailUseCase in.GetUserQuery
}

func NewGetUserDetailController() *GetUserDetailController {
	return &GetUserDetailController{
		getUserDetailUseCase: service.NewGetUserService(),
	}
}

func (c *GetUserDetailController) GetUserDetail(ctx *gin.Context) {
	userId, has := ctx.GetQuery("user_id")
	if !has {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "not found user",
		})
	}

	user, err := c.getUserDetailUseCase.GetUser(userId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": errors.Wrap(err, "fail get user detail"),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"result": user,
		"status": "ok",
	})

}
