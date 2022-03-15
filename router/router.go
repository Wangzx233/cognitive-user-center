package router

import (
	"cognitive-user-center/api"
	"cognitive-user-center/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()

	user := r.Group("/user")

	user.POST("/register", api.Register)
	user.GET("/register/code", api.SendRegisterCode)

	user.POST("/login/password", api.LoginByPassword)
	user.GET("/login/code", api.SendLoginCode)
	user.POST("/login/code", api.LoginByCode)

	user.GET("/token/access", api.RefreshToken)

	user.Use(middleware.Verify)
	user.POST("/information", api.UpdateInfo)
	user.PUT("/information", api.UpdateInfo)
	user.GET("/center/index", api.FindUserInfo)

	r.Run(":8083")
}
