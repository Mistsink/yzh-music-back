package service

type PLDefaultReq struct {
	P     uint   `form:"p,default=1" binding:"omitempty,gte=1"`
	N     uint   `form:"n,default=30" binding:"omitempty,gte=1"`
	Order string `form:"order,default=new" binding:"omitempty,oneof= new hot"`
}

type PLRecommendReq struct {
	P uint `form:"p,default=1" binding:"omitempty,gte=1"`
	N uint `form:"n,default=5" binding:"omitempty,gte=1"`
}

type PLWithTagReq struct {
	Id uint `form:"id" binding:"required,gte=0"`
	P  uint `form:"p,default=1" binding:"omitempty,gte=1"`
	N  uint `form:"n,default=30" binding:"omitempty,gte=1"`
}

type PLDetailReq struct {
	Id uint `form:"id" binding:"required,gte=0"`
	P  uint `form:"p,default=1" binding:"omitempty,gte=1"`
	N  uint `form:"n,default=30" binding:"omitempty,gte=1"`
}

type RespComItem struct {
	ImgUrl    string `json:"img_url,omitempty"`
	UserName  string `json:"user_name,omitempty"`
	TotalCnt  any    `json:"total_cnt,int,omitempty"`
	Name      string `json:"name,omitempty"`
	ListenCnt any    `json:"listened_cnt,omitempty"`
	Id        any    `json:"id,omitempty"`
	Desc      string `json:"desc,omitempty"`
	Info      string `json:"info,omitempty"`
}

//	Resp for playlist default
type RespPLDefaultItem struct {
	*RespComItem
	FavorCnt string `json:"favorite_cnt"`
	UserID   string `json:"user_id"`
}

type RespPLDefault struct {
	SongTotal int                  `json:"song_total"`
	Data      []*RespPLDefaultItem `json:"data"`
	PageSize  int                  `json:"page_size"`
	PageIndex int                  `json:"page_index"`
}

//	Resp for playlist recommend
type RespPLRecItem struct {
	*RespComItem
	Img300 string `json:"img_300"`
	Img500 string `json:"img_500"`
	Img700 string `json:"img_700"`
}

type RespPLRec struct {
	Data []*RespPLRecItem `json:"data"`
}

//	Resp for playlist tags
type RespPLTagsItemDetail struct {
	*RespComItem
}
type RespPLTagsItem struct {
	*RespComItem
	ItemDetails []*RespPLTagsItemDetail `json:"data"`
	Type        string                  `json:"type"`
	TitleImgUrl string                  `json:"title_img_url"`
}

type RespPLTags struct {
	Data []*RespPLTagsItem `json:"data"`
}

//	Resp for playlist with tag
type RespPLWithTagItem struct {
	*RespComItem
	FavorCnt string `json:"favorite_cnt"`
	UserID   string `json:"user_id"`
}

type RespPLWithTag struct {
	SongTotal int                  `json:"song_total"`
	Data      []*RespPLWithTagItem `json:"data"`
	PageSize  int                  `json:"page_size"`
	PageIndex int                  `json:"page_index"`
}

//	Resp for playlist detail
type RespPLDetail struct {
	*RespComItem
	Img300     string           `json:"img_300"`
	Img500     string           `json:"img_500"`
	Img700     string           `json:"img_700"`
	AvatarUrl  string           `json:"avatar_url"`
	IsOfficial bool             `json:"is_official"`
	Tag        string           `json:"tag"`
	MusicList  []*RespMusicInfo `json:"music_list"`
}
