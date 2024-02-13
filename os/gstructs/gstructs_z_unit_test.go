// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gstructs_test

import (
	"testing"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/os/gstructs"
	"github.com/888go/goframe/test/gtest"
)

func Test_Basic(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Id   int
			Name string `params:"name"`
			Pass string `my-tag1:"pass1" my-tag2:"pass2" params:"pass"`
		}
		var user User
		m, _ := gstructs.TagMapName(user, []string{"params"})
		t.Assert(m, g.Map{"name": "Name", "pass": "Pass"})
		m, _ = gstructs.TagMapName(&user, []string{"params"})
		t.Assert(m, g.Map{"name": "Name", "pass": "Pass"})

		m, _ = gstructs.TagMapName(&user, []string{"params", "my-tag1"})
		t.Assert(m, g.Map{"name": "Name", "pass": "Pass"})
		m, _ = gstructs.TagMapName(&user, []string{"my-tag1", "params"})
		t.Assert(m, g.Map{"name": "Name", "pass1": "Pass"})
		m, _ = gstructs.TagMapName(&user, []string{"my-tag2", "params"})
		t.Assert(m, g.Map{"name": "Name", "pass2": "Pass"})
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		type Base struct {
			Pass1 string `params:"password1"`
			Pass2 string `params:"password2"`
		}
		type UserWithBase struct {
			Id   int
			Name string
			Base `params:"base"`
		}
		user := new(UserWithBase)
		m, _ := gstructs.TagMapName(user, []string{"params"})
		t.Assert(m, g.Map{
			"base":      "Base",
			"password1": "Pass1",
			"password2": "Pass2",
		})
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		type Base struct {
			Pass1 string `params:"password1"`
			Pass2 string `params:"password2"`
		}
		type UserWithEmbeddedAttribute struct {
			Id   int
			Name string
			Base
		}
		type UserWithoutEmbeddedAttribute struct {
			Id   int
			Name string
			Pass Base
		}
		user1 := new(UserWithEmbeddedAttribute)
		user2 := new(UserWithoutEmbeddedAttribute)
		m, _ := gstructs.TagMapName(user1, []string{"params"})
		t.Assert(m, g.Map{"password1": "Pass1", "password2": "Pass2"})
		m, _ = gstructs.TagMapName(user2, []string{"params"})
		t.Assert(m, g.Map{})
	})
}

func Test_StructOfNilPointer(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Id   int
			Name string `params:"name"`
			Pass string `my-tag1:"pass1" my-tag2:"pass2" params:"pass"`
		}
		var user *User
		m, _ := gstructs.TagMapName(user, []string{"params"})
		t.Assert(m, g.Map{"name": "Name", "pass": "Pass"})
		m, _ = gstructs.TagMapName(&user, []string{"params"})
		t.Assert(m, g.Map{"name": "Name", "pass": "Pass"})

		m, _ = gstructs.TagMapName(&user, []string{"params", "my-tag1"})
		t.Assert(m, g.Map{"name": "Name", "pass": "Pass"})
		m, _ = gstructs.TagMapName(&user, []string{"my-tag1", "params"})
		t.Assert(m, g.Map{"name": "Name", "pass1": "Pass"})
		m, _ = gstructs.TagMapName(&user, []string{"my-tag2", "params"})
		t.Assert(m, g.Map{"name": "Name", "pass2": "Pass"})
	})
}

func Test_Fields(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Id   int
			Name string `params:"name"`
			Pass string `my-tag1:"pass1" my-tag2:"pass2" params:"pass"`
		}
		var user *User
		fields, _ := gstructs.Fields(gstructs.FieldsInput{
			Pointer:         user,
			RecursiveOption: 0,
		})
		t.Assert(len(fields), 3)
		t.Assert(fields[0].Name(), "Id")
		t.Assert(fields[1].Name(), "Name")
		t.Assert(fields[1].Tag("params"), "name")
		t.Assert(fields[2].Name(), "Pass")
		t.Assert(fields[2].Tag("my-tag1"), "pass1")
		t.Assert(fields[2].Tag("my-tag2"), "pass2")
		t.Assert(fields[2].Tag("params"), "pass")
	})
}

func Test_Fields_WithEmbedded1(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type B struct {
			Name string
			Age  int
		}
		type A struct {
			Site  string
			B     // 应该放在这里以验证其索引。
			Score int64
		}
		r, err := gstructs.Fields(gstructs.FieldsInput{
			Pointer:         new(A),
			RecursiveOption: gstructs.RecursiveOptionEmbeddedNoTag,
		})
		t.AssertNil(err)
		t.Assert(len(r), 4)
		t.Assert(r[0].Name(), `Site`)
		t.Assert(r[1].Name(), `Name`)
		t.Assert(r[2].Name(), `Age`)
		t.Assert(r[3].Name(), `Score`)
	})
}

