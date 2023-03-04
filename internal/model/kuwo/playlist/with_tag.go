package playlist

import (
	"github.com/Mistsink/kuwo-api/internal/model/kuwo"
)

type WithTagResp struct {
	kuwo.CommonRes
	Data WithTagRes `json:"data" yzh:"result"`
}

type WithTagRes struct {
	TotalCnt int        `json:"total" yzh:"total_cnt"`
	Data     []PlDetail `json:"data" yzh:"list"`
	Rn       int        `json:"rn" yzh:"n"`
	Pn       int        `json:"pn" yzh:"p"`
}
