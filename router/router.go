package router

import (
	"cognitive-user-center/api"
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

	user.POST("/information")

	r.Run(":8083")
}
