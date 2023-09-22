package model

import (
	"ginweb/utils/errmsg"
	"time"
)

type Leaver struct {
	ID        uint64     `json:"id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Content   string     `gorm:"type:longtext" json:"content"`
}

func CreateLeaver(data *Leaver) int {
	if data == nil {
		return errmsg.ERROR
	}

	err := db.Create(data).Error
	if err != nil {
		return errmsg.ERROR
	}

	return errmsg.SUCCESS
}
