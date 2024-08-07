// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package mysql_test

import (
	"context"
	"testing"

	gdb "github.com/888go/goframe/database/gdb"
	glog "github.com/888go/goframe/os/glog"
	gtest "github.com/888go/goframe/test/gtest"
)

func Test_Ctx(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		db, err := gdb.X取单例对象()
		t.AssertNil(err)

		err1 := db.X向主节点发送心跳()
		err2 := db.X向从节点发送心跳()
		t.Assert(err1, nil)
		t.Assert(err2, nil)

		newDb := db.X设置上下文并取副本(context.Background())
		t.AssertNE(newDb, nil)
	})
}

func Test_Ctx_Query(t *testing.T) {
	db.X取日志记录器().(*glog.Logger).X设置上下文名称("SpanId", "TraceId")
	gtest.C(t, func(t *gtest.T) {
		db.X设置调试模式(true)
		defer db.X设置调试模式(false)
		ctx := context.WithValue(context.Background(), "TraceId", "12345678")
		ctx = context.WithValue(ctx, "SpanId", "0.1")
		db.X原生SQL查询(ctx, "select 1")
	})
	gtest.C(t, func(t *gtest.T) {
		db.X设置调试模式(true)
		defer db.X设置调试模式(false)
		db.X原生SQL查询(ctx, "select 2")
	})
}

func Test_Ctx_Model(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	db.X取日志记录器().(*glog.Logger).X设置上下文名称("SpanId", "TraceId")
	gtest.C(t, func(t *gtest.T) {
		db.X设置调试模式(true)
		defer db.X设置调试模式(false)
		ctx := context.WithValue(context.Background(), "TraceId", "12345678")
		ctx = context.WithValue(ctx, "SpanId", "0.1")
		db.X创建Model对象(table).X设置上下文并取副本(ctx).X查询()
	})
	gtest.C(t, func(t *gtest.T) {
		db.X设置调试模式(true)
		defer db.X设置调试模式(false)
		db.X创建Model对象(table).X查询()
	})
}
