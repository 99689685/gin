package v1

import (
	"fmt"
	"ginweb/model"
	"ginweb/utils"
	"ginweb/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDingUser(c *gin.Context) {
	accessToken, err := model.GetDingAccessToken(utils.AppKey, utils.AppSecret)
	if err != nil {
		fmt.Printf("获取 Dingding ACCESS_TOKEN 失败：%s\n", errmsg.ErrDingUser)
		return
	}

	fmt.Printf("ACCESS_TOKEN: %s\n", accessToken)
	c.JSON(http.StatusOK, gin.H{
		"accessToken": accessToken,
		"message":     "DingDing token 获取成功",
	})
}
