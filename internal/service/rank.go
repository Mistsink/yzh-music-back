package service

type RankDetailReq struct {
	Id uint `form:"id" binding:"required,gte=1"`
	P     uint   `form:"p,default=1" binding:"omitempty,gte=1"`
	N     uint   `form:"n,default=30" binding:"omitempty,gte=1"`
}
