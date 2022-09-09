package std_search

import (
	"github.com/Mistsink/kuwo-api/internal/core/standard"
)

type MusicResp struct {
	standard.CommonRes
	Result MusicRes `json:"result" yzh:"result"`
}

type MusicRes struct {
	TotalCnt int                `json:"total_cnt" yzh:"total_cnt"`
	List     []standard.InfoRes `json:"list" yzh:"res_list"`
}
