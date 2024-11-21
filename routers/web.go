package routers

import (
	"douyin-backend/app/global/consts"
	"douyin-backend/app/global/variable"
	"douyin-backend/app/http/middleware/authorization"
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
	// 非调试模式（生产模式） 日志写到日志文件
	if variable.ConfigYml.GetBool("AppDebug") == false {
		//1.gin自行记录接口访问日志，不需要nginx，如果开启以下3行，那么请屏蔽第 34 行代码
		//gin.DisableConsoleColor()
		//f, _ := os.Create(variable.BasePath + variable.ConfigYml.GetString("Logs.GinLogName"))
		//gin.DefaultWriter = io.MultiWriter(f)

		//【生产模式】
		// 根据 gin 官方的说明：[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
		// 如果部署到生产环境，请使用以下模式：
		// 1.生产模式(release) 和开发模式的变化主要是禁用 gin 记录接口访问日志，
		// 2.go服务就必须使用nginx作为前置代理服务，这样也方便实现负载均衡
		// 3.如果程序发生 panic 等异常使用自定义的 panic 恢复中间件拦截、记录到日志
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
	auth := router.Group("base/")
	{
		auth.POST("register", validatorFactory.Create(consts.ValidatorPrefix+"Register"))
		auth.POST("login", validatorFactory.Create(consts.ValidatorPrefix+"Login"))
	}
	router.GET("message/ws", validatorFactory.Create(consts.ValidatorPrefix+"WebsocketConnect"))
	router.Use(authorization.CheckTokenAuth())
	user := router.Group("user/")
	{
		user.GET("userinfo", validatorFactory.Create(consts.ValidatorPrefix+"GetUserInfo"))
		user.GET("video-list", validatorFactory.Create(consts.ValidatorPrefix+"GetUserVideoList"))
		user.GET("panel", validatorFactory.Create(consts.ValidatorPrefix+"GetPanel"))
		user.GET("friends", validatorFactory.Create(consts.ValidatorPrefix+"GetFriends"))
		user.GET("follow", validatorFactory.Create(consts.ValidatorPrefix+"GetFollow"))
		user.GET("fans", validatorFactory.Create(consts.ValidatorPrefix+"GetFans"))

		user.POST("attention", validatorFactory.Create(consts.ValidatorPrefix+"Attention"))
		user.GET("aweme-status", validatorFactory.Create(consts.ValidatorPrefix+"AwemeStatus"))
		// 用户界面视频加载 API
		user.GET("my-video", validatorFactory.Create(consts.ValidatorPrefix+"GetMyVideo"))
		user.GET("my-private", validatorFactory.Create(consts.ValidatorPrefix+"GetMyPrivateVideo"))
		user.GET("my-like-video", validatorFactory.Create(consts.ValidatorPrefix+"GetMyLikeVideo"))
		user.GET("my-collect-video", validatorFactory.Create(consts.ValidatorPrefix+"GetMyCollectVideo"))
		user.GET("my-history-video", validatorFactory.Create(consts.ValidatorPrefix+"GetMyHistoryVideo"))
		user.GET("my-history-other", validatorFactory.Create(consts.ValidatorPrefix+"GetMyHistoryOther"))
	}
	post := router.Group("post/")
	{
		post.GET("recommended", validatorFactory.Create(consts.ValidatorPrefix+"GetPostRecommended"))
	}
	shop := router.Group("shop/")
	{
		shop.GET("recommended", validatorFactory.Create(consts.ValidatorPrefix+"GetShopRecommended"))
	}
	video := router.Group("video/")
	{
		video.POST("digg", validatorFactory.Create(consts.ValidatorPrefix+"VideoDigg"))
		video.POST("comment", validatorFactory.Create(consts.ValidatorPrefix+"VideoComment"))
		video.POST("collect", validatorFactory.Create(consts.ValidatorPrefix+"VideoCollect"))
		video.POST("share", validatorFactory.Create(consts.ValidatorPrefix+"VideoShare"))
		video.GET("comments", validatorFactory.Create(consts.ValidatorPrefix+"GetComments"))
		video.GET("recommended", validatorFactory.Create(consts.ValidatorPrefix+"GetVideoRecommended"))
		video.GET("long-recommended", validatorFactory.Create(consts.ValidatorPrefix+"GetLongVideoRecommended"))
	}

	jwt := router.Group("jwt")
	{
		jwt.POST("jsonInBlacklist", validatorFactory.Create(consts.ValidatorPrefix+"JsonInBlacklist")) // jwt加入黑名单
	}
	msg := router.Group("message")
	{

		msg.GET("all-msg", validatorFactory.Create(consts.ValidatorPrefix+"AllMsg"))
		msg.POST("send-msg", validatorFactory.Create(consts.ValidatorPrefix+"SendMsg"))
	}
	return router
}
