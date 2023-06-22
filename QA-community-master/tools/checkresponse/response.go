package checkresponse

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResSuccess(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": message,
	})
}

func ResFail(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"status":  500,
		"message": message,
	})
}
