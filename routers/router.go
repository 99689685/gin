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
		router.GET("userDetails", v1.GetUser)
		router.POST("user", v1.AddUser)
		router.PUT("user/:id", v1.EditUser)
		router.DELETE("user/:id", v1.DeleteUser)
		// 分类模块接口
		router.GET("category", v1.QueryCategory)
		router.POST("category", v1.AddCategory)
		router.PUT("category/:id", v1.EditCategory)
		router.DELETE("category/:id", v1.DeleteCategory)
		// 文章模块接口
		router.GET("article", v1.GetArt)
		router.POST("article", v1.AddArticle)
		router.PUT("article/:id", v1.EditArt)
		router.DELETE("article/:id", v1.DeleteArt)
	}
	r.Run(utils.HttpPort)
}
