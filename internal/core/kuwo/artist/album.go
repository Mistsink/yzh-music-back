package artist

import (
	"github.com/Mistsink/kuwo-api/internal/core/kuwo"
	"github.com/Mistsink/kuwo-api/internal/core/kuwo/search"
)

type AlbumResp struct {
	kuwo.CommonRes
	Data AlbumRes `json:"data"`
}

type AlbumRes struct {
	Total     string             `json:"total" yzh:"total_cnt"`
	AlbumList []search.AlbumItem `json:"albumList" yzh:"list"`
}
