package web

import (
	"douyin-backend/app/global/consts"
	"douyin-backend/app/model/video"
	"douyin-backend/app/utils/auth"
	"douyin-backend/app/utils/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/rand"
	"strconv"
	"time"
)

type VideoController struct {
}

func (v *VideoController) VideoDigg(ctx *gin.Context) {
	var uid = auth.GetUidFromToken(ctx)
	var awemeId = ctx.GetString(consts.ValidatorPrefix + "aweme_id")
	var action = ctx.GetBool(consts.ValidatorPrefix + "action")
	var awemeIDInt64, _ = strconv.ParseInt(awemeId, 10, 64)
	actionStatus := video.CreateDiggFactory("").VideoDigg(uid, awemeIDInt64, action)
	if actionStatus {
		if action {
			response.Success(ctx, consts.CurdStatusOkMsg, gin.H{
				"data": actionStatus,
				"msg":  "点赞成功",
			})
		} else {
			response.Success(ctx, consts.CurdStatusOkMsg, gin.H{
				"data": actionStatus,
				"msg":  "取消点赞成功",
			})
		}
	} else {
		if action {
			response.Fail(ctx, consts.CurdInsertFailCode, consts.CurdInsertFailMsg, gin.H{
				"data": actionStatus,
				"msg":  "点赞失败",
			})
		} else {
			response.Fail(ctx, consts.CurdInsertFailCode, consts.CurdInsertFailMsg, gin.H{
				"data": actionStatus,
				"msg":  "取消点赞失败",
			})
		}
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
		response.Success(ctx, consts.CurdStatusOkMsg, gin.H{
			"data": commentDone,
			"msg":  "评论成功",
		})
	} else {
		response.Fail(ctx, consts.CurdInsertFailCode, consts.CurdInsertFailMsg, gin.H{
			"data": commentDone,
			"msg":  "评论失败",
		})
	}

}

func (v *VideoController) VideoCollect(ctx *gin.Context) {
	var uid = auth.GetUidFromToken(ctx)
	var awemeId = ctx.GetString(consts.ValidatorPrefix + "aweme_id")
	var action = ctx.GetBool(consts.ValidatorPrefix + "action")
	var awemeIDInt64, _ = strconv.ParseInt(awemeId, 10, 64)
	actionStatus := video.CreateCollectFactory("").VideoCollect(uid, awemeIDInt64, action)
	if actionStatus {
		if action {
			response.Success(ctx, consts.CurdStatusOkMsg, gin.H{
				"data": actionStatus,
				"msg":  "收藏成功",
			})
		} else {
			response.Success(ctx, consts.CurdStatusOkMsg, gin.H{
				"data": actionStatus,
				"msg":  "取消收藏成功",
			})
		}
	} else {
		if action {
			response.Fail(ctx, consts.CurdInsertFailCode, consts.CurdInsertFailMsg, gin.H{
				"data": actionStatus,
				"msg":  "收藏失败",
			})
		} else {
			response.Fail(ctx, consts.CurdInsertFailCode, consts.CurdInsertFailMsg, gin.H{
				"data": actionStatus,
				"msg":  "取消收藏失败",
			})
		}
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
		response.Success(ctx, consts.CurdStatusOkMsg, gin.H{
			"data": shareDone,
			"msg":  "分享成功",
		})
	} else {
		response.Fail(ctx, consts.CurdInsertFailCode, consts.CurdInsertFailMsg, gin.H{
			"data": shareDone,
			"msg":  "分享失败",
		})
	}
}

func (v *VideoController) GetComments(ctx *gin.Context) {
	awemeId, _ := strconv.Atoi(ctx.Query("aweme_id"))
	comments, ok := video.CreateCommentFactory("").GetComments(int64(awemeId))
	if ok {
		response.Success(ctx, consts.CurdStatusOkMsg, comments)
	} else {
		response.Fail(ctx, consts.CurdInsertFailCode, consts.CurdInsertFailMsg, "")
	}
}

func (v *VideoController) GetHistoryOther(context *gin.Context) {
	// TODO 具体业务逻辑实现
}

func (v *VideoController) GetLongVideoRecommended(ctx *gin.Context) {
	// TODO 具体业务逻辑实现
	var uid = auth.GetUidFromToken(ctx)
	var PageNo = ctx.GetFloat64(consts.ValidatorPrefix + "pageNo")
	var PageSize = ctx.GetFloat64(consts.ValidatorPrefix + "pageSize")
	list, total, ok := video.CreateVideoFactory("").GetLongVideoRecommended(uid, int64(PageNo), int64(PageSize))
	if ok {
		response.Success(ctx, consts.CurdStatusOkMsg, gin.H{
			"total": total,
			"list":  list,
		})
	} else {
		response.Fail(ctx, consts.CurdSelectFailCode, consts.CurdSelectFailMsg, "")
	}
}

func (v *VideoController) GetVideoRecommended(ctx *gin.Context) {
	var uid = auth.GetUidFromToken(ctx)
	var Start = ctx.GetFloat64(consts.ValidatorPrefix + "start")
	var PageSize = ctx.GetFloat64(consts.ValidatorPrefix + "pageSize")
	list, total, ok := video.CreateVideoFactory("").GetVideoRecommended(uid, int64(Start), int64(PageSize))
	if ok && len(list) > 0 {
		rand.Seed(uint64(time.Now().UnixNano()))
		// 打乱切片
		rand.Shuffle(len(list), func(i, j int) {
			list[i], list[j] = list[j], list[i]
		})
		response.Success(ctx, consts.CurdStatusOkMsg, gin.H{
			"total": total,
			"list":  list,
		})
	} else {
		response.Fail(ctx, consts.CurdSelectFailCode, consts.CurdSelectFailMsg, "")
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
