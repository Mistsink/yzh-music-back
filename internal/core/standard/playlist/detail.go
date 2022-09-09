package std_playlist

import (
	"github.com/Mistsink/kuwo-api/internal/core/standard"
)

type DetailResp struct {
	standard.CommonRes
	Result DetailRes `json:"result" yzh:"result"`
}

type DetailRes struct {
	ImgUrl     string             `json:"img_url" yzh:"img_url"`
	UserImgUrl string             `json:"user_img" yzh:"user_img_url"`
	UserName   string             `json:"user_name" yzh:"user_name"`
	Img700     string             `json:"img_700" yzh:"img_700"`
	Img300     string             `json:"img_300" yzh:"img_300"`
	Img500     string             `json:"img_500" yzh:"img_500"`
	TotalCnt   int                `json:"total_cnt" yzh:"total_cnt"`
	Name       string             `json:"name" yzh:"name"`
	ListenCnt  int                `json:"listen_cnt" yzh:"lis_cnt"`
	ID         int                `json:"id" yzh:"id"`
	Tag        string             `json:"tag" yzh:"tag"`
	MusicList  []standard.InfoRes `json:"list" yzh:"list"`
}
