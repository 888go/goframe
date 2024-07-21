// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gconv_test

import (
	"math/big"
	"testing"
	"time"

	"github.com/gogf/gf/v2/container/gtype"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gconv"
)

// https://github.com/gogf/gf/issues/1227
// 
// 这段注释是引用了GitHub上gf项目的一个问题链接，编号为1227。在Go语言中，这种注释用于提供外部资源的参考或问题追踪。翻译成中文后，它仍然保留原始的URL，因为这是一个链接，无需翻译。 md5:b76d46c66d00a5ce
func Test_Issue1227(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		type StructFromIssue1227 struct {
			Name string `json:"n1"`
		}
		tests := []struct {
			name   string
			origin interface{}
			want   string
		}{
			{
				name:   "Case1",
				origin: `{"n1":"n1"}`,
				want:   "n1",
			},
			{
				name:   "Case2",
				origin: `{"name":"name"}`,
				want:   "",
			},
			{
				name:   "Case3",
				origin: `{"NaMe":"NaMe"}`,
				want:   "",
			},
			{
				name:   "Case4",
				origin: g.Map{"n1": "n1"},
				want:   "n1",
			},
			{
				name:   "Case5",
				origin: g.Map{"NaMe": "n1"},
				want:   "n1",
			},
		}
		for _, tt := range tests {
			p := StructFromIssue1227{}
			if err := gconv.Struct(tt.origin, &p); err != nil {
				t.Error(err)
			}
			t.Assert(p.Name, tt.want)
		}
	})

	// Chinese key.
	gtest.C(t, func(t *gtest.T) {
		type StructFromIssue1227 struct {
			Name string `json:"中文Key"`
		}
		tests := []struct {
			name   string
			origin interface{}
			want   string
		}{
			{
				name:   "Case1",
				origin: `{"中文Key":"n1"}`,
				want:   "n1",
			},
			{
				name:   "Case2",
				origin: `{"Key":"name"}`,
				want:   "",
			},
			{
				name:   "Case3",
				origin: `{"NaMe":"NaMe"}`,
				want:   "",
			},
			{
				name:   "Case4",
				origin: g.Map{"中文Key": "n1"},
				want:   "n1",
			},
			{
				name:   "Case5",
				origin: g.Map{"中文KEY": "n1"},
				want:   "",
			},
			{
				name:   "Case5",
				origin: g.Map{"KEY": "n1"},
				want:   "",
			},
		}
		for _, tt := range tests {
			p := StructFromIssue1227{}
			if err := gconv.Struct(tt.origin, &p); err != nil {
				t.Error(err)
			}
			//t.Log(tt)
			t.Assert(p.Name, tt.want)
		}
	})
}

// https://github.com/gogf/gf/issues/1607
// 
// 这段注释引用的是一个GitHub问题（Issue）的链接，来自gogf（一个Go语言的开发框架）项目。它可能表示开发者在讨论或记录与该框架相关的一个特定问题（编号为1607）。 md5:a5485f6c5d31a8b4
func Test_Issue1607(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		type Demo struct {
			B Float64
		}
		rat := &big.Rat{}
		rat.SetFloat64(1.5)

		var demos = make([]Demo, 1)
		err := gconv.Scan([]map[string]interface{}{
			{"A": 1, "B": rat},
		}, &demos)
		t.AssertNil(err)
		t.Assert(demos[0].B, 1.5)
	})
}

// https://github.com/gogf/gf/issues/1946
// 
// 这段注释引用的是一个GitHub问题，链接为：https://github.com/gogf/gf/issues/1946。GF（Golang Foundation）是一个与Golang（Go语言）相关的开源项目，而"issues/1946"表示该链接指向的是第1946个已知问题或讨论。 md5:daac20d464788919
func Test_Issue1946(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		type B struct {
			init *gtype.Bool
			Name string
		}
		type A struct {
			B *B
		}
		a := &A{
			B: &B{
				init: gtype.NewBool(true),
			},
		}
		err := gconv.Struct(g.Map{
			"B": g.Map{
				"Name": "init",
			},
		}, a)
		t.AssertNil(err)
		t.Assert(a.B.Name, "init")
		t.Assert(a.B.init.Val(), true)
	})
	// 它不能改变私有属性。 md5:c15d94ed6929ce70
	gtest.C(t, func(t *gtest.T) {
		type B struct {
			init *gtype.Bool
			Name string
		}
		type A struct {
			B *B
		}
		a := &A{
			B: &B{
				init: gtype.NewBool(true),
			},
		}
		err := gconv.Struct(g.Map{
			"B": g.Map{
				"init": 0,
				"Name": "init",
			},
		}, a)
		t.AssertNil(err)
		t.Assert(a.B.Name, "init")
		t.Assert(a.B.init.Val(), true)
	})
	// 它可以改变公共属性。 md5:b2c3110608923730
	gtest.C(t, func(t *gtest.T) {
		type B struct {
			Init *gtype.Bool
			Name string
		}
		type A struct {
			B *B
		}
		a := &A{
			B: &B{
				Init: gtype.NewBool(),
			},
		}
		err := gconv.Struct(g.Map{
			"B": g.Map{
				"Init": 1,
				"Name": "init",
			},
		}, a)
		t.AssertNil(err)
		t.Assert(a.B.Name, "init")
		t.Assert(a.B.Init.Val(), true)
	})
}

