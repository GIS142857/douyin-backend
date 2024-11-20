package video

import (
	"douyin-backend/app/global/variable"
	"douyin-backend/app/model"
	"github.com/goccy/go-json"
	"go.uber.org/zap"
	"golang.org/x/exp/rand"
	"gorm.io/gorm"
	"time"
)

type VideoModel struct {
	*gorm.DB        `gorm:"-" json:"-"`
	AwemeID         int64           `json:"aweme_id"`         // bigint
	VideoDesc       string          `json:"video_desc"`       // text
	CreateTime      int             `json:"create_time"`      // int
	MusicID         int64           `json:"music_id"`         // bigint
	SourceID        int64           `json:"source_id"`        // bigint
	ShareURL        string          `json:"share_url"`        // text
	StatisticsID    int64           `json:"statistics_id"`    // bigint
	Status          json.RawMessage `json:"status"`           // json (use a custom JSON type or map for unmarshaling)
	TextExtra       json.RawMessage `json:"text_extra"`       // json
	IsTop           bool            `json:"is_top"`           // tinyint(1) (boolean)
	ShareInfo       json.RawMessage `json:"share_info"`       // json
	Duration        int             `json:"duration"`         // int
	ImageInfos      json.RawMessage `json:"image_infos"`      // json
	RiskInfos       json.RawMessage `json:"risk_infos"`       // json
	Position        string          `json:"position"`         // varchar(255)
	AuthorUserID    int64           `json:"author_user_id"`   // bigint
	PreventDownload bool            `json:"prevent_download"` // tinyint(1) (boolean)
	LongVideo       json.RawMessage `json:"long_video"`       // json
	AwemeControl    json.RawMessage `json:"aweme_control"`    // json
	Images          json.RawMessage `json:"images"`           // json
	SuggestWords    json.RawMessage `json:"suggest_words"`    // json
	VideoTag        json.RawMessage `json:"video_tag"`        // json
}

func CreateVideoFactory(sqlType string) *VideoModel {
	return &VideoModel{DB: model.UseDbConn(sqlType)}
}

func (v *VideoModel) GetMyVideo(Uid, pageNo, pageSize int64) (slice []model.Video, total int64, ok bool) {
	sql1 := `
		SELECT 
			a.aweme_id,
			a.video_desc,
			a.create_time,
			json_object(
				'id', tm.id, 
				'title', tm.title,
				'author', tm.author,
				'cover_medium', tm.cover_medium, 
				'cover_thumb', tm.cover_thumb,
				'play_url', tm.play_url,
				'duration', tm.duration,
				'user_count', tm.user_count,
				'owner_nickname', tm.owner_nickname,
				'is_original', tm.is_original,
				'owner_id', tm.owner_id
			) AS music,
			json_object(
				'play_addr', ts.play_addr,
				'cover', ts.cover,
				'poster', ts.poster,
				'height', ts.height,
				'width', ts.width,
				'ratio', ts.ratio,
				'use_static_cover', ts.use_static_cover,
				'duration', ts.duration,
				'horizontal_type', ts.horizontal_type
			) AS video,
			a.share_url,
			json_object(
				'admire_count', ts2.admire_count,
				'comment_count', ts2.comment_count,
				'digg_count', ts2.digg_count,
				'collect_count', ts2.collect_count,
				'play_count', ts2.play_count,
				'share_count', ts2.share_count
			) AS statistics,
			json_object(
				'uid', tu.uid,
				'short_id', tu.short_id,
				'unique_id', tu.unique_id,
				'gender', tu.gender,
				'user_age', tu.user_age,
				'nickname', tu.nickname,
				'country', tu.country,
				'province', tu.province,
				'district', tu.district,
				'city', tu.city,
				'signature', tu.signature,
				'ip_location', tu.ip_location,
				'birthday_hide_level', tu.birthday_hide_level,
				'can_show_group_card', tu.can_show_group_card,
				'aweme_count', tu.aweme_count,
				'total_favorited', tu.total_favorited,
				'favoriting_count', tu.favoriting_count,
				'follower_count', tu.follower_count,
				'following_count', tu.following_count,
				'forward_count', tu.forward_count,
				'public_collects_count', tu.public_collects_count,
				'mplatform_followers_count', tu.mplatform_followers_count,
				'max_follower_count', tu.max_follower_count,
				'follow_status', tu.follow_status,
				'follower_status', tu.follower_status,
				'follower_request_status', tu.follower_request_status,
				'cover_colour', tu.cover_colour,
				'cover_url', tu.cover_url,
				'white_cover_url', tu.white_cover_url,
				'share_info', tu.share_info,
				'commerce_info', tu.commerce_info,
				'commerce_user_info', tu.commerce_user_info,
				'commerce_user_level', tu.commerce_user_level,
				'card_entries', tu.card_entries,
				'avatar_small', tu.avatar_small,
				'avatar_large', tu.avatar_large
			) AS author,
			a.status,
			a.text_extra,
			a.is_top,
			a.share_info,
			a.duration,
			a.image_infos,
			a.risk_infos,
			a.position,
			a.author_user_id,
			a.prevent_download,
			a.long_video,
			a.aweme_control,
			a.images,
			a.suggest_words,
			a.video_tag
		FROM tb_videos AS a
		LEFT JOIN tb_music AS tm ON a.music_id = tm.id
		LEFT JOIN tb_source AS ts ON a.source_id = ts.id
		LEFT JOIN tb_statistics AS ts2 ON a.statistics_id = ts2.id
		LEFT JOIN tb_users AS tu ON a.author_user_id = tu.uid
		WHERE a.author_user_id = ? 
			AND JSON_EXTRACT(a.status, '$.private_status') = 0
		ORDER BY a.is_top DESC, a.create_time DESC
		LIMIT ? OFFSET ?;
			`

	sql2 := `
		SELECT COUNT(*)
		FROM tb_videos as a
		WHERE a.author_user_id = ? AND JSON_EXTRACT(status, '$.private_status') = 0
		`

	offset := pageNo * pageSize
	result1 := v.Raw(sql2, Uid).Count(&total)
	result2 := v.Raw(sql1, Uid, pageSize, offset).Find(&slice)
	if result1.Error != nil || result2.Error != nil {
		variable.ZapLog.Error("GetMyVideo SQL执行出错!")
		ok = false
		return
	}
	ok = true
	return
}

