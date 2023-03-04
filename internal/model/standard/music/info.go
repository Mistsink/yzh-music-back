package std_music

import (
	"github.com/Mistsink/kuwo-api/internal/model/standard"
)

type InfoResp struct {
	standard.CommonResp
	Result standard.InfoRes `json:"result" yzh:"result"`
}
