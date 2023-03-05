package routers

import (
	"ginweb/api/v1"
	"ginweb/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.DisableConsoleColor() // 禁用控制台颜色
	gin.SetMode(utils.AppMode)
	// 初始化路由
	r := gin.Default()
	router := r.Group("api/v1")
	{
		// 用户模块口
		router.GET("user", v1.QueryUsers)
		router.POST("user", v1.AddUser)
		router.PUT("user/:id", v1.EditUser)
		router.DELETE("user/:id", v1.DeleteUser)
		// 分类模块接口

		// 文章模块接口
	}
	r.Run(utils.HttpPort)
}
