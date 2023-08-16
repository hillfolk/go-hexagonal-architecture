package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	accountcontroller "github.com/hillfolk/go-hexagonal-architecture/internal/account/adapter/in/web"
	usercontroller "github.com/hillfolk/go-hexagonal-architecture/internal/user/adapter/in/web"
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
			sendMoneyController := accountcontroller.NewSendMoneyController()
			getAccountBalanceController := accountcontroller.NewGetAccountBalanceController()
			account.POST("/send", sendMoneyController.SendMoney)
			account.GET("/:accountId/balance", getAccountBalanceController.GetAccountBalance)
		}
		user := v1.Group("/users")
		signUpController := usercontroller.NewSignUpController()
		signInController := usercontroller.NewSignInController()
		{
			user.POST("/sign-up", signUpController.SignUp)
			user.POST("/sign-in", signInController.SignIn)
		}

	}
	return router

}
