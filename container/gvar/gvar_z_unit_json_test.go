// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 泛型类_test

import (
	"math"
	"testing"
	
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/test/gtest"
)

func TestVar_Json(t *testing.T) {
	// Marshal
	单元测试类.C(t, func(t *单元测试类.T) {
		s := "i love gf"
		v := 泛型类.X创建(s)
		b1, err1 := json.Marshal(v)
		b2, err2 := json.Marshal(s)
		t.Assert(err1, err2)
		t.Assert(b1, b2)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		s := int64(math.MaxInt64)
		v := 泛型类.X创建(s)
		b1, err1 := json.Marshal(v)
		b2, err2 := json.Marshal(s)
		t.Assert(err1, err2)
		t.Assert(b1, b2)
	})

	// Unmarshal
	单元测试类.C(t, func(t *单元测试类.T) {
		s := "i love gf"
		v := 泛型类.X创建(nil)
		b, err := json.Marshal(s)
		t.AssertNil(err)

		err = json.UnmarshalUseNumber(b, v)
		t.AssertNil(err)
		t.Assert(v.String(), s)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		var v 泛型类.Var
		s := "i love gf"
		b, err := json.Marshal(s)
		t.AssertNil(err)

		err = json.UnmarshalUseNumber(b, &v)
		t.AssertNil(err)
		t.Assert(v.String(), s)
	})
}
