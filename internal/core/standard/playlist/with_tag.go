package std_playlist

import "github.com/Mistsink/kuwo-api/internal/core/standard"

type WithTagResp struct {
	standard.CommonRes
	Result WithTagRes `json:"result" yzh:"result"`
}

type WithTagRes struct {
	TotalCnt int        `json:"total_cnt" yzh:"total_cnt"`
	PlList   []PlDetail `json:"list" yzh:"list"`
	N        int        `json:"n" yzh:"n"`
	P        int        `json:"p" yzh:"p"`
}
