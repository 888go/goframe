// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package json类_test

import (
	"fmt"
	"testing"
	
	"github.com/888go/goframe/container/gmap"
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/encoding/gjson"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gconv"
)

func Test_New(t *testing.T) {
	// 通过JSON映射创建新实例
	单元测试类.C(t, func(t *单元测试类.T) {
		data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
		j := json类.X创建(data)
		t.Assert(j.X取值("n").String(), "123456789")
		t.Assert(j.X取值("m").X取Map(), g.Map{"k": "v"})
		t.Assert(j.X取值("a").Array别名(), g.Slice别名{1, 2, 3})
	})
	// 使用json数组映射创建新的（对象）
	单元测试类.C(t, func(t *单元测试类.T) {
		j := json类.X创建(`[{"a":1},{"b":2},{"c":3}]`)
		t.Assert(j.X取值(".").String(), `[{"a":1},{"b":2},{"c":3}]`)
		t.Assert(j.X取值("2.c").String(), `3`)
	})
// 使用gvar新建。
// 参考文档：https://github.com/gogf/gf/issues/1571
	单元测试类.C(t, func(t *单元测试类.T) {
		v := 泛型类.X创建(`[{"a":1},{"b":2},{"c":3}]`)
		j := json类.X创建(v)
		t.Assert(j.X取值(".").String(), `[{"a":1},{"b":2},{"c":3}]`)
		t.Assert(j.X取值("2.c").String(), `3`)
	})
	// New with gmap.
	单元测试类.C(t, func(t *单元测试类.T) {
		m := map类.X创建AnyAny并从Map(g.MapAnyAny{
			"k1": "v1",
			"k2": "v2",
		})
		j := json类.X创建(m)
		t.Assert(j.X取值("k1"), "v1")
		t.Assert(j.X取值("k2"), "v2")
		t.Assert(j.X取值("k3"), nil)
	})
}

