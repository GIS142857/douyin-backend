package post

import (
	"github.com/goccy/go-json"
	"gorm.io/gorm"
)

type Post struct {
	*gorm.DB  `gorm:"-" json:"-"`
	ID        string          `json:"id"`         // varchar(100)
	ModelType string          `json:"model_type"` // varchar(100)
	NoteCard  json.RawMessage `json:"note_card"`  // json
}
