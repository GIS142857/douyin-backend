package routers

import (
	"douyin-backend/app/global/consts"
	"douyin-backend/app/global/variable"
	"douyin-backend/app/http/middleware/cors"
	validatorFactory "douyin-backend/app/http/validator/core/factory"
	"douyin-backend/app/utils/gin_release"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func InitWebRouter() *gin.Engine {
	var router *gin.Engine
	if variable.ConfigYml.GetBool("AppDebug") == false {
		router = gin_release.ReleaseRouter()
	} else {
		// 调试模式，开启 pprof 包，便于开发阶段分析程序性能
		router = gin.Default()
		pprof.Register(router)
	}
	// TODO study
	// 设置可信任的代理服务器列表,gin (2021-11-24发布的v1.7.7版本之后出的新功能)
	if variable.ConfigYml.GetInt("HTTPServer.TrustProxies.IsOpen") == 1 {
		if err := router.SetTrustedProxies(variable.ConfigYml.GetStringSlice("HttpServer.TrustProxies.ProxyServerList")); err != nil {
			variable.ZapLog.Error(consts.GinSetTrustProxyError, zap.Error(err))
		}
	} else {
		_ = router.SetTrustedProxies(nil)
	}

	// TODO study
	//根据配置进行设置跨域
	if variable.ConfigYml.GetBool("HttpServer.AllowCrossDomain") {
		router.Use(cors.Next())
	}

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "douyin-backend")
	})

	//处理静态文件
	router.Static("/public", "./public")

	user := router.Group("user/")
	{
		user.GET("userinfo", validatorFactory.Create(consts.ValidatorPrefix+"UserInfo"))
		user.GET("video_list", validatorFactory.Create(consts.ValidatorPrefix+"VideoList"))
		user.GET("panel", validatorFactory.Create(consts.ValidatorPrefix+"Panel"))
		user.GET("friends", validatorFactory.Create(consts.ValidatorPrefix+"Friends"))
		user.GET("collect", validatorFactory.Create(consts.ValidatorPrefix+"Collect"))
	}
	// TODO 查看 post 是否需要合并到其他路由组里面
	post := router.Group("post/")
	{
		post.GET("recommended", validatorFactory.Create(consts.ValidatorPrefix+"PostRecommended"))
	}
	shop := router.Group("shop/")
	{
		shop.GET("recommended", validatorFactory.Create(consts.ValidatorPrefix+"ShopRecommended"))
	}
	video := router.Group("video/")
	{

		video.GET("like", validatorFactory.Create(consts.ValidatorPrefix+"Like"))
		video.GET("comments", validatorFactory.Create(consts.ValidatorPrefix+"Comments"))
		video.GET("star", validatorFactory.Create(consts.ValidatorPrefix+"Star"))
		video.GET("share", validatorFactory.Create(consts.ValidatorPrefix+"Share"))
		video.GET("historyOther", validatorFactory.Create(consts.ValidatorPrefix+"HistoryOther"))
		video.GET("history", validatorFactory.Create(consts.ValidatorPrefix+"History"))
		video.GET("long/recommended/", validatorFactory.Create(consts.ValidatorPrefix+"LongRecommended"))
		video.GET("my", validatorFactory.Create(consts.ValidatorPrefix+"My"))
		video.GET("private", validatorFactory.Create(consts.ValidatorPrefix+"Private"))

	}
	return router
}
