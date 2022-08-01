package std_search

import "github.com/Mistsink/kuwo-api/internal/core/standard"

type MVResp struct {
	standard.CommonRes
	Result MVRes `json:"result" yzh:"result"`
}

type MVRes struct {
	TotalCnt int      `json:"total_cnt" yzh:"total_cnt"`
	Mvlist   []MvInfo `json:"list" yzh:"res_list"`
}

type MvInfo struct {
	MvPlayCnt int        `json:"play_cnt"`
	Name      string     `json:"name" yzh:"name"`
	ID        int        `json:"id" yzh:"id"`
	ImgURL    string     `json:"img_url" yzh:"img_url"`
	Duration  Duration   `json:"duration" yzh:"duration"`
	Artist    ArtistInfo `json:"artist" yzh:"artist"`
}

type Duration struct {
	Duration    int    `json:"duration" yzh:"du_duration"`
	TimeMinutes string `json:"time_minutes" yzh:"du_time"`
}
