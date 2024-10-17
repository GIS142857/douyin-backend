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
	// user
	{
		key = consts.ValidatorPrefix + "GetUserInfo"
		containers.Set(key, user.GetUserInfo{})

		key = consts.ValidatorPrefix + "GetVideoList"
		containers.Set(key, user.GetVideoList{})

		key = consts.ValidatorPrefix + "GetPanel"
		containers.Set(key, user.GetPanel{})

		key = consts.ValidatorPrefix + "GetFriends"
		containers.Set(key, user.GetFriends{})

		key = consts.ValidatorPrefix + "GetCollect"
		containers.Set(key, user.GetCollect{})
	}
	// video
	{
		key = consts.ValidatorPrefix + "GetLike"
		containers.Set(key, video.GetLike{})

		key = consts.ValidatorPrefix + "GetComments"
		containers.Set(key, video.GetComments{})

		key = consts.ValidatorPrefix + "GetStar"
		containers.Set(key, video.GetStar{})

		key = consts.ValidatorPrefix + "GetShare"
		containers.Set(key, video.GetShare{})

		key = consts.ValidatorPrefix + "GetHistoryOther"
		containers.Set(key, video.GetHistoryOther{})

		key = consts.ValidatorPrefix + "GetHistory"
		containers.Set(key, video.GetHistory{})

		key = consts.ValidatorPrefix + "GetLongRecommended"
		containers.Set(key, video.GetLongRecommended{})

		key = consts.ValidatorPrefix + "GetMyVideo"
		containers.Set(key, video.GetMyVideo{})

		key = consts.ValidatorPrefix + "GetPrivate"
		containers.Set(key, video.GetPrivate{})
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
