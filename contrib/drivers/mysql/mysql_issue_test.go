// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package mysql_test

import (
	"context"
	"fmt"
	"testing"
	"time"
	
	"github.com/888go/goframe/database/gdb"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/text/gregex"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gmeta"
	"github.com/888go/goframe/util/guid"
)

// 这是Go语言代码中的一行注释，其内容为一个URL链接，指向GitHub上gogf/gf项目的一个问题编号1934。
// 中文翻译：
// 参考GitHub上gogf/gf项目的问题1934。
func Test_Issue1934(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		one, err := db.X创建Model对象(table).X条件(" id ", 1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["id"], 1)
	})
}

// 这是Go语言代码中的一行注释，其内容引用了GitHub上gogf/gf仓库下的第1570号issue。
// 翻译成中文：
// 引用了GitHub上gogf/gf项目中的第1570号问题。
func Test_Issue1570(t *testing.T) {
	var (
		tableUser       = "user_" + 时间类.X取文本时间戳微秒()
		tableUserDetail = "user_detail_" + 时间类.X取文本时间戳微秒()
		tableUserScores = "user_scores_" + 时间类.X取文本时间戳微秒()
	)
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUser)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUser)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  address varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUserDetail)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserDetail)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  id int(10) unsigned NOT NULL AUTO_INCREMENT,
  uid int(10) unsigned NOT NULL,
  score int(10) unsigned NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUserScores)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserScores)

	type EntityUser struct {
		Uid  int    `json:"uid"`
		Name string `json:"name"`
	}
	type EntityUserDetail struct {
		Uid     int    `json:"uid"`
		Address string `json:"address"`
	}
	type EntityUserScores struct {
		Id    int `json:"id"`
		Uid   int `json:"uid"`
		Score int `json:"score"`
	}
	type Entity struct {
		User       *EntityUser
		UserDetail *EntityUserDetail
		UserScores []*EntityUserScores
	}

	// 初始化数据
	单元测试类.C(t, func(t *单元测试类.T) {
		var err error
		for i := 1; i <= 5; i++ {
			// User.
			_, err = db.X插入(ctx, tableUser, g.Map{
				"uid":  i,
				"name": fmt.Sprintf(`name_%d`, i),
			})
			t.AssertNil(err)
			// Detail.
			_, err = db.X插入(ctx, tableUserDetail, g.Map{
				"uid":     i,
				"address": fmt.Sprintf(`address_%d`, i),
			})
			t.AssertNil(err)
			// Scores.
			for j := 1; j <= 5; j++ {
				_, err = db.X插入(ctx, tableUserScores, g.Map{
					"uid":   i,
					"score": j,
				})
				t.AssertNil(err)
			}
		}
	})

	// Result ScanList，用于包含结构体元素和指针属性。
	单元测试类.C(t, func(t *单元测试类.T) {
		var users []Entity
		// User
		err := db.X创建Model对象(tableUser).
			X条件("uid", g.Slice别名{3, 4}).
			X字段保留过滤("uid").
			X排序("uid asc").
			X查询到指针列表(&users, "User")
		t.AssertNil(err)
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].User, &EntityUser{3, ""})
		t.Assert(users[1].User, &EntityUser{4, ""})
		// Detail
		err = db.X创建Model对象(tableUserDetail).
			X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).
			X排序("uid asc").
			X查询到指针列表(&users, "UserDetail", "User", "uid:Uid")
		t.AssertNil(err)
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, &EntityUserDetail{3, "address_3"})
		t.Assert(users[1].UserDetail, &EntityUserDetail{4, "address_4"})
		// Scores
		err = db.X创建Model对象(tableUserScores).
			X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).
			X排序("id asc").
			X查询到指针列表(&users, "UserScores", "User", "uid:Uid")
		t.AssertNil(err)
		t.AssertNil(err)
		t.Assert(len(users[0].UserScores), 5)
		t.Assert(len(users[1].UserScores), 5)
		t.Assert(users[0].UserScores[0].Uid, 3)
		t.Assert(users[0].UserScores[0].Score, 1)
		t.Assert(users[0].UserScores[4].Score, 5)
		t.Assert(users[1].UserScores[0].Uid, 4)
		t.Assert(users[1].UserScores[0].Score, 1)
		t.Assert(users[1].UserScores[4].Score, 5)
	})
}

