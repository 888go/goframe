// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package json类_test

import (
	"fmt"
	"testing"

	gmap "github.com/888go/goframe/container/gmap"
	gvar "github.com/888go/goframe/container/gvar"
	gjson "github.com/888go/goframe/encoding/gjson"
	"github.com/888go/goframe/frame/g"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
	gconv "github.com/888go/goframe/util/gconv"
)

func Test_New(t *testing.T) {
	// New with json map.
	gtest.C(t, func(t *gtest.T) {
		data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
		j := gjson.X创建(data)
		t.Assert(j.X取值("n").String(), "123456789")
		t.Assert(j.X取值("m").X取Map(), g.Map{"k": "v"})
		t.Assert(j.X取值("a").Array别名(), g.Slice别名{1, 2, 3})
	})
		// 使用json数组映射创建新的。 md5:0f642e2d9a82f660
	gtest.C(t, func(t *gtest.T) {
		j := gjson.X创建(`[{"a":1},{"b":2},{"c":3}]`)
		t.Assert(j.X取值(".").String(), `[{"a":1},{"b":2},{"c":3}]`)
		t.Assert(j.X取值("2.c").String(), `3`)
	})
	// 使用gvar的新建方法。
	// https:	//github.com/gogf/gf/issues/1571
	// md5:489c7f12ac3a473f
	gtest.C(t, func(t *gtest.T) {
		v := gvar.X创建(`[{"a":1},{"b":2},{"c":3}]`)
		j := gjson.X创建(v)
		t.Assert(j.X取值(".").String(), `[{"a":1},{"b":2},{"c":3}]`)
		t.Assert(j.X取值("2.c").String(), `3`)
	})
	// New with gmap.
	gtest.C(t, func(t *gtest.T) {
		m := gmap.X创建AnyAny并从Map(g.MapAnyAny{
			"k1": "v1",
			"k2": "v2",
		})
		j := gjson.X创建(m)
		t.Assert(j.X取值("k1"), "v1")
		t.Assert(j.X取值("k2"), "v2")
		t.Assert(j.X取值("k3"), nil)
	})
		// 这段注释是指向GitHub上一个名为gf的项目中的问题号3253的链接。在Go语言中，这种注释用于提供外部资源的引用或相关问题的追踪。 md5:f4927fbc7539374d
	gtest.C(t, func(t *gtest.T) {
		type TestStruct struct {
			Result []map[string]string `json:"result"`
		}
		ts := &TestStruct{
			Result: []map[string]string{
				{
					"Name": "gf",
					"Role": "",
				},
			},
		}
		gjson.X创建(ts)
	})
}

