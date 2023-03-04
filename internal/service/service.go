package service

type PlatformTagReq struct {
	Tag string `form:"tag" binding:"required" oneof:"kuwo kugou qq netease"`
}
