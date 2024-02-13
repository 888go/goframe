// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package sqlite_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
	
	"github.com/888go/goframe/container/garray"
	"github.com/888go/goframe/database/gdb"
	"github.com/888go/goframe/encoding/gjson"
	"github.com/888go/goframe/encoding/gxml"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
)

func Test_New(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		node := db类.ConfigNode{
			Type:    "sqlite",
			Link:    文件类.X路径生成(dbDir, "test.db"),
			Charset: "utf8",
		}
		newDb, err := db类.X创建DB对象(node)
		t.AssertNil(err)
		value, err := newDb.X原生SQL查询字段值(ctx, `select 1`)
		t.AssertNil(err)
		t.Assert(value, `1`)
		t.AssertNil(newDb.X关闭数据库(ctx))
	})
}

func Test_New_Path_With_Colon(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {

		dbFilePathWithColon := 文件类.X路径生成(dbDir, "test_1") //2024-01-03 test:1改成test_1,win平台不支持此文件名
		if err := 文件类.X创建目录(dbFilePathWithColon); err != nil {
			单元测试类.Error(err)
		}
		node := db类.ConfigNode{
			Type:    "sqlite",
			Link:    fmt.Sprintf(`sqlite::@file(%s)`, 文件类.X路径生成(dbFilePathWithColon, "test.db")),
			Charset: "utf8",
		}
		newDb, err := db类.X创建DB对象(node)
		t.AssertNil(err)
		value, err := newDb.X原生SQL查询字段值(ctx, `select 1`)
		t.AssertNil(err)
		t.Assert(value, `1`)
		t.AssertNil(newDb.X关闭数据库(ctx))
	})
}

func Test_DB_Ping(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		err1 := db.X向主节点发送心跳()
		err2 := db.X向从节点发送心跳()
		t.Assert(err1, nil)
		t.Assert(err2, nil)
	})
}

func Test_DB_Query(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		_, err := db.X原生SQL查询(ctx, "SELECT ?", 1)
		t.AssertNil(err)

		_, err = db.X原生SQL查询(ctx, "SELECT ?+?", 1, 2)
		t.AssertNil(err)

		_, err = db.X原生SQL查询(ctx, "SELECT ?+?", g.Slice别名{1, 2})
		t.AssertNil(err)
	})
}

func Test_DB_Exec(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		_, err := db.X原生SQL执行(ctx, "SELECT ?", 1)
		t.AssertNil(err)
	})
}

func Test_DB_Prepare(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		st, err := db.X原生sql取参数预处理对象(ctx, "SELECT 100")
		t.AssertNil(err)

		rows, err := st.X查询()
		t.AssertNil(err)

		array, err := rows.Columns()
		t.AssertNil(err)
		t.Assert(array[0], "100")

		err = rows.Close()
		t.AssertNil(err)
	})
}

func Test_DB_Insert(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		_, err := db.X插入(ctx, table, g.Map{
			"id":          1,
			"passport":    "t1",
			"password":    "25d55ad283aa400af464c76d713c07ad",
			"nickname":    "T1",
			"create_time": 时间类.X创建并按当前时间().String(),
		})
		t.AssertNil(err)

		// normal map
		result, err := db.X插入(ctx, table, g.Map{
			"id":          "2",
			"passport":    "t2",
			"password":    "25d55ad283aa400af464c76d713c07ad",
			"nickname":    "name_2",
			"create_time": 时间类.X创建并按当前时间().String(),
		})
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)

		// struct
		type User struct {
			Id         int    `gconv:"id"`
			Passport   string `json:"passport"`
			Password   string `gconv:"password"`
			Nickname   string `gconv:"nickname"`
			CreateTime string `json:"create_time"`
		}
		timeStr := 时间类.X创建并按当前时间().String()
		result, err = db.X插入(ctx, table, User{
			Id:         3,
			Passport:   "user_3",
			Password:   "25d55ad283aa400af464c76d713c07ad",
			Nickname:   "name_3",
			CreateTime: timeStr,
		})
		t.AssertNil(err)
		n, _ = result.RowsAffected()
		t.Assert(n, 1)

		one, err := db.X创建Model对象(table).X条件("id", 3).X查询一条()
		t.AssertNil(err)

		t.Assert(one["id"].X取整数(), 3)
		t.Assert(one["passport"].String(), "user_3")
		t.Assert(one["password"].String(), "25d55ad283aa400af464c76d713c07ad")
		t.Assert(one["nickname"].String(), "name_3")
		t.Assert(one["create_time"].X取gtime时间类().String(), timeStr)

		// *struct
		timeStr = 时间类.X创建并按当前时间().String()
		result, err = db.X插入(ctx, table, &User{
			Id:         4,
			Passport:   "t4",
			Password:   "25d55ad283aa400af464c76d713c07ad",
			Nickname:   "name_4",
			CreateTime: timeStr,
		})
		t.AssertNil(err)
		n, _ = result.RowsAffected()
		t.Assert(n, 1)

		one, err = db.X创建Model对象(table).X条件("id", 4).X查询一条()
		t.AssertNil(err)
		t.Assert(one["id"].X取整数(), 4)
		t.Assert(one["passport"].String(), "t4")
		t.Assert(one["password"].String(), "25d55ad283aa400af464c76d713c07ad")
		t.Assert(one["nickname"].String(), "name_4")
		t.Assert(one["create_time"].X取gtime时间类().String(), timeStr)

		// batch with Insert
		timeStr = 时间类.X创建并按当前时间().String()
		r, err := db.X插入(ctx, table, g.Slice别名{
			g.Map{
				"id":          200,
				"passport":    "t200",
				"password":    "25d55ad283aa400af464c76d71qw07ad",
				"nickname":    "T200",
				"create_time": timeStr,
			},
			g.Map{
				"id":          300,
				"passport":    "t300",
				"password":    "25d55ad283aa400af464c76d713c07ad",
				"nickname":    "T300",
				"create_time": timeStr,
			},
		})
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 2)

		one, err = db.X创建Model对象(table).X条件("id", 200).X查询一条()
		t.AssertNil(err)
		t.Assert(one["id"].X取整数(), 200)
		t.Assert(one["passport"].String(), "t200")
		t.Assert(one["password"].String(), "25d55ad283aa400af464c76d71qw07ad")
		t.Assert(one["nickname"].String(), "T200")
		t.Assert(one["create_time"].X取gtime时间类().String(), timeStr)
	})
}

