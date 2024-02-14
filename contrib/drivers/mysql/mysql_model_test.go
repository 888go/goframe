// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package mysql_test

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
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/os/gfile"
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
			Id:       3,
			Uid:      3,
			Passport: "t3",
			Password: "25d55ad283aa400af464c76d713c07ad",
			Nickname: "name_3",
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
			Id:         1,
			Passport:   "user_10",
			Password:   "pass_10",
			Nickname:   "name_10",
			CreateTime: "2020-10-10 12:00:01",
		}
		_, err := db.X创建Model对象(table).X设置数据(data).X条件并识别主键(1).X更新()
		t.AssertNil(err)

		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
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
	table := createInitTable()
	defer dropTable(table)
	单元测试类.C(t, func(t *单元测试类.T) {
		_, err := db.X创建Model对象(table).X设置数据(g.Map{
			"id":          1,
			"uid":         1,
			"passport":    "t1",
			"password":    "25d55ad283aa400af464c76d713c07ad",
			"nickname":    "name_1",
			"create_time": 时间类.X创建并按当前时间().String(),
		}).X插入()
		t.AssertNE(err, nil)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		_, err := db.X创建Model对象(table).X设置数据(g.Map{
			"id":          1,
			"uid":         1,
			"passport":    "t1",
			"password":    "25d55ad283aa400af464c76d713c07ad",
			"nickname":    "name_1",
			"create_time": 时间类.X创建并按当前时间().String(),
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
			{"passport": "t1"},
			{"passport": "t2"},
			{"passport": "t3"},
			{"passport": "t4"},
			{"passport": "t5"},
		}).X设置批量操作行数(2).X插入()
		if err != nil {
			单元测试类.Error(err)
		}
		n, _ := result.RowsAffected()
		t.Assert(n, 5)
	})

	// batch save
	单元测试类.C(t, func(t *单元测试类.T) {
		table := createInitTable()
		defer dropTable(table)
		result, err := db.X创建Model对象(table).X查询()
		t.AssertNil(err)
		t.Assert(len(result), TableSize)
		for _, v := range result {
			v["nickname"].X设置值(v["nickname"].String() + v["id"].String())
		}
		r, e := db.X创建Model对象(table).X设置数据(result).X插入并更新已存在()
		t.Assert(e, nil)
		n, e := r.RowsAffected()
		t.Assert(e, nil)
		t.Assert(n, TableSize*2)
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
		}
		r, e := db.X创建Model对象(table).X设置数据(result).X插入并替换已存在()
		t.Assert(e, nil)
		n, e := r.RowsAffected()
		t.Assert(e, nil)
		t.Assert(n, TableSize*2)
	})
}

func Test_Model_Replace(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X设置数据(g.Map{
			"id":          1,
			"passport":    "t11",
			"password":    "25d55ad283aa400af464c76d713c07ad",
			"nickname":    "T11",
			"create_time": "2018-10-24 10:00:00",
		}).X插入并替换已存在()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
	})
}

func Test_Model_Save(t *testing.T) {
	table := createTable()
	defer dropTable(table)
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X设置数据(g.Map{
			"id":          1,
			"passport":    "t111",
			"password":    "25d55ad283aa400af464c76d713c07ad",
			"nickname":    "T111",
			"create_time": "2018-10-24 10:00:00",
		}).X插入并更新已存在()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
	})
}

func Test_Model_Update(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	// UPDATE...LIMIT
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X设置数据("nickname", "T100").X条件(1).X排序("id desc").X设置条数(2).X更新()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 2)

		v1, err := db.X创建Model对象(table).X字段保留过滤("nickname").X条件("id", 10).X查询一条值()
		t.AssertNil(err)
		t.Assert(v1.String(), "T100")

		v2, err := db.X创建Model对象(table).X字段保留过滤("nickname").X条件("id", 8).X查询一条值()
		t.AssertNil(err)
		t.Assert(v2.String(), "name_8")
	})

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
			X条件(1).X排序("id desc").X设置条数(2).
			X更新并取影响行数()
		t.AssertNil(err)
		t.Assert(n, 2)
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

func Test_Model_Fields(t *testing.T) {
	tableName1 := createInitTable()
	defer dropTable(tableName1)

	tableName2 := "user_" + 时间类.X创建并按当前时间().X取文本时间戳纳秒()
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
	    CREATE TABLE %s (
	        id         int(10) unsigned NOT NULL AUTO_INCREMENT,
	        name       varchar(45) NULL,
			age        int(10) unsigned,
	        PRIMARY KEY (id)
	    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
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
		count, err := db.X创建Model对象(table).X字段保留过滤("distinct id,nickname").X条件("id>8").X查询行数()
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

func Test_Model_Value_WithCache(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		value, err := db.X创建Model对象(table).X条件("id", 1).X缓存(db类.X缓存选项{
			X时长: time.Second * 10,
			X强制缓存:    false,
		}).X查询一条值()
		t.AssertNil(err)
		t.Assert(value.X取整数(), 0)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X设置数据(g.MapStrAny{
			"id":       1,
			"passport": fmt.Sprintf(`passport_%d`, 1),
			"password": fmt.Sprintf(`password_%d`, 1),
			"nickname": fmt.Sprintf(`nickname_%d`, 1),
		}).X插入()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		value, err := db.X创建Model对象(table).X条件("id", 1).X缓存(db类.X缓存选项{
			X时长: time.Second * 10,
			X强制缓存:    false,
		}).X查询一条值()
		t.AssertNil(err)
		t.Assert(value.X取整数(), 1)
	})
}

