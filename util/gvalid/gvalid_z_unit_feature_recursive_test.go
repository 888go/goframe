// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 效验类_test

import (
	"testing"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/test/gtest"
)

func Test_CheckStruct_Recursive_Struct(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type Pass struct {
			Pass1 string `v:"required|same:Pass2"`
			Pass2 string `v:"required|same:Pass1"`
		}
		type User struct {
			Id   int
			Name string `v:"required"`
			Pass Pass
		}
		user := &User{
			Name: "",
			Pass: Pass{
				Pass1: "1",
				Pass2: "2",
			},
		}
		err := g.X效验类().Data(user).Run(ctx)
		t.AssertNE(err, nil)
		t.Assert(err.Maps()["Name"], g.Map{"required": "The Name field is required"})
		t.Assert(err.Maps()["Pass1"], g.Map{"same": "The Pass1 value `1` must be the same as field Pass2 value `2`"})
		t.Assert(err.Maps()["Pass2"], g.Map{"same": "The Pass2 value `2` must be the same as field Pass1 value `1`"})
	})
}

func Test_CheckStruct_Recursive_Struct_WithData(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type Pass struct {
			Pass1 string `v:"required|same:Pass2"`
			Pass2 string `v:"required|same:Pass1"`
		}
		type User struct {
			Id   int
			Name string `v:"required"`
			Pass Pass
		}
		user := &User{}
		data := g.Map{
			"Name": "john",
			"Pass": g.Map{
				"Pass1": 100,
				"Pass2": 200,
			},
		}
		err := g.X效验类().Data(user).Assoc(data).Run(ctx)
		t.AssertNE(err, nil)
		t.Assert(err.Maps()["Name"], nil)
		t.Assert(err.Maps()["Pass1"], g.Map{"same": "The Pass1 value `100` must be the same as field Pass2 value `200`"})
		t.Assert(err.Maps()["Pass2"], g.Map{"same": "The Pass2 value `200` must be the same as field Pass1 value `100`"})
	})
}

func Test_CheckStruct_Recursive_SliceStruct(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type Pass struct {
			Pass1 string `v:"required|same:Pass2"`
			Pass2 string `v:"required|same:Pass1"`
		}
		type User struct {
			Id     int
			Name   string `v:"required"`
			Passes []Pass
		}
		user := &User{
			Name: "",
			Passes: []Pass{
				{
					Pass1: "1",
					Pass2: "2",
				},
				{
					Pass1: "3",
					Pass2: "4",
				},
			},
		}
		err := g.X效验类().Data(user).Run(ctx)
		g.X调试输出(err.Items())
		t.AssertNE(err, nil)
		t.Assert(err.Maps()["Name"], g.Map{"required": "The Name field is required"})
		t.Assert(err.Maps()["Pass1"], g.Map{"same": "The Pass1 value `3` must be the same as field Pass2 value `4`"})
		t.Assert(err.Maps()["Pass2"], g.Map{"same": "The Pass2 value `4` must be the same as field Pass1 value `3`"})
	})
}

func Test_CheckStruct_Recursive_SliceStruct_Bail(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type Pass struct {
			Pass1 string `v:"required|same:Pass2"`
			Pass2 string `v:"required|same:Pass1"`
		}
		type User struct {
			Id     int
			Name   string `v:"required"`
			Passes []Pass
		}
		user := &User{
			Name: "",
			Passes: []Pass{
				{
					Pass1: "1",
					Pass2: "2",
				},
				{
					Pass1: "3",
					Pass2: "4",
				},
			},
		}
		err := g.X效验类().Bail().Data(user).Run(ctx)
		g.X调试输出(err.Items())
		t.AssertNE(err, nil)
		t.Assert(err.Maps()["Name"], nil)
		t.Assert(err.Maps()["Pass1"], g.Map{"same": "The Pass1 value `1` must be the same as field Pass2 value `2`"})
		t.Assert(err.Maps()["Pass2"], nil)
	})
}

func Test_CheckStruct_Recursive_SliceStruct_Required(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type Pass struct {
			Pass1 string `v:"required|same:Pass2"`
			Pass2 string `v:"required|same:Pass1"`
		}
		type User struct {
			Id     int
			Name   string `v:"required"`
			Passes []Pass
		}
		user := &User{}
		err := g.X效验类().Data(user).Run(ctx)
		g.X调试输出(err.Items())
		t.AssertNE(err, nil)
		t.Assert(err.Maps()["Name"], g.Map{"required": "The Name field is required"})
		t.Assert(err.Maps()["Pass1"], nil)
		t.Assert(err.Maps()["Pass2"], nil)
	})
}