// 解决问题：https://github.com/gogf/gf/issues/819
func Test_DB_Insert_WithStructAndSliceAttribute(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		type Password struct {
			Salt string `json:"salt"`
			Pass string `json:"pass"`
		}
		data := g.Map{
			"id":          1,
			"passport":    "t1",
			"password":    &Password{"123", "456"},
			"nickname":    []string{"A", "B", "C"},
			"create_time": 时间类.X创建并按当前时间().String(),
		}
		_, err := db.X插入(ctx, table, data)
		t.AssertNil(err)

		one, err := db.X原生SQL查询单条记录(ctx, fmt.Sprintf("SELECT * FROM %s WHERE id=?", table), 1)
		t.AssertNil(err)
		t.Assert(one["passport"], data["passport"])
		t.Assert(one["create_time"], data["create_time"])
		t.Assert(one["nickname"], json类.X创建(data["nickname"]).X取json字节集PANI())
	})
}

func Test_DB_Insert_KeyFieldNameMapping(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
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

func Test_DB_Update_KeyFieldNameMapping(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
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
		t.Assert(one["passport"], data.Passport)
		t.Assert(one["create_time"], data.CreateTime)
		t.Assert(one["nickname"], data.Nickname)
	})
}

func Test_DB_InsertIgnore(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	单元测试类.C(t, func(t *单元测试类.T) {
		_, err := db.X插入(ctx, table, g.Map{
			"id":          1,
			"passport":    "t1",
			"password":    "25d55ad283aa400af464c76d713c07ad",
			"nickname":    "T1",
			"create_time": CreateTime,
		})
		t.AssertNE(err, nil)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		_, err := db.X插入并跳过已存在(ctx, table, g.Map{
			"id":          1,
			"passport":    "t1",
			"password":    "25d55ad283aa400af464c76d713c07ad",
			"nickname":    "T1",
			"create_time": CreateTime,
		})
		t.AssertNil(err)
	})
}

func Test_DB_BatchInsert(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		table := createTable()
		defer dropTable(table)
		r, err := db.X插入(ctx, table, g.Map数组{
			{
				"id":          2,
				"passport":    "t2",
				"password":    "25d55ad283aa400af464c76d713c07ad",
				"nickname":    "name_2",
				"create_time": 时间类.X创建并按当前时间().String(),
			},
			{
				"id":          3,
				"passport":    "user_3",
				"password":    "25d55ad283aa400af464c76d713c07ad",
				"nickname":    "name_3",
				"create_time": 时间类.X创建并按当前时间().String(),
			},
		}, 1)
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 2)

		n, _ = r.LastInsertId()
		t.Assert(n, 3)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		table := createTable()
		defer dropTable(table)
		// []interface{}
		r, err := db.X插入(ctx, table, g.Slice别名{
			g.Map{
				"id":          2,
				"passport":    "t2",
				"password":    "25d55ad283aa400af464c76d713c07ad",
				"nickname":    "name_2",
				"create_time": 时间类.X创建并按当前时间().String(),
			},
			g.Map{
				"id":          3,
				"passport":    "user_3",
				"password":    "25d55ad283aa400af464c76d713c07ad",
				"nickname":    "name_3",
				"create_time": 时间类.X创建并按当前时间().String(),
			},
		}, 1)
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 2)
	})

	// batch insert map
	单元测试类.C(t, func(t *单元测试类.T) {
		table := createTable()
		defer dropTable(table)
		result, err := db.X插入(ctx, table, g.Map{
			"id":          1,
			"passport":    "t1",
			"password":    "p1",
			"nickname":    "T1",
			"create_time": 时间类.X创建并按当前时间().String(),
		})
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
	})
}

