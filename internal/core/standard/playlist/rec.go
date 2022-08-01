package std_playlist

import "github.com/Mistsink/kuwo-api/internal/core/standard"

type RecResp struct {
	standard.CommonRes
	Result RecRes `json:"result" yzh:"result"`
}

type RecRes struct {
	List []RecItem `json:"list" yzh:"list"`
}

type RecItem struct {
	ImgUrl    string `json:"img_url" yzh:"rec_img_url"`
	UserName  string `json:"user_name" yzh:"rec_user_name"`
	Img700    string `json:"img_700" yzh:"rec_img_700"`
	Img300    string `json:"img_300" yzh:"rec_img_300"`
	Img500    string `json:"img_500" yzh:"rec_img_500"`
	TotalCnt  int    `json:"total_cnt" yzh:"rec_total_cnt"`
	Name      string `json:"name" yzh:"rec_name"`
	ListenCnt int    `json:"listen_cnt" yzh:"rec_lis_cnt"`
	ID        int    `json:"id" yzh:"rec_id"`
}
