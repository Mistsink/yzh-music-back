package search

import (
	"github.com/Mistsink/kuwo-api/internal/core/kuwo"
	"github.com/Mistsink/kuwo-api/internal/core/kuwo/music"
)

type MusicResp struct {
	kuwo.CommonRes
	Data Data `json:"data" yzh:"result"`
}

type Data struct {
	Total string                `json:"total" yzh:"total_cnt"`
	List  []music.InfoData[int] `json:"list" yzh:"res_list"`
}
