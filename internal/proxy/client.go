package proxy

import (
	"encoding/json"
	"strings"

	"github.com/Mistsink/kuwo-api/global"
	"github.com/go-resty/resty/v2"
)

var (
	proxyClient = make(map[string]*resty.Client, 0)
)

func initProxyClient() {
	if len(global.ProxySetting.ProxyItems) == 0 {
		panic("ProxySetting has no item")
	}
	for _, proxyS := range global.ProxySetting.ProxyItems {
		client := resty.New()

		if proxyS.ProxyAddr == "" {
			client.SetHeader("Cookie", proxyS.Cookie)
			client.SetHeader("csrf", proxyS.CSRF)
			client.SetHeader("Referer", proxyS.Referer)
			client.SetHeader("User-Agent", proxyS.UserAgent)
		} else {
			client.SetBaseURL(proxyS.ProxyAddr)
		}

		client.SetHeader("Accept-Encoding", "gzip, deflate, br")

		proxyClient[proxyS.Name] = client
	}
}

func sendProxy(reqOpt *ProxyReqOpt, target any) error {
	if reqOpt.method == "" {
		reqOpt.method = "GET"
	} else {
		reqOpt.method = strings.ToUpper(reqOpt.method)
	}
	tag := strings.ToLower(string(reqOpt.tag))

	client := proxyClient[tag]
	r := client.R().SetResult(target)

	res, err := r.Execute(reqOpt.method, reqOpt.url.String())

	// 若 res 的 content-type 不是 json 或 xml，则不会自动 Unmarshal
	if ctnType := res.Header().Get("Content-Type"); !strings.Contains(ctnType, "xml") && !strings.Contains(ctnType, "json") {
		err = json.Unmarshal(res.Body(), target)
	}
	return err
}
