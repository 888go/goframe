// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类_test

import (
	"fmt"
	"strings"
	"testing"
	"time"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/net/ghttp"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/guid"
)

func Test_ConfigFromMap(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m := g.Map{
			"address":         ":12345",
			"listeners":       nil,
			"readTimeout":     "60s",
			"indexFiles":      g.Slice别名{"index.php", "main.php"},
			"errorLogEnabled": true,
			"cookieMaxAge":    "1y",
			"cookieSameSite":  "lax",
			"cookieSecure":    true,
			"cookieHttpOnly":  true,
		}
		config, err := http类.X创建配置对象Map(m)
		t.AssertNil(err)
		d1, _ := time.ParseDuration(转换类.String(m["readTimeout"]))
		d2, _ := time.ParseDuration(转换类.String(m["cookieMaxAge"]))
		t.Assert(config.Address, m["address"])
		t.Assert(config.ReadTimeout, d1)
		t.Assert(config.CookieMaxAge, d2)
		t.Assert(config.IndexFiles, m["indexFiles"])
		t.Assert(config.ErrorLogEnabled, m["errorLogEnabled"])
		t.Assert(config.CookieSameSite, m["cookieSameSite"])
		t.Assert(config.CookieSecure, m["cookieSecure"])
		t.Assert(config.CookieHttpOnly, m["cookieHttpOnly"])
	})
}

func Test_SetConfigWithMap(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m := g.Map{
			"Address": ":8199",
			// "ServerRoot":       "/var/www/我的服务器根目录",
			"IndexFiles":       g.Slice别名{"index.php", "main.php"},
			"AccessLogEnabled": true,
			"ErrorLogEnabled":  true,
			"PProfEnabled":     true,
			"LogPath":          "/tmp/log/MyServerLog",
			"SessionIdName":    "MySessionId",
			"SessionPath":      "/tmp/MySessionStoragePath",
			"SessionMaxAge":    24 * time.Hour,
			"cookieSameSite":   "lax",
			"cookieSecure":     true,
			"cookieHttpOnly":   true,
		}
		s := g.Http类()
		err := s.X设置配置项Map(m)
		t.AssertNil(err)
	})
}

func Test_ClientMaxBodySize(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X创建分组路由("/", func(group *http类.RouterGroup) {
		group.X绑定POST("/", func(r *http类.Request) {
			r.Response.X写响应缓冲区(r.X取请求体文本())
		})
	})
	m := g.Map{
		"ClientMaxBodySize": "1k",
	}
	单元测试类.Assert(s.X设置配置项Map(m), nil)
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)

	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		data := make([]byte, 1056)
		for i := 0; i < 1056; i++ {
			data[i] = 'a'
		}
		t.Assert(
			文本类.X过滤首尾符并含空白(c.Post文本(ctx, "/", data)),
			`Read from request Body failed: http: request body too large`,
		)
	})
}

func Test_ClientMaxBodySize_File(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X创建分组路由("/", func(group *http类.RouterGroup) {
		group.X绑定POST("/", func(r *http类.Request) {
			r.X取上传文件对象("file")
			r.Response.X写响应缓冲区("ok")
		})
	})
	m := g.Map{
		"ErrorLogEnabled":   false,
		"ClientMaxBodySize": "1k",
	}
	单元测试类.Assert(s.X设置配置项Map(m), nil)
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)

	// ok
	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		path := 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
		data := make([]byte, 512)
		for i := 0; i < 512; i++ {
			data[i] = 'a'
		}
		t.Assert(文件类.X写入字节集(path, data), nil)
		defer 文件类.X删除(path)
		t.Assert(
			文本类.X过滤首尾符并含空白(c.Post文本(ctx, "/", "name=john&file=@file:"+path)),
			"ok",
		)
	})

	// too large
	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		path := 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
		data := make([]byte, 1056)
		for i := 0; i < 1056; i++ {
			data[i] = 'a'
		}
		t.Assert(文件类.X写入字节集(path, data), nil)
		defer 文件类.X删除(path)
		t.Assert(
			true,
			strings.Contains(
				文本类.X过滤首尾符并含空白(c.Post文本(ctx, "/", "name=john&file=@file:"+path)),
				"http: request body too large",
			),
		)
	})
}
