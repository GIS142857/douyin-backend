package web

import (
	"douyin-backend/app/global/consts"
	"douyin-backend/app/global/variable"
	"douyin-backend/app/model/user"
	"douyin-backend/app/model/video"
	userstoken "douyin-backend/app/service/users/token"
	"douyin-backend/app/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController struct {
}

func (u *UserController) Login(ctx *gin.Context) {
	var phone = ctx.GetString(consts.ValidatorPrefix + "phone")
	var password = ctx.GetString(consts.ValidatorPrefix + "password")
	userModel := user.CreateUserFactory("").Login(phone, password)
	if userModel.UID != 0 {
		userTokenFactory := userstoken.CreateUserFactory()
		if userToken, err := userTokenFactory.GenerateToken(userModel.UID, userModel.NickName, userModel.Phone, userModel.Password, variable.ConfigYml.GetInt64("Token.JwtTokenCreatedExpireAt")); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"isExist": true,
				"uid":     userModel.UID,
				"token":   userToken,
				"msg":     "用户存在",
			})
		}
	} else {
		ctx.JSON(http.StatusNoContent, gin.H{
			"isExist": false,
			"uid":     userModel.UID,
			"token":   "",
			"msg":     "用户不存在",
		})
	}
}
func (u *UserController) JsonInBlacklist(ctx *gin.Context) {
	// TODO
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

func (u *UserController) GetUserVideoList(ctx *gin.Context) {
	// TODO 具体业务逻辑实现
	uid, _ := strconv.Atoi(ctx.Query("uid"))
	videoList := video.CreateVideoFactory("").GetUserVideoList(int64(uid))
	if len(videoList) > 0 {
		ctx.JSON(http.StatusOK, videoList)
	} else {
		ctx.JSON(http.StatusNoContent, []interface{}{})
	}
}

func (u *UserController) GetPanel(ctx *gin.Context) {
	// TODO 具体业务逻辑实现

	var Uid = ctx.GetFloat64(consts.ValidatorPrefix + "uid")
	userinfo := user.CreateUserFactory("").GetPanel(int64(Uid))
	if userinfo.Uid > 0 {
		ctx.JSON(http.StatusOK, userinfo)
	} else {
		ctx.JSON(http.StatusNoContent, "")
	}
}

func (u *UserController) GetFriends(ctx *gin.Context) {
	var uid = ctx.GetFloat64(consts.ValidatorPrefix + "uid")
	friends := user.CreateUserFactory("").GetFriends(int64(uid))
	if len(friends) > 0 {
		ctx.JSON(http.StatusOK, friends)
	} else {
		ctx.JSON(http.StatusNoContent, []interface{}{})
	}
}

func (u *UserController) GetMyVideo(ctx *gin.Context) {
	var Uid = ctx.GetFloat64(consts.ValidatorPrefix + "uid")
	var PageNo = ctx.GetFloat64(consts.ValidatorPrefix + "pageNo")
	var PageSize = ctx.GetFloat64(consts.ValidatorPrefix + "pageSize")
	list, total := video.CreateVideoFactory("").GetMyVideo(int64(Uid), int64(PageNo), int64(PageSize))
	if len(list) > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"pageNo": PageNo,
			"total":  total,
			"list":   list,
		})
	} else {
		ctx.JSON(http.StatusNoContent, gin.H{
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
	list, total := video.CreateVideoFactory("").GetMyPrivateVideo(int64(Uid), int64(PageNo), int64(PageSize))
	if len(list) > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"pageNo": PageNo,
			"total":  total,
			"list":   list,
		})
	} else {
		ctx.JSON(http.StatusNoContent, gin.H{
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
	list, total := video.CreateVideoFactory("").GetMyLikeVideo(int64(Uid), int64(PageNo), int64(PageSize))
	if len(list) > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"pageNo": PageNo,
			"total":  total,
			"list":   list,
		})
	} else {
		ctx.JSON(http.StatusNoContent, gin.H{
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
	list, total := video.CreateVideoFactory("").GetMyCollectVideo(int64(Uid), int64(PageNo), int64(PageSize))
	if len(list) > 0 {
		ctx.JSON(http.StatusOK, gin.H{
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
		ctx.JSON(http.StatusNoContent, gin.H{
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

	list, total := video.CreateVideoFactory("").GetMyHistoryVideo(int64(Uid), int64(PageNo), int64(PageSize))

	if len(list) > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"pageNo": PageNo,
			"total":  total,
			"list":   list,
		})
	} else {
		ctx.JSON(http.StatusNoContent, gin.H{
			"pageNo": PageNo,
			"total":  total,
			"list":   []interface{}{}, // 返回一个空数组以确保响应一致性
		})
	}
}

func (u *UserController) GetMyHistoryOther(ctx *gin.Context) {
	response.Success(ctx, consts.CurdStatusOkMsg, "GetMyHistoryOther-ok")
}
