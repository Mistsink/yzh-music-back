package std_artist

import (
	"github.com/Mistsink/kuwo-api/internal/core/standard"
	std_music "github.com/Mistsink/kuwo-api/internal/core/standard/music"
)

type MusicListResp struct {
	standard.CommonRes
	Result MusicListRes `json:"result" yzh:"result"`
}

type MusicListRes struct {
	TotalCnt int                 `json:"total_cnt" yzh:"total_cnt"`
	List     []std_music.InfoRes `json:"list" yzh:"list"`
}
