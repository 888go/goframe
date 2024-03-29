// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 工具类_test

import (
	"bytes"
	gutil2 "github.com/gogf/gf/v2/util/gutil"
	"testing"

	"github.com/888go/goframe/gutil"
	"github.com/gogf/gf/v2/container/gtype"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gmeta"
)

func Test_Dump(t *testing.T) {
	type CommonReq struct {
		AppId      int64  `json:"appId" v:"required" in:"path" des:"应用Id" sum:"应用Id Summary"`
		ResourceId string `json:"resourceId" in:"query" des:"资源Id" sum:"资源Id Summary"`
	}
	type SetSpecInfo struct {
		StorageType string   `v:"required|in:CLOUD_PREMIUM,CLOUD_SSD,CLOUD_HSSD" des:"StorageType"`
		Shards      int32    `des:"shards 分片数" sum:"Shards Summary"`
		Params      []string `des:"默认参数(json 串-ClickHouseParams)" sum:"Params Summary"`
	}
	type CreateResourceReq struct {
		CommonReq
		gmeta.Meta `path:"/CreateResourceReq" method:"POST" tags:"default" sum:"CreateResourceReq sum"`
		Name       string
		CreatedAt  *gtime.Time
		SetMap     map[string]*SetSpecInfo
		SetSlice   []SetSpecInfo
		Handler    ghttp.HandlerFunc
		internal   string
	}
	req := &CreateResourceReq{
		CommonReq: CommonReq{
			AppId:      12345678,
			ResourceId: "tdchqy-xxx",
		},
		Name:      "john",
		CreatedAt: gtime.Now(),
		SetMap: map[string]*SetSpecInfo{
			"test1": {
				StorageType: "ssd",
				Shards:      2,
				Params:      []string{"a", "b", "c"},
			},
			"test2": {
				StorageType: "hssd",
				Shards:      10,
				Params:      []string{},
			},
		},
		SetSlice: []SetSpecInfo{
			{
				StorageType: "hssd",
				Shards:      10,
				Params:      []string{"h"},
			},
		},
	}
	gtest.C(t, func(t *gtest.T) {
		工具类.X调试输出(map[int]int{
			100: 100,
		})
		工具类.X调试输出(req)
		工具类.X调试输出(true, false)
		工具类.X调试输出(make(chan int))
		工具类.X调试输出(func() {})
		工具类.X调试输出(nil)
		工具类.X调试输出(gtype.NewInt(1))
	})
}

func Test_Dump_Map(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		buffer := bytes.NewBuffer(nil)
		m := g.Map{
			"k1": g.Map{
				"k2": "v2",
			},
		}
		工具类.X调试输出到Writer(buffer, m, 工具类.DumpOption{})
		t.Assert(buffer.String(), `{
    "k1": {
        "k2": "v2",
    },
}`)
	})
}

func TestDumpWithType(t *testing.T) {
	type CommonReq struct {
		AppId      int64  `json:"appId" v:"required" in:"path" des:"应用Id" sum:"应用Id Summary"`
		ResourceId string `json:"resourceId" in:"query" des:"资源Id" sum:"资源Id Summary"`
	}
	type SetSpecInfo struct {
		StorageType string   `v:"required|in:CLOUD_PREMIUM,CLOUD_SSD,CLOUD_HSSD" des:"StorageType"`
		Shards      int32    `des:"shards 分片数" sum:"Shards Summary"`
		Params      []string `des:"默认参数(json 串-ClickHouseParams)" sum:"Params Summary"`
	}
	type CreateResourceReq struct {
		CommonReq
		gmeta.Meta `path:"/CreateResourceReq" method:"POST" tags:"default" sum:"CreateResourceReq sum"`
		Name       string
		CreatedAt  *gtime.Time
		SetMap     map[string]*SetSpecInfo `v:"required" des:"配置Map"`
		SetSlice   []SetSpecInfo           `v:"required" des:"配置Slice"`
		Handler    ghttp.HandlerFunc
		internal   string
	}
	req := &CreateResourceReq{
		CommonReq: CommonReq{
			AppId:      12345678,
			ResourceId: "tdchqy-xxx",
		},
		Name:      "john",
		CreatedAt: gtime.Now(),
		SetMap: map[string]*SetSpecInfo{
			"test1": {
				StorageType: "ssd",
				Shards:      2,
				Params:      []string{"a", "b", "c"},
			},
			"test2": {
				StorageType: "hssd",
				Shards:      10,
				Params:      []string{},
			},
		},
		SetSlice: []SetSpecInfo{
			{
				StorageType: "hssd",
				Shards:      10,
				Params:      []string{"h"},
			},
		},
	}
	gtest.C(t, func(t *gtest.T) {
		工具类.X调试输出并带类型(map[int]int{
			100: 100,
		})
		工具类.X调试输出并带类型(req)
		工具类.X调试输出并带类型([][]byte{[]byte("hello")})
	})
}

