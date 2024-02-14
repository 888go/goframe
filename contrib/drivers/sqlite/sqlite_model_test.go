// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package sqlite_test

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"
	
	"github.com/888go/goframe/container/garray"
	"github.com/888go/goframe/container/gmap"
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/database/gdb"
	"github.com/888go/goframe/encoding/gjson"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/os/glog"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/guid"
	"github.com/888go/goframe/util/gutil"
)

func Test_Model_Insert(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		user := db.X创建Model对象(table)
		result, err := user.X设置数据(g.Map{
			"id":          1,
			"uid":         1,
			"passport":    "t1",
			"password":    "25d55ad283aa400af464c76d713c07ad",
			"nickname":    "name_1",
			"create_time": 时间类.X创建并按当前时间().String(),
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
			"create_time": 时间类.X创建并按当前时间().String(),
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
			CreateTime *时间类.Time `json:"create_time"`
		}
		// Model inserting.
		result, err = db.X创建Model对象(table).X设置数据(User{
			Id:         3,
			Uid:        3,
			Passport:   "t3",
			Password:   "25d55ad283aa400af464c76d713c07ad",
			Nickname:   "name_3",
			CreateTime: 时间类.X创建并按当前时间(),
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
			CreateTime: 时间类.X创建并按当前时间(),
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

// 解决问题：https://github.com/gogf/gf/issues/819
func Test_Model_Insert_WithStructAndSliceAttribute(t *testing.T) {
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
		_, err := db.X创建Model对象(table).X设置数据(data).X插入()
		t.AssertNil(err)

		one, err := db.X创建Model对象(table).X查询一条("id", 1)
		t.AssertNil(err)
		t.Assert(one["passport"], data["passport"])
		t.Assert(one["create_time"], data["create_time"])
		t.Assert(one["nickname"], json类.X创建(data["nickname"]).X取json字节集PANI())
	})
}

func Test_Model_Insert_KeyFieldNameMapping(t *testing.T) {
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

	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
		user := db.X创建Model对象(table)
		array := 数组类.X创建()
		for i := 1; i <= TableSize; i++ {
			array.Append别名(g.Map{
				"id":          i,
				"uid":         i,
				"passport":    fmt.Sprintf("t%d", i),
				"password":    "25d55ad283aa400af464c76d713c07ad",
				"nickname":    fmt.Sprintf("name_%d", i),
				"create_time": 时间类.X创建并按当前时间().String(),
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
	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
		table := createTable()
		defer dropTable(table)
		result, err := db.X创建Model对象(table).X设置数据(g.Map数组{
			{
				"id":          2,
				"uid":         2,
				"passport":    "t2",
				"password":    "25d55ad283aa400af464c76d713c07ad",
				"nickname":    "name_2",
				"create_time": 时间类.X创建并按当前时间().String(),
			},
			{
				"id":          3,
				"uid":         3,
				"passport":    "t3",
				"password":    "25d55ad283aa400af464c76d713c07ad",
				"nickname":    "name_3",
				"create_time": 时间类.X创建并按当前时间().String(),
			},
		}).X设置批量操作行数(1).X插入()
		if err != nil {
			单元测试类.Error(err)
		}
		n, _ := result.RowsAffected()
		t.Assert(n, 2)
	})

	// 批量插入数据，并获取最后插入的自增ID。
	单元测试类.C(t, func(t *单元测试类.T) {
		table := createTable()
		defer dropTable(table)
		result, err := db.X创建Model对象(table).X设置数据(g.Map数组{
			{"passport": "t1", "password": "25d55ad283aa400af464c76d713c07ad", "nickname": "name", "create_time": 时间类.X创建并按当前时间().String()},
			{"passport": "t2", "password": "25d55ad283aa400af464c76d713c07ad", "nickname": "name", "create_time": 时间类.X创建并按当前时间().String()},
			{"passport": "t3", "password": "25d55ad283aa400af464c76d713c07ad", "nickname": "name", "create_time": 时间类.X创建并按当前时间().String()},
			{"passport": "t4", "password": "25d55ad283aa400af464c76d713c07ad", "nickname": "name", "create_time": 时间类.X创建并按当前时间().String()},
			{"passport": "t5", "password": "25d55ad283aa400af464c76d713c07ad", "nickname": "name", "create_time": 时间类.X创建并按当前时间().String()},
		}).X设置批量操作行数(2).X插入()
		if err != nil {
			单元测试类.Error(err)
		}
		n, _ := result.RowsAffected()
		t.Assert(n, 5)
	})

	// batch replace
	单元测试类.C(t, func(t *单元测试类.T) {
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

	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
		_, err := db.X创建Model对象(table).X设置数据(g.Map{
			"id":          1,
			"passport":    "t111",
			"password":    "25d55ad283aa400af464c76d713c07ad",
			"nickname":    "T111",
			"create_time": CreateTime,
		}).X插入并更新已存在()
		t.Assert(err, ErrorSave)
	})
}

func Test_Model_Update(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
// UPDATE...LIMIT
// 使用gtest进行单元测试，参数t为测试实例
// 根据条件更新表格中的一条记录，并限制只更新两条数据
// 更新table表中nickname为"T100"且满足Where条件(如id=1)的前两条记录
// 
// gtest.C(t, func(t *gtest.T) {
//   // 执行数据库更新操作，返回结果result和可能发生的错误err
//   result, err := db.Model(table).Data("nickname", "T100").Where(1).Limit(2).Update()
//   // 断言错误应为nil，即无错误发生
//   t.AssertNil(err)
//   // 获取更新影响的行数，_表示忽略可能的错误信息
//   n, _ := result.RowsAffected()
//   // 断言实际影响的行数应为2
//   t.Assert(n, 2)
// }

// 使用db.Model方法设置数据表为table，仅查询nickname字段，并根据id为10的条件进行查询，获取查询结果赋值给v1变量，同时返回可能的错误信息并存储在err变量中
// t.AssertNil(err) 表示断言err应为空（即没有错误发生）
// t.Assert(v1.String(), "T100") 表示断言v1变量转换为字符串后与"T100"相等

// 获取表名为table的数据库模型，指定查询字段为nickname，条件为id为8的记录的值
// 并将查询结果赋值给变量v2，同时返回可能的错误信息err
// 
// v2, err := db.Model(table).Fields("nickname").Where("id", 8).Value()
// 断言错误err为nil，即查询过程中没有出现错误
// t.AssertNil(err)
// 断言v2转换为字符串后的结果为"name_8"
// t.Assert(v2.String(), "name_8")
// }) 表示闭合的测试用例或匿名函数

	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X设置数据("passport", "user_22").X条件("passport=?", "user_2").X更新()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X设置数据("passport", "user_2").X条件("passport='user_22'").X更新()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
	})

	// 更新 + 数据(字符串)
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X设置数据("passport='user_33'").X条件("passport='user_3'").X更新()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
	})
	// Update + Fields(string)
