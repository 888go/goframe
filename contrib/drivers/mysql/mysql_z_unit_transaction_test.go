// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package mysql_test

import (
	"context"
	"fmt"
	"testing"

	gdb "github.com/888go/goframe/database/gdb"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/frame/g"
	gctx "github.com/888go/goframe/os/gctx"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
)

func Test_TX_Query(t *testing.T) {
	tx, err := db.X事务开启(ctx)
	if err != nil {
		gtest.Error(err)
	}
	if _, err = tx.X原生SQL查询("SELECT ?", 1); err != nil {
		gtest.Error(err)
	}
	if _, err = tx.X原生SQL查询("SELECT ?+?", 1, 2); err != nil {
		gtest.Error(err)
	}
	if _, err = tx.X原生SQL查询("SELECT ?+?", g.Slice别名{1, 2}); err != nil {
		gtest.Error(err)
	}
	if _, err = tx.X原生SQL查询("ERROR"); err == nil {
		gtest.Error("FAIL")
	}
	if err = tx.X事务提交(); err != nil {
		gtest.Error(err)
	}
}

func Test_TX_Exec(t *testing.T) {
	tx, err := db.X事务开启(ctx)
	if err != nil {
		gtest.Error(err)
	}
	if _, err := tx.X原生SQL执行("SELECT ?", 1); err != nil {
		gtest.Error(err)
	}
	if _, err := tx.X原生SQL执行("SELECT ?+?", 1, 2); err != nil {
		gtest.Error(err)
	}
	if _, err := tx.X原生SQL执行("SELECT ?+?", g.Slice别名{1, 2}); err != nil {
		gtest.Error(err)
	}
	if _, err := tx.X原生SQL执行("ERROR"); err == nil {
		gtest.Error("FAIL")
	}
	if err := tx.X事务提交(); err != nil {
		gtest.Error(err)
	}
}

func Test_TX_Commit(t *testing.T) {
	tx, err := db.X事务开启(ctx)
	if err != nil {
		gtest.Error(err)
	}
	if err := tx.X事务提交(); err != nil {
		gtest.Error(err)
	}
}

func Test_TX_Rollback(t *testing.T) {
	tx, err := db.X事务开启(ctx)
	if err != nil {
		gtest.Error(err)
	}
	if err := tx.X事务回滚(); err != nil {
		gtest.Error(err)
	}
}

func Test_TX_Prepare(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		tx, err := db.X事务开启(ctx)
		t.AssertNil(err)

		st, err := tx.X原生sql取参数预处理对象("SELECT 100")
		t.AssertNil(err)

		rows, err := st.X查询()
		t.AssertNil(err)

		array, err := rows.Columns()
		t.AssertNil(err)
		t.Assert(array[0], "100")

		rows.Close()
		t.AssertNil(err)

		tx.X事务提交()
		t.AssertNil(err)
	})
}

func Test_TX_Insert(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		tx, err := db.X事务开启(ctx)
		if err != nil {
			gtest.Error(err)
		}
		user := tx.X创建Model对象(table)
		if _, err := user.X设置数据(g.Map{
			"id":          1,
			"passport":    "t1",
			"password":    "25d55ad283aa400af464c76d713c07ad",
			"nickname":    "T1",
			"create_time": gtime.X创建并按当前时间().String(),
		}).X插入(); err != nil {
			gtest.Error(err)
		}
		if _, err := tx.X插入(table, g.Map{
			"id":          2,
			"passport":    "t1",
			"password":    "25d55ad283aa400af464c76d713c07ad",
			"nickname":    "T1",
			"create_time": gtime.X创建并按当前时间().String(),
		}); err != nil {
			gtest.Error(err)
		}

		if n, err := tx.X创建Model对象(table).X查询行数(); err != nil {
			gtest.Error(err)
		} else {
			t.Assert(n, int64(2))
		}

		if err := tx.X事务提交(); err != nil {
			gtest.Error(err)
		}
	})
}

