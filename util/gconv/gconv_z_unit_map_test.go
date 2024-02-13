// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 转换类_test

import (
	"encoding/json"
	"testing"
	
	"gopkg.in/yaml.v3"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/gutil"
)

func Test_Map_Basic(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m1 := map[string]string{
			"k": "v",
		}
		m2 := map[int]string{
			3: "v",
		}
		m3 := map[float64]float32{
			1.22: 3.1,
		}
		t.Assert(转换类.X取Map(m1), g.Map{
			"k": "v",
		})
		t.Assert(转换类.X取Map(m2), g.Map{
			"3": "v",
		})
		t.Assert(转换类.X取Map(m3), g.Map{
			"1.22": "3.1",
		})
		t.Assert(转换类.X取Map(`{"name":"goframe"}`), g.Map{
			"name": "goframe",
		})
		t.Assert(转换类.X取Map(`{"name":"goframe"`), nil)
		t.Assert(转换类.X取Map(`{goframe}`), nil)
		t.Assert(转换类.X取Map([]byte(`{"name":"goframe"}`)), g.Map{
			"name": "goframe",
		})
		t.Assert(转换类.X取Map([]byte(`{"name":"goframe"`)), nil)
		t.Assert(转换类.X取Map([]byte(`{goframe}`)), nil)
	})
}

func Test_Map_Slice(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		slice1 := g.Slice别名{"1", "2", "3", "4"}
		slice2 := g.Slice别名{"1", "2", "3"}
		slice3 := g.Slice别名{}
		t.Assert(转换类.X取Map(slice1), g.Map{
			"1": "2",
			"3": "4",
		})
		t.Assert(转换类.X取Map(slice2), g.Map{
			"1": "2",
			"3": nil,
		})
		t.Assert(转换类.X取Map(slice3), g.Map{})
	})
}

func Test_Maps_Basic(t *testing.T) {
	params := g.Slice别名{
		g.Map{"id": 100, "name": "john"},
		g.Map{"id": 200, "name": "smith"},
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		list := 转换类.X取Map数组(params)
		t.Assert(len(list), 2)
		t.Assert(list[0]["id"], 100)
		t.Assert(list[1]["id"], 200)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		list := 转换类.SliceMap别名(params)
		t.Assert(len(list), 2)
		t.Assert(list[0]["id"], 100)
		t.Assert(list[1]["id"], 200)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		list := 转换类.SliceMapDeep别名(params)
		t.Assert(len(list), 2)
		t.Assert(list[0]["id"], 100)
		t.Assert(list[1]["id"], 200)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		type Base struct {
			Age int
		}
		type User struct {
			Id   int
			Name string
			Base
		}

		users := make([]User, 0)
		params := []g.Map{
			{"id": 1, "name": "john", "age": 18},
			{"id": 2, "name": "smith", "age": 20},
		}
		err := 转换类.SliceStruct别名(params, &users)
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].Id, params[0]["id"])
		t.Assert(users[0].Name, params[0]["name"])
		t.Assert(users[0].Age, 18)

		t.Assert(users[1].Id, params[1]["id"])
		t.Assert(users[1].Name, params[1]["name"])
		t.Assert(users[1].Age, 20)
	})
}

func Test_Maps_JsonStr(t *testing.T) {
	jsonStr := `[{"id":100, "name":"john"},{"id":200, "name":"smith"}]`
	单元测试类.C(t, func(t *单元测试类.T) {
		list := 转换类.X取Map数组(jsonStr)
		t.Assert(len(list), 2)
		t.Assert(list[0]["id"], 100)
		t.Assert(list[1]["id"], 200)

		list = 转换类.X取Map数组([]byte(jsonStr))
		t.Assert(len(list), 2)
		t.Assert(list[0]["id"], 100)
		t.Assert(list[1]["id"], 200)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(转换类.X取Map数组(`[id]`), nil)
		t.Assert(转换类.X取Map数组(`test`), nil)
		t.Assert(转换类.X取Map数组([]byte(`[id]`)), nil)
		t.Assert(转换类.X取Map数组([]byte(`test`)), nil)
	})
}

