package std_playlist

import (
	"github.com/Mistsink/kuwo-api/internal/model/standard"
)

type DefaultResp struct {
	standard.CommonResp
	Result DefaultRes `json:"data"`
}

type DefaultRes struct {
	TotalCnt int        `json:"total_cnt" yzh:"total_cnt"`
	Data     []PlDetail `json:"data" yzh:"list"`
	N        int        `json:"n" yzh:"n"`
	P        int        `json:"p" yzh:"p"`
}

type PlDetail struct {
	standard.PlatformTag
	ImgUrl     string `json:"img_url" yzh:"pld_img_url"`
	UserName   string `json:"user_name" yzh:"pld_user_name"`
	TotalCnt   int    `json:"total_cnt" yzh:"pld_total_cnt"`
	Name       string `json:"name" yzh:"pld_name"`
	ListenCnt  int    `json:"listen_cnt" yzh:"pld_lis_cnt"`
	ID         int    `json:"id" yzh:"pld_id"`
	UpdateTime int    `json:"netease_update_time" yzh:"pld_netease_update_time"`
}