func (v *VideoModel) GetMyPrivateVideo(Uid, pageNo, pageSize int64) (slice []model.Video, total int64, ok bool) {
	sql1 := `
		SELECT 
			a.aweme_id,
			a.video_desc,
			a.create_time,
			json_object(
			'id', tm.id, 
			'title', tm.title,
			'author', tm.author,
			'cover_medium', tm.cover_medium, 
			'cover_thumb', tm.cover_thumb,
			'play_url', tm.play_url,
			'duration', tm.duration,
			'user_count', tm.user_count,
			'owner_nickname', tm.owner_nickname,
			'is_original', tm.is_original,
			'owner_id', tm.owner_id) AS music,
			json_object(
			'play_addr', ts.play_addr,
			'cover', ts.cover,
			'poster', ts.poster,
			'height', ts.height,
			'width', ts.width,
			'ratio', ts.ratio,
			'use_static_cover', ts.use_static_cover,
			'duration', ts.duration,
			'horizontal_type', ts.horizontal_type) AS video,
			a.share_url,
			json_object(
			'admire_count', ts2.admire_count,
			'comment_count', ts2.comment_count,
			'digg_count', ts2.digg_count,
			'collect_count', ts2.collect_count,
			'play_count', ts2.play_count,
			'share_count', ts2.share_count) AS statistics,
		    json_object(
			'uid', tu.uid,
			'short_id', tu.short_id,
			'unique_id', tu.unique_id,
			'gender', tu.gender,
			'user_age', tu.user_age,
			'nickname', tu.nickname,
			'country', tu.country,
			'province', tu.province,
			'district', tu.district,
			'city', tu.city,
			'signature', tu.signature,
			'ip_location', tu.ip_location,
			'birthday_hide_level', tu.birthday_hide_level,
			'can_show_group_card', tu.can_show_group_card,
			'aweme_count', tu.aweme_count,
			'total_favorited', tu.total_favorited,
			'favoriting_count', tu.favoriting_count,
			'follower_count', tu.follower_count,
			'following_count', tu.following_count,
			'forward_count', tu.forward_count,
			'public_collects_count', tu.public_collects_count,
			'mplatform_followers_count', tu.mplatform_followers_count,
			'max_follower_count', tu.max_follower_count,
			'follow_status', tu.follow_status,
			'follower_status', tu.follower_status,
			'follower_request_status', tu.follower_request_status,
			'cover_colour', tu.cover_colour,
			'cover_url', tu.cover_url,
			'white_cover_url', tu.white_cover_url,
			'share_info', tu.share_info,
			'commerce_info', tu.commerce_info,
			'commerce_user_info', tu.commerce_user_info,
			'commerce_user_level', tu.commerce_user_level,
			'card_entries', tu.card_entries,
			'avatar_small', tu.avatar_small,
			'avatar_large', tu.avatar_large
		   ) AS author,
			a.status,
			a.text_extra,
			a.is_top,
			a.share_info,
			a.duration,
			a.image_infos,
			a.risk_infos,
			a.position,
			a.author_user_id,
			a.prevent_download,
			a.long_video,
			a.aweme_control,
			a.images,
			a.suggest_words,
			a.video_tag
			FROM tb_videos as a
			LEFT JOIN tb_music as tm ON a.music_id = tm.id
			LEFT JOIN tb_source as ts ON a.source_id = ts.id
			LEFT JOIN tb_statistics as ts2 ON a.statistics_id = ts2.id
			LEFT JOIN tb_users AS tu ON a.author_user_id = tu.uid
			WHERE a.author_user_id = ? AND JSON_EXTRACT(status, '$.private_status') != 0
			ORDER BY a.create_time DESC
			LIMIT ? OFFSET ?;
			`

	sql2 := `
		SELECT COUNT(*)
		FROM tb_videos as a
		WHERE a.author_user_id = ? AND JSON_EXTRACT(status, '$.private_status') != 0
		`

	offset := pageNo * pageSize
	result1 := v.Raw(sql2, Uid).Count(&total)
	result2 := v.Raw(sql1, Uid, pageSize, offset).Find(&slice)
	if result1.Error != nil || result2.Error != nil {
		variable.ZapLog.Error("GetMyPrivateVideo SQL执行出错!")
		ok = false
		return
	}
	ok = true
	return
}

