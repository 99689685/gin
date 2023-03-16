package v1

import (
	"ginweb/model"
	"ginweb/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddCategory 添加分类
func AddCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		model.CreateCategory(&data)
	}
	if code == errmsg.ErrorCateNameUsed {
		code = errmsg.ErrorCateNameUsed
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// QueryCategory 查询分类列表
func QueryCategory(c *gin.Context) {
	size, _ := strconv.Atoi(c.Query("pageSize"))
	num, _ := strconv.Atoi(c.Query("pageNum"))
	code = errmsg.SUCCESS
	data, _ := model.GetCategory(size, num)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// EditCategory 修改分类名称
func EditCategory(c *gin.Context) {
	var cate model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&cate)
	code = model.CheckCategory(cate.Name)
	if code == errmsg.SUCCESS {
		model.EditCategory(id, &cate)
	}
	if code == errmsg.ErrorCateNameUsed {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})

}

// DeleteCategory 删除分类
func DeleteCategory(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	code = model.DeleteCategory(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
