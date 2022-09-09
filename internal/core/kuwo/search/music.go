package search

import (
	"github.com/Mistsink/kuwo-api/internal/core/kuwo"
)

type MusicResp struct {
	kuwo.CommonRes
	Data Data `json:"data" yzh:"result"`
}

type Data struct {
	Total string               `json:"total" yzh:"total_cnt"`
	List  []kuwo.InfoData[int] `json:"list" yzh:"res_list"`
}
