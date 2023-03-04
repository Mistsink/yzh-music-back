package playlist

import (
	"github.com/Mistsink/kuwo-api/internal/model/kuwo"
)

type DetailResp struct {
	kuwo.CommonRes
	Data DetailRes `json:"data" yzh:"result"`
}

type DetailRes struct {
	Img        string               `json:"img" yzh:"img_url"`
	UPic       string               `json:"uPic" yzh:"user_img_url"`
	Uname      string               `json:"uname" yzh:"user_name"`
	Img700     string               `json:"img700" yzh:"img_700"`
	Img300     string               `json:"img300" yzh:"img_300"`
	UserName   string               `json:"userName"`
	Img500     string               `json:"img500" yzh:"img_500"`
	IsOfficial int                  `json:"isOfficial"`
	Total      int                  `json:"total" yzh:"total_cnt"`
	Name       string               `json:"name" yzh:"name"`
	Listencnt  int                  `json:"listencnt" yzh:"lis_cnt"`
	ID         int                  `json:"id" yzh:"id"`
	Tag        string               `json:"tag" yzh:"tag"`
	MusicList  []kuwo.InfoData[int] `json:"musicList" yzh:"list"`
	Desc       string               `json:"desc"`
	Info       string               `json:"info"`
}
