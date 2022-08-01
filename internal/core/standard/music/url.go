package std_music

import (
	"github.com/Mistsink/kuwo-api/internal/core/standard"
)

type UrlResp struct {
	standard.CommonRes
	Result Result `json:"result" yzh:"result"`
}

type Result struct {
	URL string `json:"url" yzh:"url"`
}
