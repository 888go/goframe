// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package sqlite_test

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	garray "github.com/888go/goframe/container/garray"
	gmap "github.com/888go/goframe/container/gmap"
	gvar "github.com/888go/goframe/container/gvar"
	gdb "github.com/888go/goframe/database/gdb"
	gjson "github.com/888go/goframe/encoding/gjson"
	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/frame/g"
	glog "github.com/888go/goframe/os/glog"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
	gstr "github.com/888go/goframe/text/gstr"
	guid "github.com/888go/goframe/util/guid"
	gutil "github.com/888go/goframe/util/gutil"
)

func Test_Model_Insert(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		user := db.X创建Model对象(table)
		result, err := user.X设置数据(g.Map{
			"id":          1,
			"uid":         1,
			"passport":    "t1",
			"password":    "25d55ad283aa400af464c76d713c07ad",
			"nickname":    "name_1",
			"create_time": gtime.X创建并按当前时间().String(),
		}).X插入()
		t.AssertNil(err)
		n, _ := result.LastInsertId()
		t.Assert(n, 1)

		result, err = db.X创建Model对象(table).X设置数据(g.Map{
			"id":          "2",
			"uid":         "2",
			"passport":    "t2",
			"password":    "25d55ad283aa400af464c76d713c07ad",
			"nickname":    "name_2",
			"create_time": gtime.X创建并按当前时间().String(),
		}).X插入()
		t.AssertNil(err)
		n, _ = result.RowsAffected()
		t.Assert(n, 1)

		type User struct {
			Id         int         `gconv:"id"`
			Uid        int         `gconv:"uid"`
			Passport   string      `json:"passport"`
			Password   string      `gconv:"password"`
			Nickname   string      `gconv:"nickname"`
			CreateTime *gtime.Time `json:"create_time"`
		}
		// Model inserting.
		result, err = db.X创建Model对象(table).X设置数据(User{
			Id:         3,
			Uid:        3,
			Passport:   "t3",
			Password:   "25d55ad283aa400af464c76d713c07ad",
			Nickname:   "name_3",
			CreateTime: gtime.X创建并按当前时间(),
		}).X插入()
		t.AssertNil(err)
		n, _ = result.RowsAffected()
		t.Assert(n, 1)
		value, err := db.X创建Model对象(table).X字段保留过滤("passport").X条件("id=3").X查询一条值()
		t.AssertNil(err)
		t.Assert(value.String(), "t3")

		result, err = db.X创建Model对象(table).X设置数据(&User{
			Id:         4,
			Uid:        4,
			Passport:   "t4",
			Password:   "25d55ad283aa400af464c76d713c07ad",
			Nickname:   "T4",
			CreateTime: gtime.X创建并按当前时间(),
		}).X插入()
		t.AssertNil(err)
		n, _ = result.RowsAffected()
		t.Assert(n, 1)
		value, err = db.X创建Model对象(table).X字段保留过滤("passport").X条件("id=4").X查询一条值()
		t.AssertNil(err)
		t.Assert(value.String(), "t4")

		result, err = db.X创建Model对象(table).X条件("id>?", 1).X删除()
		t.AssertNil(err)
		n, _ = result.RowsAffected()
		t.Assert(n, 3)
	})
}

//github.com/gogf/gf/issues/819. md5:205f368062ae50a5
func Test_Model_Insert_WithStructAndSliceAttribute(t *testing.T) {
	table := createTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		type Password struct {
			Salt string `json:"salt"`
			Pass string `json:"pass"`
		}
		data := g.Map{
			"id":          1,
			"passport":    "t1",
			"password":    &Password{"123", "456"},
			"nickname":    []string{"A", "B", "C"},
			"create_time": gtime.X创建并按当前时间().String(),
		}
		_, err := db.X创建Model对象(table).X设置数据(data).X插入()
		t.AssertNil(err)

		one, err := db.X创建Model对象(table).X查询一条("id", 1)
		t.AssertNil(err)
		t.Assert(one["passport"], data["passport"])
		t.Assert(one["create_time"], data["create_time"])
		t.Assert(one["nickname"], gjson.X创建(data["nickname"]).X取json字节集PANI())
	})
}

func Test_Model_Insert_KeyFieldNameMapping(t *testing.T) {
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
		_, err := db.X创建Model对象(table).X设置数据(data).X插入()
		t.AssertNil(err)

		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["passport"], data.Passport)
		t.Assert(one["create_time"], data.CreateTime)
		t.Assert(one["nickname"], data.Nickname)
	})
}

func Test_Model_Update_KeyFieldNameMapping(t *testing.T) {
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
			Id:         999999,
			Passport:   "user_10",
			Password:   "pass_10",
			Nickname:   "name_10",
			CreateTime: "2020-10-10 12:00:01",
		}
		_, err := db.X创建Model对象(table).X设置数据(data).X条件("id", 1).X更新()
		t.AssertNil(err)

		one, err := db.X创建Model对象(table).X条件("id", data.Id).X查询一条()
		t.AssertNil(err)
		t.Assert(one["passport"], data.Passport)
		t.Assert(one["create_time"], data.CreateTime)
		t.Assert(one["nickname"], data.Nickname)
	})
}

func Test_Model_Insert_Time(t *testing.T) {
	table := createTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		data := g.Map{
			"id":          1,
			"passport":    "t1",
			"password":    "p1",
			"nickname":    "n1",
			"create_time": "2020-10-10 20:09:18.334",
		}
		_, err := db.X创建Model对象(table).X设置数据(data).X插入()
		t.AssertNil(err)

		one, err := db.X创建Model对象(table).X查询一条("id", 1)
		t.AssertNil(err)
		t.Assert(one["passport"], data["passport"])
		t.Assert(one["create_time"], "2020-10-10 20:09:18")
		t.Assert(one["nickname"], data["nickname"])
	})
}

func Test_Model_BatchInsertWithArrayStruct(t *testing.T) {
	table := createTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		user := db.X创建Model对象(table)
		array := garray.X创建()
		for i := 1; i <= TableSize; i++ {
			array.Append别名(g.Map{
				"id":          i,
				"uid":         i,
				"passport":    fmt.Sprintf("t%d", i),
				"password":    "25d55ad283aa400af464c76d713c07ad",
				"nickname":    fmt.Sprintf("name_%d", i),
				"create_time": gtime.X创建并按当前时间().String(),
			})
		}

		result, err := user.X设置数据(array).X插入()
		t.AssertNil(err)
		n, _ := result.LastInsertId()
		t.Assert(n, TableSize)
	})
}

func Test_Model_InsertIgnore(t *testing.T) {
	table := createTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		_, err := db.X创建Model对象(table).X设置数据(g.Map{
			"id":          1,
			"uid":         1,
			"passport":    "t1",
			"password":    "25d55ad283aa400af464c76d713c07ad",
			"nickname":    "name_1",
			"create_time": CreateTime,
		}).X插入()
		t.AssertNil(err)
	})
	gtest.C(t, func(t *gtest.T) {
		_, err := db.X创建Model对象(table).X设置数据(g.Map{
			"id":          1,
			"uid":         1,
			"passport":    "t1",
			"password":    "25d55ad283aa400af464c76d713c07ad",
			"nickname":    "name_1",
			"create_time": CreateTime,
		}).X插入并跳过已存在()
		t.AssertNil(err)
	})
}

func Test_Model_Batch(t *testing.T) {
	// batch insert
	gtest.C(t, func(t *gtest.T) {
		table := createTable()
		defer dropTable(table)
		result, err := db.X创建Model对象(table).X设置数据(g.Map切片{
			{
				"id":          2,
				"uid":         2,
				"passport":    "t2",
				"password":    "25d55ad283aa400af464c76d713c07ad",
				"nickname":    "name_2",
				"create_time": gtime.X创建并按当前时间().String(),
			},
			{
				"id":          3,
				"uid":         3,
				"passport":    "t3",
				"password":    "25d55ad283aa400af464c76d713c07ad",
				"nickname":    "name_3",
				"create_time": gtime.X创建并按当前时间().String(),
			},
		}).X设置批量操作行数(1).X插入()
		if err != nil {
			gtest.Error(err)
		}
		n, _ := result.RowsAffected()
		t.Assert(n, 2)
	})

		// 批量插入，并获取最后插入的自增ID。 md5:b6507323b980f454
	gtest.C(t, func(t *gtest.T) {
		table := createTable()
		defer dropTable(table)
		result, err := db.X创建Model对象(table).X设置数据(g.Map切片{
			{"passport": "t1", "password": "25d55ad283aa400af464c76d713c07ad", "nickname": "name", "create_time": gtime.X创建并按当前时间().String()},
			{"passport": "t2", "password": "25d55ad283aa400af464c76d713c07ad", "nickname": "name", "create_time": gtime.X创建并按当前时间().String()},
			{"passport": "t3", "password": "25d55ad283aa400af464c76d713c07ad", "nickname": "name", "create_time": gtime.X创建并按当前时间().String()},
			{"passport": "t4", "password": "25d55ad283aa400af464c76d713c07ad", "nickname": "name", "create_time": gtime.X创建并按当前时间().String()},
			{"passport": "t5", "password": "25d55ad283aa400af464c76d713c07ad", "nickname": "name", "create_time": gtime.X创建并按当前时间().String()},
		}).X设置批量操作行数(2).X插入()
		if err != nil {
			gtest.Error(err)
		}
		n, _ := result.RowsAffected()
		t.Assert(n, 5)
	})

	// batch replace
	gtest.C(t, func(t *gtest.T) {
		table := createInitTable()
		defer dropTable(table)
		result, err := db.X创建Model对象(table).X查询()
		t.AssertNil(err)
		t.Assert(len(result), TableSize)
		for _, v := range result {
			v["nickname"].X设置值(v["nickname"].String() + v["id"].String())
			v["id"].X设置值(v["id"].X取整数() + 100)
		}
		r, e := db.X创建Model对象(table).X设置数据(result).X插入并替换已存在()
		t.Assert(e, nil)
		n, e := r.RowsAffected()
		t.Assert(e, nil)
		t.Assert(n, TableSize)
	})
}

func Test_Model_Replace(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X设置数据(g.Map{
			"id":          1,
			"passport":    "t11",
			"password":    "25d55ad283aa400af464c76d713c07ad",
			"nickname":    "T11",
			"create_time": CreateTime,
		}).X插入并替换已存在()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)

		one, err := db.X创建Model对象(table).X条件("id", 1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["id"].X取整数(), 1)
		t.Assert(one["passport"].String(), "t11")
		t.Assert(one["password"].String(), "25d55ad283aa400af464c76d713c07ad")
		t.Assert(one["nickname"].String(), "T11")
		t.Assert(one["create_time"].X取gtime时间类().String(), CreateTime)
	})
}

func Test_Model_Save(t *testing.T) {
	table := createTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *gtime.Time
		}
		var (
			user   User
			count  int
			result sql.Result
			err    error
		)

		result, err = db.X创建Model对象(table).X设置数据(g.Map{
			"id":          1,
			"passport":    "CN",
			"password":    "12345678",
			"nickname":    "oldme",
			"create_time": CreateTime,
		}).OnConflict("id").X插入并更新已存在()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)

		err = db.X创建Model对象(table).X查询到结构体指针(&user)
		t.AssertNil(err)
		t.Assert(user.Id, 1)
		t.Assert(user.Passport, "CN")
		t.Assert(user.Password, "12345678")
		t.Assert(user.NickName, "oldme")
		t.Assert(user.CreateTime.String(), CreateTime)

		_, err = db.X创建Model对象(table).X设置数据(g.Map{
			"id":          1,
			"passport":    "CN",
			"password":    "abc123456",
			"nickname":    "to be not to be",
			"create_time": CreateTime,
		}).OnConflict("id").X插入并更新已存在()
		t.AssertNil(err)

		err = db.X创建Model对象(table).X查询到结构体指针(&user)
		t.AssertNil(err)
		t.Assert(user.Passport, "CN")
		t.Assert(user.Password, "abc123456")
		t.Assert(user.NickName, "to be not to be")
		t.Assert(user.CreateTime.String(), CreateTime)

		count, err = db.X创建Model对象(table).X查询行数()
		t.AssertNil(err)
		t.Assert(count, 1)
	})
}

func Test_Model_Update(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	// 更新...限制
	// 使用gtest进行测试，传入t作为测试上下文
	// 执行如下操作：
	// 根据模型table，设置数据字段"nickname"为"T100"，
	// 并在满足条件1的情况下，限制更新操作影响的行数为2。
	// 获取更新操作的结果与错误信息。
	// 断言：期望错误为nil。
	// 计算并获取更新影响的行数，忽略此操作可能产生的错误。
	// 断言：期望更新影响的行数为2。
	// md5:cfae918cd0afb1ea

	// 通过$db$查询$table$表中id为10的nickname字段值，赋值给v1，预期可能产生错误err
	// t.AssertNil(err)：断言错误err为nil，即无错误发生
	// t.Assert(v1.String(), "T100")：断言v1转换为字符串后的值等于"T100"
	// md5:a2bbef8eea48f43a

	// 使用$db$操作数据库，根据模型$table$获取nickname字段，查询id为8的记录，并获取其值。
	// 验证错误是否为nil。
	// 验证获取到的值（v2）是否等于"name_8"。
	// }
	// md5:0005058975deac4b

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X设置数据("passport", "user_22").X条件("passport=?", "user_2").X更新()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
	})

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X设置数据("passport", "user_2").X条件("passport='user_22'").X更新()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
	})

	// Update + Data(string)
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X设置数据("passport='user_33'").X条件("passport='user_3'").X更新()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
	})
		// 更新 + Fields(字符串). md5:df4e16d13da67d5e
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X字段保留过滤("passport").X设置数据(g.Map{
			"passport": "user_44",
			"none":     "none",
		}).X条件("passport='user_4'").X更新()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
	})
}

func Test_Model_UpdateAndGetAffected(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		n, err := db.X创建Model对象(table).X设置数据("nickname", "T100").
			X条件(1).
			X更新并取影响行数()
		t.AssertNil(err)
		t.Assert(n, TableSize)
	})
}

func Test_Model_Clone(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		md := db.X创建Model对象(table).X链式安全(true).X条件("id IN(?)", g.Slice别名{1, 3})
		count, err := md.X查询行数()
		t.AssertNil(err)

		record, err := md.X链式安全(true).X排序("id DESC").X查询一条()
		t.AssertNil(err)

		result, err := md.X链式安全(true).X排序("id ASC").X查询()
		t.AssertNil(err)

		t.Assert(count, int64(2))
		t.Assert(record["id"].X取整数(), 3)
		t.Assert(len(result), 2)
		t.Assert(result[0]["id"].X取整数(), 1)
		t.Assert(result[1]["id"].X取整数(), 3)
	})
}