// 更新 + 字段(string)
	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
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

	单元测试类.C(t, func(t *单元测试类.T) {
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

	单元测试类.C(t, func(t *单元测试类.T) {
		md := db.X创建Model对象(table).X链式安全(false).X条件("id IN(?)", g.Slice别名{1, 3})
		count, err := md.X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(2))

		md.X条件("id = ?", 1)
		count, err = md.X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(1))
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		md := db.X创建Model对象(table).X链式安全(true).X条件("id IN(?)", g.Slice别名{1, 3})
		count, err := md.X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(2))

		md.X条件("id = ?", 1)
		count, err = md.X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(2))
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		md := db.X创建Model对象(table).X链式安全().X条件("id IN(?)", g.Slice别名{1, 3})
		count, err := md.X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(2))

		md.X条件("id = ?", 1)
		count, err = md.X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(2))
	})
	单元测试类.C(t, func(t *单元测试类.T) {
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

	单元测试类.C(t, func(t *单元测试类.T) {
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

	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X查询()
		t.AssertNil(err)
		t.Assert(len(result), TableSize)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件("id<0").X查询()
		t.Assert(result, nil)
		t.AssertNil(err)
	})
}

func Test_Model_AllAndCount(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	tableName2 := "user_" + 时间类.X创建并按当前时间().X取文本时间戳纳秒()
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
		单元测试类.AssertNil(err)
	}
	defer dropTable(tableName2)
	r, err := db.X插入(ctx, tableName2, g.Map{
		"id":   1,
		"name": "table2_1",
		"age":  18,
	})
	单元测试类.AssertNil(err)
	n, _ := r.RowsAffected()
	单元测试类.Assert(n, 1)

	// AllAndCount 获取所有数据及计数
	单元测试类.C(t, func(t *单元测试类.T) {
		result, count, err := db.X创建Model对象(table).X查询与行数(false)
		t.AssertNil(err)
		t.Assert(len(result), TableSize)
		t.Assert(count, TableSize)
	})
	// AllAndCount 无数据时
	单元测试类.C(t, func(t *单元测试类.T) {
		result, count, err := db.X创建Model对象(table).X条件("id<0").X查询与行数(false)
		t.Assert(result, nil)
		t.AssertNil(err)
		t.Assert(count, 0)
	})
	// AllAndCount 带分页功能
	单元测试类.C(t, func(t *单元测试类.T) {
		result, count, err := db.X创建Model对象(table).X设置分页(1, 5).X查询与行数(false)
		t.AssertNil(err)
		t.Assert(len(result), 5)
		t.Assert(count, TableSize)
	})
	// AllAndCount 正常结果时的操作
	单元测试类.C(t, func(t *单元测试类.T) {
		result, count, err := db.X创建Model对象(table).X条件("id=?", 1).X查询与行数(false)
		t.AssertNil(err)
		t.Assert(count, 1)
		t.Assert(result[0]["id"], 1)
		t.Assert(result[0]["nickname"], "name_1")
		t.Assert(result[0]["passport"], "user_1")
	})
	// AllAndCount 带有 distinct 的功能
	单元测试类.C(t, func(t *单元测试类.T) {
		result, count, err := db.X创建Model对象(table).X字段保留过滤("DISTINCT nickname").X查询与行数(true)
		t.AssertNil(err)
		t.Assert(count, TableSize)
		t.Assert(result[0]["nickname"], "name_1")
		t.AssertNil(result[0]["id"])
	})
	// AllAndCount with Join
	单元测试类.C(t, func(t *单元测试类.T) {
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
	// AllAndCount 使用 Join 并返回 CodeDbOperationError
	单元测试类.C(t, func(t *单元测试类.T) {
		all, count, err := db.X创建Model对象(table).X设置表别名("u1").
			X左连接(tableName2, "u2", "u2.id=u1.id").
			X字段保留过滤("u1.passport,u1.id,u2.name,u2.age").
			X条件("u1.id<2").
			X查询与行数(true)
		t.AssertNE(err, nil)
		t.AssertEQ(错误类.X取错误码(err), 错误码类.CodeDbOperationError)
		t.Assert(count, 0)
		t.Assert(all, nil)
	})
}

func Test_Model_Fields(t *testing.T) {
	tableName1 := createInitTable()
	defer dropTable(tableName1)

	tableName2 := "user_" + 时间类.X创建并按当前时间().X取文本时间戳纳秒()
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
		单元测试类.AssertNil(err)
	}
	defer dropTable(tableName2)

	r, err := db.X插入(ctx, tableName2, g.Map{
		"id":   1,
		"name": "table2_1",
		"age":  18,
	})
	单元测试类.AssertNil(err)
	n, _ := r.RowsAffected()
	单元测试类.Assert(n, 1)

	单元测试类.C(t, func(t *单元测试类.T) {
		all, err := db.X创建Model对象(tableName1).X设置表别名("u").X字段保留过滤("u.passport,u.id").X条件("u.id<2").X查询()
		t.AssertNil(err)
		t.Assert(len(all), 1)
		t.Assert(len(all[0]), 2)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		all, err := db.X创建Model对象(tableName1).X设置表别名("u1").
			X左连接(tableName1, "u2", "u2.id=u1.id").
			X字段保留过滤("u1.passport,u1.id,u2.id AS u2id").
			X条件("u1.id<2").
			X查询()
		t.AssertNil(err)
		t.Assert(len(all), 1)
		t.Assert(len(all[0]), 3)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
		record, err := db.X创建Model对象(table).X条件("id", 1).X查询一条()
		t.AssertNil(err)
		t.Assert(record["nickname"].String(), "name_1")
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		record, err := db.X创建Model对象(table).X条件("id", 0).X查询一条()
		t.AssertNil(err)
		t.Assert(record, nil)
	})
}

func Test_Model_Value(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		value, err := db.X创建Model对象(table).X字段保留过滤("nickname").X条件("id", 1).X查询一条值()
		t.AssertNil(err)
		t.Assert(value.String(), "name_1")
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		value, err := db.X创建Model对象(table).X字段保留过滤("nickname").X条件("id", 0).X查询一条值()
		t.AssertNil(err)
		t.Assert(value, nil)
	})
}

func Test_Model_Array(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		all, err := db.X创建Model对象(table).X条件("id", g.Slice别名{1, 2, 3}).X查询()
		t.AssertNil(err)
		t.Assert(all.X取字段数组("id"), g.Slice别名{1, 2, 3})
		t.Assert(all.X取字段数组("nickname"), g.Slice别名{"name_1", "name_2", "name_3"})
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		array, err := db.X创建Model对象(table).X字段保留过滤("nickname").X条件("id", g.Slice别名{1, 2, 3}).X查询数组()
		t.AssertNil(err)
		t.Assert(array, g.Slice别名{"name_1", "name_2", "name_3"})
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		array, err := db.X创建Model对象(table).X查询数组("nickname", "id", g.Slice别名{1, 2, 3})
		t.AssertNil(err)
		t.Assert(array, g.Slice别名{"name_1", "name_2", "name_3"})
	})
}

func Test_Model_Count(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	单元测试类.C(t, func(t *单元测试类.T) {
		count, err := db.X创建Model对象(table).X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(TableSize))
	})
	// 带缓存计数，检查内部ctx数据特性。
	单元测试类.C(t, func(t *单元测试类.T) {
		for i := 0; i < 10; i++ {
			count, err := db.X创建Model对象(table).X缓存(db类.X缓存选项{
				X时长: time.Second * 10,
				X名称:     uid类.X生成(),
				X强制缓存:    false,
			}).X查询行数()
			t.AssertNil(err)
			t.Assert(count, int64(TableSize))
		}
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		count, err := db.X创建Model对象(table).X字段排除过滤("id").X条件("id>8").X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(2))
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		count, err := db.X创建Model对象(table).X字段保留过滤("distinct id").X条件("id>8").X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(2))
	})
	// COUNT...LIMIT...
	单元测试类.C(t, func(t *单元测试类.T) {
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
		CreateTime 时间类.Time
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		var users []User
		err := db.X创建Model对象(table).X查询到结构体指针(&users)
		t.AssertNil(err)
		t.Assert(len(users), TableSize)
	})
}

func Test_Model_Struct(t *testing.T) {
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
		err := db.X创建Model对象(table).X条件("id=1").X查询到结构体指针(user)
		t.AssertNil(err)
		t.Assert(user.NickName, "name_1")
		t.Assert(user.CreateTime.String(), CreateTime)
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
		err := db.X创建Model对象(table).X条件("id=1").X查询到结构体指针(user)
		t.AssertNil(err)
		t.Assert(user.NickName, "name_1")
		t.Assert(user.CreateTime.String(), CreateTime)
	})
	// 自动创建结构体对象。
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *时间类.Time
		}
		user := (*User)(nil)
		err := db.X创建Model对象(table).X条件("id=1").X查询到结构体指针(&user)
		t.AssertNil(err)
		t.Assert(user.NickName, "name_1")
		t.Assert(user.CreateTime.String(), CreateTime)
	})
	// Just using Scan.
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *时间类.Time
		}
		user := (*User)(nil)
		err := db.X创建Model对象(table).X条件("id=1").X查询到结构体指针(&user)
		if err != nil {
			单元测试类.Error(err)
		}
		t.Assert(user.NickName, "name_1")
		t.Assert(user.CreateTime.String(), CreateTime)
	})
	// sql.ErrNoRows
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *时间类.Time
		}
		user := new(User)
		err := db.X创建Model对象(table).X条件("id=-1").X查询到结构体指针(user)
		t.Assert(err, sql.ErrNoRows)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *时间类.Time
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

	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Id         MyInt
			Passport   string
			Password   string
			NickName   string
			CreateTime 时间类.Time
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

	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime 时间类.Time
		}
		var users []User
		err := db.X创建Model对象(table).X排序("id asc").X查询到结构体指针(&users)
		if err != nil {
			单元测试类.Error(err)
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
	// 自动创建结构体切片。
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *时间类.Time
		}
		var users []*User
		err := db.X创建Model对象(table).X排序("id asc").X查询到结构体指针(&users)
		if err != nil {
			单元测试类.Error(err)
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
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *时间类.Time
		}
		var users []*User
		err := db.X创建Model对象(table).X排序("id asc").X查询到结构体指针(&users)
		if err != nil {
			单元测试类.Error(err)
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
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *时间类.Time
		}
		var users []*User
		err := db.X创建Model对象(table).X条件("id<0").X查询到结构体指针(&users)
		t.AssertNil(err)
	})
}

