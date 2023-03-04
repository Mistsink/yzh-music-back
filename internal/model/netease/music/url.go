package music

import (
	"github.com/Mistsink/kuwo-api/internal/model/netease"
)

type UrlResp struct {
	netease.CommonResp
	Data []Datum `json:"data"`
}

func (res *UrlResp) Format() (any, error) {
	return &UrlRespFormatted{
		res.CommonResp,
		res.Data[0],
	}, nil
}

type Datum struct {
	URL        string `json:"url" yzh:"url"`
	ID         int    `json:"id"`
	Br         int    `json:"br"`
	Size       int    `json:"size"`
	Md5        string `json:"md5"`
	Code       int    `json:"code"`
	Fee        int    `json:"fee"`
	Payed      int    `json:"payed"`
	Level      string `json:"level"`
	EncodeType string `json:"encodeType"`
	Time       int    `json:"time"`
}

type UrlRespFormatted struct {
	netease.CommonResp
	Data Datum `json:"data" yzh:"result"`
}