func Test_Model_Count_WithCache(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		count, err := db.X创建Model对象(table).X条件("id", 1).X缓存(db类.X缓存选项{
			X时长: time.Second * 10,
			X强制缓存:    false,
		}).X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(0))
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X设置数据(g.MapStrAny{
			"id":       1,
			"passport": fmt.Sprintf(`passport_%d`, 1),
			"password": fmt.Sprintf(`password_%d`, 1),
			"nickname": fmt.Sprintf(`nickname_%d`, 1),
		}).X插入()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		count, err := db.X创建Model对象(table).X条件("id", 1).X缓存(db类.X缓存选项{
			X时长: time.Second * 10,
			X强制缓存:    false,
		}).X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(1))
	})
}

func Test_Model_Count_All_WithCache(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		count, err := db.X创建Model对象(table).X缓存(db类.X缓存选项{
			X时长: time.Second * 10,
			X强制缓存:    false,
		}).X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(0))
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X设置数据(g.MapStrAny{
			"id":       1,
			"passport": fmt.Sprintf(`passport_%d`, 1),
			"password": fmt.Sprintf(`password_%d`, 1),
			"nickname": fmt.Sprintf(`nickname_%d`, 1),
		}).X插入()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		count, err := db.X创建Model对象(table).X缓存(db类.X缓存选项{
			X时长: time.Second * 10,
			X强制缓存:    false,
		}).X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(1))
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X设置数据(g.MapStrAny{
			"id":       2,
			"passport": fmt.Sprintf(`passport_%d`, 2),
			"password": fmt.Sprintf(`password_%d`, 2),
			"nickname": fmt.Sprintf(`nickname_%d`, 2),
		}).X插入()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		count, err := db.X创建Model对象(table).X缓存(db类.X缓存选项{
			X时长: time.Second * 10,
			X强制缓存:    false,
		}).X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(1))
	})
}

func Test_Model_CountColumn_WithCache(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		count, err := db.X创建Model对象(table).X条件("id", 1).X缓存(db类.X缓存选项{
			X时长: time.Second * 10,
			X强制缓存:    false,
		}).X查询字段行数("id")
		t.AssertNil(err)
		t.Assert(count, int64(0))
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X设置数据(g.MapStrAny{
			"id":       1,
			"passport": fmt.Sprintf(`passport_%d`, 1),
			"password": fmt.Sprintf(`password_%d`, 1),
			"nickname": fmt.Sprintf(`nickname_%d`, 1),
		}).X插入()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		count, err := db.X创建Model对象(table).X条件("id", 1).X缓存(db类.X缓存选项{
			X时长: time.Second * 10,
			X强制缓存:    false,
		}).X查询字段行数("id")
		t.AssertNil(err)
		t.Assert(count, int64(1))
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
		t.Assert(user.CreateTime.String(), "2018-10-24 10:00:00")
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
		t.Assert(user.CreateTime.String(), "2018-10-24 10:00:00")
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
		t.Assert(user.CreateTime.String(), "2018-10-24 10:00:00")
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
		t.Assert(user.CreateTime.String(), "2018-10-24 10:00:00")
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
		t.Assert(user.CreateTime.String(), "2018-10-24 10:00:00")
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
		t.Assert(users[0].CreateTime.String(), "2018-10-24 10:00:00")
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
		t.Assert(users[0].CreateTime.String(), "2018-10-24 10:00:00")
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
		t.Assert(users[0].CreateTime.String(), "2018-10-24 10:00:00")
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

	dbInvalid.X设置调试模式(true)
	defer dbInvalid.X设置调试模式(false)
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Uid      int `orm:"id"`
			Passport string
			Password string     `orm:"password"`
			Name     string     `orm:"nick_name"`
			Time     时间类.Time `orm:"create_time"`
		}
		var (
			users  []User
			buffer = bytes.NewBuffer(nil)
		)
		dbInvalid.X取日志记录器().(*日志类.Logger).X设置Writer(buffer)
		defer dbInvalid.X取日志记录器().(*日志类.Logger).X设置Writer(os.Stdout)
		dbInvalid.X创建Model对象(table).X排序("id asc").X查询到结构体指针(&users)
		// 打印buffer.String()的输出结果
		t.Assert(
			文本类.X是否包含(
				buffer.String(),
				"SELECT `id`,`Passport`,`password`,`nick_name`,`create_time` FROM `user",
			),
			true,
		)
	})

