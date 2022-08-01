package kuwo

type CommonRes struct {
	Code      int    `json:"code" yzh:"code"`
	CurTime   int    `json:"curTime"`
	Msg       string `json:"msg" yzh:"msg"`
	ProfileID string `json:"profileId"`
	ReqID     string `json:"reqId" yzh:"req_id"`
	TID       string `json:"tId"`
}
