//go:build 屏蔽单元测试

// 版权所有 2019 gf 作者（https://github.com/gogf/gf）。保留所有权利。
//
// 此源代码形式受麻省理工学院（MIT）许可证的条款约束。
// 如果未随此文件一起分发MIT许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
// md5:47e609239e0cb2bc

package mssql_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	gjson "github.com/888go/goframe/encoding/gjson"
	gxml "github.com/888go/goframe/encoding/gxml"
	"github.com/888go/goframe/frame/g"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
)

func TestTables(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		tables := []string{"t_user1", "pop", "haha"}

		for _, v := range tables {
			createTable(v)
		}

		result, err := db.X取表名称切片(context.Background())
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

		result, err = db.X取表名称切片(context.Background(), "test")
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

		for _, v := range tables {
			dropTable(v)
		}
	})
}

func TestTableFields(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		createTable("t_user")
		defer dropTable("t_user")
		var expect = map[string][]interface{}{
			"ID":          {"numeric(10,0)", false, "PRI", "", "", ""},
			"PASSPORT":    {"varchar(45)", true, "", "", "", ""},
			"PASSWORD":    {"varchar(32)", true, "", "", "", ""},
			"NICKNAME":    {"varchar(45)", true, "", "", "", ""},
			"CREATE_TIME": {"datetime", true, "", "", "", ""},
		}

		res, err := db.X取表字段信息Map(context.Background(), "t_user")
		gtest.Assert(err, nil)

		for k, v := range expect {
			_, ok := res[k]
			gtest.AssertEQ(ok, true)
			gtest.AssertEQ(res[k].X名称, k)
			gtest.AssertEQ(res[k].X类型, v[0])
			gtest.AssertEQ(res[k].Null, v[1])
			gtest.AssertEQ(res[k].Key, v[2])
			gtest.AssertEQ(res[k].Default, v[3])
			gtest.AssertEQ(res[k].X额外, v[4])
			gtest.AssertEQ(res[k].Comment, v[5])
		}

		res, err = db.X取表字段信息Map(context.Background(), "t_user", "test")
		gtest.Assert(err, nil)

		for k, v := range expect {
			_, ok := res[k]
			gtest.AssertEQ(ok, true)
			gtest.AssertEQ(res[k].X名称, k)
			gtest.AssertEQ(res[k].X类型, v[0])
			gtest.AssertEQ(res[k].Null, v[1])
			gtest.AssertEQ(res[k].Key, v[2])
			gtest.AssertEQ(res[k].Default, v[3])
			gtest.AssertEQ(res[k].X额外, v[4])
			gtest.AssertEQ(res[k].Comment, v[5])
		}
	})

	gtest.C(t, func(t *gtest.T) {
		_, err := db.X取表字段信息Map(context.Background(), "t_user t_user2")
		gtest.AssertNE(err, nil)
	})
}