func Test_CheckStruct_Recursive_MapStruct(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type Pass struct {
			Pass1 string `v:"required|same:Pass2"`
			Pass2 string `v:"required|same:Pass1"`
		}
		type User struct {
			Id     int
			Name   string `v:"required"`
			Passes map[string]Pass
		}
		user := &User{
			Name: "",
			Passes: map[string]Pass{
				"test1": {
					Pass1: "1",
					Pass2: "2",
				},
				"test2": {
					Pass1: "3",
					Pass2: "4",
				},
			},
		}
		err := g.X效验类().Data(user).Run(ctx)
		g.X调试输出(err.Items())
		t.AssertNE(err, nil)
		t.Assert(err.Maps()["Name"], g.Map{"required": "The Name field is required"})
		t.AssertNE(err.Maps()["Pass1"], nil)
		t.AssertNE(err.Maps()["Pass2"], nil)
	})
}

func Test_CheckMap_Recursive_SliceStruct(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type Pass struct {
			Pass1 string `v:"required|same:Pass2"`
			Pass2 string `v:"required|same:Pass1"`
		}
		user := g.Map{
			"Name": "",
			"Pass": []Pass{
				{
					Pass1: "1",
					Pass2: "2",
				},
				{
					Pass1: "3",
					Pass2: "4",
				},
			},
		}
		err := g.X效验类().Data(user).Run(ctx)
		g.X调试输出(err.Items())
		t.AssertNE(err, nil)
		t.Assert(err.Maps()["Name"], nil)
		t.Assert(err.Maps()["Pass1"], g.Map{"same": "The Pass1 value `3` must be the same as field Pass2 value `4`"})
		t.Assert(err.Maps()["Pass2"], g.Map{"same": "The Pass2 value `4` must be the same as field Pass1 value `3`"})
	})
}

func Test_CheckStruct_Recursively_SliceAttribute(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type Student struct {
			Name string `v:"required#Student Name is required"`
			Age  int    `v:"required"`
		}
		type Teacher struct {
			Name     string    `v:"required#Teacher Name is required"`
			Students []Student `v:"required"`
		}
		var (
			teacher = Teacher{}
			data    = g.Map{
				"name":     "john",
				"students": `[]`,
			}
		)
		err := g.X效验类().Assoc(data).Data(teacher).Run(ctx)
		t.Assert(err, `The Students field is required`)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		type Student struct {
			Name string `v:"required#Student Name is required"`
			Age  int    `v:"required"`
		}
		type Teacher struct {
			Name     string `v:"required#Teacher Name is required"`
			Students []Student
		}
		var (
			teacher = Teacher{}
			data    = g.Map{
				"name": "john",
			}
		)
		err := g.X效验类().Assoc(data).Data(teacher).Run(ctx)
		t.Assert(err, ``)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		type Student struct {
			Name string `v:"required#Student Name is required"`
			Age  int    `v:"required"`
		}
		type Teacher struct {
			Name     string    `v:"required#Teacher Name is required"`
			Students []Student `v:"required"`
		}
		var (
			teacher = Teacher{}
			data    = g.Map{
				"name":     "john",
				"students": `[{"age":2}, {"name":"jack", "age":4}]`,
			}
		)
		err := g.X效验类().Assoc(data).Data(teacher).Run(ctx)
		t.Assert(err, `Student Name is required`)
	})

	// 这是Go语言代码的GitHub issues链接，指向gogf/gf仓库下的第1864号问题。
	单元测试类.C(t, func(t *单元测试类.T) {
		type Student struct {
			Name string `v:"required"`
			Age  int
		}
		type Teacher struct {
			Name     string
			Students []*Student
		}
		var (
			teacher = Teacher{}
			data    = g.Map{
				"name":     "john",
				"students": `[{"age":2},{"name":"jack", "age":4}]`,
			}
		)
		err := g.X效验类().Assoc(data).Data(teacher).Run(ctx)
		t.Assert(err, `The Name field is required`)
	})
}

