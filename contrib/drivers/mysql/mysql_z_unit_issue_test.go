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
	"sync"
	"testing"
	"time"

	gdb "github.com/888go/goframe/database/gdb"
	"github.com/888go/goframe/frame/g"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
	gregex "github.com/888go/goframe/text/gregex"
	gstr "github.com/888go/goframe/text/gstr"
	gmeta "github.com/888go/goframe/util/gmeta"
	guid "github.com/888go/goframe/util/guid"
)

//github.com/gogf/gf/issues/1934. md5:96f55929c7ed56a0
func Test_Issue1934(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		one, err := db.X创建Model对象(table).X条件(" id ", 1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["id"], 1)
	})
}

//github.com/gogf/gf/issues/1570. md5:37966850af641bcc
func Test_Issue1570(t *testing.T) {
	var (
		tableUser       = "user_" + gtime.X取文本时间戳微秒()
		tableUserDetail = "user_detail_" + gtime.X取文本时间戳微秒()
		tableUserScores = "user_scores_" + gtime.X取文本时间戳微秒()
	)
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUser)); err != nil {
		gtest.Error(err)
	}
	defer dropTable(tableUser)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  address varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUserDetail)); err != nil {
		gtest.Error(err)
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
		gtest.Error(err)
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

	// Initialize the data.
	gtest.C(t, func(t *gtest.T) {
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

		// Result 使用具有结构体元素和指针属性的ScanList。 md5:b23d106d13859ad5
	gtest.C(t, func(t *gtest.T) {
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
			X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).
			X排序("uid asc").
			X查询到指针列表(&users, "UserDetail", "User", "uid:Uid")
		t.AssertNil(err)
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, &EntityUserDetail{3, "address_3"})
		t.Assert(users[1].UserDetail, &EntityUserDetail{4, "address_4"})
		// Scores
		err = db.X创建Model对象(tableUserScores).
			X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).
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

// 这段注释引用的是一个GitHub问题（Issue）的链接，来自gogf（一个Go语言的框架）项目。它表示这是对问题1401的讨论或参考。在中文中，这可能表示：“参见GitHub上的gogf项目问题1401”。 md5:5d32589f093beb22
func Test_Issue1401(t *testing.T) {
	var (
		table1 = "parcels"
		table2 = "parcel_items"
	)
	array := gstr.X分割并忽略空值(gtest.DataContent(`issue1401.sql`), ";")
	for _, v := range array {
		if _, err := db.X原生SQL执行(ctx, v); err != nil {
			gtest.Error(err)
		}
	}
	defer dropTable(table1)
	defer dropTable(table2)

	gtest.C(t, func(t *gtest.T) {
		type NItem struct {
			Id       int `json:"id"`
			ParcelId int `json:"parcel_id"`
		}

		type ParcelItem struct {
			gmeta.Meta `orm:"table:parcel_items"`
			NItem
		}

		type ParcelRsp struct {
			gmeta.Meta `orm:"table:parcels"`
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

// 这段注释引用的是GitHub上的一个issue（问题或讨论），gf（GoGF）是一个用Go语言编写的Web框架。 issue号1412可能是指该框架中某个特定的问题或者提出的改进请求，具体内容需要查看相关issue的详细描述。 md5:c6f20cc497b1e9a6
func Test_Issue1412(t *testing.T) {
	var (
		table1 = "parcels"
		table2 = "items"
	)
	array := gstr.X分割并忽略空值(gtest.DataContent(`issue1412.sql`), ";")
	for _, v := range array {
		if _, err := db.X原生SQL执行(ctx, v); err != nil {
			gtest.Error(err)
		}
	}
	defer dropTable(table1)
	defer dropTable(table2)

	gtest.C(t, func(t *gtest.T) {
		type Items struct {
			gmeta.Meta `orm:"table:items"`
			Id         int    `json:"id"`
			Name       string `json:"name"`
		}

		type ParcelRsp struct {
			gmeta.Meta `orm:"table:parcels"`
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

	gtest.C(t, func(t *gtest.T) {
		type Items struct {
			gmeta.Meta `orm:"table:items"`
			Id         int    `json:"id"`
			Name       string `json:"name"`
		}

		type ParcelRsp struct {
			gmeta.Meta `orm:"table:parcels"`
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

//github.com/gogf/gf/issues/1002 问题讨论. md5:2a97dfd9cd049763
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
		// 带有时间.Time参数，使用UTC时区。 md5:80f36eaa256e894c
	gtest.C(t, func(t *gtest.T) {
		t1, _ := time.Parse("2006-01-02 15:04:05", "2020-10-27 11:03:32")
		t2, _ := time.Parse("2006-01-02 15:04:05", "2020-10-27 11:03:34")
		{
			v, err := db.X创建Model对象(table).X字段保留过滤("id").X条件("create_time>? and create_time<?", t1, t2).X查询一条值()
			t.AssertNil(err)
			t.Assert(v.X取整数(), 1)
		}
	})
	// 在时间.Time参数中，+8代表时区偏移。
	// gtest.C(t, func(t *gtest.T) {
	//     	// 将当前时区更改为+8时区（东八区）。
	//     location, err := time.LoadLocation("Asia/Shanghai")
	//     t.AssertNil(err) 	// 确认加载时区无错误。
	//     t1, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-10-27 19:03:32", location)
	//     t2, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-10-27 19:03:34", location)
	//     	// 使用定义的时间段进行查询测试：
	//     {
	//         v, err := db.Model(table).Fields("id").Where("create_time>? and create_time<?", t1, t2).Value()
	//         t.AssertNil(err) 	// 确认查询无错误。
	//         t.Assert(v.Int(), 1) 	// 断言查询结果的ID为1。
	//     }
	//     {
	//         v, err := db.Model(table).Fields("id").Where("create_time>? and create_time<?", t1, t2).FindValue()
	//         t.AssertNil(err) 	// 同上，确认查询无错误。
	//         t.Assert(v.Int(), 1) 	// 断言查询结果的ID为1。
	//     }
	//     {
	//         v, err := db.Model(table).Where("create_time>? and create_time<?", t1, t2).FindValue("id")
	//         t.AssertNil(err) 	// 再次确认查询无错误。
	//         t.Assert(v.Int(), 1) 	// 继续断言查询结果的ID为1。
	//     }
	// })
	// md5:766797023d98820e
}

// 这段注释是引用了GitHub上gf框架的一个问题链接，编号为1700。在Go代码中，这种注释通常用于指向相关的讨论、问题或者需求，以便其他开发者了解代码的背景或上下文。翻译成中文后，其含义不变：. md5:a352b9ef5236ff28
func Test_Issue1700(t *testing.T) {
	table := "user_" + gtime.X创建并按当前时间().X取文本时间戳纳秒()
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
	    CREATE TABLE %s (
	        id         int(10) unsigned NOT NULL AUTO_INCREMENT,
	        user_id    int(10) unsigned NOT NULL,
	        UserId    int(10) unsigned NOT NULL,
	        PRIMARY KEY (id)
	    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	    `, table,
	)); err != nil {
		gtest.AssertNil(err)
	}
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
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

//github.com/gogf/gf/issues/1701。GF（Go Foundation）可能是Go语言的一个项目或者库，而"1701"可能是问题的编号。这个注释可能是在讨论或记录与GF项目相关的问题1701的情况。 md5:cc9c86ac60eeaf58
func Test_Issue1701(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		value, err := db.X创建Model对象(table).X字段保留过滤(gdb.Raw("if(id=1,100,null)")).X条件并识别主键(1).X查询一条值()
		t.AssertNil(err)
		t.Assert(value.String(), 100)
	})
}

// 这段注释链接指向的是GitHub上的一个Issue（问题报告）页面，来自gogf/gf（一个Go语言的框架）项目。具体来说，它可能是指1733号问题或者与该问题相关的内容。在中文中，这通常表示对某个问题、讨论或改进的引用。 md5:76faec7f21ba3b13
func Test_Issue1733(t *testing.T) {
	table := "user_" + guid.X生成()
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
	    CREATE TABLE %s (
	        id int(8) unsigned zerofill NOT NULL AUTO_INCREMENT,
	        PRIMARY KEY (id)
	    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	    `, table,
	)); err != nil {
		gtest.AssertNil(err)
	}
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
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

//github.com/gogf/gf/issues/2105. md5:579ab324e61be1fb
func Test_Issue2105(t *testing.T) {
	table := "issue2105"
	array := gstr.X分割并忽略空值(gtest.DataContent(`issue2105.sql`), ";")
	for _, v := range array {
		if _, err := db.X原生SQL执行(ctx, v); err != nil {
			gtest.Error(err)
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

	gtest.C(t, func(t *gtest.T) {
		var list []*Test
		err := db.X创建Model对象(table).X查询到结构体指针(&list)
		t.AssertNil(err)
		t.Assert(len(list), 2)
		t.Assert(len(list[0].Json), 0)
		t.Assert(len(list[1].Json), 3)
	})
}

// 这段注释是链接到一个GitHub问题的引用，该问题是关于gf（GoFrame）框架的一个问题或讨论。在GitHub仓库gf的 issues 页面中，编号2231的问题提供了更多的上下文和信息。由于注释本身没有详细内容，所以具体的翻译就是保持原样，表示这是一个与gf框架相关的问题链接。 md5:803083b8650008ce
func Test_Issue2231(t *testing.T) {
	var (
		pattern = `(\w+):([\w\-]*):(.*?)@(\w+?)\((.+?)\)/{0,1}([^\?]*)\?{0,1}(.*)`
		link    = `mysql:root:12345678@tcp(127.0.0.1:3306)/a正bc式?loc=Local&parseTime=true`
	)
	gtest.C(t, func(t *gtest.T) {
		match, err := gregex.X匹配文本(pattern, link)
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

//github.com/gogf/gf/issues/2339。"gf"可能是某个Go语言的库（golang的gopher框架）的简称，"issues/2339"表示该仓库中编号为2339的问题或者issue。这可能是一个开发者社区中关于gf库的报告、提问或者反馈。 md5:fb506ddf20da598c
func Test_Issue2339(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
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

//github.com/gogf/gf/issues/2356。"gf"可能是某个项目的代号，"gogf"可能是一个开发者的用户名，"issues/2356"表示该问题是编号为2356的 issue（通常是开发者社区中报告的问题、建议或讨论）。 md5:a688eda9a4ec7d89
func Test_Issue2356(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		table := "demo_" + guid.X生成()
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

//github.com/gogf/gf/issues/2338. md5:a504f30db0e1a70a
func Test_Issue2338(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		table1 := "demo_" + guid.X生成()
		table2 := "demo_" + guid.X生成()
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

	gtest.C(t, func(t *gtest.T) {
		table1 := "demo_" + guid.X生成()
		table2 := "demo_" + guid.X生成()
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

// 这段注释是引用了GitHub上一个名为gf的项目中的问题编号2427。在Go语言中，这种注释通常用于指向相关问题、讨论或需求的链接，以便其他开发人员了解代码背景或跟踪问题。翻译成中文后，它依然保持原样，因为这是一个网址引用，并无实际需要翻译的内容。 md5:cf1b689a44aec285
func Test_Issue2427(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		table := "demo_" + guid.X生成()
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

//github.com/gogf/gf/issues/2561。gf（Golang Foundation）是一个Go语言的开源框架，而"issues/2561"表示该仓库中编号为2561的问题或讨论。可能是用户在报告问题、请求功能或者讨论某个特定的代码问题。 md5:97cd71d9bf45e151
func Test_Issue2561(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
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

// 这段注释引用的是一个GitHub问题（issues）的链接，来自 "gf"（Go Foundation）项目，编号为2439。它可能是一个关于gf库的问题报告、讨论或者是一个已知问题的链接。具体的内容需要查看该链接才能得知。 md5:e37e02e670c04910
func Test_Issue2439(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := gstr.X分割并忽略空值(gtest.DataContent(`issue2439.sql`), ";")
		for _, v := range array {
			if _, err := db.X原生SQL执行(ctx, v); err != nil {
				gtest.Error(err)
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

//github.com/gogf/gf/issues/2782. md5:e2d84654d9404496
func Test_Issue2787(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
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

//github.com/gogf/gf/issues/2907. md5:61d8552a7d7948bb
func Test_Issue2907(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
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

// 这段注释引用的是GitHub上的一个 issue，gf（Go Foundation）是一个用Go语言编写的开源框架。3086号 issue 可能是关于gf框架的一个已知问题、错误报告、功能请求或者讨论点。具体的内容需要查看该issue的详细描述。 md5:629eedddf9f2ae76
func Test_Issue3086(t *testing.T) {
	table := "issue3086_user"
	array := gstr.X分割并忽略空值(gtest.DataContent(`issue3086.sql`), ";")
	for _, v := range array {
		if _, err := db.X原生SQL执行(ctx, v); err != nil {
			gtest.Error(err)
		}
	}
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
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
	gtest.C(t, func(t *gtest.T) {
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

// 这段注释引用的是GitHub上的一个 issue，gf（Go Foundation）是一个用Go语言编写的开源框架。"3204"可能是指issue的编号，表示这个注释是在讨论或参考该框架中的第3204个问题或请求。具体的内容需要查看issue页面以获取详细信息。 md5:36c0adae03298bd3
func Test_Issue3204(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	// where
	gtest.C(t, func(t *gtest.T) {
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
	gtest.C(t, func(t *gtest.T) {
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
		sqlArray, err = gdb.X捕捉SQL语句(ctx, func(ctx context.Context) error {
			insertId, err = db.X设置上下文并取副本(ctx).X创建Model对象(table).X设置数据(data).X插入并取ID()
			return err
		})
		t.AssertNil(err)
		t.Assert(insertId, 20)
		t.Assert(
			gstr.X是否包含(sqlArray[len(sqlArray)-1], "(`id`,`passport`) VALUES(20,'passport_20')"),
			true,
		)
	})
	// update data
	gtest.C(t, func(t *gtest.T) {
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
		sqlArray, err = gdb.X捕捉SQL语句(ctx, func(ctx context.Context) error {
			_, err = db.X设置上下文并取副本(ctx).X创建Model对象(table).X设置数据(data).X条件并识别主键(1).X更新()
			return err
		})
		t.AssertNil(err)
		t.Assert(
			gstr.X是否包含(sqlArray[len(sqlArray)-1], "SET `passport`='passport_1' WHERE `id`=1"),
			true,
		)
	})
}

//github.com/gogf/gf/issues/3218. md5:ebeb6327a156dd70
func Test_Issue3218(t *testing.T) {
	table := "issue3218_sys_config"
	array := gstr.X分割并忽略空值(gtest.DataContent(`issue3218.sql`), ";")
	for _, v := range array {
		if _, err := db.X原生SQL执行(ctx, v); err != nil {
			gtest.Error(err)
		}
	}
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
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

// 这段注释是指向GitHub上一个名为gf的项目的一个问题链接，问题编号为2552。在Go代码中，这种注释通常用于引用外部资源，如问题、讨论或文档，以便其他开发者了解代码的相关背景或上下文。 md5:23870b69cce8c4de
func Test_Issue2552_ClearTableFieldsAll(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	showTableKey := `SHOW FULL COLUMNS FROM`
	gtest.C(t, func(t *gtest.T) {
		ctx := context.Background()
		sqlArray, err := gdb.X捕捉SQL语句(ctx, func(ctx context.Context) error {
			_, err := db.X创建Model对象(table).X设置上下文并取副本(ctx).X插入(g.Map{
				"passport":    guid.X生成(),
				"password":    guid.X生成(),
				"nickname":    guid.X生成(),
				"create_time": gtime.X创建并从文本(CreateTime).String(),
			})
			return err
		})
		t.AssertNil(err)
		t.Assert(gstr.X是否包含(gstr.X连接(sqlArray, "|"), showTableKey), true)

		ctx = context.Background()
		sqlArray, err = gdb.X捕捉SQL语句(ctx, func(ctx context.Context) error {
			one, err := db.X创建Model对象(table).X设置上下文并取副本(ctx).X查询一条()
			t.Assert(len(one), 5)
			return err
		})
		t.AssertNil(err)
		t.Assert(gstr.X是否包含(gstr.X连接(sqlArray, "|"), showTableKey), false)

		_, err = db.X原生SQL执行(ctx, fmt.Sprintf("alter table %s drop column `nickname`", table))
		t.AssertNil(err)

		err = db.X取Core对象().X删除表字段所有缓存(ctx)
		t.AssertNil(err)

		ctx = context.Background()
		sqlArray, err = gdb.X捕捉SQL语句(ctx, func(ctx context.Context) error {
			one, err := db.X创建Model对象(table).X设置上下文并取副本(ctx).X查询一条()
			t.Assert(len(one), 4)
			return err
		})
		t.AssertNil(err)
		t.Assert(gstr.X是否包含(gstr.X连接(sqlArray, "|"), showTableKey), true)
	})
}

// 这段注释是指向GitHub上一个名为gf的项目的一个问题链接，问题编号为2552。在Go代码中，这种注释通常用于引用外部资源，如问题、讨论或文档，以便其他开发者了解代码的相关背景或上下文。 md5:23870b69cce8c4de
func Test_Issue2552_ClearTableFields(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	showTableKey := `SHOW FULL COLUMNS FROM`
	gtest.C(t, func(t *gtest.T) {
		ctx := context.Background()
		sqlArray, err := gdb.X捕捉SQL语句(ctx, func(ctx context.Context) error {
			_, err := db.X创建Model对象(table).X设置上下文并取副本(ctx).X插入(g.Map{
				"passport":    guid.X生成(),
				"password":    guid.X生成(),
				"nickname":    guid.X生成(),
				"create_time": gtime.X创建并从文本(CreateTime).String(),
			})
			return err
		})
		t.AssertNil(err)
		t.Assert(gstr.X是否包含(gstr.X连接(sqlArray, "|"), showTableKey), true)

		ctx = context.Background()
		sqlArray, err = gdb.X捕捉SQL语句(ctx, func(ctx context.Context) error {
			one, err := db.X创建Model对象(table).X设置上下文并取副本(ctx).X查询一条()
			t.Assert(len(one), 5)
			return err
		})
		t.AssertNil(err)
		t.Assert(gstr.X是否包含(gstr.X连接(sqlArray, "|"), showTableKey), false)

		_, err = db.X原生SQL执行(ctx, fmt.Sprintf("alter table %s drop column `nickname`", table))
		t.AssertNil(err)

		err = db.X取Core对象().X删除表字段缓存(ctx, table)
		t.AssertNil(err)

		ctx = context.Background()
		sqlArray, err = gdb.X捕捉SQL语句(ctx, func(ctx context.Context) error {
			one, err := db.X创建Model对象(table).X设置上下文并取副本(ctx).X查询一条()
			t.Assert(len(one), 4)
			return err
		})
		t.AssertNil(err)
		t.Assert(gstr.X是否包含(gstr.X连接(sqlArray, "|"), showTableKey), true)
	})
}

// 这段注释引用的是一个GitHub问题（issues）的链接，来自 "gf"（Go Foundation）项目，编号为2643。它可能是一个关于gf库的问题报告、讨论或者是一个已知问题的链接。具体的内容需要查看该链接才能了解。 md5:e98064ecba25be28
func Test_Issue2643(t *testing.T) {
	table := "issue2643"
	array := gstr.X分割并忽略空值(gtest.DataContent(`issue2643.sql`), ";")
	for _, v := range array {
		if _, err := db.X原生SQL执行(ctx, v); err != nil {
			gtest.Error(err)
		}
	}
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		var (
			expectKey1 = "SELECT s.name,replace(concat_ws(',',lpad(s.id, 6, '0'),s.name),',','') `code` FROM `issue2643` AS s"
			expectKey2 = "SELECT CASE WHEN dept='物资部' THEN '物资部' ELSE '其他' END dept,sum(s.value) FROM `issue2643` AS s GROUP BY CASE WHEN dept='物资部' THEN '物资部' ELSE '其他' END"
		)
		sqlArray, err := gdb.X捕捉SQL语句(ctx, func(ctx context.Context) error {
			db.X设置上下文并取副本(ctx).X创建Model对象(table).X设置表别名("s").X字段保留过滤(
				"s.name",
				"replace(concat_ws(',',lpad(s.id, 6, '0'),s.name),',','') `code`",
			).X查询()
			db.X设置上下文并取副本(ctx).X创建Model对象(table).X设置表别名("s").X字段保留过滤(
				"CASE WHEN dept='物资部' THEN '物资部' ELSE '其他' END dept",
				"sum(s.value)",
			).X排序分组("CASE WHEN dept='物资部' THEN '物资部' ELSE '其他' END").X查询()
			return nil
		})
		t.AssertNil(err)
		sqlContent := gstr.X连接(sqlArray, "\n")
		t.Assert(gstr.X是否包含(sqlContent, expectKey1), true)
		t.Assert(gstr.X是否包含(sqlContent, expectKey2), true)
	})
}

//github.com/gogf/gf/issues/3238。gf（GoGF）是一个用Go语言编写的高性能Web框架。这个注释可能是开发者在提到他们在gf项目中遇到的问题或提出的一个改进请求，3238号issue可能是一个已知问题的编号或者一个讨论的话题。 md5:98233bbbba37f999
func Test_Issue3238(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			_, err := db.X设置上下文并取副本(ctx).X创建Model对象(table).Hook(gdb.HookHandler{
				Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
					result, err = in.Next(ctx)
					if err != nil {
						return
					}
					var wg sync.WaitGroup
					for _, record := range result {
						wg.Add(1)
						go func(record gdb.Record) {
							defer wg.Done()
							id, _ := db.X设置上下文并取副本(ctx).X创建Model对象(table).X条件并识别主键(1).X查询一条值(`id`)
							nickname, _ := db.X设置上下文并取副本(ctx).X创建Model对象(table).X条件并识别主键(1).X查询一条值(`nickname`)
							t.Assert(id.X取整数(), 1)
							t.Assert(nickname.String(), "name_1")
						}(record)
					}
					wg.Wait()
					return
				},
			},
			).X查询()
			t.AssertNil(err)
		}
	})
}
