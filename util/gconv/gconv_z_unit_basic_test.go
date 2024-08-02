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

func Test_Basic(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		f32 := float32(123.456)
		i64 := int64(1552578474888)
		t.AssertEQ(gconv.Int(f32), int(123))
		t.AssertEQ(gconv.Int8(f32), int8(123))
		t.AssertEQ(gconv.Int16(f32), int16(123))
		t.AssertEQ(gconv.Int32(f32), int32(123))
		t.AssertEQ(gconv.Int64(f32), int64(123))
		t.AssertEQ(gconv.Int64(f32), int64(123))
		t.AssertEQ(gconv.Uint(f32), uint(123))
		t.AssertEQ(gconv.Uint8(f32), uint8(123))
		t.AssertEQ(gconv.Uint16(f32), uint16(123))
		t.AssertEQ(gconv.Uint32(f32), uint32(123))
		t.AssertEQ(gconv.Uint64(f32), uint64(123))
		t.AssertEQ(gconv.Float32(f32), float32(123.456))
		t.AssertEQ(gconv.Float64(i64), float64(i64))
		t.AssertEQ(gconv.Bool(f32), true)
		t.AssertEQ(gconv.String(f32), "123.456")
		t.AssertEQ(gconv.String(i64), "1552578474888")
	})

	gtest.C(t, func(t *gtest.T) {
		s := "-0xFF"
		t.Assert(gconv.Int(s), int64(-0xFF))
	})
}

func Test_Duration(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		d := gconv.Duration("1s")
		t.Assert(d.String(), "1s")
		t.Assert(d.Nanoseconds(), 1000000000)
	})
}

func Test_ConvertWithRefer(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(gconv.ConvertWithRefer("1", 100), 1)
		t.AssertEQ(gconv.ConvertWithRefer("1.01", 1.111), 1.01)
		t.AssertEQ(gconv.ConvertWithRefer("1.01", "1.111"), "1.01")
		t.AssertEQ(gconv.ConvertWithRefer("1.01", false), true)
		t.AssertNE(gconv.ConvertWithRefer("1.01", false), false)
	})
}
