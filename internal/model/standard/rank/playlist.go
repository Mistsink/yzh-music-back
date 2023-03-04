package std_rank

import (
	"github.com/Mistsink/kuwo-api/internal/model/standard"
)

type PlaylistResp struct {
	standard.CommonResp
	Result []PlaylistRes `json:"result" yzh:"result"`
}

type PlaylistRes struct {
	Category string `json:"category" yzh:"category_name"`
	List     []Pl   `json:"list" yzh:"list"`
}

type Pl struct {
	// SourceId string `json:"source_id" yzh:"pl_source_id"`
	// Source   string `json:"source" yzh:"source"`
	Intro      string `json:"intro" yzh:"pl_intro"`
	Name       string `json:"name" yzh:"pl_name"`
	ID         int    `json:"id"  yzh:"pl_id"`
	Pic        string `json:"img_url" yzh:"pl_img_url"`
	UpdateTime string `json:"update_time" yzh:"pl_up_time"`
}
