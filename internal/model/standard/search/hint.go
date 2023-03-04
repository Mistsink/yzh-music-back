package std_search

import (
	"github.com/Mistsink/kuwo-api/internal/model/standard"
)

type HintResp struct {
	standard.CommonResp
	Result []string `json:"result" yzh:"result"`
}
