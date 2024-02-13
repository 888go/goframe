// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 元数据类_test

import (
	"testing"
	
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/gmeta"
)

func TestMeta_Basic(t *testing.T) {
	type A struct {
		元数据类.Meta `tag:"123" orm:"456"`
		Id         int
		Name       string
	}

	单元测试类.C(t, func(t *单元测试类.T) {
		a := &A{
			Id:   100,
			Name: "john",
		}
		t.Assert(len(元数据类.Data(a)), 2)
		t.AssertEQ(元数据类.Get(a, "tag").String(), "123")
		t.AssertEQ(元数据类.Get(a, "orm").String(), "456")
		t.AssertEQ(元数据类.Get(a, "none"), nil)

		b, err := json.Marshal(a)
		t.AssertNil(err)
		t.Assert(b, `{"Id":100,"Name":"john"}`)
	})
}

func TestMeta_Convert_Map(t *testing.T) {
	type A struct {
		元数据类.Meta `tag:"123" orm:"456"`
		Id         int
		Name       string
	}

	单元测试类.C(t, func(t *单元测试类.T) {
		a := &A{
			Id:   100,
			Name: "john",
		}
		m := 转换类.X取Map(a)
		t.Assert(len(m), 2)
		t.Assert(m[`Meta`], nil)
	})
}

func TestMeta_Json(t *testing.T) {
	type A struct {
		元数据类.Meta `tag:"123" orm:"456"`
		Id         int
	}

	单元测试类.C(t, func(t *单元测试类.T) {
		a := &A{
			Id: 100,
		}
		b, err := json.Marshal(a)
		t.AssertNil(err)
		t.Assert(string(b), `{"Id":100}`)
	})
}
