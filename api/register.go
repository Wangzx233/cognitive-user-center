package api

import (
	"cognitive-user-center/logic"
	"fmt"
	"github.com/gin-gonic/gin"
)

// Register 注册接口
func Register(c *gin.Context) {
	var user logic.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"info": "parameter mismatch",
			"data": nil,
		})
		return
	}

	//验证用户类型格式
	err = user.Check()
	if err != nil {
		c.JSON(500, gin.H{
			"info": fmt.Sprint(err),
			"data": nil,
		})
		return
	}

	//验证验证码并加密保存用户
	err = user.Register()
	if err != nil {
		c.JSON(500, gin.H{
			"info": fmt.Sprint(err),
			"data": nil,
		})
		return
	}

	//生成access_token
	accessToken, err := logic.GenAccessToken(user.PhoneNumber)
	if err != nil {
		c.JSON(500, gin.H{
			"info": "err",
			"data": nil,
		})
		return
	}

	//生成refresh_token
	rt, err := logic.GenRefreshToken(user.PhoneNumber)
	if err != nil {
		c.JSON(500, gin.H{
			"info": "err",
			"data": nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"info": "success",
		"data": gin.H{
			"access_token":  accessToken,
			"refresh_token": rt,
		},
	})
}

// SendRegisterCode 发送注册验证码
func SendRegisterCode(c *gin.Context) {
	phoneNumber := c.Query("phone_number")

	//正则验证电话号格式
	err := logic.CheckPhoneNumber(phoneNumber)
	if err != nil {
		c.JSON(500, gin.H{
			"info": fmt.Sprint(err),
		})
		return
	}

	//判断是否已经登陆
	if logic.IsRegister(phoneNumber) {
		c.JSON(500, gin.H{
			"info": "phoneNumber exited",
		})
		return
	}

	//发送验证码短信并保存验证码到redis来准备验证
	err = logic.SendRegisterCode(phoneNumber)
	if err != nil {
		c.JSON(500, gin.H{
			"info": "err",
		})
		return
	}

	c.JSON(200, gin.H{
		"info": "success",
	})
}
