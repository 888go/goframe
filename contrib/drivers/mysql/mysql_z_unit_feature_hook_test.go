// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package mysql_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	gvar "github.com/888go/goframe/container/gvar"
	gdb "github.com/888go/goframe/database/gdb"
	"github.com/888go/goframe/frame/g"
	gtest "github.com/888go/goframe/test/gtest"
)

func Test_Model_Hook_Select(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		m := db.X创建Model对象(table).Hook(gdb.HookHandler{
			Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
				result, err = in.Next(ctx)
				if err != nil {
					return
				}
				for i, record := range result {
					record["test"] = gvar.X创建(100 + record["id"].X取整数())
					result[i] = record
				}
				return
			},
		})
		all, err := m.X条件(`id > 6`).X排序ASC(`id`).X查询()
		t.AssertNil(err)
		t.Assert(len(all), 4)
		t.Assert(all[0]["id"].X取整数(), 7)
		t.Assert(all[0]["test"].X取整数(), 107)
		t.Assert(all[1]["test"].X取整数(), 108)
		t.Assert(all[2]["test"].X取整数(), 109)
		t.Assert(all[3]["test"].X取整数(), 110)
	})
}

func Test_Model_Hook_Insert(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		m := db.X创建Model对象(table).Hook(gdb.HookHandler{
			Insert: func(ctx context.Context, in *gdb.HookInsertInput) (result sql.Result, err error) {
				for i, item := range in.Data {
					item["passport"] = fmt.Sprintf(`test_port_%d`, item["id"])
					item["nickname"] = fmt.Sprintf(`test_name_%d`, item["id"])
					in.Data[i] = item
				}
				return in.Next(ctx)
			},
		})
		_, err := m.X插入(g.Map{
			"id":       1,
			"nickname": "name_1",
		})
		t.AssertNil(err)
		one, err := m.X查询一条()
		t.AssertNil(err)
		t.Assert(one["id"].X取整数(), 1)
		t.Assert(one["passport"], `test_port_1`)
		t.Assert(one["nickname"], `test_name_1`)
	})
}

func Test_Model_Hook_Update(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		m := db.X创建Model对象(table).Hook(gdb.HookHandler{
			Update: func(ctx context.Context, in *gdb.HookUpdateInput) (result sql.Result, err error) {
				switch value := in.Data.(type) {
				case gdb.Map切片:
					for i, data := range value {
						data["passport"] = `port`
						data["nickname"] = `name`
						value[i] = data
					}
					in.Data = value

				case gdb.Map:
					value["passport"] = `port`
					value["nickname"] = `name`
					in.Data = value
				}
				return in.Next(ctx)
			},
		})
		_, err := m.X设置数据(g.Map{
			"nickname": "name_1",
		}).X条件并识别主键(1).X更新()
		t.AssertNil(err)

		one, err := m.X查询一条()
		t.AssertNil(err)
		t.Assert(one["id"].X取整数(), 1)
		t.Assert(one["passport"], `port`)
		t.Assert(one["nickname"], `name`)
	})
}

func Test_Model_Hook_Delete(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		m := db.X创建Model对象(table).Hook(gdb.HookHandler{
			Delete: func(ctx context.Context, in *gdb.HookDeleteInput) (result sql.Result, err error) {
				return db.X创建Model对象(table).X设置数据(g.Map{
					"nickname": `deleted`,
				}).X条件(in.Condition).X更新()
			},
		})
		_, err := m.X条件(1).X删除()
		t.AssertNil(err)

		all, err := m.X查询()
		t.AssertNil(err)
		for _, item := range all {
			t.Assert(item["nickname"].String(), `deleted`)
		}
	})
}
