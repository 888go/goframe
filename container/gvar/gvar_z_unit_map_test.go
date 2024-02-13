// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 泛型类_test

import (
	"testing"
	
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/test/gtest"
)

func TestVar_Map(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m := g.Map{
			"k1": "v1",
			"k2": "v2",
		}
		objOne := 泛型类.X创建(m, true)
		t.Assert(objOne.X取Map()["k1"], m["k1"])
		t.Assert(objOne.X取Map()["k2"], m["k2"])
	})
}

func TestVar_MapToMap(t *testing.T) {
// 将map[int]int 转换为 map[string]string
// 创建一个空的原始映射
// 这段Go代码注释翻译成中文注释为：
// ```go
// 将整数到整数的映射（map[int]int）转换为字符串到字符串的映射（map[string]string）
// 初始化一个空的原生映射
	单元测试类.C(t, func(t *单元测试类.T) {
		m1 := g.MapIntInt{}
		m2 := g.MapStrStr{}
		t.Assert(泛型类.X创建(m1).MapToMap(&m2), nil)
		t.Assert(len(m1), len(m2))
	})
	// map[int]int -> map[string]string
	单元测试类.C(t, func(t *单元测试类.T) {
		m1 := g.MapIntInt{
			1: 100,
			2: 200,
		}
		m2 := g.MapStrStr{}
		t.Assert(泛型类.X创建(m1).MapToMap(&m2), nil)
		t.Assert(m2["1"], m1[1])
		t.Assert(m2["2"], m1[2])
	})
	// 将 map[string]interface{} 类型转换为 map[string]string 类型
// 这段注释表明了代码的功能是将一个键为字符串、值为接口类型的映射（map）转换为键同样为字符串但值为字符串类型的映射。
	单元测试类.C(t, func(t *单元测试类.T) {
		m1 := g.Map{
			"k1": "v1",
			"k2": "v2",
		}
		m2 := g.MapStrStr{}
		t.Assert(泛型类.X创建(m1).MapToMap(&m2), nil)
		t.Assert(m2["k1"], m1["k1"])
		t.Assert(m2["k2"], m1["k2"])
	})
	// 将字符串到字符串的映射转换为字符串到接口的映射
	单元测试类.C(t, func(t *单元测试类.T) {
		m1 := g.MapStrStr{
			"k1": "v1",
			"k2": "v2",
		}
		m2 := g.Map{}
		t.Assert(泛型类.X创建(m1).MapToMap(&m2), nil)
		t.Assert(m2["k1"], m1["k1"])
		t.Assert(m2["k2"], m1["k2"])
	})
	// 将 map[string]interface{} 转换为 map[interface{}]interface{}
// 这段注释表明，该代码片段的功能是将键类型为字符串、值类型为接口的映射（map）转换为键和值类型都为接口的映射。
	单元测试类.C(t, func(t *单元测试类.T) {
		m1 := g.MapStrStr{
			"k1": "v1",
			"k2": "v2",
		}
		m2 := g.MapAnyAny{}
		t.Assert(泛型类.X创建(m1).MapToMap(&m2), nil)
		t.Assert(m2["k1"], m1["k1"])
		t.Assert(m2["k2"], m1["k2"])
	})
}

func TestVar_MapStrVar(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m := g.Map{
			"k1": "v1",
			"k2": "v2",
		}
		objOne := 泛型类.X创建(m, true)
		t.Assert(objOne.X取泛型类Map(), "{\"k1\":\"v1\",\"k2\":\"v2\"}")

		objEmpty := 泛型类.X创建(g.Map{})
		t.Assert(objEmpty.X取泛型类Map(), "")
	})
}