// 这是Go语言代码中的一行注释，其内容引用了GitHub上gogf/gf仓库的第1401个issue（问题）。
// 中文翻译：
// 引用了GitHub上gogf/gf项目的问题1401。
func Test_Issue1401(t *testing.T) {
	var (
		table1 = "parcels"
		table2 = "parcel_items"
	)
	array := 文本类.X分割并忽略空值(单元测试类.DataContent(`issue1401.sql`), ";")
	for _, v := range array {
		if _, err := db.X原生SQL执行(ctx, v); err != nil {
			单元测试类.Error(err)
		}
	}
	defer dropTable(table1)
	defer dropTable(table2)

	单元测试类.C(t, func(t *单元测试类.T) {
		type NItem struct {
			Id       int `json:"id"`
			ParcelId int `json:"parcel_id"`
		}

		type ParcelItem struct {
			元数据类.Meta `orm:"table:parcel_items"`
			NItem
		}

		type ParcelRsp struct {
			元数据类.Meta `orm:"table:parcels"`
			Id         int           `json:"id"`
			Items      []*ParcelItem `json:"items" orm:"with:parcel_id=Id"`
		}

		parcelDetail := &ParcelRsp{}
		err := db.X创建Model对象(table1).X关联对象(parcelDetail.Items).X条件("id", 3).X查询到结构体指针(&parcelDetail)
		t.AssertNil(err)
		t.Assert(parcelDetail.Id, 3)
		t.Assert(len(parcelDetail.Items), 1)
		t.Assert(parcelDetail.Items[0].Id, 2)
		t.Assert(parcelDetail.Items[0].ParcelId, 3)
	})
}

// 这是GitHub上gogf/gf仓库的第1412号问题链接
func Test_Issue1412(t *testing.T) {
	var (
		table1 = "parcels"
		table2 = "items"
	)
	array := 文本类.X分割并忽略空值(单元测试类.DataContent(`issue1412.sql`), ";")
	for _, v := range array {
		if _, err := db.X原生SQL执行(ctx, v); err != nil {
			单元测试类.Error(err)
		}
	}
	defer dropTable(table1)
	defer dropTable(table2)

	单元测试类.C(t, func(t *单元测试类.T) {
		type Items struct {
			元数据类.Meta `orm:"table:items"`
			Id         int    `json:"id"`
			Name       string `json:"name"`
		}

		type ParcelRsp struct {
			元数据类.Meta `orm:"table:parcels"`
			Id         int   `json:"id"`
			ItemId     int   `json:"item_id"`
			Items      Items `json:"items" orm:"with:Id=ItemId"`
		}

		entity := &ParcelRsp{}
		err := db.X创建Model对象("parcels").X关联对象(Items{}).X条件("id=3").X查询到结构体指针(&entity)
		t.AssertNil(err)
		t.Assert(entity.Id, 3)
		t.Assert(entity.ItemId, 0)
		t.Assert(entity.Items.Id, 0)
		t.Assert(entity.Items.Name, "")
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		type Items struct {
			元数据类.Meta `orm:"table:items"`
			Id         int    `json:"id"`
			Name       string `json:"name"`
		}

		type ParcelRsp struct {
			元数据类.Meta `orm:"table:parcels"`
			Id         int   `json:"id"`
			ItemId     int   `json:"item_id"`
			Items      Items `json:"items" orm:"with:Id=ItemId"`
		}

		entity := &ParcelRsp{}
		err := db.X创建Model对象("parcels").X关联对象(Items{}).X条件("id=30000").X查询到结构体指针(&entity)
		t.AssertNE(err, nil)
		t.Assert(entity.Id, 0)
		t.Assert(entity.ItemId, 0)
		t.Assert(entity.Items.Id, 0)
		t.Assert(entity.Items.Name, "")
	})
}

