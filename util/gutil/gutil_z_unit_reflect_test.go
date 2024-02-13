// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 工具类_test

import (
	"reflect"
	"testing"
	
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gutil"
)

func Test_OriginValueAndKind(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var s = "s"
		out := 工具类.OriginValueAndKind(s)
		t.Assert(out.InputKind, reflect.String)
		t.Assert(out.OriginKind, reflect.String)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var s = "s"
		out := 工具类.OriginValueAndKind(&s)
		t.Assert(out.InputKind, reflect.Ptr)
		t.Assert(out.OriginKind, reflect.String)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var s []int
		out := 工具类.OriginValueAndKind(s)
		t.Assert(out.InputKind, reflect.Slice)
		t.Assert(out.OriginKind, reflect.Slice)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var s []int
		out := 工具类.OriginValueAndKind(&s)
		t.Assert(out.InputKind, reflect.Ptr)
		t.Assert(out.OriginKind, reflect.Slice)
	})
}

func Test_OriginTypeAndKind(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var s = "s"
		out := 工具类.OriginTypeAndKind(s)
		t.Assert(out.InputKind, reflect.String)
		t.Assert(out.OriginKind, reflect.String)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var s = "s"
		out := 工具类.OriginTypeAndKind(&s)
		t.Assert(out.InputKind, reflect.Ptr)
		t.Assert(out.OriginKind, reflect.String)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var s []int
		out := 工具类.OriginTypeAndKind(s)
		t.Assert(out.InputKind, reflect.Slice)
		t.Assert(out.OriginKind, reflect.Slice)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var s []int
		out := 工具类.OriginTypeAndKind(&s)
		t.Assert(out.InputKind, reflect.Ptr)
		t.Assert(out.OriginKind, reflect.Slice)
	})
}
