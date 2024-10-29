package web

import (
	"douyin-backend/app/global/consts"
	"douyin-backend/app/model/user"
	"douyin-backend/app/model/video"
	"github.com/gin-gonic/gin"
	"strconv"
)

type VideoController struct {
}

func (u *VideoController) GetComments(ctx *gin.Context) {
	aweme_id, _ := strconv.Atoi(ctx.Query("aweme_id"))
	comments := video.CreateCommentFactory("").GetComments(int64(aweme_id))
	if len(comments) > 0 {
		ctx.JSON(consts.CurdStatusOkCode, comments)
	} else {
		ctx.JSON(consts.CurdSelectFailCode, comments)
	}
}

func (u *VideoController) GetStar(context *gin.Context) {
	// TODO 具体业务逻辑实现

	//var id = context.GetFloat64(consts.ValidatorPrefix + "id")
	//video := sv_home.CreateShortVideoFactory("").GetVideoById(int(id))
	//if video.Id != 0 {
	//	response.Success(context, consts.CurdStatusOkMsg, video)
	//} else {
	//	response.Fail(context, consts.CurdSelectFailCode, consts.CurdSelectFailMsg, "")
	//}
}

func (u *VideoController) GetShare(context *gin.Context) {
	// TODO 具体业务逻辑实现

	//var id = context.GetFloat64(consts.ValidatorPrefix + "id")
	//video := sv_home.CreateShortVideoFactory("").GetVideoById(int(id))
	//if video.Id != 0 {
	//	response.Success(context, consts.CurdStatusOkMsg, video)
	//} else {
	//	response.Fail(context, consts.CurdSelectFailCode, consts.CurdSelectFailMsg, "")
	//}
}

func (u *VideoController) GetHistoryOther(context *gin.Context) {
	// TODO 具体业务逻辑实现

	//var id = context.GetFloat64(consts.ValidatorPrefix + "id")
	//video := sv_home.CreateShortVideoFactory("").GetVideoById(int(id))
	//if video.Id != 0 {
	//	response.Success(context, consts.CurdStatusOkMsg, video)
	//} else {
	//	response.Fail(context, consts.CurdSelectFailCode, consts.CurdSelectFailMsg, "")
	//}
}

func (u *VideoController) GetLongRecommended(ctx *gin.Context) {
	// TODO 具体业务逻辑实现

}

func (u *VideoController) GetVideoRecommended(ctx *gin.Context) {
	//var Uid = ctx.GetFloat64(consts.ValidatorPrefix + "uid")
	var Start = ctx.GetFloat64(consts.ValidatorPrefix + "start")
	var PageSize = ctx.GetFloat64(consts.ValidatorPrefix + "pageSize")
	list, total := user.CreateVideoFactory("").GetVideoRecommended(int64(Start), int64(PageSize))
	if len(list) > 0 {
		ctx.JSON(consts.CurdStatusOkCode, gin.H{
			"total": total,
			"list":  list,
		})
	} else {
		ctx.JSON(consts.CurdSelectFailCode, gin.H{
			"total": total,
			"list":  []interface{}{}, // 返回一个空数组以确保响应一致性
		})
	}
}

func (u *VideoController) GetHistory(context *gin.Context) {
	// TODO 具体业务逻辑实现

	//var id = context.GetFloat64(consts.ValidatorPrefix + "id")
	//video := sv_home.CreateShortVideoFactory("").GetVideoById(int(id))
	//if video.Id != 0 {
	//	response.Success(context, consts.CurdStatusOkMsg, video)
	//} else {
	//	response.Fail(context, consts.CurdSelectFailCode, consts.CurdSelectFailMsg, "")
	//}
}
