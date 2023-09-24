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
	auth := r.Group("api/v1")
	{
		// 用户模块口
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		// 分类模块接口
		auth.POST("category", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)
		// 文章模块接口
		auth.POST("article", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArt)
		auth.DELETE("article/:id", v1.DeleteArt)
	}
	router := r.Group("api/v1")
	{
		router.GET("user", v1.QueryUsers)
		router.GET("userDetails", v1.GetUser)
		router.GET("category", v1.QueryCategory)
		router.GET("singleArticle/:id", v1.GetArtInfo)
		router.GET("articleList", v1.GetCateArt)
		router.GET("article", v1.GetArt)
		router.POST("login", v1.LoginFront)
		router.POST("user", v1.AddUser)
		router.POST("addLeaver", v1.AddLeaver)
		router.GET("getDingUser", v1.GetDingUser)
	}
	r.Run(utils.HttpPort)
}
