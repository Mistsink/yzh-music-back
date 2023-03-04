package std_rank

import (
	"github.com/Mistsink/kuwo-api/internal/model/standard"
)

type PlDetailResp struct {
	standard.CommonResp
	Result PlDetailRes `json:"result" yzh:"result"`
}

type PlDetailRes struct {
	ImgUrl     string             `json:"pl_img_url" yzh:"pl_img_url"`
	TotalCnt   int                `json:"total_cnt" yzh:"total_cnt"`
	UpdateTime string             `json:"update_time" yzh:"up_time"`
	MusicList  []standard.InfoRes `json:"list" yzh:"list"`
}
