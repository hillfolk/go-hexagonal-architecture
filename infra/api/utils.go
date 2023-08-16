package api

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type ApiResult struct {
	Success bool        `json:"success"`
	Result  interface{} `json:"result"`
	Error   error       `json:"error"`
}

func RenderFail(c *gin.Context, status int, err error) {
	if ok := errors.As(err, &err); ok {
		c.JSON(status, ApiResult{
			Success: false,
			Error:   err,
		})
	}
}

func RenderSucc(c *gin.Context, status int, result interface{}) {
	c.JSON(status, ApiResult{
		Success: true,
		Result:  result,
	})
}