func (v *VideoModel) GetMyLikeVideo(Uid, pageNo, pageSize int64) (slice []model.Video, total int64, ok bool) {
	sql1 := `
		SELECT 
			a.aweme_id,
			a.video_desc as "desc",
			a.create_time,
			json_object(
				'id', tm.id, 
				'title', tm.title,
				'author', tm.author,
				'cover_medium', tm.cover_medium, 
				'cover_thumb', tm.cover_thumb,
				'play_url', tm.play_url,
				'duration', tm.duration,
				'user_count', tm.user_count,
				'owner_nickname', tm.owner_nickname,
				'is_original', tm.is_original,
				'owner_id', tm.owner_id
			) AS music,
			json_object(
				'play_addr', ts.play_addr,
				'cover', ts.cover,
				'poster', ts.poster,
				'height', ts.height,
				'width', ts.width,
				'ratio', ts.ratio,
				'use_static_cover', ts.use_static_cover,
				'duration', ts.duration,
				'horizontal_type', ts.horizontal_type
			) AS video,
			a.share_url,
			json_object(
				'admire_count', ts2.admire_count,
				'comment_count', ts2.comment_count,
				'digg_count', ts2.digg_count,
				'collect_count', ts2.collect_count,
				'play_count', ts2.play_count,
				'share_count', ts2.share_count
			) AS statistics,
		    json_object(
			'uid', tu.uid,
			'short_id', tu.short_id,
			'unique_id', tu.unique_id,
			'gender', tu.gender,
			'user_age', tu.user_age,
			'nickname', tu.nickname,
			'country', tu.country,
			'province', tu.province,
			'district', tu.district,
			'city', tu.city,
			'signature', tu.signature,
			'ip_location', tu.ip_location,
			'birthday_hide_level', tu.birthday_hide_level,
			'can_show_group_card', tu.can_show_group_card,
			'aweme_count', tu.aweme_count,
			'total_favorited', tu.total_favorited,
			'favoriting_count', tu.favoriting_count,
			'follower_count', tu.follower_count,
			'following_count', tu.following_count,
			'forward_count', tu.forward_count,
			'public_collects_count', tu.public_collects_count,
			'mplatform_followers_count', tu.mplatform_followers_count,
			'max_follower_count', tu.max_follower_count,
			'follow_status', tu.follow_status,
			'follower_status', tu.follower_status,
			'follower_request_status', tu.follower_request_status,
			'cover_colour', tu.cover_colour,
			'cover_url', tu.cover_url,
			'white_cover_url', tu.white_cover_url,
			'share_info', tu.share_info,
			'commerce_info', tu.commerce_info,
			'commerce_user_info', tu.commerce_user_info,
			'commerce_user_level', tu.commerce_user_level,
			'card_entries', tu.card_entries,
			'avatar_small', tu.avatar_small,
			'avatar_large', tu.avatar_large
		   ) AS author,
			a.status,
			a.text_extra,
			a.is_top,
			a.share_info,
			a.duration,
			a.image_infos,
			a.risk_infos,
			a.position,
			a.author_user_id,
			a.prevent_download,
			a.long_video,
			a.aweme_control,
			a.images,
			a.suggest_words,
			a.video_tag
		FROM tb_videos AS a
		LEFT JOIN tb_music AS tm ON a.music_id = tm.id
		LEFT JOIN tb_source AS ts ON a.source_id = ts.id
		LEFT JOIN tb_statistics AS ts2 ON a.statistics_id = ts2.id
		LEFT JOIN tb_users AS tu ON a.author_user_id = tu.uid
		WHERE a.aweme_id IN (
		    SELECT td.aweme_id
		    FROM tb_diggs as td
		    WHERE td.uid = ?  -- 替换为用户的 uid
		)  
		AND JSON_EXTRACT(a.status, '$.private_status') = 0  -- 过滤条件：视频为公开状态
		ORDER BY a.create_time DESC
		LIMIT ? OFFSET ?;
		`

	sql2 := `
		SELECT COUNT(*)
		FROM tb_videos as tv
		WHERE tv.aweme_id IN (
		    SELECT td.aweme_id
		    FROM tb_diggs as td
		    WHERE td.uid = ?  -- 替换为用户的 uid
		)  
		AND JSON_EXTRACT(tv.status, '$.private_status') = 0  -- 过滤条件：视频为公开状态
		`

	offset := pageNo * pageSize
	result1 := v.Raw(sql2, Uid).Count(&total)
	result2 := v.Raw(sql1, Uid, pageSize, offset).Find(&slice)
	if result1.Error != nil || result2.Error != nil {
		variable.ZapLog.Error("GetMyLikeVideo SQL执行出错!")
		ok = false
		return
	}
	ok = true
	return
}

