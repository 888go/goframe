// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package json类_test

import (
	"testing"

	gjson "github.com/888go/goframe/encoding/gjson"
	gtest "github.com/888go/goframe/test/gtest"
)

func Test_NewWithTag(t *testing.T) {
	type User struct {
		Age  int    `xml:"age-xml"  json:"age-json"`
		Name string `xml:"name-xml" json:"name-json"`
		Addr string `xml:"addr-xml" json:"addr-json"`
	}
	data := User{
		Age:  18,
		Name: "john",
		Addr: "chengdu",
	}
	// JSON
	gtest.C(t, func(t *gtest.T) {
		j := gjson.New(data)
		t.AssertNE(j, nil)
		t.Assert(j.Get("age-xml"), nil)
		t.Assert(j.Get("age-json"), data.Age)
		t.Assert(j.Get("name-xml"), nil)
		t.Assert(j.Get("name-json"), data.Name)
		t.Assert(j.Get("addr-xml"), nil)
		t.Assert(j.Get("addr-json"), data.Addr)
	})
	// XML
	gtest.C(t, func(t *gtest.T) {
		j := gjson.NewWithTag(data, "xml")
		t.AssertNE(j, nil)
		t.Assert(j.Get("age-xml"), data.Age)
		t.Assert(j.Get("age-json"), nil)
		t.Assert(j.Get("name-xml"), data.Name)
		t.Assert(j.Get("name-json"), nil)
		t.Assert(j.Get("addr-xml"), data.Addr)
		t.Assert(j.Get("addr-json"), nil)
	})
}

func Test_New_CustomStruct(t *testing.T) {
	type Base struct {
		Id int
	}
	type User struct {
		Base
		Name string
	}
	user := new(User)
	user.Id = 1
	user.Name = "john"

	gtest.C(t, func(t *gtest.T) {
		j := gjson.New(user)
		t.AssertNE(j, nil)

		s, err := j.ToJsonString()
		t.AssertNil(err)
		t.Assert(s == `{"Id":1,"Name":"john"}` || s == `{"Name":"john","Id":1}`, true)
	})
}

func Test_New_HierarchicalStruct(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		type Me struct {
			Name     string `json:"name"`
			Score    int    `json:"score"`
			Children []Me   `json:"children"`
		}
		me := Me{
			Name:  "john",
			Score: 100,
			Children: []Me{
				{
					Name:  "Bean",
					Score: 99,
				},
				{
					Name:  "Sam",
					Score: 98,
				},
			},
		}
		j := gjson.New(me)
		t.Assert(j.Remove("children.0.score"), nil)
		t.Assert(j.Remove("children.1.score"), nil)
		t.Assert(j.MustToJsonString(), `{"children":[{"children":null,"name":"Bean"},{"children":null,"name":"Sam"}],"name":"john","score":100}`)
	})
}

func Test_NewWithOptions(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		data := []byte("[9223372036854775807, 9223372036854775806]")
		array := gjson.New(data).Var().Array()
		t.Assert(array, []uint64{9223372036854776000, 9223372036854776000})
	})
	gtest.C(t, func(t *gtest.T) {
		data := []byte("[9223372036854775807, 9223372036854775806]")
		array := gjson.NewWithOptions(data, gjson.Options{StrNumber: true}).Var().Array()
		t.Assert(array, []uint64{9223372036854775807, 9223372036854775806})
	})
}

func Test_LoadContentType(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		data := []byte("value = 79937385836643329")
		j, err := gjson.LoadContentType("toml", data)
		t.AssertNil(err)
		value := j.Get("value").Int64()
		t.Assert(value, 79937385836643329)
	})
}