func TestDoInsert(t *testing.T) {
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
		_, err := db.X插入(context.Background(), "t_user", data)
		gtest.Assert(err, nil)

	})

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
		_, err := db.X插入并更新已存在(context.Background(), "t_user", data, 10)
		gtest.AssertNE(err, nil)

		_, err = db.X插入并替换已存在(context.Background(), "t_user", data, 10)
		gtest.AssertNE(err, nil)
	})
}

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
	gtest.C(t, func(t *gtest.T) {
		_, err := db.X原生SQL执行(ctx, "SELECT ?", 1)
		t.AssertNil(err)

		_, err = db.X原生SQL执行(ctx, "ERROR")
		t.AssertNE(err, nil)
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
			"create_time": gtime.X创建并按当前时间(),
		})
		t.AssertNil(err)

		// normal map
		result, err := db.X插入(ctx, table, g.Map{
			"id":          "2",
			"passport":    "t2",
			"password":    "25d55ad283aa400af464c76d713c07ad",
			"nickname":    "name_2",
			"create_time": gtime.X创建并按当前时间(),
		})
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)

		// struct
		type User struct {
			Id         int         `gconv:"id"`
			Passport   string      `json:"passport"`
			Password   string      `gconv:"password"`
			Nickname   string      `gconv:"nickname"`
			CreateTime *gtime.Time `json:"create_time"`
		}
		timeNow := gtime.X创建("2024-10-01 12:01:01")
		result, err = db.X插入(ctx, table, User{
			Id:         3,
			Passport:   "user_3",
			Password:   "25d55ad283aa400af464c76d713c07ad",
			Nickname:   "name_3",
			CreateTime: timeNow,
		})
		t.AssertNil(err)
		n, _ = result.RowsAffected()
		t.Assert(n, 1)

		one, err := db.X创建Model对象(table).X条件("id", 3).X查询一条()
		t.AssertNil(err)
		fmt.Println(one)
		t.Assert(one["ID"].X取整数(), 3)
		t.Assert(one["PASSPORT"].String(), "user_3")
		t.Assert(one["PASSWORD"].String(), "25d55ad283aa400af464c76d713c07ad")
		t.Assert(one["NICKNAME"].String(), "name_3")
		t.AssertNE(one["CREATE_TIME"].X取gtime时间类(), nil)
		t.AssertLT(timeNow.X取纳秒时长(one["CREATE_TIME"].X取gtime时间类()), 3)

		// *struct
		timeNow = gtime.X创建并按当前时间()
		result, err = db.X插入(ctx, table, &User{
			Id:         4,
			Passport:   "t4",
			Password:   "25d55ad283aa400af464c76d713c07ad",
			Nickname:   "name_4",
			CreateTime: timeNow,
		})
		t.AssertNil(err)
		n, _ = result.RowsAffected()
		t.Assert(n, 1)

		one, err = db.X创建Model对象(table).X条件("id", 4).X查询一条()
		t.AssertNil(err)
		t.Assert(one["ID"].X取整数(), 4)
		t.Assert(one["PASSPORT"].String(), "t4")
		t.Assert(one["PASSWORD"].String(), "25d55ad283aa400af464c76d713c07ad")
		t.Assert(one["NICKNAME"].String(), "name_4")

		// batch with Insert
		timeNow = gtime.X创建并按当前时间()
		r, err := db.X插入(ctx, table, g.Slice别名{
			g.Map{
				"id":          200,
				"passport":    "t200",
				"password":    "25d55ad283aa400af464c76d71qw07ad",
				"nickname":    "T200",
				"create_time": timeNow,
			},
			g.Map{
				"id":          300,
				"passport":    "t300",
				"password":    "25d55ad283aa400af464c76d713c07ad",
				"nickname":    "T300",
				"create_time": timeNow,
			},
		})
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 2)

		one, err = db.X创建Model对象(table).X条件("id", 200).X查询一条()
		t.AssertNil(err)
		t.Assert(one["ID"].X取整数(), 200)
		t.Assert(one["PASSPORT"].String(), "t200")
		t.Assert(one["PASSWORD"].String(), "25d55ad283aa400af464c76d71qw07ad")
		t.Assert(one["NICKNAME"].String(), "T200")
	})
}

func Test_DB_Insert_KeyFieldNameMapping(t *testing.T) {
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
		t.Assert(one["PASSPORT"], data.Passport)
		t.Assert(one["CREATE_TIME"], data.CreateTime)
		t.Assert(one["NICKNAME"], data.Nickname)
	})
}

func Test_DB_Update_KeyFieldNameMapping(t *testing.T) {
	table := createInitTable()
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
			Passport:   "user_10",
			Password:   "pass_10",
			Nickname:   "name_10",
			CreateTime: "2020-10-10 12:00:01",
		}
		_, err := db.X更新(ctx, table, data, "id=1")
		t.AssertNil(err)

		one, err := db.X原生SQL查询单条记录(ctx, fmt.Sprintf("SELECT * FROM %s WHERE id=?", table), 1)
		t.AssertNil(err)
		t.Assert(one["PASSPORT"], data.Passport)
		t.Assert(one["CREATE_TIME"], data.CreateTime)
		t.Assert(one["NICKNAME"], data.Nickname)
	})
}

