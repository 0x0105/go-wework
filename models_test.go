package wework

import (
	"encoding/json"
	"net/url"
	"testing"

	c "github.com/smartystreets/goconvey/convey"
)

func TestReqAccessToken(t *testing.T) {
	c.Convey("构造一个 reqAccessToken", t, func() {
		a := reqAccessToken{
			CorpID:     "foo",
			CorpSecret: "bar",
		}

		c.Convey("序列化结果应该符合预期", func() {
			v := a.intoURLValues()

			expected := url.Values{
				"corpid":     []string{"foo"},
				"corpsecret": []string{"bar"},
			}
			c.So(v, c.ShouldResemble, expected)
		})
	})
}

func TestRespCommon(t *testing.T) {
	payloadOk := []byte(`{"errcode":0,"errmsg":"ok"}`)
	payloadErr := []byte(`{"errcode":40014,"errmsg":"invalid access_token"}`)

	c.Convey("构造一个成功响应的 respCommon", t, func() {
		var a respCommon
		err := json.Unmarshal(payloadOk, &a)

		c.Convey("应该能成功反序列化", func() {
			c.So(err, c.ShouldBeNil)

			c.Convey("反序列化之后字段应该正确对应", func() {
				c.So(a.ErrCode, c.ShouldEqual, 0)
				c.So(a.ErrMsg, c.ShouldEqual, "ok")
			})

			c.Convey("IsOk 应该为真", func() {
				c.So(a.IsOK(), c.ShouldBeTrue)
			})
		})
	})

	c.Convey("构造一个失败响应的 respCommon", t, func() {
		var a respCommon
		err := json.Unmarshal(payloadErr, &a)

		c.Convey("应该能成功反序列化", func() {
			c.So(err, c.ShouldBeNil)

			c.Convey("反序列化之后字段应该正确对应", func() {
				c.So(a.ErrCode, c.ShouldEqual, 40014)
				c.So(a.ErrMsg, c.ShouldEqual, "invalid access_token")
			})

			c.Convey("IsOk 应该为假", func() {
				c.So(a.IsOK(), c.ShouldBeFalse)
			})
		})
	})
}