func Test_DB_BatchInsert_Struct(t *testing.T) {
	// 批量插入结构体
	单元测试类.C(t, func(t *单元测试类.T) {
		table := createTable()
		defer dropTable(table)

		type User struct {
			Id         int         `c:"id"`
			Passport   string      `c:"passport"`
			Password   string      `c:"password"`
			NickName   string      `c:"nickname"`
			CreateTime *时间类.Time `c:"create_time"`
		}
		user := &User{
			Id:         1,
			Passport:   "t1",
			Password:   "p1",
			NickName:   "T1",
			CreateTime: 时间类.X创建并按当前时间(),
		}
		result, err := db.X插入(ctx, table, user)
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
	})
}

func Test_DB_Save(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		timeStr := 时间类.X创建并按当前时间().String()
		_, err := db.X插入并更新已存在(ctx, table, g.Map{
			"id":          1,
			"passport":    "t1",
			"password":    "25d55ad283aa400af464c76d713c07ad",
			"nickname":    "T11",
			"create_time": timeStr,
		})
		t.Assert(err, ErrorSave)
	})
}

func Test_DB_Replace(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		timeStr := 时间类.X创建并按当前时间().String()
		_, err := db.X插入并替换已存在(ctx, table, g.Map{
			"id":          1,
			"passport":    "t1",
			"password":    "25d55ad283aa400af464c76d713c07ad",
			"nickname":    "T11",
			"create_time": timeStr,
		})
		t.AssertNil(err)

		one, err := db.X创建Model对象(table).X条件("id", 1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["id"].X取整数(), 1)
		t.Assert(one["passport"].String(), "t1")
		t.Assert(one["password"].String(), "25d55ad283aa400af464c76d713c07ad")
		t.Assert(one["nickname"].String(), "T11")
		t.Assert(one["create_time"].X取gtime时间类().String(), timeStr)
	})
}

func Test_DB_Update(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
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

func Test_DB_GetAll(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.GetAll别名(ctx, fmt.Sprintf("SELECT * FROM %s WHERE id=?", table), 1)
		t.AssertNil(err)
		t.Assert(len(result), 1)
		t.Assert(result[0]["id"].X取整数(), 1)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.GetAll别名(ctx, fmt.Sprintf("SELECT * FROM %s WHERE id=?", table), g.Slice别名{1})
		t.AssertNil(err)
		t.Assert(len(result), 1)
		t.Assert(result[0]["id"].X取整数(), 1)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.GetAll别名(ctx, fmt.Sprintf("SELECT * FROM %s WHERE id in(?)", table), g.Slice别名{1, 2, 3})
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["id"].X取整数(), 1)
		t.Assert(result[1]["id"].X取整数(), 2)
		t.Assert(result[2]["id"].X取整数(), 3)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.GetAll别名(ctx, fmt.Sprintf("SELECT * FROM %s WHERE id in(?,?,?)", table), g.Slice别名{1, 2, 3})
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["id"].X取整数(), 1)
		t.Assert(result[1]["id"].X取整数(), 2)
		t.Assert(result[2]["id"].X取整数(), 3)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.GetAll别名(ctx, fmt.Sprintf("SELECT * FROM %s WHERE id in(?,?,?)", table), g.Slice别名{1, 2, 3}...)
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["id"].X取整数(), 1)
		t.Assert(result[1]["id"].X取整数(), 2)
		t.Assert(result[2]["id"].X取整数(), 3)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.GetAll别名(ctx, fmt.Sprintf("SELECT * FROM %s WHERE id>=? AND id <=?", table), g.Slice别名{1, 3})
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["id"].X取整数(), 1)
		t.Assert(result[1]["id"].X取整数(), 2)
		t.Assert(result[2]["id"].X取整数(), 3)
	})
}

func Test_DB_GetOne(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	单元测试类.C(t, func(t *单元测试类.T) {
		record, err := db.X原生SQL查询单条记录(ctx, fmt.Sprintf("SELECT * FROM %s WHERE passport=?", table), "user_1")
		t.AssertNil(err)
		t.Assert(record["nickname"].String(), "name_1")
	})
}

func Test_DB_GetValue(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	单元测试类.C(t, func(t *单元测试类.T) {
		value, err := db.X原生SQL查询字段值(ctx, fmt.Sprintf("SELECT id FROM %s WHERE passport=?", table), "user_3")
		t.AssertNil(err)
		t.Assert(value.X取整数(), 3)
	})
}

func Test_DB_GetCount(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	单元测试类.C(t, func(t *单元测试类.T) {
		count, err := db.X原生SQL查询字段计数(ctx, fmt.Sprintf("SELECT * FROM %s", table))
		t.AssertNil(err)
		t.Assert(count, TableSize)
	})
}

func Test_DB_GetStruct(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime 时间类.Time
		}
		user := new(User)
		err := db.X原生SQL查询到结构体指针(ctx, user, fmt.Sprintf("SELECT * FROM %s WHERE id=?", table), 3)
		t.AssertNil(err)
		t.Assert(user.NickName, "name_3")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *时间类.Time
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
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime 时间类.Time
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

	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *时间类.Time
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