// 这是Go语言代码中的一行注释，其内容为一个GitHub仓库的issue链接。
// 翻译为：
// 参考GitHub上gogf/gf项目中的第1002号问题。
func Test_Issue1002(t *testing.T) {
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
	// 函数参数中包含 where + time.Time 类型，时间使用 UTC（协调世界时）。
	单元测试类.C(t, func(t *单元测试类.T) {
		t1, _ := time.Parse("2006-01-02 15:04:05", "2020-10-27 11:03:32")
		t2, _ := time.Parse("2006-01-02 15:04:05", "2020-10-27 11:03:34")
		{
			v, err := db.X创建Model对象(table).X字段保留过滤("id").X条件("create_time>? and create_time<?", t1, t2).X查询一条值()
			t.AssertNil(err)
			t.Assert(v.X取整数(), 1)
		}
	})
// 此处使用了+8时区的时间参数。
// gtest.C(t, func(t *gtest.T) {
//// 将当前时区更改为+8时区（即中国北京时间）。
// location, err := time.LoadLocation("Asia/Shanghai")
// t.AssertNil(err) // 断言加载时区无错误
// t1, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-10-27 19:03:32", location) // 解析字符串为指定时区的time.Time类型
// t2, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-10-27 19:03:34", location)
// 
// // 使用create_time字段在t1和t2时间范围内的查询条件，获取id字段值
// {
//     v, err := db.Model(table).Fields("id").Where("create_time>? and create_time<?", t1, t2).Value()
//     t.AssertNil(err) // 断言查询过程无错误
//     t.Assert(v.Int(), 1) // 断言查询结果转换为整型后为1
// }
// 
// // 使用create_time字段在t1和t2时间范围内的查询条件，通过FindValue方法获取id字段值
// {
//     v, err := db.Model(table).Fields("id").Where("create_time>? and create_time<?", t1, t2).FindValue()
//     t.AssertNil(err) // 断言查询过程无错误
//     t.Assert(v.Int(), 1) // 断言查询结果转换为整型后为1
// }
// 
// // 使用create_time字段在t1和t2时间范围内的查询条件，通过FindValue方法并指定"id"字段获取值
// {
//     v, err := db.Model(table).Where("create_time>? and create_time<?", t1, t2).FindValue("id")
//     t.AssertNil(err) // 断言查询过程无错误
//     t.Assert(v.Int(), 1) // 断言查询结果转换为整型后为1
// }
// })
}

// 这是Go语言代码中的一行注释，其内容引用了GitHub上gogf/gf项目的一个问题 issue #1700。
// 翻译为：
// 参考GitHub上gogf/gf项目的第1700号问题。
func Test_Issue1700(t *testing.T) {
	table := "user_" + 时间类.X创建并按当前时间().X取文本时间戳纳秒()
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
	    CREATE TABLE %s (
	        id         int(10) unsigned NOT NULL AUTO_INCREMENT,
	        user_id    int(10) unsigned NOT NULL,
	        UserId    int(10) unsigned NOT NULL,
	        PRIMARY KEY (id)
	    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	    `, table,
	)); err != nil {
		单元测试类.AssertNil(err)
	}
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Id     int `orm:"id"`
			Userid int `orm:"user_id"`
			UserId int `orm:"UserId"`
		}
		_, err := db.X创建Model对象(table).X设置数据(User{
			Id:     1,
			Userid: 2,
			UserId: 3,
		}).X插入()
		t.AssertNil(err)

		one, err := db.X创建Model对象(table).X查询一条()
		t.AssertNil(err)
		t.Assert(one, g.Map{
			"id":      1,
			"user_id": 2,
			"UserId":  3,
		})

		for i := 0; i < 1000; i++ {
			var user *User
			err = db.X创建Model对象(table).X查询到结构体指针(&user)
			t.AssertNil(err)
			t.Assert(user.Id, 1)
			t.Assert(user.Userid, 2)
			t.Assert(user.UserId, 3)
		}
	})
}

// 这是Go语言代码中的一行注释，其内容为一个URL链接，指向GitHub上gogf/gf项目的一个问题 issue 1701。
// 中文翻译：
// 这是引用了GitHub上gogf/gf项目第1701号问题的链接。
func Test_Issue1701(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	单元测试类.C(t, func(t *单元测试类.T) {
		value, err := db.X创建Model对象(table).X字段保留过滤(db类.Raw("if(id=1,100,null)")).X条件并识别主键(1).X查询一条值()
		t.AssertNil(err)
		t.Assert(value.String(), 100)
	})
}

// 这是Go语言代码中的一行注释，其内容为一个URL链接，指向GitHub上gogf/gf项目的一个问题编号1733。
// 中文翻译：
// 这是引用了GitHub上gogf/gf项目第1733号问题的链接。
func Test_Issue1733(t *testing.T) {
	table := "user_" + uid类.X生成()
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
	    CREATE TABLE %s (
	        id int(8) unsigned zerofill NOT NULL AUTO_INCREMENT,
	        PRIMARY KEY (id)
	    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	    `, table,
	)); err != nil {
		单元测试类.AssertNil(err)
	}
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		for i := 1; i <= 10; i++ {
			_, err := db.X创建Model对象(table).X设置数据(g.Map{
				"id": i,
			}).X插入()
			t.AssertNil(err)
		}

		all, err := db.X创建Model对象(table).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(all), 10)
		for i := 0; i < 10; i++ {
			t.Assert(all[i]["id"].X取整数(), i+1)
		}
	})
}

