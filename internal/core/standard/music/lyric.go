package std_music

import "github.com/Mistsink/kuwo-api/internal/core/standard"

type LyricResp struct {
	standard.CommonRes
	Result LyricRes `json:"result" yzh:"result"`
}

type LyricRes struct {
	Lrclist []Lrc `json:"lyric_list" yzh:"lyric_list"`
}

type Lrc struct {
	LineLyric string `json:"lrc_line" yzh:"lrc_line"`
	Time      string `json:"lrc_time" yzh:"lrc_time"`
}
