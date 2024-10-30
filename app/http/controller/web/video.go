package web

import (
	"douyin-backend/app/global/consts"
	"douyin-backend/app/model/video"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type VideoController struct {
}

func (u *VideoController) VideoDigg(ctx *gin.Context) {
	var uid = ctx.GetString(consts.ValidatorPrefix + "uid")
	var aweme_id = ctx.GetString(consts.ValidatorPrefix + "aweme_id")
	var uidInt64, _ = strconv.ParseInt(uid, 10, 64)
	var awemeIDInt64, _ = strconv.ParseInt(aweme_id, 10, 64)
	diggDone := video.CreateVideoFactory("").VideoDigg(uidInt64, awemeIDInt64)
	if diggDone {
		ctx.JSON(http.StatusOK, gin.H{
			"data": diggDone,
			"code": consts.CurdStatusOkCode,
			"msg":  "点赞成功",
		})
	} else {
		ctx.JSON(http.StatusForbidden, gin.H{
			"data": diggDone,
			"code": consts.CurdCreatFailCode,
			"msg":  "点赞失败",
		})
	}
}

func (u *VideoController) GetComments(ctx *gin.Context) {
	aweme_id, _ := strconv.Atoi(ctx.Query("aweme_id"))
	comments := video.CreateCommentFactory("").GetComments(int64(aweme_id))
	if len(comments) > 0 {
		ctx.JSON(http.StatusOK, comments)
	} else {
		ctx.JSON(http.StatusNoContent, comments)
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

func (u *VideoController) GetLongVideoRecommended(ctx *gin.Context) {
	// TODO 具体业务逻辑实现
	var Uid = ctx.GetFloat64(consts.ValidatorPrefix + "uid")
	var PageNo = ctx.GetFloat64(consts.ValidatorPrefix + "pageNo")
	var PageSize = ctx.GetFloat64(consts.ValidatorPrefix + "pageSize")
	list, total := video.CreateVideoFactory("").GetLongVideoRecommended(int64(Uid), int64(PageNo), int64(PageSize))
	if len(list) > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"total": total,
			"list":  list,
		})
	} else {
		ctx.JSON(http.StatusNoContent, gin.H{
			"total": total,
			"list":  []interface{}{}, // 返回一个空数组以确保响应一致性
		})
	}
}

func (u *VideoController) GetVideoRecommended(ctx *gin.Context) {
	Uid, _ := strconv.Atoi(ctx.Query("uid"))
	var Start = ctx.GetFloat64(consts.ValidatorPrefix + "start")
	var PageSize = ctx.GetFloat64(consts.ValidatorPrefix + "pageSize")
	list, total := video.CreateVideoFactory("").GetVideoRecommended(int64(Uid), int64(Start), int64(PageSize))
	if len(list) > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"total": total,
			"list":  list,
		})
	} else {
		ctx.JSON(http.StatusNoContent, gin.H{
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
