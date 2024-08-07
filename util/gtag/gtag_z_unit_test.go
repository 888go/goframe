// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gtag_test

import (
	"fmt"
	"testing"

	"github.com/888go/goframe/frame/g"
	gtest "github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gtag"
	guid "github.com/888go/goframe/util/guid"
)

func Test_Set_Get(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		k := guid.X生成()
		v := guid.X生成()
		gtag.X设置值(k, v)
		t.Assert(gtag.Get(k), v)
	})
}

func Test_SetOver_Get(t *testing.T) {
	// panic by Set
	gtest.C(t, func(t *gtest.T) {
		var (
			k  = guid.X生成()
			v1 = guid.X生成()
			v2 = guid.X生成()
		)
		gtag.X设置值(k, v1)
		t.Assert(gtag.Get(k), v1)
		defer func() {
			t.AssertNE(recover(), nil)
		}()
		gtag.X设置值(k, v2)
	})
	gtest.C(t, func(t *gtest.T) {
		var (
			k  = guid.X生成()
			v1 = guid.X生成()
			v2 = guid.X生成()
		)
		gtag.SetOver(k, v1)
		t.Assert(gtag.Get(k), v1)
		gtag.SetOver(k, v2)
		t.Assert(gtag.Get(k), v2)
	})
}

func Test_Sets_Get(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			k1 = guid.X生成()
			k2 = guid.X生成()
			v1 = guid.X生成()
			v2 = guid.X生成()
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
	gtest.C(t, func(t *gtest.T) {
		var (
			k1 = guid.X生成()
			k2 = guid.X生成()
			v1 = guid.X生成()
			v2 = guid.X生成()
			v3 = guid.X生成()
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
	gtest.C(t, func(t *gtest.T) {
		var (
			k1 = guid.X生成()
			k2 = guid.X生成()
			v1 = guid.X生成()
			v2 = guid.X生成()
			v3 = guid.X生成()
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
	gtest.C(t, func(t *gtest.T) {
		var (
			k1      = guid.X生成()
			k2      = guid.X生成()
			v1      = guid.X生成()
			v2      = guid.X生成()
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
	gtest.C(t, func(t *gtest.T) {
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
