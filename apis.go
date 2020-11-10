package wework

// execGetAccessToken 获取access_token
func (c *WorkwxApp) execGetAccessToken(req reqAccessToken) (respAccessToken, error) {
	var resp respAccessToken
	err := c.executeQyapiGet("/cgi-bin/gettoken", req, &resp, false)
	if err != nil {
		return respAccessToken{}, err
	}

	return resp, nil
}

// execGetUserInfo 获取用户 id
func (c *WorkwxApp) execGetUserInfo(req reqCode) (respUserInfo, error) {
	var resp respUserInfo
	err := c.executeQyapiGet("/cgi-bin/user/getuserinfo", req, &resp, true)
	if err != nil {
		return respUserInfo{}, err
	}

	return resp, nil
}




