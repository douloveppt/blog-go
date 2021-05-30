package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg:": "success",
		"data": "login successfully",
	})
}
