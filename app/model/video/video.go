package video

import "douyin-backend/app/model"

type VideoModel struct {
	model.BaseModel
	AuthorId     int64  `json:"author_id"`
	Title        string `json:"title"`
	CoverUrl     string `json:"cover_url"`
	FileName     string `json:"file_name"`
	PlayUrl      string `json:"play_url"`
	Duration     string `json:"duration"`
	Category     int    `json:"category"`
	CategoryName string `json:"category_name"`
	FavoriteCnt  int64  `json:"favorite_cnt"`
	CommentCnt   int64  `json:"comment_cnt"`
	StarCnt      int64  `json:"star_cnt"`
	ShareCnt     int64  `json:"share_cnt"`
	DeleteAt     string `json:"delete_at"`
	FileSize     int64  `json:"file_size"`
}

func CreateShortVideoFactory(sqlType string) *VideoModel {
	return &VideoModel{BaseModel: model.BaseModel{DB: model.UseDbConn(sqlType)}}
}

func (v *VideoModel) GetVideo(pageNo int64, pageSize int64) (data Video) {
	sql := `
		SELECT *
		FROM tb_videos as a
		WHERE  a.id= ?
	`
	_ = v.Raw(sql, pageNo, pageSize).Find(&data)
	return
}