// 这是Go语言代码中的一行注释，其内容引用了GitHub上gogf/gf项目的一个问题编号2105。
// 中文翻译：
// 参考GitHub上gogf/gf项目的问题#2105
func Test_Issue2105(t *testing.T) {
	table := "issue2105"
	array := 文本类.X分割并忽略空值(单元测试类.DataContent(`issue2105.sql`), ";")
	for _, v := range array {
		if _, err := db.X原生SQL执行(ctx, v); err != nil {
			单元测试类.Error(err)
		}
	}
	defer dropTable(table)

	type JsonItem struct {
		Name  string `json:"name,omitempty"`
		Value string `json:"value,omitempty"`
	}
	type Test struct {
		Id   string      `json:"id,omitempty"`
		Json []*JsonItem `json:"json,omitempty"`
	}

	单元测试类.C(t, func(t *单元测试类.T) {
		var list []*Test
		err := db.X创建Model对象(table).X查询到结构体指针(&list)
		t.AssertNil(err)
		t.Assert(len(list), 2)
		t.Assert(len(list[0].Json), 0)
		t.Assert(len(list[1].Json), 3)
	})
}

// 这是Go语言代码中的一行注释，其内容引用了GitHub上gogf/gf仓库的第2231号议题。
// 中文翻译：
// 参考GitHub上gogf/gf项目的问题#2231
func Test_Issue2231(t *testing.T) {
	var (
		pattern = `(\w+):([\w\-]*):(.*?)@(\w+?)\((.+?)\)/{0,1}([^\?]*)\?{0,1}(.*)`
		link    = `mysql:root:12345678@tcp(127.0.0.1:3306)/a正bc式?loc=Local&parseTime=true`
	)
	单元测试类.C(t, func(t *单元测试类.T) {
		match, err := 正则类.X匹配文本(pattern, link)
		t.AssertNil(err)
		t.Assert(match[1], "mysql")
		t.Assert(match[2], "root")
		t.Assert(match[3], "12345678")
		t.Assert(match[4], "tcp")
		t.Assert(match[5], "127.0.0.1:3306")
		t.Assert(match[6], "a正bc式")
		t.Assert(match[7], "loc=Local&parseTime=true")
	})
}

// 这是Go语言代码中的一行注释，其内容为一个URL链接，指向GitHub上gogf/gf项目的一个问题（issue）：#2339。
// 翻译成中文：
// 这是Go语言代码中的一个注释，它提供了一个链接地址：https://github.com/gogf/gf/issues/2339，该链接指向GitHub上gogf/gf项目的问题编号2339。
func Test_Issue2339(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	单元测试类.C(t, func(t *单元测试类.T) {
		model1 := db.X创建Model对象(table, "u1").X条件("id between ? and ?", 1, 9)
		model2 := db.X创建Model对象("? as u2", model1)
		model3 := db.X创建Model对象("? as u3", model2)
		all2, err := model2.X条件大于("id", 6).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(all2), 3)
		t.Assert(all2[0]["id"], 7)

		all3, err := model3.X条件大于("id", 7).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(all3), 2)
		t.Assert(all3[0]["id"], 8)
	})
}

