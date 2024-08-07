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
		t.AssertEQ(gconv.X取any指针(v), &v)
	})
	gtest.C(t, func(t *gtest.T) {
		var v string = "1"
		t.AssertEQ(gconv.X取文本指针(v), &v)
	})
	gtest.C(t, func(t *gtest.T) {
		var v bool = true
		t.AssertEQ(gconv.X取布尔指针(v), &v)
	})
	gtest.C(t, func(t *gtest.T) {
		var v int = 1
		t.AssertEQ(gconv.X取整数指针(v), &v)
	})
	gtest.C(t, func(t *gtest.T) {
		var v int8 = 1
		t.AssertEQ(gconv.X取整数8位指针(v), &v)
	})
	gtest.C(t, func(t *gtest.T) {
		var v int16 = 1
		t.AssertEQ(gconv.X取整数16位指针(v), &v)
	})
	gtest.C(t, func(t *gtest.T) {
		var v int32 = 1
		t.AssertEQ(gconv.X取整数32位指针(v), &v)
	})
	gtest.C(t, func(t *gtest.T) {
		var v int64 = 1
		t.AssertEQ(gconv.X取整数64位指针(v), &v)
	})
	gtest.C(t, func(t *gtest.T) {
		var v uint = 1
		t.AssertEQ(gconv.X取正整数指针(v), &v)
	})
	gtest.C(t, func(t *gtest.T) {
		var v uint8 = 1
		t.AssertEQ(gconv.X取正整数8位指针(v), &v)
	})
	gtest.C(t, func(t *gtest.T) {
		var v uint16 = 1
		t.AssertEQ(gconv.X取正整数16位指针(v), &v)
	})
	gtest.C(t, func(t *gtest.T) {
		var v uint32 = 1
		t.AssertEQ(gconv.X取正整数32位指针(v), &v)
	})
	gtest.C(t, func(t *gtest.T) {
		var v uint64 = 1
		t.AssertEQ(gconv.X取正整数64位指针(v), &v)
	})
	gtest.C(t, func(t *gtest.T) {
		var v float32 = 1.01
		t.AssertEQ(gconv.X取小数32位指针(v), &v)
	})
	gtest.C(t, func(t *gtest.T) {
		var v float64 = 1.01
		t.AssertEQ(gconv.X取小数64位指针(v), &v)
	})
}