func Test_DB_BatchInsert(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		table := createTable()
		defer dropTable(table)
		r, err := db.X插入(ctx, table, g.Map切片{
			{
				"id":          2,
				"passport":    "t2",
				"password":    "25d55ad283aa400af464c76d713c07ad",
				"nickname":    "name_2",
				"create_time": gtime.X创建并按当前时间(),
			},
			{
				"id":          3,
				"passport":    "user_3",
				"password":    "25d55ad283aa400af464c76d713c07ad",
				"nickname":    "name_3",
				"create_time": gtime.X创建并按当前时间(),
			},
		}, 1)
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 2)

	})

	gtest.C(t, func(t *gtest.T) {
		table := createTable()
		defer dropTable(table)
		// []interface{}
		r, err := db.X插入(ctx, table, g.Slice别名{
			g.Map{
				"id":          2,
				"passport":    "t2",
				"password":    "25d55ad283aa400af464c76d713c07ad",
				"nickname":    "name_2",
				"create_time": gtime.X创建并按当前时间(),
			},
			g.Map{
				"id":          3,
				"passport":    "user_3",
				"password":    "25d55ad283aa400af464c76d713c07ad",
				"nickname":    "name_3",
				"create_time": gtime.X创建并按当前时间(),
			},
		}, 1)
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 2)
	})

	// batch insert map
	gtest.C(t, func(t *gtest.T) {
		table := createTable()
		defer dropTable(table)
		result, err := db.X插入(ctx, table, g.Map{
			"id":          1,
			"passport":    "t1",
			"password":    "p1",
			"nickname":    "T1",
			"create_time": gtime.X创建并按当前时间(),
		})
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
	})
}

func Test_DB_BatchInsert_Struct(t *testing.T) {
	// batch insert struct
	gtest.C(t, func(t *gtest.T) {
		table := createTable()
		defer dropTable(table)

		type User struct {
			Id         int         `c:"id"`
			Passport   string      `c:"passport"`
			Password   string      `c:"password"`
			NickName   string      `c:"nickname"`
			CreateTime *gtime.Time `c:"create_time"`
		}
		user := &User{
			Id:         1,
			Passport:   "t1",
			Password:   "p1",
			NickName:   "T1",
			CreateTime: gtime.X创建并按当前时间(),
		}
		result, err := db.X插入(ctx, table, user)
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
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
		t.Assert(one["ID"].X取整数(), 3)
		t.Assert(one["PASSPORT"].String(), "user_3")
		t.Assert(one["PASSWORD"].String(), "987654321")
		t.Assert(one["NICKNAME"].String(), "name_3")
	})
}

func Test_DB_GetAll(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.GetAll别名(ctx, fmt.Sprintf("SELECT * FROM %s WHERE id=?", table), 1)
		t.AssertNil(err)
		t.Assert(len(result), 1)
		t.Assert(result[0]["ID"].X取整数(), 1)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.GetAll别名(ctx, fmt.Sprintf("SELECT * FROM %s WHERE id=?", table), g.Slice别名{1})
		t.AssertNil(err)
		t.Assert(len(result), 1)
		t.Assert(result[0]["ID"].X取整数(), 1)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.GetAll别名(ctx, fmt.Sprintf("SELECT * FROM %s WHERE id in(?)", table), g.Slice别名{1, 2, 3})
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["ID"].X取整数(), 1)
		t.Assert(result[1]["ID"].X取整数(), 2)
		t.Assert(result[2]["ID"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.GetAll别名(ctx, fmt.Sprintf("SELECT * FROM %s WHERE id in(?,?,?)", table), g.Slice别名{1, 2, 3})
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["ID"].X取整数(), 1)
		t.Assert(result[1]["ID"].X取整数(), 2)
		t.Assert(result[2]["ID"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.GetAll别名(ctx, fmt.Sprintf("SELECT * FROM %s WHERE id in(?,?,?)", table), g.Slice别名{1, 2, 3}...)
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["ID"].X取整数(), 1)
		t.Assert(result[1]["ID"].X取整数(), 2)
		t.Assert(result[2]["ID"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.GetAll别名(ctx, fmt.Sprintf("SELECT * FROM %s WHERE id>=? AND id <=?", table), g.Slice别名{1, 3})
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["ID"].X取整数(), 1)
		t.Assert(result[1]["ID"].X取整数(), 2)
		t.Assert(result[2]["ID"].X取整数(), 3)
	})
}

func Test_DB_GetOne(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		record, err := db.X原生SQL查询单条记录(ctx, fmt.Sprintf("SELECT * FROM %s WHERE passport=?", table), "user_1")
		t.AssertNil(err)
		t.Assert(record["NICKNAME"].String(), "name_1")
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

func Test_DB_GetStruct(t *testing.T) {
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
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *gtime.Time
		}
		user := new(User)
		err := db.X原生SQL查询到结构体指针(ctx, user, fmt.Sprintf("SELECT * FROM %s WHERE id=?", table), 3)
		t.AssertNil(err)
		t.Assert(user.NickName, "name_3")
	})
}

