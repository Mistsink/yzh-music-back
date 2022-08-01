package service

type ArtWithTagReq struct {
	Category uint   `form:"category,default=0" binding:"omitempty,gte=0,lte=10"`
	P        uint   `form:"p,default=1" binding:"omitempty,gte=1"`
	N        uint   `form:"n,default=100" binding:"omitempty,gte=1"`
	Prefix   string `form:"prefix" binding:"omitempty,oneof=A-Z"`
}

type ArtMusicListReq struct {
	Id uint `form:"id" binding:"required,gte=1"`
	P  uint `form:"p,default=1" binding:"omitempty,gte=1"`
	N  uint `form:"n,default=100" binding:"omitempty,gte=1"`
}

type ArtRecReq struct {
	Category uint `form:"category,default=11" binding:"omitempty,oneof=11 12 13 16"`
	P        uint `form:"p,default=1" binding:"omitempty,gte=1"`
	N        uint `form:"n,default=100" binding:"omitempty,gte=1"`
}

type RespArtist struct {
	*RespComItem
}
