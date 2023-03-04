package search

import (
	"github.com/Mistsink/kuwo-api/internal/model/kuwo"
)

type HintResp struct {
	kuwo.CommonRes
	Data []string `json:"data" yzh:"result"`
}