func (v *VideoModel) GetMyCollectVideo(Uid, pageNo, pageSize int64) (slice []model.Video, total int64, ok bool) {
	sql1 := `
		SELECT 
			a.aweme_id,
			a.video_desc,
			a.create_time,
			json_object(
				'id', tm.id, 
				'title', tm.title,
				'author', tm.author,
				'cover_medium', tm.cover_medium, 
				'cover_thumb', tm.cover_thumb,
				'play_url', tm.play_url,
				'duration', tm.duration,
				'user_count', tm.user_count,
				'owner_nickname', tm.owner_nickname,
				'is_original', tm.is_original,
				'owner_id', tm.owner_id
			) AS music,
			json_object(
				'play_addr', ts.play_addr,
				'cover', ts.cover,
				'poster', ts.poster,
				'height', ts.height,
				'width', ts.width,
				'ratio', ts.ratio,
				'use_static_cover', ts.use_static_cover,
				'duration', ts.duration,
				'horizontal_type', ts.horizontal_type
			) AS video,
			a.share_url,
			json_object(
				'admire_count', ts2.admire_count,
				'comment_count', ts2.comment_count,
				'digg_count', ts2.digg_count,
				'collect_count', ts2.collect_count,
				'play_count', ts2.play_count,
				'share_count', ts2.share_count
			) AS statistics,
		    json_object(
			'uid', tu.uid,
			'short_id', tu.short_id,
			'unique_id', tu.unique_id,
			'gender', tu.gender,
			'user_age', tu.user_age,
			'nickname', tu.nickname,
			'country', tu.country,
			'province', tu.province,
			'district', tu.district,
			'city', tu.city,
			'signature', tu.signature,
			'ip_location', tu.ip_location,
			'birthday_hide_level', tu.birthday_hide_level,
			'can_show_group_card', tu.can_show_group_card,
			'aweme_count', tu.aweme_count,
			'total_favorited', tu.total_favorited,
			'favoriting_count', tu.favoriting_count,
			'follower_count', tu.follower_count,
			'following_count', tu.following_count,
			'forward_count', tu.forward_count,
			'public_collects_count', tu.public_collects_count,
			'mplatform_followers_count', tu.mplatform_followers_count,
			'max_follower_count', tu.max_follower_count,
			'follow_status', tu.follow_status,
			'follower_status', tu.follower_status,
			'follower_request_status', tu.follower_request_status,
			'cover_colour', tu.cover_colour,
			'cover_url', tu.cover_url,
			'white_cover_url', tu.white_cover_url,
			'share_info', tu.share_info,
			'commerce_info', tu.commerce_info,
			'commerce_user_info', tu.commerce_user_info,
			'commerce_user_level', tu.commerce_user_level,
			'card_entries', tu.card_entries,
			'avatar_small', tu.avatar_small,
			'avatar_large', tu.avatar_large
		   ) AS author,
			a.status,
			a.text_extra,
			a.is_top,
			a.share_info,
			a.duration,
			a.image_infos,
			a.risk_infos,
			a.position,
			a.author_user_id,
			a.prevent_download,
			a.long_video,
			a.aweme_control,
			a.images,
			a.suggest_words,
			a.video_tag
		FROM tb_videos AS a
		LEFT JOIN tb_music AS tm ON a.music_id = tm.id
		LEFT JOIN tb_source AS ts ON a.source_id = ts.id
		LEFT JOIN tb_statistics AS ts2 ON a.statistics_id = ts2.id
		LEFT JOIN tb_users AS tu ON a.author_user_id = tu.uid
		WHERE a.aweme_id IN (
		    SELECT td.aweme_id
		    FROM tb_collects as td
		    WHERE td.uid = ?  -- 替换为用户的 uid
		)  
		AND JSON_EXTRACT(a.status, '$.private_status') = 0  -- 过滤条件：视频为公开状态
		ORDER BY a.create_time DESC
		LIMIT ? OFFSET ?;
		`

	sql2 := `
		SELECT COUNT(*)
		FROM tb_videos as tv
		WHERE tv.aweme_id IN (
		    SELECT td.aweme_id
		    FROM tb_diggs as td
		    WHERE td.uid = ?  -- 替换为用户的 uid
		)  
		AND JSON_EXTRACT(tv.status, '$.private_status') = 0  -- 过滤条件：视频为公开状态
		`

	offset := pageNo * pageSize
	result1 := v.Raw(sql2, Uid).Count(&total)
	result2 := v.Raw(sql1, Uid, pageSize, offset).Find(&slice)
	if result1.Error != nil || result2.Error != nil {
		variable.ZapLog.Error("GetMyCollectVideo SQL执行出错!")
		ok = false
		return
	}
	ok = true
	return
}

