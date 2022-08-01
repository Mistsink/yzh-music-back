package std_rank

import (
	"github.com/Mistsink/kuwo-api/internal/core/standard"
	std_music "github.com/Mistsink/kuwo-api/internal/core/standard/music"
)

type PlDetailResp struct {
	standard.CommonRes
	Result PlDetailRes `json:"result" yzh:"result"`
}

type PlDetailRes struct {
	ImgUrl     string              `json:"pl_img_url" yzh:"pl_img_url"`
	TotalCnt   int                 `json:"total_cnt" yzh:"total_cnt"`
	UpdateTime string              `json:"update_time" yzh:"up_time"`
	MusicList  []std_music.InfoRes `json:"list" yzh:"list"`
}
