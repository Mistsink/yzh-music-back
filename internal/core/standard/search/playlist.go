package std_search

import "github.com/Mistsink/kuwo-api/internal/core/standard"

type PlaylistResp struct {
	standard.CommonRes
	Result PlaylistRes `json:"result" yzh:"result"`
}

type PlaylistRes struct {
	TotalCnt int            `json:"total_cnt" yzh:"total_cnt"`
	List     []PlaylistInfo `json:"list" yzh:"res_list"`
}

type PlaylistInfo struct {
	Img       string `json:"img_url" yzh:"img_url"`
	Total     string `json:"total_cnt" yzh:"total_cnt"`
	Uname     string `json:"user_name" yzh:"user_name"`
	Name      string `json:"name" yzh:"name"`
	Listencnt string `json:"listen_cnt" yzh:"listen_cnt"`
	ID        string `json:"id" yzh:"id"`
}
