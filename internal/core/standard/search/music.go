package std_search

import (
	"github.com/Mistsink/kuwo-api/internal/core/standard"
	std_music "github.com/Mistsink/kuwo-api/internal/core/standard/music"
)

type MusicResp struct {
	standard.CommonRes
	Result MusicRes `json:"result" yzh:"result"`
}

type MusicRes struct {
	TotalCnt int                 `json:"total_cnt" yzh:"total_cnt"`
	List     []std_music.InfoRes `json:"list" yzh:"res_list"`
}
