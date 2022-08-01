package service

type SearComReq struct {
	Key string `form:"key" binding:"required"`
	P   uint   `form:"p,default=1" binding:"omitempty,gte=1"`
	N   uint   `form:"n,default=30" binding:"omitempty,gte=1"`
}

type SearHintReq struct {
	Key string `form:"key" binding:"required"`
}

type RespSearMusic struct {
	TotalCnt  int              `json:"total_cnt,int,omitempty"`
	MusicList []*RespMusicInfo `json:"music_list"`
}
