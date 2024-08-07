// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/888go/goframe/frame/g"
	ghttp "github.com/888go/goframe/net/ghttp"
	gfile "github.com/888go/goframe/os/gfile"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
	gstr "github.com/888go/goframe/text/gstr"
	gconv "github.com/888go/goframe/util/gconv"
	guid "github.com/888go/goframe/util/guid"
)

func Test_ConfigFromMap(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
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
		config, err := ghttp.X创建配置对象Map(m)
		t.AssertNil(err)
		d1, _ := time.ParseDuration(gconv.String(m["readTimeout"]))
		d2, _ := time.ParseDuration(gconv.String(m["cookieMaxAge"]))
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
	gtest.C(t, func(t *gtest.T) {
		m := g.Map{
			"Address": ":8199",
									// "ServerRoot":       "/var/www/我的服务器根目录". md5:c70dfe0e8fb98ccb
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
	s := g.Http类(guid.X生成())
	s.X创建分组路由("/", func(group *ghttp.X分组路由) {
		group.X绑定POST("/", func(r *ghttp.Request) {
			r.X响应.X写响应缓冲区(r.X取请求体文本())
		})
	})
	m := g.Map{
		"ClientMaxBodySize": "1k",
	}
	gtest.Assert(s.X设置配置项Map(m), nil)
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)

	gtest.C(t, func(t *gtest.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		data := make([]byte, 1056)
		for i := 0; i < 1056; i++ {
			data[i] = 'a'
		}
		t.Assert(
			gstr.X过滤首尾符并含空白(c.Post文本(ctx, "/", data)),
			`Read from request Body failed: http: request body too large`,
		)
	})
}

func Test_ClientMaxBodySize_File(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.X创建分组路由("/", func(group *ghttp.X分组路由) {
		group.X绑定POST("/", func(r *ghttp.Request) {
			r.X取上传文件对象("file")
			r.X响应.X写响应缓冲区("ok")
		})
	})
	m := g.Map{
		"ErrorLogEnabled":   false,
		"ClientMaxBodySize": "1k",
	}
	gtest.Assert(s.X设置配置项Map(m), nil)
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)

	// ok
	gtest.C(t, func(t *gtest.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		path := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
		data := make([]byte, 512)
		for i := 0; i < 512; i++ {
			data[i] = 'a'
		}
		t.Assert(gfile.X写入字节集(path, data), nil)
		defer gfile.X删除(path)
		t.Assert(
			gstr.X过滤首尾符并含空白(c.Post文本(ctx, "/", "name=john&file=@file:"+path)),
			"ok",
		)
	})

	// too large
	gtest.C(t, func(t *gtest.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		path := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
		data := make([]byte, 1056)
		for i := 0; i < 1056; i++ {
			data[i] = 'a'
		}
		t.Assert(gfile.X写入字节集(path, data), nil)
		defer gfile.X删除(path)
		t.Assert(
			true,
			strings.Contains(
				gstr.X过滤首尾符并含空白(c.Post文本(ctx, "/", "name=john&file=@file:"+path)),
				"http: request body too large",
			),
		)
	})
}
