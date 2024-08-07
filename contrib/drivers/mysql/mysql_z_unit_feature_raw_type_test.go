// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package mysql_test

import (
	"testing"

	gdb "github.com/888go/goframe/database/gdb"
	"github.com/888go/goframe/frame/g"
	gtest "github.com/888go/goframe/test/gtest"
)

func Test_Insert_Raw(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		user := db.X创建Model对象(table)
		result, err := user.X设置数据(g.Map{
			"id":          gdb.Raw("id+2"),
			"passport":    "port_1",
			"password":    "pass_1",
			"nickname":    "name_1",
			"create_time": gdb.Raw("now()"),
		}).X插入()
		t.AssertNil(err)
		n, _ := result.LastInsertId()
		t.Assert(n, 2)
	})
}

func Test_BatchInsert_Raw(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		user := db.X创建Model对象(table)
		result, err := user.X设置数据(
			g.Map切片{
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
		).X插入()
		t.AssertNil(err)
		n, _ := result.LastInsertId()
		t.Assert(n, 4)
	})
}

func Test_Update_Raw(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		user := db.X创建Model对象(table)
		result, err := user.X设置数据(g.Map{
			"id":          gdb.Raw("id+100"),
			"create_time": gdb.Raw("now()"),
		}).X条件("id", 1).X更新()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
	})
	gtest.C(t, func(t *gtest.T) {
		user := db.X创建Model对象(table)
		n, err := user.X条件("id", 101).X查询行数()
		t.AssertNil(err)
		t.Assert(n, 1)
	})
}
