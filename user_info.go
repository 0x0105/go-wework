package wework

func (c *WorkwxApp) GetUserID(code string) (string, error) {
	rsp, err := c.execGetUserInfo(reqCode{Code: code})
	if err != nil {
		return "", nil
	}
	return rsp.UserId, nil
}
