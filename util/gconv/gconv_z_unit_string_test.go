// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 转换类_test

import (
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gconv"
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
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertEQ(转换类.String(int(123)), "123")
		t.AssertEQ(转换类.String(int(-123)), "-123")
		t.AssertEQ(转换类.String(int8(123)), "123")
		t.AssertEQ(转换类.String(int8(-123)), "-123")
		t.AssertEQ(转换类.String(int16(123)), "123")
		t.AssertEQ(转换类.String(int16(-123)), "-123")
		t.AssertEQ(转换类.String(int32(123)), "123")
		t.AssertEQ(转换类.String(int32(-123)), "-123")
		t.AssertEQ(转换类.String(int64(123)), "123")
		t.AssertEQ(转换类.String(int64(-123)), "-123")
		t.AssertEQ(转换类.String(int64(1552578474888)), "1552578474888")
		t.AssertEQ(转换类.String(int64(-1552578474888)), "-1552578474888")

		t.AssertEQ(转换类.String(uint(123)), "123")
		t.AssertEQ(转换类.String(uint8(123)), "123")
		t.AssertEQ(转换类.String(uint16(123)), "123")
		t.AssertEQ(转换类.String(uint32(123)), "123")
		t.AssertEQ(转换类.String(uint64(155257847488898765)), "155257847488898765")

		t.AssertEQ(转换类.String(float32(123.456)), "123.456")
		t.AssertEQ(转换类.String(float32(-123.456)), "-123.456")
		t.AssertEQ(转换类.String(float64(1552578474888.456)), "1552578474888.456")
		t.AssertEQ(转换类.String(float64(-1552578474888.456)), "-1552578474888.456")

		t.AssertEQ(转换类.String(true), "true")
		t.AssertEQ(转换类.String(false), "false")

		t.AssertEQ(转换类.String([]byte("bytes")), "bytes")

		t.AssertEQ(转换类.String(stringStruct1{"john"}), `{"Name":"john"}`)
		t.AssertEQ(转换类.String(&stringStruct1{"john"}), "john")

		t.AssertEQ(转换类.String(stringStruct2{"john"}), `{"Name":"john"}`)
		t.AssertEQ(转换类.String(&stringStruct2{"john"}), `{"Name":"john"}`)

		var nilTime *time.Time = nil
		t.AssertEQ(转换类.String(nilTime), "")
		var nilGTime *时间类.Time = nil
		t.AssertEQ(转换类.String(nilGTime), "")
	})
}