func Test_Dump_Slashes(t *testing.T) {
	type Req struct {
		Content string
	}
	req := &Req{
		Content: `{"name":"john", "age":18}`,
	}
	gtest.C(t, func(t *gtest.T) {
		工具类.X调试输出(req)
		工具类.X调试输出(req.Content)

		工具类.X调试输出并带类型(req)
		工具类.X调试输出并带类型(req.Content)
	})
}

// 这是Go语言代码中的一行注释，其内容为一个URL链接，指向GitHub上gogf/gf项目的一个issue（问题）页面，编号为1661。
// 中文翻译：
// 这是Go语言代码中的一个注释，它给出了一个GitHub上gogf/gf项目的Issue（问题）链接，具体问题是第1661号。
func Test_Dump_Issue1661(t *testing.T) {
	type B struct {
		ba int
		bb string
	}
	type A struct {
		aa int
		ab string
		cc []B
	}
	gtest.C(t, func(t *gtest.T) {
		var q1 []A
		var q2 []A
		q2 = make([]A, 0)
		q1 = []A{{aa: 1, ab: "1", cc: []B{{ba: 1}, {ba: 2}, {ba: 3}}}, {aa: 2, ab: "2", cc: []B{{ba: 1}, {ba: 2}, {ba: 3}}}}
		for _, q1v := range q1 {
			x := []string{"11", "22"}
			for _, iv2 := range x {
				ls := q1v
				for i := range ls.cc {
					sj := iv2
					ls.cc[i].bb = sj
				}
				q2 = append(q2, ls)
			}
		}
		buffer := bytes.NewBuffer(nil)
		工具类.X调试输出到Writer(buffer, q2, 工具类.DumpOption{})
		t.Assert(buffer.String(), `[
    {
        aa: 1,
        ab: "1",
        cc: [
            {
                ba: 1,
                bb: "22",
            },
            {
                ba: 2,
                bb: "22",
            },
            {
                ba: 3,
                bb: "22",
            },
        ],
    },
    {
        aa: 1,
        ab: "1",
        cc: [
            {
                ba: 1,
                bb: "22",
            },
            {
                ba: 2,
                bb: "22",
            },
            {
                ba: 3,
                bb: "22",
            },
        ],
    },
    {
        aa: 2,
        ab: "2",
        cc: [
            {
                ba: 1,
                bb: "22",
            },
            {
                ba: 2,
                bb: "22",
            },
            {
                ba: 3,
                bb: "22",
            },
        ],
    },
    {
        aa: 2,
        ab: "2",
        cc: [
            {
                ba: 1,
                bb: "22",
            },
            {
                ba: 2,
                bb: "22",
            },
            {
                ba: 3,
                bb: "22",
            },
        ],
    },
]`)
	})
}

func Test_Dump_Cycle_Attribute(t *testing.T) {
	type Abc struct {
		ab int
		cd *Abc
	}
	abc := Abc{ab: 3}
	abc.cd = &abc
	gtest.C(t, func(t *gtest.T) {
		buffer := bytes.NewBuffer(nil)
		g.DumpTo(buffer, abc, gutil2.DumpOption(工具类.DumpOption{}))
		t.Assert(gstr.Contains(buffer.String(), "cycle"), true)
	})
}

func Test_DumpJson(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var jsonContent = `{"a":1,"b":2}`
		工具类.X调试输出json(jsonContent)
	})
}