func Test_Model_StructsWithOrmTag(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		db.X设置调试模式(true)
		defer db.X设置调试模式(false)
		type User struct {
			Uid      int `orm:"id"`
			Passport string
			Password string     `orm:"password"`
			Name     string     `orm:"nickname"`
			Time     时间类.Time `orm:"create_time"`
		}
		var (
			users  []User
			buffer = bytes.NewBuffer(nil)
		)
		db.X取日志记录器().(*日志类.Logger).X设置Writer(buffer)
		defer db.X取日志记录器().(*日志类.Logger).X设置Writer(os.Stdout)
		db.X创建Model对象(table).X排序("id asc").X查询到结构体指针(&users)
		// 打印buffer.String()的输出结果
		t.Assert(
			文本类.X是否包含(buffer.String(), "SELECT `id`,`passport`,`password`,`nickname`,`create_time` FROM `user"),
			true,
		)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
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

	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime 时间类.Time
		}
		user := new(User)
		err := db.X创建Model对象(table).X条件("id=1").X查询到结构体指针(user)
		t.AssertNil(err)
		t.Assert(user.NickName, "name_1")
		t.Assert(user.CreateTime.String(), CreateTime)
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
		err := db.X创建Model对象(table).X条件("id=1").X查询到结构体指针(user)
		t.AssertNil(err)
		t.Assert(user.NickName, "name_1")
		t.Assert(user.CreateTime.String(), CreateTime)
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
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *时间类.Time
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
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *时间类.Time
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
	tableName2 := "user_" + 时间类.X创建并按当前时间().X取文本时间戳纳秒()
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
		单元测试类.AssertNil(err)
	}
	defer dropTable(tableName2)
	r, err := db.X插入(ctx, tableName2, g.Map{
		"id":   1,
		"name": "table2_1",
		"age":  18,
	})
	单元测试类.AssertNil(err)
	n, _ := r.RowsAffected()
	单元测试类.Assert(n, 1)

	// ScanAndCount 使用普通结构体作为结果进行扫描并计数
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *时间类.Time
		}
		user := new(User)
		var count int
		err := db.X创建Model对象(table).X条件("id=1").X查询与行数到指针(user, &count, true)
		t.AssertNil(err)
		t.Assert(user.NickName, "name_1")
		t.Assert(user.CreateTime.String(), CreateTime)
		t.Assert(count, 1)
	})
	// ScanAndCount 采用普通数组作为结果进行扫描并计数
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime 时间类.Time
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
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *时间类.Time
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
	// ScanAndCount 带分页功能
// （注：由于没有提供完整的代码上下文，这里的翻译可能不够精确。根据现有信息，“ScanAndCount with page”可以理解为这个函数或方法用于扫描数据并结合分页进行计数。）
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime 时间类.Time
		}
		var users []User
		var count int
		err := db.X创建Model对象(table).X排序("id asc").X设置分页(1, 3).X查询与行数到指针(&users, &count, true)
		t.AssertNil(err)
		t.Assert(len(users), 3)
		t.Assert(count, TableSize)
	})
	// ScanAndCount 带有唯一性的扫描并计数
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime 时间类.Time
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
	单元测试类.C(t, func(t *单元测试类.T) {
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
	// ScanAndCount 使用join并返回CodeDbOperationError
// （根据代码片段，可能的完整含义：）
// ```go
// ScanAndCount 函数在使用join操作进行扫描并计数时，如果发生数据库操作错误，则返回CodeDbOperationError错误码
// 请提供更多上下文信息以便提供更精确的翻译。
	单元测试类.C(t, func(t *单元测试类.T) {
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
		t.Assert(错误类.X取错误码(err), 错误码类.CodeDbOperationError)
		t.Assert(count, 0)
		t.AssertEQ(users, nil)
	})
}

func Test_Model_Scan_NilSliceAttrWhenNoRecordsFound(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime 时间类.Time
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

	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X排序("id DESC").X查询()
		t.AssertNil(err)
		t.Assert(len(result), TableSize)
		t.Assert(result[0]["nickname"].String(), fmt.Sprintf("name_%d", TableSize))
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X排序("NULL").X查询()
		t.AssertNil(err)
		t.Assert(len(result), TableSize)
		t.Assert(result[0]["nickname"].String(), "name_1")
	})

}

func Test_Model_GroupBy(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X排序分组("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), TableSize)
		t.Assert(result[0]["nickname"].String(), "name_1")
	})
}

func Test_Model_Data(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		table := createInitTable()
		defer dropTable(table)
		result, err := db.X创建Model对象(table).X设置数据("nickname=?", "test").X条件("id=?", 3).X更新()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		table := createTable()
		defer dropTable(table)
		users := make([]g.MapStrAny, 0)
		for i := 1; i <= 10; i++ {
			users = append(users, g.MapStrAny{
				"id":          i,
				"passport":    fmt.Sprintf(`passport_%d`, i),
				"password":    fmt.Sprintf(`password_%d`, i),
				"nickname":    fmt.Sprintf(`nickname_%d`, i),
				"create_time": 时间类.X创建并按当前时间().String(),
			})
		}
		result, err := db.X创建Model对象(table).X设置数据(users).X设置批量操作行数(2).X插入()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 10)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		table := createTable()
		defer dropTable(table)
		users := 数组类.X创建()
		for i := 1; i <= 10; i++ {
			users.Append别名(g.MapStrAny{
				"id":          i,
				"passport":    fmt.Sprintf(`passport_%d`, i),
				"password":    fmt.Sprintf(`password_%d`, i),
				"nickname":    fmt.Sprintf(`nickname_%d`, i),
				"create_time": 时间类.X创建并按当前时间().String(),
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
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件("id=? and nickname=?", 3, "name_3").X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["id"].X取整数(), 3)
	})

	// slice
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件(g.Slice别名{"id", 3}).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["id"].X取整数(), 3)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件(g.Slice别名{"id", 3, "nickname", "name_3"}).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["id"].X取整数(), 3)
	})

	// slice parameter
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件("id=? and nickname=?", g.Slice别名{3, "name_3"}).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["id"].X取整数(), 3)
	})
	// map like
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件(g.Map{
			"passport like": "user_1%",
		}).X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 2)
		t.Assert(result[0].X取Map类().X取值("id"), 1)
		t.Assert(result[1].X取Map类().X取值("id"), 10)
	})
	// map + slice 参数
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件(g.Map{
			"id":       g.Slice别名{1, 2, 3},
			"passport": g.Slice别名{"user_2", "user_3"},
		}).X条件("id=? and nickname=?", g.Slice别名{3, "name_3"}).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["id"].X取整数(), 3)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件("id=3", g.Slice别名{}).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["id"].X取整数(), 3)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件("id=?", g.Slice别名{3}).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["id"].X取整数(), 3)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件("id", 3).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["id"].X取整数(), 3)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件("id", 3).X条件("nickname", "name_3").X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件("id", 3).X条件("nickname", "name_3").X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件("id", 30).X条件或("nickname", "name_3").X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件("id", 30).X条件或("nickname", "name_3").X条件("id>?", 1).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件("id", 30).X条件或("nickname", "name_3").X条件("id>", 1).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	// slice
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件("id=? AND nickname=?", g.Slice别名{3, "name_3"}...).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件("id=? AND nickname=?", g.Slice别名{3, "name_3"}).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件("passport like ? and nickname like ?", g.Slice别名{"user_3", "name_3"}).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	// map
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件(g.Map{"id": 3, "nickname": "name_3"}).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	// map key operator
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件(g.Map{"id>": 1, "id<": 3}).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 2)
	})

	// gmap.Map
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件(map类.X创建并从Map(g.MapAnyAny{"id": 3, "nickname": "name_3"})).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	// gmap.Map 键操作器
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件(map类.X创建并从Map(g.MapAnyAny{"id>": 1, "id<": 3})).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 2)
	})

	// list map
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件(map类.X创建链表Map并从Map(g.MapAnyAny{"id": 3, "nickname": "name_3"})).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	// 列表映射键操作员
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件(map类.X创建链表Map并从Map(g.MapAnyAny{"id>": 1, "id<": 3})).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 2)
	})

	// tree map
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件(map类.X创建红黑树Map并从Map(工具类.X比较文本, g.MapAnyAny{"id": 3, "nickname": "name_3"})).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	// 树状映射键操作器
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件(map类.X创建红黑树Map并从Map(工具类.X比较文本, g.MapAnyAny{"id>": 1, "id<": 3})).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 2)
	})

	// 复杂条件 1
	单元测试类.C(t, func(t *单元测试类.T) {
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
	// 复杂条件 2
	单元测试类.C(t, func(t *单元测试类.T) {
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
	// 结构体，自动映射和过滤。
	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件("id IN(?)", g.Slice别名{1, 3}).X排序("id ASC").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 2)
		t.Assert(result[0]["id"].X取整数(), 1)
		t.Assert(result[1]["id"].X取整数(), 3)
	})
	// slice + string
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件("nickname=? AND id IN(?)", "name_3", g.Slice别名{1, 3}).X排序("id ASC").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 1)
		t.Assert(result[0]["id"].X取整数(), 3)
	})
	// slice + map
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件(g.Map{
			"id":       g.Slice别名{1, 3},
			"nickname": "name_3",
		}).X排序("id ASC").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 1)
		t.Assert(result[0]["id"].X取整数(), 3)
	})
	// slice + struct
	单元测试类.C(t, func(t *单元测试类.T) {
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

	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
		conditions := g.Map{
			"id < 4": "",
		}
		result, err := db.X创建Model对象(table).X条件(conditions).X排序("id desc").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["id"].X取整数(), 3)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
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

	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件("create_time>?", 时间类.X创建并从文本("2010-09-01")).X查询()
		t.AssertNil(err)
		t.Assert(len(result), 10)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件("create_time>?", *时间类.X创建并从文本("2010-09-01")).X查询()
		t.AssertNil(err)
		t.Assert(len(result), 10)
	})
}