func Test_Map_StructWithGConvTag(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Uid      int
			Name     string
			SiteUrl  string `gconv:"-"`
			NickName string `gconv:"nickname, omitempty"`
			Pass1    string `gconv:"password1"`
			Pass2    string `gconv:"password2"`
		}
		user1 := User{
			Uid:     100,
			Name:    "john",
			SiteUrl: "https://goframe.org",
			Pass1:   "123",
			Pass2:   "456",
		}
		user2 := &user1
		map1 := 转换类.X取Map(user1)
		map2 := 转换类.X取Map(user2)
		t.Assert(map1["Uid"], 100)
		t.Assert(map1["Name"], "john")
		t.Assert(map1["SiteUrl"], nil)
		t.Assert(map1["NickName"], nil)
		t.Assert(map1["nickname"], nil)
		t.Assert(map1["password1"], "123")
		t.Assert(map1["password2"], "456")

		t.Assert(map2["Uid"], 100)
		t.Assert(map2["Name"], "john")
		t.Assert(map2["SiteUrl"], nil)
		t.Assert(map2["NickName"], nil)
		t.Assert(map2["nickname"], nil)
		t.Assert(map2["password1"], "123")
		t.Assert(map2["password2"], "456")
	})
}

func Test_Map_StructWithJsonTag(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Uid      int
			Name     string
			SiteUrl  string `json:"-"`
			NickName string `json:"nickname, omitempty"`
			Pass1    string `json:"password1"`
			Pass2    string `json:"password2"`
		}
		user1 := User{
			Uid:     100,
			Name:    "john",
			SiteUrl: "https://goframe.org",
			Pass1:   "123",
			Pass2:   "456",
		}
		user2 := &user1
		map1 := 转换类.X取Map(user1)
		map2 := 转换类.X取Map(user2)
		t.Assert(map1["Uid"], 100)
		t.Assert(map1["Name"], "john")
		t.Assert(map1["SiteUrl"], nil)
		t.Assert(map1["NickName"], nil)
		t.Assert(map1["nickname"], nil)
		t.Assert(map1["password1"], "123")
		t.Assert(map1["password2"], "456")

		t.Assert(map2["Uid"], 100)
		t.Assert(map2["Name"], "john")
		t.Assert(map2["SiteUrl"], nil)
		t.Assert(map2["NickName"], nil)
		t.Assert(map2["nickname"], nil)
		t.Assert(map2["password1"], "123")
		t.Assert(map2["password2"], "456")
	})
}

func Test_Map_StructWithCTag(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Uid      int
			Name     string
			SiteUrl  string `c:"-"`
			NickName string `c:"nickname, omitempty"`
			Pass1    string `c:"password1"`
			Pass2    string `c:"password2"`
		}
		user1 := User{
			Uid:     100,
			Name:    "john",
			SiteUrl: "https://goframe.org",
			Pass1:   "123",
			Pass2:   "456",
		}
		user2 := &user1
		map1 := 转换类.X取Map(user1)
		map2 := 转换类.X取Map(user2)
		t.Assert(map1["Uid"], 100)
		t.Assert(map1["Name"], "john")
		t.Assert(map1["SiteUrl"], nil)
		t.Assert(map1["NickName"], nil)
		t.Assert(map1["nickname"], nil)
		t.Assert(map1["password1"], "123")
		t.Assert(map1["password2"], "456")

		t.Assert(map2["Uid"], 100)
		t.Assert(map2["Name"], "john")
		t.Assert(map2["SiteUrl"], nil)
		t.Assert(map2["NickName"], nil)
		t.Assert(map2["nickname"], nil)
		t.Assert(map2["password1"], "123")
		t.Assert(map2["password2"], "456")
	})
}