func Test_TX_BatchInsert(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		tx, err := db.X事务开启(ctx)
		if err != nil {
			gtest.Error(err)
		}
		if _, err := tx.X插入(table, g.Map切片{
			{
				"id":          2,
				"passport":    "t",
				"password":    "25d55ad283aa400af464c76d713c07ad",
				"nickname":    "T2",
				"create_time": gtime.X创建并按当前时间().String(),
			},
			{
				"id":          3,
				"passport":    "t3",
				"password":    "25d55ad283aa400af464c76d713c07ad",
				"nickname":    "T3",
				"create_time": gtime.X创建并按当前时间().String(),
			},
		}, 10); err != nil {
			gtest.Error(err)
		}
		if err := tx.X事务提交(); err != nil {
			gtest.Error(err)
		}
		if n, err := db.X创建Model对象(table).X查询行数(); err != nil {
			gtest.Error(err)
		} else {
			t.Assert(n, int64(2))
		}
	})
}

func Test_TX_BatchReplace(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		tx, err := db.X事务开启(ctx)
		if err != nil {
			gtest.Error(err)
		}
		if _, err := tx.X插入并替换已存在(table, g.Map切片{
			{
				"id":          2,
				"passport":    "USER_2",
				"password":    "PASS_2",
				"nickname":    "NAME_2",
				"create_time": gtime.X创建并按当前时间().String(),
			},
			{
				"id":          4,
				"passport":    "USER_4",
				"password":    "PASS_4",
				"nickname":    "NAME_4",
				"create_time": gtime.X创建并按当前时间().String(),
			},
		}, 10); err != nil {
			gtest.Error(err)
		}
		if err := tx.X事务提交(); err != nil {
			gtest.Error(err)
		}
		if n, err := db.X创建Model对象(table).X查询行数(); err != nil {
			gtest.Error(err)
		} else {
			t.Assert(n, int64(TableSize))
		}
		if value, err := db.X创建Model对象(table).X字段保留过滤("password").X条件("id", 2).X查询一条值(); err != nil {
			gtest.Error(err)
		} else {
			t.Assert(value.String(), "PASS_2")
		}
	})
}

func Test_TX_BatchSave(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		tx, err := db.X事务开启(ctx)
		if err != nil {
			gtest.Error(err)
		}
		if _, err := tx.X插入并更新已存在(table, g.Map切片{
			{
				"id":          4,
				"passport":    "USER_4",
				"password":    "PASS_4",
				"nickname":    "NAME_4",
				"create_time": gtime.X创建并按当前时间().String(),
			},
		}, 10); err != nil {
			gtest.Error(err)
		}
		if err := tx.X事务提交(); err != nil {
			gtest.Error(err)
		}

		if n, err := db.X创建Model对象(table).X查询行数(); err != nil {
			gtest.Error(err)
		} else {
			t.Assert(n, int64(TableSize))
		}

		if value, err := db.X创建Model对象(table).X字段保留过滤("password").X条件("id", 4).X查询一条值(); err != nil {
			gtest.Error(err)
		} else {
			t.Assert(value.String(), "PASS_4")
		}
	})
}

func Test_TX_Replace(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		tx, err := db.X事务开启(ctx)
		if err != nil {
			gtest.Error(err)
		}
		if _, err := tx.X插入并替换已存在(table, g.Map{
			"id":          1,
			"passport":    "USER_1",
			"password":    "PASS_1",
			"nickname":    "NAME_1",
			"create_time": gtime.X创建并按当前时间().String(),
		}); err != nil {
			gtest.Error(err)
		}
		if err := tx.X事务回滚(); err != nil {
			gtest.Error(err)
		}
		if value, err := db.X创建Model对象(table).X字段保留过滤("nickname").X条件("id", 1).X查询一条值(); err != nil {
			gtest.Error(err)
		} else {
			t.Assert(value.String(), "name_1")
		}
	})
}