func Test_Valid(t *testing.T) {
	data1 := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	data2 := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]`)
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gjson.X是否为有效json(data1), true)
		t.Assert(gjson.X是否为有效json(data2), false)
	})
}

func Test_Encode(t *testing.T) {
	value := g.Slice别名{1, 2, 3}
	gtest.C(t, func(t *gtest.T) {
		b, err := gjson.X变量到json字节集(value)
		t.AssertNil(err)
		t.Assert(b, []byte(`[1,2,3]`))
	})
}

func Test_Decode(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	gtest.C(t, func(t *gtest.T) {
		v, err := gjson.Json格式到变量(data)
		t.AssertNil(err)
		t.Assert(v, g.Map{
			"n": 123456789,
			"a": g.Slice别名{1, 2, 3},
			"m": g.Map{
				"k": "v",
			},
		})
	})
	gtest.C(t, func(t *gtest.T) {
		var v interface{}
		err := gjson.Json格式到变量指针(data, &v)
		t.AssertNil(err)
		t.Assert(v, g.Map{
			"n": 123456789,
			"a": g.Slice别名{1, 2, 3},
			"m": g.Map{
				"k": "v",
			},
		})
	})
	gtest.C(t, func(t *gtest.T) {
		j, err := gjson.X解码到json(data)
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
	gtest.C(t, func(t *gtest.T) {
		j, err := gjson.X解码到json(data)
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
	gtest.C(t, func(t *gtest.T) {
		j, err := gjson.X解码到json(data)
		t.AssertNil(err)
		t.Assert(j.X取值("m.a.2"), 3)
		t.Assert(j.X取值("m.v1.v2"), nil)
		j.X设置分层冲突检查(true)
		t.Assert(j.X取值("m.v1.v2"), 4)
	})
}

func Test_GetVar(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	gtest.C(t, func(t *gtest.T) {
		j, err := gjson.X解码到json(data)
		t.AssertNil(err)
		t.Assert(j.X取值("n").String(), "123456789")
		t.Assert(j.X取值("m").X取Map(), g.Map{"k": "v"})
		t.Assert(j.X取值("a").X取any切片(), g.Slice别名{1, 2, 3})
		t.Assert(j.X取值("a").Slice别名(), g.Slice别名{1, 2, 3})
		t.Assert(j.X取值("a").Array别名(), g.Slice别名{1, 2, 3})
	})
}

func Test_GetMap(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	gtest.C(t, func(t *gtest.T) {
		j, err := gjson.X解码到json(data)
		t.AssertNil(err)
		t.Assert(j.X取值("n").X取Map(), nil)
		t.Assert(j.X取值("m").X取Map(), g.Map{"k": "v"})
		t.Assert(j.X取值("a").X取Map(), g.Map{"1": "2", "3": nil})
	})
}

func Test_GetJson(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	gtest.C(t, func(t *gtest.T) {
		j, err := gjson.X解码到json(data)
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
	gtest.C(t, func(t *gtest.T) {
		j, err := gjson.X解码到json(data)
		t.AssertNil(err)
		t.Assert(j.X取值("n").Array别名(), g.X切片{123456789})
		t.Assert(j.X取值("m").Array别名(), g.X切片{g.Map{"k": "v"}})
		t.Assert(j.X取值("a").Array别名(), g.X切片{1, 2, 3})
	})
}

func Test_GetString(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	gtest.C(t, func(t *gtest.T) {
		j, err := gjson.X解码到json(data)
		t.AssertNil(err)
		t.AssertEQ(j.X取值("n").String(), "123456789")
		t.AssertEQ(j.X取值("m").String(), `{"k":"v"}`)
		t.AssertEQ(j.X取值("a").String(), `[1,2,3]`)
		t.AssertEQ(j.X取值("i").String(), "")
	})
}

func Test_GetStrings(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	gtest.C(t, func(t *gtest.T) {
		j, err := gjson.X解码到json(data)
		t.AssertNil(err)
		t.AssertEQ(j.X取值("n").X取文本切片(), g.SliceStr别名{"123456789"})
		t.AssertEQ(j.X取值("m").X取文本切片(), g.SliceStr别名{`{"k":"v"}`})
		t.AssertEQ(j.X取值("a").X取文本切片(), g.SliceStr别名{"1", "2", "3"})
		t.AssertEQ(j.X取值("i").X取文本切片(), nil)
	})
}

func Test_GetInterfaces(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	gtest.C(t, func(t *gtest.T) {
		j, err := gjson.X解码到json(data)
		t.AssertNil(err)
		t.AssertEQ(j.X取值("n").X取any切片(), g.X切片{123456789})
		t.AssertEQ(j.X取值("m").X取any切片(), g.X切片{g.Map{"k": "v"}})
		t.AssertEQ(j.X取值("a").X取any切片(), g.X切片{1, 2, 3})
	})
}

func Test_Len(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		p := gjson.X创建(nil)
		p.X加入("a", 1)
		p.X加入("a", 2)
		t.Assert(p.X取长度("a"), 2)
	})
	gtest.C(t, func(t *gtest.T) {
		p := gjson.X创建(nil)
		p.X加入("a.b", 1)
		p.X加入("a.c", 2)
		t.Assert(p.X取长度("a"), 2)
	})
	gtest.C(t, func(t *gtest.T) {
		p := gjson.X创建(nil)
		p.X设置值("a", 1)
		t.Assert(p.X取长度("a"), -1)
	})
}

func Test_Append(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		p := gjson.X创建(nil)
		p.X加入("a", 1)
		p.X加入("a", 2)
		t.Assert(p.X取值("a"), g.Slice别名{1, 2})
	})
	gtest.C(t, func(t *gtest.T) {
		p := gjson.X创建(nil)
		p.X加入("a.b", 1)
		p.X加入("a.c", 2)
		t.Assert(p.X取值("a").X取Map(), g.Map{
			"b": g.Slice别名{1},
			"c": g.Slice别名{2},
		})
	})
	gtest.C(t, func(t *gtest.T) {
		p := gjson.X创建(nil)
		p.X设置值("a", 1)
		err := p.X加入("a", 2)
		t.AssertNE(err, nil)
		t.Assert(p.X取值("a"), 1)
	})
}

func Test_RawArray(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		j := gjson.X创建(nil)
		t.AssertNil(j.X设置值("0", 1))
		t.AssertNil(j.X设置值("1", 2))
		t.Assert(j.X取json文本PANI(), `[1,2]`)
	})

	gtest.C(t, func(t *gtest.T) {
		j := gjson.X创建(nil)
		t.AssertNil(j.X加入(".", 1))
		t.AssertNil(j.X加入(".", 2))
		t.Assert(j.X取json文本PANI(), `[1,2]`)
	})
}

func TestJson_ToJson(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		p := gjson.X创建(1)
		s, e := p.X取json文本()
		t.Assert(e, nil)
		t.Assert(s, "1")
	})
	gtest.C(t, func(t *gtest.T) {
		p := gjson.X创建("a")
		s, e := p.X取json文本()
		t.Assert(e, nil)
		t.Assert(s, `"a"`)
	})
}

func TestJson_Default(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		j := gjson.X创建(nil)
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
		t.AssertEQ(j.X取值("no", g.Slice别名{1, 2, 3}).X取整数切片(), g.SliceInt别名{1, 2, 3})
		t.AssertEQ(j.X取值("no", g.Slice别名{1, 2, 3}).X取小数切片(), []float64{1, 2, 3})
		t.AssertEQ(j.X取值("no", g.Map{"k": "v"}).X取Map(), g.Map{"k": "v"})
		t.AssertEQ(j.X取值("no", 123.456).X取小数64位(), float64(123.456))
		t.AssertEQ(j.X取对象("no", g.Map{"k": "v"}).X取值("k").String(), "v")
		t.AssertEQ(j.X取对象切片("no", g.Slice别名{
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
	gtest.C(t, func(t *gtest.T) {
		j := gjson.X创建(`{"name":"gf"}`)
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
	gtest.C(t, func(t *gtest.T) {
		name := struct {
			Name string
		}{}
		j := gjson.X创建(`{"name":"gf","time":"2019-06-12"}`)
		t.Assert(j.Interface().(g.Map)["name"], "gf")
		t.Assert(j.X取值("name1").X取Map(), nil)
		t.Assert(j.X取对象("name1"), nil)
		t.Assert(j.X取对象切片("name1"), nil)
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

		j = gjson.X创建(`{"person":{"name":"gf"}}`)
		err = j.X取值("person").X取结构体指针(&name)
		t.AssertNil(err)
		t.Assert(name.Name, "gf")

		j = gjson.X创建(`{"name":"gf""}`)
		// j.Dump()
		t.AssertNil(err)

		j = gjson.X创建(`[1,2,3]`)
		t.Assert(len(j.X取泛型类().Array别名()), 3)
	})
}

