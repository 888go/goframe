// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试当前目录下所有.go文件，并执行所有benchmark测试

package 文本类_test

import (
	"net/url"
	"testing"
	
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/888go/goframe/gstr"
)

func Test_Parse(t *testing.T) {
	// cover test
	gtest.C(t, func(t *gtest.T) {
		// empty
		m, err := 文本类.X参数解析("")
		t.AssertNil(err)
		t.Assert(m, nil)
		// invalid
		m, err = 文本类.X参数解析("a&b")
		t.AssertNil(err)
		t.Assert(m, make(map[string]interface{}))
		// special key
		m, err = 文本类.X参数解析(" =1& b=2&   c =3")
		t.AssertNil(err)
		t.Assert(m, map[string]interface{}{"b": "2", "c_": "3"})
		m, err = 文本类.X参数解析("c[=3")
		t.AssertNil(err)
		t.Assert(m, map[string]interface{}{"c_": "3"})
		m, err = 文本类.X参数解析("v[a][a]a=m")
		t.AssertNil(err)
		t.Assert(m, g.Map{
			"v": g.Map{
				"a": g.Map{
					"a": "m",
				},
			},
		})
		// 原始注释：v[][a]=m&v[][b]=b => map["v"]:[{"a":"m","b":"b"}]
// 翻译后的中文注释：
// 当表达式为 v[][a]=m 且 v[][b]=b 时，转换为映射格式为 map["v"]:[{"a":"m","b":"b"}]
// 其中，v 是一个二维数组或切片，"a" 和 "m"、"b" 和 "b" 分别表示键值对，
// 将这些键值对组合在内嵌的 JSON 对象中，并以数组形式存储在 map 的 "v" 键下。
		m, err = 文本类.X参数解析("v[][a]=m&v[][b]=b")
		t.AssertNil(err)
		t.Assert(m, g.Map{
			"v": g.Slice{
				g.Map{
					"a": "m",
					"b": "b",
				},
			},
		})
		// 原始注释：v[][a]=m&v[][a]=b => map["v"]:[{"a":"m"},{"a":"b"}]
// 中文注释：当表达式为 v[][a]=m 与 v[][a]=b 时，可以转换为映射形式，
// 即 map 中键为 "v" 的值是一个数组，数组包含两个结构体元素，每个结构体中字段 a 的值分别为 "m" 和 "b"：
// map["v"] := [{"a": "m"}, {"a": "b"}]
		m, err = 文本类.X参数解析("v[][a]=m&v[][a]=b")
		t.AssertNil(err)
		t.Assert(m, g.Map{
			"v": g.Slice{
				g.Map{
					"a": "m",
				},
				g.Map{
					"a": "b",
				},
			},
		})
		// error
		m, err = 文本类.X参数解析("v=111&v[]=m&v[]=a&v[]=b")
		t.Log(err)
		t.AssertNE(err, nil)
		m, err = 文本类.X参数解析("v=111&v[a]=m&v[a]=a")
		t.Log(err)
		t.AssertNE(err, nil)
		_, err = 文本类.X参数解析("%Q=%Q&b")
		t.Log(err)
		t.AssertNE(err, nil)
		_, err = 文本类.X参数解析("a=%Q&b")
		t.Log(err)
		t.AssertNE(err, nil)
		_, err = 文本类.X参数解析("v[a][a]=m&v[][a]=b")
		t.Log(err)
		t.AssertNE(err, nil)
	})

	// url
	gtest.C(t, func(t *gtest.T) {
		s := "goframe.org/index?name=john&score=100"
		u, err := url.Parse(s)
		t.AssertNil(err)
		m, err := 文本类.X参数解析(u.RawQuery)
		t.AssertNil(err)
		t.Assert(m["name"], "john")
		t.Assert(m["score"], "100")

		// name overwrite
		m, err = 文本类.X参数解析("a=1&a=2")
		t.AssertNil(err)
		t.Assert(m, g.Map{
			"a": 2,
		})
		// slice
		m, err = 文本类.X参数解析("a[]=1&a[]=2")
		t.AssertNil(err)
		t.Assert(m, g.Map{
			"a": g.Slice{"1", "2"},
		})
		// map
		m, err = 文本类.X参数解析("a=1&b=2&c=3")
		t.AssertNil(err)
		t.Assert(m, g.Map{
			"a": "1",
			"b": "2",
			"c": "3",
		})
		m, err = 文本类.X参数解析("a=1&a=2&c=3")
		t.AssertNil(err)
		t.Assert(m, g.Map{
			"a": "2",
			"c": "3",
		})
		// map
		m, err = 文本类.X参数解析("m[a]=1&m[b]=2&m[c]=3")
		t.AssertNil(err)
		t.Assert(m, g.Map{
			"m": g.Map{
				"a": "1",
				"b": "2",
				"c": "3",
			},
		})
		m, err = 文本类.X参数解析("m[a]=1&m[a]=2&m[b]=3")
		t.AssertNil(err)
		t.Assert(m, g.Map{
			"m": g.Map{
				"a": "2",
				"b": "3",
			},
		})
		// map - slice
		m, err = 文本类.X参数解析("m[a][]=1&m[a][]=2")
		t.AssertNil(err)
		t.Assert(m, g.Map{
			"m": g.Map{
				"a": g.Slice{"1", "2"},
			},
		})
		m, err = 文本类.X参数解析("m[a][b][]=1&m[a][b][]=2")
		t.AssertNil(err)
		t.Assert(m, g.Map{
			"m": g.Map{
				"a": g.Map{
					"b": g.Slice{"1", "2"},
				},
			},
		})
		// map - complicated
		m, err = 文本类.X参数解析("m[a1][b1][c1][d1]=1&m[a2][b2]=2&m[a3][b3][c3]=3")
		t.AssertNil(err)
		t.Assert(m, g.Map{
			"m": g.Map{
				"a1": g.Map{
					"b1": g.Map{
						"c1": g.Map{
							"d1": "1",
						},
					},
				},
				"a2": g.Map{
					"b2": "2",
				},
				"a3": g.Map{
					"b3": g.Map{
						"c3": "3",
					},
				},
			},
		})
	})
}
