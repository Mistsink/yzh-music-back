package search

import "github.com/Mistsink/kuwo-api/internal/core/kuwo"

type PlaylistResp struct {
	kuwo.CommonRes
	Data PlaylistRes `json:"data" yzh:"result"`
}

type PlaylistRes struct {
	Total string         `json:"total" yzh:"total_cnt"`
	List  []PlaylistInfo `json:"list" yzh:"res_list"`
}

type PlaylistInfo struct {
	Img       string `json:"img" yzh:"img_url"`
	Total     string `json:"total" yzh:"total_cnt"`
	Uname     string `json:"uname" yzh:"user_name"`
	Name      string `json:"name" yzh:"name"`
	Listencnt string `json:"listencnt" yzh:"listen_cnt"`
	ID        string `json:"id" yzh:"id"`
}