// 设置数据库调试模式为开启状态
// db.SetDebug(true)
// 在函数结束时，确保关闭数据库调试模式
// defer db.SetDebug(false)
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
		t.Assert(user.CreateTime.String(), "2018-10-24 10:00:00")
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
		t.Assert(user.CreateTime.String(), "2018-10-24 10:00:00")
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
		t.Assert(users[0].CreateTime.String(), "2018-10-24 10:00:00")
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
		t.Assert(users[0].CreateTime.String(), "2018-10-24 10:00:00")
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

	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X排序(db类.Raw("field(id, 10,1,2,3,4,5,6,7,8,9)")).X查询()
		t.AssertNil(err)
		t.Assert(len(result), TableSize)
		t.Assert(result[0]["nickname"].String(), "name_10")
		t.Assert(result[1]["nickname"].String(), "name_1")
		t.Assert(result[2]["nickname"].String(), "name_2")
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
				"id":       i,
				"passport": fmt.Sprintf(`passport_%d`, i),
				"password": fmt.Sprintf(`password_%d`, i),
				"nickname": fmt.Sprintf(`nickname_%d`, i),
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
				"id":       i,
				"passport": fmt.Sprintf(`passport_%d`, i),
				"password": fmt.Sprintf(`password_%d`, i),
				"nickname": fmt.Sprintf(`nickname_%d`, i),
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
	table := createInitTable()
	defer dropTable(table)

	// DELETE...LIMIT
	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件(1).X设置条数(2).X删除()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 2)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件(1).X删除()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, TableSize-2)
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
		r, err := db.X创建Model对象(table).X字段保留过滤("id, passport").X设置数据(g.Map{
			"id":       1,
			"passport": "1",
			"password": "1",
			"nickname": "1",
		}).X插入()
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 1)
		one, err := db.X创建Model对象(table).X条件("id", 1).X查询一条()
		t.AssertNil(err)
		t.AssertNE(one["password"].String(), "1")
		t.AssertNE(one["nickname"].String(), "1")
		t.Assert(one["passport"].String(), "1")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		table := createTable()
		defer dropTable(table)
		r, err := db.X创建Model对象(table).X过滤空值数据().X设置数据(g.Map{
			"id":       1,
			"passport": 0,
			"password": 0,
			"nickname": "1",
		}).X插入()
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 1)
		one, err := db.X创建Model对象(table).X条件("id", 1).X查询一条()
		t.AssertNil(err)
		t.AssertNE(one["passport"].String(), "0")
		t.AssertNE(one["password"].String(), "0")
		t.Assert(one["nickname"].String(), "1")
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

	// Save
	单元测试类.C(t, func(t *单元测试类.T) {
		table := createTable()
		defer dropTable(table)
		r, err := db.X创建Model对象(table).X字段保留过滤("id, passport").X设置数据(g.Map{
			"id":       1,
			"passport": "1",
			"password": "1",
			"nickname": "1",
		}).X插入并更新已存在()
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 1)
		one, err := db.X创建Model对象(table).X条件("id", 1).X查询一条()
		t.AssertNil(err)
		t.AssertNE(one["password"].String(), "1")
		t.AssertNE(one["nickname"].String(), "1")
		t.Assert(one["passport"].String(), "1")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		table := createTable()
		defer dropTable(table)
		_, err := db.X创建Model对象(table).X过滤空值数据().X设置数据(g.Map{
			"id":       1,
			"passport": 0,
			"password": 0,
			"nickname": "1",
		}).X插入并更新已存在()
		t.AssertNil(err)
		one, err := db.X创建Model对象(table).X条件("id", 1).X查询一条()
		t.AssertNil(err)
		t.AssertNE(one["passport"].String(), "0")
		t.AssertNE(one["password"].String(), "0")
		t.Assert(one["nickname"].String(), "1")

		_, err = db.X创建Model对象(table).X设置数据(g.Map{
			"id":       1,
			"passport": 0,
			"password": 0,
			"nickname": "1",
		}).X插入并更新已存在()
		t.AssertNil(err)
		one, err = db.X创建Model对象(table).X条件("id", 1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["passport"].String(), "0")
		t.Assert(one["password"].String(), "0")
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

func Test_Model_Option_List(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		table := createTable()
		defer dropTable(table)
		r, err := db.X创建Model对象(table).X字段保留过滤("id, password").X设置数据(g.Map数组{
			g.Map{
				"id":       1,
				"passport": "1",
				"password": "1",
				"nickname": "1",
			},
			g.Map{
				"id":       2,
				"passport": "2",
				"password": "2",
				"nickname": "2",
			},
		}).X插入并更新已存在()
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 2)
		list, err := db.X创建Model对象(table).X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(list), 2)
		t.Assert(list[0]["id"].String(), "1")
		t.Assert(list[0]["nickname"].String(), "")
		t.Assert(list[0]["passport"].String(), "")
		t.Assert(list[0]["password"].String(), "1")

		t.Assert(list[1]["id"].String(), "2")
		t.Assert(list[1]["nickname"].String(), "")
		t.Assert(list[1]["passport"].String(), "")
		t.Assert(list[1]["password"].String(), "2")
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		table := createTable()
		defer dropTable(table)
		r, err := db.X创建Model对象(table).X过滤空值().X字段保留过滤("id, password").X设置数据(g.Map数组{
			g.Map{
				"id":       1,
				"passport": "1",
				"password": 0,
				"nickname": "1",
			},
			g.Map{
				"id":       2,
				"passport": "2",
				"password": "2",
				"nickname": "2",
			},
		}).X插入并更新已存在()
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 2)
		list, err := db.X创建Model对象(table).X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(list), 2)
		t.Assert(list[0]["id"].String(), "1")
		t.Assert(list[0]["nickname"].String(), "")
		t.Assert(list[0]["passport"].String(), "")
		t.Assert(list[0]["password"].String(), "0")

		t.Assert(list[1]["id"].String(), "2")
		t.Assert(list[1]["nickname"].String(), "")
		t.Assert(list[1]["passport"].String(), "")
		t.Assert(list[1]["password"].String(), "2")
	})
}

