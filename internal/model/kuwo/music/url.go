package music

import (
	"github.com/Mistsink/kuwo-api/internal/model/kuwo"
)

type UrlResp struct {
	kuwo.CommonRes
	Data UrlData `json:"data" yzh:"result"`
}

type UrlData struct {
	URL string `json:"url" yzh:"url"`
}