func Test_Valid(t *testing.T) {
	data1 := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	data2 := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]`)
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(json类.X是否为有效json(data1), true)
		t.Assert(json类.X是否为有效json(data2), false)
	})
}

func Test_Encode(t *testing.T) {
	value := g.Slice别名{1, 2, 3}
	单元测试类.C(t, func(t *单元测试类.T) {
		b, err := json类.X变量到json字节集(value)
		t.AssertNil(err)
		t.Assert(b, []byte(`[1,2,3]`))
	})
}

func Test_Decode(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	单元测试类.C(t, func(t *单元测试类.T) {
		v, err := json类.Json格式到变量(data)
		t.AssertNil(err)
		t.Assert(v, g.Map{
			"n": 123456789,
			"a": g.Slice别名{1, 2, 3},
			"m": g.Map{
				"k": "v",
			},
		})
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var v interface{}
		err := json类.Json格式到变量指针(data, &v)
		t.AssertNil(err)
		t.Assert(v, g.Map{
			"n": 123456789,
			"a": g.Slice别名{1, 2, 3},
			"m": g.Map{
				"k": "v",
			},
		})
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		j, err := json类.X解码到json(data)
		t.AssertNil(err)
		t.Assert(j.X取值("n").String(), "123456789")
		t.Assert(j.X取值("m").X取Map(), g.Map{"k": "v"})
		t.Assert(j.X取值("m.k"), "v")
		t.Assert(j.X取值("a").Array别名(), g.Slice别名{1, 2, 3})
		t.Assert(j.X取值("a.1").X取整数(), 2)
	})
}

func Test_SplitChar(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	单元测试类.C(t, func(t *单元测试类.T) {
		j, err := json类.X解码到json(data)
		j.X设置参数分隔符(byte('#'))
		t.AssertNil(err)
		t.Assert(j.X取值("n").String(), "123456789")
		t.Assert(j.X取值("m").X取Map(), g.Map{"k": "v"})
		t.Assert(j.X取值("m#k").String(), "v")
		t.Assert(j.X取值("a").Array别名(), g.Slice别名{1, 2, 3})
		t.Assert(j.X取值("a#1").X取整数(), 2)
	})
}

func Test_ViolenceCheck(t *testing.T) {
	data := []byte(`{"m":{"a":[1,2,3], "v1.v2":"4"}}`)
	单元测试类.C(t, func(t *单元测试类.T) {
		j, err := json类.X解码到json(data)
		t.AssertNil(err)
		t.Assert(j.X取值("m.a.2"), 3)
		t.Assert(j.X取值("m.v1.v2"), nil)
		j.X设置分层冲突检查(true)
		t.Assert(j.X取值("m.v1.v2"), 4)
	})
}

func Test_GetVar(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	单元测试类.C(t, func(t *单元测试类.T) {
		j, err := json类.X解码到json(data)
		t.AssertNil(err)
		t.Assert(j.X取值("n").String(), "123456789")
		t.Assert(j.X取值("m").X取Map(), g.Map{"k": "v"})
		t.Assert(j.X取值("a").X取any数组(), g.Slice别名{1, 2, 3})
		t.Assert(j.X取值("a").Slice别名(), g.Slice别名{1, 2, 3})
		t.Assert(j.X取值("a").Array别名(), g.Slice别名{1, 2, 3})
	})
}

func Test_GetMap(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	单元测试类.C(t, func(t *单元测试类.T) {
		j, err := json类.X解码到json(data)
		t.AssertNil(err)
		t.Assert(j.X取值("n").X取Map(), nil)
		t.Assert(j.X取值("m").X取Map(), g.Map{"k": "v"})
		t.Assert(j.X取值("a").X取Map(), g.Map{"1": "2", "3": nil})
	})
}

func Test_GetJson(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	单元测试类.C(t, func(t *单元测试类.T) {
		j, err := json类.X解码到json(data)
		t.AssertNil(err)
		j2 := j.X取对象("m")
		t.AssertNE(j2, nil)
		t.Assert(j2.X取值("k"), "v")
		t.Assert(j2.X取值("a"), nil)
		t.Assert(j2.X取值("n"), nil)
	})
}

func Test_GetArray(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	单元测试类.C(t, func(t *单元测试类.T) {
		j, err := json类.X解码到json(data)
		t.AssertNil(err)
		t.Assert(j.X取值("n").Array别名(), g.X数组{123456789})
		t.Assert(j.X取值("m").Array别名(), g.X数组{g.Map{"k": "v"}})
		t.Assert(j.X取值("a").Array别名(), g.X数组{1, 2, 3})
	})
}

func Test_GetString(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	单元测试类.C(t, func(t *单元测试类.T) {
		j, err := json类.X解码到json(data)
		t.AssertNil(err)
		t.AssertEQ(j.X取值("n").String(), "123456789")
		t.AssertEQ(j.X取值("m").String(), `{"k":"v"}`)
		t.AssertEQ(j.X取值("a").String(), `[1,2,3]`)
		t.AssertEQ(j.X取值("i").String(), "")
	})
}

func Test_GetStrings(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	单元测试类.C(t, func(t *单元测试类.T) {
		j, err := json类.X解码到json(data)
		t.AssertNil(err)
		t.AssertEQ(j.X取值("n").X取文本数组(), g.SliceStr别名{"123456789"})
		t.AssertEQ(j.X取值("m").X取文本数组(), g.SliceStr别名{`{"k":"v"}`})
		t.AssertEQ(j.X取值("a").X取文本数组(), g.SliceStr别名{"1", "2", "3"})
		t.AssertEQ(j.X取值("i").X取文本数组(), nil)
	})
}

func Test_GetInterfaces(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	单元测试类.C(t, func(t *单元测试类.T) {
		j, err := json类.X解码到json(data)
		t.AssertNil(err)
		t.AssertEQ(j.X取值("n").X取any数组(), g.X数组{123456789})
		t.AssertEQ(j.X取值("m").X取any数组(), g.X数组{g.Map{"k": "v"}})
		t.AssertEQ(j.X取值("a").X取any数组(), g.X数组{1, 2, 3})
	})
}

func Test_Len(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		p := json类.X创建(nil)
		p.X加入("a", 1)
		p.X加入("a", 2)
		t.Assert(p.X取长度("a"), 2)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		p := json类.X创建(nil)
		p.X加入("a.b", 1)
		p.X加入("a.c", 2)
		t.Assert(p.X取长度("a"), 2)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		p := json类.X创建(nil)
		p.X设置值("a", 1)
		t.Assert(p.X取长度("a"), -1)
	})
}

func Test_Append(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		p := json类.X创建(nil)
		p.X加入("a", 1)
		p.X加入("a", 2)
		t.Assert(p.X取值("a"), g.Slice别名{1, 2})
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		p := json类.X创建(nil)
		p.X加入("a.b", 1)
		p.X加入("a.c", 2)
		t.Assert(p.X取值("a").X取Map(), g.Map{
			"b": g.Slice别名{1},
			"c": g.Slice别名{2},
		})
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		p := json类.X创建(nil)
		p.X设置值("a", 1)
		err := p.X加入("a", 2)
		t.AssertNE(err, nil)
		t.Assert(p.X取值("a"), 1)
	})
}

func Test_RawArray(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		j := json类.X创建(nil)
		t.AssertNil(j.X设置值("0", 1))
		t.AssertNil(j.X设置值("1", 2))
		t.Assert(j.X取json文本PANI(), `[1,2]`)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		j := json类.X创建(nil)
		t.AssertNil(j.X加入(".", 1))
		t.AssertNil(j.X加入(".", 2))
		t.Assert(j.X取json文本PANI(), `[1,2]`)
	})
}

func TestJson_ToJson(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		p := json类.X创建(1)
		s, e := p.X取json文本()
		t.Assert(e, nil)
		t.Assert(s, "1")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		p := json类.X创建("a")
		s, e := p.X取json文本()
		t.Assert(e, nil)
		t.Assert(s, `"a"`)
	})
}

func TestJson_Default(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		j := json类.X创建(nil)
		t.AssertEQ(j.X取值("no", 100).X取整数(), 100)
		t.AssertEQ(j.X取值("no", 100).String(), "100")
		t.AssertEQ(j.X取值("no", "on").X取布尔(), true)
		t.AssertEQ(j.X取值("no", 100).X取整数(), 100)
		t.AssertEQ(j.X取值("no", 100).X取整数8位(), int8(100))
		t.AssertEQ(j.X取值("no", 100).X取整数16位(), int16(100))
		t.AssertEQ(j.X取值("no", 100).X取整数32位(), int32(100))
		t.AssertEQ(j.X取值("no", 100).X取整数64位(), int64(100))
		t.AssertEQ(j.X取值("no", 100).X取正整数(), uint(100))
		t.AssertEQ(j.X取值("no", 100).X取正整数8位(), uint8(100))
		t.AssertEQ(j.X取值("no", 100).X取正整数16位(), uint16(100))
		t.AssertEQ(j.X取值("no", 100).X取正整数32位(), uint32(100))
		t.AssertEQ(j.X取值("no", 100).X取正整数64位(), uint64(100))
		t.AssertEQ(j.X取值("no", 123.456).X取小数32位(), float32(123.456))
		t.AssertEQ(j.X取值("no", 123.456).X取小数64位(), float64(123.456))
		t.AssertEQ(j.X取值("no", g.Slice别名{1, 2, 3}).Array别名(), g.Slice别名{1, 2, 3})
		t.AssertEQ(j.X取值("no", g.Slice别名{1, 2, 3}).X取整数数组(), g.SliceInt别名{1, 2, 3})
		t.AssertEQ(j.X取值("no", g.Slice别名{1, 2, 3}).X取小数数组(), []float64{1, 2, 3})
		t.AssertEQ(j.X取值("no", g.Map{"k": "v"}).X取Map(), g.Map{"k": "v"})
		t.AssertEQ(j.X取值("no", 123.456).X取小数64位(), float64(123.456))
		t.AssertEQ(j.X取对象("no", g.Map{"k": "v"}).X取值("k").String(), "v")
		t.AssertEQ(j.X取对象数组("no", g.Slice别名{
			g.Map{"k1": "v1"},
			g.Map{"k2": "v2"},
			g.Map{"k3": "v3"},
		})[0].X取值("k1").String(), "v1")
		t.AssertEQ(j.X取对象Map("no", g.Map{
			"m1": g.Map{"k1": "v1"},
			"m2": g.Map{"k2": "v2"},
		})["m2"].X取值("k2").String(), "v2")
	})
}

func Test_Convert(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		j := json类.X创建(`{"name":"gf"}`)
		arr, err := j.X取xml字节集()
		t.AssertNil(err)
		t.Assert(string(arr), "<name>gf</name>")
		arr, err = j.X取xml字节集并格式化()
		t.AssertNil(err)
		t.Assert(string(arr), "<name>gf</name>")
		str, err := j.X取xml文本()
		t.AssertNil(err)
		t.Assert(str, "<name>gf</name>")
		str, err = j.X取xml文本并格式化()
		t.AssertNil(err)
		t.Assert(str, "<name>gf</name>")

		arr, err = j.X取json字节集并格式化()
		t.AssertNil(err)
		t.Assert(string(arr), "{\n\t\"name\": \"gf\"\n}")
		str, err = j.X取json文本并格式化()
		t.AssertNil(err)
		t.Assert(string(arr), "{\n\t\"name\": \"gf\"\n}")

		arr, err = j.X取YAML字节集()
		t.AssertNil(err)
		t.Assert(string(arr), "name: gf\n")
		str, err = j.X取YAML文本()
		t.AssertNil(err)
		t.Assert(string(arr), "name: gf\n")

		arr, err = j.X取TOML字节集()
		t.AssertNil(err)
		t.Assert(string(arr), "name = \"gf\"\n")
		str, err = j.X取TOML文本()
		t.AssertNil(err)
		t.Assert(string(arr), "name = \"gf\"\n")
	})
}

func Test_Convert2(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		name := struct {
			Name string
		}{}
		j := json类.X创建(`{"name":"gf","time":"2019-06-12"}`)
		t.Assert(j.Interface().(g.Map)["name"], "gf")
		t.Assert(j.X取值("name1").X取Map(), nil)
		t.Assert(j.X取对象("name1"), nil)
		t.Assert(j.X取对象数组("name1"), nil)
		t.Assert(j.X取对象Map("name1"), nil)
		t.Assert(j.X是否存在("name1"), false)
		t.Assert(j.X取值("name1").X是否为Nil(), true)
		t.Assert(j.X取值("name").X是否为Nil(), false)
		t.Assert(j.X取长度("name1"), -1)
		t.Assert(j.X取值("time").X取时间类().Format("2006-01-02"), "2019-06-12")
		t.Assert(j.X取值("time").X取gtime时间类().X取格式文本("Y-m-d"), "2019-06-12")
		t.Assert(j.X取值("time").X取时长().String(), "0s")

		err := j.X取泛型类().X取结构体指针(&name)
		t.AssertNil(err)
		t.Assert(name.Name, "gf")
		// j.Dump()
		t.AssertNil(err)

		j = json类.X创建(`{"person":{"name":"gf"}}`)
		err = j.X取值("person").X取结构体指针(&name)
		t.AssertNil(err)
		t.Assert(name.Name, "gf")

		j = json类.X创建(`{"name":"gf""}`)
		// j.Dump()
		t.AssertNil(err)

		j = json类.X创建(`[1,2,3]`)
		t.Assert(len(j.X取泛型类().Array别名()), 3)
	})
}

