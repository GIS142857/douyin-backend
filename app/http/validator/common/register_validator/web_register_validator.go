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

		key = consts.ValidatorPrefix + "GetMyVideo"
		containers.Set(key, user.GetMyVideo{})

		key = consts.ValidatorPrefix + "GetMyPrivateVideo"
		containers.Set(key, user.GetMyPrivateVideo{})

		key = consts.ValidatorPrefix + "GetMyLikeVideo"
		containers.Set(key, user.GetMyLikeVideo{})

		key = consts.ValidatorPrefix + "GetMyCollectVideo"
		containers.Set(key, user.GetMyCollectVideo{})

	}
	// video
	{
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
