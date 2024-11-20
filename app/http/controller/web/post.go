package web

import (
	"douyin-backend/app/global/consts"
	"douyin-backend/app/model/post"
	"douyin-backend/app/utils/auth"
	"douyin-backend/app/utils/response"
	"github.com/gin-gonic/gin"
)

type PostController struct {
}

func (u *PostController) GetPostRecommended(ctx *gin.Context) {
	var uid = auth.GetUidFromToken(ctx)
	var pageNo = ctx.GetFloat64(consts.ValidatorPrefix + "pageNo")
	var pageSize = ctx.GetFloat64(consts.ValidatorPrefix + "pageSize")
	list, total, ok := post.CreatePostFactory("").GetPostRecommended(uid, int64(pageNo), int64(pageSize))
	if !ok {
		response.Fail(ctx, consts.CurdSelectFailCode, consts.CurdSelectFailMsg, "获取post推荐失败!")
	} else {
		response.Success(ctx, consts.CurdStatusOkMsg, gin.H{
			"pageNo": pageNo,
			"total":  total,
			"list":   list,
		})
	}
}