func Test_Model_OmitEmpty(t *testing.T) {
	table := fmt.Sprintf(`table_%s`, 时间类.X取文本时间戳纳秒())
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
    CREATE TABLE IF NOT EXISTS %s (
        id int(10) unsigned NOT NULL AUTO_INCREMENT,
        name varchar(45) NOT NULL,
        PRIMARY KEY (id)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, table)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		_, err := db.X创建Model对象(table).X过滤空值().X设置数据(g.Map{
			"id":   1,
			"name": "",
		}).X插入并更新已存在()
		t.AssertNE(err, nil)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		_, err := db.X创建Model对象(table).X过滤空值数据().X设置数据(g.Map{
			"id":   1,
			"name": "",
		}).X插入并更新已存在()
		t.AssertNE(err, nil)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		_, err := db.X创建Model对象(table).X过滤空值条件().X设置数据(g.Map{
			"id":   1,
			"name": "",
		}).X插入并更新已存在()
		t.AssertNil(err)
	})
}

func Test_Model_OmitNil(t *testing.T) {
	table := fmt.Sprintf(`table_%s`, 时间类.X取文本时间戳纳秒())
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
    CREATE TABLE IF NOT EXISTS %s (
        id int(10) unsigned NOT NULL AUTO_INCREMENT,
        name varchar(45) NOT NULL,
        PRIMARY KEY (id)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, table)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		_, err := db.X创建Model对象(table).X过滤Nil().X设置数据(g.Map{
			"id":   1,
			"name": nil,
		}).X插入并更新已存在()
		t.AssertNE(err, nil)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		_, err := db.X创建Model对象(table).X过滤Nil().X设置数据(g.Map{
			"id":   1,
			"name": "",
		}).X插入并更新已存在()
		t.AssertNil(err)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		_, err := db.X创建Model对象(table).X过滤Nil条件().X设置数据(g.Map{
			"id":   1,
			"name": "",
		}).X插入并更新已存在()
		t.AssertNil(err)
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
	return
	单元测试类.C(t, func(t *单元测试类.T) {
		table := createInitTable()
		defer dropTable(table)
		r, err := db.X创建Model对象(table).X过滤空值().X设置数据("nickname", 1).X条件(g.Map{"id": 1, "passport": ""}).X更新()
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 1)

		v, err := db.X创建Model对象(table).X条件("id", 1).X字段保留过滤("nickname").X查询一条值()
		t.AssertNil(err)
		t.Assert(v.String(), "1")
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

func Test_Model_FieldsEx_WithReservedWords(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			table      = "fieldsex_test_table"
			sqlTpcPath = 单元测试类.DataPath("reservedwords_table_tpl.sql")
			sqlContent = 文件类.X读文本(sqlTpcPath)
		)
		t.AssertNE(sqlContent, "")
		if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(sqlContent, table)); err != nil {
			t.AssertNil(err)
		}
		defer dropTable(table)
		_, err := db.X创建Model对象(table).X字段排除过滤("content").X查询一条()
		t.AssertNil(err)
	})
}

