// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package session类_test

import (
	"context"
	"testing"
	"time"

	"github.com/888go/goframe/frame/g"
	gsession "github.com/888go/goframe/os/gsession"
	gtest "github.com/888go/goframe/test/gtest"
)

func Test_StorageMemory(t *testing.T) {
	storage := gsession.NewStorageMemory()
	manager := gsession.New(time.Second, storage)
	sessionId := ""
	gtest.C(t, func(t *gtest.T) {
		s := manager.New(context.TODO())
		defer s.Close()
		s.X设置值("k1", "v1")
		s.X设置值("k2", "v2")
		s.SetMap(g.Map{
			"k3": "v3",
			"k4": "v4",
		})
		t.Assert(s.IsDirty(), true)
		sessionId = s.MustId()
	})

	time.Sleep(500 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		s := manager.New(context.TODO(), sessionId)
		t.Assert(s.MustGet("k1"), "v1")
		t.Assert(s.MustGet("k2"), "v2")
		t.Assert(s.MustGet("k3"), "v3")
		t.Assert(s.MustGet("k4"), "v4")
		t.Assert(len(s.MustData()), 4)
		t.Assert(s.MustData()["k1"], "v1")
		t.Assert(s.MustData()["k4"], "v4")
		t.Assert(s.MustId(), sessionId)
		t.Assert(s.MustSize(), 4)
		t.Assert(s.MustContains("k1"), true)
		t.Assert(s.MustContains("k3"), true)
		t.Assert(s.MustContains("k5"), false)
		s.Remove("k4")
		t.Assert(s.MustSize(), 3)
		t.Assert(s.MustContains("k3"), true)
		t.Assert(s.MustContains("k4"), false)
		s.RemoveAll()
		t.Assert(s.MustSize(), 0)
		t.Assert(s.MustContains("k1"), false)
		t.Assert(s.MustContains("k2"), false)
		s.SetMap(g.Map{
			"k5": "v5",
			"k6": "v6",
		})
		t.Assert(s.MustSize(), 2)
		t.Assert(s.MustContains("k5"), true)
		t.Assert(s.MustContains("k6"), true)
	})

	time.Sleep(1000 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		s := manager.New(context.TODO(), sessionId)
		t.Assert(s.MustSize(), 0)
		t.Assert(s.MustGet("k5"), nil)
		t.Assert(s.MustGet("k6"), nil)
	})
}
