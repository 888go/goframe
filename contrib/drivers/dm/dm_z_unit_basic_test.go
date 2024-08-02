//go:build 屏蔽单元测试

// 版权所有 2019 gf 作者（https://github.com/gogf/gf）。保留所有权利。
//
// 此源代码形式受麻省理工学院（MIT）许可证的条款约束。
// 如果未随此文件一起分发MIT许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
// md5:47e609239e0cb2bc

package dm_test

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/888go/goframe/frame/g"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
)

func Test_DB_Ping(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		err1 := dblink.PingMaster()
		err2 := dblink.PingSlave()
		t.Assert(err1, nil)
		t.Assert(err2, nil)
	})
}

func TestTables(t *testing.T) {
	tables := []string{"A_tables", "A_tables2"}
	for _, v := range tables {
		createInitTable(v)
	}
	gtest.C(t, func(t *gtest.T) {
		result, err := db.Tables(ctx)
		gtest.Assert(err, nil)

		for i := 0; i < len(tables); i++ {
			find := false
			for j := 0; j < len(result); j++ {
				if strings.ToUpper(tables[i]) == result[j] {
					find = true
					break
				}
			}
			gtest.AssertEQ(find, true)
		}

		result, err = dblink.Tables(ctx)
		gtest.Assert(err, nil)
		for i := 0; i < len(tables); i++ {
			find := false
			for j := 0; j < len(result); j++ {
				if strings.ToUpper(tables[i]) == result[j] {
					find = true
					break
				}
			}
			gtest.AssertEQ(find, true)
		}
	})
}

// 该测试用例的测试场景索引（完全匹配字段）在达梦数据库中是一个保留关键字，不能作为字段名存在。
// 如果之前从mysql迁移过来的数据结构中有一个索引（完全匹配字段），也会被允许。
// 但是，在处理索引（完全匹配字段）时，适配器会自动添加安全字符。
// 原则上，如果你直接使用达梦数据库初始化，而不是从mysql迁移数据结构，就不会出现此类问题。
// 即使如此，适配器也已经考虑到了这种情况。
// md5:bdfc451ff291c639
func TestTablesFalse(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		tables := []string{"A_tables", "A_tables2"}
		for _, v := range tables {
			_, err := createTableFalse(v)
			gtest.Assert(err, fmt.Errorf("createTableFalse"))
			// createTable(v)
		}
	})
}

func TestTableFields(t *testing.T) {
	tables := "A_tables"
	createInitTable(tables)
	gtest.C(t, func(t *gtest.T) {
		var expect = map[string][]interface{}{
			"ID":           {"BIGINT", false},
			"ACCOUNT_NAME": {"VARCHAR", false},
			"PWD_RESET":    {"TINYINT", false},
			"ATTR_INDEX":   {"INT", true},
			"DELETED":      {"INT", false},
			"CREATED_TIME": {"TIMESTAMP", false},
		}

		_, err := dbErr.TableFields(ctx, "Fields")
		gtest.AssertNE(err, nil)

		res, err := db.TableFields(ctx, tables)
		gtest.Assert(err, nil)

		for k, v := range expect {
			_, ok := res[k]
			gtest.AssertEQ(ok, true)

			gtest.AssertEQ(res[k].Name, k)
			gtest.Assert(res[k].Type, v[0])
			gtest.Assert(res[k].Null, v[1])
		}

	})

	gtest.C(t, func(t *gtest.T) {
		_, err := db.TableFields(ctx, "t_user t_user2")
		gtest.AssertNE(err, nil)
	})
}

