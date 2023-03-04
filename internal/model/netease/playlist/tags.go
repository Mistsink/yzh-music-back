package playlist

import (
	"github.com/Mistsink/kuwo-api/internal/model/netease"
	"strconv"
)

var (
	Categorys = map[int]string{
		0: "语种",
		1: "风格",
		2: "场景",
		3: "情感",
		4: "主题",
	}
)

type TagResp struct {
	netease.CommonResp
	Tags []Tag `json:"tags"`
}

type Tag struct {
	ID       int    `json:"id" yzh:"tag_id"`
	Name     string `json:"name" yzh:"tag_name"`
	Type     int    `json:"type"`
	Category int    `json:"category"`
	Hot      bool   `json:"hot"`
}

func (r *TagResp) Format() (any, error) {
	ret := &TagsRespFormatted{
		CommonResp: r.CommonResp,
	}

	_catmap := make(map[int][]Tag, 4)
	for _, rawTag := range r.Tags {
		_catmap[rawTag.Category] = append(_catmap[rawTag.Category], rawTag)
	}
	for cat, tags := range _catmap {
		ret.Result = append(ret.Result, TagsResFormatted{
			Name: Categorys[cat],
			ID:   strconv.Itoa(cat),
			Data: tags,
		})
	}

	return ret, nil
}

type TagsRespFormatted struct {
	netease.CommonResp
	Result []TagsResFormatted `yzh:"result"`
}

type TagsResFormatted struct {
	Name string `yzh:"name"`
	ID   string `yzh:"fake_id"`
	Data []Tag  `yzh:"list"`
}