func Test_TX_Save(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		tx, err := db.X事务开启(ctx)
		if err != nil {
			gtest.Error(err)
		}
		if _, err := tx.X插入并更新已存在(table, g.Map{
			"id":          1,
			"passport":    "USER_1",
			"password":    "PASS_1",
			"nickname":    "NAME_1",
			"create_time": gtime.X创建并按当前时间().String(),
		}); err != nil {
			gtest.Error(err)
		}
		if err := tx.X事务提交(); err != nil {
			gtest.Error(err)
		}
		if value, err := db.X创建Model对象(table).X字段保留过滤("nickname").X条件("id", 1).X查询一条值(); err != nil {
			gtest.Error(err)
		} else {
			t.Assert(value.String(), "NAME_1")
		}
	})
}

func Test_TX_Update(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		tx, err := db.X事务开启(ctx)
		if err != nil {
			gtest.Error(err)
		}
		if result, err := tx.X更新(table, "create_time='2019-10-24 10:00:00'", "id=3"); err != nil {
			gtest.Error(err)
		} else {
			n, _ := result.RowsAffected()
			t.Assert(n, 1)
		}
		if err := tx.X事务提交(); err != nil {
			gtest.Error(err)
		}
		_, err = tx.X创建Model对象(table).X字段保留过滤("create_time").X条件("id", 3).X查询一条值()
		t.AssertNE(err, nil)

		if value, err := db.X创建Model对象(table).X字段保留过滤("create_time").X条件("id", 3).X查询一条值(); err != nil {
			gtest.Error(err)
		} else {
			t.Assert(value.String(), "2019-10-24 10:00:00")
		}
	})
}

func Test_TX_GetAll(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		tx, err := db.X事务开启(ctx)
		if err != nil {
			gtest.Error(err)
		}
		if result, err := tx.GetAll别名(fmt.Sprintf("SELECT * FROM %s WHERE id=?", table), 1); err != nil {
			gtest.Error(err)
		} else {
			t.Assert(len(result), 1)
		}
		if err := tx.X事务提交(); err != nil {
			gtest.Error(err)
		}
	})
}

func Test_TX_GetOne(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		tx, err := db.X事务开启(ctx)
		if err != nil {
			gtest.Error(err)
		}
		if record, err := tx.X原生SQL查询单条记录(fmt.Sprintf("SELECT * FROM %s WHERE passport=?", table), "user_2"); err != nil {
			gtest.Error(err)
		} else {
			if record == nil {
				gtest.Error("FAIL")
			}
			t.Assert(record["nickname"].String(), "name_2")
		}
		if err := tx.X事务提交(); err != nil {
			gtest.Error(err)
		}
	})
}

func Test_TX_GetValue(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		tx, err := db.X事务开启(ctx)
		if err != nil {
			gtest.Error(err)
		}
		if value, err := tx.X原生SQL查询字段值(fmt.Sprintf("SELECT id FROM %s WHERE passport=?", table), "user_3"); err != nil {
			gtest.Error(err)
		} else {
			t.Assert(value.X取整数(), 3)
		}
		if err := tx.X事务提交(); err != nil {
			gtest.Error(err)
		}
	})
}

func Test_TX_GetCount(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		tx, err := db.X事务开启(ctx)
		if err != nil {
			gtest.Error(err)
		}
		if count, err := tx.X原生SQL查询字段计数("SELECT * FROM " + table); err != nil {
			gtest.Error(err)
		} else {
			t.Assert(count, int64(TableSize))
		}
		if err := tx.X事务提交(); err != nil {
			gtest.Error(err)
		}
	})
}

func Test_TX_GetStruct(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		tx, err := db.X事务开启(ctx)
		if err != nil {
			gtest.Error(err)
		}
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime gtime.Time
		}
		user := new(User)
		if err := tx.X原生SQL查询单条到结构体指针(user, fmt.Sprintf("SELECT * FROM %s WHERE id=?", table), 3); err != nil {
			gtest.Error(err)
		}
		t.Assert(user.NickName, "name_3")
		t.Assert(user.CreateTime.String(), "2018-10-24 10:00:00")
		if err := tx.X事务提交(); err != nil {
			gtest.Error(err)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		tx, err := db.X事务开启(ctx)
		if err != nil {
			gtest.Error(err)
		}
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *gtime.Time
		}
		user := new(User)
		if err := tx.X原生SQL查询单条到结构体指针(user, fmt.Sprintf("SELECT * FROM %s WHERE id=?", table), 3); err != nil {
			gtest.Error(err)
		}
		t.Assert(user.NickName, "name_3")
		t.Assert(user.CreateTime.String(), "2018-10-24 10:00:00")
		if err := tx.X事务提交(); err != nil {
			gtest.Error(err)
		}
	})
}

