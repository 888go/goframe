//go:build 屏蔽单元测试

// 版权所有 2019 gf 作者（https://github.com/gogf/gf）。保留所有权利。
//
// 此源代码形式受麻省理工学院（MIT）许可证的条款约束。
// 如果未随此文件一起分发MIT许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
// md5:47e609239e0cb2bc

package dm_test

import (
	"testing"
	"time"

	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
	gstr "github.com/888go/goframe/text/gstr"
)

func Test_Issue2594(t *testing.T) {
	table := "HANDLE_INFO"
	array := gstr.SplitAndTrim(gtest.DataContent(`issue`, `2594`, `sql.sql`), ";")
	for _, v := range array {
		if _, err := db.Exec(ctx, v); err != nil {
			gtest.Error(err)
		}
	}
	defer dropTable(table)

	type HandleValueMysql struct {
		Index int64  `orm:"index"`
		Type  string `orm:"type"`
		Data  []byte `orm:"data"`
	}
	type HandleInfoMysql struct {
		Id         int                `orm:"id,primary" json:"id"`
		SubPrefix  string             `orm:"sub_prefix"`
		Prefix     string             `orm:"prefix"`
		HandleName string             `orm:"handle_name"`
		CreateTime time.Time          `orm:"create_time"`
		UpdateTime time.Time          `orm:"update_time"`
		Value      []HandleValueMysql `orm:"value"`
	}

	gtest.C(t, func(t *gtest.T) {
		var h1 = HandleInfoMysql{
			SubPrefix:  "p_",
			Prefix:     "m_",
			HandleName: "name",
			CreateTime: gtime.Now().FormatTo("Y-m-d H:i:s").Time,
			UpdateTime: gtime.Now().FormatTo("Y-m-d H:i:s").Time,
			Value: []HandleValueMysql{
				{
					Index: 10,
					Type:  "t1",
					Data:  []byte("abc"),
				},
				{
					Index: 20,
					Type:  "t2",
					Data:  []byte("def"),
				},
			},
		}
		_, err := db.Model(table).OmitEmptyData().Insert(h1)
		t.AssertNil(err)

		var h2 HandleInfoMysql
		err = db.Model(table).Scan(&h2)
		t.AssertNil(err)

		h1.Id = 1
		t.Assert(h1, h2)
	})
}