// 这是GitHub上gogf/gf仓库的第2356号issue
func Test_Issue2356(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		table := "demo_" + uid类.X生成()
		if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
	    CREATE TABLE %s (
	        id BIGINT(20) UNSIGNED NOT NULL DEFAULT '0',
	        PRIMARY KEY (id)
	    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	    `, table,
		)); err != nil {
			t.AssertNil(err)
		}
		defer dropTable(table)

		if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`INSERT INTO %s (id) VALUES (18446744073709551615);`, table)); err != nil {
			t.AssertNil(err)
		}

		one, err := db.X创建Model对象(table).X查询一条()
		t.AssertNil(err)
		t.AssertEQ(one["id"].X取值(), uint64(18446744073709551615))
	})
}

// 这是GitHub上gogf/gf仓库的第2338个issue
func Test_Issue2338(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		table1 := "demo_" + uid类.X生成()
		table2 := "demo_" + uid类.X生成()
		if _, err := db.X切换数据库(TestSchema1).X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
    id        int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'User ID',
    nickname  varchar(45) DEFAULT NULL COMMENT 'User Nickname',
    create_at datetime(6) DEFAULT NULL COMMENT 'Created Time',
    update_at datetime(6) DEFAULT NULL COMMENT 'Updated Time',
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	    `, table1,
		)); err != nil {
			t.AssertNil(err)
		}
		if _, err := db.X切换数据库(TestSchema2).X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
    id        int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'User ID',
    nickname  varchar(45) DEFAULT NULL COMMENT 'User Nickname',
    create_at datetime(6) DEFAULT NULL COMMENT 'Created Time',
    update_at datetime(6) DEFAULT NULL COMMENT 'Updated Time',
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	    `, table2,
		)); err != nil {
			t.AssertNil(err)
		}
		defer dropTableWithDb(db.X切换数据库(TestSchema1), table1)
		defer dropTableWithDb(db.X切换数据库(TestSchema2), table2)

		var err error
		_, err = db.X切换数据库(TestSchema1).X创建Model对象(table1).X插入(g.Map{
			"id":       1,
			"nickname": "name_1",
		})
		t.AssertNil(err)

		_, err = db.X切换数据库(TestSchema2).X创建Model对象(table2).X插入(g.Map{
			"id":       1,
			"nickname": "name_2",
		})
		t.AssertNil(err)

		tableName1 := fmt.Sprintf(`%s.%s`, TestSchema1, table1)
		tableName2 := fmt.Sprintf(`%s.%s`, TestSchema2, table2)
		all, err := db.X创建Model对象(tableName1).X设置表别名(`a`).
			X左连接(tableName2+" b", `a.id=b.id`).
			X字段保留过滤(`a.id`, `b.nickname`).X查询()
		t.AssertNil(err)
		t.Assert(len(all), 1)
		t.Assert(all[0]["nickname"], "name_2")
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		table1 := "demo_" + uid类.X生成()
		table2 := "demo_" + uid类.X生成()
		if _, err := db.X切换数据库(TestSchema1).X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
    id        int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'User ID',
    nickname  varchar(45) DEFAULT NULL COMMENT 'User Nickname',
    create_at datetime(6) DEFAULT NULL COMMENT 'Created Time',
    update_at datetime(6) DEFAULT NULL COMMENT 'Updated Time',
    deleted_at datetime(6) DEFAULT NULL COMMENT 'Deleted Time',
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	    `, table1,
		)); err != nil {
			t.AssertNil(err)
		}
		if _, err := db.X切换数据库(TestSchema2).X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
    id        int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'User ID',
    nickname  varchar(45) DEFAULT NULL COMMENT 'User Nickname',
    create_at datetime(6) DEFAULT NULL COMMENT 'Created Time',
    update_at datetime(6) DEFAULT NULL COMMENT 'Updated Time',
    deleted_at datetime(6) DEFAULT NULL COMMENT 'Deleted Time',
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	    `, table2,
		)); err != nil {
			t.AssertNil(err)
		}
		defer dropTableWithDb(db.X切换数据库(TestSchema1), table1)
		defer dropTableWithDb(db.X切换数据库(TestSchema2), table2)

		var err error
		_, err = db.X切换数据库(TestSchema1).X创建Model对象(table1).X插入(g.Map{
			"id":       1,
			"nickname": "name_1",
		})
		t.AssertNil(err)

		_, err = db.X切换数据库(TestSchema2).X创建Model对象(table2).X插入(g.Map{
			"id":       1,
			"nickname": "name_2",
		})
		t.AssertNil(err)

		tableName1 := fmt.Sprintf(`%s.%s`, TestSchema1, table1)
		tableName2 := fmt.Sprintf(`%s.%s`, TestSchema2, table2)
		all, err := db.X创建Model对象(tableName1).X设置表别名(`a`).
			X左连接(tableName2+" b", `a.id=b.id`).
			X字段保留过滤(`a.id`, `b.nickname`).X查询()
		t.AssertNil(err)
		t.Assert(len(all), 1)
		t.Assert(all[0]["nickname"], "name_2")
	})
}

// 这是Go语言代码中的一行注释，其内容引用了GitHub上gogf/gf仓库下的第2427个issue（问题）。
// 翻译成中文：
// 引用了GitHub上gogf/gf项目中的第2427个问题。
func Test_Issue2427(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		table := "demo_" + uid类.X生成()
		if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
    id        int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'User ID',
    passport  varchar(45) NOT NULL COMMENT 'User Passport',
    password  varchar(45) NOT NULL COMMENT 'User Password',
    nickname  varchar(45) NOT NULL COMMENT 'User Nickname',
    create_at datetime(6) DEFAULT NULL COMMENT 'Created Time',
    update_at datetime(6) DEFAULT NULL COMMENT 'Updated Time',
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	    `, table,
		)); err != nil {
			t.AssertNil(err)
		}
		defer dropTable(table)

		_, err1 := db.X创建Model对象(table).X删除()
		t.Assert(err1, `there should be WHERE condition statement for DELETE operation`)

		_, err2 := db.X创建Model对象(table).X条件(g.Map{}).X删除()
		t.Assert(err2, `there should be WHERE condition statement for DELETE operation`)

		_, err3 := db.X创建Model对象(table).X条件(1).X删除()
		t.AssertNil(err3)
	})
}

