package register_validator

import (
	"douyin-backend/app/core/container"
	"douyin-backend/app/global/consts"
	"douyin-backend/app/http/validator/web/douyin/post"
	"douyin-backend/app/http/validator/web/douyin/shop"
	"douyin-backend/app/http/validator/web/douyin/user"
	"douyin-backend/app/http/validator/web/douyin/video"
)

func WebRegisterValidator() {
	containers := container.CreateContainersFactory()

	var key string
	// jwt
	{
		key = consts.ValidatorPrefix + "JsonInBlacklist"
		containers.Set(key, user.JsonInBlacklist{})
	}

	// user
	{
		key = consts.ValidatorPrefix + "UserLogin"
		containers.Set(key, user.Login{})

		key = consts.ValidatorPrefix + "GetUserInfo"
		containers.Set(key, user.GetUserInfo{})

		key = consts.ValidatorPrefix + "GetUserVideoList"
		containers.Set(key, user.GetUserVideoList{})

		key = consts.ValidatorPrefix + "GetPanel"
		containers.Set(key, user.GetPanel{})

		key = consts.ValidatorPrefix + "GetFriends"
		containers.Set(key, user.GetFriends{})

		key = consts.ValidatorPrefix + "GetMyVideo"
		containers.Set(key, user.GetMyVideo{})

		key = consts.ValidatorPrefix + "GetMyPrivateVideo"
		containers.Set(key, user.GetMyPrivateVideo{})

		key = consts.ValidatorPrefix + "GetMyLikeVideo"
		containers.Set(key, user.GetMyLikeVideo{})

		key = consts.ValidatorPrefix + "GetMyCollectVideo"
		containers.Set(key, user.GetMyCollectVideo{})

		key = consts.ValidatorPrefix + "GetMyHistoryVideo"
		containers.Set(key, user.GetMyHistoryVideo{})

		key = consts.ValidatorPrefix + "GetMyHistoryOther"
		containers.Set(key, user.GetMyHistoryOther{})

	}
	// video
	{
		key = consts.ValidatorPrefix + "GetVideoRecommended"
		containers.Set(key, video.GetVideoRecommended{})

		key = consts.ValidatorPrefix + "GetLongVideoRecommended"
		containers.Set(key, video.GetLongVideoRecommended{})

		key = consts.ValidatorPrefix + "GetComments"
		containers.Set(key, video.GetComments{})

		key = consts.ValidatorPrefix + "GetStar"
		containers.Set(key, video.GetStar{})

		key = consts.ValidatorPrefix + "GetShare"
		containers.Set(key, video.GetShare{})

		key = consts.ValidatorPrefix + "VideoDigg"
		containers.Set(key, video.VideoDigg{})

		key = consts.ValidatorPrefix + "VideoComment"
		containers.Set(key, video.VideoComment{})

		key = consts.ValidatorPrefix + "VideoCollect"
		containers.Set(key, video.VideoCollect{})

		key = consts.ValidatorPrefix + "VideoShare"
		containers.Set(key, video.VideoShare{})
	}
	// shop
	{
		key = consts.ValidatorPrefix + "GetShopRecommended"
		containers.Set(key, shop.GetShopRecommended{})
	}

	// post
	{
		key = consts.ValidatorPrefix + "GetPostRecommended"
		containers.Set(key, post.GetPostRecommended{})
	}
}
