package artist

import (
	"github.com/Mistsink/kuwo-api/internal/model/netease"
)

type MusicListResp struct {
	netease.CommonResp
	Songs []*netease.Song `json:"songs" yzh:"list"`
	More  bool            `json:"more"`
	Total int             `json:"total" yzh:"total_cnt"`
}

func (r *MusicListResp) Format() (any, error) {
	ret := &MusicListRespFormatted{
		CommonResp: r.CommonResp,
	}

	ret.Result = MusicListItemFormatted{
		TotalCnt: r.Total,
	}

	ret.Result.List = make([]netease.SongFormatted, len(r.Songs), len(r.Songs))
	for i, rawSong := range r.Songs {
		if songFormatted, err := rawSong.Format(); err != nil {
			return nil, err
		} else {
			ret.Result.List[i] = *(songFormatted.(*netease.SongFormatted))
		}
	}

	return ret, nil
}

type MusicListRespFormatted struct {
	netease.CommonResp
	Result MusicListItemFormatted `yzh:"result"`
}

type MusicListItemFormatted struct {
	TotalCnt int                     `yzh:"total_cnt"`
	List     []netease.SongFormatted `yzh:"list"`
}
