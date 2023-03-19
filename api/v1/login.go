package v1

import (
	"ginweb/middleware"
	"ginweb/model"
	"ginweb/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

//func Login(c *gin.Context) {
//	var data *model.User
//	c.ShouldBindJSON(&data)
//	var token string
//	var code int
//	code = model.CheckLogin(data)
//	if code == errmsg.SUCCESS {
//		token, code = middleware.SetToken(data.Username)
//		c.JSON(http.StatusOK, gin.H{
//			"status":  code,
//			"message": errmsg.GetErrMsg(code),
//			"token":   token,
//			"role":    data.Role,
//		})
//	}
//}

// LoginFront 前台登录
func LoginFront(c *gin.Context) {
	var data model.User
	c.ShouldBindJSON(&data)
	var token string
	var code int
	code = model.CheckLogin(data.Username, data.Password)
	if code == errmsg.SUCCESS {
		token, code = middleware.SetToken(data.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"user":    data.Username,
		"id":      data.ID,
		"token":   token,
		"message": errmsg.GetErrMsg(code),
	})
}
