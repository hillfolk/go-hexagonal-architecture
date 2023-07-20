package web

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/hillfolk/go-hexagonal-architecture/internal/account/adapter/in/web"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{MaxAge: 60 * 60 * 24}) // expire in a day
	router.Use(sessions.Sessions("server-session", store))

	router.GET("/health", health)

	v1 := router.Group("/api/v1")
	{

		account := v1.Group("/accounts")
		{
			sendMoneyController := web.NewSendMoneyController()
			account.POST("/send", sendMoneyController.SendMoney)
		}
	}
	return router

}
