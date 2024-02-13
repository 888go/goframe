// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gtag_test

import (
	"fmt"
	"testing"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gtag"
	"github.com/888go/goframe/util/guid"
)

func Test_Set_Get(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		k := uid类.X生成()
		v := uid类.X生成()
		gtag.X设置值(k, v)
		t.Assert(gtag.Get(k), v)
	})
}

func Test_SetOver_Get(t *testing.T) {
	// panic by Set
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			k  = uid类.X生成()
			v1 = uid类.X生成()
			v2 = uid类.X生成()
		)
		gtag.X设置值(k, v1)
		t.Assert(gtag.Get(k), v1)
		defer func() {
			t.AssertNE(recover(), nil)
		}()
		gtag.X设置值(k, v2)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			k  = uid类.X生成()
			v1 = uid类.X生成()
			v2 = uid类.X生成()
		)
		gtag.SetOver(k, v1)
		t.Assert(gtag.Get(k), v1)
		gtag.SetOver(k, v2)
		t.Assert(gtag.Get(k), v2)
	})
}

func Test_Sets_Get(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			k1 = uid类.X生成()
			k2 = uid类.X生成()
			v1 = uid类.X生成()
			v2 = uid类.X生成()
		)
		gtag.Sets(g.MapStrStr{
			k1: v1,
			k2: v2,
		})
		t.Assert(gtag.Get(k1), v1)
		t.Assert(gtag.Get(k2), v2)
	})
}

func Test_SetsOver_Get(t *testing.T) {
	// panic by Sets
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			k1 = uid类.X生成()
			k2 = uid类.X生成()
			v1 = uid类.X生成()
			v2 = uid类.X生成()
			v3 = uid类.X生成()
		)
		gtag.Sets(g.MapStrStr{
			k1: v1,
			k2: v2,
		})
		t.Assert(gtag.Get(k1), v1)
		t.Assert(gtag.Get(k2), v2)
		defer func() {
			t.AssertNE(recover(), nil)
		}()
		gtag.Sets(g.MapStrStr{
			k1: v3,
			k2: v3,
		})
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			k1 = uid类.X生成()
			k2 = uid类.X生成()
			v1 = uid类.X生成()
			v2 = uid类.X生成()
			v3 = uid类.X生成()
		)
		gtag.SetsOver(g.MapStrStr{
			k1: v1,
			k2: v2,
		})
		t.Assert(gtag.Get(k1), v1)
		t.Assert(gtag.Get(k2), v2)
		gtag.SetsOver(g.MapStrStr{
			k1: v3,
			k2: v3,
		})
		t.Assert(gtag.Get(k1), v3)
		t.Assert(gtag.Get(k2), v3)
	})
}

func Test_Parse(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			k1      = uid类.X生成()
			k2      = uid类.X生成()
			v1      = uid类.X生成()
			v2      = uid类.X生成()
			content = fmt.Sprintf(`this is {%s} and {%s}`, k1, k2)
			expect  = fmt.Sprintf(`this is %s and %s`, v1, v2)
		)
		gtag.Sets(g.MapStrStr{
			k1: v1,
			k2: v2,
		})
		t.Assert(gtag.Parse(content), expect)
	})
}

func Test_SetGlobalEnums(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		oldEnumsJson, err := gtag.GetGlobalEnums()
		t.AssertNil(err)

		err = gtag.SetGlobalEnums(`{"k8s.io/apimachinery/pkg/api/resource.Format": [
        "BinarySI",
        "DecimalExponent",
        "DecimalSI"
    ]}`)
		t.AssertNil(err)
		t.Assert(gtag.GetEnumsByType("k8s.io/apimachinery/pkg/api/resource.Format"), `[
        "BinarySI",
        "DecimalExponent",
        "DecimalSI"
    ]`)
		t.AssertNil(gtag.SetGlobalEnums(oldEnumsJson))
	})
}