func Test_DB_Query(t *testing.T) {
	tableName := "A_tables"
	createInitTable(tableName)
	gtest.C(t, func(t *gtest.T) {
		// createTable(tableName)
		_, err := db.Query(ctx, fmt.Sprintf("SELECT * from %s", tableName))
		t.AssertNil(err)

		resTwo := make([]User, 0)
		err = db.Model(tableName).Scan(&resTwo)
		t.AssertNil(err)

		resThree := make([]User, 0)
		model := db.Model(tableName)
		model.Where("id", g.Slice{1, 2, 3, 4})
				// 使用model的Where方法，传入一个SQL条件："account_name"字段包含"%list%"字符串。 md5:008bda1d9a70b4f2
		model.Where("deleted", 0).Order("pwd_reset desc")
		_, err = model.Count()
		t.AssertNil(err)
		err = model.Page(2, 2).Scan(&resThree)
		t.AssertNil(err)
	})
}

func TestModelSave(t *testing.T) {
	table := createTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id          int
			AccountName string
			AttrIndex   int
		}
		var (
			user   User
			count  int
			result sql.Result
			err    error
		)
		db.SetDebug(true)

		result, err = db.Model(table).Data(g.Map{
			"id":          1,
			"accountName": "ac1",
			"attrIndex":   100,
		}).OnConflict("id").Save()

		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)

		err = db.Model(table).Scan(&user)
		t.AssertNil(err)
		t.Assert(user.Id, 1)
		t.Assert(user.AccountName, "ac1")
		t.Assert(user.AttrIndex, 100)

		_, err = db.Model(table).Data(g.Map{
			"id":          1,
			"accountName": "ac2",
			"attrIndex":   200,
		}).OnConflict("id").Save()
		t.AssertNil(err)

		err = db.Model(table).Scan(&user)
		t.AssertNil(err)
		t.Assert(user.AccountName, "ac2")
		t.Assert(user.AttrIndex, 200)

		count, err = db.Model(table).Count()
		t.AssertNil(err)
		t.Assert(count, 1)
	})
}

func TestModelInsert(t *testing.T) {
		// g.Model.insert 不会丢失默认非空列. md5:475cfebfed134e1b
	table := "A_tables"
	createInitTable(table)
	gtest.C(t, func(t *gtest.T) {
		i := 200
		data := User{
			ID:          int64(i),
			AccountName: fmt.Sprintf(`A%dtwo`, i),
			PwdReset:    0,
			AttrIndex:   99,
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
		}
				// 使用TestDBName数据库的模式，根据table模型和data数据执行插入操作，返回一个表示是否成功的空值和错误信息。 md5:665c442bb4e1be49
		_, err := db.Model(table).Insert(&data)
		gtest.Assert(err, nil)
	})

	gtest.C(t, func(t *gtest.T) {
		i := 201
		data := User{
			ID:          int64(i),
			AccountName: fmt.Sprintf(`A%dtwoONE`, i),
			PwdReset:    1,
			CreatedTime: time.Now(),
			AttrIndex:   98,
			UpdatedTime: time.Now(),
		}
				// 使用TestDBName数据库的模式，根据table模型和data数据执行插入操作，返回一个表示是否成功的空值和错误信息。 md5:665c442bb4e1be49
		_, err := db.Model(table).Data(&data).Insert()
		gtest.Assert(err, nil)
	})
}

func TestDBInsert(t *testing.T) {
	table := "A_tables"
	createInitTable("A_tables")
	gtest.C(t, func(t *gtest.T) {
		i := 300
		data := g.Map{
			"ID":           i,
			"ACCOUNT_NAME": fmt.Sprintf(`A%dthress`, i),
			"PWD_RESET":    3,
			"ATTR_INDEX":   98,
			"CREATED_TIME": gtime.Now(),
			"UPDATED_TIME": gtime.Now(),
		}
		_, err := db.Insert(ctx, table, &data)
		gtest.Assert(err, nil)
	})
}

func Test_DB_Exec(t *testing.T) {
	createInitTable("A_tables")
	gtest.C(t, func(t *gtest.T) {
		_, err := db.Exec(ctx, "SELECT ? from dual", 1)
		t.AssertNil(err)

		_, err = db.Exec(ctx, "ERROR")
		t.AssertNE(err, nil)
	})
}