func Test_Model_WherePri(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	// primary key
	单元测试类.C(t, func(t *单元测试类.T) {
		one, err := db.X创建Model对象(table).X条件并识别主键(3).X查询一条()
		t.AssertNil(err)
		t.AssertNE(one, nil)
		t.Assert(one["id"].X取整数(), 3)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		all, err := db.X创建Model对象(table).X条件并识别主键(g.Slice别名{3, 9}).X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(all), 2)
		t.Assert(all[0]["id"].X取整数(), 3)
		t.Assert(all[1]["id"].X取整数(), 9)
	})

	// string
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("id=? and nickname=?", 3, "name_3").X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["id"].X取整数(), 3)
	})
	// slice parameter
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("id=? and nickname=?", g.Slice别名{3, "name_3"}).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["id"].X取整数(), 3)
	})
	// map like
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键(g.Map{
			"passport like": "user_1%",
		}).X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 2)
		t.Assert(result[0].X取Map类().X取值("id"), 1)
		t.Assert(result[1].X取Map类().X取值("id"), 10)
	})
	// map + slice 参数
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键(g.Map{
			"id":       g.Slice别名{1, 2, 3},
			"passport": g.Slice别名{"user_2", "user_3"},
		}).X条件("id=? and nickname=?", g.Slice别名{3, "name_3"}).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["id"].X取整数(), 3)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键(g.Map{
			"id":       g.Slice别名{1, 2, 3},
			"passport": g.Slice别名{"user_2", "user_3"},
		}).X条件或("nickname=?", g.Slice别名{"name_4"}).X条件("id", 3).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["id"].X取整数(), 2)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("id=3", g.Slice别名{}).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["id"].X取整数(), 3)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("id=?", g.Slice别名{3}).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["id"].X取整数(), 3)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("id", 3).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["id"].X取整数(), 3)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("id", 3).X条件并识别主键("nickname", "name_3").X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("id", 3).X条件("nickname", "name_3").X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("id", 30).X条件或("nickname", "name_3").X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("id", 30).X条件或("nickname", "name_3").X条件("id>?", 1).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("id", 30).X条件或("nickname", "name_3").X条件("id>", 1).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	// slice
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("id=? AND nickname=?", g.Slice别名{3, "name_3"}...).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("id=? AND nickname=?", g.Slice别名{3, "name_3"}).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("passport like ? and nickname like ?", g.Slice别名{"user_3", "name_3"}).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	// map
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键(g.Map{"id": 3, "nickname": "name_3"}).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	// map key operator
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键(g.Map{"id>": 1, "id<": 3}).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 2)
	})

	// gmap.Map
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键(map类.X创建并从Map(g.MapAnyAny{"id": 3, "nickname": "name_3"})).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	// gmap.Map 键操作器
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键(map类.X创建并从Map(g.MapAnyAny{"id>": 1, "id<": 3})).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 2)
	})

	// list map
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键(map类.X创建链表Map并从Map(g.MapAnyAny{"id": 3, "nickname": "name_3"})).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	// 列表映射键操作员
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键(map类.X创建链表Map并从Map(g.MapAnyAny{"id>": 1, "id<": 3})).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 2)
	})

	// tree map
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键(map类.X创建红黑树Map并从Map(工具类.X比较文本, g.MapAnyAny{"id": 3, "nickname": "name_3"})).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
	// 树状映射键操作器
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键(map类.X创建红黑树Map并从Map(工具类.X比较文本, g.MapAnyAny{"id>": 1, "id<": 3})).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 2)
	})

	// 复杂条件 1
	单元测试类.C(t, func(t *单元测试类.T) {
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
	// 复杂条件 2
	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("id IN(?)", g.Slice别名{1, 3}).X排序("id ASC").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 2)
		t.Assert(result[0]["id"].X取整数(), 1)
		t.Assert(result[1]["id"].X取整数(), 3)
	})
	// slice + string
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("nickname=? AND id IN(?)", "name_3", g.Slice别名{1, 3}).X排序("id ASC").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 1)
		t.Assert(result[0]["id"].X取整数(), 3)
	})
	// slice + map
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键(g.Map{
			"id":       g.Slice别名{1, 3},
			"nickname": "name_3",
		}).X排序("id ASC").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 1)
		t.Assert(result[0]["id"].X取整数(), 3)
	})
	// slice + struct
	单元测试类.C(t, func(t *单元测试类.T) {
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
// 创建并初始化表格，将结果赋值给table变量
// table := createInitTable()
// 在函数结束时（defer关键字确保的）执行dropTable函数，传入table作为参数，用于删除已创建的表格
// defer dropTable(table)

// DELETE...LIMIT
// 这段代码引用了 GitHub 上 go-sqlite3 库的一个 PR（#802）
// 使用 gtest 单元测试框架进行测试
// 测试内容如下：
// 针对 table 模型，使用 Where 条件为 1 并限制删除数量为 2 的记录执行删除操作
// gtest.C 函数中进行单元测试的主体部分：
// t 表示当前测试环境
// 调用 db.Model(table).Where(1).Limit(2).Delete() 删除符合条件的记录，将结果和错误信息分别赋值给 result 和 err
// 断言 err 为空（即无错误发生）
// 获取被影响的行数并赋值给 n，这里忽略了可能的错误信息
// 最后断言受影响的行数 n 等于 2，即成功删除了两条记录

	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X设置分页(3, 3).X排序("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["id"], 7)
		t.Assert(result[1]["id"], 8)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
		table := createTable()
		defer dropTable(table)
		r, err := db.X创建Model对象(table).X字段保留过滤("id, passport", "password", "create_time").X设置数据(g.Map{
			"id":          1,
			"passport":    "1",
			"password":    "1",
			"nickname":    "1",
			"create_time": 时间类.X创建并按当前时间().String(),
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
	单元测试类.C(t, func(t *单元测试类.T) {
		table := createTable()
		defer dropTable(table)
		r, err := db.X创建Model对象(table).X过滤空值数据().X设置数据(g.Map{
			"id":          1,
			"passport":    "1",
			"password":    "1",
			"nickname":    "",
			"create_time": 时间类.X创建并按当前时间().String(),
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
	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
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

	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
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
	noPrefixName := fmt.Sprintf(`%s_%d`, TableName, 时间类.X取时间戳纳秒())
	table := TableNamePrefix + noPrefixName
	createInitTableWithDb(db, table)
	defer dropTable(table)
	// Select.
	单元测试类.C(t, func(t *单元测试类.T) {
		r, err := db.X创建Model对象(noPrefixName).X条件("id in (?)", g.Slice别名{1, 2}).X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(r), 2)
		t.Assert(r[0]["id"], "1")
		t.Assert(r[1]["id"], "2")
	})
	// 使用别名进行选择。
	单元测试类.C(t, func(t *单元测试类.T) {
		r, err := db.X创建Model对象(noPrefixName+" as u").X条件("u.id in (?)", g.Slice别名{1, 2}).X排序("u.id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(r), 2)
		t.Assert(r[0]["id"], "1")
		t.Assert(r[1]["id"], "2")
	})
	// 使用别名选择到结构体。
	单元测试类.C(t, func(t *单元测试类.T) {
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
	// 使用别名和连接语句进行选择。
	单元测试类.C(t, func(t *单元测试类.T) {
		r, err := db.X创建Model对象(noPrefixName+" as u1").X左连接(noPrefixName+" as u2", "u2.id=u1.id").X条件("u1.id in (?)", g.Slice别名{1, 2}).X排序("u1.id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(r), 2)
		t.Assert(r[0]["id"], "1")
		t.Assert(r[1]["id"], "2")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
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

	单元测试类.C(t, func(t *单元测试类.T) {
		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["id"], 1)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
		subQuery := fmt.Sprintf("select * from `%s`", table)
		r, err := db.X创建Model对象(table, "t1").X字段保留过滤("t2.id").X左连接(subQuery, "t2", "t2.id=t1.id").X查询数组()
		t.AssertNil(err)
		t.Assert(len(r), TableSize)
		t.Assert(r[0], "1")
		t.Assert(r[TableSize-1], TableSize)
	})
}

func Test_Model_Cache(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		one, err := db.X创建Model对象(table).X缓存(db类.X缓存选项{
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

		one, err = db.X创建Model对象(table).X缓存(db类.X缓存选项{
			X时长: time.Second,
			X名称:     "test1",
			X强制缓存:    false,
		}).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["passport"], "user_1")

		time.Sleep(time.Second * 2)

		one, err = db.X创建Model对象(table).X缓存(db类.X缓存选项{
			X时长: time.Second,
			X名称:     "test1",
			X强制缓存:    false,
		}).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["passport"], "user_100")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		one, err := db.X创建Model对象(table).X缓存(db类.X缓存选项{
			X时长: time.Second,
			X名称:     "test2",
			X强制缓存:    false,
		}).X条件并识别主键(2).X查询一条()
		t.AssertNil(err)
		t.Assert(one["passport"], "user_2")

		r, err := db.X创建Model对象(table).X设置数据("passport", "user_200").X缓存(db类.X缓存选项{
			X时长: -1,
			X名称:     "test2",
			X强制缓存:    false,
		}).X条件并识别主键(2).X更新()
		t.AssertNil(err)
		n, err := r.RowsAffected()
		t.AssertNil(err)
		t.Assert(n, 1)

		one, err = db.X创建Model对象(table).X缓存(db类.X缓存选项{
			X时长: time.Second,
			X名称:     "test2",
			X强制缓存:    false,
		}).X条件并识别主键(2).X查询一条()
		t.AssertNil(err)
		t.Assert(one["passport"], "user_200")
	})
	// transaction.
	单元测试类.C(t, func(t *单元测试类.T) {
		// 为id 3创建缓存
		one, err := db.X创建Model对象(table).X缓存(db类.X缓存选项{
			X时长: time.Second,
			X名称:     "test3",
			X强制缓存:    false,
		}).X条件并识别主键(3).X查询一条()
		t.AssertNil(err)
		t.Assert(one["passport"], "user_3")

		r, err := db.X创建Model对象(table).X设置数据("passport", "user_300").X缓存(db类.X缓存选项{
			X时长: time.Second,
			X名称:     "test3",
			X强制缓存:    false,
		}).X条件并识别主键(3).X更新()
		t.AssertNil(err)
		n, err := r.RowsAffected()
		t.AssertNil(err)
		t.Assert(n, 1)

		err = db.X事务(context.TODO(), func(ctx context.Context, tx db类.TX) error {
			one, err := tx.X创建Model对象(table).X缓存(db类.X缓存选项{
				X时长: time.Second,
				X名称:     "test3",
				X强制缓存:    false,
			}).X条件并识别主键(3).X查询一条()
			t.AssertNil(err)
			t.Assert(one["passport"], "user_300")
			return nil
		})
		t.AssertNil(err)

		one, err = db.X创建Model对象(table).X缓存(db类.X缓存选项{
			X时长: time.Second,
			X名称:     "test3",
			X强制缓存:    false,
		}).X条件并识别主键(3).X查询一条()
		t.AssertNil(err)
		t.Assert(one["passport"], "user_3")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		// 为id 4创建缓存
		one, err := db.X创建Model对象(table).X缓存(db类.X缓存选项{
			X时长: time.Second,
			X名称:     "test4",
			X强制缓存:    false,
		}).X条件并识别主键(4).X查询一条()
		t.AssertNil(err)
		t.Assert(one["passport"], "user_4")

		r, err := db.X创建Model对象(table).X设置数据("passport", "user_400").X缓存(db类.X缓存选项{
			X时长: time.Second,
			X名称:     "test3",
			X强制缓存:    false,
		}).X条件并识别主键(4).X更新()
		t.AssertNil(err)
		n, err := r.RowsAffected()
		t.AssertNil(err)
		t.Assert(n, 1)

		err = db.X事务(context.TODO(), func(ctx context.Context, tx db类.TX) error {
			// 缓存功能已禁用。
			one, err := tx.X创建Model对象(table).X缓存(db类.X缓存选项{
				X时长: time.Second,
				X名称:     "test4",
				X强制缓存:    false,
			}).X条件并识别主键(4).X查询一条()
			t.AssertNil(err)
			t.Assert(one["passport"], "user_400")
			// Update the cache.
			r, err := tx.X创建Model对象(table).X设置数据("passport", "user_4000").
				X缓存(db类.X缓存选项{
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
		one, err = db.X创建Model对象(table).X缓存(db类.X缓存选项{
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

	单元测试类.C(t, func(t *单元测试类.T) {
		all, err := db.X创建Model对象(table).X条件("id > 1").X排序分组("id").X设置分组条件("id > 8").X查询()
		t.AssertNil(err)
		t.Assert(len(all), 2)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		all, err := db.X创建Model对象(table).X条件("id > 1").X排序分组("id").X设置分组条件("id > ?", 8).X查询()
		t.AssertNil(err)
		t.Assert(len(all), 2)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		all, err := db.X创建Model对象(table).X条件("id > ?", 1).X排序分组("id").X设置分组条件("id > ?", 8).X查询()
		t.AssertNil(err)
		t.Assert(len(all), 2)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		all, err := db.X创建Model对象(table).X条件("id > ?", 1).X排序分组("id").X设置分组条件("id", 8).X查询()
		t.AssertNil(err)
		t.Assert(len(all), 1)
	})
}

func Test_Model_Distinct(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		all, err := db.X创建Model对象(table, "t").X字段保留过滤("distinct t.id").X条件("id > 1").X排序分组("id").X设置分组条件("id > 8").X查询()
		t.AssertNil(err)
		t.Assert(len(all), 2)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		count, err := db.X创建Model对象(table).X条件("id > 1").X设置去重().X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(9))
	})
}

func Test_Model_Min_Max(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		value, err := db.X创建Model对象(table, "t").X字段保留过滤("min(t.id)").X条件("id > 1").X查询一条值()
		t.AssertNil(err)
		t.Assert(value.X取整数(), 2)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		value, err := db.X创建Model对象(table, "t").X字段保留过滤("max(t.id)").X条件("id > 1").X查询一条值()
		t.AssertNil(err)
		t.Assert(value.X取整数(), 10)
	})
}

func Test_Model_Fields_AutoMapping(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		value, err := db.X创建Model对象(table).X字段保留过滤("ID").X条件("id", 2).X查询一条值()
		t.AssertNil(err)
		t.Assert(value.X取整数(), 2)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		value, err := db.X创建Model对象(table).X字段保留过滤("NICK_NAME").X条件("id", 2).X查询一条值()
		t.AssertNil(err)
		t.Assert(value.String(), "name_2")
	})
	// Map
	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
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

// "id":          i, // "id"字段：存储变量i的值
// "passport":    fmt.Sprintf(`user_%d`, i), // "passport"字段：格式化输出字符串，形式为"user_数字"，其中数字为变量i的值
// "password":    fmt.Sprintf(`pass_%d`, i), // "password"字段：格式化输出字符串，形式为"pass_数字"，其中数字为变量i的值
// "nickname":    fmt.Sprintf(`name_%d`, i), // "nickname"字段：格式化输出字符串，形式为"name_数字"，其中数字为变量i的值
// "create_time": gtime.NewFromStr(CreateTime).String(), // "create_time"字段：通过CreateTime字符串创建一个gtime.Time对象，并获取该时间对象的字符串表示形式
// （注：这里的`CreateTime`应是一个符合日期时间格式的字符串）

	单元测试类.C(t, func(t *单元测试类.T) {
		value, err := db.X创建Model对象(table).X字段排除过滤("Passport, Password, NickName, CreateTime").X条件("id", 2).X查询一条值()
		t.AssertNil(err)
		t.Assert(value.X取整数(), 2)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		value, err := db.X创建Model对象(table).X字段排除过滤("ID, Passport, Password, CreateTime").X条件("id", 2).X查询一条值()
		t.AssertNil(err)
		t.Assert(value.String(), "name_2")
	})
	// Map
	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
		one, err := db.X创建Model对象(table).X字段保留过滤(A{}).X条件("id", 2).X查询一条()
		t.AssertNil(err)
		t.Assert(len(one), 2)
		t.Assert(one["passport"], "user_2")
		t.Assert(one["password"], "pass_2")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		one, err := db.X创建Model对象(table).X字段保留过滤(&A{}).X条件("id", 2).X查询一条()
		t.AssertNil(err)
		t.Assert(len(one), 2)
		t.Assert(one["passport"], "user_2")
		t.Assert(one["password"], "pass_2")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		one, err := db.X创建Model对象(table).X字段保留过滤(B{}).X条件("id", 2).X查询一条()
		t.AssertNil(err)
		t.Assert(len(one), 3)
		t.Assert(one["passport"], "user_2")
		t.Assert(one["password"], "pass_2")
		t.Assert(one["nickname"], "name_2")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件(`id`, g.Slice别名{}).X查询()
		t.AssertNil(err)
		t.Assert(len(result), 0)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件(`id in(?)`, g.Slice别名{}).X查询()
		t.AssertNil(err)
		t.Assert(len(result), 0)
	})
}

func Test_Model_HasTable(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertNil(db.X取Core对象().X删除所有表查询缓存(ctx))
		result, err := db.X取Core对象().X是否存在表名(table)
		t.Assert(result, true)
		t.AssertNil(err)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertNil(db.X取Core对象().X删除所有表查询缓存(ctx))
		result, err := db.X取Core对象().X是否存在表名("table12321")
		t.Assert(result, false)
		t.AssertNil(err)
	})
}

func Test_Model_HasField(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X是否存在字段("id")
		t.Assert(result, true)
		t.AssertNil(err)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X是否存在字段("id123")
		t.Assert(result, false)
		t.AssertNil(err)
	})
}

// 问题：https://github.com/gogf/gf/issues/1002
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
	单元测试类.AssertNil(err)
	n, _ := result.RowsAffected()
	单元测试类.Assert(n, 1)

	// where + string.
	单元测试类.C(t, func(t *单元测试类.T) {
		v, err := db.X创建Model对象(table).X字段保留过滤("id").X条件("create_time>'2020-10-27 19:03:32' and create_time<'2020-10-27 19:03:34'").X查询一条值()
		t.AssertNil(err)
		t.Assert(v.X取整数(), 1)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		v, err := db.X创建Model对象(table).X字段保留过滤("id").X条件("create_time>'2020-10-27 19:03:32' and create_time<'2020-10-27 19:03:34'").X查询一条值()
		t.AssertNil(err)
		t.Assert(v.X取整数(), 1)
	})
	// where + 字符串参数。
	单元测试类.C(t, func(t *单元测试类.T) {
		v, err := db.X创建Model对象(table).X字段保留过滤("id").X条件("create_time>? and create_time<?", "2020-10-27 19:03:32", "2020-10-27 19:03:34").X查询一条值()
		t.AssertNil(err)
		t.Assert(v.X取整数(), 1)
	})
	// where + gtime.Time 参数
	单元测试类.C(t, func(t *单元测试类.T) {
		v, err := db.X创建Model对象(table).X字段保留过滤("id").X条件("create_time>? and create_time<?", 时间类.X创建("2020-10-27 19:03:32"), 时间类.X创建("2020-10-27 19:03:34")).X查询一条值()
		t.AssertNil(err)
		t.Assert(v.X取整数(), 1)
	})
// TODO
// 对包含time.Time参数的where条件进行测试，时间采用UTC格式。
// 使用gtest.C进行单元测试，传入t参数作为测试环境上下文。
// 
// 首先，解析两个时间字符串为time.Time类型：
// t1表示"2020-10-27 11:03:32"
// t2表示"2020-10-27 11:03:34"
// 
// 然后执行如下代码块：
// 通过db.Model(table).Fields("id")设置SQL查询模型和字段（这里为表名为table的表中的id字段），
// 并添加where条件：create_time>? and create_time<?，其中问号占位符分别被t1和t2替换。
// 执行查询并获取查询结果值v及可能的错误err。
// 验证错误err应为nil，即查询无错误发生。
// 验证查询结果值v转换为整数后为1。
}

func createTableForTimeZoneTest() string {
	tableName := "user_" + 时间类.X创建并按当前时间().X取文本时间戳纳秒()
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
		单元测试类.Fatal(err)
	}
	return tableName
}

// 这是GitHub上gogf/gf项目的一个问题链接，具体为第1012号问题
func Test_TimeZoneInsert(t *testing.T) {
	tableName := createTableForTimeZoneTest()
	defer dropTable(tableName)

	tokyoLoc, err := time.LoadLocation("Asia/Tokyo")
	单元测试类.AssertNil(err)

	CreateTime := "2020-11-22 12:23:45"
	UpdateTime := "2020-11-22 13:23:45"
	DeleteTime := "2020-11-22 14:23:45"
	type User struct {
		Id        int         `json:"id"`
		CreatedAt *时间类.Time `json:"created_at"`
		UpdatedAt 时间类.Time  `json:"updated_at"`
		DeletedAt time.Time   `json:"deleted_at"`
	}
	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", CreateTime, tokyoLoc)
	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", UpdateTime, tokyoLoc)
	t3, _ := time.ParseInLocation("2006-01-02 15:04:05", DeleteTime, tokyoLoc)
	u := &User{
		Id:        1,
		CreatedAt: 时间类.X创建(t1.UTC()),
		UpdatedAt: *时间类.X创建(t2.UTC()),
		DeletedAt: t3.UTC(),
	}

	单元测试类.C(t, func(t *单元测试类.T) {
		_, _ = db.X创建Model对象(tableName).X禁用时间自动更新特性().X插入(u)
		userEntity := &User{}
		err := db.X创建Model对象(tableName).X条件("id", 1).X禁用时间自动更新特性().X查询到结构体指针(&userEntity)
		t.AssertNil(err)
// TODO
// t.Assert(userEntity.CreatedAt.String(), "2020-11-22 11:23:45") // 待办：断言userEntity的CreatedAt字段转换为字符串后，其值应为"2020-11-22 11:23:45"
// t.Assert(userEntity.UpdatedAt.String(), "2020-11-22 12:23:45") // 待办：断言userEntity的UpdatedAt字段转换为字符串后，其值应为"2020-11-22 12:23:45"
// t.Assert(gtime.NewFromTime(userEntity.DeletedAt).String(), "2020-11-22 13:23:45") // 待办：断言根据userEntity的DeletedAt字段创建的新时间对象转换为字符串后，其值应为"2020-11-22 13:23:45"
	})
}

func Test_Model_Fields_Map_Struct(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	// map
	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
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

	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件包含("id", g.Slice别名{1, 2, 3, 4}).X条件包含("id", g.Slice别名{3, 4, 5}).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 2)
		t.Assert(result[0]["id"], 3)
		t.Assert(result[1]["id"], 4)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件包含("id", g.Slice别名{}).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 0)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X过滤空值条件().X条件包含("id", g.Slice别名{}).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), TableSize)
	})
}