func Test_DB_GetArray(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	单元测试类.C(t, func(t *单元测试类.T) {
		array, err := db.X原生SQL查询数组(ctx, fmt.Sprintf("SELECT id FROM %s WHERE id>?", table), 1)
		t.AssertNil(err)
		t.Assert(len(array), TableSize-1)
		for i, v := range array {
			t.Assert(v.X取整数(), i+2)
		}
	})
}

func Test_DB_GetScan(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime 时间类.Time
		}
		user := new(User)
		err := db.X原生SQL查询到结构体指针(ctx, user, fmt.Sprintf("SELECT * FROM %s WHERE id=?", table), 3)
		t.AssertNil(err)
		t.Assert(user.NickName, "name_3")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime 时间类.Time
		}
		var user *User
		err := db.X原生SQL查询到结构体指针(ctx, &user, fmt.Sprintf("SELECT * FROM %s WHERE id=?", table), 3)
		t.AssertNil(err)
		t.Assert(user.NickName, "name_3")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *时间类.Time
		}
		user := new(User)
		err := db.X原生SQL查询到结构体指针(ctx, user, fmt.Sprintf("SELECT * FROM %s WHERE id=?", table), 3)
		t.AssertNil(err)
		t.Assert(user.NickName, "name_3")
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime 时间类.Time
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

	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *时间类.Time
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
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X删除(ctx, table, 1)
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, TableSize)
	})
}

func Test_DB_Time(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X插入(ctx, table, g.Map{
			"id":          200,
			"passport":    "t200",
			"password":    "123456",
			"nickname":    "T200",
			"create_time": time.Now(),
		})
		if err != nil {
			单元测试类.Error(err)
		}
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
		value, err := db.X原生SQL查询字段值(ctx, fmt.Sprintf("select `passport` from `%s` where id=?", table), 200)
		t.AssertNil(err)
		t.Assert(value.String(), "t200")
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		t1 := time.Now()
		result, err := db.X插入(ctx, table, g.Map{
			"id":          300,
			"passport":    "t300",
			"password":    "123456",
			"nickname":    "T300",
			"create_time": &t1,
		})
		if err != nil {
			单元测试类.Error(err)
		}
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
		value, err := db.X原生SQL查询字段值(ctx, fmt.Sprintf("select `passport` from `%s` where id=?", table), 300)
		t.AssertNil(err)
		t.Assert(value.String(), "t300")
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X删除(ctx, table, 1)
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 2)
	})
}

func Test_DB_ToJson(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	_, err := db.X更新(ctx, table, "create_time='2010-10-10 00:00:01'", "id=?", 1)
	单元测试类.AssertNil(err)

	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X字段保留过滤("*").X条件("id =? ", 1).X查询()
		if err != nil {
			单元测试类.Fatal(err)
		}

		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime string
		}

		users := make([]User, 0)

		err = result.X取数组结构体指针(users)
		t.AssertNE(err, nil)

		err = result.X取数组结构体指针(&users)
		if err != nil {
			单元测试类.Fatal(err)
		}

		// ToJson
		resultJson, err := json类.X加载并自动识别格式(result.X取json())
		if err != nil {
			单元测试类.Fatal(err)
		}

		t.Assert(users[0].Id, resultJson.X取值("0.id").X取整数())
		t.Assert(users[0].Passport, resultJson.X取值("0.passport").String())
		t.Assert(users[0].Password, resultJson.X取值("0.password").String())
		t.Assert(users[0].NickName, resultJson.X取值("0.nickname").String())
		t.Assert(users[0].CreateTime, resultJson.X取值("0.create_time").String())

		result = nil
		t.Assert(result.X取数组结构体指针(&users), sql.ErrNoRows)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X字段保留过滤("*").X条件("id =? ", 1).X查询一条()
		if err != nil {
			单元测试类.Fatal(err)
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
			单元测试类.Fatal(err)
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
	单元测试类.AssertNil(err)

	单元测试类.C(t, func(t *单元测试类.T) {
		record, err := db.X创建Model对象(table).X字段保留过滤("*").X条件("id = ?", 1).X查询一条()
		if err != nil {
			单元测试类.Fatal(err)
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
			单元测试类.Fatal(err)
		}

		result, err := xml类.Decode([]byte(record.X取xml("doc")))
		if err != nil {
			单元测试类.Fatal(err)
		}

		resultXml := result["doc"].(map[string]interface{})
		if v, ok := resultXml["id"]; ok {
			t.Assert(user.Id, v)
		} else {
			单元测试类.Fatal("FAIL")
		}

		if v, ok := resultXml["passport"]; ok {
			t.Assert(user.Passport, v)
		} else {
			单元测试类.Fatal("FAIL")
		}

		if v, ok := resultXml["password"]; ok {
			t.Assert(user.Password, v)
		} else {
			单元测试类.Fatal("FAIL")
		}

		if v, ok := resultXml["nickname"]; ok {
			t.Assert(user.NickName, v)
		} else {
			单元测试类.Fatal("FAIL")
		}

		if v, ok := resultXml["create_time"]; ok {
			t.Assert(user.CreateTime, v)
		} else {
			单元测试类.Fatal("FAIL")
		}
	})
}

