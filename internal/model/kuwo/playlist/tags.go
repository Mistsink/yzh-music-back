package playlist

import (
	"github.com/Mistsink/kuwo-api/internal/model/kuwo"
)

type TagResp struct {
	kuwo.CommonRes
	Data []TagRes `json:"data" yzh:"result"`
}

type TagRes struct {
	Img     string    `json:"img"`
	Mdigest string    `json:"mdigest"`
	Data    []TagItem `json:"data" yzh:"list"`
	Name    string    `json:"name" yzh:"name"`
	ID      string    `json:"id" yzh:"fake_id"`
	Type    string    `json:"type"`
	Img1    string    `json:"img1"`
}

type TagItem struct {
	Extend string `json:"extend"`
	Img    string `json:"img" yzh:"tag_img_url"`
	Digest string `json:"digest"`
	Name   string `json:"name" yzh:"tag_name"`
	Isnew  string `json:"isnew"`
	ID     string `json:"id" yzh:"tag_id"`
}
