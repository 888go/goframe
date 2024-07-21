	// 版权归GoFrame作者(https:	//goframe.org)所有。保留所有权利。
	//
	// 本源代码形式受MIT许可证条款约束。
	// 如果未随本文件一同分发MIT许可证副本，
	// 您可以在https:	//github.com/gogf/gf处获取。
	// md5:a9832f33b234e3f3

package gmeta_test

import (
	"testing"

	"github.com/gogf/gf/v2/internal/json"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gmeta"
)

func TestMeta_Basic(t *testing.T) {
	type A struct {
		gmeta.Meta `tag:"123" orm:"456"`
		Id         int
		Name       string
	}

	gtest.C(t, func(t *gtest.T) {
		a := &A{
			Id:   100,
			Name: "john",
		}
		t.Assert(len(gmeta.Data(a)), 2)
		t.AssertEQ(gmeta.Get(a, "tag").String(), "123")
		t.AssertEQ(gmeta.Get(a, "orm").String(), "456")
		t.AssertEQ(gmeta.Get(a, "none"), nil)

		b, err := json.Marshal(a)
		t.AssertNil(err)
		t.Assert(b, `{"Id":100,"Name":"john"}`)
	})
}

func TestMeta_Convert_Map(t *testing.T) {
	type A struct {
		gmeta.Meta `tag:"123" orm:"456"`
		Id         int
		Name       string
	}

	gtest.C(t, func(t *gtest.T) {
		a := &A{
			Id:   100,
			Name: "john",
		}
		m := gconv.Map(a)
		t.Assert(len(m), 2)
		t.Assert(m[`Meta`], nil)
	})
}

func TestMeta_Json(t *testing.T) {
	type A struct {
		gmeta.Meta `tag:"123" orm:"456"`
		Id         int
	}

	gtest.C(t, func(t *gtest.T) {
		a := &A{
			Id: 100,
		}
		b, err := json.Marshal(a)
		t.AssertNil(err)
		t.Assert(string(b), `{"Id":100}`)
	})
}
