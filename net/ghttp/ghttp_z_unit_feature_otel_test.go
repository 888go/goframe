// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package http类_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/888go/goframe/frame/g"
	ghttp "github.com/888go/goframe/net/ghttp"
	"github.com/888go/goframe/net/gtrace"
	gtest "github.com/888go/goframe/test/gtest"
	guid "github.com/888go/goframe/util/guid"
)

func Test_OTEL_TraceID(t *testing.T) {
	var (
		traceId string
	)
	s := g.Server(guid.S())
	s.BindHandler("/", func(r *ghttp.Request) {
		traceId = gtrace.GetTraceID(r.Context())
		r.Response.Write(r.GetUrl())
	})
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		prefix := fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort())
		client := g.Client()
		client.SetBrowserMode(true)
		client.SetPrefix(prefix)
		res, err := client.Get(ctx, "/")
		t.AssertNil(err)
		defer res.Close()

		t.Assert(res.Header.Get("Trace-Id"), traceId)
	})
}
