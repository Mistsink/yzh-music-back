package std_search

import (
	"github.com/Mistsink/kuwo-api/internal/model/standard"
)

type AlbumResp struct {
	standard.CommonResp
	Result AlbumRes `json:"result" yzh:"result"`
}

type AlbumRes struct {
	TotalCnt  int                  `json:"total_cnt" yzh:"total_cnt"`
	AlbumList []standard.AlbumItem `json:"list" yzh:"list"`
}
