			// 版权归GoFrame作者(https:			//goframe.org)所有。保留所有权利。
			//
			// 本源代码形式受MIT许可证条款约束。
			// 如果未随本文件一同分发MIT许可证副本，
			// 您可以在https:			//github.com/gogf/gf处获取。
			// md5:a9832f33b234e3f3
package gurl_test//bm:url类_test

import (
	"net/url"
	"testing"

	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/test/gtest"
)

var (
	urlStr       = `https://golang.org/x/crypto?go-get=1 +`
	urlEncode    = `https%3A%2F%2Fgolang.org%2Fx%2Fcrypto%3Fgo-get%3D1+%2B`
	rawUrlEncode = `https%3A%2F%2Fgolang.org%2Fx%2Fcrypto%3Fgo-get%3D1%20%2B`
)

func TestEncodeAndDecode(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gurl.Encode(urlStr), urlEncode)

		res, err := gurl.Decode(urlEncode)
		if err != nil {
			t.Errorf("decode failed. %v", err)
			return
		}
		t.Assert(res, urlStr)
	})
}

func TestRowEncodeAndDecode(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gurl.RawEncode(urlStr), rawUrlEncode)

		res, err := gurl.RawDecode(rawUrlEncode)
		if err != nil {
			t.Errorf("decode failed. %v", err)
			return
		}
		t.Assert(res, urlStr)
	})
}

func TestBuildQuery(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		src := url.Values{
			"a": {"a2", "a1"},
			"b": {"b2", "b1"},
			"c": {"c1", "c2"},
		}
		expect := "a=a2&a=a1&b=b2&b=b1&c=c1&c=c2"
		t.Assert(gurl.BuildQuery(src), expect)
	})
}

func TestParseURL(t *testing.T) {
	src := `http://username:password@hostname:9090/path?arg=value#anchor`
	expect := map[string]string{
		"scheme":   "http",
		"host":     "hostname",
		"port":     "9090",
		"user":     "username",
		"pass":     "password",
		"path":     "/path",
		"query":    "arg=value",
		"fragment": "anchor",
	}

	gtest.C(t, func(t *gtest.T) {
		component := 0
		for k, v := range []string{"all", "scheme", "host", "port", "user", "pass", "path", "query", "fragment"} {
			if v == "all" {
				component = -1
			} else {
				component = 1 << (uint(k - 1))
			}

			res, err := gurl.ParseURL(src, component)
			if err != nil {
				t.Errorf("ParseURL failed. component:%v, err:%v", component, err)
				return
			}

			if v == "all" {
				t.Assert(res, expect)
			} else {
				t.Assert(res[v], expect[v])
			}

		}
	})
}
