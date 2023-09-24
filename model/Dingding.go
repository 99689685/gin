package model

import (
	"fmt"
	"github.com/goccy/go-json"
	"io/ioutil"
	"net/http"
)

// DingTalkUser 钉钉接口返回的用户信息数据结构
type DingTalkUser struct {
	UserID     string `json:"userid"`
	Name       string `json:"name"`
	Mobile     string `json:"mobile"`
	Department []int  `json:"department"`
}

// DingTalkUserList 钉钉接口返回的用户列表数据结构
type DingTalkUserList struct {
	Result   DingTalkUser `json:"result"`
	HasMore  bool         `json:"has_more"`
	NextPage int          `json:"next_page"`
}

type AccessTokenResponse struct {
	Errcode      int    `json:"errcode"`
	Errmsg       string `json:"errmsg"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

func GetDingAccessToken(appKey, appSecret string) (string, error) {
	url := fmt.Sprintf("https://oapi.dingtalk.com/gettoken?appkey=%s&appsecret=%s", appKey, appSecret)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result AccessTokenResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	if result.Errcode != 0 {
		return "", fmt.Errorf("获取 ACCESS_TOKEN 失败：%s", result.Errmsg)
	}

	return result.AccessToken, nil
}
