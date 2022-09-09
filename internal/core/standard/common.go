package standard

type CommonRes struct {
	Code  int    `json:"code" yzh:"code"`
	Msg   string `json:"msg" yzh:"msg"`
	ReqID string `json:"req_id" yzh:"req_id"`
}

// Album
type AlbumItem struct {
	ReleaseDate string     `json:"release_date" yzh:"release_date"`
	Album       AlbumInfo  `json:"album" yzh:"album"`
	Artist      ArtistInfo `json:"artist" yzh:"artist"`
}

type AlbumInfo struct {
	Albuminfo string `json:"info" yzh:"al_info"`
	ImgURL    string `json:"img_url" yzh:"al_img_url"`
	Name      string `json:"name" yzh:"al_name"`
	ID        int    `json:"id" yzh:"al_id"`
}

type ArtistInfo struct {
	Name string `json:"name" yzh:"art_name"`
	ID   int    `json:"id" yzh:"art_id"`
}

// Music Info
type InfoRes struct {
	ImgURL      string   `json:"img_url" yzh:"img_url"`
	Name        string   `json:"name" yzh:"name"`
	ID          int      `json:"id" yzh:"id"`
	Score       int      `json:"score" yzh:"score"`
	HasLossLess bool     `json:"has_loss_less" yzh:"has_loss_less"`
	ReleaseDate string   `json:"release_date" yzh:"release_date"`
	Img120      string   `json:"img_120" yzh:"img_120"`
	Album       Album    `json:"album" yzh:"album"`
	Artist      Artist   `json:"artist" yzh:"artist"`
	Duration    Duration `json:"duration" yzh:"duration"`
	HasMV       int      `json:"has_mv" yzh:"has_mv"`
	MVId        int      `json:"mv_id" yzh:"mv_id"`
}

type Album struct {
	ImgURL string `json:"img_url" yzh:"al_img_url"`
	Name   string `json:"name" yzh:"al_name"`
	ID     int    `json:"id" yzh:"al_id"`
}

type Artist struct {
	Name string `json:"name" yzh:"art_name"`
	ID   int    `json:"id" yzh:"art_id"`
}

type Duration struct {
	Duration    int    `json:"duration" yzh:"du_duration"`
	TimeMinutes string `json:"time_minutes" yzh:"du_time"`
}