func Test_Basic(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		j := json类.X创建(`{"name":"gf","time":"2019-06-12"}`)
		j.X设置分层冲突检查(true)
		t.Assert(j.X取值(""), nil)
		t.Assert(j.X取值(".").Interface().(g.Map)["name"], "gf")
		t.Assert(j.X取值(".").Interface().(g.Map)["name1"], nil)
		j.X设置分层冲突检查(false)
		t.Assert(j.X取值(".").Interface().(g.Map)["name"], "gf")

		err := j.X设置值("name", "gf1")
		t.AssertNil(err)
		t.Assert(j.X取值("name"), "gf1")

		j = json类.X创建(`[1,2,3]`)
		err = j.X设置值("\"0\".1", 11)
		t.AssertNil(err)
		t.Assert(j.X取值("1"), 11)

		j = json类.X创建(`[1,2,3]`)
		err = j.X设置值("11111111111111111111111", 11)
		t.AssertNE(err, nil)

		j = json类.X创建(`[1,2,3]`)
		err = j.X删除("1")
		t.AssertNil(err)
		t.Assert(j.X取值("0"), 1)
		t.Assert(len(j.X取泛型类().Array别名()), 2)

		j = json类.X创建(`[1,2,3]`)
		// 如果索引0处的元素被删除，其下一个元素将会移动到索引0的位置。
		t.Assert(j.X删除("0"), nil)
		t.Assert(j.X删除("0"), nil)
		t.Assert(j.X删除("0"), nil)
		t.Assert(j.X取值("0"), nil)
		t.Assert(len(j.X取泛型类().Array别名()), 0)

		j = json类.X创建(`[1,2,3]`)
		err = j.X删除("3")
		t.AssertNil(err)
		t.Assert(j.X取值("0"), 1)
		t.Assert(len(j.X取泛型类().Array别名()), 3)

		j = json类.X创建(`[1,2,3]`)
		err = j.X删除("0.3")
		t.AssertNil(err)
		t.Assert(j.X取值("0"), 1)

		j = json类.X创建(`[1,2,3]`)
		err = j.X删除("0.a")
		t.AssertNil(err)
		t.Assert(j.X取值("0"), 1)

		name := struct {
			Name string
		}{Name: "gf"}
		j = json类.X创建(name)
		t.Assert(j.X取值("Name"), "gf")
		err = j.X删除("Name")
		t.AssertNil(err)
		t.Assert(j.X取值("Name"), nil)

		err = j.X设置值("Name", "gf1")
		t.AssertNil(err)
		t.Assert(j.X取值("Name"), "gf1")

		j = json类.X创建(nil)
		err = j.X删除("Name")
		t.AssertNil(err)
		t.Assert(j.X取值("Name"), nil)

		j = json类.X创建(name)
		t.Assert(j.X取值("Name"), "gf")
		err = j.X设置值("Name1", g.Map{"Name": "gf1"})
		t.AssertNil(err)
		t.Assert(j.X取值("Name1").Interface().(g.Map)["Name"], "gf1")
		err = j.X设置值("Name2", g.Slice别名{1, 2, 3})
		t.AssertNil(err)
		t.Assert(j.X取值("Name2").Interface().(g.Slice别名)[0], 1)
		err = j.X设置值("Name3", name)
		t.AssertNil(err)
		t.Assert(j.X取值("Name3").Interface().(g.Map)["Name"], "gf")
		err = j.X设置值("Name4", &name)
		t.AssertNil(err)
		t.Assert(j.X取值("Name4").Interface().(g.Map)["Name"], "gf")
		arr := [3]int{1, 2, 3}
		err = j.X设置值("Name5", arr)
		t.AssertNil(err)
		t.Assert(j.X取值("Name5").Interface().(g.X数组)[0], 1)

	})
}

