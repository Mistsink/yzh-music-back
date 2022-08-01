package search

import "github.com/Mistsink/kuwo-api/internal/core/kuwo"

type HintResp struct {
	kuwo.CommonRes
	Data []string `json:"data" yzh:"result"`
}