func Test_Map_PrivateAttribute(t *testing.T) {
	type User struct {
		Id   int
		name string
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		user := &User{1, "john"}
		t.Assert(转换类.X取Map(user), g.Map{"Id": 1})
	})
}

func Test_Map_Embedded(t *testing.T) {
	type Base struct {
		Id int
	}
	type User struct {
		Base
		Name string
	}
	type UserDetail struct {
		User
		Brief string
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		user := &User{}
		user.Id = 1
		user.Name = "john"

		m := 转换类.X取Map(user)
		t.Assert(len(m), 2)
		t.Assert(m["Id"], user.Id)
		t.Assert(m["Name"], user.Name)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		user := &UserDetail{}
		user.Id = 1
		user.Name = "john"
		user.Brief = "john guo"

		m := 转换类.X取Map(user)
		t.Assert(len(m), 3)
		t.Assert(m["Id"], user.Id)
		t.Assert(m["Name"], user.Name)
		t.Assert(m["Brief"], user.Brief)
	})
}

func Test_Map_Embedded2(t *testing.T) {
	type Ids struct {
		Id  int `c:"id"`
		Uid int `c:"uid"`
	}
	type Base struct {
		Ids
		CreateTime string `c:"create_time"`
	}
	type User struct {
		Base
		Passport string `c:"passport"`
		Password string `c:"password"`
		Nickname string `c:"nickname"`
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		user := new(User)
		user.Id = 100
		user.Nickname = "john"
		user.CreateTime = "2019"
		m := 转换类.X取Map(user)
		t.Assert(m["id"], "100")
		t.Assert(m["nickname"], user.Nickname)
		t.Assert(m["create_time"], "2019")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		user := new(User)
		user.Id = 100
		user.Nickname = "john"
		user.CreateTime = "2019"
		m := 转换类.X取Map_递归(user)
		t.Assert(m["id"], user.Id)
		t.Assert(m["nickname"], user.Nickname)
		t.Assert(m["create_time"], user.CreateTime)
	})
}

func Test_MapDeep2(t *testing.T) {
	type A struct {
		F string
		G string
	}

	type B struct {
		A
		H string
	}

	type C struct {
		A A
		F string
	}

	type D struct {
		I A
		F string
	}

	单元测试类.C(t, func(t *单元测试类.T) {
		b := new(B)
		c := new(C)
		d := new(D)
		mb := 转换类.X取Map_递归(b)
		mc := 转换类.X取Map_递归(c)
		md := 转换类.X取Map_递归(d)
		t.Assert(工具类.MapContains(mb, "F"), true)
		t.Assert(工具类.MapContains(mb, "G"), true)
		t.Assert(工具类.MapContains(mb, "H"), true)
		t.Assert(工具类.MapContains(mc, "A"), true)
		t.Assert(工具类.MapContains(mc, "F"), true)
		t.Assert(工具类.MapContains(mc, "G"), false)
		t.Assert(工具类.MapContains(md, "F"), true)
		t.Assert(工具类.MapContains(md, "I"), true)
		t.Assert(工具类.MapContains(md, "H"), false)
		t.Assert(工具类.MapContains(md, "G"), false)
	})
}

func Test_MapDeep3(t *testing.T) {
	type Base struct {
		Id   int    `c:"id"`
		Date string `c:"date"`
	}
	type User struct {
		UserBase Base   `c:"base"`
		Passport string `c:"passport"`
		Password string `c:"password"`
		Nickname string `c:"nickname"`
	}

	单元测试类.C(t, func(t *单元测试类.T) {
		user := &User{
			UserBase: Base{
				Id:   1,
				Date: "2019-10-01",
			},
			Passport: "john",
			Password: "123456",
			Nickname: "JohnGuo",
		}
		m := 转换类.X取Map_递归(user)
		t.Assert(m, g.Map{
			"base": g.Map{
				"id":   user.UserBase.Id,
				"date": user.UserBase.Date,
			},
			"passport": user.Passport,
			"password": user.Password,
			"nickname": user.Nickname,
		})
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		user := &User{
			UserBase: Base{
				Id:   1,
				Date: "2019-10-01",
			},
			Passport: "john",
			Password: "123456",
			Nickname: "JohnGuo",
		}
		m := 转换类.X取Map(user)
		t.Assert(m, g.Map{
			"base":     user.UserBase,
			"passport": user.Passport,
			"password": user.Password,
			"nickname": user.Nickname,
		})
	})
}

