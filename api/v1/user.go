package v1

import (
	"ginweb/model"
	"ginweb/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var code int

// AddUser 添加用户
func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
	}
	if code == errmsg.ErrorUsernameUsed {
		code = errmsg.ErrorUsernameUsed
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// QueryUsers 查询用户列表
func QueryUsers(c *gin.Context) {
	size, _ := strconv.Atoi(c.Query("pageSize"))
	num, _ := strconv.Atoi(c.Query("pageNum"))
	username := c.Query("username")
	code = errmsg.SUCCESS
	data, _ := model.GetUsers(username, size, num)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetUser 查询单个用户
func GetUser(c *gin.Context) {
	//id, _ := strconv.Atoi(c.Param("id"))
	id := c.Query("id")
	var maps = make(map[string]interface{})
	code = errmsg.SUCCESS
	data, code := model.FindUser(id)
	maps["username"] = data.Username
	maps["role"] = data.Role
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    maps,
		"message": errmsg.GetErrMsg(code),
	})

}

// EditUser 修改用户
func EditUser(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.EditUser(id, &data)
	}
	if code == errmsg.ErrorUsernameUsed {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})

}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	code = model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
