package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/Mistsink/kuwo-api/global"
	"github.com/Mistsink/kuwo-api/pkg/decoder"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type requestHeaders struct {
	cookie          string
	csrf            string
	host            string
	referer         string
	contentEncoding string
	// reqId           string
	httpStatus string
	userAgent  string
}

var (
	Req    *requestHeaders
	client *http.Client
	flag   bool
)

func init() {
	if global.AppSetting != nil {
		setup()
		flag = true
	}
}

func setup() {
	Req = &requestHeaders{
		cookie:          global.AppSetting.Cookie,
		csrf:            global.AppSetting.CSRF,
		host:            global.AppSetting.Host,
		referer:         global.AppSetting.Referer,
		contentEncoding: "gzip, deflate, br",
		httpStatus:      "1",
		userAgent:       global.AppSetting.UserAgent,
	}
	client = &http.Client{}
}

func getRaw(c *gin.Context) (*http.Response, error) {
	if !flag {
		setup()
	}

	req, err := http.NewRequest("GET", createUrl(c.Request.URL).String(), nil)
	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("http.NewRequest err: %v", err))
		global.Logger.Error(c, "http.NewRequest err:", err)
		return nil, err
	}

	injectReqHeader(req, c)

	res, err := client.Do(req)
	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("client.Get err: %v", err))
		global.Logger.Error(c, "client.Get err:", err)
		return nil, err
	}
	return res, nil
}

func SaveTmpFile(c *gin.Context, resp *http.Response, fileName string) io.Reader {
	reader, _ := decoder.NewDecoder(resp.Header.Get("Content-Encoding"), resp.Body)
	p, _ := io.ReadAll(reader)
	file, err := os.OpenFile("storage/tmp/"+
		fileName+".json", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		global.Logger.Error(c, err)
	}
	defer file.Close()

	_, _ = file.Write(p)

	return bytes.NewReader(p)
}

func injectReqHeader(req *http.Request, c *gin.Context) {
	req.Header = c.Request.Header
	req.Header.Set("Cookie", Req.cookie)
	req.Header.Set("csrf", Req.csrf)
	req.Header.Set("User-Agent", Req.userAgent)
	req.Header.Set("Host", Req.host)

	if !strings.Contains(req.Header.Get("Referer"), "kuwo") {
		req.Header.Set("Referer", Req.referer)
	}

	//	ensure response header ['content-encoding'] is not nil
	req.Header.Add("Accept-Encoding", Req.contentEncoding)

}

func createUrl(originUrl *url.URL) *url.URL {
	query := originUrl.Query()
	injectQuery(query, uuid.NewV4().String())
	u, _ := url.Parse(originUrl.Scheme + "://" + originUrl.Host + originUrl.Path + "?" + query.Encode())
	return u
}

func injectQuery(queries url.Values, reqId string) {
	queries.Set("httpsStatus", Req.httpStatus)
	queries.Set("reqId", reqId)
}

func rewriteUrl(c *gin.Context, u *url.URL) {
	c.Request.Host = u.Host
	c.Request.URL = u
}

func parseUrl(_url string, c *gin.Context) (u *url.URL, err error) {
	u, err = url.Parse(_url)
	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("url.Parse err: %v", err))
		global.Logger.Error(c, "url.Parse err:", err)
	}
	return
}

func unmarshal(c *gin.Context, raw *http.Response, v any) error {
	contentEncoding := raw.Header.Get("Content-Encoding")
	reader, err := decoder.NewDecoder(contentEncoding, raw.Body)
	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("decoder.NewDecoder err: %v", err))
		global.Logger.Error(c, "decoder.NewDecoder err:", err)
		return err
	}

	if err = json.NewDecoder(reader).Decode(v); err != nil {
		err = errors.Wrap(err, fmt.Sprintf("raw.Decode(json.NewDecoder(reader)) err: %v", err))
		global.Logger.Errorf(c, "raw.Decode(json.NewDecoder(reader)) err: %v", err)
		return err
	}

	return nil
}