func Test_TX_GetStructs(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		tx, err := db.X事务开启(ctx)
		if err != nil {
			gtest.Error(err)
		}
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime gtime.Time
		}
		var users []User
		if err := tx.X原生SQL查询到结构体切片指针(&users, fmt.Sprintf("SELECT * FROM %s WHERE id>=?", table), 1); err != nil {
			gtest.Error(err)
		}
		t.Assert(len(users), TableSize)
		t.Assert(users[0].Id, 1)
		t.Assert(users[1].Id, 2)
		t.Assert(users[2].Id, 3)
		t.Assert(users[0].NickName, "name_1")
		t.Assert(users[1].NickName, "name_2")
		t.Assert(users[2].NickName, "name_3")
		t.Assert(users[2].CreateTime.String(), "2018-10-24 10:00:00")
		if err := tx.X事务提交(); err != nil {
			gtest.Error(err)
		}
	})

	gtest.C(t, func(t *gtest.T) {
		tx, err := db.X事务开启(ctx)
		if err != nil {
			gtest.Error(err)
		}
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *gtime.Time
		}
		var users []User
		if err := tx.X原生SQL查询到结构体切片指针(&users, fmt.Sprintf("SELECT * FROM %s WHERE id>=?", table), 1); err != nil {
			gtest.Error(err)
		}
		t.Assert(len(users), TableSize)
		t.Assert(users[0].Id, 1)
		t.Assert(users[1].Id, 2)
		t.Assert(users[2].Id, 3)
		t.Assert(users[0].NickName, "name_1")
		t.Assert(users[1].NickName, "name_2")
		t.Assert(users[2].NickName, "name_3")
		t.Assert(users[2].CreateTime.String(), "2018-10-24 10:00:00")
		if err := tx.X事务提交(); err != nil {
			gtest.Error(err)
		}
	})
}

func Test_TX_GetScan(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		tx, err := db.X事务开启(ctx)
		if err != nil {
			gtest.Error(err)
		}
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime gtime.Time
		}
		user := new(User)
		if err := tx.X原生SQL查询到结构体指针(user, fmt.Sprintf("SELECT * FROM %s WHERE id=?", table), 3); err != nil {
			gtest.Error(err)
		}
		t.Assert(user.NickName, "name_3")
		t.Assert(user.CreateTime.String(), "2018-10-24 10:00:00")
		if err := tx.X事务提交(); err != nil {
			gtest.Error(err)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		tx, err := db.X事务开启(ctx)
		if err != nil {
			gtest.Error(err)
		}
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *gtime.Time
		}
		user := new(User)
		if err := tx.X原生SQL查询到结构体指针(user, fmt.Sprintf("SELECT * FROM %s WHERE id=?", table), 3); err != nil {
			gtest.Error(err)
		}
		t.Assert(user.NickName, "name_3")
		t.Assert(user.CreateTime.String(), "2018-10-24 10:00:00")
		if err := tx.X事务提交(); err != nil {
			gtest.Error(err)
		}
	})

	gtest.C(t, func(t *gtest.T) {
		tx, err := db.X事务开启(ctx)
		if err != nil {
			gtest.Error(err)
		}
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime gtime.Time
		}
		var users []User
		if err := tx.X原生SQL查询到结构体指针(&users, fmt.Sprintf("SELECT * FROM %s WHERE id>=?", table), 1); err != nil {
			gtest.Error(err)
		}
		t.Assert(len(users), TableSize)
		t.Assert(users[0].Id, 1)
		t.Assert(users[1].Id, 2)
		t.Assert(users[2].Id, 3)
		t.Assert(users[0].NickName, "name_1")
		t.Assert(users[1].NickName, "name_2")
		t.Assert(users[2].NickName, "name_3")
		t.Assert(users[2].CreateTime.String(), "2018-10-24 10:00:00")
		if err := tx.X事务提交(); err != nil {
			gtest.Error(err)
		}
	})

	gtest.C(t, func(t *gtest.T) {
		tx, err := db.X事务开启(ctx)
		if err != nil {
			gtest.Error(err)
		}
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *gtime.Time
		}
		var users []User
		if err := tx.X原生SQL查询到结构体指针(&users, fmt.Sprintf("SELECT * FROM %s WHERE id>=?", table), 1); err != nil {
			gtest.Error(err)
		}
		t.Assert(len(users), TableSize)
		t.Assert(users[0].Id, 1)
		t.Assert(users[1].Id, 2)
		t.Assert(users[2].Id, 3)
		t.Assert(users[0].NickName, "name_1")
		t.Assert(users[1].NickName, "name_2")
		t.Assert(users[2].NickName, "name_3")
		t.Assert(users[2].CreateTime.String(), "2018-10-24 10:00:00")
		if err := tx.X事务提交(); err != nil {
			gtest.Error(err)
		}
	})
}