func (v *VideoModel) GetMyHistoryVideo(Uid, pageNo, pageSize int64) (slice []model.Video, total int64, ok bool) {
	sql1 := `
		SELECT 
			a.aweme_id,
			a.video_desc as "desc",
			a.create_time,
			json_object(
				'id', tm.id, 
				'title', tm.title,
				'author', tm.author,
				'cover_medium', tm.cover_medium, 
				'cover_thumb', tm.cover_thumb,
				'play_url', tm.play_url,
				'duration', tm.duration,
				'user_count', tm.user_count,
				'owner_nickname', tm.owner_nickname,
				'is_original', tm.is_original,
				'owner_id', tm.owner_id
			) AS music,
			json_object(
				'play_addr', ts.play_addr,
				'cover', ts.cover,
				'poster', ts.poster,
				'height', ts.height,
				'width', ts.width,
				'ratio', ts.ratio,
				'use_static_cover', ts.use_static_cover,
				'duration', ts.duration,
				'horizontal_type', ts.horizontal_type
			) AS video,
			a.share_url,
			json_object(
				'admire_count', ts2.admire_count,
				'comment_count', ts2.comment_count,
				'digg_count', ts2.digg_count,
				'collect_count', ts2.collect_count,
				'play_count', ts2.play_count,
				'share_count', ts2.share_count
			) AS statistics,
		    json_object(
			'uid', tu.uid,
			'short_id', tu.short_id,
			'unique_id', tu.unique_id,
			'gender', tu.gender,
			'user_age', tu.user_age,
			'nickname', tu.nickname,
			'country', tu.country,
			'province', tu.province,
			'district', tu.district,
			'city', tu.city,
			'signature', tu.signature,
			'ip_location', tu.ip_location,
			'birthday_hide_level', tu.birthday_hide_level,
			'can_show_group_card', tu.can_show_group_card,
			'aweme_count', tu.aweme_count,
			'total_favorited', tu.total_favorited,
			'favoriting_count', tu.favoriting_count,
			'follower_count', tu.follower_count,
			'following_count', tu.following_count,
			'forward_count', tu.forward_count,
			'public_collects_count', tu.public_collects_count,
			'mplatform_followers_count', tu.mplatform_followers_count,
			'max_follower_count', tu.max_follower_count,
			'follow_status', tu.follow_status,
			'follower_status', tu.follower_status,
			'follower_request_status', tu.follower_request_status,
			'cover_colour', tu.cover_colour,
			'cover_url', tu.cover_url,
			'white_cover_url', tu.white_cover_url,
			'share_info', tu.share_info,
			'commerce_info', tu.commerce_info,
			'commerce_user_info', tu.commerce_user_info,
			'commerce_user_level', tu.commerce_user_level,
			'card_entries', tu.card_entries,
			'avatar_small', tu.avatar_small,
			'avatar_large', tu.avatar_large
		   ) AS author,
			a.status,
			a.text_extra,
			a.is_top,
			a.share_info,
			a.duration,
			a.image_infos,
			a.risk_infos,
			a.position,
			a.author_user_id,
			a.prevent_download,
			a.long_video,
			a.aweme_control,
			a.images,
			a.suggest_words,
			a.video_tag
		FROM tb_videos AS a
		LEFT JOIN tb_music AS tm ON a.music_id = tm.id
		LEFT JOIN tb_source AS ts ON a.source_id = ts.id
		LEFT JOIN tb_statistics AS ts2 ON a.statistics_id = ts2.id
		LEFT JOIN tb_users AS tu ON a.author_user_id = tu.uid
		WHERE a.aweme_id IN (
		    SELECT th.aweme_id
		    FROM tb_history as th
		    WHERE th.uid = ?  -- 替换为用户的 uid
		)  
		AND JSON_EXTRACT(a.status, '$.private_status') = 0  -- 过滤条件：视频为公开状态
		ORDER BY a.create_time DESC
		LIMIT ? OFFSET ?;
		`

	sql2 := `
		SELECT COUNT(*)
		FROM tb_videos as tv
		WHERE tv.aweme_id IN (
		    SELECT td.aweme_id
		    FROM tb_history as td
		    WHERE td.uid = ?  -- 替换为用户的 uid
		)  
		AND JSON_EXTRACT(tv.status, '$.private_status') = 0  -- 过滤条件：视频为公开状态
		`

	offset := pageNo * pageSize
	result1 := v.Raw(sql2, Uid).Count(&total)
	result2 := v.Raw(sql1, Uid, pageSize, offset).Find(&slice)
	if result1.Error != nil || result2.Error != nil {
		variable.ZapLog.Error("GetMyHistoryVideo SQL执行出错!")
		ok = false
		return
	}
	ok = true
	return
}

