package proxy

import (
	"fmt"
	"strconv"

	"github.com/Mistsink/kuwo-api/global"
	kuwo "github.com/Mistsink/kuwo-api/internal/model/kuwo/playlist"
	netease "github.com/Mistsink/kuwo-api/internal/model/netease/playlist"
	std "github.com/Mistsink/kuwo-api/internal/model/standard/playlist"
	"github.com/Mistsink/kuwo-api/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func PLDefault(c *gin.Context, param *service.PLDefaultReq) (*std.DefaultResp, error) {
	newUrl := "http://www.kuwo.cn/api/www/classify/playlist/getRcmPlayList"
	newUrl = fmt.Sprintf(
		"%s?pn=%d&rn=%d&order=%s",
		newUrl, param.P, param.N, param.Order)

	rawRet := &kuwo.DefaultResp{}
	ret := &std.DefaultResp{}

	err := getRespByNewUrl(c, newUrl, rawRet, ret)
	return ret, err
}

func PLRecommend(c *gin.Context, param *service.PLRecommendReq) (*std.RecResp, error) {
	newUrl := "http://www.kuwo.cn/api/www/rcm/index/playlist"
	newUrl = fmt.Sprintf(
		"%s?pn=%d&rn=%d&httpsStatus=1",
		newUrl, param.P, param.N)
	rawRet := &kuwo.RecResp{}
	ret := &std.RecResp{}

	err := getRespByNewUrl(c, newUrl, rawRet, ret)
	return ret, err
}

func PLTags(c *gin.Context, _ any) (*std.TagResp, error) {
	var (
		err error
		url string
		raw any
		ret = &std.TagResp{}
	)

	for _, _tag := range global.ProxySetting.AvailableTags {
		_ret := &std.TagResp{}

		tag := PlatformTag(_tag)
		switch tag {
		case Tkuwo:
			url = "http://www.kuwo.cn/api/www/playlist/getTagList"
			raw = &kuwo.TagResp{}
		case Tnetease:
			url = "/playlist/highquality/tags"
			raw = &netease.TagResp{}
		}

		e := sendAndFormat(&ProxyReqOpt{
			tag:    tag,
			rawUrl: url,
		}, raw, _ret)
		if e != nil {
			err = errors.Wrap(err, e.Error())
		}

		for i := range _ret.Result {
			_ret.Result[i].Tag = _tag
		}

		ret.Result = append(ret.Result, _ret.Result...)
	}

	return ret, err
}

func PLWithTag(c *gin.Context, param *service.PLWithTagReq) (*std.WithTagResp, error) {
	var (
		err error
		url string
		raw any
		ret = &std.WithTagResp{}
		tag = PlatformTag(param.Tag)
	)

	switch tag {
	case Tkuwo:
		url = "http://www.kuwo.cn/api/www/classify/playlist/getTagPlayList"
		url = fmt.Sprintf("%s?pn=%d&rn=%d&id=%s",
			url, param.P, param.N, param.Id)
		raw = &kuwo.WithTagResp{}
	case Tnetease:
		url = fmt.Sprintf("/top/playlist/highquality?limit=%d&cat=%s",
			param.N, param.Id)
		if param.Before != 0 {
			url += "&before=" + strconv.Itoa(int(param.Before))
		}
		raw = &netease.WithTagResp{}
	}

	e := sendAndFormat(&ProxyReqOpt{
		tag:    tag,
		rawUrl: url,
	}, raw, ret)
	if e != nil {
		err = errors.Wrap(err, e.Error())
	}

	for i := range ret.Result.PlList {
		ret.Result.PlList[i].Tag = string(tag)
	}

	return ret, err
}

func PLDetail(c *gin.Context, param *service.PLDetailReq) (*std.DetailResp, error) {
	var (
		err error
		url string
		raw any
		ret = &std.DetailResp{}
		tag = PlatformTag(param.Tag)
	)

	switch tag {
	case Tkuwo:
		url = "http://www.kuwo.cn/api/www/playlist/playListInfo"
		url = fmt.Sprintf("%s?pn=%d&rn=%d&pid=%d&httpsStatus=1",
			url, param.P, param.N, param.Id)
		raw = &kuwo.DetailResp{}
	case Tnetease: // TODO: 分为两次请求：获取歌单详情 -> 获取歌单所有歌曲
		// 暂时只获取所有歌曲，会确实很多歌单的信息：user、name、id、img 等
		url = fmt.Sprintf("/playlist/track/all?id=%d", param.Id)
		raw = &netease.DetailResp{}
	}

	e := sendAndFormat(&ProxyReqOpt{
		tag:    tag,
		rawUrl: url,
	}, raw, ret)
	if e != nil {
		err = errors.Wrap(err, e.Error())
		return nil, err
	}

	for i := range ret.Result.MusicList {
		ret.Result.MusicList[i].Tag = string(tag)
	}

	return ret, err
}

func PLAlbumDetail(c *gin.Context, param *service.PLDetailReq) (*std.DetailResp, error) {
	var (
		err error
		url string
		raw any
		ret = &std.DetailResp{}
		tag = PlatformTag(param.Tag)
	)

	switch tag {
	case Tkuwo:
		url = "http://www.kuwo.cn/api/www/album/albumInfo"
		url = fmt.Sprintf("%s?pn=%d&rn=%d&albumId=%d&httpsStatus=1",
			url, param.P, param.N, param.Id)
		raw = &kuwo.DetailResp{}
	case Tnetease: // TODO: 分为两次请求：获取歌单详情 -> 获取歌单所有歌曲
		// 暂时只获取所有歌曲，会确实很多歌单的信息：user、name、id、img 等
		url = fmt.Sprintf("/album?id=%d", param.Id)
		raw = &netease.DetailResp{}
	}

	e := sendAndFormat(&ProxyReqOpt{
		tag:    tag,
		rawUrl: url,
	}, raw, ret)
	if e != nil {
		err = errors.Wrap(err, e.Error())
		return nil, err
	}

	for i := range ret.Result.MusicList {
		ret.Result.MusicList[i].Tag = string(tag)
	}

	return ret, err
}
