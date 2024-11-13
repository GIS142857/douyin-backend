package web

import (
	"douyin-backend/app/global/consts"
	"douyin-backend/app/model/shop"
	"douyin-backend/app/utils/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ShopController struct {
}

func (u *ShopController) GetShopRecommended(ctx *gin.Context) {
	// TODO 具体业务逻辑实现
	var uid = auth.GetUidFromToken(ctx)
	var pageNo = ctx.GetFloat64(consts.ValidatorPrefix + "pageNo")
	var pageSize = ctx.GetFloat64(consts.ValidatorPrefix + "pageSize")
	list, total := shop.CreateShopFactory("").GetShopRecommended(uid, int64(pageNo), int64(pageSize))
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