func Test_DB_Insert(t *testing.T) {
	table := "A_tables"
	createInitTable(table)
	gtest.C(t, func(t *gtest.T) {
		timeNow := time.Now()
		// normal map
		_, err := db.Insert(ctx, table, g.Map{
			"ID":           1000,
			"ACCOUNT_NAME": "map1",
			"CREATED_TIME": timeNow,
			"UPDATED_TIME": timeNow,
		})
		t.AssertNil(err)

		result, err := db.Insert(ctx, table, g.Map{
			"ID":           "2000",
			"ACCOUNT_NAME": "map2",
			"CREATED_TIME": timeNow,
			"UPDATED_TIME": timeNow,
		})
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)

		result, err = db.Insert(ctx, table, g.Map{
			"ID":           3000,
			"ACCOUNT_NAME": "map3",
			"CREATED_TIME": timeNow,
			"UPDATED_TIME": timeNow,
		})
		t.AssertNil(err)
		n, _ = result.RowsAffected()
		t.Assert(n, 1)

		// struct
		result, err = db.Insert(ctx, table, User{
			ID:          4000,
			AccountName: "struct_4",
			CreatedTime: timeNow,
			UpdatedTime: timeNow,
		})
		t.AssertNil(err)
		n, _ = result.RowsAffected()
		t.Assert(n, 1)

		ones, err := db.Model(table).Where("ID", 4000).All()
		t.AssertNil(err)
		t.Assert(ones[0]["ID"].Int(), 4000)
		t.Assert(ones[0]["ACCOUNT_NAME"].String(), "struct_4")
		// 待办事项：问题2
		// 这是DM（可能是某个项目或模块的缩写）的bug。
		// 断言one["CREATED_TIME"]的GTime转换后字符串与timeStr相等。
		// md5:6c078020ce38de99

		// *struct
		result, err = db.Insert(ctx, table, &User{
			ID:          5000,
			AccountName: "struct_5",
			CreatedTime: timeNow,
			UpdatedTime: timeNow,
		})
		t.AssertNil(err)
		n, _ = result.RowsAffected()
		t.Assert(n, 1)

		one, err := db.Model(table).Where("ID", 5000).One()
		t.AssertNil(err)
		t.Assert(one["ID"].Int(), 5000)
		t.Assert(one["ACCOUNT_NAME"].String(), "struct_5")

		// batch with Insert
		r, err := db.Insert(ctx, table, g.Slice{
			g.Map{
				"ID":           6000,
				"ACCOUNT_NAME": "t6000",
				"CREATED_TIME": timeNow,
				"UPDATED_TIME": timeNow,
			},
			g.Map{
				"ID":           6001,
				"ACCOUNT_NAME": "t6001",
				"CREATED_TIME": timeNow,
				"UPDATED_TIME": timeNow,
			},
		})
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 2)

		one, err = db.Model(table).Where("ID", 6000).One()
		t.AssertNil(err)
		t.Assert(one["ID"].Int(), 6000)
		t.Assert(one["ACCOUNT_NAME"].String(), "t6000")
	})
}

