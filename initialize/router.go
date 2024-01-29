package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"uuid_srv/api"
	"uuid_srv/middleware"
)

// 初始化总路由

func Routers() *gin.Engine {
	Router := gin.Default()

	// 跨域，如需跨域可以打开下面的注释
	Router.Use(middleware.Cors()) // @tips 直接放行全部跨域请求

	// 注册路由
	api.Routers(Router)

	fmt.Println("router register success")
	return Router
}