func Test_Fields_WithEmbedded2(t *testing.T) {
	type MetaNode struct {
		Id          uint   `orm:"id,primary"  description:""`
		Capacity    string `orm:"capacity"    description:"Capacity string"`
		Allocatable string `orm:"allocatable" description:"Allocatable string"`
		Status      string `orm:"status"      description:"Status string"`
	}
	type MetaNodeZone struct {
		Nodes    uint
		Clusters uint
		Disk     uint
		Cpu      uint
		Memory   uint
		Zone     string
	}

	type MetaNodeItem struct {
		MetaNode
		Capacity    []MetaNodeZone `dc:"Capacity []MetaNodeZone"`
		Allocatable []MetaNodeZone `dc:"Allocatable []MetaNodeZone"`
	}

	单元测试类.C(t, func(t *单元测试类.T) {
		r, err := gstructs.Fields(gstructs.FieldsInput{
			Pointer:         new(MetaNodeItem),
			RecursiveOption: gstructs.RecursiveOptionEmbeddedNoTag,
		})
		t.AssertNil(err)
		t.Assert(len(r), 4)
		t.Assert(r[0].Name(), `Id`)
		t.Assert(r[1].Name(), `Capacity`)
		t.Assert(r[1].TagStr(), `dc:"Capacity []MetaNodeZone"`)
		t.Assert(r[2].Name(), `Allocatable`)
		t.Assert(r[2].TagStr(), `dc:"Allocatable []MetaNodeZone"`)
		t.Assert(r[3].Name(), `Status`)
	})
}

// 当存在嵌套结构体时，过滤重复的字段。
func Test_Fields_WithEmbedded_Filter(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type B struct {
			Name string
			Age  int
		}
		type A struct {
			Name  string
			Site  string
			Age   string
			B     // 应该放在这里以验证其索引。
			Score int64
		}
		r, err := gstructs.Fields(gstructs.FieldsInput{
			Pointer:         new(A),
			RecursiveOption: gstructs.RecursiveOptionEmbeddedNoTag,
		})
		t.AssertNil(err)
		t.Assert(len(r), 4)
		t.Assert(r[0].Name(), `Name`)
		t.Assert(r[1].Name(), `Site`)
		t.Assert(r[2].Name(), `Age`)
		t.Assert(r[3].Name(), `Score`)
	})
}

func Test_FieldMap(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Id   int
			Name string `params:"name"`
			Pass string `my-tag1:"pass1" my-tag2:"pass2" params:"pass"`
		}
		var user *User
		m, _ := gstructs.FieldMap(gstructs.FieldMapInput{
			Pointer:          user,
			PriorityTagArray: []string{"params"},
			RecursiveOption:  gstructs.RecursiveOptionEmbedded,
		})
		t.Assert(len(m), 3)
		_, ok := m["Id"]
		t.Assert(ok, true)
		_, ok = m["Name"]
		t.Assert(ok, false)
		_, ok = m["name"]
		t.Assert(ok, true)
		_, ok = m["Pass"]
		t.Assert(ok, false)
		_, ok = m["pass"]
		t.Assert(ok, true)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Id   int
			Name string `params:"name"`
			Pass string `my-tag1:"pass1" my-tag2:"pass2" params:"pass"`
		}
		var user *User
		m, _ := gstructs.FieldMap(gstructs.FieldMapInput{
			Pointer:          user,
			PriorityTagArray: nil,
			RecursiveOption:  gstructs.RecursiveOptionEmbedded,
		})
		t.Assert(len(m), 3)
		_, ok := m["Id"]
		t.Assert(ok, true)
		_, ok = m["Name"]
		t.Assert(ok, true)
		_, ok = m["name"]
		t.Assert(ok, false)
		_, ok = m["Pass"]
		t.Assert(ok, true)
		_, ok = m["pass"]
		t.Assert(ok, false)
	})
}

func Test_StructType(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type B struct {
			Name string
		}
		type A struct {
			B
		}
		r, err := gstructs.StructType(new(A))
		t.AssertNil(err)
		t.Assert(r.Signature(), `github.com/888go/goframe/os/gstructs_test/gstructs_test.A`)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		type B struct {
			Name string
		}
		type A struct {
			B
		}
		r, err := gstructs.StructType(new(A).B)
		t.AssertNil(err)
		t.Assert(r.Signature(), `github.com/888go/goframe/os/gstructs_test/gstructs_test.B`)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		type B struct {
			Name string
		}
		type A struct {
			*B
		}
		r, err := gstructs.StructType(new(A).B)
		t.AssertNil(err)
		t.Assert(r.String(), `gstructs_test.B`)
	})
	// Error.
	单元测试类.C(t, func(t *单元测试类.T) {
		type B struct {
			Name string
		}
		type A struct {
			*B
			Id int
		}
		_, err := gstructs.StructType(new(A).Id)
		t.AssertNE(err, nil)
	})
}

func Test_StructTypeBySlice(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type B struct {
			Name string
		}
		type A struct {
			Array []*B
		}
		r, err := gstructs.StructType(new(A).Array)
		t.AssertNil(err)
		t.Assert(r.Signature(), `github.com/888go/goframe/os/gstructs_test/gstructs_test.B`)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		type B struct {
			Name string
		}
		type A struct {
			Array []B
		}
		r, err := gstructs.StructType(new(A).Array)
		t.AssertNil(err)
		t.Assert(r.Signature(), `github.com/888go/goframe/os/gstructs_test/gstructs_test.B`)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		type B struct {
			Name string
		}
		type A struct {
			Array *[]B
		}
		r, err := gstructs.StructType(new(A).Array)
		t.AssertNil(err)
		t.Assert(r.Signature(), `github.com/888go/goframe/os/gstructs_test/gstructs_test.B`)
	})
}

