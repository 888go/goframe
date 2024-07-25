// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package mysql_test

import (
	"context"
	"testing"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/test/gtest"
)

func Test_Ctx(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		db, err := gdb.Instance()
		t.AssertNil(err)

		err1 := db.PingMaster()
		err2 := db.PingSlave()
		t.Assert(err1, nil)
		t.Assert(err2, nil)

		newDb := db.Ctx(context.Background())
		t.AssertNE(newDb, nil)
	})
}

func Test_Ctx_Query(t *testing.T) {
	db.GetLogger().(*glog.Logger).SetCtxKeys("SpanId", "TraceId")
	gtest.C(t, func(t *gtest.T) {
		db.SetDebug(true)
		defer db.SetDebug(false)
		ctx := context.WithValue(context.Background(), "TraceId", "12345678")
		ctx = context.WithValue(ctx, "SpanId", "0.1")
		db.Query(ctx, "select 1")
	})
	gtest.C(t, func(t *gtest.T) {
		db.SetDebug(true)
		defer db.SetDebug(false)
		db.Query(ctx, "select 2")
	})
}

func Test_Ctx_Model(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	db.GetLogger().(*glog.Logger).SetCtxKeys("SpanId", "TraceId")
	gtest.C(t, func(t *gtest.T) {
		db.SetDebug(true)
		defer db.SetDebug(false)
		ctx := context.WithValue(context.Background(), "TraceId", "12345678")
		ctx = context.WithValue(ctx, "SpanId", "0.1")
		db.Model(table).Ctx(ctx).All()
	})
	gtest.C(t, func(t *gtest.T) {
		db.SetDebug(true)
		defer db.SetDebug(false)
		db.Model(table).All()
	})
}
