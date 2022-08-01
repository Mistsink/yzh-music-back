package std_search

import "github.com/Mistsink/kuwo-api/internal/core/standard"

type ArtistResp struct {
	standard.CommonRes
	Result ArtistRes `json:"result" yzh:"result"`
}

type ArtistRes struct {
	TotalCnt   int                `json:"total_cnt" yzh:"total_cnt"`
	ArtistList []ArtistDetailInfo `json:"list" yzh:"res_list"`
}

type ArtistDetailInfo struct {
	Country  string `json:"country" yzh:"art_country_name"`
	ImgURL   string `json:"img_url" yzh:"art_img_url"`
	MusicNum int    `json:"music_num" yzh:"art_music_num"`
	Img120   string `json:"img_120" yzh:"art_img_120"`
	Name     string `json:"name" yzh:"art_name"`
	Img70    string `json:"img_70" yzh:"art_img_70"`
	ID       int    `json:"id" yzh:"art_id"`
	Img300   string `json:"img_300" yzh:"art_img_300"`
}
