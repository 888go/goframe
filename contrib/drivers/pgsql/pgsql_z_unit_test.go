//go:build 屏蔽单元测试

// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package pgsql_test

import (
	"context"
	"testing"

	gdb "github.com/888go/goframe/database/gdb"
	"github.com/888go/goframe/frame/g"
	gctx "github.com/888go/goframe/os/gctx"
	gtest "github.com/888go/goframe/test/gtest"

	"github.com/888go/goframe/contrib/drivers/pgsql"
)

func Test_LastInsertId(t *testing.T) {
	// err not nil
	gtest.C(t, func(t *gtest.T) {
		_, err := db.X创建Model对象("notexist").X插入(g.Map切片{
			{"name": "user1"},
			{"name": "user2"},
			{"name": "user3"},
		})
		t.AssertNE(err, nil)
	})

	gtest.C(t, func(t *gtest.T) {
		tableName := createTable()
		defer dropTable(tableName)
		res, err := db.X创建Model对象(tableName).X插入(g.Map切片{
			{"passport": "user1", "password": "pwd", "nickname": "nickname", "create_time": CreateTime},
			{"passport": "user2", "password": "pwd", "nickname": "nickname", "create_time": CreateTime},
			{"passport": "user3", "password": "pwd", "nickname": "nickname", "create_time": CreateTime},
		})
		t.AssertNil(err)
		lastInsertId, err := res.LastInsertId()
		t.AssertNil(err)
		t.Assert(lastInsertId, int64(3))
		rowsAffected, err := res.RowsAffected()
		t.AssertNil(err)
		t.Assert(rowsAffected, int64(3))
	})
}

func Test_TxLastInsertId(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		tableName := createTable()
		defer dropTable(tableName)
		err := db.X事务(context.TODO(), func(ctx context.Context, tx gdb.TX) error {
			// user
			res, err := tx.X创建Model对象(tableName).X插入(g.Map切片{
				{"passport": "user1", "password": "pwd", "nickname": "nickname", "create_time": CreateTime},
				{"passport": "user2", "password": "pwd", "nickname": "nickname", "create_time": CreateTime},
				{"passport": "user3", "password": "pwd", "nickname": "nickname", "create_time": CreateTime},
			})
			t.AssertNil(err)
			lastInsertId, err := res.LastInsertId()
			t.AssertNil(err)
			t.AssertEQ(lastInsertId, int64(3))
			rowsAffected, err := res.RowsAffected()
			t.AssertNil(err)
			t.AssertEQ(rowsAffected, int64(3))

			res1, err := tx.X创建Model对象(tableName).X插入(g.Map切片{
				{"passport": "user4", "password": "pwd", "nickname": "nickname", "create_time": CreateTime},
				{"passport": "user5", "password": "pwd", "nickname": "nickname", "create_time": CreateTime},
			})
			t.AssertNil(err)
			lastInsertId1, err := res1.LastInsertId()
			t.AssertNil(err)
			t.AssertEQ(lastInsertId1, int64(5))
			rowsAffected1, err := res1.RowsAffected()
			t.AssertNil(err)
			t.AssertEQ(rowsAffected1, int64(2))
			return nil

		})
		t.AssertNil(err)
	})
}

func Test_Driver_DoFilter(t *testing.T) {
	var (
		ctx    = gctx.X创建()
		driver = pgsql.Driver{}
	)
	gtest.C(t, func(t *gtest.T) {
		var data = g.Map{
			`select * from user where (role)::jsonb ?| 'admin'`: `select * from user where (role)::jsonb ?| 'admin'`,
			`select * from user where (role)::jsonb ?| '?'`:     `select * from user where (role)::jsonb ?| '$2'`,
			`select * from user where (role)::jsonb &? '?'`:     `select * from user where (role)::jsonb &? '$2'`,
			`select * from user where (role)::jsonb ? '?'`:      `select * from user where (role)::jsonb ? '$2'`,
			`select * from user where '?'`:                      `select * from user where '$1'`,
		}
		for k, v := range data {
			newSql, _, err := driver.X底层DoFilter(ctx, nil, k, nil)
			t.AssertNil(err)
			t.Assert(newSql, v)
		}
	})
}
