// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 转换类_test

import (
	"testing"

	gtest "github.com/888go/goframe/test/gtest"
	gconv "github.com/888go/goframe/util/gconv"
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
