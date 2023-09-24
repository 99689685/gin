package model

import (
	"bytes"
	"ginweb/utils"
	"ginweb/utils/errmsg"
	"github.com/goccy/go-json"
	"net/http"
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
	err = sendDingTalkMessage("新的留言信息", data.Content)
	if err != nil {
		return 0
	}
	return errmsg.SUCCESS
}

func sendDingTalkMessage(title, message string) error {
	url := utils.RobotWebhook
	data := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"title": title,
			"text":  message,
		},
	}
	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(payload))
	if err != nil {
		return err
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			// 错误处理
		}
	}()

	return nil
}