func Test_DB_ToStringMap(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	_, err := db.X更新(ctx, table, "create_time='2010-10-10 00:00:01'", "id=?", 1)
	单元测试类.AssertNil(err)
	单元测试类.C(t, func(t *单元测试类.T) {
		id := "1"
		result, err := db.X创建Model对象(table).X字段保留过滤("*").X条件("id = ?", 1).X查询()
		if err != nil {
			单元测试类.Fatal(err)
		}

		type t_user struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime string
		}

		t_users := make([]t_user, 0)
		err = result.X取数组结构体指针(&t_users)
		if err != nil {
			单元测试类.Fatal(err)
		}

		resultStringMap := result.X取字段MapStr("id")
		t.Assert(t_users[0].Id, resultStringMap[id]["id"])
		t.Assert(t_users[0].Passport, resultStringMap[id]["passport"])
		t.Assert(t_users[0].Password, resultStringMap[id]["password"])
		t.Assert(t_users[0].NickName, resultStringMap[id]["nickname"])
		t.Assert(t_users[0].CreateTime, resultStringMap[id]["create_time"])
	})
}

func Test_DB_ToIntMap(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	_, err := db.X更新(ctx, table, "create_time='2010-10-10 00:00:01'", "id=?", 1)
	单元测试类.AssertNil(err)

	单元测试类.C(t, func(t *单元测试类.T) {
		id := 1
		result, err := db.X创建Model对象(table).X字段保留过滤("*").X条件("id = ?", id).X查询()
		if err != nil {
			单元测试类.Fatal(err)
		}

		type t_user struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime string
		}

		t_users := make([]t_user, 0)
		err = result.X取数组结构体指针(&t_users)
		if err != nil {
			单元测试类.Fatal(err)
		}

		resultIntMap := result.X取字段MapInt("id")
		t.Assert(t_users[0].Id, resultIntMap[id]["id"])
		t.Assert(t_users[0].Passport, resultIntMap[id]["passport"])
		t.Assert(t_users[0].Password, resultIntMap[id]["password"])
		t.Assert(t_users[0].NickName, resultIntMap[id]["nickname"])
		t.Assert(t_users[0].CreateTime, resultIntMap[id]["create_time"])
	})
}

func Test_DB_ToUintMap(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	_, err := db.X更新(ctx, table, "create_time='2010-10-10 00:00:01'", "id=?", 1)
	单元测试类.AssertNil(err)

	单元测试类.C(t, func(t *单元测试类.T) {
		id := 1
		result, err := db.X创建Model对象(table).X字段保留过滤("*").X条件("id = ?", id).X查询()
		if err != nil {
			单元测试类.Fatal(err)
		}

		type t_user struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime string
		}

		t_users := make([]t_user, 0)
		err = result.X取数组结构体指针(&t_users)
		if err != nil {
			单元测试类.Fatal(err)
		}

		resultUintMap := result.X取字段MapUint("id")
		t.Assert(t_users[0].Id, resultUintMap[uint(id)]["id"])
		t.Assert(t_users[0].Passport, resultUintMap[uint(id)]["passport"])
		t.Assert(t_users[0].Password, resultUintMap[uint(id)]["password"])
		t.Assert(t_users[0].NickName, resultUintMap[uint(id)]["nickname"])
		t.Assert(t_users[0].CreateTime, resultUintMap[uint(id)]["create_time"])
	})
}

func Test_DB_ToStringRecord(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	_, err := db.X更新(ctx, table, "create_time='2010-10-10 00:00:01'", "id=?", 1)
	单元测试类.AssertNil(err)

	单元测试类.C(t, func(t *单元测试类.T) {
		id := 1
		ids := "1"
		result, err := db.X创建Model对象(table).X字段保留过滤("*").X条件("id = ?", id).X查询()
		if err != nil {
			单元测试类.Fatal(err)
		}

		type t_user struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime string
		}

		t_users := make([]t_user, 0)
		err = result.X取数组结构体指针(&t_users)
		if err != nil {
			单元测试类.Fatal(err)
		}

		resultStringRecord := result.RecordKeyStr("id")
		t.Assert(t_users[0].Id, resultStringRecord[ids]["id"].X取整数())
		t.Assert(t_users[0].Passport, resultStringRecord[ids]["passport"].String())
		t.Assert(t_users[0].Password, resultStringRecord[ids]["password"].String())
		t.Assert(t_users[0].NickName, resultStringRecord[ids]["nickname"].String())
		t.Assert(t_users[0].CreateTime, resultStringRecord[ids]["create_time"].String())
	})
}