func Test_DB_GetStructs(t *testing.T) {
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
		var users []User
		err := db.X原生SQL查询到结构体指针(ctx, &users, fmt.Sprintf("SELECT * FROM %s WHERE id>?", table), 1)
		t.AssertNil(err)
		t.Assert(len(users), TableSize-1)
		t.Assert(users[0].Id, 2)
		t.Assert(users[1].Id, 3)
		t.Assert(users[2].Id, 4)
		t.Assert(users[0].NickName, "name_2")
		t.Assert(users[1].NickName, "name_3")
		t.Assert(users[2].NickName, "name_4")
	})

	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *gtime.Time
		}
		var users []User
		err := db.X原生SQL查询到结构体指针(ctx, &users, fmt.Sprintf("SELECT * FROM %s WHERE id>?", table), 1)
		t.AssertNil(err)
		t.Assert(len(users), TableSize-1)
		t.Assert(users[0].Id, 2)
		t.Assert(users[1].Id, 3)
		t.Assert(users[2].Id, 4)
		t.Assert(users[0].NickName, "name_2")
		t.Assert(users[1].NickName, "name_3")
		t.Assert(users[2].NickName, "name_4")
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
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime gtime.Time
		}
		var user *User
		err := db.X原生SQL查询到结构体指针(ctx, &user, fmt.Sprintf("SELECT * FROM %s WHERE id=?", table), 3)
		t.AssertNil(err)
		t.Assert(user.NickName, "name_3")
	})
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *gtime.Time
		}
		user := new(User)
		err := db.X原生SQL查询到结构体指针(ctx, user, fmt.Sprintf("SELECT * FROM %s WHERE id=?", table), 3)
		t.AssertNil(err)
		t.Assert(user.NickName, "name_3")
	})

	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime gtime.Time
		}
		var users []User
		err := db.X原生SQL查询到结构体指针(ctx, &users, fmt.Sprintf("SELECT * FROM %s WHERE id>?", table), 1)
		t.AssertNil(err)
		t.Assert(len(users), TableSize-1)
		t.Assert(users[0].Id, 2)
		t.Assert(users[1].Id, 3)
		t.Assert(users[2].Id, 4)
		t.Assert(users[0].NickName, "name_2")
		t.Assert(users[1].NickName, "name_3")
		t.Assert(users[2].NickName, "name_4")
	})

	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *gtime.Time
		}
		var users []User
		err := db.X原生SQL查询到结构体指针(ctx, &users, fmt.Sprintf("SELECT * FROM %s WHERE id>?", table), 1)
		t.AssertNil(err)
		t.Assert(len(users), TableSize-1)
		t.Assert(users[0].Id, 2)
		t.Assert(users[1].Id, 3)
		t.Assert(users[2].Id, 4)
		t.Assert(users[0].NickName, "name_2")
		t.Assert(users[1].NickName, "name_3")
		t.Assert(users[2].NickName, "name_4")
	})
}

func Test_DB_Delete(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X删除(ctx, table, "1=1")
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, TableSize)
	})
}

func Test_DB_Time(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X插入(ctx, table, g.Map{
			"id":          200,
			"passport":    "t200",
			"password":    "123456",
			"nickname":    "T200",
			"create_time": time.Now(),
		})
		if err != nil {
			gtest.Error(err)
		}
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
		value, err := db.X原生SQL查询字段值(ctx, fmt.Sprintf("select passport from %s where id=?", table), 200)
		t.AssertNil(err)
		t.Assert(value.String(), "t200")
	})

	gtest.C(t, func(t *gtest.T) {
		t1 := time.Now()
		result, err := db.X插入(ctx, table, g.Map{
			"id":          300,
			"passport":    "t300",
			"password":    "123456",
			"nickname":    "T300",
			"create_time": &t1,
		})
		if err != nil {
			gtest.Error(err)
		}
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
		value, err := db.X原生SQL查询字段值(ctx, fmt.Sprintf("select passport from %s where id=?", table), 300)
		t.AssertNil(err)
		t.Assert(value.String(), "t300")
	})

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X删除(ctx, table, "1=1")
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 2)
	})
}