func Test_Model_Prefix(t *testing.T) {
	db := dbPrefix
	table := fmt.Sprintf(`%s_%d`, TableName, 时间类.X取时间戳纳秒())
	createInitTableWithDb(db, TableNamePrefix1+table)
	defer dropTable(TableNamePrefix1 + table)
	// Select.
	单元测试类.C(t, func(t *单元测试类.T) {
		r, err := db.X创建Model对象(table).X条件("id in (?)", g.Slice别名{1, 2}).X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(r), 2)
		t.Assert(r[0]["id"], "1")
		t.Assert(r[1]["id"], "2")
	})
	// 使用别名进行选择。
	单元测试类.C(t, func(t *单元测试类.T) {
		r, err := db.X创建Model对象(table+" as u").X条件("u.id in (?)", g.Slice别名{1, 2}).X排序("u.id asc").X查询()
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
		err := db.X创建Model对象(table+" u").X条件("u.id in (?)", g.Slice别名{1, 5}).X排序("u.id asc").X查询到结构体指针(&users)
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].Id, 1)
		t.Assert(users[1].Id, 5)
	})
	// 使用别名和连接语句进行选择。
	单元测试类.C(t, func(t *单元测试类.T) {
		r, err := db.X创建Model对象(table+" as u1").X左连接(table+" as u2", "u2.id=u1.id").X条件("u1.id in (?)", g.Slice别名{1, 2}).X排序("u1.id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(r), 2)
		t.Assert(r[0]["id"], "1")
		t.Assert(r[1]["id"], "2")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		r, err := db.X创建Model对象(table).X设置表别名("u1").X左连接(table+" as u2", "u2.id=u1.id").X条件("u1.id in (?)", g.Slice别名{1, 2}).X排序("u1.id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(r), 2)
		t.Assert(r[0]["id"], "1")
		t.Assert(r[1]["id"], "2")
	})
}

func Test_Model_Schema1(t *testing.T) {
	// db.SetDebug(true)

	db = db.X切换数据库(TestSchema1)
	table := fmt.Sprintf(`%s_%s`, TableName, 时间类.X取文本时间戳纳秒())
	createInitTableWithDb(db, table)
	db = db.X切换数据库(TestSchema2)
	createInitTableWithDb(db, table)
	defer func() {
		db = db.X切换数据库(TestSchema1)
		dropTableWithDb(db, table)
		db = db.X切换数据库(TestSchema2)
		dropTableWithDb(db, table)
		db = db.X切换数据库(TestSchema1)
	}()
	// Method.
	单元测试类.C(t, func(t *单元测试类.T) {
		db = db.X切换数据库(TestSchema1)
		r, err := db.X创建Model对象(table).X更新(g.Map{"nickname": "name_100"}, "id=1")
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 1)

		v, err := db.X创建Model对象(table).X查询一条值("nickname", "id=1")
		t.AssertNil(err)
		t.Assert(v.String(), "name_100")

		db = db.X切换数据库(TestSchema2)
		v, err = db.X创建Model对象(table).X查询一条值("nickname", "id=1")
		t.AssertNil(err)
		t.Assert(v.String(), "name_1")
	})
	// Model.
	单元测试类.C(t, func(t *单元测试类.T) {
		v, err := db.X创建Model对象(table).X切换数据库(TestSchema1).X查询一条值("nickname", "id=2")
		t.AssertNil(err)
		t.Assert(v.String(), "name_2")

		r, err := db.X创建Model对象(table).X切换数据库(TestSchema1).X更新(g.Map{"nickname": "name_200"}, "id=2")
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 1)

		v, err = db.X创建Model对象(table).X切换数据库(TestSchema1).X查询一条值("nickname", "id=2")
		t.AssertNil(err)
		t.Assert(v.String(), "name_200")

		v, err = db.X创建Model对象(table).X切换数据库(TestSchema2).X查询一条值("nickname", "id=2")
		t.AssertNil(err)
		t.Assert(v.String(), "name_2")

		v, err = db.X创建Model对象(table).X切换数据库(TestSchema1).X查询一条值("nickname", "id=2")
		t.AssertNil(err)
		t.Assert(v.String(), "name_200")
	})
	// Model.
	单元测试类.C(t, func(t *单元测试类.T) {
		i := 1000
		_, err := db.X创建Model对象(table).X切换数据库(TestSchema1).X插入(g.Map{
			"id":               i,
			"passport":         fmt.Sprintf(`user_%d`, i),
			"password":         fmt.Sprintf(`pass_%d`, i),
			"nickname":         fmt.Sprintf(`name_%d`, i),
			"create_time":      时间类.X创建并从文本("2018-10-24 10:00:00").String(),
			"none-exist-field": 1,
		})
		t.AssertNil(err)

		v, err := db.X创建Model对象(table).X切换数据库(TestSchema1).X查询一条值("nickname", "id=?", i)
		t.AssertNil(err)
		t.Assert(v.String(), "name_1000")

		v, err = db.X创建Model对象(table).X切换数据库(TestSchema2).X查询一条值("nickname", "id=?", i)
		t.AssertNil(err)
		t.Assert(v.String(), "")
	})
}

func Test_Model_Schema2(t *testing.T) {
	// db.SetDebug(true)

	db = db.X切换数据库(TestSchema1)
	table := fmt.Sprintf(`%s_%s`, TableName, 时间类.X取文本时间戳纳秒())
	createInitTableWithDb(db, table)
	db = db.X切换数据库(TestSchema2)
	createInitTableWithDb(db, table)
	defer func() {
		db = db.X切换数据库(TestSchema1)
		dropTableWithDb(db, table)
		db = db.X切换数据库(TestSchema2)
		dropTableWithDb(db, table)

		db = db.X切换数据库(TestSchema1)
	}()
	// Schema.
	单元测试类.C(t, func(t *单元测试类.T) {
		v, err := db.X切换数据库(TestSchema1).X创建Model对象(table).X查询一条值("nickname", "id=2")
		t.AssertNil(err)
		t.Assert(v.String(), "name_2")

		r, err := db.X切换数据库(TestSchema1).X创建Model对象(table).X更新(g.Map{"nickname": "name_200"}, "id=2")
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 1)

		v, err = db.X切换数据库(TestSchema1).X创建Model对象(table).X查询一条值("nickname", "id=2")
		t.AssertNil(err)
		t.Assert(v.String(), "name_200")

		v, err = db.X切换数据库(TestSchema2).X创建Model对象(table).X查询一条值("nickname", "id=2")
		t.AssertNil(err)
		t.Assert(v.String(), "name_2")

		v, err = db.X切换数据库(TestSchema1).X创建Model对象(table).X查询一条值("nickname", "id=2")
		t.AssertNil(err)
		t.Assert(v.String(), "name_200")
	})
	// Schema.
	单元测试类.C(t, func(t *单元测试类.T) {
		i := 1000
		_, err := db.X切换数据库(TestSchema1).X创建Model对象(table).X插入(g.Map{
			"id":               i,
			"passport":         fmt.Sprintf(`user_%d`, i),
			"password":         fmt.Sprintf(`pass_%d`, i),
			"nickname":         fmt.Sprintf(`name_%d`, i),
			"create_time":      时间类.X创建并从文本("2018-10-24 10:00:00").String(),
			"none-exist-field": 1,
		})
		t.AssertNil(err)

		v, err := db.X切换数据库(TestSchema1).X创建Model对象(table).X查询一条值("nickname", "id=?", i)
		t.AssertNil(err)
		t.Assert(v.String(), "name_1000")

		v, err = db.X切换数据库(TestSchema2).X创建Model对象(table).X查询一条值("nickname", "id=?", i)
		t.AssertNil(err)
		t.Assert(v.String(), "")
	})
}

