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

	gtest.C(t, func(t *gtest.T) {
		m := db.Model(table).Hook(gdb.HookHandler{
			Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
				result, err = in.Next(ctx)
				if err != nil {
					return
				}
				for i, record := range result {
					record["test"] = gvar.New(100 + record["id"].Int())
					result[i] = record
				}
				return
			},
		})
		all, err := m.Where(`id > 6`).OrderAsc(`id`).All()
		t.AssertNil(err)
		t.Assert(len(all), 4)
		t.Assert(all[0]["id"].Int(), 7)
		t.Assert(all[0]["test"].Int(), 107)
		t.Assert(all[1]["test"].Int(), 108)
		t.Assert(all[2]["test"].Int(), 109)
		t.Assert(all[3]["test"].Int(), 110)
	})
}

func Test_Model_Hook_Insert(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		m := db.Model(table).Hook(gdb.HookHandler{
			Insert: func(ctx context.Context, in *gdb.HookInsertInput) (result sql.Result, err error) {
				for i, item := range in.Data {
					item["passport"] = fmt.Sprintf(`test_port_%d`, item["id"])
					item["nickname"] = fmt.Sprintf(`test_name_%d`, item["id"])
					in.Data[i] = item
				}
				return in.Next(ctx)
			},
		})
		_, err := m.Insert(g.Map{
			"id":       1,
			"nickname": "name_1",
		})
		t.AssertNil(err)
		one, err := m.One()
		t.AssertNil(err)
		t.Assert(one["id"].Int(), 1)
		t.Assert(one["passport"], `test_port_1`)
		t.Assert(one["nickname"], `test_name_1`)
	})
}

func Test_Model_Hook_Update(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		m := db.Model(table).Hook(gdb.HookHandler{
			Update: func(ctx context.Context, in *gdb.HookUpdateInput) (result sql.Result, err error) {
				switch value := in.Data.(type) {
				case gdb.List:
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
		_, err := m.Data(g.Map{
			"nickname": "name_1",
		}).WherePri(1).Update()
		t.AssertNil(err)

		one, err := m.One()
		t.AssertNil(err)
		t.Assert(one["id"].Int(), 1)
		t.Assert(one["passport"], `port`)
		t.Assert(one["nickname"], `name`)
	})
}

func Test_Model_Hook_Delete(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		m := db.Model(table).Hook(gdb.HookHandler{
			Delete: func(ctx context.Context, in *gdb.HookDeleteInput) (result sql.Result, err error) {
				return db.Model(table).Data(g.Map{
					"nickname": `deleted`,
				}).Where(in.Condition).Update()
			},
		})
		_, err := m.Where(1).Delete()
		t.AssertNil(err)

		all, err := m.All()
		t.AssertNil(err)
		for _, item := range all {
			t.Assert(item["nickname"].String(), `deleted`)
		}
	})
}