func Test_Model_WhereNotIn(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
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

	单元测试类.C(t, func(t *单元测试类.T) {
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

	单元测试类.C(t, func(t *单元测试类.T) {
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

	单元测试类.C(t, func(t *单元测试类.T) {
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

	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件取范围以外("id", 2, 8).X条件取范围以外("id", 3, 100).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 1)
		t.Assert(result[0]["id"], 1)
	})
}

func Test_Model_WhereOrBetween(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
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

	单元测试类.C(t, func(t *单元测试类.T) {
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

	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件模糊匹配以外("nickname", "name%").X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 0)
	})
}

func Test_Model_WhereOrLike(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
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

	单元测试类.C(t, func(t *单元测试类.T) {
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

	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件NULL值("nickname").X条件NULL值("passport").X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 0)
	})
}

func Test_Model_WhereNotNull(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
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

	单元测试类.C(t, func(t *单元测试类.T) {
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

	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件小于("id", 3).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 2)
		t.Assert(result[0]["id"], 1)
	})
}

func Test_Model_WhereLTE(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件小于等于("id", 3).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["id"], 1)
	})
}

func Test_Model_WhereGT(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件大于("id", 8).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 2)
		t.Assert(result[0]["id"], 9)
	})
}

func Test_Model_WhereGTE(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件大于等于("id", 8).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["id"], 8)
	})
}

