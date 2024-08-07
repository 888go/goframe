// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 工具类_test

import (
	"bytes"
	"testing"

	gtype "github.com/888go/goframe/container/gtype"
	"github.com/888go/goframe/frame/g"
	ghttp "github.com/888go/goframe/net/ghttp"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
	gstr "github.com/888go/goframe/text/gstr"
	gmeta "github.com/888go/goframe/util/gmeta"
	gutil "github.com/888go/goframe/util/gutil"
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
		CreatedAt: gtime.X创建并按当前时间(),
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
		gutil.X调试输出(map[int]int{
			100: 100,
		})
		gutil.X调试输出(req)
		gutil.X调试输出(true, false)
		gutil.X调试输出(make(chan int))
		gutil.X调试输出(func() {})
		gutil.X调试输出(nil)
		gutil.X调试输出(gtype.NewInt(1))
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
		gutil.X调试输出到Writer(buffer, m, gutil.DumpOption{})
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
		CreatedAt: gtime.X创建并按当前时间(),
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
		gutil.X调试输出并带类型(map[int]int{
			100: 100,
		})
		gutil.X调试输出并带类型(req)
		gutil.X调试输出并带类型([][]byte{[]byte("hello")})
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
		gutil.X调试输出(req)
		gutil.X调试输出(req.Content)

		gutil.X调试输出并带类型(req)
		gutil.X调试输出并带类型(req.Content)
	})
}

// 这段注释链接指向的是GitHub上的一个 issues（问题或讨论），GF（GoGF）是一个Go语言的框架。具体来说，这是关于GF框架的一个Issue，编号为1661，可能是开发者社区中报告的问题、建议或讨论的内容。要了解详情，可以点击链接查看相关讨论。 md5:2af841e765567898
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
		gutil.X调试输出到Writer(buffer, q2, gutil.DumpOption{})
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
		g.X调试输出到Writer(buffer, abc, gutil.DumpOption{})
		t.Assert(gstr.X是否包含(buffer.String(), "cycle"), true)
	})
}

func Test_DumpJson(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var jsonContent = `{"a":1,"b":2}`
		gutil.X调试输出json(jsonContent)
	})
}