func Test_Model_FieldsExStruct(t *testing.T) {
	table := createTable()
	defer dropTable(table)
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Id       int    `orm:"id"       json:"id"`
			Passport string `orm:"password" json:"pass_port"`
			Password string `orm:"password" json:"password"`
			NickName string `orm:"nickname" json:"nick__name"`
		}
		user := &User{
			Id:       1,
			Passport: "111",
			Password: "222",
			NickName: "333",
		}
		r, err := db.X创建Model对象(table).X字段排除过滤("create_time, password").X过滤空值().X设置数据(user).X插入()
		t.AssertNil(err)
		n, err := r.RowsAffected()
		t.AssertNil(err)
		t.Assert(n, 1)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Id       int    `orm:"id"       json:"id"`
			Passport string `orm:"password" json:"pass_port"`
			Password string `orm:"password" json:"password"`
			NickName string `orm:"nickname" json:"nick__name"`
		}
		users := make([]*User, 0)
		for i := 100; i < 110; i++ {
			users = append(users, &User{
				Id:       i,
				Passport: fmt.Sprintf(`passport_%d`, i),
				Password: fmt.Sprintf(`password_%d`, i),
				NickName: fmt.Sprintf(`nickname_%d`, i),
			})
		}
		r, err := db.X创建Model对象(table).X字段排除过滤("create_time, password").
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
		r, err := db.X创建Model对象(table).X过滤空值().X设置数据(user).X条件并识别主键(1).X更新()
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
		all, err := db.X创建Model对象(table).X条件("id > 1").X设置分组条件("id > 8").X查询()
		t.AssertNil(err)
		t.Assert(len(all), 2)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		all, err := db.X创建Model对象(table).X条件("id > 1").X设置分组条件("id > ?", 8).X查询()
		t.AssertNil(err)
		t.Assert(len(all), 2)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		all, err := db.X创建Model对象(table).X条件("id > ?", 1).X设置分组条件("id > ?", 8).X查询()
		t.AssertNil(err)
		t.Assert(len(all), 2)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		all, err := db.X创建Model对象(table).X条件("id > ?", 1).X设置分组条件("id", 8).X查询()
		t.AssertNil(err)
		t.Assert(len(all), 1)
	})
}

func Test_Model_Distinct(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		all, err := db.X创建Model对象(table, "t").X字段保留过滤("distinct t.id").X条件("id > 1").X设置分组条件("id > 8").X查询()
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

// "id":          i, // "id"字段：i的值
// "passport":    fmt.Sprintf(`user_%d`, i), // "passport"字段：格式化输出字符串，形如"user_1"，其中%d用i的值替换
// "password":    fmt.Sprintf(`pass_%d`, i), // "password"字段：格式化输出字符串，形如"pass_1"，其中%d用i的值替换
// "nickname":    fmt.Sprintf(`name_%d`, i), // "nickname"字段：格式化输出字符串，形如"name_1"，其中%d用i的值替换
// "create_time": gtime.NewFromStr("2018-10-24 10:00:00").String(), // "create_time"字段：创建一个时间对象，使用"2018-10-24 10:00:00"字符串初始化，并将其转换为字符串表示形式

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

func Test_Model_NullField(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Id       int
			Passport *string
		}
		data := g.Map{
			"id":       1,
			"passport": nil,
		}
		result, err := db.X创建Model对象(table).X设置数据(data).X插入()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)

		var user *User
		err = one.X取结构体指针(&user)
		t.AssertNil(err)
		t.Assert(user.Id, data["id"])
		t.Assert(user.Passport, data["passport"])
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

