package api

import (
	"cognitive-user-center/dao"
	"cognitive-user-center/logic"
	"github.com/gin-gonic/gin"
)

// RefreshToken 根据refresh_token刷新access_token和refresh_token
func RefreshToken(c *gin.Context) {
	oldRt := c.Query("refresh_token")
	phoneNumber, err := dao.GetPhoneNumberByRT(oldRt)
	if err != nil {
		c.JSON(500, gin.H{
			"info": "err",
			"data": nil,
		})
		return
	}

	//生成新的access_token
	accessToken, err := logic.GenAccessToken(phoneNumber)
	if err != nil {
		c.JSON(500, gin.H{
			"info": "err",
			"data": nil,
		})
		return
	}

	//生成新的refresh_token
	rt, err := logic.GenRefreshToken(phoneNumber)
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
