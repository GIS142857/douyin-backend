package register_validator

import (
	"douyin-backend/app/core/container"
	"douyin-backend/app/global/consts"
	"douyin-backend/app/http/validator/common/websocket"
	"douyin-backend/app/http/validator/web/douyin/message"
	"douyin-backend/app/http/validator/web/douyin/post"
	"douyin-backend/app/http/validator/web/douyin/shop"
	"douyin-backend/app/http/validator/web/douyin/upload"
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

	// base
	{
		key = consts.ValidatorPrefix + "Login"
		containers.Set(key, user.Login{})

		key = consts.ValidatorPrefix + "Register"
		containers.Set(key, user.Register{})
	}
	// upload
	{
		key = consts.ValidatorPrefix + "Avatar"
		containers.Set(key, upload.Avatar{})

		key = consts.ValidatorPrefix + "Cover"
		containers.Set(key, upload.Cover{})

		key = consts.ValidatorPrefix + "Video"
		containers.Set(key, upload.Video{})
	}

	// user
	{
		key = consts.ValidatorPrefix + "GetUserInfo"
		containers.Set(key, user.GetUserInfo{})

		key = consts.ValidatorPrefix + "UpdateInfo"
		containers.Set(key, user.UpdateInfo{})

		key = consts.ValidatorPrefix + "GetUserVideoList"
		containers.Set(key, user.GetUserVideoList{})

		key = consts.ValidatorPrefix + "GetPanel"
		containers.Set(key, user.GetPanel{})

		key = consts.ValidatorPrefix + "Attention"
		containers.Set(key, user.Attention{})

		key = consts.ValidatorPrefix + "AwemeStatus"
		containers.Set(key, user.AwemeStatus{})

		key = consts.ValidatorPrefix + "GetFriends"
		containers.Set(key, user.GetFriends{})

		key = consts.ValidatorPrefix + "GetFollow"
		containers.Set(key, user.GetFollow{})

		key = consts.ValidatorPrefix + "GetFans"
		containers.Set(key, user.GetFans{})

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
	// msg
	{
		key = consts.ValidatorPrefix + "WebsocketConnect"
		containers.Set(key, websocket.Connect{})

		key = consts.ValidatorPrefix + "AllMsg"
		containers.Set(key, message.AllMsg{})

		key = consts.ValidatorPrefix + "SendMsg"
		containers.Set(key, message.SendMsg{})
	}

}
