package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func UserFunc() func(c *gin.Context) {
	return func(c *gin.Context) {
		fmt.Println("Executing UserFunc")
		c.JSON(200, gin.H{"message": "UserFunc executed"})
	}
}
