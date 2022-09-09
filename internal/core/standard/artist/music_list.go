package std_artist

import (
	"github.com/Mistsink/kuwo-api/internal/core/standard"
)

type MusicListResp struct {
	standard.CommonRes
	Result MusicListRes `json:"result" yzh:"result"`
}

type MusicListRes struct {
	TotalCnt int                `json:"total_cnt" yzh:"total_cnt"`
	List     []standard.InfoRes `json:"list" yzh:"list"`
}