func createTableForTimeZoneTest() string {
	tableName := "user_" + 时间类.X创建并按当前时间().X取文本时间戳纳秒()
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
	    CREATE TABLE %s (
	        id          int(10) unsigned NOT NULL AUTO_INCREMENT,
	        passport    varchar(45) NULL,
	        password    char(32) NULL,
	        nickname    varchar(45) NULL,
	        created_at timestamp(6) NULL,
 			updated_at timestamp(6) NULL,
			deleted_at timestamp(6) NULL,
	        PRIMARY KEY (id)
	    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
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
		t.Assert(userEntity.CreatedAt.String(), "2020-11-22 11:23:45")
		t.Assert(userEntity.UpdatedAt.String(), "2020-11-22 12:23:45")
		t.Assert(时间类.X创建并按Time(userEntity.DeletedAt).String(), "2020-11-22 13:23:45")
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

func Test_Model_WhereOrNull(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := db.X创建Model对象(table).X条件或NULL值("nickname").X条件或NULL值("passport").X排序ASC("id").X排序随机().X查询()
		t.AssertNil(err)
		t.Assert(len(result), 0)
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
			"id":       1,
			"passport": "user_1",
			"password": "pass_1",
			"nickname": "name_1",
		}).X插入并取ID()
		t.AssertNil(err)
		t.Assert(id, 1)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		id, err := db.X创建Model对象(table).X设置数据(g.Map{
			"passport": "user_2",
			"password": "pass_2",
			"nickname": "name_2",
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

func Test_Model_OnDuplicate(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	// string.
	单元测试类.C(t, func(t *单元测试类.T) {
		data := g.Map{
			"id":          1,
			"passport":    "pp1",
			"password":    "pw1",
			"nickname":    "n1",
			"create_time": "2016-06-06",
		}
		_, err := db.X创建Model对象(table).X设置插入冲突更新字段("passport,password").X设置数据(data).X插入并更新已存在()
		t.AssertNil(err)
		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["passport"], data["passport"])
		t.Assert(one["password"], data["password"])
		t.Assert(one["nickname"], "name_1")
	})

	// slice.
	单元测试类.C(t, func(t *单元测试类.T) {
		data := g.Map{
			"id":          1,
			"passport":    "pp1",
			"password":    "pw1",
			"nickname":    "n1",
			"create_time": "2016-06-06",
		}
		_, err := db.X创建Model对象(table).X设置插入冲突更新字段(g.Slice别名{"passport", "password"}).X设置数据(data).X插入并更新已存在()
		t.AssertNil(err)
		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["passport"], data["passport"])
		t.Assert(one["password"], data["password"])
		t.Assert(one["nickname"], "name_1")
	})

	// map.
	单元测试类.C(t, func(t *单元测试类.T) {
		data := g.Map{
			"id":          1,
			"passport":    "pp1",
			"password":    "pw1",
			"nickname":    "n1",
			"create_time": "2016-06-06",
		}
		_, err := db.X创建Model对象(table).X设置插入冲突更新字段(g.Map{
			"passport": "nickname",
			"password": "nickname",
		}).X设置数据(data).X插入并更新已存在()
		t.AssertNil(err)
		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["passport"], data["nickname"])
		t.Assert(one["password"], data["nickname"])
		t.Assert(one["nickname"], "name_1")
	})

	// map+raw.
	单元测试类.C(t, func(t *单元测试类.T) {
		data := g.MapStrStr{
			"id":          "1",
			"passport":    "pp1",
			"password":    "pw1",
			"nickname":    "n1",
			"create_time": "2016-06-06",
		}
		_, err := db.X创建Model对象(table).X设置插入冲突更新字段(g.Map{
			"passport": db类.Raw("CONCAT(VALUES(`passport`), '1')"),
			"password": db类.Raw("CONCAT(VALUES(`password`), '2')"),
		}).X设置数据(data).X插入并更新已存在()
		t.AssertNil(err)
		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["passport"], data["passport"]+"1")
		t.Assert(one["password"], data["password"]+"2")
		t.Assert(one["nickname"], "name_1")
	})
}