func Test_DB_BatchInsert(t *testing.T) {
	table := "A_tables"
	createInitTable(table)
	gtest.C(t, func(t *gtest.T) {
		r, err := db.Insert(ctx, table, g.List{
			{
				"ID":           400,
				"ACCOUNT_NAME": "list_400",
				"CREATE_TIME":  gtime.Now(),
				"UPDATED_TIME": gtime.Now(),
			},
			{
				"ID":           401,
				"ACCOUNT_NAME": "list_401",
				"CREATE_TIME":  gtime.Now(),
				"UPDATED_TIME": gtime.Now(),
			},
		}, 1)
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 2)
	})

	gtest.C(t, func(t *gtest.T) {
		// []interface{}
		r, err := db.Insert(ctx, table, g.Slice{
			g.Map{
				"ID":           500,
				"ACCOUNT_NAME": "500_batch_500",
				"CREATE_TIME":  gtime.Now(),
				"UPDATED_TIME": gtime.Now(),
			},
			g.Map{
				"ID":           501,
				"ACCOUNT_NAME": "501_batch_501",
				"CREATE_TIME":  gtime.Now(),
				"UPDATED_TIME": gtime.Now(),
			},
		}, 1)
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 2)
	})

	// batch insert map
	gtest.C(t, func(t *gtest.T) {
		result, err := db.Insert(ctx, table, g.Map{
			"ID":           600,
			"ACCOUNT_NAME": "600_batch_600",
			"CREATE_TIME":  gtime.Now(),
			"UPDATED_TIME": gtime.Now(),
		})
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
	})
}

func Test_DB_BatchInsert_Struct(t *testing.T) {
	// batch insert struct
	table := "A_tables"
	createInitTable(table)
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		user := &User{
			ID:          700,
			AccountName: "BatchInsert_Struct_700",
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
		}
		result, err := db.Model(table).Insert(user)
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
	})
}

func Test_DB_Update(t *testing.T) {
	table := "A_tables"
	createInitTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.Update(ctx, table, "pwd_reset=7", "id=7")
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)

		one, err := db.Model(table).Where("ID", 7).One()
		t.AssertNil(err)
		t.Assert(one["ID"].Int(), 7)
		t.Assert(one["ACCOUNT_NAME"].String(), "name_7")
		t.Assert(one["PWD_RESET"].String(), "7")
	})
}

func Test_DB_GetAll(t *testing.T) {
	table := "A_tables"
	createInitTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.GetAll(ctx, fmt.Sprintf("SELECT * FROM %s WHERE id=?", table), 1)
		t.AssertNil(err)
		t.Assert(len(result), 1)
		t.Assert(result[0]["ID"].Int(), 1)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.GetAll(ctx, fmt.Sprintf("SELECT * FROM %s WHERE id=?", table), g.Slice{1})
		t.AssertNil(err)
		t.Assert(len(result), 1)
		t.Assert(result[0]["ID"].Int(), 1)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.GetAll(ctx, fmt.Sprintf("SELECT * FROM %s WHERE id in(?)", table), g.Slice{1, 2, 3})
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["ID"].Int(), 1)
		t.Assert(result[1]["ID"].Int(), 2)
		t.Assert(result[2]["ID"].Int(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.GetAll(ctx, fmt.Sprintf("SELECT * FROM %s WHERE id in(?,?,?)", table), g.Slice{1, 2, 3})
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["ID"].Int(), 1)
		t.Assert(result[1]["ID"].Int(), 2)
		t.Assert(result[2]["ID"].Int(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.GetAll(ctx, fmt.Sprintf("SELECT * FROM %s WHERE id in(?,?,?)", table), g.Slice{1, 2, 3}...)
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["ID"].Int(), 1)
		t.Assert(result[1]["ID"].Int(), 2)
		t.Assert(result[2]["ID"].Int(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.GetAll(ctx, fmt.Sprintf("SELECT * FROM %s WHERE id>=? AND id <=?", table), g.Slice{1, 3})
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["ID"].Int(), 1)
		t.Assert(result[1]["ID"].Int(), 2)
		t.Assert(result[2]["ID"].Int(), 3)
	})
}

func Test_DB_GetOne(t *testing.T) {
	table := "A_tables"
	createInitTable(table)
	gtest.C(t, func(t *gtest.T) {
		record, err := db.GetOne(ctx, fmt.Sprintf("SELECT * FROM %s WHERE account_name=?", table), "name_4")
		t.AssertNil(err)
		t.Assert(record["ACCOUNT_NAME"].String(), "name_4")
	})
}

