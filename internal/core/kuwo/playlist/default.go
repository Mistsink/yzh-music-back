package playlist

import "github.com/Mistsink/kuwo-api/internal/core/kuwo"

type DefaultResp struct {
	kuwo.CommonRes
	Data DefaultRes `json:"data"`
}

type DefaultRes struct {
	Total int        `json:"total" yzh:"total_cnt"`
	Data  []PlDetail `json:"data" yzh:"list"`
	Rn    int        `json:"rn" yzh:"n"`
	Pn    int        `json:"pn" yzh:"p"`
}

type PlDetail struct {
	Img          string `json:"img" yzh:"pld_img_url"`
	Uname        string `json:"uname" yzh:"pld_user_name"`
	LosslessMark string `json:"lossless_mark"`
	Favorcnt     string `json:"favorcnt"`
	Isnew        string `json:"isnew"`
	Extend       string `json:"extend"`
	Uid          string `json:"uid"`
	Total        string `json:"total" yzh:"pld_total_cnt"`
	Commentcnt   string `json:"commentcnt"`
	Imgscript    string `json:"imgscript"`
	Digest       string `json:"digest"`
	Name         string `json:"name" yzh:"pld_name"`
	Listencnt    string `json:"listencnt" yzh:"pld_lis_cnt"`
	ID           string `json:"id" yzh:"pld_id"`
	Attribute    string `json:"attribute"`
	RadioID      string `json:"radio_id"`
	Desc         string `json:"desc"`
	Info         string `json:"info"`
}
