package playlist

import "github.com/Mistsink/kuwo-api/internal/core/kuwo"

// import (
// 	"encoding/json"
// 	"github.com/Mistsink/kuwo-api/global"
// 	"github.com/Mistsink/kuwo-api/internal/proxy/raw_dao"
// 	"github.com/Mistsink/kuwo-api/internal/service"
// 	"github.com/Mistsink/kuwo-api/pkg/app"
// 	"github.com/Mistsink/kuwo-api/pkg/errcode"
// 	"github.com/Mistsink/kuwo-api/pkg/util"
// 	"github.com/gin-gonic/gin"
// )
// Generated by https://quicktype.io

type RecResp struct {
	kuwo.CommonRes
	Data RecRes `json:"data" yzh:"result"`
}

type RecRes struct {
	List []RecItem `json:"list" yzh:"list"`
}

type RecItem struct {
	Img       string `json:"img" yzh:"rec_img_url"`
	Uname     string `json:"uname" yzh:"rec_user_name"`
	Img700    string `json:"img700" yzh:"rec_img_700"`
	Img300    string `json:"img300" yzh:"rec_img_300"`
	UserName  string `json:"userName"`
	Img500    string `json:"img500" yzh:"rec_img_500"`
	Total     int    `json:"total" yzh:"rec_total_cnt"`
	Name      string `json:"name" yzh:"rec_name"`
	Listencnt int    `json:"listencnt" yzh:"rec_lis_cnt"`
	ID        int    `json:"id" yzh:"rec_id"`
	Desc      string `json:"desc"`
	Info      string `json:"info"`
	Tag       string `json:"tag,omitempty"`
}

// type RawPLRecItem struct {
// 	*raw_dao.RawComItem
// 	Img300 string `json:"img300"`
// 	Img500 string `json:"img500"`
// 	Img700 string `json:"img700"`
// }

// type RawPLRec struct {
// 	*raw_dao.RawComResp
// 	Data struct {
// 		Data []*RawPLRecItem `json:"list"`
// 	} `json:"data"`
// }

// func (r *RawPLRec) Decode(decoder *json.Decoder) error {
// 	return decoder.Decode(r)
// }

// func (r *RawPLRec) TransCode() int {
// 	return raw_dao.TransCode(r.Code)
// }

// func (r *RawPLRec) TransItems() interface{} {
// 	respItems := make([]*service.RespPLRecItem, len(r.Data.Data))

// 	for i, v := range r.Data.Data {
// 		respItems[i] = &service.RespPLRecItem{}
// 		respItems[i].RespComItem = &service.RespComItem{}
// 		util.StructAssign(respItems[i], v)
// 		util.StructAssign(respItems[i].RespComItem, v.RawComItem)
// 	}
// 	return respItems
// }

// func (r *RawPLRec) TransToResp(c *gin.Context) *app.RespJSON {
// 	if r.TransCode() != errcode.Success.Code {
// 		return &app.RespJSON{
// 			Code:  r.TransCode(),
// 			Msg:   r.Msg,
// 			ReqID: r.ReqID,
// 		}
// 	}

// 	var resItems []*service.RespPLRecItem
// 	var ok bool
// 	if resItems, ok = r.TransItems().([]*service.RespPLRecItem); !ok {
// 		global.Logger.Fatal(c, "r.TransItems().([]*RespPLDefaultItem) err")
// 	}

// 	res := &service.RespPLRec{
// 		Data: resItems,
// 	}
// 	ret := &app.RespJSON{
// 		Code:   r.TransCode(),
// 		Msg:    r.Msg,
// 		Result: res,
// 		ReqID:  r.ReqID,
// 	}

// 	return ret
// }

// func NewPLRecRawResp() raw_dao.RawResp {
// 	return &RawPLRec{}
// }