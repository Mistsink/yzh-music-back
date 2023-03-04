package playlist

import (
	"github.com/Mistsink/kuwo-api/internal/model/netease"
)

type DetailResp struct {
	netease.CommonResp
	Songs []netease.Song `json:"songs"`
}

func (r *DetailResp) Format() (any, error) {
	ret := &DetailRespFormatted{
		CommonResp: r.CommonResp,
	}

	ret.Result = DetailResFormatted{
		TotalCnt:  len(r.Songs),
		MusicList: make([]netease.SongFormatted, len(r.Songs), len(r.Songs)),
	}

	for i, rawSong := range r.Songs {
		ret.Result.MusicList[i] = *rawSong.Format()
	}
	return ret, nil
}

type DetailRespFormatted struct {
	netease.CommonResp
	Result DetailResFormatted `yzh:"result"`
}

type DetailResFormatted struct {
	TotalCnt  int                     `yzh:"total_cnt"`
	MusicList []netease.SongFormatted `yzh:"list"`
}