func Test_Model_Safe(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		md := db.X创建Model对象(table).X链式安全(false).X条件("id IN(?)", g.Slice别名{1, 3})
		count, err := md.X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(2))

		md.X条件("id = ?", 1)
		count, err = md.X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(1))
	})
	gtest.C(t, func(t *gtest.T) {
		md := db.X创建Model对象(table).X链式安全(true).X条件("id IN(?)", g.Slice别名{1, 3})
		count, err := md.X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(2))

		md.X条件("id = ?", 1)
		count, err = md.X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(2))
	})

	gtest.C(t, func(t *gtest.T) {
		md := db.X创建Model对象(table).X链式安全().X条件("id IN(?)", g.Slice别名{1, 3})
		count, err := md.X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(2))

		md.X条件("id = ?", 1)
		count, err = md.X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(2))
	})
	gtest.C(t, func(t *gtest.T) {
		md1 := db.X创建Model对象(table).X链式安全()
		md2 := md1.X条件("id in (?)", g.Slice别名{1, 3})
		count, err := md2.X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(2))

		all, err := md2.X查询()
		t.AssertNil(err)
		t.Assert(len(all), 2)

		all, err = md2.X设置分页(1, 10).X查询()
		t.AssertNil(err)
		t.Assert(len(all), 2)
	})

	gtest.C(t, func(t *gtest.T) {
		table := createInitTable()
		defer dropTable(table)

		md1 := db.X创建Model对象(table).X条件("id>", 0).X链式安全()
		md2 := md1.X条件("id in (?)", g.Slice别名{1, 3})
		md3 := md1.X条件("id in (?)", g.Slice别名{4, 5, 6})

		// 1,3
		count, err := md2.X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(2))

		all, err := md2.X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(all), 2)
		t.Assert(all[0]["id"].X取整数(), 1)
		t.Assert(all[1]["id"].X取整数(), 3)

		all, err = md2.X设置分页(1, 10).X查询()
		t.AssertNil(err)
		t.Assert(len(all), 2)

		// 4,5,6
		count, err = md3.X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(3))

		all, err = md3.X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(all), 3)
		t.Assert(all[0]["id"].X取整数(), 4)
		t.Assert(all[1]["id"].X取整数(), 5)
		t.Assert(all[2]["id"].X取整数(), 6)

		all, err = md3.X设置分页(1, 10).X查询()
		t.AssertNil(err)
		t.Assert(len(all), 3)
	})
}

func Test_Model_All(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X查询()
		t.AssertNil(err)
		t.Assert(len(result), TableSize)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("id<0").X查询()
		t.Assert(result, nil)
		t.AssertNil(err)
	})
}

func Test_Model_AllAndCount(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	tableName2 := "user_" + gtime.X创建并按当前时间().X取文本时间戳纳秒()
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
	CREATE TABLE %s (
		id         INTEGER       PRIMARY KEY AUTOINCREMENT
									UNIQUE
									NOT NULL,
		name       varchar(45) NULL,
		age        int(10)
	);
	`, tableName2,
	)); err != nil {
		gtest.AssertNil(err)
	}
	defer dropTable(tableName2)
	r, err := db.X插入(ctx, tableName2, g.Map{
		"id":   1,
		"name": "table2_1",
		"age":  18,
	})
	gtest.AssertNil(err)
	n, _ := r.RowsAffected()
	gtest.Assert(n, 1)

			// 使用所有数据的AllAndCount. md5:04233fbd8b956565
	gtest.C(t, func(t *gtest.T) {
		result, count, err := db.X创建Model对象(table).X查询与行数(false)
		t.AssertNil(err)
		t.Assert(len(result), TableSize)
		t.Assert(count, TableSize)
	})
		// AllAndCount 无数据情况. md5:78116cd399301bd7
	gtest.C(t, func(t *gtest.T) {
		result, count, err := db.X创建Model对象(table).X条件("id<0").X查询与行数(false)
		t.Assert(result, nil)
		t.AssertNil(err)
		t.Assert(count, 0)
	})
	// AllAndCount with page
	gtest.C(t, func(t *gtest.T) {
		result, count, err := db.X创建Model对象(table).X设置分页(1, 5).X查询与行数(false)
		t.AssertNil(err)
		t.Assert(len(result), 5)
		t.Assert(count, TableSize)
	})
			// AllAndCount 返回正常结果. md5:d132fb7fcbc86207
	gtest.C(t, func(t *gtest.T) {
		result, count, err := db.X创建Model对象(table).X条件("id=?", 1).X查询与行数(false)
		t.AssertNil(err)
		t.Assert(count, 1)
		t.Assert(result[0]["id"], 1)
		t.Assert(result[0]["nickname"], "name_1")
		t.Assert(result[0]["passport"], "user_1")
	})
			// 所有唯一项并计数. md5:ecb27c1ddcd9a325
	gtest.C(t, func(t *gtest.T) {
		result, count, err := db.X创建Model对象(table).X字段保留过滤("DISTINCT nickname").X查询与行数(true)
		t.AssertNil(err)
		t.Assert(count, TableSize)
		t.Assert(result[0]["nickname"], "name_1")
		t.AssertNil(result[0]["id"])
	})
	// AllAndCount with Join
	gtest.C(t, func(t *gtest.T) {
		all, count, err := db.X创建Model对象(table).X设置表别名("u1").
			X左连接(tableName2, "u2", "u2.id=u1.id").
			X字段保留过滤("u1.passport,u1.id,u2.name,u2.age").
			X条件("u1.id<2").
			X查询与行数(false)
		t.AssertNil(err)
		t.Assert(len(all), 1)
		t.Assert(len(all[0]), 4)
		t.Assert(all[0]["id"], 1)
		t.Assert(all[0]["age"], 18)
		t.Assert(all[0]["name"], "table2_1")
		t.Assert(all[0]["passport"], "user_1")
		t.Assert(count, 1)
	})
		// AllAndCount 与 Join 方法返回 CodeDbOperationError. md5:e59618ae9d29f9f5
	gtest.C(t, func(t *gtest.T) {
		all, count, err := db.X创建Model对象(table).X设置表别名("u1").
			X左连接(tableName2, "u2", "u2.id=u1.id").
			X字段保留过滤("u1.passport,u1.id,u2.name,u2.age").
			X条件("u1.id<2").
			X查询与行数(true)
		t.AssertNE(err, nil)
		t.AssertEQ(gerror.X取错误码(err), gcode.CodeDbOperationError)
		t.Assert(count, 0)
		t.Assert(all, nil)
	})
}

func Test_Model_Fields(t *testing.T) {
	tableName1 := createInitTable()
	defer dropTable(tableName1)

	tableName2 := "user_" + gtime.X创建并按当前时间().X取文本时间戳纳秒()
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
	CREATE TABLE %s (
		id         INTEGER       PRIMARY KEY AUTOINCREMENT
									UNIQUE
									NOT NULL,
		name       varchar(45) NULL,
		age        int(10)
	);
	`, tableName2,
	)); err != nil {
		gtest.AssertNil(err)
	}
	defer dropTable(tableName2)

	r, err := db.X插入(ctx, tableName2, g.Map{
		"id":   1,
		"name": "table2_1",
		"age":  18,
	})
	gtest.AssertNil(err)
	n, _ := r.RowsAffected()
	gtest.Assert(n, 1)

	gtest.C(t, func(t *gtest.T) {
		all, err := db.X创建Model对象(tableName1).X设置表别名("u").X字段保留过滤("u.passport,u.id").X条件("u.id<2").X查询()
		t.AssertNil(err)
		t.Assert(len(all), 1)
		t.Assert(len(all[0]), 2)
	})
	gtest.C(t, func(t *gtest.T) {
		all, err := db.X创建Model对象(tableName1).X设置表别名("u1").
			X左连接(tableName1, "u2", "u2.id=u1.id").
			X字段保留过滤("u1.passport,u1.id,u2.id AS u2id").
			X条件("u1.id<2").
			X查询()
		t.AssertNil(err)
		t.Assert(len(all), 1)
		t.Assert(len(all[0]), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		all, err := db.X创建Model对象(tableName1).X设置表别名("u1").
			X左连接(tableName2, "u2", "u2.id=u1.id").
			X字段保留过滤("u1.passport,u1.id,u2.name,u2.age").
			X条件("u1.id<2").
			X查询()
		t.AssertNil(err)
		t.Assert(len(all), 1)
		t.Assert(len(all[0]), 4)
		t.Assert(all[0]["id"], 1)
		t.Assert(all[0]["age"], 18)
		t.Assert(all[0]["name"], "table2_1")
		t.Assert(all[0]["passport"], "user_1")
	})
}

func Test_Model_One(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		record, err := db.X创建Model对象(table).X条件("id", 1).X查询一条()
		t.AssertNil(err)
		t.Assert(record["nickname"].String(), "name_1")
	})

	gtest.C(t, func(t *gtest.T) {
		record, err := db.X创建Model对象(table).X条件("id", 0).X查询一条()
		t.AssertNil(err)
		t.Assert(record, nil)
	})
}

func Test_Model_Value(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		value, err := db.X创建Model对象(table).X字段保留过滤("nickname").X条件("id", 1).X查询一条值()
		t.AssertNil(err)
		t.Assert(value.String(), "name_1")
	})

	gtest.C(t, func(t *gtest.T) {
		value, err := db.X创建Model对象(table).X字段保留过滤("nickname").X条件("id", 0).X查询一条值()
		t.AssertNil(err)
		t.Assert(value, nil)
	})
}

func Test_Model_Array(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		all, err := db.X创建Model对象(table).X条件("id", g.Slice别名{1, 2, 3}).X查询()
		t.AssertNil(err)
		t.Assert(all.X取字段切片("id"), g.Slice别名{1, 2, 3})
		t.Assert(all.X取字段切片("nickname"), g.Slice别名{"name_1", "name_2", "name_3"})
	})
	gtest.C(t, func(t *gtest.T) {
		array, err := db.X创建Model对象(table).X字段保留过滤("nickname").X条件("id", g.Slice别名{1, 2, 3}).X查询切片()
		t.AssertNil(err)
		t.Assert(array, g.Slice别名{"name_1", "name_2", "name_3"})
	})
	gtest.C(t, func(t *gtest.T) {
		array, err := db.X创建Model对象(table).X查询切片("nickname", "id", g.Slice别名{1, 2, 3})
		t.AssertNil(err)
		t.Assert(array, g.Slice别名{"name_1", "name_2", "name_3"})
	})
}

func Test_Model_Count(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		count, err := db.X创建Model对象(table).X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(TableSize))
	})
		// 使用缓存计数，检查内部上下文数据特性。 md5:fa8263fd899afcec
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 10; i++ {
			count, err := db.X创建Model对象(table).X缓存(gdb.CacheOption{
				X时长: time.Second * 10,
				X名称:     guid.X生成(),
				X强制缓存:    false,
			}).X查询行数()
			t.AssertNil(err)
			t.Assert(count, int64(TableSize))
		}
	})
	gtest.C(t, func(t *gtest.T) {
		count, err := db.X创建Model对象(table).X字段排除过滤("id").X条件("id>8").X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(2))
	})
	gtest.C(t, func(t *gtest.T) {
		count, err := db.X创建Model对象(table).X字段保留过滤("distinct id").X条件("id>8").X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(2))
	})
	// COUNT...LIMIT...
	gtest.C(t, func(t *gtest.T) {
		count, err := db.X创建Model对象(table).X设置分页(1, 2).X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(TableSize))
	})
}

func Test_Model_Select(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	type User struct {
		Id         int
		Passport   string
		Password   string
		NickName   string
		CreateTime gtime.Time
	}
	gtest.C(t, func(t *gtest.T) {
		var users []User
		err := db.X创建Model对象(table).X查询到结构体指针(&users)
		t.AssertNil(err)
		t.Assert(len(users), TableSize)
	})
}

func Test_Model_Struct(t *testing.T) {
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
		err := db.X创建Model对象(table).X条件("id=1").X查询到结构体指针(user)
		t.AssertNil(err)
		t.Assert(user.NickName, "name_1")
		t.Assert(user.CreateTime.String(), CreateTime)
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
		err := db.X创建Model对象(table).X条件("id=1").X查询到结构体指针(user)
		t.AssertNil(err)
		t.Assert(user.NickName, "name_1")
		t.Assert(user.CreateTime.String(), CreateTime)
	})
		// 自动创建结构体对象。 md5:4b196dfc1321dc30
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *gtime.Time
		}
		user := (*User)(nil)
		err := db.X创建Model对象(table).X条件("id=1").X查询到结构体指针(&user)
		t.AssertNil(err)
		t.Assert(user.NickName, "name_1")
		t.Assert(user.CreateTime.String(), CreateTime)
	})
	// Just using Scan.
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *gtime.Time
		}
		user := (*User)(nil)
		err := db.X创建Model对象(table).X条件("id=1").X查询到结构体指针(&user)
		if err != nil {
			gtest.Error(err)
		}
		t.Assert(user.NickName, "name_1")
		t.Assert(user.CreateTime.String(), CreateTime)
	})
	// sql.ErrNoRows
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *gtime.Time
		}
		user := new(User)
		err := db.X创建Model对象(table).X条件("id=-1").X查询到结构体指针(user)
		t.Assert(err, sql.ErrNoRows)
	})
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *gtime.Time
		}
		var user *User
		err := db.X创建Model对象(table).X条件("id=-1").X查询到结构体指针(&user)
		t.AssertNil(err)
	})
}

func Test_Model_Struct_CustomType(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	type MyInt int

	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         MyInt
			Passport   string
			Password   string
			NickName   string
			CreateTime gtime.Time
		}
		user := new(User)
		err := db.X创建Model对象(table).X条件("id=1").X查询到结构体指针(user)
		t.AssertNil(err)
		t.Assert(user.NickName, "name_1")
		t.Assert(user.CreateTime.String(), CreateTime)
	})
}

func Test_Model_Structs(t *testing.T) {
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
		err := db.X创建Model对象(table).X排序("id asc").X查询到结构体指针(&users)
		if err != nil {
			gtest.Error(err)
		}
		t.Assert(len(users), TableSize)
		t.Assert(users[0].Id, 1)
		t.Assert(users[1].Id, 2)
		t.Assert(users[2].Id, 3)
		t.Assert(users[0].NickName, "name_1")
		t.Assert(users[1].NickName, "name_2")
		t.Assert(users[2].NickName, "name_3")
		t.Assert(users[0].CreateTime.String(), CreateTime)
	})
		// 自动创建结构体切片。 md5:78598f0d7f20b815
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *gtime.Time
		}
		var users []*User
		err := db.X创建Model对象(table).X排序("id asc").X查询到结构体指针(&users)
		if err != nil {
			gtest.Error(err)
		}
		t.Assert(len(users), TableSize)
		t.Assert(users[0].Id, 1)
		t.Assert(users[1].Id, 2)
		t.Assert(users[2].Id, 3)
		t.Assert(users[0].NickName, "name_1")
		t.Assert(users[1].NickName, "name_2")
		t.Assert(users[2].NickName, "name_3")
		t.Assert(users[0].CreateTime.String(), CreateTime)
	})
	// Just using Scan.
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *gtime.Time
		}
		var users []*User
		err := db.X创建Model对象(table).X排序("id asc").X查询到结构体指针(&users)
		if err != nil {
			gtest.Error(err)
		}
		t.Assert(len(users), TableSize)
		t.Assert(users[0].Id, 1)
		t.Assert(users[1].Id, 2)
		t.Assert(users[2].Id, 3)
		t.Assert(users[0].NickName, "name_1")
		t.Assert(users[1].NickName, "name_2")
		t.Assert(users[2].NickName, "name_3")
		t.Assert(users[0].CreateTime.String(), CreateTime)
	})
	// sql.ErrNoRows
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *gtime.Time
		}
		var users []*User
		err := db.X创建Model对象(table).X条件("id<0").X查询到结构体指针(&users)
		t.AssertNil(err)
	})
}

