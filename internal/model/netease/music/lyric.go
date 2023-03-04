package music

import (
	"fmt"
	"github.com/Mistsink/kuwo-api/internal/model/netease"
	"github.com/Mistsink/kuwo-api/internal/proxy/copy"
	"strings"
	"time"
)

type LyricResp struct {
	netease.CommonResp
	Sgc     bool   `json:"sgc"`
	Sfy     bool   `json:"sfy"`
	Qfy     bool   `json:"qfy"`
	Lrc     Klyric `json:"lrc"`
	Klyric  Klyric `json:"klyric"`
	Tlyric  Klyric `json:"tlyric"`
	Romalrc Klyric `json:"romalrc"`
}

type LyricRespFormatted struct {
	netease.CommonResp
	Result LyricRes `yzh:"result"`
}
type LyricRes struct {
	Lrclist []Lrc `yzh:"lyric_list"`
}

type Lrc struct {
	LineLyric string `yzh:"lrc_line"`
	Time      string `yzh:"lrc_time"` //	转成 秒 做单位的数字
}

func (l *LyricResp) Format() (any, error) {
	ret := &LyricRespFormatted{}
	_ = copy.CopyByTag(ret, l, "yzh")

	// tidy lyric

	// [00:04.050]\n[00:12.570]难以忘记初次见你\n
	rawLyrics := strings.Split(l.Lrc.Lyric, "\n")
	rawLyrics = rawLyrics[:len(rawLyrics)-1]
	lyrics := make([]Lrc, len(rawLyrics), len(rawLyrics))

	for i, rawLyric := range rawLyrics {
		raws := strings.SplitN(rawLyric[1:], "]", 2)

		rawTime := []byte(raws[0])
		rawTime = append(rawTime, 's')

		//	将 时间中的 ： 替换成具有时间含义的字母
		_replacedStr, _cnt := []uint8{'m', 'h'}, 0
		for i := len(rawTime) - 1; i >= 2; i-- {
			if rawTime[i] == ':' {
				rawTime[i] = _replacedStr[_cnt]
				_cnt++
				if _cnt >= 2 {
					break
				}
			}
		}

		du, err := time.ParseDuration(string(rawTime))
		if err != nil {
			return nil, err
		}

		lyrics[i] = Lrc{
			LineLyric: raws[1],
			Time:      fmt.Sprintf("%f", du.Seconds()),
		}
	}
	ret.Result.Lrclist = lyrics

	return ret, nil
}

type Klyric struct {
	Version int    `json:"version"`
	Lyric   string `json:"lyric"`
}
