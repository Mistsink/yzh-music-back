package netease

import (
	"github.com/Mistsink/kuwo-api/internal/proxy/copy"
)

type CommonResp struct {
	Code int    `json:"code" yzh:"code"`
	Msg  string `json:"msg" yzh:"msg"`
}

type SongFormatted struct {
	Name        string        `json:"name" yzh:"name"`
	ID          int           `json:"id" yzh:"id"`
	Ar          Artist        `json:"ar"`
	Alia        []interface{} `json:"alia"`
	Pop         int           `json:"pop" yzh:"score"` //	歌曲热度
	Fee         int           `json:"fee"`
	Al          Album         `json:"al" yzh:"album"`
	Dt          int           `json:"dt" yzh:"du_duration"`
	Mv          int           `json:"mv" yzh:"mv_id"`
	PublishTime int           `json:"publishTime" yzh:"release_date"`
}

func (s *Song) Format() *SongFormatted {
	ret := &SongFormatted{}

	_ = copy.CopyByTag(ret, s, "yzh")
	ret.Ar = s.Ar[0]
	return ret
}

type Song struct {
	Name        string        `json:"name" yzh:"name"`
	ID          int           `json:"id" yzh:"id"`
	Pop         int           `json:"pop" yzh:"score"` //	歌曲热度
	Al          Album         `json:"al" yzh:"album"`
	Dt          int           `json:"dt" yzh:"du_duration"`
	Mv          int           `json:"mv" yzh:"mv_id"`
	PublishTime int           `json:"publishTime" yzh:"release_date"`
	Ar          []Artist      `json:"ar"`
	Alia        []interface{} `json:"alia"`
	Fee         int           `json:"fee"`
	// H                    H             `json:"h"`
	// M                    H             `json:"m"`
	// L                    H             `json:"l"`
	// Sq                   H             `json:"sq"`
	Hr                   interface{}   `json:"hr"`
	A                    interface{}   `json:"a"`
	CD                   string        `json:"cd"`
	No                   int           `json:"no"`
	RtURL                interface{}   `json:"rtUrl"`
	Ftype                int           `json:"ftype"`
	RtUrls               []interface{} `json:"rtUrls"`
	DjID                 int           `json:"djId"`
	Copyright            int           `json:"copyright"`
	SID                  int           `json:"s_id"`
	Mark                 int           `json:"mark"`
	OriginCoverType      int           `json:"originCoverType"`
	OriginSongSimpleData interface{}   `json:"originSongSimpleData"`
	TagPicList           interface{}   `json:"tagPicList"`
	ResourceState        bool          `json:"resourceState"`
	Version              int           `json:"version"`
	SongJumpInfo         interface{}   `json:"songJumpInfo"`
	EntertainmentTags    interface{}   `json:"entertainmentTags"`
	AwardTags            interface{}   `json:"awardTags"`
	Single               int           `json:"single"`
	NoCopyrightRcmd      interface{}   `json:"noCopyrightRcmd"`
	Mst                  int           `json:"mst"`
	Cp                   int           `json:"cp"`
	Rtype                int           `json:"rtype"`
	Rurl                 interface{}   `json:"rurl"`
	Tns                  []string      `json:"tns"`
}

type Album struct {
	ID          int           `json:"id" yzh:"al_id"`
	Name        string        `json:"name" yzh:"al_name"`
	PicURL      string        `json:"picUrl" yzh:"al_img_url"`
	Tns         []interface{} `json:"tns"`
	PicStr      string        `json:"pic_str"`
	PublishTime int           `json:"publishTime"`
	Pic         float64       `json:"pic"`
}

type Artist struct {
	ID        int           `json:"id" yzh:"art_id"`
	Name      string        `json:"name" yzh:"art_name"`
	Tns       []interface{} `json:"tns"`
	AlbumSize int           `json:"albumSize"`
	Alias     []interface{} `json:"alias"`
}

type H struct {
	Br   int `json:"br"`
	Fid  int `json:"fid"`
	Size int `json:"size"`
	Vd   int `json:"vd"`
	Sr   int `json:"sr"`
}