func Test_DB_ToIntRecord(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	_, err := db.X更新(ctx, table, "create_time='2010-10-10 00:00:01'", "id=?", 1)
	单元测试类.AssertNil(err)

	单元测试类.C(t, func(t *单元测试类.T) {
		id := 1
		result, err := db.X创建Model对象(table).X字段保留过滤("*").X条件("id = ?", id).X查询()
		if err != nil {
			单元测试类.Fatal(err)
		}

		type t_user struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime string
		}

		t_users := make([]t_user, 0)
		err = result.X取数组结构体指针(&t_users)
		if err != nil {
			单元测试类.Fatal(err)
		}

		resultIntRecord := result.RecordKeyInt("id")
		t.Assert(t_users[0].Id, resultIntRecord[id]["id"].X取整数())
		t.Assert(t_users[0].Passport, resultIntRecord[id]["passport"].String())
		t.Assert(t_users[0].Password, resultIntRecord[id]["password"].String())
		t.Assert(t_users[0].NickName, resultIntRecord[id]["nickname"].String())
		t.Assert(t_users[0].CreateTime, resultIntRecord[id]["create_time"].String())
	})
}

func Test_DB_ToUintRecord(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	_, err := db.X更新(ctx, table, "create_time='2010-10-10 00:00:01'", "id=?", 1)
	单元测试类.AssertNil(err)

	单元测试类.C(t, func(t *单元测试类.T) {
		id := 1
		result, err := db.X创建Model对象(table).X字段保留过滤("*").X条件("id = ?", id).X查询()
		if err != nil {
			单元测试类.Fatal(err)
		}

		type t_user struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime string
		}

		t_users := make([]t_user, 0)
		err = result.X取数组结构体指针(&t_users)
		if err != nil {
			单元测试类.Fatal(err)
		}

		resultUintRecord := result.RecordKeyUint("id")
		t.Assert(t_users[0].Id, resultUintRecord[uint(id)]["id"].X取整数())
		t.Assert(t_users[0].Passport, resultUintRecord[uint(id)]["passport"].String())
		t.Assert(t_users[0].Password, resultUintRecord[uint(id)]["password"].String())
		t.Assert(t_users[0].NickName, resultUintRecord[uint(id)]["nickname"].String())
		t.Assert(t_users[0].CreateTime, resultUintRecord[uint(id)]["create_time"].String())
	})
}