func TestType_FieldKeys(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type B struct {
			Id   int
			Name string
		}
		type A struct {
			Array []*B
		}
		r, err := gstructs.StructType(new(A).Array)
		t.AssertNil(err)
		t.Assert(r.FieldKeys(), g.Slice别名{"Id", "Name"})
	})
}

func TestType_TagMap(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type A struct {
			Id   int    `d:"123" description:"I love gf"`
			Name string `v:"required" description:"应用Id"`
		}
		r, err := gstructs.Fields(gstructs.FieldsInput{
			Pointer:         new(A),
			RecursiveOption: 0,
		})
		t.AssertNil(err)

		t.Assert(len(r), 2)
		t.Assert(r[0].TagMap()["d"], `123`)
		t.Assert(r[0].TagMap()["description"], `I love gf`)
		t.Assert(r[1].TagMap()["v"], `required`)
		t.Assert(r[1].TagMap()["description"], `应用Id`)
	})
}

func TestType_TagJsonName(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type A struct {
			Name string `json:"name,omitempty"`
		}
		r, err := gstructs.Fields(gstructs.FieldsInput{
			Pointer:         new(A),
			RecursiveOption: 0,
		})
		t.AssertNil(err)

		t.Assert(len(r), 1)
		t.Assert(r[0].TagJsonName(), `name`)
	})
}

func TestType_TagDefault(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type A struct {
			Name  string `default:"john"`
			Name2 string `d:"john"`
		}
		r, err := gstructs.Fields(gstructs.FieldsInput{
			Pointer:         new(A),
			RecursiveOption: 0,
		})
		t.AssertNil(err)

		t.Assert(len(r), 2)
		t.Assert(r[0].TagDefault(), `john`)
		t.Assert(r[1].TagDefault(), `john`)
	})
}

func TestType_TagParam(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type A struct {
			Name  string `param:"name"`
			Name2 string `p:"name"`
		}
		r, err := gstructs.Fields(gstructs.FieldsInput{
			Pointer:         new(A),
			RecursiveOption: 0,
		})
		t.AssertNil(err)

		t.Assert(len(r), 2)
		t.Assert(r[0].TagParam(), `name`)
		t.Assert(r[1].TagParam(), `name`)
	})
}

func TestType_TagValid(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type A struct {
			Name  string `valid:"required"`
			Name2 string `v:"required"`
		}
		r, err := gstructs.Fields(gstructs.FieldsInput{
			Pointer:         new(A),
			RecursiveOption: 0,
		})
		t.AssertNil(err)

		t.Assert(len(r), 2)
		t.Assert(r[0].TagValid(), `required`)
		t.Assert(r[1].TagValid(), `required`)
	})
}

func TestType_TagDescription(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type A struct {
			Name  string `description:"my name"`
			Name2 string `des:"my name"`
			Name3 string `dc:"my name"`
		}
		r, err := gstructs.Fields(gstructs.FieldsInput{
			Pointer:         new(A),
			RecursiveOption: 0,
		})
		t.AssertNil(err)

		t.Assert(len(r), 3)
		t.Assert(r[0].TagDescription(), `my name`)
		t.Assert(r[1].TagDescription(), `my name`)
		t.Assert(r[2].TagDescription(), `my name`)
	})
}

func TestType_TagSummary(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type A struct {
			Name  string `summary:"my name"`
			Name2 string `sum:"my name"`
			Name3 string `sm:"my name"`
		}
		r, err := gstructs.Fields(gstructs.FieldsInput{
			Pointer:         new(A),
			RecursiveOption: 0,
		})
		t.AssertNil(err)

		t.Assert(len(r), 3)
		t.Assert(r[0].TagSummary(), `my name`)
		t.Assert(r[1].TagSummary(), `my name`)
		t.Assert(r[2].TagSummary(), `my name`)
	})
}

func TestType_TagAdditional(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type A struct {
			Name  string `additional:"my name"`
			Name2 string `ad:"my name"`
		}
		r, err := gstructs.Fields(gstructs.FieldsInput{
			Pointer:         new(A),
			RecursiveOption: 0,
		})
		t.AssertNil(err)

		t.Assert(len(r), 2)
		t.Assert(r[0].TagAdditional(), `my name`)
		t.Assert(r[1].TagAdditional(), `my name`)
	})
}

func TestType_TagExample(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type A struct {
			Name  string `example:"john"`
			Name2 string `eg:"john"`
		}
		r, err := gstructs.Fields(gstructs.FieldsInput{
			Pointer:         new(A),
			RecursiveOption: 0,
		})
		t.AssertNil(err)

		t.Assert(len(r), 2)
		t.Assert(r[0].TagExample(), `john`)
		t.Assert(r[1].TagExample(), `john`)
	})
}
