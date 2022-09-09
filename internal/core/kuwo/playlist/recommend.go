package playlist

import "github.com/Mistsink/kuwo-api/internal/core/kuwo"

type RecResp struct {
	kuwo.CommonRes
	Data RecRes `json:"data" yzh:"result"`
}

type RecRes struct {
	List []RecItem `json:"list" yzh:"list"`
}

type RecItem struct {
	Img       string `json:"img" yzh:"rec_img_url"`
	Uname     string `json:"uname" yzh:"rec_user_name"`
	Img700    string `json:"img700" yzh:"rec_img_700"`
	Img300    string `json:"img300" yzh:"rec_img_300"`
	UserName  string `json:"userName"`
	Img500    string `json:"img500" yzh:"rec_img_500"`
	Total     int    `json:"total" yzh:"rec_total_cnt"`
	Name      string `json:"name" yzh:"rec_name"`
	Listencnt int    `json:"listencnt" yzh:"rec_lis_cnt"`
	ID        int    `json:"id" yzh:"rec_id"`
	Desc      string `json:"desc"`
	Info      string `json:"info"`
	Tag       string `json:"tag,omitempty"`
}