func Test_DB_GetValue(t *testing.T) {
	table := "A_tables"
	createInitTable(table)
	gtest.C(t, func(t *gtest.T) {
		value, err := db.GetValue(ctx, fmt.Sprintf("SELECT id FROM %s WHERE account_name=?", table), "name_2")
		t.AssertNil(err)
		t.Assert(value.Int(), 2)
	})
}

func Test_DB_GetCount(t *testing.T) {
	table := "A_tables"
	createInitTable(table)
	gtest.C(t, func(t *gtest.T) {
		count, err := db.GetCount(ctx, fmt.Sprintf("SELECT * FROM %s", table))
		t.AssertNil(err)
		t.Assert(count, 10)
	})
}

func Test_DB_GetStruct(t *testing.T) {
	table := "A_tables"
	createInitTable(table)
	gtest.C(t, func(t *gtest.T) {
		user := new(User)
		err := db.GetScan(ctx, user, fmt.Sprintf("SELECT * FROM %s WHERE id=?", table), 3)
		t.AssertNil(err)
		t.Assert(user.AccountName, "name_3")
	})
	gtest.C(t, func(t *gtest.T) {
		user := new(User)
		err := db.GetScan(ctx, user, fmt.Sprintf("SELECT * FROM %s WHERE id=?", table), 2)
		t.AssertNil(err)
		t.Assert(user.AccountName, "name_2")
	})
}

func Test_DB_GetStructs(t *testing.T) {
	table := "A_tables"
	createInitTable(table)
	gtest.C(t, func(t *gtest.T) {
		var users []User
		err := db.GetScan(ctx, &users, fmt.Sprintf("SELECT * FROM %s WHERE id>?", table), 4)
		t.AssertNil(err)
		t.Assert(users[0].ID, 5)
		t.Assert(users[1].ID, 6)
		t.Assert(users[2].ID, 7)
		t.Assert(users[0].AccountName, "name_5")
		t.Assert(users[1].AccountName, "name_6")
		t.Assert(users[2].AccountName, "name_7")
	})
}

func Test_DB_GetScan(t *testing.T) {
	table := "A_tables"
	createInitTable(table)
	gtest.C(t, func(t *gtest.T) {
		user := new(User)
		err := db.GetScan(ctx, user, fmt.Sprintf("SELECT * FROM %s WHERE id=?", table), 3)
		t.AssertNil(err)
		t.Assert(user.AccountName, "name_3")
	})
	gtest.C(t, func(t *gtest.T) {
		var user *User
		err := db.GetScan(ctx, &user, fmt.Sprintf("SELECT * FROM %s WHERE id=?", table), 3)
		t.AssertNil(err)
		t.Assert(user.AccountName, "name_3")
	})
	gtest.C(t, func(t *gtest.T) {
		var users []User
		err := db.GetScan(ctx, &users, fmt.Sprintf("SELECT * FROM %s WHERE id<?", table), 4)
		t.AssertNil(err)
		t.Assert(users[0].ID, 1)
		t.Assert(users[1].ID, 2)
		t.Assert(users[2].ID, 3)
		t.Assert(users[0].AccountName, "name_1")
		t.Assert(users[1].AccountName, "name_2")
		t.Assert(users[2].AccountName, "name_3")
	})
}

func Test_DB_Delete(t *testing.T) {
	table := "A_tables"
	createInitTable(table)
	gtest.C(t, func(t *gtest.T) {
		result, err := db.Delete(ctx, table, "id=32")
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 0)
	})

	gtest.C(t, func(t *gtest.T) {
		result, err := db.Model(table).Where("id", 33).Delete()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 0)
	})
}

func Test_Empty_Slice_Argument(t *testing.T) {
	table := "A_tables"
	createInitTable(table)
	gtest.C(t, func(t *gtest.T) {
		result, err := db.GetAll(ctx, fmt.Sprintf(`select * from %s where id in(?)`, table), g.Slice{})
		t.AssertNil(err)
		t.Assert(len(result), 0)
	})
}
