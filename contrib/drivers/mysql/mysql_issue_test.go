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

	gtest.C(t, func(t *gtest.T) {
		one, err := db.Model(table).Where(" id ", 1).One()
		t.AssertNil(err)
		t.Assert(one["id"], 1)
	})
}

// 这是Go语言代码中的一行注释，其内容引用了GitHub上gogf/gf仓库下的第1570号issue。
// 翻译成中文：
// 引用了GitHub上gogf/gf项目中的第1570号问题。
func Test_Issue1570(t *testing.T) {
	var (
		tableUser       = "user_" + gtime.TimestampMicroStr()
		tableUserDetail = "user_detail_" + gtime.TimestampMicroStr()
		tableUserScores = "user_scores_" + gtime.TimestampMicroStr()
	)
	if _, err := db.Exec(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUser)); err != nil {
		gtest.Error(err)
	}
	defer dropTable(tableUser)

	if _, err := db.Exec(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  address varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUserDetail)); err != nil {
		gtest.Error(err)
	}
	defer dropTable(tableUserDetail)

	if _, err := db.Exec(ctx, fmt.Sprintf(`
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

	// 初始化数据
	gtest.C(t, func(t *gtest.T) {
		var err error
		for i := 1; i <= 5; i++ {
			// User.
			_, err = db.Insert(ctx, tableUser, g.Map{
				"uid":  i,
				"name": fmt.Sprintf(`name_%d`, i),
			})
			t.AssertNil(err)
			// Detail.
			_, err = db.Insert(ctx, tableUserDetail, g.Map{
				"uid":     i,
				"address": fmt.Sprintf(`address_%d`, i),
			})
			t.AssertNil(err)
			// Scores.
			for j := 1; j <= 5; j++ {
				_, err = db.Insert(ctx, tableUserScores, g.Map{
					"uid":   i,
					"score": j,
				})
				t.AssertNil(err)
			}
		}
	})

	// Result ScanList，用于包含结构体元素和指针属性。
	gtest.C(t, func(t *gtest.T) {
		var users []Entity
		// User
		err := db.Model(tableUser).
			Where("uid", g.Slice{3, 4}).
			Fields("uid").
			Order("uid asc").
			ScanList(&users, "User")
		t.AssertNil(err)
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].User, &EntityUser{3, ""})
		t.Assert(users[1].User, &EntityUser{4, ""})
		// Detail
		err = db.Model(tableUserDetail).
			Where("uid", gdb.ListItemValues(users, "User", "Uid")).
			Order("uid asc").
			ScanList(&users, "UserDetail", "User", "uid:Uid")
		t.AssertNil(err)
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, &EntityUserDetail{3, "address_3"})
		t.Assert(users[1].UserDetail, &EntityUserDetail{4, "address_4"})
		// Scores
		err = db.Model(tableUserScores).
			Where("uid", gdb.ListItemValues(users, "User", "Uid")).
			Order("id asc").
			ScanList(&users, "UserScores", "User", "uid:Uid")
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
	array := gstr.SplitAndTrim(gtest.DataContent(`issue1401.sql`), ";")
	for _, v := range array {
		if _, err := db.Exec(ctx, v); err != nil {
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
		err := db.Model(table1).With(parcelDetail.Items).Where("id", 3).Scan(&parcelDetail)
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
	array := gstr.SplitAndTrim(gtest.DataContent(`issue1412.sql`), ";")
	for _, v := range array {
		if _, err := db.Exec(ctx, v); err != nil {
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
		err := db.Model("parcels").With(Items{}).Where("id=3").Scan(&entity)
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
		err := db.Model("parcels").With(Items{}).Where("id=30000").Scan(&entity)
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

	result, err := db.Model(table).Data(g.Map{
		"id":          1,
		"passport":    "port_1",
		"password":    "pass_1",
		"nickname":    "name_2",
		"create_time": "2020-10-27 19:03:33",
	}).Insert()
	gtest.AssertNil(err)
	n, _ := result.RowsAffected()
	gtest.Assert(n, 1)

	// where + string.
	gtest.C(t, func(t *gtest.T) {
		v, err := db.Model(table).Fields("id").Where("create_time>'2020-10-27 19:03:32' and create_time<'2020-10-27 19:03:34'").Value()
		t.AssertNil(err)
		t.Assert(v.Int(), 1)
	})
	gtest.C(t, func(t *gtest.T) {
		v, err := db.Model(table).Fields("id").Where("create_time>'2020-10-27 19:03:32' and create_time<'2020-10-27 19:03:34'").Value()
		t.AssertNil(err)
		t.Assert(v.Int(), 1)
	})
	// where + 字符串参数。
	gtest.C(t, func(t *gtest.T) {
		v, err := db.Model(table).Fields("id").Where("create_time>? and create_time<?", "2020-10-27 19:03:32", "2020-10-27 19:03:34").Value()
		t.AssertNil(err)
		t.Assert(v.Int(), 1)
	})
	// where + gtime.Time 参数
	gtest.C(t, func(t *gtest.T) {
		v, err := db.Model(table).Fields("id").Where("create_time>? and create_time<?", gtime.New("2020-10-27 19:03:32"), gtime.New("2020-10-27 19:03:34")).Value()
		t.AssertNil(err)
		t.Assert(v.Int(), 1)
	})
	// 函数参数中包含 where + time.Time 类型，时间使用 UTC（协调世界时）。
	gtest.C(t, func(t *gtest.T) {
		t1, _ := time.Parse("2006-01-02 15:04:05", "2020-10-27 11:03:32")
		t2, _ := time.Parse("2006-01-02 15:04:05", "2020-10-27 11:03:34")
		{
			v, err := db.Model(table).Fields("id").Where("create_time>? and create_time<?", t1, t2).Value()
			t.AssertNil(err)
			t.Assert(v.Int(), 1)
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
	table := "user_" + gtime.Now().TimestampNanoStr()
	if _, err := db.Exec(ctx, fmt.Sprintf(`
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
		_, err := db.Model(table).Data(User{
			Id:     1,
			Userid: 2,
			UserId: 3,
		}).Insert()
		t.AssertNil(err)

		one, err := db.Model(table).One()
		t.AssertNil(err)
		t.Assert(one, g.Map{
			"id":      1,
			"user_id": 2,
			"UserId":  3,
		})

		for i := 0; i < 1000; i++ {
			var user *User
			err = db.Model(table).Scan(&user)
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
	gtest.C(t, func(t *gtest.T) {
		value, err := db.Model(table).Fields(gdb.Raw("if(id=1,100,null)")).WherePri(1).Value()
		t.AssertNil(err)
		t.Assert(value.String(), 100)
	})
}

// 这是Go语言代码中的一行注释，其内容为一个URL链接，指向GitHub上gogf/gf项目的一个问题编号1733。
// 中文翻译：
// 这是引用了GitHub上gogf/gf项目第1733号问题的链接。
func Test_Issue1733(t *testing.T) {
	table := "user_" + guid.S()
	if _, err := db.Exec(ctx, fmt.Sprintf(`
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
			_, err := db.Model(table).Data(g.Map{
				"id": i,
			}).Insert()
			t.AssertNil(err)
		}

		all, err := db.Model(table).OrderAsc("id").All()
		t.AssertNil(err)
		t.Assert(len(all), 10)
		for i := 0; i < 10; i++ {
			t.Assert(all[i]["id"].Int(), i+1)
		}
	})
}

// 这是Go语言代码中的一行注释，其内容引用了GitHub上gogf/gf项目的一个问题编号2105。
// 中文翻译：
// 参考GitHub上gogf/gf项目的问题#2105
func Test_Issue2105(t *testing.T) {
	table := "issue2105"
	array := gstr.SplitAndTrim(gtest.DataContent(`issue2105.sql`), ";")
	for _, v := range array {
		if _, err := db.Exec(ctx, v); err != nil {
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
		err := db.Model(table).Scan(&list)
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
	gtest.C(t, func(t *gtest.T) {
		match, err := gregex.MatchString(pattern, link)
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
	gtest.C(t, func(t *gtest.T) {
		model1 := db.Model(table, "u1").Where("id between ? and ?", 1, 9)
		model2 := db.Model("? as u2", model1)
		model3 := db.Model("? as u3", model2)
		all2, err := model2.WhereGT("id", 6).OrderAsc("id").All()
		t.AssertNil(err)
		t.Assert(len(all2), 3)
		t.Assert(all2[0]["id"], 7)

		all3, err := model3.WhereGT("id", 7).OrderAsc("id").All()
		t.AssertNil(err)
		t.Assert(len(all3), 2)
		t.Assert(all3[0]["id"], 8)
	})
}

// 这是GitHub上gogf/gf仓库的第2356号issue
func Test_Issue2356(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		table := "demo_" + guid.S()
		if _, err := db.Exec(ctx, fmt.Sprintf(`
	    CREATE TABLE %s (
	        id BIGINT(20) UNSIGNED NOT NULL DEFAULT '0',
	        PRIMARY KEY (id)
	    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	    `, table,
		)); err != nil {
			t.AssertNil(err)
		}
		defer dropTable(table)

		if _, err := db.Exec(ctx, fmt.Sprintf(`INSERT INTO %s (id) VALUES (18446744073709551615);`, table)); err != nil {
			t.AssertNil(err)
		}

		one, err := db.Model(table).One()
		t.AssertNil(err)
		t.AssertEQ(one["id"].Val(), uint64(18446744073709551615))
	})
}

// 这是GitHub上gogf/gf仓库的第2338个issue
func Test_Issue2338(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		table1 := "demo_" + guid.S()
		table2 := "demo_" + guid.S()
		if _, err := db.Schema(TestSchema1).Exec(ctx, fmt.Sprintf(`
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
		if _, err := db.Schema(TestSchema2).Exec(ctx, fmt.Sprintf(`
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
		defer dropTableWithDb(db.Schema(TestSchema1), table1)
		defer dropTableWithDb(db.Schema(TestSchema2), table2)

		var err error
		_, err = db.Schema(TestSchema1).Model(table1).Insert(g.Map{
			"id":       1,
			"nickname": "name_1",
		})
		t.AssertNil(err)

		_, err = db.Schema(TestSchema2).Model(table2).Insert(g.Map{
			"id":       1,
			"nickname": "name_2",
		})
		t.AssertNil(err)

		tableName1 := fmt.Sprintf(`%s.%s`, TestSchema1, table1)
		tableName2 := fmt.Sprintf(`%s.%s`, TestSchema2, table2)
		all, err := db.Model(tableName1).As(`a`).
			LeftJoin(tableName2+" b", `a.id=b.id`).
			Fields(`a.id`, `b.nickname`).All()
		t.AssertNil(err)
		t.Assert(len(all), 1)
		t.Assert(all[0]["nickname"], "name_2")
	})

	gtest.C(t, func(t *gtest.T) {
		table1 := "demo_" + guid.S()
		table2 := "demo_" + guid.S()
		if _, err := db.Schema(TestSchema1).Exec(ctx, fmt.Sprintf(`
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
		if _, err := db.Schema(TestSchema2).Exec(ctx, fmt.Sprintf(`
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
		defer dropTableWithDb(db.Schema(TestSchema1), table1)
		defer dropTableWithDb(db.Schema(TestSchema2), table2)

		var err error
		_, err = db.Schema(TestSchema1).Model(table1).Insert(g.Map{
			"id":       1,
			"nickname": "name_1",
		})
		t.AssertNil(err)

		_, err = db.Schema(TestSchema2).Model(table2).Insert(g.Map{
			"id":       1,
			"nickname": "name_2",
		})
		t.AssertNil(err)

		tableName1 := fmt.Sprintf(`%s.%s`, TestSchema1, table1)
		tableName2 := fmt.Sprintf(`%s.%s`, TestSchema2, table2)
		all, err := db.Model(tableName1).As(`a`).
			LeftJoin(tableName2+" b", `a.id=b.id`).
			Fields(`a.id`, `b.nickname`).All()
		t.AssertNil(err)
		t.Assert(len(all), 1)
		t.Assert(all[0]["nickname"], "name_2")
	})
}

// 这是Go语言代码中的一行注释，其内容引用了GitHub上gogf/gf仓库下的第2427个issue（问题）。
// 翻译成中文：
// 引用了GitHub上gogf/gf项目中的第2427个问题。
func Test_Issue2427(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		table := "demo_" + guid.S()
		if _, err := db.Exec(ctx, fmt.Sprintf(`
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

		_, err1 := db.Model(table).Delete()
		t.Assert(err1, `there should be WHERE condition statement for DELETE operation`)

		_, err2 := db.Model(table).Where(g.Map{}).Delete()
		t.Assert(err2, `there should be WHERE condition statement for DELETE operation`)

		_, err3 := db.Model(table).Where(1).Delete()
		t.AssertNil(err3)
	})
}

// 这是GitHub上gogf/gf仓库的第2561个issue链接
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
		data := g.Slice{
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
		result, err := db.Model(table).Data(data).Insert()
		t.AssertNil(err)
		m, _ := result.LastInsertId()
		t.Assert(m, 3)

		n, _ := result.RowsAffected()
		t.Assert(n, 3)

		one, err := db.Model(table).WherePri(1).One()
		t.AssertNil(err)
		t.Assert(one[`id`], `1`)
		t.Assert(one[`passport`], `user_1`)
		t.Assert(one[`password`], ``)
		t.Assert(one[`nickname`], ``)
		t.Assert(one[`create_time`], ``)

		one, err = db.Model(table).WherePri(2).One()
		t.AssertNil(err)
		t.Assert(one[`id`], `2`)
		t.Assert(one[`passport`], ``)
		t.Assert(one[`password`], `pass_2`)
		t.Assert(one[`nickname`], ``)
		t.Assert(one[`create_time`], ``)

		one, err = db.Model(table).WherePri(3).One()
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
	gtest.C(t, func(t *gtest.T) {
		array := gstr.SplitAndTrim(gtest.DataContent(`issue2439.sql`), ";")
		for _, v := range array {
			if _, err := db.Exec(ctx, v); err != nil {
				gtest.Error(err)
			}
		}
		defer dropTable("a")
		defer dropTable("b")
		defer dropTable("c")

		orm := db.Model("a")
		orm = orm.InnerJoin(
			"c", "a.id=c.id",
		)
		orm = orm.InnerJoinOnField("b", "id")
		whereFormat := fmt.Sprintf(
			"(`%s`.`%s` LIKE ?) ",
			"b", "name",
		)
		orm = orm.WhereOrf(
			whereFormat,
			"%a%",
		)
		r, err := orm.All()
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

	gtest.C(t, func(t *gtest.T) {
		m := db.Model("user")

		condWhere, _ := m.Builder().
			Where("id", "").
			Where(m.Builder().
				Where("nickname", "foo").
				WhereOr("password", "abc123")).
			Where("passport", "pp").
			Build()
		t.Assert(condWhere, "(`id`=?) AND (((`nickname`=?) OR (`password`=?))) AND (`passport`=?)")

		condWhere, _ = m.OmitEmpty().Builder().
			Where("id", "").
			Where(m.Builder().
				Where("nickname", "foo").
				WhereOr("password", "abc123")).
			Where("passport", "pp").
			Build()
		t.Assert(condWhere, "((`nickname`=?) OR (`password`=?)) AND (`passport`=?)")

		condWhere, _ = m.OmitEmpty().Builder().
			Where(m.Builder().
				Where("nickname", "foo").
				WhereOr("password", "abc123")).
			Where("id", "").
			Where("passport", "pp").
			Build()
		t.Assert(condWhere, "((`nickname`=?) OR (`password`=?)) AND (`passport`=?)")
	})
}

// 这是GitHub上gogf/gf仓库的第2907号问题链接
func Test_Issue2907(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		var (
			orm = db.Model(table)
			err error
		)

		orm = orm.WherePrefixNotIn(
			table,
			"id",
			[]int{
				1,
				2,
			},
		)
		all, err := orm.OrderAsc("id").All()
		t.AssertNil(err)
		t.Assert(len(all), TableSize-2)
		t.Assert(all[0]["id"], 3)
	})
}

// 这是GitHub上gogf/gf仓库的第3086个issue链接
func Test_Issue3086(t *testing.T) {
	table := "issue3086_user"
	array := gstr.SplitAndTrim(gtest.DataContent(`issue3086.sql`), ";")
	for _, v := range array {
		if _, err := db.Exec(ctx, v); err != nil {
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
		data := g.Slice{
			User{
				Id:       nil,
				Passport: "user_1",
			},
			User{
				Id:       2,
				Passport: "user_2",
			},
		}
		_, err := db.Model(table).Data(data).Batch(10).Insert()
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
		data := g.Slice{
			User{
				Id:       1,
				Passport: "user_1",
			},
			User{
				Id:       2,
				Passport: "user_2",
			},
		}
		result, err := db.Model(table).Data(data).Batch(10).Insert()
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
		all, err := db.Model(table).Where(where).All()
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
		sqlArray, err = gdb.CatchSQL(ctx, func(ctx context.Context) error {
			insertId, err = db.Ctx(ctx).Model(table).Data(data).InsertAndGetId()
			return err
		})
		t.AssertNil(err)
		t.Assert(insertId, 20)
		t.Assert(
			gstr.Contains(sqlArray[len(sqlArray)-1], "(`id`,`passport`) VALUES(20,'passport_20')"),
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
		sqlArray, err = gdb.CatchSQL(ctx, func(ctx context.Context) error {
			_, err = db.Ctx(ctx).Model(table).Data(data).WherePri(1).Update()
			return err
		})
		t.AssertNil(err)
		t.Assert(
			gstr.Contains(sqlArray[len(sqlArray)-1], "SET `passport`='passport_1' WHERE `id`=1"),
			true,
		)
	})
}

// 这是GitHub上gogf/gf仓库的第3218个issue
func Test_Issue3218(t *testing.T) {
	table := "issue3218_sys_config"
	array := gstr.SplitAndTrim(gtest.DataContent(`issue3218.sql`), ";")
	for _, v := range array {
		if _, err := db.Exec(ctx, v); err != nil {
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
		err := db.Model(table).Scan(&configData)
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
