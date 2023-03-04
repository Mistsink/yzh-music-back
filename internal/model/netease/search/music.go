package search

import (
	"github.com/Mistsink/kuwo-api/internal/model/netease"
)

type MusicResp struct {
	netease.CommonResp
	Result MusicResult `json:"result"`
}

func (r *MusicResp) Format() (any, error) {
	ret := &MusicRespFormatted{
		CommonResp: r.CommonResp,
	}

	ret.Result = MusicResultFormatted{
		TotalCnt: r.Result.SongCount,
		List:     make([]netease.SongFormatted, len(r.Result.Songs), len(r.Result.Songs)),
	}

	for i, rawSong := range r.Result.Songs {
		if songs, e := rawSong.Format(); e != nil {
			return nil, e
		} else {
			ret.Result.List[i] = *(songs.(*netease.SongFormatted))
		}
	}
	return ret, nil
}

type MusicRespFormatted struct {
	netease.CommonResp
	Result MusicResultFormatted `yzh:"result"`
}

type MusicResultFormatted struct {
	TotalCnt int                     `yzh:"total_cnt"`
	List     []netease.SongFormatted `yzh:"res_list"`
}

type MusicResult struct {
	SearchQcReminder interface{}    `json:"searchQcReminder"`
	Songs            []netease.Song `json:"songs"`
	SongCount        int            `json:"songCount"`
}
