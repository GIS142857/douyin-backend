package register_validator

import (
	"douyin-backend/app/core/container"
	"douyin-backend/app/global/consts"
	"douyin-backend/app/http/validator/web/douyin/shop"
	"douyin-backend/app/http/validator/web/douyin/user"
	"douyin-backend/app/http/validator/web/douyin/video"
)

func WebRegisterValidator() {
	containers := container.CreateContainersFactory()

	var key string
	// user
	{
		key = consts.ValidatorPrefix + "UserInfo"
		containers.Set(key, user.GetUserInfo{})

		key = consts.ValidatorPrefix + "VideoList"
		containers.Set(key, user.GetVideoList{})

		key = consts.ValidatorPrefix + "Panel"
		containers.Set(key, user.GetPanel{})

		key = consts.ValidatorPrefix + "Friends"
		containers.Set(key, user.GetFriends{})

		key = consts.ValidatorPrefix + "Collect"
		containers.Set(key, user.GetCollect{})
	}
	// video
	{
		key = consts.ValidatorPrefix + "Like"
		containers.Set(key, video.GetLike{})

		key = consts.ValidatorPrefix + "Comments"
		containers.Set(key, video.GetComments{})

		key = consts.ValidatorPrefix + "Star"
		containers.Set(key, video.GetStar{})

		key = consts.ValidatorPrefix + "Share"
		containers.Set(key, video.GetShare{})

		key = consts.ValidatorPrefix + "HistoryOther"
		containers.Set(key, video.GetHistoryOther{})

		key = consts.ValidatorPrefix + "History"
		containers.Set(key, video.GetHistory{})

		key = consts.ValidatorPrefix + "LongRecommended"
		containers.Set(key, video.GetLongRecommended{})

		key = consts.ValidatorPrefix + "My"
		containers.Set(key, video.GetMy{})

		key = consts.ValidatorPrefix + "Private"
		containers.Set(key, video.GetPrivate{})
	}
	// shop
	{
		key = consts.ValidatorPrefix + "ShopRecommended"
		containers.Set(key, shop.GetRecommended{})
	}
}
