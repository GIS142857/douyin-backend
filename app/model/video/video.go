package video

import (
	"douyin-backend/app/global/variable"
	"douyin-backend/app/model"
	"encoding/json"
	"gorm.io/gorm"
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

func CreateShortVideoFactory(sqlType string) *VideoModel {
	return &VideoModel{DB: model.UseDbConn(sqlType)}
}

func (v *VideoModel) GetMyVideo(Uid, pageNo, pageSize int64) (slice []Video, total int64) {
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
			WHERE a.author_user_id = ?
			ORDER BY a.is_top DESC, a.create_time DESC
			LIMIT ? OFFSET ?;
			`

	sql2 := `
		SELECT COUNT(*)
		FROM tb_videos as a
		WHERE a.author_user_id = ?
		`

	offset := pageNo * pageSize
	v.Raw(sql2, Uid).Count(&total)
	v.Raw(sql1, Uid, pageSize, offset).Find(&slice)

	if len(slice) > 0 {
		return
	} else {
		variable.ZapLog.Error("GetMyVideo 查询出错!")
		return
	}
}
