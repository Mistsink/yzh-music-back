package rank

import (
	"github.com/Mistsink/kuwo-api/internal/core/kuwo"
	"github.com/Mistsink/kuwo-api/internal/core/kuwo/music"
)

type PlDetailResp struct {
	kuwo.CommonRes
	Data PlDetailRes `json:"data" yzh:"result"`
}

type PlDetailRes struct {
	Img       string                `json:"img" yzh:"pl_img_url"`
	Num       string                `json:"num" yzh:"total_cnt"`
	Pub       string                `json:"pub" yzh:"up_time"`
	MusicList []music.InfoData[int] `json:"musicList" yzh:"list"`
}
