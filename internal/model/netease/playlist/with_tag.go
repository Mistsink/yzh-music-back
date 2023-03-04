package playlist

import (
	"github.com/Mistsink/kuwo-api/internal/model/netease"
	"github.com/Mistsink/kuwo-api/internal/proxy/copy"
)

type WithTagResp struct {
	netease.CommonResp
	Playlists []Playlist `json:"playlists" yzh:"list"`
	More      bool       `json:"more"`
	Lasttime  int        `json:"lasttime"`
	Total     int        `json:"total" yzh:"total_cnt"`
}

func (r *WithTagResp) Format() (any, error) {
	ret := &WithTagRespFormatted{}

	if err := copy.CopyByTag(ret, r, "yzh"); err != nil {
		return nil, err
	}

	ret.Result.N = len(r.Playlists)

	return ret, nil
}

type WithTagRespFormatted struct {
	netease.CommonResp
	Result WithTagResFormatted `yzh:"result"`
}

type WithTagResFormatted struct {
	TotalCnt int        `yzh:"total_cnt"`
	PlList   []Playlist `yzh:"list"`
	N        int        `yzh:"n"`
}
type Playlist struct {
	Name                  string      `json:"name" yzh:"pld_name"`
	ID                    int         `json:"id" yzh:"pld_id"`
	PlayCount             int         `json:"playCount" yzh:"pld_lis_cnt"`
	CoverImgURL           string      `json:"coverImgUrl" yzh:"pld_img_url"`
	UpdateTime            int         `json:"updateTime" yzh:"pld_netease_update_time"`
	TrackNumberUpdateTime int         `json:"trackNumberUpdateTime"`
	Status                int         `json:"status"`
	UserID                int         `json:"userId"`
	CreateTime            int         `json:"createTime"`
	SubscribedCount       int         `json:"subscribedCount"`
	TrackCount            int         `json:"trackCount"`
	CloudTrackCount       int         `json:"cloudTrackCount"`
	CoverImgID            float64     `json:"coverImgId"`
	Description           string      `json:"description"`
	Tags                  []string    `json:"tags"`
	TrackUpdateTime       int         `json:"trackUpdateTime"`
	SpecialType           int         `json:"specialType"`
	TotalDuration         int         `json:"totalDuration"`
	Creator               Creator     `json:"creator"`
	Tracks                interface{} `json:"tracks"`
	Subscribers           []Creator   `json:"subscribers"`
	Subscribed            bool        `json:"subscribed"`
	CommentThreadID       string      `json:"commentThreadId"`
	NewImported           bool        `json:"newImported"`
	AdType                int         `json:"adType"`
	HighQuality           bool        `json:"highQuality"`
	Privacy               int         `json:"privacy"`
	Ordered               bool        `json:"ordered"`
	Anonimous             bool        `json:"anonimous"`
	CoverStatus           int         `json:"coverStatus"`
	RecommendInfo         interface{} `json:"recommendInfo"`
	ShareCount            int         `json:"shareCount"`
	CoverImgIDStr         string      `json:"coverImgId_str"`
	CommentCount          int         `json:"commentCount"`
	Copywriter            string      `json:"copywriter"`
	Tag                   string      `json:"tag"`
}

type Creator struct {
	Nickname              string        `json:"nickname" yzh:"pld_user_name"`
	DefaultAvatar         bool          `json:"defaultAvatar"`
	Province              int           `json:"province"`
	AuthStatus            int           `json:"authStatus"`
	Followed              bool          `json:"followed"`
	AvatarURL             string        `json:"avatarUrl"`
	AccountStatus         int           `json:"accountStatus"`
	Gender                int           `json:"gender"`
	City                  int           `json:"city"`
	Birthday              int           `json:"birthday"`
	UserID                int           `json:"userId"`
	UserType              int           `json:"userType"`
	Signature             string        `json:"signature"`
	Description           string        `json:"description"`
	DetailDescription     string        `json:"detailDescription"`
	AvatarImgID           float64       `json:"avatarImgId"`
	BackgroundImgID       float64       `json:"backgroundImgId"`
	BackgroundURL         string        `json:"backgroundUrl"`
	Authority             int           `json:"authority"`
	Mutual                bool          `json:"mutual"`
	ExpertTags            []string      `json:"expertTags"`
	Experts               interface{}   `json:"experts"`
	DjStatus              int           `json:"djStatus"`
	VipType               int           `json:"vipType"`
	RemarkName            interface{}   `json:"remarkName"`
	AuthenticationTypes   int           `json:"authenticationTypes"`
	AvatarDetail          *AvatarDetail `json:"avatarDetail"`
	BackgroundImgIDStr    string        `json:"backgroundImgIdStr"`
	AvatarImgIDStr        string        `json:"avatarImgIdStr"`
	Anchor                bool          `json:"anchor"`
	CreatorAvatarImgIDStr string        `json:"avatarImgId_str"`
}

type AvatarDetail struct {
	UserType        int    `json:"userType"`
	IdentityLevel   int    `json:"identityLevel"`
	IdentityIconURL string `json:"identityIconUrl"`
}