// 这是GitHub上gogf/gf仓库的第2561个issue链接
func Test_Issue2561(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			g.Meta     `orm:"do:true"`
			Id         interface{}
			Passport   interface{}
			Password   interface{}
			Nickname   interface{}
			CreateTime interface{}
		}
		data := g.Slice别名{
			User{
				Id:       1,
				Passport: "user_1",
			},
			User{
				Id:       2,
				Password: "pass_2",
			},
			User{
				Id:       3,
				Password: "pass_3",
			},
		}
		result, err := db.X创建Model对象(table).X设置数据(data).X插入()
		t.AssertNil(err)
		m, _ := result.LastInsertId()
		t.Assert(m, 3)

		n, _ := result.RowsAffected()
		t.Assert(n, 3)

		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one[`id`], `1`)
		t.Assert(one[`passport`], `user_1`)
		t.Assert(one[`password`], ``)
		t.Assert(one[`nickname`], ``)
		t.Assert(one[`create_time`], ``)

		one, err = db.X创建Model对象(table).X条件并识别主键(2).X查询一条()
		t.AssertNil(err)
		t.Assert(one[`id`], `2`)
		t.Assert(one[`passport`], ``)
		t.Assert(one[`password`], `pass_2`)
		t.Assert(one[`nickname`], ``)
		t.Assert(one[`create_time`], ``)

		one, err = db.X创建Model对象(table).X条件并识别主键(3).X查询一条()
		t.AssertNil(err)
		t.Assert(one[`id`], `3`)
		t.Assert(one[`passport`], ``)
		t.Assert(one[`password`], `pass_3`)
		t.Assert(one[`nickname`], ``)
		t.Assert(one[`create_time`], ``)
	})
}