func Test_CheckStruct_Recursively_SliceAttribute_WithTypeAlias(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type ParamsItemBase struct {
			Component string `v:"required" dc:"组件名称"`
			Params    string `v:"required" dc:"配置参数(一般是JSON)"`
			Version   uint64 `v:"required" dc:"参数版本"`
		}
		type ParamsItem = ParamsItemBase
		type ParamsModifyReq struct {
			Revision  uint64       `v:"required"`
			BizParams []ParamsItem `v:"required"`
		}
		var (
			req  = ParamsModifyReq{}
			data = g.Map{
				"Revision":  "1",
				"BizParams": `[{}]`,
			}
		)
		err := g.X效验类().Assoc(data).Data(req).Run(ctx)
		t.Assert(err, `The Component field is required; The Params field is required; The Version field is required`)
	})
}

func Test_CheckStruct_Recursively_MapAttribute(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type Student struct {
			Name string `v:"required#Student Name is required"`
			Age  int    `v:"required"`
		}
		type Teacher struct {
			Name     string             `v:"required#Teacher Name is required"`
			Students map[string]Student `v:"required"`
		}
		var (
			teacher = Teacher{}
			data    = g.Map{
				"name":     "john",
				"students": `{"john":{"age":18}}`,
			}
		)
		err := g.X效验类().Assoc(data).Data(teacher).Run(ctx)
		t.Assert(err, `Student Name is required`)
	})
}

// 这是golang代码中的一行注释，其内容引用了GitHub上gf项目的第1983号issue。
// 中文翻译：
// 参考GitHub上gf项目中的第1983号问题。
func Test_Issue1983(t *testing.T) {
	// 错误：因为在Teacher结构体中的属性Student是一个已初始化的结构体，它具有默认值。
	单元测试类.C(t, func(t *单元测试类.T) {
		type Student struct {
			Name string `v:"required"`
			Age  int
		}
		type Teacher struct {
			Students Student
		}
		var (
			teacher = Teacher{}
			data    = g.Map{
				"students": nil,
			}
		)
		err := g.X效验类().Assoc(data).Data(teacher).Run(ctx)
		t.Assert(err, `The Name field is required`)
	})
	// 与upper相同，它不受关联值的影响。
	单元测试类.C(t, func(t *单元测试类.T) {
		type Student struct {
			Name string `v:"required"`
			Age  int
		}
		type Teacher struct {
			Students Student
		}
		var (
			teacher = Teacher{}
			data    = g.Map{
				"name":     "john",
				"students": nil,
			}
		)
		err := g.X效验类().Assoc(data).Data(teacher).Run(ctx)
		t.Assert(err, `The Name field is required`)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		type Student struct {
			Name string `v:"required"`
			Age  int
		}
		type Teacher struct {
			Students *Student
		}
		var (
			teacher = Teacher{}
			data    = g.Map{
				"students": nil,
			}
		)
		err := g.X效验类().Assoc(data).Data(teacher).Run(ctx)
		t.AssertNil(err)
	})
}

// 这是Go语言代码的URL注释，指向GitHub上gogf/gf仓库的第1921号问题。
func Test_Issue1921(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type SearchOption struct {
			Size int `v:"max:100"`
		}
		type SearchReq struct {
			Option *SearchOption `json:"option,omitempty"`
		}

		var (
			req = SearchReq{
				Option: &SearchOption{
					Size: 10000,
				},
			}
		)
		err := g.X效验类().Data(req).Run(ctx)
		t.Assert(err, "The Size value `10000` must be equal or lesser than 100")
	})
}

// 这是Go语言代码中的一行注释，其内容为一个GitHub网址链接，指向gogf/gf项目下的第2011号问题。 
// 翻译：// 参见GitHub上gogf/gf项目中的第2011号问题：https://github.com/gogf/gf/issues/2011
func Test_Issue2011(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type Student struct {
			Name string `v:"required|min-length:6"`
			Age  int
		}
		type Teacher struct {
			Student *Student
		}
		var (
			teacher = Teacher{}
			data    = g.Map{
				"student": g.Map{
					"name": "john",
				},
			}
		)
		err := g.X效验类().Assoc(data).Data(teacher).Run(ctx)
		t.Assert(err, "The Name value `john` length must be equal or greater than 6")
	})
}