func Test_DB_TableField(t *testing.T) {
	name := "field_test"
	dropTable(name)
	defer dropTable(name)
	_, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
	CREATE TABLE %s (
		field_tinyint  tinyint(8) NULL ,
		field_int  int(8) NULL ,
		field_integer  integer(8) NULL ,
		field_bigint  bigint(8) NULL ,
		field_real  real(8,0) NULL ,
		field_double  double(12,2) NULL ,
		field_varchar  varchar(10) NULL ,
		field_varbinary  varbinary(255) NULL
	);
	`, name))
	if err != nil {
		单元测试类.Fatal(err)
	}

	data := db类.Map{
		"field_tinyint":   1,
		"field_int":       2,
		"field_integer":   3,
		"field_bigint":    4,
		"field_real":      123,
		"field_double":    123.25,
		"field_varchar":   "abc",
		"field_varbinary": "aaa",
	}
	res, err := db.X创建Model对象(name).X设置数据(data).X插入()
	if err != nil {
		单元测试类.Fatal(err)
	}

	n, err := res.RowsAffected()
	if err != nil {
		单元测试类.Fatal(err)
	} else {
		单元测试类.Assert(n, 1)
	}

	result, err := db.X创建Model对象(name).X字段保留过滤("*").X条件("field_int = ?", 2).X查询()
	if err != nil {
		单元测试类.Fatal(err)
	}

	单元测试类.Assert(result[0], data)
}

func Test_DB_Prefix(t *testing.T) {
	db := dbPrefix
	noPrefixName := fmt.Sprintf(`%s_%d`, TableName, 时间类.X取时间戳纳秒())
	table := TableNamePrefix + noPrefixName
	createTableWithDb(db, table)
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		id := 10000
		result, err := db.X插入(ctx, noPrefixName, g.Map{
			"id":          id,
			"passport":    fmt.Sprintf(`user_%d`, id),
			"password":    fmt.Sprintf(`pass_%d`, id),
			"nickname":    fmt.Sprintf(`name_%d`, id),
			"create_time": 时间类.X创建并从文本(CreateTime).String(),
		})
		t.AssertNil(err)

		n, e := result.RowsAffected()
		t.Assert(e, nil)
		t.Assert(n, 1)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		id := 10000
		result, err := db.X插入并替换已存在(ctx, noPrefixName, g.Map{
			"id":          id,
			"passport":    fmt.Sprintf(`user_%d`, id),
			"password":    fmt.Sprintf(`pass_%d`, id),
			"nickname":    fmt.Sprintf(`name_%d`, id),
			"create_time": 时间类.X创建并按当前时间().String(),
		})
		t.AssertNil(err)

		n, e := result.RowsAffected()
		t.Assert(e, nil)
		t.Assert(n, 1)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		id := 10000
		result, err := db.X更新(ctx, noPrefixName, g.Map{
			"id":          id,
			"passport":    fmt.Sprintf(`user_%d`, id),
			"password":    fmt.Sprintf(`pass_%d`, id),
			"nickname":    fmt.Sprintf(`name_%d`, id),
			"create_time": 时间类.X创建并从文本("2018-10-24 10:00:03").String(),
		}, "id=?", id)
		t.AssertNil(err)

		n, e := result.RowsAffected()
		t.Assert(e, nil)
		t.Assert(n, 1)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		id := 10000
		result, err := db.X删除(ctx, noPrefixName, "id=?", id)
		t.AssertNil(err)

		n, e := result.RowsAffected()
		t.Assert(e, nil)
		t.Assert(n, 1)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建(true)
		for i := 1; i <= TableSize; i++ {
			array.Append别名(g.Map{
				"id":          i,
				"passport":    fmt.Sprintf(`user_%d`, i),
				"password":    fmt.Sprintf(`pass_%d`, i),
				"nickname":    fmt.Sprintf(`name_%d`, i),
				"create_time": 时间类.X创建并从文本(CreateTime).String(),
			})
		}

		result, err := db.X插入(ctx, noPrefixName, array.X取切片())
		t.AssertNil(err)

		n, e := result.RowsAffected()
		t.Assert(e, nil)
		t.Assert(n, TableSize)
	})
}

func Test_Model_InnerJoin(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
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

func Test_Empty_Slice_Argument(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.GetAll别名(ctx, fmt.Sprintf(`select * from %s where id in(?)`, table), g.Slice别名{})
		t.AssertNil(err)
		t.Assert(len(result), 0)
	})
}

// 更新计数器测试
func Test_DB_UpdateCounter(t *testing.T) {
	tableName := "gf_update_counter_test_" + 时间类.X取文本时间戳纳秒()
	_, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id INTEGER	PRIMARY KEY AUTOINCREMENT
					UNIQUE
					NOT NULL,
		views  int(8) DEFAULT '0'  NOT NULL ,
		updated_time int(10) DEFAULT '0' NOT NULL
	);
	`, tableName))
	if err != nil {
		单元测试类.Fatal(err)
	}
	defer dropTable(tableName)

	单元测试类.C(t, func(t *单元测试类.T) {
		insertData := g.Map{
			"id":           1,
			"views":        0,
			"updated_time": 0,
		}
		_, err = db.X插入(ctx, tableName, insertData)
		t.AssertNil(err)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		gdbCounter := &db类.Counter{
			Field: "id",
			Value: 1,
		}
		updateData := g.Map{
			"views": gdbCounter,
		}
		result, err := db.X更新(ctx, tableName, updateData, "id", 1)
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
		one, err := db.X创建Model对象(tableName).X条件("id", 1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["id"].X取整数(), 1)
		t.Assert(one["views"].X取整数(), 2)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		gdbCounter := &db类.Counter{
			Field: "views",
			Value: -1,
		}
		updateData := g.Map{
			"views":        gdbCounter,
			"updated_time": 时间类.X创建并按当前时间().Unix(),
		}
		result, err := db.X更新(ctx, tableName, updateData, "id", 1)
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
		one, err := db.X创建Model对象(tableName).X条件("id", 1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["id"].X取整数(), 1)
		t.Assert(one["views"].X取整数(), 1)
	})
}

func Test_DB_Ctx_Logger(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type TraceId string
		defer db.X设置调试模式(db.X取调试模式())
		db.X设置调试模式(true)
		ctx := context.WithValue(context.Background(), TraceId("Trace-Id"), "123456789")
		_, err := db.X原生SQL查询(ctx, "SELECT 1")
		t.AssertNil(err)
	})
}

