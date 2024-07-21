	// 版权归GoFrame作者(https:	//goframe.org)所有。保留所有权利。
	//
	// 本源代码形式受MIT许可证条款约束。
	// 如果未随本文件一同分发MIT许可证副本，
	// 您可以在https:	//github.com/gogf/gf处获取。
	// md5:a9832f33b234e3f3

package reflection_test

import (
	"reflect"
	"testing"

	"github.com/gogf/gf/v2/internal/reflection"
	"github.com/gogf/gf/v2/test/gtest"
)

func Test_OriginValueAndKind(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var s = "s"
		out := reflection.OriginValueAndKind(s)
		t.Assert(out.InputKind, reflect.String)
		t.Assert(out.OriginKind, reflect.String)
	})
	gtest.C(t, func(t *gtest.T) {
		var s = "s"
		out := reflection.OriginValueAndKind(&s)
		t.Assert(out.InputKind, reflect.Ptr)
		t.Assert(out.OriginKind, reflect.String)
	})
	gtest.C(t, func(t *gtest.T) {
		var s []int
		out := reflection.OriginValueAndKind(s)
		t.Assert(out.InputKind, reflect.Slice)
		t.Assert(out.OriginKind, reflect.Slice)
	})
	gtest.C(t, func(t *gtest.T) {
		var s []int
		out := reflection.OriginValueAndKind(&s)
		t.Assert(out.InputKind, reflect.Ptr)
		t.Assert(out.OriginKind, reflect.Slice)
	})
}

func Test_OriginTypeAndKind(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var s = "s"
		out := reflection.OriginTypeAndKind(s)
		t.Assert(out.InputKind, reflect.String)
		t.Assert(out.OriginKind, reflect.String)
	})
	gtest.C(t, func(t *gtest.T) {
		var s = "s"
		out := reflection.OriginTypeAndKind(&s)
		t.Assert(out.InputKind, reflect.Ptr)
		t.Assert(out.OriginKind, reflect.String)
	})
	gtest.C(t, func(t *gtest.T) {
		var s []int
		out := reflection.OriginTypeAndKind(s)
		t.Assert(out.InputKind, reflect.Slice)
		t.Assert(out.OriginKind, reflect.Slice)
	})
	gtest.C(t, func(t *gtest.T) {
		var s []int
		out := reflection.OriginTypeAndKind(&s)
		t.Assert(out.InputKind, reflect.Ptr)
		t.Assert(out.OriginKind, reflect.Slice)
	})
}