func Test_Model_StructsWithOrmTag(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		db.X设置调试模式(true)
		defer db.X设置调试模式(false)
		type User struct {
			Uid      int `orm:"id"`
			Passport string
			Password string     `orm:"password"`
			Name     string     `orm:"nickname"`
			Time     gtime.Time `orm:"create_time"`
		}
		var (
			users  []User
			buffer = bytes.NewBuffer(nil)
		)
		db.X取日志记录器().(*glog.Logger).X设置Writer(buffer)
		defer db.X取日志记录器().(*glog.Logger).X设置Writer(os.Stdout)
		db.X创建Model对象(table).X排序("id asc").X查询到结构体指针(&users)
				// 打印出buffer的内容字符串。 md5:3d49298f0e6d7a25
		t.Assert(
			gstr.X是否包含(buffer.String(), "SELECT `id`,`passport`,`password`,`nickname`,`create_time` FROM `user"),
			true,
		)
	})

	gtest.C(t, func(t *gtest.T) {
		type A struct {
			Passport string
			Password string
		}
		type B struct {
			A
			NickName string
		}
		one, err := db.X创建Model对象(table).X字段保留过滤(&B{}).X条件("id", 2).X查询一条()
		t.AssertNil(err)
		t.Assert(len(one), 3)
		t.Assert(one["nickname"], "name_2")
		t.Assert(one["passport"], "user_2")
		t.Assert(one["password"], "pass_2")
	})
}

func Test_Model_Scan(t *testing.T) {
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
		err := db.X创建Model对象(table).X条件("id=1").X查询到结构体指针(user)
		t.AssertNil(err)
		t.Assert(user.NickName, "name_1")
		t.Assert(user.CreateTime.String(), CreateTime)
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
		err := db.X创建Model对象(table).X条件("id=1").X查询到结构体指针(user)
		t.AssertNil(err)
		t.Assert(user.NickName, "name_1")
		t.Assert(user.CreateTime.String(), CreateTime)
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
		err := db.X创建Model对象(table).X排序("id asc").X查询到结构体指针(&users)
		t.AssertNil(err)
		t.Assert(len(users), TableSize)
		t.Assert(users[0].Id, 1)
		t.Assert(users[1].Id, 2)
		t.Assert(users[2].Id, 3)
		t.Assert(users[0].NickName, "name_1")
		t.Assert(users[1].NickName, "name_2")
		t.Assert(users[2].NickName, "name_3")
		t.Assert(users[0].CreateTime.String(), CreateTime)
	})
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *gtime.Time
		}
		var users []*User
		err := db.X创建Model对象(table).X排序("id asc").X查询到结构体指针(&users)
		t.AssertNil(err)
		t.Assert(len(users), TableSize)
		t.Assert(users[0].Id, 1)
		t.Assert(users[1].Id, 2)
		t.Assert(users[2].Id, 3)
		t.Assert(users[0].NickName, "name_1")
		t.Assert(users[1].NickName, "name_2")
		t.Assert(users[2].NickName, "name_3")
		t.Assert(users[0].CreateTime.String(), CreateTime)
	})
	// sql.ErrNoRows
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *gtime.Time
		}
		var (
			user  = new(User)
			users = new([]*User)
		)
		err1 := db.X创建Model对象(table).X条件("id < 0").X查询到结构体指针(user)
		err2 := db.X创建Model对象(table).X条件("id < 0").X查询到结构体指针(users)
		t.Assert(err1, sql.ErrNoRows)
		t.Assert(err2, nil)
	})
}

