// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gtrace_test

import (
	"context"
	"net/http"
	"strings"
	"testing"

	gcompress "github.com/888go/goframe/encoding/gcompress"

	"github.com/888go/goframe/net/gtrace"
	gtest "github.com/888go/goframe/test/gtest"
	gstr "github.com/888go/goframe/text/gstr"
)

func TestWithTraceID(t *testing.T) {
	var (
		ctx  = context.Background()
		uuid = `a323f910-f690-11ec-963d-79c0b7fcf119`
	)
	gtest.C(t, func(t *gtest.T) {
		newCtx, err := gtrace.WithTraceID(ctx, uuid)
		t.AssertNE(err, nil)
		t.Assert(newCtx, ctx)
	})
	gtest.C(t, func(t *gtest.T) {
		var traceId = gstr.X替换(uuid, "-", "")
		newCtx, err := gtrace.WithTraceID(ctx, traceId)
		t.AssertNil(err)
		t.AssertNE(newCtx, ctx)
		t.Assert(gtrace.GetTraceID(ctx), "")
		t.Assert(gtrace.GetTraceID(newCtx), traceId)
	})
}

func TestWithUUID(t *testing.T) {
	var (
		ctx  = context.Background()
		uuid = `a323f910-f690-11ec-963d-79c0b7fcf119`
	)
	gtest.C(t, func(t *gtest.T) {
		newCtx, err := gtrace.WithTraceID(ctx, uuid)
		t.AssertNE(err, nil)
		t.Assert(newCtx, ctx)
	})
	gtest.C(t, func(t *gtest.T) {
		newCtx, err := gtrace.WithUUID(ctx, uuid)
		t.AssertNil(err)
		t.AssertNE(newCtx, ctx)
		t.Assert(gtrace.GetTraceID(ctx), "")
		t.Assert(gtrace.GetTraceID(newCtx), gstr.X替换(uuid, "-", ""))
	})
}

func TestSafeContent(t *testing.T) {
	var (
		defText    = "中"
		shortData  = strings.Repeat(defText, gtrace.MaxContentLogSize()-1)
		standData  = strings.Repeat(defText, gtrace.MaxContentLogSize())
		longData   = strings.Repeat(defText, gtrace.MaxContentLogSize()+1)
		header     = http.Header{}
		gzipHeader = http.Header{
			"Content-Encoding": []string{"gzip"},
		}
	)

	// safe content
	gtest.C(t, func(t *gtest.T) {

		t1, err := gtrace.SafeContentForHttp([]byte(shortData), header)
		t.AssertNil(err)
		t.Assert(t1, shortData)
		t.Assert(gtrace.SafeContent([]byte(shortData)), shortData)

		t2, err := gtrace.SafeContentForHttp([]byte(standData), header)
		t.AssertNil(err)
		t.Assert(t2, standData)
		t.Assert(gtrace.SafeContent([]byte(standData)), standData)

		t3, err := gtrace.SafeContentForHttp([]byte(longData), header)
		t.AssertNil(err)
		t.Assert(t3, standData+"...")
		t.Assert(gtrace.SafeContent([]byte(longData)), standData+"...")
	})

	// compress content
	var (
		compressShortData, _ = gcompress.Gzip压缩字节集([]byte(shortData))
		compressStandData, _ = gcompress.Gzip压缩字节集([]byte(standData))
		compressLongData, _  = gcompress.Gzip压缩字节集([]byte(longData))
	)
	gtest.C(t, func(t *gtest.T) {

		t1, err := gtrace.SafeContentForHttp(compressShortData, gzipHeader)
		t.AssertNil(err)
		t.Assert(t1, shortData)

		t2, err := gtrace.SafeContentForHttp(compressStandData, gzipHeader)
		t.AssertNil(err)
		t.Assert(t2, standData)

		t3, err := gtrace.SafeContentForHttp(compressLongData, gzipHeader)
		t.AssertNil(err)
		t.Assert(t3, standData+"...")
	})
}
