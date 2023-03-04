package service

type MusicUrlReq struct {
	PlatformTagReq
	Id uint `form:"id" binding:"required,gte=1"`
}

type MusicCommentReq struct {
	PlatformTagReq
	Id     uint   `form:"id" binding:"required,gte=1"`
	Type   string `form:"type,default=get_rec_comment" binding:"omitempty,oneof=get_rec_comment get_comment"`
	P      uint   `form:"p,default=1" binding:"omitempty,gte=1"`
	N      uint   `form:"n,default=30" binding:"omitempty,gte=1"`
	Digest uint   `form:"digest,default=15" binding:"omitempty,oneof=2 7 8 15"`
}

// type RespAlbum struct {
// 	*RespComItem
// }
//
// type RespDuration struct {
// 	Duration    int    `json:"duration"`
// 	TimeMinutes string `json:"time_minutes"`
// }
//
// type RespMusicInfo struct {
// 	*RespComItem
// 	Score string `json:"score"`
// 	//	无损
// 	HasLossLess bool          `json:"has_loss_less"`
// 	ReleaseDate string        `json:"release_date"`
// 	Img120      string        `json:"img_120"`
// 	Album       *RespAlbum    `json:"album"`
// 	Artist      *RespArtist   `json:"artist"`
// 	Duration    *RespDuration `json:"duration"`
// }
//
// type RespMusicUrl struct {
// 	Url string `json:"url"`
// }