func TestJson_Var(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		data := []byte("[9223372036854775807, 9223372036854775806]")
		array := json类.X创建(data).X取泛型类().Array别名()
		t.Assert(array, []uint64{9223372036854776000, 9223372036854776000})
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		data := []byte("[9223372036854775807, 9223372036854775806]")
		array := json类.X创建并按选项(data, json类.Options{StrNumber: true}).X取泛型类().Array别名()
		t.Assert(array, []uint64{9223372036854775807, 9223372036854775806})
	})
}

func TestJson_IsNil(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		j := json类.X创建(nil)
		t.Assert(j.X是否为Nil(), true)
	})
}

func TestJson_Set_With_Struct(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		v := json类.X创建(g.Map{
			"user1": g.Map{"name": "user1"},
			"user2": g.Map{"name": "user2"},
			"user3": g.Map{"name": "user3"},
		})
		user1 := v.X取对象("user1")
		t.AssertNil(user1.X设置值("id", 111))
		t.AssertNil(v.X设置值("user1", user1))
		t.Assert(v.X取值("user1.id"), 111)
	})
}

func TestJson_Options(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type S struct {
			Id int64
		}
		s := S{
			Id: 53687091200,
		}
		m := make(map[string]interface{})
		t.AssertNil(json类.Json格式到变量指针(json类.X变量到json字节集PANI(s), &m, json类.Options{
			StrNumber: false,
		}))
		t.Assert(fmt.Sprintf(`%v`, m["Id"]), `5.36870912e+10`)
		t.AssertNil(json类.Json格式到变量指针(json类.X变量到json字节集PANI(s), &m, json类.Options{
			StrNumber: true,
		}))
		t.Assert(fmt.Sprintf(`%v`, m["Id"]), `53687091200`)
	})
}

