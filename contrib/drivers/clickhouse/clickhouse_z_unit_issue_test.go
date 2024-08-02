//go:build 屏蔽单元测试

// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package clickhouse

import (
	"github.com/shopspring/decimal"
	"testing"

	gvar "github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/frame/g"
	gtest "github.com/888go/goframe/test/gtest"
	gconv "github.com/888go/goframe/util/gconv"
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
