package middleware

import "github.com/gin-gonic/gin"

func Verify(c *gin.Context) {
	phoneNumber := c.GetHeader("phone_number")
	if phoneNumber == "" {
		c.JSON(402, gin.H{
			"info": "access_token wrong",
		})
		c.Abort()
	} else {
		c.Set("phone_number", phoneNumber)
		c.Next()
	}
}