func (v *VideoModel) GetMyHistoryOther(Uid, pageNo, pageSize int64) {
	return
}

func (v *VideoModel) GetVideoRecommended(Uid, start, pageSize int64) (slice []model.Video, total int64, ok bool) {
	sql1 := `
		SELECT 
			a.aweme_id,
			a.video_desc,
			a.create_time,
			json_object(
				'id', tm.id, 
				'title', tm.title,
				'author', tm.author,
				'cover_medium', tm.cover_medium, 
				'cover_thumb', tm.cover_thumb,
				'play_url', tm.play_url,
				'duration', tm.duration,
				'user_count', tm.user_count,
				'owner_nickname', tm.owner_nickname,
				'is_original', tm.is_original,
				'owner_id', tm.owner_id
			) AS music,
			json_object(
				'play_addr', ts.play_addr,
				'cover', ts.cover,
				'poster', ts.poster,
				'height', ts.height,
				'width', ts.width,
				'ratio', ts.ratio,
				'use_static_cover', ts.use_static_cover,
				'duration', ts.duration,
				'horizontal_type', ts.horizontal_type
			) AS video,
			a.share_url,
			json_object(
				'admire_count', ts2.admire_count,
				'comment_count', ts2.comment_count,
				'digg_count', ts2.digg_count,
				'collect_count', ts2.collect_count,
				'play_count', ts2.play_count,
				'share_count', ts2.share_count
			) AS statistics,
			json_object(
				'uid', tu.uid,
				'short_id', tu.short_id,
				'unique_id', tu.unique_id,
				'gender', tu.gender,
				'user_age', tu.user_age,
				'nickname', tu.nickname,
				'country', tu.country,
				'province', tu.province,
				'district', tu.district,
				'city', tu.city,
				'signature', tu.signature,
				'ip_location', tu.ip_location,
				'birthday_hide_level', tu.birthday_hide_level,
				'can_show_group_card', tu.can_show_group_card,
				'aweme_count', tu.aweme_count,
				'total_favorited', tu.total_favorited,
				'favoriting_count', tu.favoriting_count,
				'follower_count', tu.follower_count,
				'following_count', tu.following_count,
				'forward_count', tu.forward_count,
				'public_collects_count', tu.public_collects_count,
				'mplatform_followers_count', tu.mplatform_followers_count,
				'max_follower_count', tu.max_follower_count,
				'follow_status', tu.follow_status,
				'follower_status', tu.follower_status,
				'follower_request_status', tu.follower_request_status,
				'cover_colour', tu.cover_colour,
				'cover_url', tu.cover_url,
				'white_cover_url', tu.white_cover_url,
				'share_info', tu.share_info,
				'commerce_info', tu.commerce_info,
				'commerce_user_info', tu.commerce_user_info,
				'commerce_user_level', tu.commerce_user_level,
				'card_entries', tu.card_entries,
				'avatar_small', tu.avatar_small,
				'avatar_large', tu.avatar_large
			) AS author,
			a.status,
			a.text_extra,
			a.is_top,
			a.share_info,
			a.duration,
			a.image_infos,
			a.risk_infos,
			a.position,
			a.author_user_id,
			a.prevent_download,
			a.long_video,
			a.aweme_control,
			a.images,
			a.suggest_words,
			a.video_tag
		FROM tb_videos AS a
		LEFT JOIN tb_music AS tm ON a.music_id = tm.id
		LEFT JOIN tb_source AS ts ON a.source_id = ts.id
		LEFT JOIN tb_statistics AS ts2 ON a.statistics_id = ts2.id
		LEFT JOIN tb_users AS tu ON a.author_user_id = tu.uid
		WHERE JSON_EXTRACT(a.status, '$.private_status') = 0
		LIMIT ? OFFSET ?;
			`

	sql2 := `
		SELECT COUNT(*)
		FROM tb_videos as a
		WHERE JSON_EXTRACT(status, '$.private_status') = 0
		`

	//offset := start * pageSize
	result1 := v.Raw(sql2).Count(&total)
	if total <= pageSize {
		variable.ZapLog.Error("GetVideoRecommended 可用数据少于 pageSize!")
		return
	}
	rand.Seed(uint64(time.Now().UnixNano()))
	randomStart := rand.Intn(int(total - pageSize + 1))
	result2 := v.Raw(sql1, pageSize, randomStart).Find(&slice)

	if result1.Error != nil || result2.Error != nil {
		variable.ZapLog.Error("GetVideoRecommended SQL执行出错!")
		ok = false
		return
	}
	ok = true
	return
}

