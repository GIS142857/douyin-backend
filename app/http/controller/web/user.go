package web

import (
	"douyin-backend/app/global/consts"
	"douyin-backend/app/model/user"
	"douyin-backend/app/utils/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (u *UserController) GetUserInfo(context *gin.Context) {
	// TODO 具体业务逻辑实现
	response.Success(context, consts.CurdStatusOkMsg, "GetUserInfo-ok")
	//var id = context.GetFloat64(consts.ValidatorPrefix + "id")
	//video := sv_home.CreateShortVideoFactory("").GetVideoById(int(id))
	//if video.Id != 0 {
	//	response.Success(context, consts.CurdStatusOkMsg, video)
	//} else {
	//	response.Fail(context, consts.CurdSelectFailCode, consts.CurdSelectFailMsg, "")
	//}
}

func (u *UserController) GetVideoList(context *gin.Context) {
	// TODO 具体业务逻辑实现
	response.Success(context, consts.CurdStatusOkMsg, "GetVideoList-ok")
	//var id = context.GetFloat64(consts.ValidatorPrefix + "id")
	//video := sv_home.CreateShortVideoFactory("").GetVideoById(int(id))
	//if video.Id != 0 {
	//	response.Success(context, consts.CurdStatusOkMsg, video)
	//} else {
	//	response.Fail(context, consts.CurdSelectFailCode, consts.CurdSelectFailMsg, "")
	//}
}

func (u *UserController) GetPanel(ctx *gin.Context) {
	// TODO 具体业务逻辑实现

	var Uid = ctx.GetFloat64(consts.ValidatorPrefix + "uid")
	userinfo := user.CreateUserFactory("").GetPanel(int64(Uid))
	if userinfo.Uid > 0 {
		ctx.JSON(consts.CurdStatusOkCode, userinfo)
	} else {
		ctx.JSON(consts.CurdSelectFailCode, "")
	}
}

func (u *UserController) GetFriends(context *gin.Context) {
	// TODO 具体业务逻辑实现
	response.Success(context, consts.CurdStatusOkMsg, "GetFriends-ok")
	//var id = context.GetFloat64(consts.ValidatorPrefix + "id")
	//video := sv_home.CreateShortVideoFactory("").GetVideoById(int(id))
	//if video.Id != 0 {
	//	response.Success(context, consts.CurdStatusOkMsg, video)
	//} else {
	//	response.Fail(context, consts.CurdSelectFailCode, consts.CurdSelectFailMsg, "")
	//}
}

func (u *UserController) GetMyVideo(ctx *gin.Context) {
	var Uid = ctx.GetFloat64(consts.ValidatorPrefix + "uid")
	var PageNo = ctx.GetFloat64(consts.ValidatorPrefix + "pageNo")
	var PageSize = ctx.GetFloat64(consts.ValidatorPrefix + "pageSize")
	list, total := user.CreateVideoFactory("").GetMyVideo(int64(Uid), int64(PageNo), int64(PageSize))
	if len(list) > 0 {
		ctx.JSON(consts.CurdStatusOkCode, gin.H{
			"pageNo": PageNo,
			"total":  total,
			"list":   list,
		})
	} else {
		ctx.JSON(consts.CurdSelectFailCode, gin.H{
			"pageNo": PageNo,
			"total":  total,
			"list":   []interface{}{}, // 返回一个空数组以确保响应一致性
		})
	}
}

func (u *UserController) GetMyPrivateVideo(ctx *gin.Context) {
	var Uid = ctx.GetFloat64(consts.ValidatorPrefix + "uid")
	var PageNo = ctx.GetFloat64(consts.ValidatorPrefix + "pageNo")
	var PageSize = ctx.GetFloat64(consts.ValidatorPrefix + "pageSize")
	list, total := user.CreateVideoFactory("").GetMyPrivateVideo(int64(Uid), int64(PageNo), int64(PageSize))
	if len(list) > 0 {
		ctx.JSON(consts.CurdStatusOkCode, gin.H{
			"pageNo": PageNo,
			"total":  total,
			"list":   list,
		})
	} else {
		ctx.JSON(consts.CurdSelectFailCode, gin.H{
			"pageNo": PageNo,
			"total":  total,
			"list":   []interface{}{}, // 返回一个空数组以确保响应一致性
		})
	}
}

func (u *UserController) GetMyLikeVideo(ctx *gin.Context) {
	var Uid = ctx.GetFloat64(consts.ValidatorPrefix + "uid")
	var PageNo = ctx.GetFloat64(consts.ValidatorPrefix + "pageNo")
	var PageSize = ctx.GetFloat64(consts.ValidatorPrefix + "pageSize")
	list, total := user.CreateVideoFactory("").GetMyLikeVideo(int64(Uid), int64(PageNo), int64(PageSize))
	if len(list) > 0 {
		ctx.JSON(consts.CurdStatusOkCode, gin.H{
			"pageNo": PageNo,
			"total":  total,
			"list":   list,
		})
	} else {
		ctx.JSON(consts.CurdSelectFailCode, gin.H{
			"pageNo": PageNo,
			"total":  total,
			"list":   []interface{}{}, // 返回一个空数组以确保响应一致性
		})
	}
}

func (u *UserController) GetMyCollectVideo(ctx *gin.Context) {
	var Uid = ctx.GetFloat64(consts.ValidatorPrefix + "uid")
	var PageNo = ctx.GetFloat64(consts.ValidatorPrefix + "pageNo")
	var PageSize = ctx.GetFloat64(consts.ValidatorPrefix + "pageSize")
	list, total := user.CreateVideoFactory("").GetMyCollectVideo(int64(Uid), int64(PageNo), int64(PageSize))
	if len(list) > 0 {
		ctx.JSON(consts.CurdStatusOkCode, gin.H{
			"video": gin.H{
				"pageNo": PageNo,
				"total":  total,
				"list":   list,
			},
			"music": gin.H{
				"list":  []interface{}{},
				"total": 0,
			},
		})
	} else {
		ctx.JSON(consts.CurdSelectFailCode, gin.H{
			"video": gin.H{
				"pageNo": PageNo,
				"total":  total,
				"list":   []interface{}{}, // 返回一个空数组以确保响应一致性
			},
			"music": gin.H{
				"list":  []interface{}{},
				"total": 0,
			},
		})
	}
}

func (u *UserController) GetMyHistoryVideo(ctx *gin.Context) {
	var Uid = ctx.GetFloat64(consts.ValidatorPrefix + "uid")
	var PageNo = ctx.GetFloat64(consts.ValidatorPrefix + "pageNo")
	var PageSize = ctx.GetFloat64(consts.ValidatorPrefix + "pageSize")

	list, total := user.CreateVideoFactory("").GetMyHistoryVideo(int64(Uid), int64(PageNo), int64(PageSize))

	if len(list) > 0 {
		ctx.JSON(consts.CurdStatusOkCode, gin.H{
			"pageNo": PageNo,
			"total":  total,
			"list":   list,
		})
	} else {
		ctx.JSON(consts.CurdSelectFailCode, gin.H{
			"pageNo": PageNo,
			"total":  total,
			"list":   []interface{}{}, // 返回一个空数组以确保响应一致性
		})
	}
}

func (u *UserController) GetMyHistoryOther(ctx *gin.Context) {
	response.Success(ctx, consts.CurdStatusOkMsg, "GetMyHistoryOther-ok")
}