func Test_Basic(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		j := gjson.X创建(`{"name":"gf","time":"2019-06-12"}`)
		j.X设置分层冲突检查(true)
		t.Assert(j.X取值(""), nil)
		t.Assert(j.X取值(".").Interface().(g.Map)["name"], "gf")
		t.Assert(j.X取值(".").Interface().(g.Map)["name1"], nil)
		j.X设置分层冲突检查(false)
		t.Assert(j.X取值(".").Interface().(g.Map)["name"], "gf")

		err := j.X设置值("name", "gf1")
		t.AssertNil(err)
		t.Assert(j.X取值("name"), "gf1")

		j = gjson.X创建(`[1,2,3]`)
		err = j.X设置值("\"0\".1", 11)
		t.AssertNil(err)
		t.Assert(j.X取值("1"), 11)

		j = gjson.X创建(`[1,2,3]`)
		err = j.X设置值("11111111111111111111111", 11)
		t.AssertNE(err, nil)

		j = gjson.X创建(`[1,2,3]`)
		err = j.X删除("1")
		t.AssertNil(err)
		t.Assert(j.X取值("0"), 1)
		t.Assert(len(j.X取泛型类().Array别名()), 2)

		j = gjson.X创建(`[1,2,3]`)
				// 如果索引 0 被删除，其下一个项目将位于索引 0。 md5:935e43cb97250f0e
		t.Assert(j.X删除("0"), nil)
		t.Assert(j.X删除("0"), nil)
		t.Assert(j.X删除("0"), nil)
		t.Assert(j.X取值("0"), nil)
		t.Assert(len(j.X取泛型类().Array别名()), 0)

		j = gjson.X创建(`[1,2,3]`)
		err = j.X删除("3")
		t.AssertNil(err)
		t.Assert(j.X取值("0"), 1)
		t.Assert(len(j.X取泛型类().Array别名()), 3)

		j = gjson.X创建(`[1,2,3]`)
		err = j.X删除("0.3")
		t.AssertNil(err)
		t.Assert(j.X取值("0"), 1)

		j = gjson.X创建(`[1,2,3]`)
		err = j.X删除("0.a")
		t.AssertNil(err)
		t.Assert(j.X取值("0"), 1)

		name := struct {
			Name string
		}{Name: "gf"}
		j = gjson.X创建(name)
		t.Assert(j.X取值("Name"), "gf")
		err = j.X删除("Name")
		t.AssertNil(err)
		t.Assert(j.X取值("Name"), nil)

		err = j.X设置值("Name", "gf1")
		t.AssertNil(err)
		t.Assert(j.X取值("Name"), "gf1")

		j = gjson.X创建(nil)
		err = j.X删除("Name")
		t.AssertNil(err)
		t.Assert(j.X取值("Name"), nil)

		j = gjson.X创建(name)
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
		t.Assert(j.X取值("Name5").Interface().(g.X切片)[0], 1)

	})
}