func Test_Model_OnDuplicateEx(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	// string.
	单元测试类.C(t, func(t *单元测试类.T) {
		data := g.Map{
			"id":          1,
			"passport":    "pp1",
			"password":    "pw1",
			"nickname":    "n1",
			"create_time": "2016-06-06",
		}
		_, err := db.X创建Model对象(table).X设置插入冲突不更新字段("nickname,create_time").X设置数据(data).X插入并更新已存在()
		t.AssertNil(err)
		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["passport"], data["passport"])
		t.Assert(one["password"], data["password"])
		t.Assert(one["nickname"], "name_1")
	})

	// slice.
	单元测试类.C(t, func(t *单元测试类.T) {
		data := g.Map{
			"id":          1,
			"passport":    "pp1",
			"password":    "pw1",
			"nickname":    "n1",
			"create_time": "2016-06-06",
		}
		_, err := db.X创建Model对象(table).X设置插入冲突不更新字段(g.Slice别名{"nickname", "create_time"}).X设置数据(data).X插入并更新已存在()
		t.AssertNil(err)
		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["passport"], data["passport"])
		t.Assert(one["password"], data["password"])
		t.Assert(one["nickname"], "name_1")
	})

	// map.
	单元测试类.C(t, func(t *单元测试类.T) {
		data := g.Map{
			"id":          1,
			"passport":    "pp1",
			"password":    "pw1",
			"nickname":    "n1",
			"create_time": "2016-06-06",
		}
		_, err := db.X创建Model对象(table).X设置插入冲突不更新字段(g.Map{
			"nickname":    "nickname",
			"create_time": "nickname",
		}).X设置数据(data).X插入并更新已存在()
		t.AssertNil(err)
		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["passport"], data["passport"])
		t.Assert(one["password"], data["password"])
		t.Assert(one["nickname"], "name_1")
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
		type Input struct {
			Id []int
		}
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
			Id:       1,
			Passport: "user_1",
			Password: "pass_1",
			Nickname: "name_1",
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

// 从GoFrame v1.16.0开始，此功能不再使用，因为过滤特性会自动启用。
// func Test_Model_Insert_KeyFieldNameMapping_Error(t *testing.T) {
// 	// 创建测试表
// 	table := createTable()
// 	// 在测试结束后删除测试表
// 	defer dropTable(table)
//
// 	// 使用gtest进行单元测试
// 	gtest.C(t, func(t *gtest.T) {
//  	// 定义User结构体
//  	type User struct {
//  		Id             int
//  		Passport       string
//  		Password       string
//  		Nickname       string
//  		CreateTime     string
//  		NoneExistField string // 不存在的字段
//  	}
//  	// 初始化用户数据
//  	data := User{
//  		Id:         1,
//  		Passport:   "user_1",
//  		Password:   "pass_1",
//  		Nickname:   "name_1",
//  		CreateTime: "2020-10-10 12:00:01",
//  	}
//  	// 尝试将数据插入到指定表中
//  	result, err := db.Model(table).Data(data).Insert()
//  	// 断言错误不为空
//  	t.AssertNE(err, nil)
//  })
// }

func Test_Model_Fields_AutoFilterInJoinStatement(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var err error
		table1 := "user"
		table2 := "score"
		table3 := "info"
		if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
		   id int(11) NOT NULL AUTO_INCREMENT,
		   name varchar(500) NOT NULL DEFAULT '',
		 PRIMARY KEY (id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;
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
			id int(11) NOT NULL AUTO_INCREMENT,
			user_id int(11) NOT NULL DEFAULT 0,
		    number varchar(500) NOT NULL DEFAULT '',
		 PRIMARY KEY (id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;
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
			id int(11) NOT NULL AUTO_INCREMENT,
			user_id int(11) NOT NULL DEFAULT 0,
		    description varchar(500) NOT NULL DEFAULT '',
		 PRIMARY KEY (id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;
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
		table1 = 时间类.X取文本时间戳纳秒() + "_table1"
		table2 = 时间类.X取文本时间戳纳秒() + "_table2"
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
		table1 = 时间类.X取文本时间戳纳秒() + "_table1"
		table2 = 时间类.X取文本时间戳纳秒() + "_table2"
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
		table1 = 时间类.X取文本时间戳纳秒() + "_table1"
		table2 = 时间类.X取文本时间戳纳秒() + "_table2"
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

func Test_Builder_OmitEmptyWhere(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	单元测试类.C(t, func(t *单元测试类.T) {
		count, err := db.X创建Model对象(table).X条件("id", 1).X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(1))
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		count, err := db.X创建Model对象(table).X条件("id", 0).X过滤空值条件().X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(TableSize))
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		builder := db.X创建Model对象(table).X过滤空值条件().X创建组合条件()
		count, err := db.X创建Model对象(table).X条件(
			builder.X条件("id", 0),
		).X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(TableSize))
	})
}

func Test_Scan_Nil_Result_Error(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	type S struct {
		Id    int
		Name  string
		Age   int
		Score int
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		var s *S
		err := db.X创建Model对象(table).X条件("id", 1).X查询到结构体指针(&s)
		t.AssertNil(err)
		t.Assert(s.Id, 1)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var s *S
		err := db.X创建Model对象(table).X条件("id", 100).X查询到结构体指针(&s)
		t.AssertNil(err)
		t.Assert(s, nil)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var s S
		err := db.X创建Model对象(table).X条件("id", 100).X查询到结构体指针(&s)
		t.Assert(err, sql.ErrNoRows)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var ss []*S
		err := db.X创建Model对象(table).X查询到结构体指针(&ss)
		t.AssertNil(err)
		t.Assert(len(ss), TableSize)
	})
	// 如果结果为空，则返回错误。
	单元测试类.C(t, func(t *单元测试类.T) {
		var ss = make([]*S, 10)
		err := db.X创建Model对象(table).X条件大于("id", 100).X查询到结构体指针(&ss)
		t.Assert(err, sql.ErrNoRows)
	})
}

func Test_Model_FixGdbJoin(t *testing.T) {
	array := 文本类.X分割并忽略空值(单元测试类.DataContent(`fix_gdb_join.sql`), ";")
	for _, v := range array {
		if _, err := db.X原生SQL执行(ctx, v); err != nil {
			单元测试类.Error(err)
		}
	}
	defer dropTable(`common_resource`)
	defer dropTable(`managed_resource`)
	defer dropTable(`rules_template`)
	defer dropTable(`resource_mark`)
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertNil(db.X取Core对象().X删除所有表查询缓存(ctx))
		sqlSlice, err := db类.X捕捉SQL语句(ctx, func(ctx context.Context) error {
			orm := db.X创建Model对象(`managed_resource`).X设置上下文并取副本(ctx).
				X左连接相同字段(`common_resource`, `resource_id`).
				X左连接带比较运算符(`resource_mark`, `resource_mark_id`, `=`, `id`).
				X左连接带比较运算符(`rules_template`, `rule_template_id`, `=`, `template_id`).
				X字段保留过滤并带前缀(
					`managed_resource`,
					"resource_id", "user", "status", "status_message", "safe_publication", "rule_template_id",
					"created_at", "comments", "expired_at", "resource_mark_id", "instance_id", "resource_name",
					"pay_mode").
				X字段保留过滤并带前缀(`resource_mark`, "mark_name", "color").
				X字段保留过滤并带前缀(`rules_template`, "name").
				X字段保留过滤并带前缀(`common_resource`, `src_instance_id`, "database_kind", "source_type", "ip", "port")
			all, err := orm.X排序ASC("src_instance_id").X查询()
			t.Assert(len(all), 4)
			t.Assert(all[0]["pay_mode"], 1)
			t.Assert(all[0]["src_instance_id"], 2)
			t.Assert(all[3]["instance_id"], "dmcins-jxy0x75m")
			t.Assert(all[3]["src_instance_id"], "vdb-6b6m3u1u")
			t.Assert(all[3]["resource_mark_id"], "11")
			return err
		})
		t.AssertNil(err)

		t.Assert(单元测试类.DataContent(`fix_gdb_join_expect.sql`), sqlSlice[len(sqlSlice)-1])
	})
}
