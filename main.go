package main

import (
	"ginweb/model"
	"ginweb/routers"
)

func main() {
	//// 引用数据库
	model.InitDb()
	//// 路由地址
	routers.InitRouter()
}