func Test_TX_Delete(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		table := createInitTable()
		defer dropTable(table)
		tx, err := db.X事务开启(ctx)
		if err != nil {
			gtest.Error(err)
		}
		if _, err := tx.X删除(table, 1); err != nil {
			gtest.Error(err)
		}
		if err := tx.X事务提交(); err != nil {
			gtest.Error(err)
		}
		if n, err := db.X创建Model对象(table).X查询行数(); err != nil {
			gtest.Error(err)
		} else {
			t.Assert(n, int64(0))
		}

		t.Assert(tx.X是否已关闭(), true)
	})

	gtest.C(t, func(t *gtest.T) {
		table := createInitTable()
		defer dropTable(table)
		tx, err := db.X事务开启(ctx)
		if err != nil {
			gtest.Error(err)
		}
		if _, err := tx.X删除(table, 1); err != nil {
			gtest.Error(err)
		}
		if n, err := tx.X创建Model对象(table).X查询行数(); err != nil {
			gtest.Error(err)
		} else {
			t.Assert(n, int64(0))
		}
		if err := tx.X事务回滚(); err != nil {
			gtest.Error(err)
		}
		if n, err := db.X创建Model对象(table).X查询行数(); err != nil {
			gtest.Error(err)
		} else {
			t.Assert(n, int64(TableSize))
			t.AssertNE(n, int64(0))
		}

		t.Assert(tx.X是否已关闭(), true)
	})
}

func Test_Transaction(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		ctx := context.TODO()
		err := db.X事务(ctx, func(ctx context.Context, tx gdb.TX) error {
			if _, err := tx.X设置上下文并取副本(ctx).X插入并替换已存在(table, g.Map{
				"id":          1,
				"passport":    "USER_1",
				"password":    "PASS_1",
				"nickname":    "NAME_1",
				"create_time": gtime.X创建并按当前时间().String(),
			}); err != nil {
				t.Error(err)
			}
			t.Assert(tx.X是否已关闭(), false)
			return gerror.X创建("error")
		})
		t.AssertNE(err, nil)

		if value, err := db.X创建Model对象(table).X设置上下文并取副本(ctx).X字段保留过滤("nickname").X条件("id", 1).X查询一条值(); err != nil {
			gtest.Error(err)
		} else {
			t.Assert(value.String(), "name_1")
		}
	})

	gtest.C(t, func(t *gtest.T) {
		ctx := context.TODO()
		err := db.X事务(ctx, func(ctx context.Context, tx gdb.TX) error {
			if _, err := tx.X插入并替换已存在(table, g.Map{
				"id":          1,
				"passport":    "USER_1",
				"password":    "PASS_1",
				"nickname":    "NAME_1",
				"create_time": gtime.X创建并按当前时间().String(),
			}); err != nil {
				t.Error(err)
			}
			return nil
		})
		t.AssertNil(err)

		if value, err := db.X创建Model对象(table).X字段保留过滤("nickname").X条件("id", 1).X查询一条值(); err != nil {
			gtest.Error(err)
		} else {
			t.Assert(value.String(), "NAME_1")
		}
	})
}

