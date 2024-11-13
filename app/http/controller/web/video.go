package web

import (
	"douyin-backend/app/global/consts"
	"douyin-backend/app/model/video"
	"douyin-backend/app/utils/auth"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/rand"
	"net/http"
	"strconv"
	"time"
)

type VideoController struct {
}

func (v *VideoController) VideoDigg(ctx *gin.Context) {
	var uid = auth.GetUidFromToken(ctx)
	var awemeId = ctx.GetString(consts.ValidatorPrefix + "aweme_id")
	var awemeIDInt64, _ = strconv.ParseInt(awemeId, 10, 64)
	diggDone := video.CreateDiggFactory("").VideoDigg(uid, awemeIDInt64)
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

func (v *VideoController) VideoComment(ctx *gin.Context) {
	var ipLocation = ctx.GetString(consts.ValidatorPrefix + "ip_location")
	var awemeId = ctx.GetString(consts.ValidatorPrefix + "aweme_id")
	var content = ctx.GetString(consts.ValidatorPrefix + "content")
	var uid = auth.GetUidFromToken(ctx)
	var shortId = ctx.GetString(consts.ValidatorPrefix + "short_id")
	var uniqueId = ctx.GetString(consts.ValidatorPrefix + "unique_id")
	var signature = ctx.GetString(consts.ValidatorPrefix + "signature")
	var nickname = ctx.GetString(consts.ValidatorPrefix + "nickname")
	var avatar = ctx.GetString(consts.ValidatorPrefix + "avatar")
	var awemeIDInt64, _ = strconv.ParseInt(awemeId, 10, 64)
	commentDone := video.CreateCommentFactory("").VideoComment(uid, awemeIDInt64, ipLocation, content, shortId, uniqueId, signature, nickname, avatar)
	if commentDone {
		ctx.JSON(http.StatusOK, gin.H{
			"data": commentDone,
			"code": consts.CurdStatusOkCode,
			"msg":  "评论成功",
		})
	} else {
		ctx.JSON(http.StatusForbidden, gin.H{
			"data": commentDone,
			"code": consts.CurdCreatFailCode,
			"msg":  "评论失败",
		})
	}

}

func (v *VideoController) VideoCollect(ctx *gin.Context) {
	var uid = auth.GetUidFromToken(ctx)
	var awemeId = ctx.GetString(consts.ValidatorPrefix + "aweme_id")
	var awemeIDInt64, _ = strconv.ParseInt(awemeId, 10, 64)
	diggDone := video.CreateCollectFactory("").VideoCollect(uid, awemeIDInt64)
	if diggDone {
		ctx.JSON(http.StatusOK, gin.H{
			"data": diggDone,
			"code": consts.CurdStatusOkCode,
			"msg":  "收藏成功",
		})
	} else {
		ctx.JSON(http.StatusForbidden, gin.H{
			"data": diggDone,
			"code": consts.CurdCreatFailCode,
			"msg":  "收藏失败",
		})
	}
}

func (v *VideoController) VideoShare(ctx *gin.Context) {
	var uid = auth.GetUidFromToken(ctx)
	var awemeId = ctx.GetString(consts.ValidatorPrefix + "aweme_id")
	var message = ctx.GetString(consts.ValidatorPrefix + "message")
	var shareUidList = ctx.GetString(consts.ValidatorPrefix + "share_uid_list")
	var awemeIDInt64, _ = strconv.ParseInt(awemeId, 10, 64)
	shareDone := video.CreateShareFactory("").VideoShare(uid, awemeIDInt64, message, shareUidList)
	if shareDone {
		ctx.JSON(http.StatusOK, gin.H{
			"data": shareDone,
			"code": consts.CurdStatusOkCode,
			"msg":  "分享成功",
		})
	} else {
		ctx.JSON(http.StatusForbidden, gin.H{
			"data": shareDone,
			"code": consts.CurdCreatFailCode,
			"msg":  "分享失败",
		})
	}
}

func (v *VideoController) GetComments(ctx *gin.Context) {
	awemeId, _ := strconv.Atoi(ctx.Query("aweme_id"))
	comments := video.CreateCommentFactory("").GetComments(int64(awemeId))
	if len(comments) > 0 {
		ctx.JSON(http.StatusOK, comments)
	} else {
		ctx.JSON(http.StatusNoContent, comments)
	}
}

func (v *VideoController) GetStar(context *gin.Context) {
	// TODO 具体业务逻辑实现

	//var id = context.GetFloat64(consts.ValidatorPrefix + "id")
	//video := sv_home.CreateShortVideoFactory("").GetVideoById(int(id))
	//if video.Id != 0 {
	//	response.Success(context, consts.CurdStatusOkMsg, video)
	//} else {
	//	response.Fail(context, consts.CurdSelectFailCode, consts.CurdSelectFailMsg, "")
	//}
}

func (v *VideoController) GetShare(context *gin.Context) {
	// TODO 具体业务逻辑实现

	//var id = context.GetFloat64(consts.ValidatorPrefix + "id")
	//video := sv_home.CreateShortVideoFactory("").GetVideoById(int(id))
	//if video.Id != 0 {
	//	response.Success(context, consts.CurdStatusOkMsg, video)
	//} else {
	//	response.Fail(context, consts.CurdSelectFailCode, consts.CurdSelectFailMsg, "")
	//}
}

func (v *VideoController) GetHistoryOther(context *gin.Context) {
	// TODO 具体业务逻辑实现

	//var id = context.GetFloat64(consts.ValidatorPrefix + "id")
	//video := sv_home.CreateShortVideoFactory("").GetVideoById(int(id))
	//if video.Id != 0 {
	//	response.Success(context, consts.CurdStatusOkMsg, video)
	//} else {
	//	response.Fail(context, consts.CurdSelectFailCode, consts.CurdSelectFailMsg, "")
	//}
}

func (v *VideoController) GetLongVideoRecommended(ctx *gin.Context) {
	// TODO 具体业务逻辑实现
	var uid = auth.GetUidFromToken(ctx)
	var PageNo = ctx.GetFloat64(consts.ValidatorPrefix + "pageNo")
	var PageSize = ctx.GetFloat64(consts.ValidatorPrefix + "pageSize")
	list, total := video.CreateVideoFactory("").GetLongVideoRecommended(uid, int64(PageNo), int64(PageSize))
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

func (v *VideoController) GetVideoRecommended(ctx *gin.Context) {
	var uid = auth.GetUidFromToken(ctx)
	var Start = ctx.GetFloat64(consts.ValidatorPrefix + "start")
	var PageSize = ctx.GetFloat64(consts.ValidatorPrefix + "pageSize")
	list, total := video.CreateVideoFactory("").GetVideoRecommended(uid, int64(Start), int64(PageSize))
	if len(list) > 0 {
		rand.Seed(uint64(time.Now().UnixNano()))
		// 打乱切片
		rand.Shuffle(len(list), func(i, j int) {
			list[i], list[j] = list[j], list[i]
		})
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

func (v *VideoController) GetHistory(context *gin.Context) {
	// TODO 具体业务逻辑实现

	//var id = context.GetFloat64(consts.ValidatorPrefix + "id")
	//video := sv_home.CreateShortVideoFactory("").GetVideoById(int(id))
	//if video.Id != 0 {
	//	response.Success(context, consts.CurdStatusOkMsg, video)
	//} else {
	//	response.Fail(context, consts.CurdSelectFailCode, consts.CurdSelectFailMsg, "")
	//}
}