func (v *VideoModel) GetLongVideoRecommended(Uid, PageNo, pageSize int64) (slice []model.Video, total int64, ok bool) {
	sql1 := `
		SELECT 
			a.aweme_id,
			a.video_desc,
			a.create_time,
			json_object(
				'id', tm.id, 
				'title', tm.title,
				'author', tm.author,
				'cover_medium', tm.cover_medium, 
				'cover_thumb', tm.cover_thumb,
				'play_url', tm.play_url,
				'duration', tm.duration,
				'user_count', tm.user_count,
				'owner_nickname', tm.owner_nickname,
				'is_original', tm.is_original,
				'owner_id', tm.owner_id
			) AS music,
			json_object(
				'play_addr', ts.play_addr,
				'cover', ts.cover,
				'poster', ts.poster,
				'height', ts.height,
				'width', ts.width,
				'ratio', ts.ratio,
				'use_static_cover', ts.use_static_cover,
				'duration', ts.duration,
				'horizontal_type', ts.horizontal_type
			) AS video,
			a.share_url,
			json_object(
				'admire_count', ts2.admire_count,
				'comment_count', ts2.comment_count,
				'digg_count', ts2.digg_count,
				'collect_count', ts2.collect_count,
				'play_count', ts2.play_count,
				'share_count', ts2.share_count
			) AS statistics,
			json_object(
				'uid', tu.uid,
				'short_id', tu.short_id,
				'unique_id', tu.unique_id,
				'gender', tu.gender,
				'user_age', tu.user_age,
				'nickname', tu.nickname,
				'country', tu.country,
				'province', tu.province,
				'district', tu.district,
				'city', tu.city,
				'signature', tu.signature,
				'ip_location', tu.ip_location,
				'birthday_hide_level', tu.birthday_hide_level,
				'can_show_group_card', tu.can_show_group_card,
				'aweme_count', tu.aweme_count,
				'total_favorited', tu.total_favorited,
				'favoriting_count', tu.favoriting_count,
				'follower_count', tu.follower_count,
				'following_count', tu.following_count,
				'forward_count', tu.forward_count,
				'public_collects_count', tu.public_collects_count,
				'mplatform_followers_count', tu.mplatform_followers_count,
				'max_follower_count', tu.max_follower_count,
				'follow_status', tu.follow_status,
				'follower_status', tu.follower_status,
				'follower_request_status', tu.follower_request_status,
				'cover_colour', tu.cover_colour,
				'cover_url', tu.cover_url,
				'white_cover_url', tu.white_cover_url,
				'share_info', tu.share_info,
				'commerce_info', tu.commerce_info,
				'commerce_user_info', tu.commerce_user_info,
				'commerce_user_level', tu.commerce_user_level,
				'card_entries', tu.card_entries,
				'avatar_small', tu.avatar_small,
				'avatar_large', tu.avatar_large
			) AS author,
			a.status,
			a.text_extra,
			a.is_top,
			a.share_info,
			a.duration,
			a.image_infos,
			a.risk_infos,
			a.position,
			a.author_user_id,
			a.prevent_download,
			a.long_video,
			a.aweme_control,
			a.images,
			a.suggest_words,
			a.video_tag
		FROM tb_videos AS a
		LEFT JOIN tb_music AS tm ON a.music_id = tm.id
		LEFT JOIN tb_source AS ts ON a.source_id = ts.id
		LEFT JOIN tb_statistics AS ts2 ON a.statistics_id = ts2.id
		LEFT JOIN tb_users AS tu ON a.author_user_id = tu.uid
		WHERE JSON_EXTRACT(a.status, '$.private_status') = 0
		LIMIT ? OFFSET ?;
			`

	sql2 := `
		SELECT COUNT(*)
		FROM tb_videos as a
		WHERE JSON_EXTRACT(status, '$.private_status') = 0
		`

	offset := PageNo * pageSize
	result1 := v.Raw(sql2).Count(&total)
	result2 := v.Raw(sql1, pageSize, offset).Find(&slice)

	if result1.Error != nil || result2.Error != nil {
		variable.ZapLog.Error("GetLongVideoRecommended SQL执行出错!")
		ok = false
		return
	}
	ok = true
	return
}

