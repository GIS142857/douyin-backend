package main

import (
	"douyin-backend/app/global/variable" // 项目编译之前加载全局变量
	_ "douyin-backend/bootstrap"         // 项目初始化
	"douyin-backend/routers"
)

// 后端路由启动入口
func main() {
	router := routers.InitWebRouter()
	//fmt.Println(router.RouterGroup.Handlers)
	_ = router.Run(variable.ConfigYml.GetString("HttpServer.Web.Port"))
}
