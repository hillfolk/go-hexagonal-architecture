package web

import (
	"github.com/gin-gonic/gin"
	"github.com/hillfolk/go-hexagonal-architecture/internal/user/adapter/in/web/response"
	"github.com/hillfolk/go-hexagonal-architecture/internal/user/application/port/in"
	"github.com/hillfolk/go-hexagonal-architecture/internal/user/application/service"
	"github.com/hillfolk/go-hexagonal-architecture/internal/user/domain"
	"github.com/samber/lo"
	"net/http"
)

type GetUserListController struct {
	getUserListUseCase in.GetUserListQuery
}

func NewGetUserListController() *GetUserListController {
	return &GetUserListController{
		getUserListUseCase: service.NewGetUserService(),
	}
}

func (c *GetUserListController) GetUserList(context *gin.Context) {
	users, err := c.getUserListUseCase.GetUserList()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	items := lo.Map(users, func(user *domain.User, idx int) *response.UserResponse {
		return &response.UserResponse{
			UserId: string(user.UserId),
			Email:  user.Email,
		}
	})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"items":  items,
		"status": "ok",
	})
}
