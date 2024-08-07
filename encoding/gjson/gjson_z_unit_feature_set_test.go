// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package json类_test

import (
	"bytes"
	"testing"

	garray "github.com/888go/goframe/container/garray"
	gjson "github.com/888go/goframe/encoding/gjson"
	"github.com/888go/goframe/frame/g"
	gtest "github.com/888go/goframe/test/gtest"
	gstr "github.com/888go/goframe/text/gstr"
)

func Test_Set1(t *testing.T) {
	e := []byte(`{"k1":{"k11":[1,2,3]},"k2":"v2"}`)
	p := gjson.X创建(map[string]string{
		"k1": "v1",
		"k2": "v2",
	})
	p.X设置值("k1.k11", []int{1, 2, 3})
	if c, err := p.X取json字节集(); err == nil {

		if !bytes.Equal(c, []byte(`{"k1":{"k11":[1,2,3]},"k2":"v2"}`)) {
			t.Error("expect:", string(e))
		}
	} else {
		t.Error(err)
	}
}

func Test_Set2(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		e := `[[null,1]]`
		p := gjson.X创建([]string{"a"})
		p.X设置值("0.1", 1)
		s := p.X取json文本PANI()
		t.Assert(s, e)
	})
}

func Test_Set3(t *testing.T) {
	e := []byte(`{"kv":{"k1":"v1"}}`)
	p := gjson.X创建([]string{"a"})
	p.X设置值("kv", map[string]string{
		"k1": "v1",
	})
	if c, err := p.X取json字节集(); err == nil {
		if !bytes.Equal(c, e) {
			t.Error("expect:", string(e))
		}
	} else {
		t.Error(err)
	}
}

func Test_Set4(t *testing.T) {
	e := []byte(`["a",[{"k1":"v1"}]]`)
	p := gjson.X创建([]string{"a"})
	p.X设置值("1.0", map[string]string{
		"k1": "v1",
	})
	if c, err := p.X取json字节集(); err == nil {

		if !bytes.Equal(c, e) {
			t.Error("expect:", string(e))
		}
	} else {
		t.Error(err)
	}
}

func Test_Set5(t *testing.T) {
	e := []byte(`[[[[[[[[[[[[[[[[[[[[[1,2,3]]]]]]]]]]]]]]]]]]]]]`)
	p := gjson.X创建([]string{"a"})
	p.X设置值("0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0", []int{1, 2, 3})
	if c, err := p.X取json字节集(); err == nil {

		if !bytes.Equal(c, e) {
			t.Error("expect:", string(e))
		}
	} else {
		t.Error(err)
	}
}

func Test_Set6(t *testing.T) {
	e := []byte(`["a",[1,2,3]]`)
	p := gjson.X创建([]string{"a"})
	p.X设置值("1", []int{1, 2, 3})
	if c, err := p.X取json字节集(); err == nil {

		if !bytes.Equal(c, e) {
			t.Error("expect:", string(e))
		}
	} else {
		t.Error(err)
	}
}

func Test_Set7(t *testing.T) {
	e := []byte(`{"0":[null,[1,2,3]],"k1":"v1","k2":"v2"}`)
	p := gjson.X创建(map[string]string{
		"k1": "v1",
		"k2": "v2",
	})
	p.X设置值("0.1", []int{1, 2, 3})
	if c, err := p.X取json字节集(); err == nil {

		if !bytes.Equal(c, e) {
			t.Error("expect:", string(e))
		}
	} else {
		t.Error(err)
	}
}

func Test_Set8(t *testing.T) {
	e := []byte(`{"0":[[[[[[null,[1,2,3]]]]]]],"k1":"v1","k2":"v2"}`)
	p := gjson.X创建(map[string]string{
		"k1": "v1",
		"k2": "v2",
	})
	p.X设置值("0.0.0.0.0.0.1", []int{1, 2, 3})
	if c, err := p.X取json字节集(); err == nil {

		if !bytes.Equal(c, e) {
			t.Error("expect:", string(e))
		}
	} else {
		t.Error(err)
	}
}

func Test_Set9(t *testing.T) {
	e := []byte(`{"k1":[null,[1,2,3]],"k2":"v2"}`)
	p := gjson.X创建(map[string]string{
		"k1": "v1",
		"k2": "v2",
	})
	p.X设置值("k1.1", []int{1, 2, 3})
	if c, err := p.X取json字节集(); err == nil {

		if !bytes.Equal(c, e) {
			t.Error("expect:", string(e))
		}
	} else {
		t.Error(err)
	}
}

func Test_Set10(t *testing.T) {
	e := []byte(`{"a":{"b":{"c":1}}}`)
	p := gjson.X创建(nil)
	p.X设置值("a.b.c", 1)
	if c, err := p.X取json字节集(); err == nil {

		if !bytes.Equal(c, e) {
			t.Error("expect:", string(e))
		}
	} else {
		t.Error(err)
	}
}

func Test_Set11(t *testing.T) {
	e := []byte(`{"a":{"b":{}}}`)
	p, _ := gjson.X加载并自动识别格式([]byte(`{"a":{"b":{"c":1}}}`))
	p.X删除("a.b.c")
	if c, err := p.X取json字节集(); err == nil {

		if !bytes.Equal(c, e) {
			t.Error("expect:", string(e))
		}
	} else {
		t.Error(err)
	}
}

