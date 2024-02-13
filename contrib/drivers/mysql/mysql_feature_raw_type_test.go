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

	单元测试类.C(t, func(t *单元测试类.T) {
		user := db.X创建Model对象(table)
		result, err := user.X设置数据(g.Map{
			"id":          db类.Raw("id+2"),
			"passport":    "port_1",
			"password":    "pass_1",
			"nickname":    "name_1",
			"create_time": db类.Raw("now()"),
		}).X插入()
		t.AssertNil(err)
		n, _ := result.LastInsertId()
		t.Assert(n, 2)
	})
}

func Test_BatchInsert_Raw(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		user := db.X创建Model对象(table)
		result, err := user.X设置数据(
			g.Map数组{
				g.Map{
					"id":          db类.Raw("id+2"),
					"passport":    "port_2",
					"password":    "pass_2",
					"nickname":    "name_2",
					"create_time": db类.Raw("now()"),
				},
				g.Map{
					"id":          db类.Raw("id+4"),
					"passport":    "port_4",
					"password":    "pass_4",
					"nickname":    "name_4",
					"create_time": db类.Raw("now()"),
				},
			},
		).X插入()
		t.AssertNil(err)
		n, _ := result.LastInsertId()
		t.Assert(n, 4)
	})
}

func Test_Update_Raw(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		user := db.X创建Model对象(table)
		result, err := user.X设置数据(g.Map{
			"id":          db类.Raw("id+100"),
			"create_time": db类.Raw("now()"),
		}).X条件("id", 1).X更新()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		user := db.X创建Model对象(table)
		n, err := user.X条件("id", 101).X查询行数()
		t.AssertNil(err)
		t.Assert(n, 1)
	})
}
