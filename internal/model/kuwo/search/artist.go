package search

import (
	"github.com/Mistsink/kuwo-api/internal/model/kuwo"
)

type ArtistResp struct {
	kuwo.CommonRes
	Data ArtistRes `json:"data" yzh:"result"`
}

type ArtistRes struct {
	Total string       `json:"total" yzh:"total_cnt"`
	List  []ArtistInfo `json:"list" yzh:"res_list"`
}

type ArtistInfo struct {
	Country  string `json:"country" yzh:"art_country_name"`
	Pic      string `json:"pic" yzh:"art_img_url"`
	MusicNum int    `json:"musicNum" yzh:"art_music_num"`
	Pic120   string `json:"pic120" yzh:"art_img_120"`
	IsStar   int    `json:"isStar"`
	Name     string `json:"name" yzh:"art_name"`
	Pic70    string `json:"pic70" yzh:"art_img_70"`
	ID       int    `json:"id" yzh:"art_id"`
	Pic300   string `json:"pic300" yzh:"art_img_300"`
}
