// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/net/ghttp"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/guid"
)

type testWrapStdHTTPStruct struct {
	T    *单元测试类.T
	text string
}

func (t *testWrapStdHTTPStruct) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	t.T.Assert(req.Method, "POST")
	t.T.Assert(req.URL.Path, "/api/wraph")
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, t.text)
}

func Test_Server_Wrap_Handler(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s := g.Http类(uid类.X生成())
		str1 := "hello"
		str2 := "hello again"
		s.X创建分组路由("/api", func(group *http类.X分组路由) {
			group.X绑定GET("/wrapf", http类.WrapF(func(w http.ResponseWriter, req *http.Request) {
				t.Assert(req.Method, "GET")
				t.Assert(req.URL.Path, "/api/wrapf")
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprint(w, str1)
			}))

			group.X绑定POST("/wraph", http类.WrapH(&testWrapStdHTTPStruct{t, str2}))
		})

		s.SetDumpRouterMap(false)
		s.X开始监听()
		defer s.X关闭当前服务()

		time.Sleep(100 * time.Millisecond)
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d/api", s.X取已监听端口()))

		response, err := client.Get响应对象(ctx, "/wrapf")
		t.AssertNil(err)
		defer response.X关闭()
		t.Assert(response.StatusCode, http.StatusBadRequest)
		t.Assert(response.X取响应文本(), str1)

		response2, err := client.Post响应对象(ctx, "/wraph")
		t.AssertNil(err)
		defer response2.X关闭()
		t.Assert(response2.StatusCode, http.StatusInternalServerError)
		t.Assert(response2.X取响应文本(), str2)
	})
}