func Test_Model_WhereOrLT(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
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

	单元测试类.C(t, func(t *单元测试类.T) {
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

	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件大于("id", 8).X条件或大于("id", 7).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["id"], 8)
	})
}

func Test_Model_WhereOrGTE(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件大于等于("id", 8).X条件或大于等于("id", 7).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 4)
		t.Assert(result[0]["id"], 7)
	})
}

func Test_Model_Min_Max_Avg_Sum(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X查询最小值("id")
		t.AssertNil(err)
		t.Assert(result, 1)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X查询最大值("id")
		t.AssertNil(err)
		t.Assert(result, TableSize)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X查询平均值("id")
		t.AssertNil(err)
		t.Assert(result, 5.5)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X查询求和("id")
		t.AssertNil(err)
		t.Assert(result, 55)
	})
}

func Test_Model_CountColumn(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X查询字段行数("id")
		t.AssertNil(err)
		t.Assert(result, TableSize)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件包含("id", g.Slice别名{1, 2, 3}).X查询字段行数("id")
		t.AssertNil(err)
		t.Assert(result, 3)
	})
}

func Test_Model_InsertAndGetId(t *testing.T) {
	table := createTable()
	defer dropTable(table)
	单元测试类.C(t, func(t *单元测试类.T) {
		id, err := db.X创建Model对象(table).X设置数据(g.Map{
			"id":          1,
			"passport":    "user_1",
			"password":    "pass_1",
			"nickname":    "name_1",
			"create_time": 时间类.X创建并按当前时间().String(),
		}).X插入并取ID()
		t.AssertNil(err)
		t.Assert(id, 1)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		id, err := db.X创建Model对象(table).X设置数据(g.Map{
			"passport":    "user_2",
			"password":    "pass_2",
			"nickname":    "name_2",
			"create_time": 时间类.X创建并按当前时间().String(),
		}).X插入并取ID()
		t.AssertNil(err)
		t.Assert(id, 2)
	})
}