func TestJson_Var(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		data := []byte("[9223372036854775807, 9223372036854775806]")
		array := gjson.X创建(data).X取泛型类().Array别名()
		t.Assert(array, []uint64{9223372036854776000, 9223372036854776000})
	})
	gtest.C(t, func(t *gtest.T) {
		data := []byte("[9223372036854775807, 9223372036854775806]")
		array := gjson.X创建并按选项(data, gjson.Options{StrNumber: true}).X取泛型类().Array别名()
		t.Assert(array, []uint64{9223372036854775807, 9223372036854775806})
	})
}

func TestJson_IsNil(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		j := gjson.X创建(nil)
		t.Assert(j.X是否为Nil(), true)
	})
}

func TestJson_Set_With_Struct(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v := gjson.X创建(g.Map{
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
	gtest.C(t, func(t *gtest.T) {
		type S struct {
			Id int64
		}
		s := S{
			Id: 53687091200,
		}
		m := make(map[string]interface{})
		t.AssertNil(gjson.Json格式到变量指针(gjson.X变量到json字节集PANI(s), &m, gjson.Options{
			StrNumber: false,
		}))
		t.Assert(fmt.Sprintf(`%v`, m["Id"]), `5.36870912e+10`)
		t.AssertNil(gjson.Json格式到变量指针(gjson.X变量到json字节集PANI(s), &m, gjson.Options{
			StrNumber: true,
		}))
		t.Assert(fmt.Sprintf(`%v`, m["Id"]), `53687091200`)
	})
}

// 这段注释引用的是一个GitHub问题的链接，来自gogf（一个Go语言的优秀库）项目。"gf"是"Golang Foundation"的缩写，它表示这是一个关于gogf库的问题编号为1617的讨论或报告。具体的内容需要查看链接以获取详细信息。 md5:b2aec94e8fc5f5be
func Test_Issue1617(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		type MyJsonName struct {
			F中文   int64 `json:"F中文"`
			F英文   int64 `json:"F英文"`
			F法文   int64 `json:"F法文"`
			F西班牙语 int64 `json:"F西班牙语"`
		}
		jso := `{"F中文":1,"F英文":2,"F法文":3,"F西班牙语":4}`
		var a MyJsonName
		json, err := gjson.X解码到json(jso)
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

//github.com/gogf/gf/issues/1747. md5:6ee5dc419dd3705e
func Test_Issue1747(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var j *gjson.Json
		err := gconv.Struct(gvar.X创建("[1, 2, 336371793314971759]"), &j)
		t.AssertNil(err)
		t.Assert(j.X取值("2"), `336371793314971759`)
	})
}

// 指向GoGF框架问题2520. md5:ed6150ec52dbee88
func Test_Issue2520(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		type test struct {
			Unique *gvar.Var `json:"unique"`
		}

		t2 := test{Unique: gvar.X创建(gtime.Date())}
		t.Assert(gjson.X变量到json文本PANI(t2), gjson.X创建(t2).X取json文本PANI())
	})
}
