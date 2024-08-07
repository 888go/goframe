// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 使用go test命令运行当前目录下所有.go文件的性能测试，模式为匹配所有函数. md5:b546d3aaffaebd06

package 文本类_test

import (
	"net/url"
	"testing"

	"github.com/888go/goframe/frame/g"
	gtest "github.com/888go/goframe/test/gtest"
	gstr "github.com/888go/goframe/text/gstr"
)

func Test_Parse(t *testing.T) {
	// cover test
	gtest.C(t, func(t *gtest.T) {
		// empty
		m, err := gstr.X参数解析("")
		t.AssertNil(err)
		t.Assert(m, nil)
		// invalid
		m, err = gstr.X参数解析("a&b")
		t.AssertNil(err)
		t.Assert(m, make(map[string]interface{}))
		// special key
		m, err = gstr.X参数解析(" =1& b=2&   c =3")
		t.AssertNil(err)
		t.Assert(m, map[string]interface{}{"b": "2", "c_": "3"})
		m, err = gstr.X参数解析("c[=3")
		t.AssertNil(err)
		t.Assert(m, map[string]interface{}{"c_": "3"})
		m, err = gstr.X参数解析("v[a][a]a=m")
		t.AssertNil(err)
		t.Assert(m, g.Map{
			"v": g.Map{
				"a": g.Map{
					"a": "m",
				},
			},
		})
				// v[][a]=m&v[][b]=b 翻译为：v 中的任意元素的 a 键值为 m，b 键值为 b => map["v"]：[{"a":"m","b":"b"}]. md5:4c02c4944234a046
		m, err = gstr.X参数解析("v[][a]=m&v[][b]=b")
		t.AssertNil(err)
		t.Assert(m, g.Map{
			"v": g.Slice别名{
				g.Map{
					"a": "m",
					"b": "b",
				},
			},
		})
				// 当 v 数组的某个元素的 "a" 键对应值为 m 和 b 时， => 结果映射为 "v": [{"a": "m"}, {"a": "b"}]。 md5:c5ab84cca5b2e02d
		m, err = gstr.X参数解析("v[][a]=m&v[][a]=b")
		t.AssertNil(err)
		t.Assert(m, g.Map{
			"v": g.Slice别名{
				g.Map{
					"a": "m",
				},
				g.Map{
					"a": "b",
				},
			},
		})
		// error
		m, err = gstr.X参数解析("v=111&v[]=m&v[]=a&v[]=b")
		t.Log(err)
		t.AssertNE(err, nil)
		m, err = gstr.X参数解析("v=111&v[a]=m&v[a]=a")
		t.Log(err)
		t.AssertNE(err, nil)
		_, err = gstr.X参数解析("%Q=%Q&b")
		t.Log(err)
		t.AssertNE(err, nil)
		_, err = gstr.X参数解析("a=%Q&b")
		t.Log(err)
		t.AssertNE(err, nil)
		_, err = gstr.X参数解析("v[a][a]=m&v[][a]=b")
		t.Log(err)
		t.AssertNE(err, nil)
	})

	// url
	gtest.C(t, func(t *gtest.T) {
		s := "goframe.org/index?name=john&score=100"
		u, err := url.Parse(s)
		t.AssertNil(err)
		m, err := gstr.X参数解析(u.RawQuery)
		t.AssertNil(err)
		t.Assert(m["name"], "john")
		t.Assert(m["score"], "100")

		// name overwrite
		m, err = gstr.X参数解析("a=1&a=2")
		t.AssertNil(err)
		t.Assert(m, g.Map{
			"a": 2,
		})
		// slice
		m, err = gstr.X参数解析("a[]=1&a[]=2")
		t.AssertNil(err)
		t.Assert(m, g.Map{
			"a": g.Slice别名{"1", "2"},
		})
		// map
		m, err = gstr.X参数解析("a=1&b=2&c=3")
		t.AssertNil(err)
		t.Assert(m, g.Map{
			"a": "1",
			"b": "2",
			"c": "3",
		})
		m, err = gstr.X参数解析("a=1&a=2&c=3")
		t.AssertNil(err)
		t.Assert(m, g.Map{
			"a": "2",
			"c": "3",
		})
		// map
		m, err = gstr.X参数解析("m[a]=1&m[b]=2&m[c]=3")
		t.AssertNil(err)
		t.Assert(m, g.Map{
			"m": g.Map{
				"a": "1",
				"b": "2",
				"c": "3",
			},
		})
		m, err = gstr.X参数解析("m[a]=1&m[a]=2&m[b]=3")
		t.AssertNil(err)
		t.Assert(m, g.Map{
			"m": g.Map{
				"a": "2",
				"b": "3",
			},
		})
		// map - slice
		m, err = gstr.X参数解析("m[a][]=1&m[a][]=2")
		t.AssertNil(err)
		t.Assert(m, g.Map{
			"m": g.Map{
				"a": g.Slice别名{"1", "2"},
			},
		})
		m, err = gstr.X参数解析("m[a][b][]=1&m[a][b][]=2")
		t.AssertNil(err)
		t.Assert(m, g.Map{
			"m": g.Map{
				"a": g.Map{
					"b": g.Slice别名{"1", "2"},
				},
			},
		})
		// map - complicated
		m, err = gstr.X参数解析("m[a1][b1][c1][d1]=1&m[a2][b2]=2&m[a3][b3][c3]=3")
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