func Test_Model_ScanAndCount(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	tableName2 := "user_" + gtime.X创建并按当前时间().X取文本时间戳纳秒()
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
	CREATE TABLE %s (
		id         INTEGER       PRIMARY KEY AUTOINCREMENT
									UNIQUE
									NOT NULL,
		name       varchar(45) NULL,
		age        int(10)
	);
	`, tableName2,
	)); err != nil {
		gtest.AssertNil(err)
	}
	defer dropTable(tableName2)
	r, err := db.X插入(ctx, tableName2, g.Map{
		"id":   1,
		"name": "table2_1",
		"age":  18,
	})
	gtest.AssertNil(err)
	n, _ := r.RowsAffected()
	gtest.Assert(n, 1)

			// 使用普通结构体结果的ScanAndCount. md5:941b5fec0e73797f
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *gtime.Time
		}
		user := new(User)
		var count int
		err := db.X创建Model对象(table).X条件("id=1").X查询与行数到指针(user, &count, true)
		t.AssertNil(err)
		t.Assert(user.NickName, "name_1")
		t.Assert(user.CreateTime.String(), CreateTime)
		t.Assert(count, 1)
	})
			// ScanAndCount 使用常规数组作为结果. md5:640a035a18ac03db
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime gtime.Time
		}
		var users []User
		var count int
		err := db.X创建Model对象(table).X排序("id asc").X查询与行数到指针(&users, &count, true)
		t.AssertNil(err)
		t.Assert(len(users), TableSize)
		t.Assert(users[0].Id, 1)
		t.Assert(users[1].Id, 2)
		t.Assert(users[2].Id, 3)
		t.Assert(users[0].NickName, "name_1")
		t.Assert(users[1].NickName, "name_2")
		t.Assert(users[2].NickName, "name_3")
		t.Assert(users[0].CreateTime.String(), CreateTime)
		t.Assert(count, len(users))
	})
	// sql.ErrNoRows
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *gtime.Time
		}
		var (
			user  = new(User)
			users = new([]*User)
		)
		var count1 int
		var count2 int
		err1 := db.X创建Model对象(table).X条件("id < 0").X查询与行数到指针(user, &count1, true)
		err2 := db.X创建Model对象(table).X条件("id < 0").X查询与行数到指针(users, &count2, true)
		t.Assert(count1, 0)
		t.Assert(count2, 0)
		t.Assert(err1, nil)
		t.Assert(err2, nil)
	})
	// ScanAndCount with page
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime gtime.Time
		}
		var users []User
		var count int
		err := db.X创建Model对象(table).X排序("id asc").X设置分页(1, 3).X查询与行数到指针(&users, &count, true)
		t.AssertNil(err)
		t.Assert(len(users), 3)
		t.Assert(count, TableSize)
	})
		// 使用distinct进行扫描和计数. md5:5afa1e02dbecba67
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime gtime.Time
		}
		var users []User
		var count int
		err = db.X创建Model对象(table).X字段保留过滤("distinct id").X查询与行数到指针(&users, &count, true)
		t.AssertNil(err)
		t.Assert(len(users), 10)
		t.Assert(count, TableSize)
		t.Assert(users[0].Id, 1)
	})
	// ScanAndCount with join
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id       int
			Passport string
			Name     string
			Age      int
		}
		var users []User
		var count int
		err = db.X创建Model对象(table).X设置表别名("u1").
			X左连接(tableName2, "u2", "u2.id=u1.id").
			X字段保留过滤("u1.passport,u1.id,u2.name,u2.age").
			X条件("u1.id<2").
			X查询与行数到指针(&users, &count, false)
		t.AssertNil(err)
		t.Assert(len(users), 1)
		t.Assert(count, 1)
		t.AssertEQ(users[0].Name, "table2_1")
	})
			// 使用连接执行ScanAndCount，返回CodeDbOperationError. md5:28f0d53619e4ce12
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id       int
			Passport string
			Name     string
			Age      int
		}
		var users []User
		var count int
		err = db.X创建Model对象(table).X设置表别名("u1").
			X左连接(tableName2, "u2", "u2.id=u1.id").
			X字段保留过滤("u1.passport,u1.id,u2.name,u2.age").
			X条件("u1.id<2").
			X查询与行数到指针(&users, &count, true)
		t.AssertNE(err, nil)
		t.Assert(gerror.X取错误码(err), gcode.CodeDbOperationError)
		t.Assert(count, 0)
		t.AssertEQ(users, nil)
	})
}

func Test_Model_Scan_NilSliceAttrWhenNoRecordsFound(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime gtime.Time
		}
		type Response struct {
			Users []User `json:"users"`
		}
		var res Response
		err := db.X创建Model对象(table).X查询到结构体指针(&res.Users)
		t.AssertNil(err)
		t.Assert(res.Users, nil)
	})
}

func Test_Model_OrderBy(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X排序("id DESC").X查询()
		t.AssertNil(err)
		t.Assert(len(result), TableSize)
		t.Assert(result[0]["nickname"].String(), fmt.Sprintf("name_%d", TableSize))
	})

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X排序("NULL").X查询()
		t.AssertNil(err)
		t.Assert(len(result), TableSize)
		t.Assert(result[0]["nickname"].String(), "name_1")
	})

}

func Test_Model_GroupBy(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X排序分组("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), TableSize)
		t.Assert(result[0]["nickname"].String(), "name_1")
	})
}

func Test_Model_Data(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		table := createInitTable()
		defer dropTable(table)
		result, err := db.X创建Model对象(table).X设置数据("nickname=?", "test").X条件("id=?", 3).X更新()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
	})
	gtest.C(t, func(t *gtest.T) {
		table := createTable()
		defer dropTable(table)
		users := make([]g.MapStrAny, 0)
		for i := 1; i <= 10; i++ {
			users = append(users, g.MapStrAny{
				"id":          i,
				"passport":    fmt.Sprintf(`passport_%d`, i),
				"password":    fmt.Sprintf(`password_%d`, i),
				"nickname":    fmt.Sprintf(`nickname_%d`, i),
				"create_time": gtime.X创建并按当前时间().String(),
			})
		}
		result, err := db.X创建Model对象(table).X设置数据(users).X设置批量操作行数(2).X插入()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 10)
	})
	gtest.C(t, func(t *gtest.T) {
		table := createTable()
		defer dropTable(table)
		users := garray.X创建()
		for i := 1; i <= 10; i++ {
			users.Append别名(g.MapStrAny{
				"id":          i,
				"passport":    fmt.Sprintf(`passport_%d`, i),
				"password":    fmt.Sprintf(`password_%d`, i),
				"nickname":    fmt.Sprintf(`nickname_%d`, i),
				"create_time": gtime.X创建并按当前时间().String(),
			})
		}
		result, err := db.X创建Model对象(table).X设置数据(users).X设置批量操作行数(2).X插入()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 10)
	})
}

func Test_Model_Where(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	// string
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("id=? and nickname=?", 3, "name_3").X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["id"].X取整数(), 3)
	})

	// slice
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件(g.Slice别名{"id", 3}).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["id"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件(g.Slice别名{"id", 3, "nickname", "name_3"}).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["id"].X取整数(), 3)
	})

	// slice parameter
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("id=? and nickname=?", g.Slice别名{3, "name_3"}).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["id"].X取整数(), 3)
	})
	// map like
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件(g.Map{
			"passport like": "user_1%",
		}).X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 2)
		t.Assert(result[0].X取Map类().X取值("id"), 1)
		t.Assert(result[1].X取Map类().X取值("id"), 10)
	})
	// map + slice parameter
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件(g.Map{
			"id":       g.Slice别名{1, 2, 3},
			"passport": g.Slice别名{"user_2", "user_3"},
		}).X条件("id=? and nickname=?", g.Slice别名{3, "name_3"}).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["id"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("id=3", g.Slice别名{}).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["id"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("id=?", g.Slice别名{3}).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["id"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("id", 3).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["id"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("id", 3).X条件("nickname", "name_3").X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("id", 3).X条件("nickname", "name_3").X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("id", 30).X条件或("nickname", "name_3").X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("id", 30).X条件或("nickname", "name_3").X条件("id>?", 1).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("id", 30).X条件或("nickname", "name_3").X条件("id>", 1).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	// slice
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("id=? AND nickname=?", g.Slice别名{3, "name_3"}...).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("id=? AND nickname=?", g.Slice别名{3, "name_3"}).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("passport like ? and nickname like ?", g.Slice别名{"user_3", "name_3"}).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	// map
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件(g.Map{"id": 3, "nickname": "name_3"}).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	// map key operator
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件(g.Map{"id>": 1, "id<": 3}).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 2)
	})

	// gmap.Map
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件(gmap.X创建并从Map(g.MapAnyAny{"id": 3, "nickname": "name_3"})).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	// gmap.Map key operator
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件(gmap.X创建并从Map(g.MapAnyAny{"id>": 1, "id<": 3})).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 2)
	})

	// list map
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件(gmap.X创建链表Map并从Map(g.MapAnyAny{"id": 3, "nickname": "name_3"})).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	// list map key operator
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件(gmap.X创建链表Map并从Map(g.MapAnyAny{"id>": 1, "id<": 3})).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 2)
	})

	// tree map
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件(gmap.X创建红黑树Map并从Map(gutil.X比较文本, g.MapAnyAny{"id": 3, "nickname": "name_3"})).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	// tree map key operator
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件(gmap.X创建红黑树Map并从Map(gutil.X比较文本, g.MapAnyAny{"id>": 1, "id<": 3})).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 2)
	})

	// complicated where 1
	gtest.C(t, func(t *gtest.T) {
		// db.SetDebug(true)
		conditions := g.Map{
			"nickname like ?":    "%name%",
			"id between ? and ?": g.Slice别名{1, 3},
			"id > 0":             nil,
			"create_time > 0":    nil,
			"id":                 g.Slice别名{1, 2, 3},
		}
		result, err := db.X创建Model对象(table).X条件(conditions).X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["id"].X取整数(), 1)
	})
	// complicated where 2
	gtest.C(t, func(t *gtest.T) {
		// db.SetDebug(true)
		conditions := g.Map{
			"nickname like ?":    "%name%",
			"id between ? and ?": g.Slice别名{1, 3},
			"id >= ?":            1,
			"create_time > ?":    0,
			"id in(?)":           g.Slice别名{1, 2, 3},
		}
		result, err := db.X创建Model对象(table).X条件(conditions).X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["id"].X取整数(), 1)
	})
		// 结构体，自动映射和过滤。 md5:8edea55227b914af
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id       int
			Nickname string
		}
		result, err := db.X创建Model对象(table).X条件(User{3, "name_3"}).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)

		result, err = db.X创建Model对象(table).X条件(&User{3, "name_3"}).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	// slice single
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("id IN(?)", g.Slice别名{1, 3}).X排序("id ASC").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 2)
		t.Assert(result[0]["id"].X取整数(), 1)
		t.Assert(result[1]["id"].X取整数(), 3)
	})
	// slice + string
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("nickname=? AND id IN(?)", "name_3", g.Slice别名{1, 3}).X排序("id ASC").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 1)
		t.Assert(result[0]["id"].X取整数(), 3)
	})
	// slice + map
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件(g.Map{
			"id":       g.Slice别名{1, 3},
			"nickname": "name_3",
		}).X排序("id ASC").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 1)
		t.Assert(result[0]["id"].X取整数(), 3)
	})
	// slice + struct
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Ids      []int  `json:"id"`
			Nickname string `gconv:"nickname"`
		}
		result, err := db.X创建Model对象(table).X条件(User{
			Ids:      []int{1, 3},
			Nickname: "name_3",
		}).X排序("id ASC").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 1)
		t.Assert(result[0]["id"].X取整数(), 3)
	})
}

func Test_Model_Where_ISNULL_1(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		// db.SetDebug(true)
		result, err := db.X创建Model对象(table).X设置数据("nickname", nil).X条件("id", 2).X更新()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)

		one, err := db.X创建Model对象(table).X条件("nickname", nil).X查询一条()
		t.AssertNil(err)
		t.Assert(one.X是否为空(), false)
		t.Assert(one["id"], 2)
	})
}

func Test_Model_Where_ISNULL_2(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	// complicated one.
	gtest.C(t, func(t *gtest.T) {
		// db.SetDebug(true)
		conditions := g.Map{
			"nickname like ?":    "%name%",
			"id between ? and ?": g.Slice别名{1, 3},
			"id > 0":             nil,
			"create_time > 0":    nil,
			"id":                 g.Slice别名{1, 2, 3},
		}
		result, err := db.X创建Model对象(table).X条件(conditions).X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["id"].X取整数(), 1)
	})
}

func Test_Model_Where_OmitEmpty(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		conditions := g.Map{
			"id < 4": "",
		}
		result, err := db.X创建Model对象(table).X条件(conditions).X排序("id desc").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["id"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		conditions := g.Map{
			"id < 4": "",
		}
		result, err := db.X创建Model对象(table).X条件(conditions).X过滤空值().X排序("id desc").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 10)
		t.Assert(result[0]["id"].X取整数(), 10)
	})
}

func Test_Model_Where_GTime(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("create_time>?", gtime.X创建并从文本("2010-09-01")).X查询()
		t.AssertNil(err)
		t.Assert(len(result), 10)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("create_time>?", *gtime.X创建并从文本("2010-09-01")).X查询()
		t.AssertNil(err)
		t.Assert(len(result), 10)
	})
}

func Test_Model_WherePri(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	// primary key
	gtest.C(t, func(t *gtest.T) {
		one, err := db.X创建Model对象(table).X条件并识别主键(3).X查询一条()
		t.AssertNil(err)
		t.AssertNE(one, nil)
		t.Assert(one["id"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		all, err := db.X创建Model对象(table).X条件并识别主键(g.Slice别名{3, 9}).X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(all), 2)
		t.Assert(all[0]["id"].X取整数(), 3)
		t.Assert(all[1]["id"].X取整数(), 9)
	})

	// string
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("id=? and nickname=?", 3, "name_3").X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["id"].X取整数(), 3)
	})
	// slice parameter
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("id=? and nickname=?", g.Slice别名{3, "name_3"}).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["id"].X取整数(), 3)
	})
	// map like
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键(g.Map{
			"passport like": "user_1%",
		}).X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 2)
		t.Assert(result[0].X取Map类().X取值("id"), 1)
		t.Assert(result[1].X取Map类().X取值("id"), 10)
	})
	// map + slice parameter
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键(g.Map{
			"id":       g.Slice别名{1, 2, 3},
			"passport": g.Slice别名{"user_2", "user_3"},
		}).X条件("id=? and nickname=?", g.Slice别名{3, "name_3"}).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["id"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键(g.Map{
			"id":       g.Slice别名{1, 2, 3},
			"passport": g.Slice别名{"user_2", "user_3"},
		}).X条件或("nickname=?", g.Slice别名{"name_4"}).X条件("id", 3).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["id"].X取整数(), 2)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("id=3", g.Slice别名{}).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["id"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("id=?", g.Slice别名{3}).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["id"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("id", 3).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["id"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("id", 3).X条件并识别主键("nickname", "name_3").X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("id", 3).X条件("nickname", "name_3").X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("id", 30).X条件或("nickname", "name_3").X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("id", 30).X条件或("nickname", "name_3").X条件("id>?", 1).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("id", 30).X条件或("nickname", "name_3").X条件("id>", 1).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	// slice
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("id=? AND nickname=?", g.Slice别名{3, "name_3"}...).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("id=? AND nickname=?", g.Slice别名{3, "name_3"}).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("passport like ? and nickname like ?", g.Slice别名{"user_3", "name_3"}).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	// map
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键(g.Map{"id": 3, "nickname": "name_3"}).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	// map key operator
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键(g.Map{"id>": 1, "id<": 3}).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 2)
	})

	// gmap.Map
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键(gmap.X创建并从Map(g.MapAnyAny{"id": 3, "nickname": "name_3"})).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	// gmap.Map key operator
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键(gmap.X创建并从Map(g.MapAnyAny{"id>": 1, "id<": 3})).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 2)
	})

	// list map
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键(gmap.X创建链表Map并从Map(g.MapAnyAny{"id": 3, "nickname": "name_3"})).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	// list map key operator
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键(gmap.X创建链表Map并从Map(g.MapAnyAny{"id>": 1, "id<": 3})).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 2)
	})

	// tree map
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键(gmap.X创建红黑树Map并从Map(gutil.X比较文本, g.MapAnyAny{"id": 3, "nickname": "name_3"})).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	// tree map key operator
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键(gmap.X创建红黑树Map并从Map(gutil.X比较文本, g.MapAnyAny{"id>": 1, "id<": 3})).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 2)
	})

	// complicated where 1
	gtest.C(t, func(t *gtest.T) {
		// db.SetDebug(true)
		conditions := g.Map{
			"nickname like ?":    "%name%",
			"id between ? and ?": g.Slice别名{1, 3},
			"id > 0":             nil,
			"create_time > 0":    nil,
			"id":                 g.Slice别名{1, 2, 3},
		}
		result, err := db.X创建Model对象(table).X条件并识别主键(conditions).X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["id"].X取整数(), 1)
	})
	// complicated where 2
	gtest.C(t, func(t *gtest.T) {
		// db.SetDebug(true)
		conditions := g.Map{
			"nickname like ?":    "%name%",
			"id between ? and ?": g.Slice别名{1, 3},
			"id >= ?":            1,
			"create_time > ?":    0,
			"id in(?)":           g.Slice别名{1, 2, 3},
		}
		result, err := db.X创建Model对象(table).X条件并识别主键(conditions).X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["id"].X取整数(), 1)
	})
	// struct
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id       int    `json:"id"`
			Nickname string `gconv:"nickname"`
		}
		result, err := db.X创建Model对象(table).X条件并识别主键(User{3, "name_3"}).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)

		result, err = db.X创建Model对象(table).X条件并识别主键(&User{3, "name_3"}).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	// slice single
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("id IN(?)", g.Slice别名{1, 3}).X排序("id ASC").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 2)
		t.Assert(result[0]["id"].X取整数(), 1)
		t.Assert(result[1]["id"].X取整数(), 3)
	})
	// slice + string
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("nickname=? AND id IN(?)", "name_3", g.Slice别名{1, 3}).X排序("id ASC").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 1)
		t.Assert(result[0]["id"].X取整数(), 3)
	})
	// slice + map
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键(g.Map{
			"id":       g.Slice别名{1, 3},
			"nickname": "name_3",
		}).X排序("id ASC").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 1)
		t.Assert(result[0]["id"].X取整数(), 3)
	})
	// slice + struct
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Ids      []int  `json:"id"`
			Nickname string `gconv:"nickname"`
		}
		result, err := db.X创建Model对象(table).X条件并识别主键(User{
			Ids:      []int{1, 3},
			Nickname: "name_3",
		}).X排序("id ASC").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 1)
		t.Assert(result[0]["id"].X取整数(), 3)
	})
}

func Test_Model_Delete(t *testing.T) {
	// table := createInitTable() 	// 创建初始化表
	// defer dropTable(table)    	// 延迟执行，删除表
	// md5:b569b2401cb8568d

	// DELETE...LIMIT
	// 参考: https:	//github.com/mattn/go-sqlite3/pull/802
	// gtest.C(t, func(t *gtest.T) {
	// 	删除结果, err := db.Model(table).Where(1).Limit(2).Delete()
	// 	t.AssertNil(err)
	// 	影响行数, _ := result.RowsAffected()
	// 	t.Assert(影响行数, 2)
	// })
	// md5:63b42e136740eea6

	gtest.C(t, func(t *gtest.T) {
		table := createInitTable()
		defer dropTable(table)
		result, err := db.X创建Model对象(table).X条件(1).X删除()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, TableSize)
	})
}

func Test_Model_Offset(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X设置条数(2).Offset(5).X排序("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 2)
		t.Assert(result[0]["id"], 6)
		t.Assert(result[1]["id"], 7)
	})
}

func Test_Model_Page(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X设置分页(3, 3).X排序("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["id"], 7)
		t.Assert(result[1]["id"], 8)
	})
	gtest.C(t, func(t *gtest.T) {
		model := db.X创建Model对象(table).X链式安全().X排序("id")
		all, err := model.X设置分页(3, 3).X查询()
		t.AssertNil(err)
		count, err := model.X查询行数()
		t.AssertNil(err)
		t.Assert(len(all), 3)
		t.Assert(all[0]["id"], "7")
		t.Assert(count, int64(TableSize))
	})
}

func Test_Model_Option_Map(t *testing.T) {
	// Insert
	gtest.C(t, func(t *gtest.T) {
		table := createTable()
		defer dropTable(table)
		r, err := db.X创建Model对象(table).X字段保留过滤("id, passport", "password", "create_time").X设置数据(g.Map{
			"id":          1,
			"passport":    "1",
			"password":    "1",
			"nickname":    "1",
			"create_time": gtime.X创建并按当前时间().String(),
		}).X插入()
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 1)
		one, err := db.X创建Model对象(table).X条件("id", 1).X查询一条()
		t.AssertNil(err)
		t.AssertNE(one["password"].String(), "2")
		t.AssertNE(one["nickname"].String(), "2")
		t.Assert(one["passport"].String(), "1")
	})
	gtest.C(t, func(t *gtest.T) {
		table := createTable()
		defer dropTable(table)
		r, err := db.X创建Model对象(table).X过滤空值数据().X设置数据(g.Map{
			"id":          1,
			"passport":    "1",
			"password":    "1",
			"nickname":    "",
			"create_time": gtime.X创建并按当前时间().String(),
		}).X插入()
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 1)
		one, err := db.X创建Model对象(table).X条件("id", 1).X查询一条()
		t.AssertNil(err)
		t.AssertNE(one["passport"].String(), "0")
		t.AssertNE(one["password"].String(), "0")
		t.Assert(one["nickname"].String(), "")
	})

	// Replace
	gtest.C(t, func(t *gtest.T) {
		table := createInitTable()
		defer dropTable(table)
		_, err := db.X创建Model对象(table).X过滤空值数据().X设置数据(g.Map{
			"id":       1,
			"passport": 0,
			"password": 0,
			"nickname": "1",
		}).X插入并替换已存在()
		t.AssertNil(err)
		one, err := db.X创建Model对象(table).X条件("id", 1).X查询一条()
		t.AssertNil(err)
		t.AssertNE(one["passport"].String(), "0")
		t.AssertNE(one["password"].String(), "0")
		t.Assert(one["nickname"].String(), "1")
	})

	// Update
	gtest.C(t, func(t *gtest.T) {
		table := createInitTable()
		defer dropTable(table)

		r, err := db.X创建Model对象(table).X设置数据(g.Map{"nickname": ""}).X条件("id", 1).X更新()
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 1)

		_, err = db.X创建Model对象(table).X过滤空值数据().X设置数据(g.Map{"nickname": ""}).X条件("id", 2).X更新()
		t.AssertNE(err, nil)

		r, err = db.X创建Model对象(table).X过滤空值().X设置数据(g.Map{"nickname": "", "password": "123"}).X条件("id", 3).X更新()
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 1)

		_, err = db.X创建Model对象(table).X过滤空值().X字段保留过滤("nickname").X设置数据(g.Map{"nickname": "", "password": "123"}).X条件("id", 4).X更新()
		t.AssertNE(err, nil)

		r, err = db.X创建Model对象(table).X过滤空值().
			X字段保留过滤("password").X设置数据(g.Map{
			"nickname": "",
			"passport": "123",
			"password": "456",
		}).X条件("id", 5).X更新()
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 1)

		one, err := db.X创建Model对象(table).X条件("id", 5).X查询一条()
		t.AssertNil(err)
		t.Assert(one["password"], "456")
		t.AssertNE(one["passport"].String(), "")
		t.AssertNE(one["passport"].String(), "123")
	})
}

func Test_Model_Option_Where(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		table := createInitTable()
		defer dropTable(table)
		r, err := db.X创建Model对象(table).X过滤空值().X设置数据("nickname", 1).X条件(g.Map{"id": 0, "passport": ""}).X条件(1).X更新()
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, TableSize)
	})
}

func Test_Model_Where_MultiSliceArguments(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		r, err := db.X创建Model对象(table).X条件(g.Map{
			"id":       g.Slice别名{1, 2, 3, 4},
			"passport": g.Slice别名{"user_2", "user_3", "user_4"},
			"nickname": g.Slice别名{"name_2", "name_4"},
			"id >= 4":  nil,
		}).X查询()
		t.AssertNil(err)
		t.Assert(len(r), 1)
		t.Assert(r[0]["id"], 4)
	})

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件(g.Map{
			"id":       g.Slice别名{1, 2, 3},
			"passport": g.Slice别名{"user_2", "user_3"},
		}).X条件或("nickname=?", g.Slice别名{"name_4"}).X条件("id", 3).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["id"].X取整数(), 2)
	})
}

func Test_Model_FieldsEx(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	// Select.
	gtest.C(t, func(t *gtest.T) {
		r, err := db.X创建Model对象(table).X字段排除过滤("create_time, id").X条件("id in (?)", g.Slice别名{1, 2}).X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(r), 2)
		t.Assert(len(r[0]), 3)
		t.Assert(r[0]["id"], "")
		t.Assert(r[0]["passport"], "user_1")
		t.Assert(r[0]["password"], "pass_1")
		t.Assert(r[0]["nickname"], "name_1")
		t.Assert(r[0]["create_time"], "")
		t.Assert(r[1]["id"], "")
		t.Assert(r[1]["passport"], "user_2")
		t.Assert(r[1]["password"], "pass_2")
		t.Assert(r[1]["nickname"], "name_2")
		t.Assert(r[1]["create_time"], "")
	})
	// Update.
	gtest.C(t, func(t *gtest.T) {
		r, err := db.X创建Model对象(table).X字段排除过滤("password").X设置数据(g.Map{"nickname": "123", "password": "456"}).X条件("id", 3).X更新()
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 1)

		one, err := db.X创建Model对象(table).X条件("id", 3).X查询一条()
		t.AssertNil(err)
		t.Assert(one["nickname"], "123")
		t.AssertNE(one["password"], "456")
	})
}

func Test_Model_Prefix(t *testing.T) {
	db := dbPrefix
	noPrefixName := fmt.Sprintf(`%s_%d`, TableName, gtime.X取时间戳纳秒())
	table := TableNamePrefix + noPrefixName
	createInitTableWithDb(db, table)
	defer dropTable(table)
	// Select.
	gtest.C(t, func(t *gtest.T) {
		r, err := db.X创建Model对象(noPrefixName).X条件("id in (?)", g.Slice别名{1, 2}).X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(r), 2)
		t.Assert(r[0]["id"], "1")
		t.Assert(r[1]["id"], "2")
	})
	// Select with alias.
	gtest.C(t, func(t *gtest.T) {
		r, err := db.X创建Model对象(noPrefixName+" as u").X条件("u.id in (?)", g.Slice别名{1, 2}).X排序("u.id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(r), 2)
		t.Assert(r[0]["id"], "1")
		t.Assert(r[1]["id"], "2")
	})
		// 用别名选择到结构体。 md5:86d27c7f5b555a89
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id       int
			Passport string
			Password string
			NickName string
		}
		var users []User
		err := db.X创建Model对象(noPrefixName+" u").X条件("u.id in (?)", g.Slice别名{1, 5}).X排序("u.id asc").X查询到结构体指针(&users)
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].Id, 1)
		t.Assert(users[1].Id, 5)
	})
		// 使用别名和连接语句进行选择。 md5:5ae27281997ad29c
	gtest.C(t, func(t *gtest.T) {
		r, err := db.X创建Model对象(noPrefixName+" as u1").X左连接(noPrefixName+" as u2", "u2.id=u1.id").X条件("u1.id in (?)", g.Slice别名{1, 2}).X排序("u1.id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(r), 2)
		t.Assert(r[0]["id"], "1")
		t.Assert(r[1]["id"], "2")
	})
	gtest.C(t, func(t *gtest.T) {
		r, err := db.X创建Model对象(noPrefixName).X设置表别名("u1").X左连接(noPrefixName+" as u2", "u2.id=u1.id").X条件("u1.id in (?)", g.Slice别名{1, 2}).X排序("u1.id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(r), 2)
		t.Assert(r[0]["id"], "1")
		t.Assert(r[1]["id"], "2")
	})
}

func Test_Model_FieldsExStruct(t *testing.T) {
	table := createTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id       int       `orm:"id"       json:"id"`
			Passport string    `orm:"passport" json:"pass_port"`
			Password string    `orm:"password" json:"password"`
			NickName string    `orm:"nickname" json:"nick__name"`
			Time     time.Time `orm:"create_time" `
		}
		user := &User{
			Id:       1,
			Passport: "111",
			Password: "222",
			NickName: "333",
			Time:     time.Now(),
		}
		r, err := db.X创建Model对象(table).X字段排除过滤("nickname").X过滤空值().X设置数据(user).X插入()
		t.AssertNil(err)
		n, err := r.RowsAffected()
		t.AssertNil(err)
		t.Assert(n, 1)
	})
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id       int       `orm:"id"       json:"id"`
			Passport string    `orm:"passport" json:"pass_port"`
			Password string    `orm:"password" json:"password"`
			NickName string    `orm:"nickname" json:"nick__name"`
			Time     time.Time `orm:"create_time" `
		}
		users := make([]*User, 0)
		for i := 100; i < 110; i++ {
			users = append(users, &User{
				Id:       i,
				Passport: fmt.Sprintf(`passport_%d`, i),
				Password: fmt.Sprintf(`password_%d`, i),
				NickName: fmt.Sprintf(`nickname_%d`, i),
				Time:     time.Now(),
			})
		}
		r, err := db.X创建Model对象(table).X字段排除过滤("nickname").
			X过滤空值().
			X设置批量操作行数(2).
			X设置数据(users).
			X插入()
		t.AssertNil(err)
		n, err := r.RowsAffected()
		t.AssertNil(err)
		t.Assert(n, 10)
	})
}

func Test_Model_OmitEmpty_Time(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id       int       `orm:"id"       json:"id"`
			Passport string    `orm:"password" json:"pass_port"`
			Password string    `orm:"password" json:"password"`
			Time     time.Time `orm:"create_time" `
		}
		user := &User{
			Id:       1,
			Passport: "111",
			Password: "222",
			Time:     time.Time{},
		}
		r, err := db.X创建Model对象(table).X过滤空值().X设置数据(user).X条件("id", 1).X更新()
		t.AssertNil(err)
		n, err := r.RowsAffected()
		t.AssertNil(err)
		t.Assert(n, 1)
	})
}

func Test_Result_Chunk(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		r, err := db.X创建Model对象(table).X排序("id asc").X查询()
		t.AssertNil(err)
		chunks := r.X分割(3)
		t.Assert(len(chunks), 4)
		t.Assert(chunks[0][0]["id"].X取整数(), 1)
		t.Assert(chunks[1][0]["id"].X取整数(), 4)
		t.Assert(chunks[2][0]["id"].X取整数(), 7)
		t.Assert(chunks[3][0]["id"].X取整数(), 10)
	})
}

func Test_Model_DryRun(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	db.X设置空跑特性(true)
	defer db.X设置空跑特性(false)

	gtest.C(t, func(t *gtest.T) {
		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["id"], 1)
	})
	gtest.C(t, func(t *gtest.T) {
		r, err := db.X创建Model对象(table).X设置数据("passport", "port_1").X条件并识别主键(1).X更新()
		t.AssertNil(err)
		n, err := r.RowsAffected()
		t.AssertNil(err)
		t.Assert(n, 0)
	})
}

func Test_Model_Join_SubQuery(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		subQuery := fmt.Sprintf("select * from `%s`", table)
		r, err := db.X创建Model对象(table, "t1").X字段保留过滤("t2.id").X左连接(subQuery, "t2", "t2.id=t1.id").X查询切片()
		t.AssertNil(err)
		t.Assert(len(r), TableSize)
		t.Assert(r[0], "1")
		t.Assert(r[TableSize-1], TableSize)
	})
}

func Test_Model_Cache(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		one, err := db.X创建Model对象(table).X缓存(gdb.CacheOption{
			X时长: time.Second,
			X名称:     "test1",
			X强制缓存:    false,
		}).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["passport"], "user_1")

		r, err := db.X创建Model对象(table).X设置数据("passport", "user_100").X条件并识别主键(1).X更新()
		t.AssertNil(err)
		n, err := r.RowsAffected()
		t.AssertNil(err)
		t.Assert(n, 1)

		one, err = db.X创建Model对象(table).X缓存(gdb.CacheOption{
			X时长: time.Second,
			X名称:     "test1",
			X强制缓存:    false,
		}).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["passport"], "user_1")

		time.Sleep(time.Second * 2)

		one, err = db.X创建Model对象(table).X缓存(gdb.CacheOption{
			X时长: time.Second,
			X名称:     "test1",
			X强制缓存:    false,
		}).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["passport"], "user_100")
	})
	gtest.C(t, func(t *gtest.T) {
		one, err := db.X创建Model对象(table).X缓存(gdb.CacheOption{
			X时长: time.Second,
			X名称:     "test2",
			X强制缓存:    false,
		}).X条件并识别主键(2).X查询一条()
		t.AssertNil(err)
		t.Assert(one["passport"], "user_2")

		r, err := db.X创建Model对象(table).X设置数据("passport", "user_200").X缓存(gdb.CacheOption{
			X时长: -1,
			X名称:     "test2",
			X强制缓存:    false,
		}).X条件并识别主键(2).X更新()
		t.AssertNil(err)
		n, err := r.RowsAffected()
		t.AssertNil(err)
		t.Assert(n, 1)

		one, err = db.X创建Model对象(table).X缓存(gdb.CacheOption{
			X时长: time.Second,
			X名称:     "test2",
			X强制缓存:    false,
		}).X条件并识别主键(2).X查询一条()
		t.AssertNil(err)
		t.Assert(one["passport"], "user_200")
	})
	// transaction.
	gtest.C(t, func(t *gtest.T) {
		// make cache for id 3
		one, err := db.X创建Model对象(table).X缓存(gdb.CacheOption{
			X时长: time.Second,
			X名称:     "test3",
			X强制缓存:    false,
		}).X条件并识别主键(3).X查询一条()
		t.AssertNil(err)
		t.Assert(one["passport"], "user_3")

		r, err := db.X创建Model对象(table).X设置数据("passport", "user_300").X缓存(gdb.CacheOption{
			X时长: time.Second,
			X名称:     "test3",
			X强制缓存:    false,
		}).X条件并识别主键(3).X更新()
		t.AssertNil(err)
		n, err := r.RowsAffected()
		t.AssertNil(err)
		t.Assert(n, 1)

		err = db.X事务(context.TODO(), func(ctx context.Context, tx gdb.TX) error {
			one, err := tx.X创建Model对象(table).X缓存(gdb.CacheOption{
				X时长: time.Second,
				X名称:     "test3",
				X强制缓存:    false,
			}).X条件并识别主键(3).X查询一条()
			t.AssertNil(err)
			t.Assert(one["passport"], "user_300")
			return nil
		})
		t.AssertNil(err)

		one, err = db.X创建Model对象(table).X缓存(gdb.CacheOption{
			X时长: time.Second,
			X名称:     "test3",
			X强制缓存:    false,
		}).X条件并识别主键(3).X查询一条()
		t.AssertNil(err)
		t.Assert(one["passport"], "user_3")
	})
	gtest.C(t, func(t *gtest.T) {
		// make cache for id 4
		one, err := db.X创建Model对象(table).X缓存(gdb.CacheOption{
			X时长: time.Second,
			X名称:     "test4",
			X强制缓存:    false,
		}).X条件并识别主键(4).X查询一条()
		t.AssertNil(err)
		t.Assert(one["passport"], "user_4")

		r, err := db.X创建Model对象(table).X设置数据("passport", "user_400").X缓存(gdb.CacheOption{
			X时长: time.Second,
			X名称:     "test3",
			X强制缓存:    false,
		}).X条件并识别主键(4).X更新()
		t.AssertNil(err)
		n, err := r.RowsAffected()
		t.AssertNil(err)
		t.Assert(n, 1)

		err = db.X事务(context.TODO(), func(ctx context.Context, tx gdb.TX) error {
						// 缓存功能已禁用。 md5:96110ddd3191b243
			one, err := tx.X创建Model对象(table).X缓存(gdb.CacheOption{
				X时长: time.Second,
				X名称:     "test4",
				X强制缓存:    false,
			}).X条件并识别主键(4).X查询一条()
			t.AssertNil(err)
			t.Assert(one["passport"], "user_400")
			// Update the cache.
			r, err := tx.X创建Model对象(table).X设置数据("passport", "user_4000").
				X缓存(gdb.CacheOption{
					X时长: -1,
					X名称:     "test4",
					X强制缓存:    false,
				}).X条件并识别主键(4).X更新()
			t.AssertNil(err)
			n, err := r.RowsAffected()
			t.AssertNil(err)
			t.Assert(n, 1)
			return nil
		})
		t.AssertNil(err)
		// Read from db.
		one, err = db.X创建Model对象(table).X缓存(gdb.CacheOption{
			X时长: time.Second,
			X名称:     "test4",
			X强制缓存:    false,
		}).X条件并识别主键(4).X查询一条()
		t.AssertNil(err)
		t.Assert(one["passport"], "user_4000")
	})
}

func Test_Model_Having(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		all, err := db.X创建Model对象(table).X条件("id > 1").X排序分组("id").X设置分组条件("id > 8").X查询()
		t.AssertNil(err)
		t.Assert(len(all), 2)
	})
	gtest.C(t, func(t *gtest.T) {
		all, err := db.X创建Model对象(table).X条件("id > 1").X排序分组("id").X设置分组条件("id > ?", 8).X查询()
		t.AssertNil(err)
		t.Assert(len(all), 2)
	})
	gtest.C(t, func(t *gtest.T) {
		all, err := db.X创建Model对象(table).X条件("id > ?", 1).X排序分组("id").X设置分组条件("id > ?", 8).X查询()
		t.AssertNil(err)
		t.Assert(len(all), 2)
	})
	gtest.C(t, func(t *gtest.T) {
		all, err := db.X创建Model对象(table).X条件("id > ?", 1).X排序分组("id").X设置分组条件("id", 8).X查询()
		t.AssertNil(err)
		t.Assert(len(all), 1)
	})
}

func Test_Model_Distinct(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		all, err := db.X创建Model对象(table, "t").X字段保留过滤("distinct t.id").X条件("id > 1").X排序分组("id").X设置分组条件("id > 8").X查询()
		t.AssertNil(err)
		t.Assert(len(all), 2)
	})
	gtest.C(t, func(t *gtest.T) {
		count, err := db.X创建Model对象(table).X条件("id > 1").X设置去重().X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(9))
	})
}

func Test_Model_Min_Max(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		value, err := db.X创建Model对象(table, "t").X字段保留过滤("min(t.id)").X条件("id > 1").X查询一条值()
		t.AssertNil(err)
		t.Assert(value.X取整数(), 2)
	})
	gtest.C(t, func(t *gtest.T) {
		value, err := db.X创建Model对象(table, "t").X字段保留过滤("max(t.id)").X条件("id > 1").X查询一条值()
		t.AssertNil(err)
		t.Assert(value.X取整数(), 10)
	})
}

func Test_Model_Fields_AutoMapping(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		value, err := db.X创建Model对象(table).X字段保留过滤("ID").X条件("id", 2).X查询一条值()
		t.AssertNil(err)
		t.Assert(value.X取整数(), 2)
	})

	gtest.C(t, func(t *gtest.T) {
		value, err := db.X创建Model对象(table).X字段保留过滤("NICK_NAME").X条件("id", 2).X查询一条值()
		t.AssertNil(err)
		t.Assert(value.String(), "name_2")
	})
	// Map
	gtest.C(t, func(t *gtest.T) {
		one, err := db.X创建Model对象(table).X字段保留过滤(g.Map{
			"ID":        1,
			"NICK_NAME": 1,
		}).X条件("id", 2).X查询一条()
		t.AssertNil(err)
		t.Assert(len(one), 2)
		t.Assert(one["id"], 2)
		t.Assert(one["nickname"], "name_2")
	})
	// Struct
	gtest.C(t, func(t *gtest.T) {
		type T struct {
			ID       int
			NICKNAME int
		}
		one, err := db.X创建Model对象(table).X字段保留过滤(&T{
			ID:       0,
			NICKNAME: 0,
		}).X条件("id", 2).X查询一条()
		t.AssertNil(err)
		t.Assert(len(one), 2)
		t.Assert(one["id"], 2)
		t.Assert(one["nickname"], "name_2")
	})
}

func Test_Model_FieldsEx_AutoMapping(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	// "id":          i, 	// 用户ID
	// "passport":    fmt.Sprintf("user_%d", i), 	// 通行证（格式为"user_编号")
	// "password":    fmt.Sprintf("pass_%d", i), 	// 密码（格式为"pass_编号")
	// "nickname":    fmt.Sprintf("name_%d", i), 	// 昵称（格式为"name_编号")
	// "create_time": gtime.NewFromStr(CreateTime).String(), 	// 创建时间（将CreateTime字符串转换为gtime格式并转为字符串）
	// md5:ddd0764dc67c4e9f

	gtest.C(t, func(t *gtest.T) {
		value, err := db.X创建Model对象(table).X字段排除过滤("Passport, Password, NickName, CreateTime").X条件("id", 2).X查询一条值()
		t.AssertNil(err)
		t.Assert(value.X取整数(), 2)
	})

	gtest.C(t, func(t *gtest.T) {
		value, err := db.X创建Model对象(table).X字段排除过滤("ID, Passport, Password, CreateTime").X条件("id", 2).X查询一条值()
		t.AssertNil(err)
		t.Assert(value.String(), "name_2")
	})
	// Map
	gtest.C(t, func(t *gtest.T) {
		one, err := db.X创建Model对象(table).X字段排除过滤(g.Map{
			"Passport":   1,
			"Password":   1,
			"CreateTime": 1,
		}).X条件("id", 2).X查询一条()
		t.AssertNil(err)
		t.Assert(len(one), 2)
		t.Assert(one["id"], 2)
		t.Assert(one["nickname"], "name_2")
	})
	// Struct
	gtest.C(t, func(t *gtest.T) {
		type T struct {
			Passport   int
			Password   int
			CreateTime int
		}
		one, err := db.X创建Model对象(table).X字段排除过滤(&T{
			Passport:   0,
			Password:   0,
			CreateTime: 0,
		}).X条件("id", 2).X查询一条()
		t.AssertNil(err)
		t.Assert(len(one), 2)
		t.Assert(one["id"], 2)
		t.Assert(one["nickname"], "name_2")
	})
}

func Test_Model_Fields_Struct(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	type A struct {
		Passport string
		Password string
	}
	type B struct {
		A
		NickName string
	}
	gtest.C(t, func(t *gtest.T) {
		one, err := db.X创建Model对象(table).X字段保留过滤(A{}).X条件("id", 2).X查询一条()
		t.AssertNil(err)
		t.Assert(len(one), 2)
		t.Assert(one["passport"], "user_2")
		t.Assert(one["password"], "pass_2")
	})
	gtest.C(t, func(t *gtest.T) {
		one, err := db.X创建Model对象(table).X字段保留过滤(&A{}).X条件("id", 2).X查询一条()
		t.AssertNil(err)
		t.Assert(len(one), 2)
		t.Assert(one["passport"], "user_2")
		t.Assert(one["password"], "pass_2")
	})
	gtest.C(t, func(t *gtest.T) {
		one, err := db.X创建Model对象(table).X字段保留过滤(B{}).X条件("id", 2).X查询一条()
		t.AssertNil(err)
		t.Assert(len(one), 3)
		t.Assert(one["passport"], "user_2")
		t.Assert(one["password"], "pass_2")
		t.Assert(one["nickname"], "name_2")
	})
	gtest.C(t, func(t *gtest.T) {
		one, err := db.X创建Model对象(table).X字段保留过滤(&B{}).X条件("id", 2).X查询一条()
		t.AssertNil(err)
		t.Assert(len(one), 3)
		t.Assert(one["passport"], "user_2")
		t.Assert(one["password"], "pass_2")
		t.Assert(one["nickname"], "name_2")
	})
}

func Test_Model_Empty_Slice_Argument(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件(`id`, g.Slice别名{}).X查询()
		t.AssertNil(err)
		t.Assert(len(result), 0)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件(`id in(?)`, g.Slice别名{}).X查询()
		t.AssertNil(err)
		t.Assert(len(result), 0)
	})
}

func Test_Model_HasTable(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		t.AssertNil(db.X取Core对象().X删除所有表查询缓存(ctx))
		result, err := db.X取Core对象().X是否存在表名(table)
		t.Assert(result, true)
		t.AssertNil(err)
	})

	gtest.C(t, func(t *gtest.T) {
		t.AssertNil(db.X取Core对象().X删除所有表查询缓存(ctx))
		result, err := db.X取Core对象().X是否存在表名("table12321")
		t.Assert(result, false)
		t.AssertNil(err)
	})
}

func Test_Model_HasField(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X是否存在字段("id")
		t.Assert(result, true)
		t.AssertNil(err)
	})

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X是否存在字段("id123")
		t.Assert(result, false)
		t.AssertNil(err)
	})
}

//github.com/gogf/gf/issues/1002. md5:2b9ad829e9523427
func Test_Model_Issue1002(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	result, err := db.X创建Model对象(table).X设置数据(g.Map{
		"id":          1,
		"passport":    "port_1",
		"password":    "pass_1",
		"nickname":    "name_2",
		"create_time": "2020-10-27 19:03:33",
	}).X插入()
	gtest.AssertNil(err)
	n, _ := result.RowsAffected()
	gtest.Assert(n, 1)

	// where + string.
	gtest.C(t, func(t *gtest.T) {
		v, err := db.X创建Model对象(table).X字段保留过滤("id").X条件("create_time>'2020-10-27 19:03:32' and create_time<'2020-10-27 19:03:34'").X查询一条值()
		t.AssertNil(err)
		t.Assert(v.X取整数(), 1)
	})
	gtest.C(t, func(t *gtest.T) {
		v, err := db.X创建Model对象(table).X字段保留过滤("id").X条件("create_time>'2020-10-27 19:03:32' and create_time<'2020-10-27 19:03:34'").X查询一条值()
		t.AssertNil(err)
		t.Assert(v.X取整数(), 1)
	})
		// where + 字符串参数。 md5:cb1db92222691d4d
	gtest.C(t, func(t *gtest.T) {
		v, err := db.X创建Model对象(table).X字段保留过滤("id").X条件("create_time>? and create_time<?", "2020-10-27 19:03:32", "2020-10-27 19:03:34").X查询一条值()
		t.AssertNil(err)
		t.Assert(v.X取整数(), 1)
	})
		// 其中包含 gtime.Time 类型的参数。 md5:3bd9bb993dd2cc53
	gtest.C(t, func(t *gtest.T) {
		v, err := db.X创建Model对象(table).X字段保留过滤("id").X条件("create_time>? and create_time<?", gtime.X创建("2020-10-27 19:03:32"), gtime.X创建("2020-10-27 19:03:34")).X查询一条值()
		t.AssertNil(err)
		t.Assert(v.X取整数(), 1)
	})
	// 待办事项
	// 在这里使用 + time.Time 参数，采用 UTC 时间。
	// gtest.C(t, func(t *gtest.T) {
	//   t1, _ := time.Parse("2006-01-02 15:04:05", "2020-10-27 11:03:32") 	// 解析时间字符串为 t1
	//   t2, _ := time.Parse("2006-01-02 15:04:05", "2020-10-27 11:03:34") 	// 解析时间字符串为 t2
	//   {
	//     v, err := db.Model(table).Fields("id").Where("create_time>? and create_time<?", t1, t2).Value() 	// 查询创建时间在 t1 和 t2 之间记录的 id
	//     t.AssertNil(err) 	// 断言 err 为空，即查询无错误
	//     t.Assert(v.Int(), 1) 	// 断言查询结果的整数值为 1
	//   }
	// })
	// md5:6089a1ebb4983ace
}

func createTableForTimeZoneTest() string {
	tableName := "user_" + gtime.X创建并按当前时间().X取文本时间戳纳秒()
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id INTEGER	PRIMARY KEY AUTOINCREMENT
					UNIQUE
					NOT NULL,
		passport    varchar(45) NULL,
		password    char(32) NULL,
		nickname    varchar(45) NULL,
		created_at timestamp NULL,
		updated_at timestamp NULL,
		deleted_at timestamp NULL
	);
	`, tableName,
	)); err != nil {
		gtest.Fatal(err)
	}
	return tableName
}

