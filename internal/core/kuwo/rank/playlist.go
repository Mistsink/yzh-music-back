package rank

import "github.com/Mistsink/kuwo-api/internal/core/kuwo"

type PlaylistResp struct {
	kuwo.CommonRes
	Data []PlaylistRes `json:"data" yzh:"result"`
}

type PlaylistRes struct {
	Name string `json:"name" yzh:"category_name"`
	List []Pl   `json:"list" yzh:"list"`
}

type Pl struct {
	Sourceid string `json:"sourceid" yzh:"pl_id"`
	Intro    string `json:"intro" yzh:"pl_intro"`
	Name     string `json:"name" yzh:"pl_name"`
	// ID       string `json:"id" yzh:"pl_id"`
	Source string `json:"source" yzh:"pl_source"`
	Pic    string `json:"pic" yzh:"pl_img_url"`
	Pub    string `json:"pub" yzh:"pl_up_time"`
}