// 这是GitHub上gogf/gf项目的一个问题链接，编号为2439。
func Test_Issue2439(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 文本类.X分割并忽略空值(单元测试类.DataContent(`issue2439.sql`), ";")
		for _, v := range array {
			if _, err := db.X原生SQL执行(ctx, v); err != nil {
				单元测试类.Error(err)
			}
		}
		defer dropTable("a")
		defer dropTable("b")
		defer dropTable("c")

		orm := db.X创建Model对象("a")
		orm = orm.X内连接(
			"c", "a.id=c.id",
		)
		orm = orm.X内连接相同字段("b", "id")
		whereFormat := fmt.Sprintf(
			"(`%s`.`%s` LIKE ?) ",
			"b", "name",
		)
		orm = orm.X条件或格式化(
			whereFormat,
			"%a%",
		)
		r, err := orm.X查询()
		t.AssertNil(err)
		t.Assert(len(r), 1)
		t.Assert(r[0]["id"], 2)
		t.Assert(r[0]["name"], "a")
	})
}

// 这是Go语言代码中的一行注释，其内容为一个URL链接，指向GitHub上gogf/gf项目的一个issue（问题）讨论页面，编号为2782。
// 中文翻译：
// 这是Go语言代码中的一个注释，它提供了一个链接至GitHub上gogf/gf项目第2782号问题的讨论页面。
func Test_Issue2787(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		m := db.X创建Model对象("user")

		condWhere, _ := m.X创建组合条件().
			X条件("id", "").
			X条件(m.X创建组合条件().
				X条件("nickname", "foo").
				X条件或("password", "abc123")).
			X条件("passport", "pp").
			X生成条件字符串及参数()
		t.Assert(condWhere, "(`id`=?) AND (((`nickname`=?) OR (`password`=?))) AND (`passport`=?)")

		condWhere, _ = m.X过滤空值().X创建组合条件().
			X条件("id", "").
			X条件(m.X创建组合条件().
				X条件("nickname", "foo").
				X条件或("password", "abc123")).
			X条件("passport", "pp").
			X生成条件字符串及参数()
		t.Assert(condWhere, "((`nickname`=?) OR (`password`=?)) AND (`passport`=?)")

		condWhere, _ = m.X过滤空值().X创建组合条件().
			X条件(m.X创建组合条件().
				X条件("nickname", "foo").
				X条件或("password", "abc123")).
			X条件("id", "").
			X条件("passport", "pp").
			X生成条件字符串及参数()
		t.Assert(condWhere, "((`nickname`=?) OR (`password`=?)) AND (`passport`=?)")
	})
}

// 这是GitHub上gogf/gf仓库的第2907号问题链接
func Test_Issue2907(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			orm = db.X创建Model对象(table)
			err error
		)

		orm = orm.X条件不包含并带前缀(
			table,
			"id",
			[]int{
				1,
				2,
			},
		)
		all, err := orm.X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(all), TableSize-2)
		t.Assert(all[0]["id"], 3)
	})
}

// 这是GitHub上gogf/gf仓库的第3086个issue链接
func Test_Issue3086(t *testing.T) {
	table := "issue3086_user"
	array := 文本类.X分割并忽略空值(单元测试类.DataContent(`issue3086.sql`), ";")
	for _, v := range array {
		if _, err := db.X原生SQL执行(ctx, v); err != nil {
			单元测试类.Error(err)
		}
	}
	defer dropTable(table)
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			g.Meta     `orm:"do:true"`
			Id         interface{}
			Passport   interface{}
			Password   interface{}
			Nickname   interface{}
			CreateTime interface{}
		}
		data := g.Slice别名{
			User{
				Id:       nil,
				Passport: "user_1",
			},
			User{
				Id:       2,
				Passport: "user_2",
			},
		}
		_, err := db.X创建Model对象(table).X设置数据(data).X设置批量操作行数(10).X插入()
		t.AssertNE(err, nil)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			g.Meta     `orm:"do:true"`
			Id         interface{}
			Passport   interface{}
			Password   interface{}
			Nickname   interface{}
			CreateTime interface{}
		}
		data := g.Slice别名{
			User{
				Id:       1,
				Passport: "user_1",
			},
			User{
				Id:       2,
				Passport: "user_2",
			},
		}
		result, err := db.X创建Model对象(table).X设置数据(data).X设置批量操作行数(10).X插入()
		t.AssertNil(err)
		n, err := result.RowsAffected()
		t.AssertNil(err)
		t.Assert(n, 2)
	})
}

