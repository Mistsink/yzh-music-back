package std_search

import "github.com/Mistsink/kuwo-api/internal/core/standard"

type AlbumResp struct {
	standard.CommonRes
	Result AlbumRes `json:"result" yzh:"result"`
}

type AlbumRes struct {
	TotalCnt  int         `json:"total_cnt" yzh:"total_cnt"`
	AlbumList []AlbumItem `json:"list" yzh:"res_list"`
}

type AlbumItem struct {
	ReleaseDate string     `json:"release_date" yzh:"release_date"`
	Album       AlbumInfo  `json:"album" yzh:"album"`
	Artist      ArtistInfo `json:"artist" yzh:"artist"`
}

type AlbumInfo struct {
	Albuminfo string `json:"info" yzh:"al_info"`
	ImgURL    string `json:"img_url" yzh:"al_img_url"`
	Name      string `json:"name" yzh:"al_name"`
	ID        int    `json:"id" yzh:"al_id"`
}

type ArtistInfo struct {
	Name string `json:"name" yzh:"art_name"`
	ID   int    `json:"id" yzh:"art_id"`
}
