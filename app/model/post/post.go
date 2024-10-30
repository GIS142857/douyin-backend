package post

import (
	"douyin-backend/app/model"
	"github.com/goccy/go-json"
	"gorm.io/gorm"
)

type PostModel struct {
	*gorm.DB  `gorm:"-" json:"-"`
	ID        string          `json:"id"`         // varchar(100)
	ModelType string          `json:"model_type"` // varchar(100)
	NoteCard  json.RawMessage `json:"note_card"`  // json
}

func CreatePostFactory(sqlType string) *PostModel {
	return &PostModel{DB: model.UseDbConn(sqlType)}
}

func (u *PostModel) GetPostRecommended(uid, pageNo, pageSize int64) (slice []Post, total int64) {
	sql1 := `
		SELECT *
		from tb_posts as tu
		LIMIT ? OFFSET ?;`
	sql2 := `
		SELECT COUNT(*)
		FROM tb_posts as a;
		`

	offset := pageNo * pageSize
	u.Raw(sql2).Count(&total)
	u.Raw(sql1, pageSize, offset).Find(&slice)
	return
}
