//go:build 屏蔽单元测试

// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package clickhouse_test

import (
	"fmt"
	"testing"

	"github.com/888go/goframe/frame/g"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
)

func Test_DB_Ping(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		err1 := db.X向主节点发送心跳()
		err2 := db.X向从节点发送心跳()
		t.Assert(err1, nil)
		t.Assert(err2, nil)
	})
}

func Test_DB_Query(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		_, err := db.X原生SQL查询(ctx, "SELECT ?", 1)
		t.AssertNil(err)

		_, err = db.X原生SQL查询(ctx, "SELECT ?+?", 1, 2)
		t.AssertNil(err)

		_, err = db.X原生SQL查询(ctx, "SELECT ?+?", g.Slice别名{1, 2})
		t.AssertNil(err)

		_, err = db.X原生SQL查询(ctx, "ERROR")
		t.AssertNE(err, nil)
	})
}

func Test_DB_Exec(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		_, err := db.X原生SQL执行(ctx, fmt.Sprintf("select * from %s ", table))
		t.AssertNil(err)
	})
}

func Test_DB_Insert(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		_, err := db.X插入(ctx, table, g.Map{
			"id":          uint64(1),
			"passport":    "t1",
			"password":    "25d55ad283aa400af464c76d713c07ad",
			"nickname":    "T1",
			"create_time": gtime.X创建并按当前时间(),
		})
		t.AssertNil(err)
		answer, err := db.GetAll别名(ctx, fmt.Sprintf("SELECT * FROM %s WHERE id=?", table), 1)
		t.AssertNil(err)
		t.Assert(len(answer), 1)
		t.Assert(answer[0]["passport"], "t1")
		t.Assert(answer[0]["password"], "25d55ad283aa400af464c76d713c07ad")
		t.Assert(answer[0]["nickname"], "T1")

		// normal map
		_, err = db.X插入(ctx, table, g.Map{
			"id":          uint64(2),
			"passport":    "t2",
			"password":    "25d55ad283aa400af464c76d713c07ad",
			"nickname":    "name_2",
			"create_time": gtime.X创建并按当前时间(),
		})
		t.AssertNil(err)

		answer, err = db.GetAll别名(ctx, fmt.Sprintf("SELECT * FROM %s WHERE id=?", table), 2)
		t.AssertNil(err)
		t.Assert(len(answer), 1)
		t.Assert(answer[0]["passport"], "t2")
		t.Assert(answer[0]["password"], "25d55ad283aa400af464c76d713c07ad")
		t.Assert(answer[0]["nickname"], "name_2")
	})
}

func Test_DB_Save(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		createTable("t_user")
		defer dropTable("t_user")

		i := 10
		data := g.Map{
			"id":          i,
			"passport":    fmt.Sprintf(`t%d`, i),
			"password":    fmt.Sprintf(`p%d`, i),
			"nickname":    fmt.Sprintf(`T%d`, i),
			"create_time": gtime.X创建并按当前时间(),
		}
		_, err := db.X插入并更新已存在(ctx, "t_user", data, 10)
		gtest.AssertNE(err, nil)
	})
}

func Test_DB_Replace(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		createTable("t_user")
		defer dropTable("t_user")

		i := 10
		data := g.Map{
			"id":          i,
			"passport":    fmt.Sprintf(`t%d`, i),
			"password":    fmt.Sprintf(`p%d`, i),
			"nickname":    fmt.Sprintf(`T%d`, i),
			"create_time": gtime.X创建并按当前时间(),
		}
		_, err := db.X插入并替换已存在(ctx, "t_user", data, 10)
		gtest.AssertNE(err, nil)
	})
}

