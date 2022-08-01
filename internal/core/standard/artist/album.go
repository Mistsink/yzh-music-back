package std_artist

import (
	"github.com/Mistsink/kuwo-api/internal/core/standard"
	std_search "github.com/Mistsink/kuwo-api/internal/core/standard/search"
)

type AlbumResp struct {
	standard.CommonRes
	Result AlbumRes `json:"result" yzh:"result"`
}

type AlbumRes struct {
	TotalCnt  int                    `json:"total_cnt" yzh:"total_cnt"`
	AlbumList []std_search.AlbumItem `json:"list" yzh:"list"`
}
