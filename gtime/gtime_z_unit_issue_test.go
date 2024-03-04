// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gtime_test

import (
	"testing"
	
	"github.com/888go/goframe/gtime"
	"github.com/gogf/gf/v2/test/gtest"
)

// 这是Go语言代码中的一行注释，其内容引用了GitHub上gogf/gf项目的一个问题链接：https://github.com/gogf/gf/issues/1681
// 翻译为：
// 参考gogf/gf项目在GitHub上的第1681号问题：https://github.com/gogf/gf/issues/1681
func Test_Issue1681(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gtime.New("2022-03-08T03:01:14-07:00").Local().Time, gtime.New("2022-03-08T10:01:14Z").Local().Time)
		t.Assert(gtime.New("2022-03-08T03:01:14-08:00").Local().Time, gtime.New("2022-03-08T11:01:14Z").Local().Time)
		t.Assert(gtime.New("2022-03-08T03:01:14-09:00").Local().Time, gtime.New("2022-03-08T12:01:14Z").Local().Time)
		t.Assert(gtime.New("2022-03-08T03:01:14+08:00").Local().Time, gtime.New("2022-03-07T19:01:14Z").Local().Time)
	})
}

// 这是GitHub上gogf/gf仓库的第2803个issue的链接
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
