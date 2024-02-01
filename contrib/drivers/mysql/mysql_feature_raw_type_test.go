// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package mysql_test
import (
	"testing"
	
	"github.com/888go/goframe/database/gdb"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/test/gtest"
	)

func Test_Insert_Raw(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		user := db.Model(table)
		result, err := user.Data(g.Map{
			"id":          gdb.Raw("id+2"),
			"passport":    "port_1",
			"password":    "pass_1",
			"nickname":    "name_1",
			"create_time": gdb.Raw("now()"),
		}).Insert()
		t.AssertNil(err)
		n, _ := result.LastInsertId()
		t.Assert(n, 2)
	})
}

func Test_BatchInsert_Raw(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		user := db.Model(table)
		result, err := user.Data(
			g.List{
				g.Map{
					"id":          gdb.Raw("id+2"),
					"passport":    "port_2",
					"password":    "pass_2",
					"nickname":    "name_2",
					"create_time": gdb.Raw("now()"),
				},
				g.Map{
					"id":          gdb.Raw("id+4"),
					"passport":    "port_4",
					"password":    "pass_4",
					"nickname":    "name_4",
					"create_time": gdb.Raw("now()"),
				},
			},
		).Insert()
		t.AssertNil(err)
		n, _ := result.LastInsertId()
		t.Assert(n, 4)
	})
}

func Test_Update_Raw(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		user := db.Model(table)
		result, err := user.Data(g.Map{
			"id":          gdb.Raw("id+100"),
			"create_time": gdb.Raw("now()"),
		}).Where("id", 1).Update()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
	})
	gtest.C(t, func(t *gtest.T) {
		user := db.Model(table)
		n, err := user.Where("id", 101).Count()
		t.AssertNil(err)
		t.Assert(n, 1)
	})
}