func Test_Transaction_Panic(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		ctx := context.TODO()
		err := db.X事务(ctx, func(ctx context.Context, tx gdb.TX) error {
			if _, err := tx.X插入并替换已存在(table, g.Map{
				"id":          1,
				"passport":    "USER_1",
				"password":    "PASS_1",
				"nickname":    "NAME_1",
				"create_time": gtime.X创建并按当前时间().String(),
			}); err != nil {
				t.Error(err)
			}
			panic("error")
			return nil
		})
		t.AssertNE(err, nil)

		if value, err := db.X创建Model对象(table).X字段保留过滤("nickname").X条件("id", 1).X查询一条值(); err != nil {
			gtest.Error(err)
		} else {
			t.Assert(value.String(), "name_1")
		}
	})
}

func Test_Transaction_Nested_Begin_Rollback_Commit(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		tx, err := db.X事务开启(ctx)
		t.AssertNil(err)
		// tx begin.
		err = tx.X事务开启()
		t.AssertNil(err)
		// tx rollback.
		_, err = tx.X创建Model对象(table).X设置数据(g.Map{
			"id":       1,
			"passport": "user_1",
			"password": "pass_1",
			"nickname": "name_1",
		}).X插入()
		err = tx.X事务回滚()
		t.AssertNil(err)
		// tx commit.
		_, err = tx.X创建Model对象(table).X设置数据(g.Map{
			"id":       2,
			"passport": "user_2",
			"password": "pass_2",
			"nickname": "name_2",
		}).X插入()
		err = tx.X事务提交()
		t.AssertNil(err)
		// check data.
		all, err := db.X创建Model对象(table).X查询()
		t.AssertNil(err)
		t.Assert(len(all), 1)
		t.Assert(all[0]["id"], 2)
	})
}

func Test_Transaction_Nested_TX_Transaction_UseTX(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	db.X设置调试模式(true)
	defer db.X设置调试模式(false)

	gtest.C(t, func(t *gtest.T) {
		var (
			err error
			ctx = context.TODO()
		)
		err = db.X事务(ctx, func(ctx context.Context, tx gdb.TX) error {
			// commit
			err = tx.X事务(ctx, func(ctx context.Context, tx gdb.TX) error {
				err = tx.X事务(ctx, func(ctx context.Context, tx gdb.TX) error {
					err = tx.X事务(ctx, func(ctx context.Context, tx gdb.TX) error {
						err = tx.X事务(ctx, func(ctx context.Context, tx gdb.TX) error {
							err = tx.X事务(ctx, func(ctx context.Context, tx gdb.TX) error {
								_, err = tx.X创建Model对象(table).X设置数据(g.Map{
									"id":          1,
									"passport":    "USER_1",
									"password":    "PASS_1",
									"nickname":    "NAME_1",
									"create_time": gtime.X创建并按当前时间().String(),
								}).X插入()
								t.AssertNil(err)
								return err
							})
							t.AssertNil(err)
							return err
						})
						t.AssertNil(err)
						return err
					})
					t.AssertNil(err)
					return err
				})
				t.AssertNil(err)
				return err
			})
			t.AssertNil(err)
			// rollback
			err = tx.X事务(ctx, func(ctx context.Context, tx gdb.TX) error {
				_, err = tx.X创建Model对象(table).X设置数据(g.Map{
					"id":          2,
					"passport":    "USER_2",
					"password":    "PASS_2",
					"nickname":    "NAME_2",
					"create_time": gtime.X创建并按当前时间().String(),
				}).X插入()
				t.AssertNil(err)
				panic("error")
				return err
			})
			t.AssertNE(err, nil)
			return nil
		})
		t.AssertNil(err)

		all, err := db.X设置上下文并取副本(ctx).X创建Model对象(table).X查询()
		t.AssertNil(err)
		t.Assert(len(all), 1)
		t.Assert(all[0]["id"], 1)

		// another record.
		err = db.X事务(ctx, func(ctx context.Context, tx gdb.TX) error {
			// commit
			err = tx.X事务(ctx, func(ctx context.Context, tx gdb.TX) error {
				err = tx.X事务(ctx, func(ctx context.Context, tx gdb.TX) error {
					err = tx.X事务(ctx, func(ctx context.Context, tx gdb.TX) error {
						err = tx.X事务(ctx, func(ctx context.Context, tx gdb.TX) error {
							err = tx.X事务(ctx, func(ctx context.Context, tx gdb.TX) error {
								_, err = tx.X创建Model对象(table).X设置数据(g.Map{
									"id":          3,
									"passport":    "USER_1",
									"password":    "PASS_1",
									"nickname":    "NAME_1",
									"create_time": gtime.X创建并按当前时间().String(),
								}).X插入()
								t.AssertNil(err)
								return err
							})
							t.AssertNil(err)
							return err
						})
						t.AssertNil(err)
						return err
					})
					t.AssertNil(err)
					return err
				})
				t.AssertNil(err)
				return err
			})
			t.AssertNil(err)
			// rollback
			err = tx.X事务(ctx, func(ctx context.Context, tx gdb.TX) error {
				_, err = tx.X创建Model对象(table).X设置数据(g.Map{
					"id":          4,
					"passport":    "USER_2",
					"password":    "PASS_2",
					"nickname":    "NAME_2",
					"create_time": gtime.X创建并按当前时间().String(),
				}).X插入()
				t.AssertNil(err)
				panic("error")
				return err
			})
			t.AssertNE(err, nil)
			return nil
		})
		t.AssertNil(err)

		all, err = db.X设置上下文并取副本(ctx).X创建Model对象(table).X查询()
		t.AssertNil(err)
		t.Assert(len(all), 2)
		t.Assert(all[0]["id"], 1)
		t.Assert(all[1]["id"], 3)
	})
}

