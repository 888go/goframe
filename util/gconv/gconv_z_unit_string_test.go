// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gconv_test

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gconv"
	"testing"
	"time"
)

type stringStruct1 struct {
	Name string
}

type stringStruct2 struct {
	Name string
}

func (s *stringStruct1) String() string {
	return s.Name
}

func Test_String(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(gconv.String(int(123)), "123")
		t.AssertEQ(gconv.String(int(-123)), "-123")
		t.AssertEQ(gconv.String(int8(123)), "123")
		t.AssertEQ(gconv.String(int8(-123)), "-123")
		t.AssertEQ(gconv.String(int16(123)), "123")
		t.AssertEQ(gconv.String(int16(-123)), "-123")
		t.AssertEQ(gconv.String(int32(123)), "123")
		t.AssertEQ(gconv.String(int32(-123)), "-123")
		t.AssertEQ(gconv.String(int64(123)), "123")
		t.AssertEQ(gconv.String(int64(-123)), "-123")
		t.AssertEQ(gconv.String(int64(1552578474888)), "1552578474888")
		t.AssertEQ(gconv.String(int64(-1552578474888)), "-1552578474888")

		t.AssertEQ(gconv.String(uint(123)), "123")
		t.AssertEQ(gconv.String(uint8(123)), "123")
		t.AssertEQ(gconv.String(uint16(123)), "123")
		t.AssertEQ(gconv.String(uint32(123)), "123")
		t.AssertEQ(gconv.String(uint64(155257847488898765)), "155257847488898765")

		t.AssertEQ(gconv.String(float32(123.456)), "123.456")
		t.AssertEQ(gconv.String(float32(-123.456)), "-123.456")
		t.AssertEQ(gconv.String(float64(1552578474888.456)), "1552578474888.456")
		t.AssertEQ(gconv.String(float64(-1552578474888.456)), "-1552578474888.456")

		t.AssertEQ(gconv.String(true), "true")
		t.AssertEQ(gconv.String(false), "false")

		t.AssertEQ(gconv.String([]byte("bytes")), "bytes")

		t.AssertEQ(gconv.String(stringStruct1{"john"}), `{"Name":"john"}`)
		t.AssertEQ(gconv.String(&stringStruct1{"john"}), "john")

		t.AssertEQ(gconv.String(stringStruct2{"john"}), `{"Name":"john"}`)
		t.AssertEQ(gconv.String(&stringStruct2{"john"}), `{"Name":"john"}`)

		var nilTime *time.Time = nil
		t.AssertEQ(gconv.String(nilTime), "")
		var nilGTime *gtime.Time = nil
		t.AssertEQ(gconv.String(nilGTime), "")
	})
}
