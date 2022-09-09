package std_search

import "github.com/Mistsink/kuwo-api/internal/core/standard"

type AlbumResp struct {
	standard.CommonRes
	Result AlbumRes `json:"result" yzh:"result"`
}

type AlbumRes struct {
	TotalCnt  int                  `json:"total_cnt" yzh:"total_cnt"`
	AlbumList []standard.AlbumItem `json:"list" yzh:"list"`
}
