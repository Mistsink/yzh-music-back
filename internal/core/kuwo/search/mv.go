package search

import "github.com/Mistsink/kuwo-api/internal/core/kuwo"

type MVResp struct {
	kuwo.CommonRes
	Data MVRes `json:"data"`
}

type MVRes struct {
	Total  string   `json:"total" yzh:"total_cnt"`
	Mvlist []Mvlist `json:"mvlist" yzh:"res_list"`
}

type Mvlist struct {
	Duration        int    `json:"duration" yzh:"du_duration"`
	Artist          string `json:"artist" yzh:"art_name"`
	MvPlayCnt       int    `json:"mvPlayCnt" yzh:"play_cnt"`
	Name            string `json:"name" yzh:"name"`
	Online          int    `json:"online"`
	Artistid        int    `json:"artistid" yzh:"art_id"`
	SongTimeMinutes string `json:"songTimeMinutes" yzh:"du_time"`
	ID              int    `json:"id" yzh:"id"`
	Pic             string `json:"pic" yzh:"img_url"`
}
