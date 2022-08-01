package std_search

import "github.com/Mistsink/kuwo-api/internal/core/standard"

type HintResp struct {
	standard.CommonRes
	Result []string `json:"result" yzh:"result"`
}
