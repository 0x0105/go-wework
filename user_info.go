package wework

import (
	"fmt"
)

func (c *WorkwxApp) GetUserID(code string) (string, error) {
	rsp, err := c.execGetUserInfo(reqCode{Code: code})
	if err != nil {
		return "", nil
	}
	if !rsp.IsOK() {
		return "", fmt.Errorf("get user info failed, errcode:%d, errormsg:%s", rsp.ErrCode, rsp.ErrMsg)
	}
	return rsp.UserId, nil
}