func Test_Model_Increment_Decrement(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件("id", 1).X更新增量("id", 100)
		t.AssertNil(err)
		rows, _ := result.RowsAffected()
		t.Assert(rows, 1)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件("id", 101).X更新减量("id", 10)
		t.AssertNil(err)
		rows, _ := result.RowsAffected()
		t.Assert(rows, 1)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		count, err := db.X创建Model对象(table).X条件("id", 91).X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(1))
	})
}

func Test_Model_Raw(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
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

	单元测试类.C(t, func(t *单元测试类.T) {
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

	单元测试类.C(t, func(t *单元测试类.T) {
		m := db.X创建Model对象(table).X链式安全().X处理函数(
			func(m *db类.Model) *db类.Model {
				return m.X设置分页(0, 3)
			},
			func(m *db类.Model) *db类.Model {
				return m.X条件("id", g.Slice别名{1, 2, 3, 4, 5, 6})
			},
			func(m *db类.Model) *db类.Model {
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

	单元测试类.C(t, func(t *单元测试类.T) {
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

	单元测试类.C(t, func(t *单元测试类.T) {
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

	单元测试类.C(t, func(t *单元测试类.T) {
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

	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
		count, err := db.X创建Model对象(table).X条件("id", 0).X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(0))
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		count, err := db.X创建Model对象(table).X过滤空值条件().X条件("id", 0).X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(TableSize))
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		count, err := db.X创建Model对象(table).X过滤空值条件().X条件("id", 0).X条件("nickname", "").X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(TableSize))
	})
	// Slice where.
	单元测试类.C(t, func(t *单元测试类.T) {
		count, err := db.X创建Model对象(table).X条件("id", g.Slice别名{1, 2, 3}).X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(3))
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		count, err := db.X创建Model对象(table).X条件("id", g.Slice别名{}).X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(0))
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		count, err := db.X创建Model对象(table).X过滤空值条件().X条件("id", g.Slice别名{}).X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(TableSize))
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		count, err := db.X创建Model对象(table).X条件("id", g.Slice别名{}).X过滤空值条件().X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(TableSize))
	})
	// Struct Where.
	单元测试类.C(t, func(t *单元测试类.T) {
		type Input struct {
			Id   []int
			Name []string
		}
		count, err := db.X创建Model对象(table).X条件(Input{}).X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(0))
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		type Input struct {
			Id   []int
			Name []string
		}
		count, err := db.X创建Model对象(table).X条件(Input{}).X过滤空值条件().X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(TableSize))
	})
	// Map Where.
	单元测试类.C(t, func(t *单元测试类.T) {
		count, err := db.X创建Model对象(table).X条件(g.Map{
			"id":       []int{},
			"nickname": []string{},
		}).X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(0))
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		count, err := db.X创建Model对象(table).X条件(g.Map{
			"id": []int{},
		}).X过滤空值条件().X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(TableSize))
	})
}

