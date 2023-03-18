package v1

import (
	"ginweb/model"
	"ginweb/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddArticle 添加文章

func AddArticle(c *gin.Context) {
	var data *model.Article
	_ = c.ShouldBindJSON(&data)
	code := model.CreateArt(data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询分类下的所有文章

// 查询单个文章信息

// GetArt 查询文章列表
func GetArt(c *gin.Context) {
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

// EditArt 修改文章
func EditArt(c *gin.Context) {
	var data *model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code := model.EditArt(id, data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})

}

// DeleteArt 删除文章
func DeleteArt(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteArt(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