func Test_DB_ToJson(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	_, err := db.X更新(ctx, table, "create_time='2010-10-10 00:00:01'", "id=?", 1)
	gtest.AssertNil(err)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X字段保留过滤("*").X条件("id =? ", 1).X查询()
		if err != nil {
			gtest.Fatal(err)
		}

		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime string
		}

		users := make([]User, 0)

		err = result.X取切片结构体指针(users)
		t.AssertNE(err, nil)

		err = result.X取切片结构体指针(&users)
		if err != nil {
			gtest.Fatal(err)
		}

		// ToJson
		resultJson, err := gjson.X加载并自动识别格式(result.X取json())
		if err != nil {
			gtest.Fatal(err)
		}

		t.Assert(users[0].Id, resultJson.X取值("0.ID").X取整数())
		t.Assert(users[0].Passport, resultJson.X取值("0.PASSPORT").String())
		t.Assert(users[0].Password, resultJson.X取值("0.PASSWORD").String())
		t.Assert(users[0].NickName, resultJson.X取值("0.NICKNAME").String())
		t.Assert(users[0].CreateTime, resultJson.X取值("0.CREATE_TIME").String())

		result = nil
		t.Assert(result.X取切片结构体指针(&users), sql.ErrNoRows)
	})

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X字段保留过滤("*").X条件("id =? ", 1).X查询一条()
		if err != nil {
			gtest.Fatal(err)
		}

		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime string
		}

		users := User{}

		err = result.X取结构体指针(&users)
		if err != nil {
			gtest.Fatal(err)
		}

		result = nil
		err = result.X取结构体指针(&users)
		t.AssertNE(err, nil)
	})
}

func Test_DB_ToXml(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	_, err := db.X更新(ctx, table, "create_time='2010-10-10 00:00:01'", "id=?", 1)
	gtest.AssertNil(err)

	gtest.C(t, func(t *gtest.T) {
		record, err := db.X创建Model对象(table).X字段保留过滤("*").X条件("id = ?", 1).X查询一条()
		if err != nil {
			gtest.Fatal(err)
		}

		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime string
		}

		user := User{}
		err = record.X取结构体指针(&user)
		if err != nil {
			gtest.Fatal(err)
		}

		result, err := gxml.Decode([]byte(record.X取xml("doc")))
		if err != nil {
			gtest.Fatal(err)
		}

		resultXml := result["doc"].(map[string]interface{})
		if v, ok := resultXml["ID"]; ok {
			t.Assert(user.Id, v)
		} else {
			gtest.Fatal("FAIL")
		}

		if v, ok := resultXml["PASSPORT"]; ok {
			t.Assert(user.Passport, v)
		} else {
			gtest.Fatal("FAIL")
		}

		if v, ok := resultXml["PASSWORD"]; ok {
			t.Assert(user.Password, v)
		} else {
			gtest.Fatal("FAIL")
		}

		if v, ok := resultXml["NICKNAME"]; ok {
			t.Assert(user.NickName, v)
		} else {
			gtest.Fatal("FAIL")
		}

		if v, ok := resultXml["CREATE_TIME"]; ok {
			t.Assert(user.CreateTime, v)
		} else {
			gtest.Fatal("FAIL")
		}
	})
}

func Test_DB_ToStringMap(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	_, err := db.X更新(ctx, table, "create_time='2010-10-10 00:00:01'", "id=?", 1)
	gtest.AssertNil(err)
	gtest.C(t, func(t *gtest.T) {
		id := "1"
		result, err := db.X创建Model对象(table).X字段保留过滤("*").X条件("id = ?", 1).X查询()
		if err != nil {
			gtest.Fatal(err)
		}

		type t_user struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime string
		}

		t_users := make([]t_user, 0)
		err = result.X取切片结构体指针(&t_users)
		if err != nil {
			gtest.Fatal(err)
		}

		resultStringMap := result.X取字段MapStr("ID")
		t.Assert(t_users[0].Id, resultStringMap[id]["ID"])
		t.Assert(t_users[0].Passport, resultStringMap[id]["PASSPORT"])
		t.Assert(t_users[0].Password, resultStringMap[id]["PASSWORD"])
		t.Assert(t_users[0].NickName, resultStringMap[id]["NICKNAME"])
		t.Assert(t_users[0].CreateTime, resultStringMap[id]["CREATE_TIME"])
	})
}

