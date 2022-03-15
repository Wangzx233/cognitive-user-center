package api

import (
	"cognitive-user-center/logic"
	"fmt"
	"github.com/gin-gonic/gin"
)

// UpdateInfo 完善用户信息和更新用户信息公用接口
func UpdateInfo(c *gin.Context) {
	var userInfo logic.UserInfo
	err := c.BindJSON(&userInfo)
	if err != nil {
		c.JSON(400, gin.H{
			"info": "parameter mismatch",
		})
		return
	}

	phoneNumber, _ := c.Get("phone_number")
	userInfo.PhoneNumber = phoneNumber.(string)

	// 保存用户信息
	err = userInfo.SaveInfo()
	if err != nil {
		c.JSON(500, gin.H{
			"info": fmt.Sprint(err),
		})
		return
	}

	c.JSON(200, gin.H{
		"info": "success",
	})
}

// FindUserInfo 通过电话号获取用户信息
func FindUserInfo(c *gin.Context) {
	phoneNumber, _ := c.Get("phone_number")

	var UserInfo logic.UserInfo
	UserInfo.PhoneNumber = phoneNumber.(string)
	err := UserInfo.FindInfo()
	if err != nil {
		c.JSON(500, gin.H{
			"info": fmt.Sprint(err),
			"data": UserInfo,
		})
		return
	}

	c.JSON(200, gin.H{
		"info": "success",
		"data": UserInfo,
	})
}
