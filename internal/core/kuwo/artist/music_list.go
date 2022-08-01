package artist

import (
	"github.com/Mistsink/kuwo-api/internal/core/kuwo"
	"github.com/Mistsink/kuwo-api/internal/core/kuwo/music"
)

type MusicListResp struct {
	kuwo.CommonRes
	Result MusicListRes `json:"data" yzh:"result"`
}

type MusicListRes struct {
	Total int                      `json:"total" yzh:"total_cnt"`
	List  []music.InfoData[string] `json:"list" yzh:"list"`
}
