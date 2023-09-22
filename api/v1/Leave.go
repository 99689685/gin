package v1

import (
	"ginweb/model"
	"ginweb/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddLeaver(c *gin.Context) {
	var data *model.Leaver
	_ = c.ShouldBindJSON(&data)
	code := model.CreateLeaver(data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}
