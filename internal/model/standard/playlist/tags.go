package std_playlist

import (
	"github.com/Mistsink/kuwo-api/internal/model/standard"
)

type TagResp struct {
	standard.CommonResp
	Result []TagRes `json:"result" yzh:"result"`
}

type TagRes struct {
	standard.PlatformTag
	Data []TagItem `json:"list" yzh:"list"`
	Name string    `json:"name" yzh:"name"`
	ID   string    `json:"fake_id" yzh:"fake_id"`
}

type TagItem struct {
	Img  string `json:"img_url" yzh:"tag_img_url"`
	Name string `json:"name" yzh:"tag_name"`
	ID   string `json:"id" yzh:"tag_id"`
}
