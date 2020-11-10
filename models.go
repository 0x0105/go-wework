package wework

import (
	"net/url"
)

type reqAccessToken struct {
	CorpID     string
	CorpSecret string
}

var _ urlValuer = reqAccessToken{}

func (x reqAccessToken) intoURLValues() url.Values {
	return url.Values{
		"corpid":     {x.CorpID},
		"corpsecret": {x.CorpSecret},
	}
}

type respCommon struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// IsOK 响应体是否为一次成功请求的响应
//
// 实现依据: https://work.weixin.qq.com/api/doc#10013
//
// > 企业微信所有接口，返回包里都有errcode、errmsg。
// > 开发者需根据errcode是否为0判断是否调用成功(errcode意义请见全局错误码)。
// > 而errmsg仅作参考，后续可能会有变动，因此不可作为是否调用成功的判据。
func (x *respCommon) IsOK() bool {
	return x.ErrCode == 0
}

type respAccessToken struct {
	respCommon

	AccessToken   string `json:"access_token"`
	ExpiresInSecs int64  `json:"expires_in"`
}




// type reqUserInfo struct {
// 	UserTicketInfo UserTicketInfo
// }
//
// type UserTicketInfo struct {
// 	UserTicket string `json:"user_ticket"`
// }
//
// func (x reqUserDetail) intoBody() ([]byte, error) {
// 	result, err := json.Marshal(x.UserTicketInfo)
// 	if err != nil {
// 		// should never happen unless OOM or similar bad things
// 		// TODO: error_chain
// 		return nil, err
// 	}
//
// 	return result, nil
// }
//
// var _ bodyer = reqUserDetail{}
//
//
//
// type respUserDetail struct {
// 	respCommon
// 	// 成员UserID
// 	UserId   string `json:"UserId"`
// 	// 成员UserID
// 	Name string  `json:"name"`
// }



type reqCode struct {
	Code     string
}

var _ urlValuer = reqCode{}

func (x reqCode) intoURLValues() url.Values {
	return url.Values{
		"code":     {x.Code},
	}
}

type respUserInfo struct {
	respCommon
	// 成员UserID
	UserId   string `json:"UserId"`
}