func Test_DB_GetAll(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.GetAll别名(ctx, fmt.Sprintf("SELECT * FROM %s WHERE id=?", table), 1)
		t.AssertNil(err)
		t.Assert(len(result), 1)
		t.Assert(result[0]["id"].X取整数(), 1)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.GetAll别名(ctx, fmt.Sprintf("SELECT * FROM %s WHERE id=?", table), g.Slice别名{1})
		t.AssertNil(err)
		t.Assert(len(result), 1)
		t.Assert(result[0]["id"].X取整数(), 1)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.GetAll别名(ctx, fmt.Sprintf("SELECT * FROM %s WHERE id in(?)", table), g.Slice别名{1, 2, 3})
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["id"].X取整数(), 1)
		t.Assert(result[1]["id"].X取整数(), 2)
		t.Assert(result[2]["id"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.GetAll别名(ctx, fmt.Sprintf("SELECT * FROM %s WHERE id in(?,?,?)", table), g.Slice别名{1, 2, 3})
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["id"].X取整数(), 1)
		t.Assert(result[1]["id"].X取整数(), 2)
		t.Assert(result[2]["id"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.GetAll别名(ctx, fmt.Sprintf("SELECT * FROM %s WHERE id in(?,?,?)", table), g.Slice别名{1, 2, 3}...)
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["id"].X取整数(), 1)
		t.Assert(result[1]["id"].X取整数(), 2)
		t.Assert(result[2]["id"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.GetAll别名(ctx, fmt.Sprintf("SELECT * FROM %s WHERE id>=? AND id <=?", table), g.Slice别名{1, 3})
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["id"].X取整数(), 1)
		t.Assert(result[1]["id"].X取整数(), 2)
		t.Assert(result[2]["id"].X取整数(), 3)
	})
}

func Test_DB_GetOne(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         uint64
			Passport   string
			Password   string
			Nickname   string
			CreateTime *gtime.Time
		}
		data := User{
			Id:         uint64(1),
			Passport:   "user_1",
			Password:   "pass_1",
			Nickname:   "name_1",
			CreateTime: gtime.X创建并按当前时间(),
		}
		_, err := db.X插入(ctx, table, data)
		t.AssertNil(err)

		one, err := db.X原生SQL查询单条记录(ctx, fmt.Sprintf("SELECT * FROM %s WHERE id=?", table), 1)
		t.AssertNil(err)
		t.Assert(one["passport"], data.Passport)
		t.Assert(one["create_time"], data.CreateTime)
		t.Assert(one["nickname"], data.Nickname)
	})
}

func Test_DB_GetValue(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		value, err := db.X原生SQL查询字段值(ctx, fmt.Sprintf("SELECT id FROM %s WHERE passport=?", table), "user_3")
		t.AssertNil(err)
		t.Assert(value.X取整数(), 3)
	})
}

func Test_DB_GetCount(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		count, err := db.X原生SQL查询字段计数(ctx, fmt.Sprintf("SELECT * FROM %s", table))
		t.AssertNil(err)
		t.Assert(count, TableSize)
	})
}

func Test_DB_GetArray(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		array, err := db.X原生SQL查询切片(ctx, fmt.Sprintf("SELECT password FROM %s", table))
		t.AssertNil(err)
		arrays := make([]string, 0)
		for i := 1; i <= TableSize; i++ {
			arrays = append(arrays, fmt.Sprintf(`pass_%d`, i))
		}
		t.Assert(array, arrays)
	})
}

func Test_DB_GetScan(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime gtime.Time
		}
		user := new(User)
		err := db.X原生SQL查询到结构体指针(ctx, user, fmt.Sprintf("SELECT * FROM %s WHERE id=?", table), 3)
		t.AssertNil(err)
		t.Assert(user.NickName, "name_3")
	})
}

func Test_DB_Update(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		_, err := db.X更新(ctx, table, "password='123456'", "id=3")
		t.AssertNE(err, nil)

		one, err := db.X创建Model对象(table).X条件("id", 3).X查询一条()
		t.AssertNil(err)
		t.AssertNE(one["password"].String(), "123456")

		t.Assert(one["id"].X取整数(), 3)
		t.Assert(one["passport"].String(), "user_3")
		t.Assert(one["nickname"].String(), "name_3")
	})
}

func Test_DB_Delete(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		_, err := db.X删除(ctx, table, "id>3")
		t.AssertNE(err, nil)

	})
}

func Test_DB_Tables(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		tables := []string{"t_user1", "pop", "haha"}

		for _, v := range tables {
			createTable(v)
		}

		result, err := db.X取表名称切片(ctx)
		gtest.Assert(err, nil)

		for i := 0; i < len(tables); i++ {
			find := false
			for j := 0; j < len(result); j++ {
				if tables[i] == result[j] {
					find = true
					break
				}
			}
			gtest.AssertEQ(find, true)
		}
	})
}

func Test_DB_TableFields(t *testing.T) {
	table := createInitTable("user")
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		field, err := db.X取表字段信息Map(ctx, "user")
		gtest.AssertNil(err)
		gtest.AssertEQ(len(field), 5)
		gtest.AssertNQ(field, nil)
	})
}
