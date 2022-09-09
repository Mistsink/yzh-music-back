package music

import (
	"github.com/Mistsink/kuwo-api/internal/core/kuwo"
)

type InfoResp struct {
	kuwo.CommonRes
	Data *kuwo.InfoData[int] `json:"data" yzh:"result"`
}
