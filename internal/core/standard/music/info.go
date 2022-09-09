package std_music

import "github.com/Mistsink/kuwo-api/internal/core/standard"

type InfoResp struct {
	standard.CommonRes
	Result standard.InfoRes `json:"result" yzh:"result"`
}