func Test_MapDeepWithAttributeTag(t *testing.T) {
	type Ids struct {
		Id  int `c:"id"`
		Uid int `c:"uid"`
	}
	type Base struct {
		Ids        `json:"ids"`
		CreateTime string `c:"create_time"`
	}
	type User struct {
		Base     `json:"base"`
		Passport string `c:"passport"`
		Password string `c:"password"`
		Nickname string `c:"nickname"`
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		user := new(User)
		user.Id = 100
		user.Nickname = "john"
		user.CreateTime = "2019"
		m := 转换类.X取Map(user)
		t.Assert(m["id"], "")
		t.Assert(m["nickname"], user.Nickname)
		t.Assert(m["create_time"], "")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		user := new(User)
		user.Id = 100
		user.Nickname = "john"
		user.CreateTime = "2019"
		m := 转换类.X取Map_递归(user)
		t.Assert(m["base"].(map[string]interface{})["ids"].(map[string]interface{})["id"], user.Id)
		t.Assert(m["nickname"], user.Nickname)
		t.Assert(m["base"].(map[string]interface{})["create_time"], user.CreateTime)
	})
}

func Test_MapDeepWithNestedMapAnyAny(t *testing.T) {
	type User struct {
		ExtraAttributes g.Map `c:"extra_attributes"`
	}

	单元测试类.C(t, func(t *单元测试类.T) {
		user := &User{
			ExtraAttributes: g.Map{
				"simple_attribute": 123,
				"map_string_attribute": g.Map{
					"inner_value": 456,
				},
				"map_interface_attribute": g.MapAnyAny{
					"inner_value": 456,
					123:           "integer_key_should_be_converted_to_string",
				},
			},
		}
		m := 转换类.X取Map_递归(user)
		t.Assert(m, g.Map{
			"extra_attributes": g.Map{
				"simple_attribute": 123,
				"map_string_attribute": g.Map{
					"inner_value": user.ExtraAttributes["map_string_attribute"].(g.Map)["inner_value"],
				},
				"map_interface_attribute": g.Map{
					"inner_value": user.ExtraAttributes["map_interface_attribute"].(g.MapAnyAny)["inner_value"],
					"123":         "integer_key_should_be_converted_to_string",
				},
			},
		})
	})

	type Outer struct {
		OuterStruct map[string]interface{} `c:"outer_struct" yaml:"outer_struct"`
		Field3      map[string]interface{} `c:"field3" yaml:"field3"`
	}

	单元测试类.C(t, func(t *单元测试类.T) {
		problemYaml := []byte(`
outer_struct:
  field1: &anchor1
    inner1: 123
    inner2: 345
  field2: 
    inner3: 456
    inner4: 789
    <<: *anchor1
field3:
  123: integer_key
`)
		parsed := &Outer{}

		err := yaml.Unmarshal(problemYaml, parsed)
		t.AssertNil(err)

		_, err = json.Marshal(parsed)
		t.AssertNil(err)

		converted := 转换类.X取Map_递归(parsed)
		jsonData, err := json.Marshal(converted)
		t.AssertNil(err)

		t.Assert(string(jsonData), `{"field3":{"123":"integer_key"},"outer_struct":{"field1":{"inner1":123,"inner2":345},"field2":{"inner1":123,"inner2":345,"inner3":456,"inner4":789}}}`)
	})
}

func TestMapStrStr(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(转换类.X取文本Map(map[string]string{"k": "v"}), map[string]string{"k": "v"})
		t.Assert(转换类.X取文本Map(`{}`), nil)
	})
}

