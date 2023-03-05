package routers

import (
	"ginweb/api/v1"
	"ginweb/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	// 初始化路由
	r := gin.Default()
	router := r.Group("api/v1")
	{
		// 用户模块口
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.QueryUsers)
		router.PUT("users:id", v1.EditUser)
		router.DELETE("users/:id ", v1.DeleteUser)
		// 分类模块接口

		// 文章模块接口
	}
	r.Run(utils.HttpPort)
}