func Test_Transaction_Nested_TX_Transaction_UseDB(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	// 将数据库设置为调试模式
	// 使用defer语句确保在函数返回前将数据库的调试模式重置为false
	// md5:b9225b2fca692b91

	gtest.C(t, func(t *gtest.T) {
		var (
			err error
			ctx = context.TODO()
		)
		err = db.X事务(ctx, func(ctx context.Context, tx gdb.TX) error {
			// commit
			err = db.X事务(ctx, func(ctx context.Context, tx gdb.TX) error {
				err = db.X事务(ctx, func(ctx context.Context, tx gdb.TX) error {
					err = db.X事务(ctx, func(ctx context.Context, tx gdb.TX) error {
						err = db.X事务(ctx, func(ctx context.Context, tx gdb.TX) error {
							err = db.X事务(ctx, func(ctx context.Context, tx gdb.TX) error {
								_, err = db.X创建Model对象(table).X设置上下文并取副本(ctx).X设置数据(g.Map{
									"id":          1,
									"passport":    "USER_1",
									"password":    "PASS_1",
									"nickname":    "NAME_1",
									"create_time": gtime.X创建并按当前时间().String(),
								}).X插入()
								t.AssertNil(err)
								return err
							})
							t.AssertNil(err)
							return err
						})
						t.AssertNil(err)
						return err
					})
					t.AssertNil(err)
					return err
				})
				t.AssertNil(err)
				return err
			})
			t.AssertNil(err)

			// rollback
			err = db.X事务(ctx, func(ctx context.Context, tx gdb.TX) error {
				_, err = tx.X创建Model对象(table).X设置上下文并取副本(ctx).X设置数据(g.Map{
					"id":          2,
					"passport":    "USER_2",
					"password":    "PASS_2",
					"nickname":    "NAME_2",
					"create_time": gtime.X创建并按当前时间().String(),
				}).X插入()
				t.AssertNil(err)
								// panic 会使这个事务回滚。 md5:f790e985b287f5dc
				panic("error")
				return err
			})
			t.AssertNE(err, nil)
			return nil
		})
		t.AssertNil(err)
		all, err := db.X创建Model对象(table).X查询()
		t.AssertNil(err)
		t.Assert(len(all), 1)
		t.Assert(all[0]["id"], 1)

		err = db.X事务(ctx, func(ctx context.Context, tx gdb.TX) error {
			// commit
			err = db.X事务(ctx, func(ctx context.Context, tx gdb.TX) error {
				err = db.X事务(ctx, func(ctx context.Context, tx gdb.TX) error {
					err = db.X事务(ctx, func(ctx context.Context, tx gdb.TX) error {
						err = db.X事务(ctx, func(ctx context.Context, tx gdb.TX) error {
							err = db.X事务(ctx, func(ctx context.Context, tx gdb.TX) error {
								_, err = db.X创建Model对象(table).X设置上下文并取副本(ctx).X设置数据(g.Map{
									"id":          3,
									"passport":    "USER_1",
									"password":    "PASS_1",
									"nickname":    "NAME_1",
									"create_time": gtime.X创建并按当前时间().String(),
								}).X插入()
								t.AssertNil(err)
								return err
							})
							t.AssertNil(err)
							return err
						})
						t.AssertNil(err)
						return err
					})
					t.AssertNil(err)
					return err
				})
				t.AssertNil(err)
				return err
			})
			t.AssertNil(err)

			// rollback
			err = db.X事务(ctx, func(ctx context.Context, tx gdb.TX) error {
				_, err = tx.X创建Model对象(table).X设置上下文并取副本(ctx).X设置数据(g.Map{
					"id":          4,
					"passport":    "USER_2",
					"password":    "PASS_2",
					"nickname":    "NAME_2",
					"create_time": gtime.X创建并按当前时间().String(),
				}).X插入()
				t.AssertNil(err)
								// panic 会使这个事务回滚。 md5:f790e985b287f5dc
				panic("error")
				return err
			})
			t.AssertNE(err, nil)
			return nil
		})
		t.AssertNil(err)

		all, err = db.X创建Model对象(table).X查询()
		t.AssertNil(err)
		t.Assert(len(all), 2)
		t.Assert(all[0]["id"], 1)
		t.Assert(all[1]["id"], 3)
	})
}