func Test_DB_ToIntMap(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	_, err := db.X更新(ctx, table, "create_time='2010-10-10 00:00:01'", "id=?", 1)
	gtest.AssertNil(err)

	gtest.C(t, func(t *gtest.T) {
		id := 1
		result, err := db.X创建Model对象(table).X字段保留过滤("*").X条件("id = ?", id).X查询()
		if err != nil {
			gtest.Fatal(err)
		}

		type t_user struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime string
		}

		t_users := make([]t_user, 0)
		err = result.X取切片结构体指针(&t_users)
		if err != nil {
			gtest.Fatal(err)
		}

		resultIntMap := result.X取字段MapInt("ID")
		t.Assert(t_users[0].Id, resultIntMap[id]["ID"])
		t.Assert(t_users[0].Passport, resultIntMap[id]["PASSPORT"])
		t.Assert(t_users[0].Password, resultIntMap[id]["PASSWORD"])
		t.Assert(t_users[0].NickName, resultIntMap[id]["NICKNAME"])
		t.Assert(t_users[0].CreateTime, resultIntMap[id]["CREATE_TIME"])
	})
}

func Test_DB_ToUintMap(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	_, err := db.X更新(ctx, table, "create_time='2010-10-10 00:00:01'", "id=?", 1)
	gtest.AssertNil(err)

	gtest.C(t, func(t *gtest.T) {
		id := 1
		result, err := db.X创建Model对象(table).X字段保留过滤("*").X条件("id = ?", id).X查询()
		if err != nil {
			gtest.Fatal(err)
		}

		type t_user struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime string
		}

		t_users := make([]t_user, 0)
		err = result.X取切片结构体指针(&t_users)
		if err != nil {
			gtest.Fatal(err)
		}

		resultUintMap := result.X取字段MapUint("ID")
		t.Assert(t_users[0].Id, resultUintMap[uint(id)]["ID"])
		t.Assert(t_users[0].Passport, resultUintMap[uint(id)]["PASSPORT"])
		t.Assert(t_users[0].Password, resultUintMap[uint(id)]["PASSWORD"])
		t.Assert(t_users[0].NickName, resultUintMap[uint(id)]["NICKNAME"])
		t.Assert(t_users[0].CreateTime, resultUintMap[uint(id)]["CREATE_TIME"])
	})
}

func Test_DB_ToStringRecord(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	_, err := db.X更新(ctx, table, "create_time='2010-10-10 00:00:01'", "id=?", 1)
	gtest.AssertNil(err)

	gtest.C(t, func(t *gtest.T) {
		id := 1
		ids := "1"
		result, err := db.X创建Model对象(table).X字段保留过滤("*").X条件("id = ?", id).X查询()
		if err != nil {
			gtest.Fatal(err)
		}

		type t_user struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime string
		}

		t_users := make([]t_user, 0)
		err = result.X取切片结构体指针(&t_users)
		if err != nil {
			gtest.Fatal(err)
		}

		resultStringRecord := result.RecordKeyStr("ID")
		t.Assert(t_users[0].Id, resultStringRecord[ids]["ID"].X取整数())
		t.Assert(t_users[0].Passport, resultStringRecord[ids]["PASSPORT"].String())
		t.Assert(t_users[0].Password, resultStringRecord[ids]["PASSWORD"].String())
		t.Assert(t_users[0].NickName, resultStringRecord[ids]["NICKNAME"].String())
		t.Assert(t_users[0].CreateTime, resultStringRecord[ids]["CREATE_TIME"].String())
	})
}

func Test_DB_ToIntRecord(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	_, err := db.X更新(ctx, table, "create_time='2010-10-10 00:00:01'", "id=?", 1)
	gtest.AssertNil(err)

	gtest.C(t, func(t *gtest.T) {
		id := 1
		result, err := db.X创建Model对象(table).X字段保留过滤("*").X条件("id = ?", id).X查询()
		if err != nil {
			gtest.Fatal(err)
		}

		type t_user struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime string
		}

		t_users := make([]t_user, 0)
		err = result.X取切片结构体指针(&t_users)
		if err != nil {
			gtest.Fatal(err)
		}

		resultIntRecord := result.RecordKeyInt("ID")
		t.Assert(t_users[0].Id, resultIntRecord[id]["ID"].X取整数())
		t.Assert(t_users[0].Passport, resultIntRecord[id]["PASSPORT"].String())
		t.Assert(t_users[0].Password, resultIntRecord[id]["PASSWORD"].String())
		t.Assert(t_users[0].NickName, resultIntRecord[id]["NICKNAME"].String())
		t.Assert(t_users[0].CreateTime, resultIntRecord[id]["CREATE_TIME"].String())
	})
}