func TestMapStrStrDeep(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(转换类.X取文本Map_递归(map[string]string{"k": "v"}), map[string]string{"k": "v"})
		t.Assert(转换类.X取文本Map_递归(`{"k":"v"}`), map[string]string{"k": "v"})
		t.Assert(转换类.X取文本Map_递归(`{}`), nil)
	})
}

func TestMapsDeep(t *testing.T) {
	jsonStr := `[{"id":100, "name":"john"},{"id":200, "name":"smith"}]`
	params := g.Slice别名{
		g.Map{"id": 100, "name": "john"},
		g.Map{"id": 200, "name": "smith"},
	}

	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(转换类.X取Map数组_递归(nil), nil)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		list := 转换类.X取Map数组_递归(params)
		t.Assert(len(list), 2)
		t.Assert(list[0]["id"], 100)
		t.Assert(list[1]["id"], 200)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		list := 转换类.X取Map数组_递归(jsonStr)
		t.Assert(len(list), 2)
		t.Assert(list[0]["id"], 100)
		t.Assert(list[1]["id"], 200)

		list = 转换类.X取Map数组_递归([]byte(jsonStr))
		t.Assert(len(list), 2)
		t.Assert(list[0]["id"], 100)
		t.Assert(list[1]["id"], 200)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(转换类.X取Map数组_递归(`[id]`), nil)
		t.Assert(转换类.X取Map数组_递归(`test`), nil)
		t.Assert(转换类.X取Map数组_递归([]byte(`[id]`)), nil)
		t.Assert(转换类.X取Map数组_递归([]byte(`test`)), nil)
		t.Assert(转换类.X取Map数组_递归([]string{}), nil)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		stringInterfaceMapList := make([]map[string]interface{}, 0)
		stringInterfaceMapList = append(stringInterfaceMapList, map[string]interface{}{"id": 100})
		stringInterfaceMapList = append(stringInterfaceMapList, map[string]interface{}{"id": 200})
		list := 转换类.X取Map数组_递归(stringInterfaceMapList)
		t.Assert(len(list), 2)
		t.Assert(list[0]["id"], 100)
		t.Assert(list[1]["id"], 200)

		list = 转换类.X取Map数组_递归([]byte(jsonStr))
		t.Assert(len(list), 2)
		t.Assert(list[0]["id"], 100)
		t.Assert(list[1]["id"], 200)
	})
}

func TestMapWithJsonOmitEmpty(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type S struct {
			Key   string      `json:",omitempty"`
			Value interface{} `json:",omitempty"`
		}
		s := S{
			Key:   "",
			Value: 1,
		}
		m1 := 转换类.X取Map(s)
		t.Assert(m1, g.Map{
			"Key":   "",
			"Value": 1,
		})

		m2 := 转换类.X取Map(s, 转换类.MapOption{
			Deep:      false,
			OmitEmpty: true,
			Tags:      nil,
		})
		t.Assert(m2, g.Map{
			"Value": 1,
		})
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		type ProductConfig struct {
			Pid      int `v:"required" json:"pid,omitempty"`
			TimeSpan int `v:"required" json:"timeSpan,omitempty"`
		}
		type CreateGoodsDetail struct {
			ProductConfig
			AutoRenewFlag int `v:"required" json:"autoRenewFlag"`
		}
		s := &CreateGoodsDetail{
			ProductConfig: ProductConfig{
				Pid:      1,
				TimeSpan: 0,
			},
			AutoRenewFlag: 0,
		}
		m1 := 转换类.X取Map(s)
		t.Assert(m1, g.Map{
			"pid":           1,
			"timeSpan":      0,
			"autoRenewFlag": 0,
		})

		m2 := 转换类.X取Map(s, 转换类.MapOption{
			Deep:      false,
			OmitEmpty: true,
			Tags:      nil,
		})
		t.Assert(m2, g.Map{
			"pid":           1,
			"autoRenewFlag": 0,
		})
	})
}
