package artist

import (
	"fmt"
	"github.com/Mistsink/kuwo-api/internal/model/netease"
	"time"
)

type AlbumResp struct {
	netease.CommonResp
	Artist netease.Artist  `json:"artist"`
	Albums []netease.Album `json:"hotAlbums"`
	More   bool            `json:"more"`
}

func (r *AlbumResp) Format() (any, error) {
	ret := &AlbumRespFormatted{
		CommonResp: r.CommonResp,
	}

	ret.Result = AlbumResFormatted{
		TotalCnt: r.Artist.AlbumSize,
	}

	ret.Result.AlbumItems = make([]AlbumItemFormatted, len(r.Albums), len(r.Albums))
	for i, rawAlbum := range r.Albums {
		// album := &AlbumItemFormatted{}
		//	unix 时间戳 转成 格式化字符串
		year, mon, day := time.UnixMilli(int64(rawAlbum.PublishTime)).Date()

		album := AlbumItemFormatted{
			fmt.Sprintf("%d-%d-%d", year, mon, day),
			rawAlbum,
			r.Artist,
		}

		ret.Result.AlbumItems[i] = album
	}

	return ret, nil
}

type AlbumRespFormatted struct {
	netease.CommonResp
	Result AlbumResFormatted `yzh:"result"`
}

type AlbumResFormatted struct {
	TotalCnt   int                  `yzh:"total_cnt"`
	AlbumItems []AlbumItemFormatted `yzh:"list"`
}

type AlbumItemFormatted struct {
	ReleaseDate string         `yzh:"release_date"`
	Album       netease.Album  `yzh:"album"`
	Artist      netease.Artist `yzh:"artist"`
}

type Artist struct {
	Img1V1ID    float64  `json:"img1v1Id"`
	TopicPerson int      `json:"topicPerson"`
	Followed    bool     `json:"followed"`
	Alias       []string `json:"alias"`
	PicID       float64  `json:"picId"`
	BriefDesc   string   `json:"briefDesc"`
	MusicSize   int      `json:"musicSize"`
	AlbumSize   int      `json:"albumSize"`
	PicURL      string   `json:"picUrl"`
	Img1V1URL   string   `json:"img1v1Url"`
	Trans       string   `json:"trans"`
	Name        string   `json:"name"`
	ID          int      `json:"id"`
	PicIDStr    any      `json:"picId_str,omitempty"`
	Img1V1IDStr string   `json:"img1v1Id_str"`
}

type Album struct {
	Songs           []interface{} `json:"songs"`
	Paid            bool          `json:"paid"`
	OnSale          bool          `json:"onSale"`
	Mark            int           `json:"mark"`
	AwardTags       interface{}   `json:"awardTags"`
	BlurPicURL      string        `json:"blurPicUrl"`
	Alias           []string      `json:"alias"`
	Artists         []Artist      `json:"artists"`
	CopyrightID     int           `json:"copyrightId"`
	PicID           float64       `json:"picId"`
	Artist          Artist        `json:"artist"`
	PublishTime     int           `json:"publishTime"`
	Company         string        `json:"company"`
	BriefDesc       string        `json:"briefDesc"`
	PicURL          string        `json:"picUrl"`
	CommentThreadID string        `json:"commentThreadId"`
	Pic             float64       `json:"pic"`
	CompanyID       int           `json:"companyId"`
	Tags            string        `json:"tags"`
	Description     string        `json:"description"`
	Status          int           `json:"status"`
	SubType         string        `json:"subType"`
	Name            string        `json:"name"`
	ID              int           `json:"id"`
	Type            string        `json:"type"`
	Size            int           `json:"size"`
	PicIDStr        string        `json:"picId_str"`
	IsSub           bool          `json:"isSub"`
}
