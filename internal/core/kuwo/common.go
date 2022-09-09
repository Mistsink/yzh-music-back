package kuwo

type CommonRes struct {
	Code      int    `json:"code" yzh:"code"`
	CurTime   int    `json:"curTime"`
	Msg       string `json:"msg" yzh:"msg"`
	ProfileID string `json:"profileId"`
	ReqID     string `json:"reqId" yzh:"req_id"`
	TID       string `json:"tId"`
}

// Music Info
type InfoData[T comparable] struct {
	Musicrid          string        `json:"musicrid"`
	Barrage           string        `json:"barrage"`
	Artist            string        `json:"artist" yzh:"art_name"`
	Mvpayinfo         *Mvpayinfo[T] `json:"mvpayinfo" yzh:"mv_info"`
	Pic               string        `json:"pic" yzh:"img_url"`
	Isstar            int           `json:"isstar"`
	Rid               int           `json:"rid" yzh:"id"`
	UpPCStr           string        `json:"upPcStr"`
	Duration          int           `json:"duration" yzh:"du_duration"`
	Score100          string        `json:"score100" yzh:"score"`
	AdSubtype         string        `json:"ad_subtype"`
	ContentType       string        `json:"content_type"`
	MvPlayCnt         int           `json:"mvPlayCnt"`
	Track             int           `json:"track"`
	HasLossless       bool          `json:"hasLossless" yzh:"has_loss_less"`
	Hasmv             int           `json:"hasmv" yzh:"has_mv"`
	ReleaseDate       string        `json:"releaseDate" yzh:"release_date"`
	Album             string        `json:"album" yzh:"al_name"`
	Albumid           any           `json:"albumid" yzh:"al_id"`
	Pay               string        `json:"pay"`
	Artistid          int           `json:"artistid" yzh:"art_id"`
	Albumpic          string        `json:"albumpic" yzh:"al_img_url"`
	Originalsongtype  int           `json:"originalsongtype"`
	SongTimeMinutes   string        `json:"songTimeMinutes" yzh:"du_time"`
	IsListenFee       bool          `json:"isListenFee"`
	MvUpPCStr         string        `json:"mvUpPcStr"`
	Pic120            string        `json:"pic120" yzh:"img_120"`
	Albuminfo         string        `json:"albuminfo"`
	Name              string        `json:"name" yzh:"name"`
	Online            int           `json:"online"`
	PayInfo           *PayInfo      `json:"payInfo"`
	TmeMusicianAdtype string        `json:"tme_musician_adtype"`
}

type Mvpayinfo[T comparable] struct {
	Play any `json:"play"`
	Vid  T   `json:"vid" yzh:"mv_id"`
	Down int `json:"down"`
}

type PayInfo struct {
	Play          string   `json:"play"`
	Nplay         string   `json:"nplay"`
	OverseasNplay string   `json:"overseas_nplay"`
	LocalEncrypt  string   `json:"local_encrypt"`
	Limitfree     any      `json:"limitfree"`
	RefrainStart  any      `json:"refrain_start"`
	FeeType       *FeeType `json:"feeType"`
	Down          string   `json:"down"`
	Ndown         string   `json:"ndown"`
	Download      string   `json:"download"`
	OverseasNdown string   `json:"overseas_ndown"`
}

type FeeType struct {
	Song string `json:"song"`
	Vip  string `json:"vip"`
}
