// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package mysql_test

import (
	"context"
	"testing"
	
	"github.com/888go/goframe/database/gdb"
	"github.com/888go/goframe/os/glog"
	"github.com/888go/goframe/test/gtest"
)

func Test_Ctx(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		db, err := db类.X取单例对象()
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
	db.X取日志记录器().(*日志类.Logger).X设置上下文名称("SpanId", "TraceId")
	单元测试类.C(t, func(t *单元测试类.T) {
		db.X设置调试模式(true)
		defer db.X设置调试模式(false)
		ctx := context.WithValue(context.Background(), "TraceId", "12345678")
		ctx = context.WithValue(ctx, "SpanId", "0.1")
		db.X原生SQL查询(ctx, "select 1")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		db.X设置调试模式(true)
		defer db.X设置调试模式(false)
		db.X原生SQL查询(ctx, "select 2")
	})
}

func Test_Ctx_Model(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	db.X取日志记录器().(*日志类.Logger).X设置上下文名称("SpanId", "TraceId")
	单元测试类.C(t, func(t *单元测试类.T) {
		db.X设置调试模式(true)
		defer db.X设置调试模式(false)
		ctx := context.WithValue(context.Background(), "TraceId", "12345678")
		ctx = context.WithValue(ctx, "SpanId", "0.1")
		db.X创建Model对象(table).X设置上下文并取副本(ctx).X查询()
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		db.X设置调试模式(true)
		defer db.X设置调试模式(false)
		db.X创建Model对象(table).X查询()
	})
}
