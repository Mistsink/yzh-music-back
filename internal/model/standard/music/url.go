package std_music

import (
	"github.com/Mistsink/kuwo-api/internal/model/standard"
)

type UrlResp struct {
	standard.CommonResp
	Result Result `json:"result" yzh:"result"`
}

type Result struct {
	standard.PlatformTag
	URL string `json:"url" yzh:"url"`
}
