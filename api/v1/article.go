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

// GetCateArt 查询分类下的所有文章
func GetCateArt(c *gin.Context) {
	size, _ := strconv.Atoi(c.Query("pageSize"))
	num, _ := strconv.Atoi(c.Query("pageNum"))
	id, _ := strconv.Atoi(c.Query("id"))
	if size >= 100 {
		size = 100
	}
	if size <= 0 {
		size = 10
	}
	if num == 0 {
		num = 1
	}

	data, code := model.GetCateArt(id, size, num)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetArtInfo 查询单个文章信息
func GetArtInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetArtInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  data,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetArt 查询文章列表
func GetArt(c *gin.Context) {
	size, _ := strconv.Atoi(c.Query("pageSize"))
	num, _ := strconv.Atoi(c.Query("pageNum"))
	data, code := model.GetArt(size, num)
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