// 这段注释指的是在GitHub上的一个gf项目（Golang Fast Foundation，一个Go语言的优秀库）中的Issue 1012。"Issue"通常在GitHub上表示一个问题、错误报告或者改进的请求。所以，这个注释可能是在指有关gf库的一个已知问题或者开发者希望解决的问题，链接指向了该问题的具体页面。 md5:d21c0bba53139335
func Test_TimeZoneInsert(t *testing.T) {
	tableName := createTableForTimeZoneTest()
	defer dropTable(tableName)

	tokyoLoc, err := time.LoadLocation("Asia/Tokyo")
	gtest.AssertNil(err)

	CreateTime := "2020-11-22 12:23:45"
	UpdateTime := "2020-11-22 13:23:45"
	DeleteTime := "2020-11-22 14:23:45"
	type User struct {
		Id        int         `json:"id"`
		CreatedAt *gtime.Time `json:"created_at"`
		UpdatedAt gtime.Time  `json:"updated_at"`
		DeletedAt time.Time   `json:"deleted_at"`
	}
	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", CreateTime, tokyoLoc)
	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", UpdateTime, tokyoLoc)
	t3, _ := time.ParseInLocation("2006-01-02 15:04:05", DeleteTime, tokyoLoc)
	u := &User{
		Id:        1,
		CreatedAt: gtime.X创建(t1.UTC()),
		UpdatedAt: *gtime.X创建(t2.UTC()),
		DeletedAt: t3.UTC(),
	}

	gtest.C(t, func(t *gtest.T) {
		_, _ = db.X创建Model对象(tableName).Unscoped().X插入(u)
		userEntity := &User{}
		err := db.X创建Model对象(tableName).X条件("id", 1).Unscoped().X查询到结构体指针(&userEntity)
		t.AssertNil(err)
		// TODO
		// t.Assert(userEntity.CreatedAt.String(), "2020-11-22 11:23:45") 		// 断言用户实体的创建时间字符串为 "2020-11-22 11:23:45"
		// t.Assert(userEntity.UpdatedAt.String(), "2020-11-22 12:23:45") 		// 断言用户实体的更新时间字符串为 "2020-11-22 12:23:45"
		// t.Assert(gtime.NewFromTime(userEntity.DeletedAt).String(), "2020-11-22 13:23:45") 		// 断言用户实体的删除时间（转换为gtime类型）字符串为 "2020-11-22 13:23:45"
		// md5:8ad9ae5f1d9029d0
	})
}

