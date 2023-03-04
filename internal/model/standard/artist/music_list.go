package std_artist

import (
	"github.com/Mistsink/kuwo-api/internal/model/standard"
)

type MusicListResp struct {
	standard.CommonResp
	Result MusicListRes `json:"result" yzh:"result"`
}

type MusicListRes struct {
	standard.PlatformTag
	TotalCnt int                `json:"total_cnt" yzh:"total_cnt"`
	List     []standard.InfoRes `json:"list" yzh:"list"`
}
