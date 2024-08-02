// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package mysql_test

import (
	"testing"

	gjson "github.com/888go/goframe/encoding/gjson"
	"github.com/888go/goframe/frame/g"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
	gmeta "github.com/888go/goframe/util/gmeta"
)

func Test_Model_Builder(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		m := db.Model(table)
		b := m.Builder()

		all, err := m.Where(
			b.Where("id", g.Slice{1, 2, 3}).WhereOr("id", g.Slice{4, 5, 6}),
		).All()
		t.AssertNil(err)
		t.Assert(len(all), 6)
	})

	// Where And
	gtest.C(t, func(t *gtest.T) {
		m := db.Model(table)
		b := m.Builder()

		all, err := m.Where(
			b.Where("id", g.Slice{1, 2, 3}).WhereOr("id", g.Slice{4, 5, 6}),
		).Where(
			b.Where("id", g.Slice{2, 3}).WhereOr("id", g.Slice{5, 6}),
		).Where(
			b.Where("id", g.Slice{3}).Where("id", g.Slice{1, 2, 3}),
		).All()
		t.AssertNil(err)
		t.Assert(len(all), 1)
	})

	// Where Or
	gtest.C(t, func(t *gtest.T) {
		m := db.Model(table)
		b := m.Builder()

		all, err := m.WhereOr(
			b.Where("id", g.Slice{1, 2, 3}).WhereOr("id", g.Slice{4, 5, 6}),
		).WhereOr(
			b.Where("id", g.Slice{2, 3}).WhereOr("id", g.Slice{5, 6}),
		).WhereOr(
			b.Where("id", g.Slice{3}).Where("id", g.Slice{1, 2, 3}),
		).All()
		t.AssertNil(err)
		t.Assert(len(all), 6)
	})

			// 使用具有*gtime.Time类型字段的结构体的Where方法. md5:ec0563987be5b60d
	gtest.C(t, func(t *gtest.T) {
		m := db.Model(table)
		b := m.Builder()

		type Query struct {
			Id       interface{}
			Nickname *gtime.Time
		}

		where, args := b.Where(&Query{Id: 1}).Build()
		t.Assert(where, "`id`=? AND `nickname` IS NULL")
		t.Assert(args, []interface{}{1})
	})

			// 对于具有*gjson.Json类型字段的结构体. md5:2f9bb2360683fbd5
	gtest.C(t, func(t *gtest.T) {
		m := db.Model(table)
		b := m.Builder()

		type Query struct {
			Id       interface{}
			Nickname *gjson.Json
		}

		where, args := b.Where(&Query{Id: 1}).Build()
		t.Assert(where, "`id`=? AND `nickname` IS NULL")
		t.Assert(args, []interface{}{1})
	})

		// 哪里有使用了gf命令生成，且字段类型为*gtime.Time的结构体？. md5:da1a0a7e8081d4af
	gtest.C(t, func(t *gtest.T) {
		m := db.Model(table)
		b := m.Builder()

		type Query struct {
			gmeta.Meta `orm:"do:true"`
			Id         interface{}
			Nickname   *gtime.Time
		}

		where, args := b.Where(&Query{Id: 1}).Build()
		t.Assert(where, "`id`=?")
		t.Assert(args, []interface{}{1})
	})

		// 使用带有 *gjson.Json 类型字段的结构体进行操作，该结构体由 gf cli 自动生成。 md5:06bfa64a33e7e5c6
	gtest.C(t, func(t *gtest.T) {
		m := db.Model(table)
		b := m.Builder()

		type Query struct {
			gmeta.Meta `orm:"do:true"`
			Id         interface{}
			Nickname   *gjson.Json
		}

		where, args := b.Where(&Query{Id: 1}).Build()
		t.Assert(where, "`id`=?")
		t.Assert(args, []interface{}{1})
	})
}

func Test_Safe_Builder(t *testing.T) {
			// 测试m.Builder()是否在链式调用时安全. md5:429d30328a75265b
	gtest.C(t, func(t *gtest.T) {
		b := db.Model().Builder()
		b.Where("id", 1)
		_, args := b.Build()
		t.AssertNil(args)

		b = b.Where("id", 1)
		_, args = b.Build()
		t.Assert(args, g.Slice{1})
	})
}