// 这是Go语言代码中的一行注释，引用了GitHub上gogf/gf仓库的一个Issue（问题）1617号。
// 翻译为：
// 参考GitHub上gogf/gf项目的问题1617
func Test_Issue1617(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type MyJsonName struct {
			F中文   int64 `json:"F中文"`
			F英文   int64 `json:"F英文"`
			F法文   int64 `json:"F法文"`
			F西班牙语 int64 `json:"F西班牙语"`
		}
		jso := `{"F中文":1,"F英文":2,"F法文":3,"F西班牙语":4}`
		var a MyJsonName
		json, err := json类.X解码到json(jso)
		t.AssertNil(err)
		err = json.X取结构体指针(&a)
		t.AssertNil(err)
		t.Assert(a, MyJsonName{
			F中文:   1,
			F英文:   2,
			F法文:   3,
			F西班牙语: 4,
		})
	})
}

// 这是golang代码中的一行注释，其内容为一个URL链接，指向了GitHub上gogf/gf项目的一个问题编号1747。
// 翻译为：
// 参考gogf/gf项目在GitHub上的第1747号问题：https://github.com/gogf/gf/issues/1747
func Test_Issue1747(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var j *json类.Json
		err := 转换类.Struct(泛型类.X创建("[1, 2, 336371793314971759]"), &j)
		t.AssertNil(err)
		t.Assert(j.X取值("2"), `336371793314971759`)
	})
}

// 这是Go语言代码中的一行注释，其内容引用了GitHub上gogf/gf仓库的第2520个issue。
// 中文翻译：
// 参考GitHub上gogf/gf项目下的第2520个问题。
func Test_Issue2520(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type test struct {
			Unique *泛型类.Var `json:"unique"`
		}

		t2 := test{Unique: 泛型类.X创建(时间类.Date())}
		t.Assert(json类.X变量到json文本PANI(t2), json类.X创建(t2).X取json文本PANI())
	})
}