// 这是Go语言代码中的一行注释，其内容引用了GitHub上gogf/gf项目的一个问题链接：https://github.com/gogf/gf/issues/3204
// 翻译成中文：
// 这指向了GitHub上gogf/gf项目的一个问题，编号为3204
func Test_Issue3204(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	// where
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			g.Meta     `orm:"do:true"`
			Id         interface{} `orm:"id,omitempty"`
			Passport   interface{} `orm:"passport,omitempty"`
			Password   interface{} `orm:"password,omitempty"`
			Nickname   interface{} `orm:"nickname,omitempty"`
			CreateTime interface{} `orm:"create_time,omitempty"`
		}
		where := User{
			Id:       2,
			Passport: "",
		}
		all, err := db.X创建Model对象(table).X条件(where).X查询()
		t.AssertNil(err)
		t.Assert(len(all), 1)
		t.Assert(all[0]["id"], 2)
	})
	// data
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			g.Meta     `orm:"do:true"`
			Id         interface{} `orm:"id,omitempty"`
			Passport   interface{} `orm:"passport,omitempty"`
			Password   interface{} `orm:"password,omitempty"`
			Nickname   interface{} `orm:"nickname,omitempty"`
			CreateTime interface{} `orm:"create_time,omitempty"`
		}
		var (
			err      error
			sqlArray []string
			insertId int64
			data     = User{
				Id:       20,
				Passport: "passport_20",
				Password: "",
			}
		)
		sqlArray, err = db类.X捕捉SQL语句(ctx, func(ctx context.Context) error {
			insertId, err = db.X设置上下文并取副本(ctx).X创建Model对象(table).X设置数据(data).X插入并取ID()
			return err
		})
		t.AssertNil(err)
		t.Assert(insertId, 20)
		t.Assert(
			文本类.X是否包含(sqlArray[len(sqlArray)-1], "(`id`,`passport`) VALUES(20,'passport_20')"),
			true,
		)
	})
	// update data
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			g.Meta     `orm:"do:true"`
			Id         interface{} `orm:"id,omitempty"`
			Passport   interface{} `orm:"passport,omitempty"`
			Password   interface{} `orm:"password,omitempty"`
			Nickname   interface{} `orm:"nickname,omitempty"`
			CreateTime interface{} `orm:"create_time,omitempty"`
		}
		var (
			err      error
			sqlArray []string
			data     = User{
				Passport: "passport_1",
				Password: "",
				Nickname: "",
			}
		)
		sqlArray, err = db类.X捕捉SQL语句(ctx, func(ctx context.Context) error {
			_, err = db.X设置上下文并取副本(ctx).X创建Model对象(table).X设置数据(data).X条件并识别主键(1).X更新()
			return err
		})
		t.AssertNil(err)
		t.Assert(
			文本类.X是否包含(sqlArray[len(sqlArray)-1], "SET `passport`='passport_1' WHERE `id`=1"),
			true,
		)
	})
}

// 这是GitHub上gogf/gf仓库的第3218个issue
func Test_Issue3218(t *testing.T) {
	table := "issue3218_sys_config"
	array := 文本类.X分割并忽略空值(单元测试类.DataContent(`issue3218.sql`), ";")
	for _, v := range array {
		if _, err := db.X原生SQL执行(ctx, v); err != nil {
			单元测试类.Error(err)
		}
	}
	defer dropTable(table)
	单元测试类.C(t, func(t *单元测试类.T) {
		type SysConfigInfo struct {
			Name  string            `json:"name"`
			Value map[string]string `json:"value"`
		}
		var configData *SysConfigInfo
		err := db.X创建Model对象(table).X查询到结构体指针(&configData)
		t.AssertNil(err)
		t.Assert(configData, &SysConfigInfo{
			Name: "site",
			Value: map[string]string{
				"fixed_page": "",
				"site_name":  "22",
				"version":    "22",
				"banned_ip":  "22",
				"filings":    "2222",
			},
		})
	})
}
