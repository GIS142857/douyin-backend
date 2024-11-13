package web

import (
	"douyin-backend/app/global/consts"
	"douyin-backend/app/model/post"
	"douyin-backend/app/utils/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PostController struct {
}

func (u *PostController) GetPostRecommended(ctx *gin.Context) {
	var uid = auth.GetUidFromToken(ctx)
	var pageNo = ctx.GetFloat64(consts.ValidatorPrefix + "pageNo")
	var pageSize = ctx.GetFloat64(consts.ValidatorPrefix + "pageSize")
	list, total := post.CreatePostFactory("").GetPostRecommended(uid, int64(pageNo), int64(pageSize))
	if len(list) > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"pageNo": pageNo,
			"total":  total,
			"list":   list,
		})
	} else {
		ctx.JSON(http.StatusNoContent, gin.H{
			"pageNo": pageNo,
			"total":  total,
			"list":   []interface{}{}, // 返回一个空数组以确保响应一致性
		})
	}
}
