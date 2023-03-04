package search

import (
	"github.com/Mistsink/kuwo-api/internal/model/kuwo"
)

type AlbumResp struct {
	kuwo.CommonRes
	Data AlbumRes `json:"data"`
}

type AlbumRes struct {
	Total     string      `json:"total" yzh:"total_cnt"`
	AlbumList []AlbumItem `json:"albumList" yzh:"list"`
}

type AlbumItem struct {
	ContentType string `json:"content_type"`
	Albuminfo   string `json:"albuminfo" yzh:"al_info"`
	Artist      string `json:"artist" yzh:"art_name"`
	ReleaseDate string `json:"releaseDate" yzh:"release_date"`
	Album       string `json:"album" yzh:"al_name"`
	Albumid     int    `json:"albumid" yzh:"al_id"`
	Artistid    int    `json:"artistid" yzh:"art_id"`
	Pic         string `json:"pic" yzh:"al_img_url"`
	Isstar      int    `json:"isstar"`
	Lang        string `json:"lang"`
}
