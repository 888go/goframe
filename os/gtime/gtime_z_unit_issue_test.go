// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 时间类_test

import (
	"testing"
	"time"

	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
)

//github.com/gogf/gf/issues/1681. md5:2eac1ca19dcb940c
func Test_Issue1681(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gtime.New("2022-03-08T03:01:14-07:00").Local().Time, gtime.New("2022-03-08T10:01:14Z").Local().Time)
		t.Assert(gtime.New("2022-03-08T03:01:14-08:00").Local().Time, gtime.New("2022-03-08T11:01:14Z").Local().Time)
		t.Assert(gtime.New("2022-03-08T03:01:14-09:00").Local().Time, gtime.New("2022-03-08T12:01:14Z").Local().Time)
		t.Assert(gtime.New("2022-03-08T03:01:14+08:00").Local().Time, gtime.New("2022-03-07T19:01:14Z").Local().Time)
	})
}

//github.com/gogf/gf/issues/2803。gf可能是Go语言的一个库（gogf）的简称，issue号2803表示该库中的一个已知问题或特性请求。具体的内容需要查看该issue页面以获取详细信息。 md5:1ee3164a38e80927
func Test_Issue2803(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		newTime := gtime.New("2023-07-26").LayoutTo("2006-01")
		t.Assert(newTime.Year(), 2023)
		t.Assert(newTime.Month(), 7)
		t.Assert(newTime.Day(), 1)
		t.Assert(newTime.Hour(), 0)
		t.Assert(newTime.Minute(), 0)
		t.Assert(newTime.Second(), 0)
	})
}

//github.com/gogf/gf/issues/3558。这可能是关于GF（Golang Foundation）项目的一个问题或者建议，3558是该问题或拉取请求的编号。具体的内容需要查看相关页面以了解详细信息。 md5:53534467109d62b9
func Test_Issue3558(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timeStr := "1880-10-24T00:00:00+08:05"
		gfTime := gtime.NewFromStr(timeStr)
		t.Assert(gfTime.Year(), 1880)
		t.Assert(gfTime.Month(), 10)
		t.Assert(gfTime.Day(), 24)
		t.Assert(gfTime.Hour(), 0)
		t.Assert(gfTime.Minute(), 0)
		t.Assert(gfTime.Second(), 0)

		stdTime, err := time.Parse(time.RFC3339, timeStr)
		t.AssertNil(err)
		stdTimeFormat := stdTime.Format("2006-01-02 15:04:05")
		gfTimeFormat := gfTime.Format("Y-m-d H:i:s")
		t.Assert(gfTimeFormat, stdTimeFormat)
	})
	gtest.C(t, func(t *gtest.T) {
		timeStr := "1880-10-24T00:00:00-08:05"
		gfTime := gtime.NewFromStr(timeStr)
		t.Assert(gfTime.Year(), 1880)
		t.Assert(gfTime.Month(), 10)
		t.Assert(gfTime.Day(), 24)
		t.Assert(gfTime.Hour(), 0)
		t.Assert(gfTime.Minute(), 0)
		t.Assert(gfTime.Second(), 0)
		stdTime, err := time.Parse(time.RFC3339, timeStr)
		t.AssertNil(err)
		stdTimeFormat := stdTime.Format("2006-01-02 15:04:05")
		gfTimeFormat := gfTime.Format("Y-m-d H:i:s")
		t.Assert(gfTimeFormat, stdTimeFormat)
	})
}