func Test_DB_ToUintRecord(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	_, err := db.X更新(ctx, table, "create_time='2010-10-10 00:00:01'", "id=?", 1)
	gtest.AssertNil(err)

	gtest.C(t, func(t *gtest.T) {
		id := 1
		result, err := db.X创建Model对象(table).X字段保留过滤("*").X条件("id = ?", id).X查询()
		if err != nil {
			gtest.Fatal(err)
		}

		type t_user struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime string
		}

		t_users := make([]t_user, 0)
		err = result.X取切片结构体指针(&t_users)
		if err != nil {
			gtest.Fatal(err)
		}

		resultUintRecord := result.RecordKeyUint("ID")
		t.Assert(t_users[0].Id, resultUintRecord[uint(id)]["ID"].X取整数())
		t.Assert(t_users[0].Passport, resultUintRecord[uint(id)]["PASSPORT"].String())
		t.Assert(t_users[0].Password, resultUintRecord[uint(id)]["PASSWORD"].String())
		t.Assert(t_users[0].NickName, resultUintRecord[uint(id)]["NICKNAME"].String())
		t.Assert(t_users[0].CreateTime, resultUintRecord[uint(id)]["CREATE_TIME"].String())
	})
}

func Test_Model_InnerJoin(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		table1 := createInitTable("user1")
		table2 := createInitTable("user2")

		defer dropTable(table1)
		defer dropTable(table2)

		res, err := db.X创建Model对象(table1).X条件("id > ?", 5).X删除()
		if err != nil {
			t.Fatal(err)
		}

		n, err := res.RowsAffected()
		if err != nil {
			t.Fatal(err)
		}

		t.Assert(n, 5)

		result, err := db.X创建Model对象(table1+" u1").X内连接(table2+" u2", "u1.id = u2.id").X排序("u1.id").X查询()
		if err != nil {
			t.Fatal(err)
		}

		t.Assert(len(result), 5)

		result, err = db.X创建Model对象(table1+" u1").X内连接(table2+" u2", "u1.id = u2.id").X条件("u1.id > ?", 1).X排序("u1.id").X查询()
		if err != nil {
			t.Fatal(err)
		}

		t.Assert(len(result), 4)
	})
}

func Test_Model_LeftJoin(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		table1 := createInitTable("user1")
		table2 := createInitTable("user2")

		defer dropTable(table1)
		defer dropTable(table2)

		res, err := db.X创建Model对象(table2).X条件("id > ?", 3).X删除()
		if err != nil {
			t.Fatal(err)
		}

		n, err := res.RowsAffected()
		if err != nil {
			t.Fatal(err)
		} else {
			t.Assert(n, 7)
		}

		result, err := db.X创建Model对象(table1+" u1").X左连接(table2+" u2", "u1.id = u2.id").X查询()
		if err != nil {
			t.Fatal(err)
		}

		t.Assert(len(result), 10)

		result, err = db.X创建Model对象(table1+" u1").X左连接(table2+" u2", "u1.id = u2.id").X条件("u1.id > ? ", 2).X查询()
		if err != nil {
			t.Fatal(err)
		}

		t.Assert(len(result), 8)
	})
}

func Test_Model_RightJoin(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		table1 := createInitTable("user1")
		table2 := createInitTable("user2")

		defer dropTable(table1)
		defer dropTable(table2)

		res, err := db.X创建Model对象(table1).X条件("id > ?", 3).X删除()
		if err != nil {
			t.Fatal(err)
		}

		n, err := res.RowsAffected()
		if err != nil {
			t.Fatal(err)
		}

		t.Assert(n, 7)

		result, err := db.X创建Model对象(table1+" u1").X右连接(table2+" u2", "u1.id = u2.id").X查询()
		if err != nil {
			t.Fatal(err)
		}
		t.Assert(len(result), 10)

		result, err = db.X创建Model对象(table1+" u1").X右连接(table2+" u2", "u1.id = u2.id").X条件("u1.id > 2").X查询()
		if err != nil {
			t.Fatal(err)
		}
		t.Assert(len(result), 1)
	})
}

func Test_Empty_Slice_Argument(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		result, err := db.GetAll别名(ctx, fmt.Sprintf(`select * from %s where id in(?)`, table), g.Slice别名{})
		t.AssertNil(err)
		t.Assert(len(result), 0)
	})
}
