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
	return &InfoRespFormatted{
		r.CommonResp,
		r.Songs[0].Format(),
	}, nil
}
