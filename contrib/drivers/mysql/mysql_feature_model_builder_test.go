// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package mysql_test

import (
	"testing"
	
	"github.com/888go/goframe/encoding/gjson"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gmeta"
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

	// Where 针对具有 *gtime.Time 类型字段的结构体
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

	// Where 函数，用于处理具有字段类型为 *gjson.Json 的结构体
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

	// Where 函数配合使用具有 *gtime.Time 类型字段的结构体，该结构体由 gf cli 生成
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

	// Where 针对由 gf cli 生成的拥有 *gjson.Json 类型字段的 do 结构体
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
	// 测试m.Builder()是否支持链式安全调用
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
