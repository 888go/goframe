// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 时间类_test

import (
	"testing"
	
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
)

// 这是Go语言代码中的一行注释，其内容引用了GitHub上gogf/gf项目的一个问题链接：https://github.com/gogf/gf/issues/1681
// 翻译为：
// 参考gogf/gf项目在GitHub上的第1681号问题：https://github.com/gogf/gf/issues/1681
func Test_Issue1681(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(时间类.X创建("2022-03-08T03:01:14-07:00").X取本地时区().Time, 时间类.X创建("2022-03-08T10:01:14Z").X取本地时区().Time)
		t.Assert(时间类.X创建("2022-03-08T03:01:14-08:00").X取本地时区().Time, 时间类.X创建("2022-03-08T11:01:14Z").X取本地时区().Time)
		t.Assert(时间类.X创建("2022-03-08T03:01:14-09:00").X取本地时区().Time, 时间类.X创建("2022-03-08T12:01:14Z").X取本地时区().Time)
		t.Assert(时间类.X创建("2022-03-08T03:01:14+08:00").X取本地时区().Time, 时间类.X创建("2022-03-07T19:01:14Z").X取本地时区().Time)
	})
}

// 这是GitHub上gogf/gf仓库的第2803个issue的链接
func Test_Issue2803(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		newTime := 时间类.X创建("2023-07-26").X设置Layout格式("2006-01")
		t.Assert(newTime.Year(), 2023)
		t.Assert(newTime.X取月份(), 7)
		t.Assert(newTime.Day(), 1)
		t.Assert(newTime.Hour(), 0)
		t.Assert(newTime.Minute(), 0)
		t.Assert(newTime.X取秒(), 0)
	})
}
