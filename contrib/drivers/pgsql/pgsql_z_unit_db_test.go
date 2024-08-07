//go:build 屏蔽单元测试

// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package pgsql_test

import (
	"fmt"
	"testing"

	"github.com/888go/goframe/frame/g"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
)

func Test_DB_Query(t *testing.T) {
	table := createTable("name")
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		_, err := db.X原生SQL查询(ctx, fmt.Sprintf("select * from %s ", table))
		t.AssertNil(err)
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
			"id":          1,
			"passport":    "t1",
			"password":    "25d55ad283aa400af464c76d713c07ad",
			"nickname":    "T1",
			"create_time": gtime.X创建并按当前时间().String(),
		})
		t.AssertNil(err)
		answer, err := db.GetAll别名(ctx, fmt.Sprintf("SELECT * FROM %s WHERE id=?", table), 1)
		t.AssertNil(err)
		t.Assert(len(answer), 1)
		t.Assert(answer[0]["passport"], "t1")
		t.Assert(answer[0]["password"], "25d55ad283aa400af464c76d713c07ad")
		t.Assert(answer[0]["nickname"], "T1")

		// normal map
		result, err := db.X插入(ctx, table, g.Map{
			"id":          "2",
			"passport":    "t2",
			"password":    "25d55ad283aa400af464c76d713c07ad",
			"nickname":    "name_2",
			"create_time": gtime.X创建并按当前时间().String(),
		})
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
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
			"create_time": gtime.X创建并按当前时间().String(),
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
			"create_time": gtime.X创建并按当前时间().String(),
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
			Id         int
			Passport   string
			Password   string
			Nickname   string
			CreateTime string
		}
		data := User{
			Id:         1,
			Passport:   "user_1",
			Password:   "pass_1",
			Nickname:   "name_1",
			CreateTime: "2020-10-10 12:00:01",
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
		result, err := db.X更新(ctx, table, "password='987654321'", "id=3")
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)

		one, err := db.X创建Model对象(table).X条件("id", 3).X查询一条()
		t.AssertNil(err)
		t.Assert(one["id"].X取整数(), 3)
		t.Assert(one["passport"].String(), "user_3")
		t.Assert(one["password"].String(), "987654321")
		t.Assert(one["nickname"].String(), "name_3")
	})
}

func Test_DB_Delete(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X删除(ctx, table, "id>3")
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 7)
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
	gtest.C(t, func(t *gtest.T) {
		table := createTable()
		defer dropTable(table)

		var expect = map[string][]interface{}{
			// []string: 索引类型 为空情况 键 默认值 注释
			// id 是 bigserial，所以默认值是一个 pgsql 函数
			// md5:42642899667c8094
			"id":          {0, "int8", false, "pri", fmt.Sprintf("nextval('%s_id_seq'::regclass)", table), ""},
			"passport":    {1, "varchar", false, "", nil, ""},
			"password":    {2, "varchar", false, "", nil, ""},
			"nickname":    {3, "varchar", false, "", nil, ""},
			"create_time": {4, "timestamp", false, "", nil, ""},
		}

		res, err := db.X取表字段信息Map(ctx, table)
		gtest.Assert(err, nil)

		for k, v := range expect {
			_, ok := res[k]
			gtest.AssertEQ(ok, true)

			gtest.AssertEQ(res[k].Index, v[0])
			gtest.AssertEQ(res[k].X名称, k)
			gtest.AssertEQ(res[k].X类型, v[1])
			gtest.AssertEQ(res[k].Null, v[2])
			gtest.AssertEQ(res[k].Key, v[3])
			gtest.AssertEQ(res[k].Default, v[4])
			gtest.AssertEQ(res[k].Comment, v[5])
		}
	})
}