func Test_Model_Fields_Map_Struct(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	// map
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X字段保留过滤(g.Map{
			"ID":         1,
			"PASSPORT":   1,
			"NONE_EXIST": 1,
		}).X条件("id", 1).X查询一条()
		t.AssertNil(err)
		t.Assert(len(result), 2)
		t.Assert(result["id"], 1)
		t.Assert(result["passport"], "user_1")
	})
	// struct
	gtest.C(t, func(t *gtest.T) {
		type A struct {
			ID       int
			PASSPORT string
			XXX_TYPE int
		}
		a := A{}
		err := db.X创建Model对象(table).X字段保留过滤(a).X条件("id", 1).X查询到结构体指针(&a)
		t.AssertNil(err)
		t.Assert(a.ID, 1)
		t.Assert(a.PASSPORT, "user_1")
		t.Assert(a.XXX_TYPE, 0)
	})
	// *struct
	gtest.C(t, func(t *gtest.T) {
		type A struct {
			ID       int
			PASSPORT string
			XXX_TYPE int
		}
		var a *A
		err := db.X创建Model对象(table).X字段保留过滤(a).X条件("id", 1).X查询到结构体指针(&a)
		t.AssertNil(err)
		t.Assert(a.ID, 1)
		t.Assert(a.PASSPORT, "user_1")
		t.Assert(a.XXX_TYPE, 0)
	})
	// **struct
	gtest.C(t, func(t *gtest.T) {
		type A struct {
			ID       int
			PASSPORT string
			XXX_TYPE int
		}
		var a *A
		err := db.X创建Model对象(table).X字段保留过滤(&a).X条件("id", 1).X查询到结构体指针(&a)
		t.AssertNil(err)
		t.Assert(a.ID, 1)
		t.Assert(a.PASSPORT, "user_1")
		t.Assert(a.XXX_TYPE, 0)
	})
}

func Test_Model_WhereIn(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件包含("id", g.Slice别名{1, 2, 3, 4}).X条件包含("id", g.Slice别名{3, 4, 5}).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 2)
		t.Assert(result[0]["id"], 3)
		t.Assert(result[1]["id"], 4)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件包含("id", g.Slice别名{}).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 0)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X过滤空值条件().X条件包含("id", g.Slice别名{}).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), TableSize)
	})
}

func Test_Model_WhereNotIn(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件不包含("id", g.Slice别名{1, 2, 3, 4}).X条件不包含("id", g.Slice别名{3, 4, 5}).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 5)
		t.Assert(result[0]["id"], 6)
		t.Assert(result[1]["id"], 7)
	})
}

func Test_Model_WhereOrIn(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件或包含("id", g.Slice别名{1, 2, 3, 4}).X条件或包含("id", g.Slice别名{3, 4, 5}).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 5)
		t.Assert(result[0]["id"], 1)
		t.Assert(result[4]["id"], 5)
	})
}

func Test_Model_WhereOrNotIn(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件或不包含("id", g.Slice别名{1, 2, 3, 4}).X条件或不包含("id", g.Slice别名{3, 4, 5}).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 8)
		t.Assert(result[0]["id"], 1)
		t.Assert(result[4]["id"], 7)
	})
}

func Test_Model_WhereBetween(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件取范围("id", 1, 4).X条件取范围("id", 3, 5).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 2)
		t.Assert(result[0]["id"], 3)
		t.Assert(result[1]["id"], 4)
	})
}

func Test_Model_WhereNotBetween(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件取范围以外("id", 2, 8).X条件取范围以外("id", 3, 100).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 1)
		t.Assert(result[0]["id"], 1)
	})
}

func Test_Model_WhereOrBetween(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件或取范围("id", 1, 4).X条件或取范围("id", 3, 5).X排序Desc("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 5)
		t.Assert(result[0]["id"], 5)
		t.Assert(result[4]["id"], 1)
	})
}

func Test_Model_WhereOrNotBetween(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	// db.SetDebug(true)
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件或取范围以外("id", 1, 4).X条件或取范围以外("id", 3, 5).X排序Desc("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 8)
		t.Assert(result[0]["id"], 10)
		t.Assert(result[4]["id"], 6)
	})
}

func Test_Model_WhereLike(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件模糊匹配("nickname", "name%").X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), TableSize)
		t.Assert(result[0]["id"], 1)
		t.Assert(result[TableSize-1]["id"], TableSize)
	})
}

func Test_Model_WhereNotLike(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件模糊匹配以外("nickname", "name%").X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 0)
	})
}

