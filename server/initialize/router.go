package initialize

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "myblog/docs"
	"myblog/global"
	"myblog/router"
	"net/http"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.StaticFS(global.GCONFIG.Local.Path, http.Dir(global.GCONFIG.Local.Path)) //用户上传图片的存放路径
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GLOG.Info("register swagger handler")
	PublicGroup := Router.Group("v1")
	{
		router.InitBaseRouter(PublicGroup)
	}
	global.GLOG.Info("register routers successfully")
	return Router
}
