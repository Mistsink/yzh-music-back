package artist

import (
	"github.com/Mistsink/kuwo-api/internal/model/kuwo"
)

type MusicListResp struct {
	kuwo.CommonRes
	Result MusicListRes `json:"data" yzh:"result"`
}

type MusicListRes struct {
	Total int                     `json:"total" yzh:"total_cnt"`
	List  []kuwo.InfoData[string] `json:"list" yzh:"list"`
}
