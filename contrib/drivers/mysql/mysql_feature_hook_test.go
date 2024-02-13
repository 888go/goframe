// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package mysql_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/database/gdb"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/test/gtest"
)

func Test_Model_Hook_Select(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		m := db.X创建Model对象(table).Hook(db类.HookHandler{
			Select: func(ctx context.Context, in *db类.HookSelectInput) (result db类.Result, err error) {
				result, err = in.Next(ctx)
				if err != nil {
					return
				}
				for i, record := range result {
					record["test"] = 泛型类.X创建(100 + record["id"].X取整数())
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

	单元测试类.C(t, func(t *单元测试类.T) {
		m := db.X创建Model对象(table).Hook(db类.HookHandler{
			Insert: func(ctx context.Context, in *db类.HookInsertInput) (result sql.Result, err error) {
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

	单元测试类.C(t, func(t *单元测试类.T) {
		m := db.X创建Model对象(table).Hook(db类.HookHandler{
			Update: func(ctx context.Context, in *db类.HookUpdateInput) (result sql.Result, err error) {
				switch value := in.Data.(type) {
				case db类.Map数组:
					for i, data := range value {
						data["passport"] = `port`
						data["nickname"] = `name`
						value[i] = data
					}
					in.Data = value

				case db类.Map:
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

	单元测试类.C(t, func(t *单元测试类.T) {
		m := db.X创建Model对象(table).Hook(db类.HookHandler{
			Delete: func(ctx context.Context, in *db类.HookDeleteInput) (result sql.Result, err error) {
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