// 这是GitHub上gogf/gf仓库中关于第1387号问题的链接
func Test_Model_GTime_DefaultValue(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			Nickname   string
			CreateTime *时间类.Time
		}
		data := User{
			Id:         1,
			Passport:   "user_1",
			Password:   "pass_1",
			Nickname:   "name_1",
			CreateTime: 时间类.X创建并按当前时间(),
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

// 在函数内部使用filter不会影响外部的值。
func Test_Model_Insert_Filter(t *testing.T) {
	// map
	单元测试类.C(t, func(t *单元测试类.T) {
		table := createTable()
		defer dropTable(table)
		data := g.Map{
			"id":          1,
			"uid":         1,
			"passport":    "t1",
			"password":    "25d55ad283aa400af464c76d713c07ad",
			"nickname":    "name_1",
			"create_time": 时间类.X创建并按当前时间().String(),
		}
		result, err := db.X创建Model对象(table).X设置数据(data).X插入()
		t.AssertNil(err)
		n, _ := result.LastInsertId()
		t.Assert(n, 1)

		t.Assert(data["uid"], 1)
	})
	// slice
	单元测试类.C(t, func(t *单元测试类.T) {
		table := createTable()
		defer dropTable(table)
		data := g.Map数组{
			g.Map{
				"id":          1,
				"uid":         1,
				"passport":    "t1",
				"password":    "25d55ad283aa400af464c76d713c07ad",
				"nickname":    "name_1",
				"create_time": 时间类.X创建并按当前时间().String(),
			},
			g.Map{
				"id":          2,
				"uid":         2,
				"passport":    "t1",
				"password":    "25d55ad283aa400af464c76d713c07ad",
				"nickname":    "name_1",
				"create_time": 时间类.X创建并按当前时间().String(),
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
	单元测试类.C(t, func(t *单元测试类.T) {
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
				CreateTime: 时间类.X创建并按当前时间().String(),
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

// 从GoFrame v1.16.0版本开始，此功能不再使用，因为过滤功能会自动启用。
func Test_Model_Insert_KeyFieldNameMapping_Error(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
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
		table1 = "table1_" + 时间类.X取文本时间戳纳秒()
		table2 = "table2_" + 时间类.X取文本时间戳纳秒()
	)
	createInitTable(table1)
	defer dropTable(table1)
	createInitTable(table2)
	defer dropTable(table2)

	单元测试类.C(t, func(t *单元测试类.T) {
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
		table1 = "table1_" + 时间类.X取文本时间戳纳秒()
		table2 = "table2_" + 时间类.X取文本时间戳纳秒()
	)
	createInitTable(table1)
	defer dropTable(table1)
	createInitTable(table2)
	defer dropTable(table2)

	单元测试类.C(t, func(t *单元测试类.T) {
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
		table1 = "table1_" + 时间类.X取文本时间戳纳秒()
		table2 = "table2_" + 时间类.X取文本时间戳纳秒()
	)
	createInitTable(table1)
	defer dropTable(table1)
	createInitTable(table2)
	defer dropTable(table2)

	单元测试类.C(t, func(t *单元测试类.T) {
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

// 这是Go语言代码中的一行注释，其内容引用了GitHub上gogf/gf仓库的第1159号问题。
// 中文翻译：
// 参考GitHub上gogf/gf项目的问题1159。
func Test_ScanList_NoRecreate_PtrAttribute(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
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
		r1 := db类.Result{
			db类.Record{
				"id":   泛型类.X创建(1),
				"name": 泛型类.X创建("john"),
				"age":  泛型类.X创建(16),
			},
			db类.Record{
				"id":   泛型类.X创建(2),
				"name": 泛型类.X创建("smith"),
				"age":  泛型类.X创建(18),
			},
		}
		err = r1.X取指针列表(&s, "One")
		t.AssertNil(err)
		t.Assert(len(s), 2)
		t.Assert(s[0].One.Name, "john")
		t.Assert(s[0].One.Age, 16)
		t.Assert(s[1].One.Name, "smith")
		t.Assert(s[1].One.Age, 18)

		r2 := db类.Result{
			db类.Record{
				"id":  泛型类.X创建(1),
				"age": 泛型类.X创建(20),
			},
			db类.Record{
				"id":  泛型类.X创建(2),
				"age": 泛型类.X创建(21),
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

// 这是Go语言代码中的一行注释，其内容引用了GitHub上gogf/gf仓库的第1159号问题。
// 中文翻译：
// 参考GitHub上gogf/gf项目的问题1159。
func Test_ScanList_NoRecreate_StructAttribute(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
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
		r1 := db类.Result{
			db类.Record{
				"id":   泛型类.X创建(1),
				"name": 泛型类.X创建("john"),
				"age":  泛型类.X创建(16),
			},
			db类.Record{
				"id":   泛型类.X创建(2),
				"name": 泛型类.X创建("smith"),
				"age":  泛型类.X创建(18),
			},
		}
		err = r1.X取指针列表(&s, "One")
		t.AssertNil(err)
		t.Assert(len(s), 2)
		t.Assert(s[0].One.Name, "john")
		t.Assert(s[0].One.Age, 16)
		t.Assert(s[1].One.Name, "smith")
		t.Assert(s[1].One.Age, 18)

		r2 := db类.Result{
			db类.Record{
				"id":  泛型类.X创建(1),
				"age": 泛型类.X创建(20),
			},
			db类.Record{
				"id":  泛型类.X创建(2),
				"age": 泛型类.X创建(21),
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

// 这是Go语言代码中的一行注释，其内容引用了GitHub上gogf/gf仓库的第1159号问题。
// 中文翻译：
// 参考GitHub上gogf/gf项目的问题1159。
func Test_ScanList_NoRecreate_SliceAttribute_Ptr(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
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
		r1 := db类.Result{
			db类.Record{
				"id":   泛型类.X创建(1),
				"name": 泛型类.X创建("john"),
				"age":  泛型类.X创建(16),
			},
			db类.Record{
				"id":   泛型类.X创建(2),
				"name": 泛型类.X创建("smith"),
				"age":  泛型类.X创建(18),
			},
		}
		err = r1.X取指针列表(&s, "One")
		t.AssertNil(err)
		t.Assert(len(s), 2)
		t.Assert(s[0].One.Name, "john")
		t.Assert(s[0].One.Age, 16)
		t.Assert(s[1].One.Name, "smith")
		t.Assert(s[1].One.Age, 18)

		r2 := db类.Result{
			db类.Record{
				"id":   泛型类.X创建(100),
				"pid":  泛型类.X创建(1),
				"age":  泛型类.X创建(30),
				"name": 泛型类.X创建("john"),
			},
			db类.Record{
				"id":   泛型类.X创建(200),
				"pid":  泛型类.X创建(1),
				"age":  泛型类.X创建(31),
				"name": 泛型类.X创建("smith"),
			},
		}
		err = r2.X取指针列表(&s, "Many", "One", "pid:Id")
		// 使用 %+v 格式化输出错误信息，会包含错误类型和详细堆栈信息
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

		r3 := db类.Result{
			db类.Record{
				"id":  泛型类.X创建(100),
				"pid": 泛型类.X创建(1),
				"age": 泛型类.X创建(40),
			},
			db类.Record{
				"id":  泛型类.X创建(200),
				"pid": 泛型类.X创建(1),
				"age": 泛型类.X创建(41),
			},
		}
		err = r3.X取指针列表(&s, "Many", "One", "pid:Id")
		// 使用 %+v 格式化输出错误信息，会包含错误类型和详细堆栈信息
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

// 这是Go语言代码中的一行注释，其内容引用了GitHub上gogf/gf仓库的第1159号问题。
// 中文翻译：
// 参考GitHub上gogf/gf项目的问题1159。
func Test_ScanList_NoRecreate_SliceAttribute_Struct(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
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
		r1 := db类.Result{
			db类.Record{
				"id":   泛型类.X创建(1),
				"name": 泛型类.X创建("john"),
				"age":  泛型类.X创建(16),
			},
			db类.Record{
				"id":   泛型类.X创建(2),
				"name": 泛型类.X创建("smith"),
				"age":  泛型类.X创建(18),
			},
		}
		err = r1.X取指针列表(&s, "One")
		t.AssertNil(err)
		t.Assert(len(s), 2)
		t.Assert(s[0].One.Name, "john")
		t.Assert(s[0].One.Age, 16)
		t.Assert(s[1].One.Name, "smith")
		t.Assert(s[1].One.Age, 18)

		r2 := db类.Result{
			db类.Record{
				"id":   泛型类.X创建(100),
				"pid":  泛型类.X创建(1),
				"age":  泛型类.X创建(30),
				"name": 泛型类.X创建("john"),
			},
			db类.Record{
				"id":   泛型类.X创建(200),
				"pid":  泛型类.X创建(1),
				"age":  泛型类.X创建(31),
				"name": 泛型类.X创建("smith"),
			},
		}
		err = r2.X取指针列表(&s, "Many", "One", "pid:Id")
		// 使用 %+v 格式化输出错误信息，会包含错误类型和详细堆栈信息
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

		r3 := db类.Result{
			db类.Record{
				"id":  泛型类.X创建(100),
				"pid": 泛型类.X创建(1),
				"age": 泛型类.X创建(40),
			},
			db类.Record{
				"id":  泛型类.X创建(200),
				"pid": 泛型类.X创建(1),
				"age": 泛型类.X创建(41),
			},
		}
		err = r3.X取指针列表(&s, "Many", "One", "pid:Id")
		// 使用 %+v 格式化输出错误信息，会包含错误类型和详细堆栈信息
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
	单元测试类.C(t, func(t *单元测试类.T) {
		r := db类.Result{
			db类.Record{"id": 泛型类.X创建(nil), "name": 泛型类.X创建("john")},
			db类.Record{"id": 泛型类.X创建(nil), "name": 泛型类.X创建("smith")},
		}
		array := make([]*B, 2)
		err := r.X取数组结构体指针(&array)
		t.AssertNil(err)
		t.Assert(array[0].Id, 0)
		t.Assert(array[1].Id, 0)
		t.Assert(array[0].Name, "john")
		t.Assert(array[1].Name, "smith")
	})
}