func Test_Set12(t *testing.T) {
	e := []byte(`[0,1]`)
	p := gjson.X创建(nil)
	p.X设置值("0", 0)
	p.X设置值("1", 1)
	if c, err := p.X取json字节集(); err == nil {

		if !bytes.Equal(c, e) {
			t.Error("expect:", string(e))
		}
	} else {
		t.Error(err)
	}
}

func Test_Set13(t *testing.T) {
	e := []byte(`{"array":[0,1]}`)
	p := gjson.X创建(nil)
	p.X设置值("array.0", 0)
	p.X设置值("array.1", 1)
	if c, err := p.X取json字节集(); err == nil {

		if !bytes.Equal(c, e) {
			t.Error("expect:", string(e))
		}
	} else {
		t.Error(err)
	}
}

func Test_Set14(t *testing.T) {
	e := []byte(`{"f":{"a":1}}`)
	p := gjson.X创建(nil)
	p.X设置值("f", "m")
	p.X设置值("f.a", 1)
	if c, err := p.X取json字节集(); err == nil {

		if !bytes.Equal(c, e) {
			t.Error("expect:", string(e))
		}
	} else {
		t.Error(err)
	}
}

func Test_Set15(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		j := gjson.X创建(nil)

		t.Assert(j.X设置值("root.0.k1", "v1"), nil)
		t.Assert(j.X设置值("root.1.k2", "v2"), nil)
		t.Assert(j.X设置值("k", "v"), nil)

		s, err := j.X取json文本()
		t.AssertNil(err)
		t.Assert(
			gstr.X是否包含(s, `"root":[{"k1":"v1"},{"k2":"v2"}`) ||
				gstr.X是否包含(s, `"root":[{"k2":"v2"},{"k1":"v1"}`),
			true,
		)
		t.Assert(
			gstr.X是否包含(s, `{"k":"v"`) ||
				gstr.X是否包含(s, `"k":"v"}`),
			true,
		)
	})
}

func Test_Set16(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		j := gjson.X创建(nil)

		t.Assert(j.X设置值("processors.0.set.0value", "1"), nil)
		t.Assert(j.X设置值("processors.0.set.0field", "2"), nil)
		t.Assert(j.X设置值("description", "3"), nil)

		s, err := j.X取json文本()
		t.AssertNil(err)
		t.Assert(
			gstr.X是否包含(s, `"processors":[{"set":{"0field":"2","0value":"1"}}]`) ||
				gstr.X是否包含(s, `"processors":[{"set":{"0value":"1","0field":"2"}}]`),
			true,
		)
		t.Assert(
			gstr.X是否包含(s, `{"description":"3"`) || gstr.X是否包含(s, `"description":"3"}`),
			true,
		)
	})
}

func Test_Set17(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		j := gjson.X创建(nil)

		t.Assert(j.X设置值("0.k1", "v1"), nil)
		t.Assert(j.X设置值("1.k2", "v2"), nil)
				// 覆盖之前的切片。 md5:7ecc228788fbb89e
		t.Assert(j.X设置值("k", "v"), nil)

		s, err := j.X取json文本()
		t.AssertNil(err)
		t.Assert(s, `{"k":"v"}`)
	})
}

func Test_Set18(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		j := gjson.X创建(nil)

		t.Assert(j.X设置值("0.1.k1", "v1"), nil)
		t.Assert(j.X设置值("0.2.k2", "v2"), nil)
		s, err := j.X取json文本()
		t.AssertNil(err)
		t.Assert(s, `[[null,{"k1":"v1"},{"k2":"v2"}]]`)
	})
}

func Test_Set19(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		j := gjson.X创建(nil)

		t.Assert(j.X设置值("0.1.1.k1", "v1"), nil)
		t.Assert(j.X设置值("0.2.1.k2", "v2"), nil)
		s, err := j.X取json文本()
		t.AssertNil(err)
		t.Assert(s, `[[null,[null,{"k1":"v1"}],[null,{"k2":"v2"}]]]`)
	})
}

func Test_Set20(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		j := gjson.X创建(nil)

		t.Assert(j.X设置值("k1", "v1"), nil)
		t.Assert(j.X设置值("k2", g.Slice别名{1, 2, 3}), nil)
		t.Assert(j.X设置值("k2.1", 20), nil)
		t.Assert(j.X设置值("k2.2", g.Map{"k3": "v3"}), nil)
		s, err := j.X取json文本()
		t.AssertNil(err)
		t.Assert(gstr.X切片是否存在(
			g.SliceStr别名{
				`{"k1":"v1","k2":[1,20,{"k3":"v3"}]}`,
				`{"k2":[1,20,{"k3":"v3"}],"k1":"v1"}`,
			},
			s,
		), true)
	})
}

func Test_Set_GArray(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		j := gjson.X创建(nil)
		arr := garray.X创建().Append别名("test")
		t.AssertNil(j.X设置值("arr", arr))
		t.Assert(j.X取值("arr").Array别名(), g.Slice别名{"test"})
	})
}

func Test_Set_WithEmptyStruct(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		j := gjson.X创建(&struct{}{})
		t.AssertNil(j.X设置值("aa", "123"))
		t.Assert(j.X取json文本PANI(), `{"aa":"123"}`)
	})
}