func Test_Model_WhereOrLike(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件或模糊匹配("nickname", "namexxx%").X条件或模糊匹配("nickname", "name%").X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), TableSize)
		t.Assert(result[0]["id"], 1)
		t.Assert(result[TableSize-1]["id"], TableSize)
	})
}

func Test_Model_WhereOrNotLike(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件或模糊匹配以外("nickname", "namexxx%").X条件或模糊匹配以外("nickname", "name%").X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), TableSize)
		t.Assert(result[0]["id"], 1)
		t.Assert(result[TableSize-1]["id"], TableSize)
	})
}

func Test_Model_WhereNull(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件NULL值("nickname").X条件NULL值("passport").X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 0)
	})
}

func Test_Model_WhereNotNull(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件非Null("nickname").X条件非Null("passport").X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), TableSize)
		t.Assert(result[0]["id"], 1)
		t.Assert(result[TableSize-1]["id"], TableSize)
	})
}

func Test_Model_WhereOrNotNull(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件或非Null("nickname").X条件或非Null("passport").X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), TableSize)
		t.Assert(result[0]["id"], 1)
		t.Assert(result[TableSize-1]["id"], TableSize)
	})
}

func Test_Model_WhereLT(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件小于("id", 3).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 2)
		t.Assert(result[0]["id"], 1)
	})
}

func Test_Model_WhereLTE(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件小于等于("id", 3).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["id"], 1)
	})
}

func Test_Model_WhereGT(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件大于("id", 8).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 2)
		t.Assert(result[0]["id"], 9)
	})
}

func Test_Model_WhereGTE(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件大于等于("id", 8).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["id"], 8)
	})
}

func Test_Model_WhereOrLT(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件小于("id", 3).X条件或小于("id", 4).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["id"], 1)
		t.Assert(result[2]["id"], 3)
	})
}

func Test_Model_WhereOrLTE(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件小于等于("id", 3).X条件或小于等于("id", 4).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 4)
		t.Assert(result[0]["id"], 1)
		t.Assert(result[3]["id"], 4)
	})
}

func Test_Model_WhereOrGT(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件大于("id", 8).X条件或大于("id", 7).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["id"], 8)
	})
}

func Test_Model_WhereOrGTE(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件大于等于("id", 8).X条件或大于等于("id", 7).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 4)
		t.Assert(result[0]["id"], 7)
	})
}

func Test_Model_Min_Max_Avg_Sum(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X查询最小值("id")
		t.AssertNil(err)
		t.Assert(result, 1)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X查询最大值("id")
		t.AssertNil(err)
		t.Assert(result, TableSize)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X查询平均值("id")
		t.AssertNil(err)
		t.Assert(result, 5.5)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X查询求和("id")
		t.AssertNil(err)
		t.Assert(result, 55)
	})
}

func Test_Model_CountColumn(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X查询字段行数("id")
		t.AssertNil(err)
		t.Assert(result, TableSize)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件包含("id", g.Slice别名{1, 2, 3}).X查询字段行数("id")
		t.AssertNil(err)
		t.Assert(result, 3)
	})
}

func Test_Model_InsertAndGetId(t *testing.T) {
	table := createTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		id, err := db.X创建Model对象(table).X设置数据(g.Map{
			"id":          1,
			"passport":    "user_1",
			"password":    "pass_1",
			"nickname":    "name_1",
			"create_time": gtime.X创建并按当前时间().String(),
		}).X插入并取ID()
		t.AssertNil(err)
		t.Assert(id, 1)
	})
	gtest.C(t, func(t *gtest.T) {
		id, err := db.X创建Model对象(table).X设置数据(g.Map{
			"passport":    "user_2",
			"password":    "pass_2",
			"nickname":    "name_2",
			"create_time": gtime.X创建并按当前时间().String(),
		}).X插入并取ID()
		t.AssertNil(err)
		t.Assert(id, 2)
	})
}

func Test_Model_Increment_Decrement(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("id", 1).X更新增量("id", 100)
		t.AssertNil(err)
		rows, _ := result.RowsAffected()
		t.Assert(rows, 1)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("id", 101).X更新减量("id", 10)
		t.AssertNil(err)
		rows, _ := result.RowsAffected()
		t.Assert(rows, 1)
	})
	gtest.C(t, func(t *gtest.T) {
		count, err := db.X创建Model对象(table).X条件("id", 91).X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(1))
	})
}

func Test_Model_Raw(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		all, err := db.
			X原生SQL(fmt.Sprintf("select * from %s where id in (?)", table), g.Slice别名{1, 5, 7, 8, 9, 10}).
			X条件小于("id", 8).
			X条件包含("id", g.Slice别名{1, 2, 3, 4, 5, 6, 7}).
			X排序Desc("id").
			X设置条数(2).
			X查询()
		t.AssertNil(err)
		t.Assert(len(all), 2)
		t.Assert(all[0]["id"], 7)
		t.Assert(all[1]["id"], 5)
	})

	gtest.C(t, func(t *gtest.T) {
		count, err := db.
			X原生SQL(fmt.Sprintf("select * from %s where id in (?)", table), g.Slice别名{1, 5, 7, 8, 9, 10}).
			X条件小于("id", 8).
			X条件包含("id", g.Slice别名{1, 2, 3, 4, 5, 6, 7}).
			X排序Desc("id").
			X设置条数(2).
			X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(6))
	})
}

func Test_Model_Handler(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		m := db.X创建Model对象(table).X链式安全().X处理函数(
			func(m *gdb.Model) *gdb.Model {
				return m.X设置分页(0, 3)
			},
			func(m *gdb.Model) *gdb.Model {
				return m.X条件("id", g.Slice别名{1, 2, 3, 4, 5, 6})
			},
			func(m *gdb.Model) *gdb.Model {
				return m.X排序Desc("id")
			},
		)
		all, err := m.X查询()
		t.AssertNil(err)
		t.Assert(len(all), 3)
		t.Assert(all[0]["id"], 6)
		t.Assert(all[2]["id"], 4)
	})
}

func Test_Model_FieldCount(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		all, err := db.X创建Model对象(table).X字段保留过滤("id").X字段追加计数("id", "total").X排序分组("id").X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(all), TableSize)
		t.Assert(all[0]["id"], 1)
		t.Assert(all[0]["total"].X取整数(), 1)
	})
}

func Test_Model_FieldMax(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		all, err := db.X创建Model对象(table).X字段保留过滤("id").X字段追加最大值("id", "total").X排序分组("id").X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(all), TableSize)
		t.Assert(all[0]["id"], 1)
		t.Assert(all[0]["total"].X取整数(), 1)
	})
}

func Test_Model_FieldMin(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		all, err := db.X创建Model对象(table).X字段保留过滤("id").X字段追加最小值("id", "total").X排序分组("id").X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(all), TableSize)
		t.Assert(all[0]["id"], 1)
		t.Assert(all[0]["total"].X取整数(), 1)
	})
}

func Test_Model_FieldAvg(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		all, err := db.X创建Model对象(table).X字段保留过滤("id").X字段追加平均值("id", "total").X排序分组("id").X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(all), TableSize)
		t.Assert(all[0]["id"], 1)
		t.Assert(all[0]["total"].X取整数(), 1)
	})
}

func Test_Model_OmitEmptyWhere(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	// Basic type where.
	gtest.C(t, func(t *gtest.T) {
		count, err := db.X创建Model对象(table).X条件("id", 0).X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(0))
	})
	gtest.C(t, func(t *gtest.T) {
		count, err := db.X创建Model对象(table).X过滤空值条件().X条件("id", 0).X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(TableSize))
	})
	gtest.C(t, func(t *gtest.T) {
		count, err := db.X创建Model对象(table).X过滤空值条件().X条件("id", 0).X条件("nickname", "").X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(TableSize))
	})
	// Slice where.
	gtest.C(t, func(t *gtest.T) {
		count, err := db.X创建Model对象(table).X条件("id", g.Slice别名{1, 2, 3}).X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(3))
	})
	gtest.C(t, func(t *gtest.T) {
		count, err := db.X创建Model对象(table).X条件("id", g.Slice别名{}).X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(0))
	})
	gtest.C(t, func(t *gtest.T) {
		count, err := db.X创建Model对象(table).X过滤空值条件().X条件("id", g.Slice别名{}).X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(TableSize))
	})
	gtest.C(t, func(t *gtest.T) {
		count, err := db.X创建Model对象(table).X条件("id", g.Slice别名{}).X过滤空值条件().X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(TableSize))
	})
	// Struct Where.
	gtest.C(t, func(t *gtest.T) {
		type Input struct {
			Id   []int
			Name []string
		}
		count, err := db.X创建Model对象(table).X条件(Input{}).X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(0))
	})
	gtest.C(t, func(t *gtest.T) {
		type Input struct {
			Id   []int
			Name []string
		}
		count, err := db.X创建Model对象(table).X条件(Input{}).X过滤空值条件().X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(TableSize))
	})
	// Map Where.
	gtest.C(t, func(t *gtest.T) {
		count, err := db.X创建Model对象(table).X条件(g.Map{
			"id":       []int{},
			"nickname": []string{},
		}).X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(0))
	})
	gtest.C(t, func(t *gtest.T) {
		count, err := db.X创建Model对象(table).X条件(g.Map{
			"id": []int{},
		}).X过滤空值条件().X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(TableSize))
	})
}

// 这段注释链接指向的是GitHub上的一个Issue，GF（Go Foundation）是一个Go语言的库或框架。"1387"可能是Issue的编号。具体的内容需要查看该链接才能得知，大致意思是关于GF项目在1387号问题上的讨论、报告了一个错误或者提出了一个特性请求。 md5:7c877c3e7a856cb1
func Test_Model_GTime_DefaultValue(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			Nickname   string
			CreateTime *gtime.Time
		}
		data := User{
			Id:         1,
			Passport:   "user_1",
			Password:   "pass_1",
			Nickname:   "name_1",
			CreateTime: gtime.X创建并按当前时间(),
		}
		// Insert
		_, err := db.X创建Model对象(table).X设置数据(data).X插入()
		t.AssertNil(err)

		// Select
		var (
			user *User
		)
		err = db.X创建Model对象(table).X查询到结构体指针(&user)
		t.AssertNil(err)
		t.Assert(user.Passport, data.Passport)
		t.Assert(user.Password, data.Password)
		t.Assert(user.CreateTime, data.CreateTime)
		t.Assert(user.Nickname, data.Nickname)

		// Insert
		user.Id = 2
		_, err = db.X创建Model对象(table).X设置数据(user).X插入()
		t.AssertNil(err)
	})
}

// 使用过滤器不会影响函数内部的外部值。 md5:857585fd480ebfc6
func Test_Model_Insert_Filter(t *testing.T) {
	// map
	gtest.C(t, func(t *gtest.T) {
		table := createTable()
		defer dropTable(table)
		data := g.Map{
			"id":          1,
			"uid":         1,
			"passport":    "t1",
			"password":    "25d55ad283aa400af464c76d713c07ad",
			"nickname":    "name_1",
			"create_time": gtime.X创建并按当前时间().String(),
		}
		result, err := db.X创建Model对象(table).X设置数据(data).X插入()
		t.AssertNil(err)
		n, _ := result.LastInsertId()
		t.Assert(n, 1)

		t.Assert(data["uid"], 1)
	})
	// slice
	gtest.C(t, func(t *gtest.T) {
		table := createTable()
		defer dropTable(table)
		data := g.Map切片{
			g.Map{
				"id":          1,
				"uid":         1,
				"passport":    "t1",
				"password":    "25d55ad283aa400af464c76d713c07ad",
				"nickname":    "name_1",
				"create_time": gtime.X创建并按当前时间().String(),
			},
			g.Map{
				"id":          2,
				"uid":         2,
				"passport":    "t1",
				"password":    "25d55ad283aa400af464c76d713c07ad",
				"nickname":    "name_1",
				"create_time": gtime.X创建并按当前时间().String(),
			},
		}

		result, err := db.X创建Model对象(table).X设置数据(data).X插入()
		t.AssertNil(err)
		n, _ := result.LastInsertId()
		t.Assert(n, 2)

		t.Assert(data[0]["uid"], 1)
		t.Assert(data[1]["uid"], 2)
	})
}

func Test_Model_Embedded_Filter(t *testing.T) {
	table := createTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		type Base struct {
			Id         int
			Uid        int
			CreateTime string
			NoneExist  string
		}
		type User struct {
			Base
			Passport string
			Password string
			Nickname string
		}
		result, err := db.X创建Model对象(table).X设置数据(User{
			Passport: "john-test",
			Password: "123456",
			Nickname: "John",
			Base: Base{
				Id:         100,
				Uid:        100,
				CreateTime: gtime.X创建并按当前时间().String(),
			},
		}).X插入()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)

		var user *User
		err = db.X创建Model对象(table).X字段保留过滤(user).X条件("id=100").X查询到结构体指针(&user)
		t.AssertNil(err)
		t.Assert(user.Passport, "john-test")
		t.Assert(user.Id, 100)
	})
}

// 从GoFrame v1.16.0开始，此功能不再使用，因为过滤功能已自动启用。 md5:a491426db314e6d6
func Test_Model_Insert_KeyFieldNameMapping_Error(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id             int
			Passport       string
			Password       string
			Nickname       string
			CreateTime     string
			NoneExistField string
		}
		data := User{
			Id:         1,
			Passport:   "user_1",
			Password:   "pass_1",
			Nickname:   "name_1",
			CreateTime: "2020-10-10 12:00:01",
		}
		_, err := db.X创建Model对象(table).X设置数据(data).X插入()
		t.AssertNil(err)
	})
}

