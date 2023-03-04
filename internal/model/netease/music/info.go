package music

import (
	"github.com/Mistsink/kuwo-api/internal/model/netease"
)

type InfoResp struct {
	netease.CommonResp
	Songs []netease.Song `json:"songs"`
}
type InfoRespFormatted struct {
	netease.CommonResp
	Song *netease.SongFormatted `json:"songs" yzh:"result"`
}

func (r *InfoResp) Format() (any, error) {
	songsFormatted, err := r.Songs[0].Format()
	return &InfoRespFormatted{
		r.CommonResp,
		songsFormatted.(*netease.SongFormatted),
	}, err
}
