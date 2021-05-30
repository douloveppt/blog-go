package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Tags Base
// @Summary 用户登录
// @Success 200 {string} string "{"code":0,"msg":"success","data":"login"}"
// @Router /v1/base/login [post]
func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg:": "success",
		"data": "login successfully",
	})
}