// 对所有类型进行测试。
// 参考文档：https://www.sqlite.org/datatype3.html
func Test_Types(t *testing.T) {
	tableName := "types_" + 时间类.X取文本时间戳纳秒()
	单元测试类.C(t, func(t *单元测试类.T) {
		if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
    CREATE TABLE IF NOT EXISTS %s (
		id INTEGER	PRIMARY KEY AUTOINCREMENT
					UNIQUE
					NOT NULL,
		%s blob NOT NULL,
		%s binary(8) NOT NULL,
		%s date NOT NULL,
		%s time NOT NULL,
		%s timestamp(6) NOT NULL,
		%s decimal(5,2) NOT NULL,
		%s double NOT NULL,
		%s tinyint(1) NOT NULL,
		%s bool NOT NULL
	);
	`,
			tableName,
			"`blob`",
			"`binary`",
			"`date`",
			"`time`",
			"`timestamp`",
			"`decimal`",
			"`double`",
			"`tinyint`",
			"`bool`")); err != nil {
			单元测试类.Error(err)
		}
		defer dropTable(tableName)
		data := g.Map{
			"id":        1,
			"blob":      "i love gf",
			"binary":    []byte("abcdefgh"),
			"date":      "1880-10-24",
			"time":      "10:00:01",
			"timestamp": "2022-02-14 12:00:01.123456",
			"decimal":   -123.456,
			"double":    -123.456,
			"tinyint":   true,
			"bool":      false,
		}
		r, err := db.X创建Model对象(tableName).X设置数据(data).X插入()
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 1)

		one, err := db.X创建Model对象(tableName).X查询一条()
		t.AssertNil(err)
		t.Assert(one["id"].X取整数(), 1)
		t.Assert(one["blob"].String(), data["blob"])
		t.Assert(one["binary"].String(), data["binary"])
		t.Assert(one["date"].String(), data["date"])
		t.Assert(one["time"].String(), `10:00:01`)
		t.Assert(one["timestamp"].X取gtime时间类().X取格式文本(`Y-m-d H:i:s.u`), `2022-02-14 12:00:01.123`)
		t.Assert(one["decimal"].String(), data["decimal"]) // 在SQLite中，值的数据类型与其自身关联，而不是与其容器关联。
		t.Assert(one["double"].String(), data["double"])
		t.Assert(one["tinyint"].X取布尔(), data["tinyint"])

		type T struct {
			Id        int
			Blob      []byte
			Binary    []byte
			Date      *时间类.Time
			Time      *时间类.Time
			Timestamp *时间类.Time
			Decimal   float64
			Double    float64
			Bit       int8
			TinyInt   bool
		}
		var obj *T
		err = db.X创建Model对象(tableName).X查询到结构体指针(&obj)
		t.AssertNil(err)
		t.Assert(obj.Id, 1)
		t.Assert(obj.Blob, data["blob"])
		t.Assert(obj.Binary, data["binary"])
		t.Assert(obj.Date.X取格式文本("Y-m-d"), data["date"])
		t.Assert(obj.Time.String(), `10:00:01`)
		t.Assert(obj.Timestamp.X取格式文本(`Y-m-d H:i:s.u`), `2022-02-14 12:00:01.123`)
		t.Assert(obj.Decimal, data["decimal"])
		t.Assert(obj.Double, data["double"])
		t.Assert(obj.TinyInt, data["tinyint"])
	})
}

func Test_TableFields(t *testing.T) {

	单元测试类.C(t, func(t *单元测试类.T) {
		tableName := "fields_" + 时间类.X取文本时间戳纳秒()
		createTable(tableName)
		defer dropTable(tableName)
		var expect = map[string][]interface{}{
			// fields 字段名称  类型 是否可为空 主键 默认值 额外信息 注释
			"id":          {"INTEGER", false, "pri", nil, "", ""},
			"passport":    {"VARCHAR(45)", false, "", "passport", "", ""},
			"password":    {"VARCHAR(128)", false, "", "password", "", ""},
			"nickname":    {"VARCHAR(45)", true, "", nil, "", ""},
			"create_time": {"DATETIME", true, "", nil, "", ""},
		}

		res, err := db.X取表字段信息Map(context.Background(), tableName)
		单元测试类.Assert(err, nil)

		for k, v := range expect {
			_, ok := res[k]
			单元测试类.AssertEQ(ok, true)
			单元测试类.AssertEQ(res[k].Name, k)
			单元测试类.AssertEQ(res[k].Type, v[0])
			单元测试类.AssertEQ(res[k].Null, v[1])
			单元测试类.AssertEQ(res[k].Key, v[2])
			单元测试类.AssertEQ(res[k].Default, v[3])
			单元测试类.AssertEQ(res[k].Extra, v[4])
			单元测试类.AssertEQ(res[k].Comment, v[5])
		}

	})

	单元测试类.C(t, func(t *单元测试类.T) {
		_, err := db.X取表字段信息Map(context.Background(), "t1 t2")
		单元测试类.AssertNE(err, nil)
	})
}

func Test_TableNameIsKeyword(t *testing.T) {
	table := createInitTable(TableNameWhichIsKeyword)
	defer dropTable(table)
	_, err := db.X更新(ctx, table, "create_time='2010-10-10 00:00:01'", "id=?", 1)
	单元测试类.AssertNil(err)

	单元测试类.C(t, func(t *单元测试类.T) {
		id := 1
		result, err := db.X创建Model对象(table).X字段保留过滤("*").X条件("id = ?", id).X查询()
		if err != nil {
			单元测试类.Fatal(err)
		}

		type t_user struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime string
		}

		t_users := make([]t_user, 0)
		err = result.X取数组结构体指针(&t_users)
		if err != nil {
			单元测试类.Fatal(err)
		}

		resultIntMap := result.X取字段MapInt("id")
		t.Assert(t_users[0].Id, resultIntMap[id]["id"])
		t.Assert(t_users[0].Passport, resultIntMap[id]["passport"])
		t.Assert(t_users[0].Password, resultIntMap[id]["password"])
		t.Assert(t_users[0].NickName, resultIntMap[id]["nickname"])
		t.Assert(t_users[0].CreateTime, resultIntMap[id]["create_time"])
	})
}