func (v *VideoModel) GetUserVideoList(Uid int64) (slice []model.Video, ok bool) {
	sql1 := `
		SELECT 
			a.aweme_id,
			a.video_desc,
			a.create_time,
			json_object(
				'id', tm.id, 
				'title', tm.title,
				'author', tm.author,
				'cover_medium', tm.cover_medium, 
				'cover_thumb', tm.cover_thumb,
				'play_url', tm.play_url,
				'duration', tm.duration,
				'user_count', tm.user_count,
				'owner_nickname', tm.owner_nickname,
				'is_original', tm.is_original,
				'owner_id', tm.owner_id
			) AS music,
			json_object(
				'play_addr', ts.play_addr,
				'cover', ts.cover,
				'poster', ts.poster,
				'height', ts.height,
				'width', ts.width,
				'ratio', ts.ratio,
				'use_static_cover', ts.use_static_cover,
				'duration', ts.duration,
				'horizontal_type', ts.horizontal_type
			) AS video,
			a.share_url,
			json_object(
				'admire_count', ts2.admire_count,
				'comment_count', ts2.comment_count,
				'digg_count', ts2.digg_count,
				'collect_count', ts2.collect_count,
				'play_count', ts2.play_count,
				'share_count', ts2.share_count
			) AS statistics,
			json_object(
				'uid', tu.uid,
				'short_id', tu.short_id,
				'unique_id', tu.unique_id,
				'gender', tu.gender,
				'user_age', tu.user_age,
				'nickname', tu.nickname,
				'country', tu.country,
				'province', tu.province,
				'district', tu.district,
				'city', tu.city,
				'signature', tu.signature,
				'ip_location', tu.ip_location,
				'birthday_hide_level', tu.birthday_hide_level,
				'can_show_group_card', tu.can_show_group_card,
				'aweme_count', tu.aweme_count,
				'total_favorited', tu.total_favorited,
				'favoriting_count', tu.favoriting_count,
				'follower_count', tu.follower_count,
				'following_count', tu.following_count,
				'forward_count', tu.forward_count,
				'public_collects_count', tu.public_collects_count,
				'mplatform_followers_count', tu.mplatform_followers_count,
				'max_follower_count', tu.max_follower_count,
				'follow_status', tu.follow_status,
				'follower_status', tu.follower_status,
				'follower_request_status', tu.follower_request_status,
				'cover_colour', tu.cover_colour,
				'cover_url', tu.cover_url,
				'white_cover_url', tu.white_cover_url,
				'share_info', tu.share_info,
				'commerce_info', tu.commerce_info,
				'commerce_user_info', tu.commerce_user_info,
				'commerce_user_level', tu.commerce_user_level,
				'card_entries', tu.card_entries,
				'avatar_small', tu.avatar_small,
				'avatar_large', tu.avatar_large
			) AS author,
			a.status,
			a.text_extra,
			a.is_top,
			a.share_info,
			a.duration,
			a.image_infos,
			a.risk_infos,
			a.position,
			a.author_user_id,
			a.prevent_download,
			a.long_video,
			a.aweme_control,
			a.images,
			a.suggest_words,
			a.video_tag
		FROM tb_videos AS a
		LEFT JOIN tb_music AS tm ON a.music_id = tm.id
		LEFT JOIN tb_source AS ts ON a.source_id = ts.id
		LEFT JOIN tb_statistics AS ts2 ON a.statistics_id = ts2.id
		LEFT JOIN tb_users AS tu ON a.author_user_id = tu.uid
		WHERE a.author_user_id = ? 
			AND JSON_EXTRACT(a.status, '$.private_status') = 0
		ORDER BY a.is_top DESC, a.create_time DESC;`
	result := v.Raw(sql1, Uid).Find(&slice)
	if result.Error != nil {
		variable.ZapLog.Error("GetUserVideoList SQL执行出错!", zap.Error(result.Error))
		ok = false
		return
	}
	ok = true
	return
}