func Test_Transaction_Nested_SavePoint_RollbackTo(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		tx, err := db.X事务开启(ctx)
		t.AssertNil(err)
		// tx save point.
		_, err = tx.X创建Model对象(table).X设置数据(g.Map{
			"id":       1,
			"passport": "user_1",
			"password": "pass_1",
			"nickname": "name_1",
		}).X插入()
		err = tx.X保存事务点("MyPoint")
		t.AssertNil(err)
		_, err = tx.X创建Model对象(table).X设置数据(g.Map{
			"id":       2,
			"passport": "user_2",
			"password": "pass_2",
			"nickname": "name_2",
		}).X插入()
		// tx rollback to.
		err = tx.X回滚事务点("MyPoint")
		t.AssertNil(err)
		// tx commit.
		err = tx.X事务提交()
		t.AssertNil(err)

		// check data.
		all, err := db.X创建Model对象(table).X查询()
		t.AssertNil(err)
		t.Assert(len(all), 1)
		t.Assert(all[0]["id"], 1)
	})
}

func Test_Transaction_Method(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		var err error
		err = db.X事务(gctx.X创建(), func(ctx context.Context, tx gdb.TX) error {
			_, err = db.X创建Model对象(table).X设置上下文并取副本(ctx).X设置数据(g.Map{
				"id":          1,
				"passport":    "t1",
				"password":    "25d55ad283aa400af464c76d713c07ad",
				"nickname":    "T1",
				"create_time": gtime.X创建并按当前时间().String(),
			}).X插入()
			t.AssertNil(err)

			_, err = db.X设置上下文并取副本(ctx).X原生SQL执行(ctx, fmt.Sprintf(
				"insert into %s(`passport`,`password`,`nickname`,`create_time`,`id`) "+
					"VALUES('t2','25d55ad283aa400af464c76d713c07ad','T2','2021-08-25 21:53:00',2) ",
				table))
			t.AssertNil(err)
			return gerror.X创建("rollback")
		})
		t.AssertNE(err, nil)

		count, err := db.X创建Model对象(table).X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(0))
	})
}