// https://github.com/gogf/gf/issues/2381
// 
// 这个注释引用的是GitHub上的一个issue（问题或讨论），gf是一个Go语言的框架。具体来说，它指向了gf项目在github上编号为2381的问题或讨论。可能是开发者在提到某个与gf框架相关的问题时使用的链接，以便他人可以查看详细信息。 md5:d90662fa74c58cc1
func Test_Issue2381(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		type Inherit struct {
			Id        int64       `json:"id"          description:"Id"`
			Flag      *gjson.Json `json:"flag"        description:"标签"`
			Title     string      `json:"title"       description:"标题"`
			CreatedAt *gtime.Time `json:"createdAt"   description:"创建时间"`
		}
		type Test1 struct {
			Inherit
		}
		type Test2 struct {
			Inherit
		}
		var (
			a1 Test1
			a2 Test2
		)

		a1 = Test1{
			Inherit{
				Id:        2,
				Flag:      gjson.New("[1, 2]"),
				Title:     "测试",
				CreatedAt: gtime.Now(),
			},
		}
		err := gconv.Scan(a1, &a2)
		t.AssertNil(err)
		t.Assert(a1.Id, a2.Id)
		t.Assert(a1.Title, a2.Title)
		t.Assert(a1.CreatedAt, a2.CreatedAt)
		t.Assert(a1.Flag.String(), a2.Flag.String())
	})
}

// https://github.com/gogf/gf/issues/2391
// 
// 这段注释引用的是一个GitHub问题或 issue，它位于gogf（一个Go语言的框架）的仓库中。具体来说，这是第2391个问题或讨论。"gf"可能是 "goframe" 的简称，这是一个用Go语言编写的Web开发框架。这个链接指向开发者可以在其中查看、报告问题或参与讨论的地方。 md5:ba8514e6f11d7f01
func Test_Issue2391(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		type Inherit struct {
			Ids   []int
			Ids2  []int64
			Flag  *gjson.Json
			Title string
		}

		type Test1 struct {
			Inherit
		}
		type Test2 struct {
			Inherit
		}

		var (
			a1 Test1
			a2 Test2
		)

		a1 = Test1{
			Inherit{
				Ids:   []int{1, 2, 3},
				Ids2:  []int64{4, 5, 6},
				Flag:  gjson.New("[\"1\", \"2\"]"),
				Title: "测试",
			},
		}

		err := gconv.Scan(a1, &a2)
		t.AssertNil(err)
		t.Assert(a1.Ids, a2.Ids)
		t.Assert(a1.Ids2, a2.Ids2)
		t.Assert(a1.Title, a2.Title)
		t.Assert(a1.Flag.String(), a2.Flag.String())
	})
}

// 关于 issue #2395 的讨论，请访问：https://github.com/gogf/gf/issues/2395. md5:f21adc84f980809e
func Test_Issue2395(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		type Test struct {
			Num int
		}
		var ()
		obj := Test{Num: 0}
		t.Assert(gconv.Interfaces(obj), []interface{}{obj})
	})
}

// https://github.com/gogf/gf/issues/2371
// 
// 这段注释是在引用一个GitHub上的问题链接，该问题是关于gf框架的。在中文中可以翻译为：
// 
// https://github.com/gogf/gf/issues/2371
// 
// 这段注释是在引用一个GitHub上的问题链接，该问题是关于gf框架的。在中文中可以翻译为：
// 
// https://github.com/gogf/gf/issues/2371 有关gf框架的GitHub问题讨论. md5:c428de079857b105
func Test_Issue2371(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			s = struct {
				Time time.Time `json:"time"`
			}{}
			jsonMap = map[string]interface{}{"time": "2022-12-15 16:11:34"}
		)

		err := gconv.Struct(jsonMap, &s)
		t.AssertNil(err)
		t.Assert(s.Time.UTC(), `2022-12-15 08:11:34 +0000 UTC`)
	})
}

func Test_Issue2901(t *testing.T) {
	type GameApp2 struct {
		ForceUpdateTime *time.Time
	}
	gtest.C(t, func(t *gtest.T) {
		src := map[string]interface{}{
			"FORCE_UPDATE_TIME": time.Now(),
		}
		m := GameApp2{}
		err := gconv.Scan(src, &m)
		t.AssertNil(err)
	})
}
