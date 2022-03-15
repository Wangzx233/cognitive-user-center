package api

import (
	"cognitive-user-center/logic"
	"fmt"
	"github.com/gin-gonic/gin"
)

// LoginByPassword 密码登陆
func LoginByPassword(c *gin.Context) {
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

	//验证电话和密码
	err = user.LoginByPassword()
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

// LoginByCode 验证码登陆
func LoginByCode(c *gin.Context) {
	var user logic.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"info": "parameter mismatch",
			"data": nil,
		})
		return
	}

	//验证电话格式
	err = logic.CheckPhoneNumber(user.PhoneNumber)
	if err != nil {
		c.JSON(400, gin.H{
			"info": fmt.Sprint(err),
			"data": nil,
		})
		return
	}

	//判断是否存在用户
	if !logic.IsRegister(user.PhoneNumber) {
		c.JSON(500, gin.H{
			"info": "phoneNumber don't exited",
			"data": nil,
		})
		return
	}

	err = user.LoginByCode()
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

// SendLoginCode 发送登陆验证码短信
func SendLoginCode(c *gin.Context) {
	phoneNumber := c.Query("phone_number")

	//正则验证电话号格式
	err := logic.CheckPhoneNumber(phoneNumber)
	if err != nil {
		c.JSON(500, gin.H{
			"info": fmt.Sprint(err),
		})
		return
	}

	//判断是否存在用户
	if !logic.IsRegister(phoneNumber) {
		c.JSON(500, gin.H{
			"info": "phoneNumber don't exited",
		})
		return
	}

	//发送验证码短信并保存到redis
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
