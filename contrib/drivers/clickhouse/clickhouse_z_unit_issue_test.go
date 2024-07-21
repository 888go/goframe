//go:build 屏蔽单元测试

// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package clickhouse

import (
	"github.com/shopspring/decimal"
	"testing"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gconv"
)

func Test_Issue2584(t *testing.T) {
	type TDecimal struct {
		F1 *decimal.Decimal `json:"f1"`
	}

	gtest.C(t, func(t *gtest.T) {
		var (
			p1    = TDecimal{}
			data1 = g.Map{"f1": gvar.New(1111.111)}
			err   = gconv.Scan(data1, &p1)
		)
		t.AssertNil(err)
		t.Assert(p1.F1, 1111.111)
	})

	gtest.C(t, func(t *gtest.T) {
		var (
			p2    = TDecimal{}
			data2 = g.Map{"f1": gvar.New("2222.222")}
			err   = gconv.Scan(data2, &p2)
		)
		t.AssertNil(err)
		t.Assert(p2.F1, 2222.222)
	})
}
