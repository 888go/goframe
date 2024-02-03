// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gconv_test

import (
	"testing"
	
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gconv"
)

func Test_Ptr_Functions(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var v interface{} = 1
		t.AssertEQ(gconv.PtrAny(v), &v)
	})
	gtest.C(t, func(t *gtest.T) {
		var v string = "1"
		t.AssertEQ(gconv.PtrString(v), &v)
	})
	gtest.C(t, func(t *gtest.T) {
		var v bool = true
		t.AssertEQ(gconv.PtrBool(v), &v)
	})
	gtest.C(t, func(t *gtest.T) {
		var v int = 1
		t.AssertEQ(gconv.PtrInt(v), &v)
	})
	gtest.C(t, func(t *gtest.T) {
		var v int8 = 1
		t.AssertEQ(gconv.PtrInt8(v), &v)
	})
	gtest.C(t, func(t *gtest.T) {
		var v int16 = 1
		t.AssertEQ(gconv.PtrInt16(v), &v)
	})
	gtest.C(t, func(t *gtest.T) {
		var v int32 = 1
		t.AssertEQ(gconv.PtrInt32(v), &v)
	})
	gtest.C(t, func(t *gtest.T) {
		var v int64 = 1
		t.AssertEQ(gconv.PtrInt64(v), &v)
	})
	gtest.C(t, func(t *gtest.T) {
		var v uint = 1
		t.AssertEQ(gconv.PtrUint(v), &v)
	})
	gtest.C(t, func(t *gtest.T) {
		var v uint8 = 1
		t.AssertEQ(gconv.PtrUint8(v), &v)
	})
	gtest.C(t, func(t *gtest.T) {
		var v uint16 = 1
		t.AssertEQ(gconv.PtrUint16(v), &v)
	})
	gtest.C(t, func(t *gtest.T) {
		var v uint32 = 1
		t.AssertEQ(gconv.PtrUint32(v), &v)
	})
	gtest.C(t, func(t *gtest.T) {
		var v uint64 = 1
		t.AssertEQ(gconv.PtrUint64(v), &v)
	})
	gtest.C(t, func(t *gtest.T) {
		var v float32 = 1.01
		t.AssertEQ(gconv.PtrFloat32(v), &v)
	})
	gtest.C(t, func(t *gtest.T) {
		var v float64 = 1.01
		t.AssertEQ(gconv.PtrFloat64(v), &v)
	})
}
