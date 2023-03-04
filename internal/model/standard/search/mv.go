package std_search

import (
	"github.com/Mistsink/kuwo-api/internal/model/standard"
)

type MVResp struct {
	standard.CommonResp
	Result MVRes `json:"result" yzh:"result"`
}

type MVRes struct {
	TotalCnt int      `json:"total_cnt" yzh:"total_cnt"`
	Mvlist   []MvInfo `json:"list" yzh:"res_list"`
}

type MvInfo struct {
	MvPlayCnt int                 `json:"play_cnt"`
	Name      string              `json:"name" yzh:"name"`
	ID        int                 `json:"id" yzh:"id"`
	ImgURL    string              `json:"img_url" yzh:"img_url"`
	Duration  standard.Duration   `json:"duration" yzh:"duration"`
	Artist    standard.ArtistInfo `json:"artist" yzh:"artist"`
}