func Test_Model_Fields_AutoFilterInJoinStatement(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var err error
		table1 := "user"
		table2 := "score"
		table3 := "info"
		if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			id INTEGER	PRIMARY KEY AUTOINCREMENT
						UNIQUE
						NOT NULL,
			name varchar(500) NOT NULL DEFAULT ''
		);
		`, table1,
		)); err != nil {
			t.AssertNil(err)
		}
		defer dropTable(table1)
		_, err = db.X创建Model对象(table1).X插入(g.Map{
			"id":   1,
			"name": "john",
		})
		t.AssertNil(err)

		if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			id INTEGER	PRIMARY KEY AUTOINCREMENT
						UNIQUE
						NOT NULL,
			user_id int(11) NOT NULL DEFAULT 0,
			number varchar(500) NOT NULL DEFAULT ''
		);
	    `, table2,
		)); err != nil {
			t.AssertNil(err)
		}
		defer dropTable(table2)
		_, err = db.X创建Model对象(table2).X插入(g.Map{
			"id":      1,
			"user_id": 1,
			"number":  "n",
		})
		t.AssertNil(err)

		if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			id INTEGER	PRIMARY KEY AUTOINCREMENT
						UNIQUE
						NOT NULL,
			user_id int(11) NOT NULL DEFAULT 0,
			description varchar(500) NOT NULL DEFAULT ''
		);
		`, table3,
		)); err != nil {
			t.AssertNil(err)
		}
		defer dropTable(table3)
		_, err = db.X创建Model对象(table3).X插入(g.Map{
			"id":          1,
			"user_id":     1,
			"description": "brief",
		})
		t.AssertNil(err)

		one, err := db.X创建Model对象("user").
			X条件("user.id", 1).
			X字段保留过滤("score.number,user.name").
			X左连接("score", "user.id=score.user_id").
			X左连接("info", "info.id=info.user_id").
			X排序("user.id asc").
			X查询一条()
		t.AssertNil(err)
		t.Assert(len(one), 2)
		t.Assert(one["name"].String(), "john")
		t.Assert(one["number"].String(), "n")

		one, err = db.X创建Model对象("user").
			X左连接("score", "user.id=score.user_id").
			X左连接("info", "info.id=info.user_id").
			X字段保留过滤("score.number,user.name").
			X查询一条()
		t.AssertNil(err)
		t.Assert(len(one), 2)
		t.Assert(one["name"].String(), "john")
		t.Assert(one["number"].String(), "n")
	})
}

func Test_Model_WherePrefix(t *testing.T) {
	var (
		table1 = "table1_" + gtime.X取文本时间戳纳秒()
		table2 = "table2_" + gtime.X取文本时间戳纳秒()
	)
	createInitTable(table1)
	defer dropTable(table1)
	createInitTable(table2)
	defer dropTable(table2)

	gtest.C(t, func(t *gtest.T) {
		r, err := db.X创建Model对象(table1).
			X字段保留过滤并带前缀(table1, "*").
			X左连接相同字段(table2, "id").
			X条件带前缀(table2, g.Map{
				"id": g.Slice别名{1, 2},
			}).
			X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(r), 2)
		t.Assert(r[0]["id"], "1")
		t.Assert(r[1]["id"], "2")
	})
}

func Test_Model_WhereOrPrefix(t *testing.T) {
	var (
		table1 = "table1_" + gtime.X取文本时间戳纳秒()
		table2 = "table2_" + gtime.X取文本时间戳纳秒()
	)
	createInitTable(table1)
	defer dropTable(table1)
	createInitTable(table2)
	defer dropTable(table2)

	gtest.C(t, func(t *gtest.T) {
		r, err := db.X创建Model对象(table1).
			X字段保留过滤并带前缀(table1, "*").
			X左连接相同字段(table2, "id").
			X条件或并带前缀(table1, g.Map{
				"id": g.Slice别名{1, 2},
			}).
			X条件或并带前缀(table2, g.Map{
				"id": g.Slice别名{8, 9},
			}).
			X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(r), 4)
		t.Assert(r[0]["id"], "1")
		t.Assert(r[1]["id"], "2")
		t.Assert(r[2]["id"], "8")
		t.Assert(r[3]["id"], "9")
	})
}

func Test_Model_WherePrefixLike(t *testing.T) {
	var (
		table1 = "table1_" + gtime.X取文本时间戳纳秒()
		table2 = "table2_" + gtime.X取文本时间戳纳秒()
	)
	createInitTable(table1)
	defer dropTable(table1)
	createInitTable(table2)
	defer dropTable(table2)

	gtest.C(t, func(t *gtest.T) {
		r, err := db.X创建Model对象(table1).
			X字段保留过滤并带前缀(table1, "*").
			X左连接相同字段(table2, "id").
			X条件带前缀(table1, g.Map{
				"id": g.Slice别名{1, 2, 3},
			}).
			X条件带前缀(table2, g.Map{
				"id": g.Slice别名{3, 4, 5},
			}).
			X条件模糊匹配并带前缀(table2, "nickname", "name%").
			X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(r), 1)
		t.Assert(r[0]["id"], "3")
	})
}

// 这段注释链接指向的是GitHub上的一个 issues（问题或讨论），来自gogf（GoGF）项目。它表示这个注释与 issue #1159 相关，可能是对某个特定问题、错误报告、功能请求或者讨论的引用。具体的内容需要查看该issue页面以获取详细信息。 md5:ef2c3285217b52b1
func Test_ScanList_NoRecreate_PtrAttribute(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		type S1 struct {
			Id    int
			Name  string
			Age   int
			Score int
		}
		type S3 struct {
			One *S1
		}
		var (
			s   []*S3
			err error
		)
		r1 := gdb.Result{
			gdb.Record{
				"id":   gvar.X创建(1),
				"name": gvar.X创建("john"),
				"age":  gvar.X创建(16),
			},
			gdb.Record{
				"id":   gvar.X创建(2),
				"name": gvar.X创建("smith"),
				"age":  gvar.X创建(18),
			},
		}
		err = r1.X取指针列表(&s, "One")
		t.AssertNil(err)
		t.Assert(len(s), 2)
		t.Assert(s[0].One.Name, "john")
		t.Assert(s[0].One.Age, 16)
		t.Assert(s[1].One.Name, "smith")
		t.Assert(s[1].One.Age, 18)

		r2 := gdb.Result{
			gdb.Record{
				"id":  gvar.X创建(1),
				"age": gvar.X创建(20),
			},
			gdb.Record{
				"id":  gvar.X创建(2),
				"age": gvar.X创建(21),
			},
		}
		err = r2.X取指针列表(&s, "One", "One", "id:Id")
		t.AssertNil(err)
		t.Assert(len(s), 2)
		t.Assert(s[0].One.Name, "john")
		t.Assert(s[0].One.Age, 20)
		t.Assert(s[1].One.Name, "smith")
		t.Assert(s[1].One.Age, 21)
	})
}

// 这段注释链接指向的是GitHub上的一个 issues（问题或讨论），来自gogf（GoGF）项目。它表示这个注释与 issue #1159 相关，可能是对某个特定问题、错误报告、功能请求或者讨论的引用。具体的内容需要查看该issue页面以获取详细信息。 md5:ef2c3285217b52b1
func Test_ScanList_NoRecreate_StructAttribute(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		type S1 struct {
			Id    int
			Name  string
			Age   int
			Score int
		}
		type S3 struct {
			One S1
		}
		var (
			s   []*S3
			err error
		)
		r1 := gdb.Result{
			gdb.Record{
				"id":   gvar.X创建(1),
				"name": gvar.X创建("john"),
				"age":  gvar.X创建(16),
			},
			gdb.Record{
				"id":   gvar.X创建(2),
				"name": gvar.X创建("smith"),
				"age":  gvar.X创建(18),
			},
		}
		err = r1.X取指针列表(&s, "One")
		t.AssertNil(err)
		t.Assert(len(s), 2)
		t.Assert(s[0].One.Name, "john")
		t.Assert(s[0].One.Age, 16)
		t.Assert(s[1].One.Name, "smith")
		t.Assert(s[1].One.Age, 18)

		r2 := gdb.Result{
			gdb.Record{
				"id":  gvar.X创建(1),
				"age": gvar.X创建(20),
			},
			gdb.Record{
				"id":  gvar.X创建(2),
				"age": gvar.X创建(21),
			},
		}
		err = r2.X取指针列表(&s, "One", "One", "id:Id")
		t.AssertNil(err)
		t.Assert(len(s), 2)
		t.Assert(s[0].One.Name, "john")
		t.Assert(s[0].One.Age, 20)
		t.Assert(s[1].One.Name, "smith")
		t.Assert(s[1].One.Age, 21)
	})
}

// 这段注释链接指向的是GitHub上的一个 issues（问题或讨论），来自gogf（GoGF）项目。它表示这个注释与 issue #1159 相关，可能是对某个特定问题、错误报告、功能请求或者讨论的引用。具体的内容需要查看该issue页面以获取详细信息。 md5:ef2c3285217b52b1
func Test_ScanList_NoRecreate_SliceAttribute_Ptr(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		type S1 struct {
			Id    int
			Name  string
			Age   int
			Score int
		}
		type S2 struct {
			Id    int
			Pid   int
			Name  string
			Age   int
			Score int
		}
		type S3 struct {
			One  *S1
			Many []*S2
		}
		var (
			s   []*S3
			err error
		)
		r1 := gdb.Result{
			gdb.Record{
				"id":   gvar.X创建(1),
				"name": gvar.X创建("john"),
				"age":  gvar.X创建(16),
			},
			gdb.Record{
				"id":   gvar.X创建(2),
				"name": gvar.X创建("smith"),
				"age":  gvar.X创建(18),
			},
		}
		err = r1.X取指针列表(&s, "One")
		t.AssertNil(err)
		t.Assert(len(s), 2)
		t.Assert(s[0].One.Name, "john")
		t.Assert(s[0].One.Age, 16)
		t.Assert(s[1].One.Name, "smith")
		t.Assert(s[1].One.Age, 18)

		r2 := gdb.Result{
			gdb.Record{
				"id":   gvar.X创建(100),
				"pid":  gvar.X创建(1),
				"age":  gvar.X创建(30),
				"name": gvar.X创建("john"),
			},
			gdb.Record{
				"id":   gvar.X创建(200),
				"pid":  gvar.X创建(1),
				"age":  gvar.X创建(31),
				"name": gvar.X创建("smith"),
			},
		}
		err = r2.X取指针列表(&s, "Many", "One", "pid:Id")
		// fmt.Printf("%+v", err)
		t.AssertNil(err)
		t.Assert(len(s), 2)
		t.Assert(s[0].One.Name, "john")
		t.Assert(s[0].One.Age, 16)
		t.Assert(len(s[0].Many), 2)
		t.Assert(s[0].Many[0].Name, "john")
		t.Assert(s[0].Many[0].Age, 30)
		t.Assert(s[0].Many[1].Name, "smith")
		t.Assert(s[0].Many[1].Age, 31)

		t.Assert(s[1].One.Name, "smith")
		t.Assert(s[1].One.Age, 18)
		t.Assert(len(s[1].Many), 0)

		r3 := gdb.Result{
			gdb.Record{
				"id":  gvar.X创建(100),
				"pid": gvar.X创建(1),
				"age": gvar.X创建(40),
			},
			gdb.Record{
				"id":  gvar.X创建(200),
				"pid": gvar.X创建(1),
				"age": gvar.X创建(41),
			},
		}
		err = r3.X取指针列表(&s, "Many", "One", "pid:Id")
		// fmt.Printf("%+v", err)
		t.AssertNil(err)
		t.Assert(len(s), 2)
		t.Assert(s[0].One.Name, "john")
		t.Assert(s[0].One.Age, 16)
		t.Assert(len(s[0].Many), 2)
		t.Assert(s[0].Many[0].Name, "john")
		t.Assert(s[0].Many[0].Age, 40)
		t.Assert(s[0].Many[1].Name, "smith")
		t.Assert(s[0].Many[1].Age, 41)

		t.Assert(s[1].One.Name, "smith")
		t.Assert(s[1].One.Age, 18)
		t.Assert(len(s[1].Many), 0)
	})
}

// 这段注释链接指向的是GitHub上的一个 issues（问题或讨论），来自gogf（GoGF）项目。它表示这个注释与 issue #1159 相关，可能是对某个特定问题、错误报告、功能请求或者讨论的引用。具体的内容需要查看该issue页面以获取详细信息。 md5:ef2c3285217b52b1
func Test_ScanList_NoRecreate_SliceAttribute_Struct(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		type S1 struct {
			Id    int
			Name  string
			Age   int
			Score int
		}
		type S2 struct {
			Id    int
			Pid   int
			Name  string
			Age   int
			Score int
		}
		type S3 struct {
			One  S1
			Many []S2
		}
		var (
			s   []S3
			err error
		)
		r1 := gdb.Result{
			gdb.Record{
				"id":   gvar.X创建(1),
				"name": gvar.X创建("john"),
				"age":  gvar.X创建(16),
			},
			gdb.Record{
				"id":   gvar.X创建(2),
				"name": gvar.X创建("smith"),
				"age":  gvar.X创建(18),
			},
		}
		err = r1.X取指针列表(&s, "One")
		t.AssertNil(err)
		t.Assert(len(s), 2)
		t.Assert(s[0].One.Name, "john")
		t.Assert(s[0].One.Age, 16)
		t.Assert(s[1].One.Name, "smith")
		t.Assert(s[1].One.Age, 18)

		r2 := gdb.Result{
			gdb.Record{
				"id":   gvar.X创建(100),
				"pid":  gvar.X创建(1),
				"age":  gvar.X创建(30),
				"name": gvar.X创建("john"),
			},
			gdb.Record{
				"id":   gvar.X创建(200),
				"pid":  gvar.X创建(1),
				"age":  gvar.X创建(31),
				"name": gvar.X创建("smith"),
			},
		}
		err = r2.X取指针列表(&s, "Many", "One", "pid:Id")
		// fmt.Printf("%+v", err)
		t.AssertNil(err)
		t.Assert(len(s), 2)
		t.Assert(s[0].One.Name, "john")
		t.Assert(s[0].One.Age, 16)
		t.Assert(len(s[0].Many), 2)
		t.Assert(s[0].Many[0].Name, "john")
		t.Assert(s[0].Many[0].Age, 30)
		t.Assert(s[0].Many[1].Name, "smith")
		t.Assert(s[0].Many[1].Age, 31)

		t.Assert(s[1].One.Name, "smith")
		t.Assert(s[1].One.Age, 18)
		t.Assert(len(s[1].Many), 0)

		r3 := gdb.Result{
			gdb.Record{
				"id":  gvar.X创建(100),
				"pid": gvar.X创建(1),
				"age": gvar.X创建(40),
			},
			gdb.Record{
				"id":  gvar.X创建(200),
				"pid": gvar.X创建(1),
				"age": gvar.X创建(41),
			},
		}
		err = r3.X取指针列表(&s, "Many", "One", "pid:Id")
		// fmt.Printf("%+v", err)
		t.AssertNil(err)
		t.Assert(len(s), 2)
		t.Assert(s[0].One.Name, "john")
		t.Assert(s[0].One.Age, 16)
		t.Assert(len(s[0].Many), 2)
		t.Assert(s[0].Many[0].Name, "john")
		t.Assert(s[0].Many[0].Age, 40)
		t.Assert(s[0].Many[1].Name, "smith")
		t.Assert(s[0].Many[1].Age, 41)

		t.Assert(s[1].One.Name, "smith")
		t.Assert(s[1].One.Age, 18)
		t.Assert(len(s[1].Many), 0)
	})
}

func TestResult_Structs1(t *testing.T) {
	type A struct {
		Id int `orm:"id"`
	}
	type B struct {
		*A
		Name string
	}
	gtest.C(t, func(t *gtest.T) {
		r := gdb.Result{
			gdb.Record{"id": gvar.X创建(nil), "name": gvar.X创建("john")},
			gdb.Record{"id": gvar.X创建(nil), "name": gvar.X创建("smith")},
		}
		array := make([]*B, 2)
		err := r.X取切片结构体指针(&array)
		t.AssertNil(err)
		t.Assert(array[0].Id, 0)
		t.Assert(array[1].Id, 0)
		t.Assert(array[0].Name, "john")
		t.Assert(array[1].Name, "smith")
	})
}
