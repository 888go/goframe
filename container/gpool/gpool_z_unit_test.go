// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 对象复用类_test

import (
	"errors"
	"testing"
	"time"

	gpool "github.com/888go/goframe/container/gpool"
	"github.com/888go/goframe/frame/g"
	gtest "github.com/888go/goframe/test/gtest"
)

var nf gpool.NewFunc = func() (i interface{}, e error) {
	return "hello", nil
}

var assertIndex int = 0

var ef gpool.ExpireFunc = func(i interface{}) {
	assertIndex++
	gtest.Assert(i, assertIndex)
}

func Test_Gpool(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		//
		// expire = 0
		p1 := gpool.X创建(0, nf)
		p1.X入栈(1)
		p1.X入栈(2)
		time.Sleep(1 * time.Second)
		// test won't be timeout
		v1, err1 := p1.X出栈()
		t.Assert(err1, nil)
		t.AssertIN(v1, g.Slice别名{1, 2})
		// test clear
		p1.X清空()
		t.Assert(p1.X取数量(), 0)
		// test newFunc
		v1, err1 = p1.X出栈()
		t.Assert(err1, nil)
		t.Assert(v1, "hello")
		// put data again
		p1.X入栈(3)
		p1.X入栈(4)
		v1, err1 = p1.X出栈()
		t.Assert(err1, nil)
		t.AssertIN(v1, g.Slice别名{3, 4})
		// test close
		p1.X关闭()
		v1, err1 = p1.X出栈()
		t.Assert(err1, nil)
		t.Assert(v1, "hello")
	})

	gtest.C(t, func(t *gtest.T) {
		//
		// expire > 0
		p2 := gpool.X创建(2*time.Second, nil, ef)
		for index := 0; index < 10; index++ {
			p2.X入栈(index)
		}
		t.Assert(p2.X取数量(), 10)
		v2, err2 := p2.X出栈()
		t.Assert(err2, nil)
		t.Assert(v2, 0)
						// 测试超时过期函数. md5:f6b3719e6bec20e3
		time.Sleep(3 * time.Second)
		v2, err2 = p2.X出栈()
		t.Assert(err2, errors.New("pool is empty"))
		t.Assert(v2, nil)
		// test close expireFunc
		for index := 0; index < 10; index++ {
			p2.X入栈(index)
		}
		t.Assert(p2.X取数量(), 10)
		v2, err2 = p2.X出栈()
		t.Assert(err2, nil)
		t.Assert(v2, 0)
		assertIndex = 0
		p2.X关闭()
		time.Sleep(3 * time.Second)
		t.AssertNE(p2.X入栈(1), nil)
	})

	gtest.C(t, func(t *gtest.T) {
		//
		// expire < 0
		p3 := gpool.X创建(-1, nil)
		v3, err3 := p3.X出栈()
		t.Assert(err3, errors.New("pool is empty"))
		t.Assert(v3, nil)
	})

	gtest.C(t, func(t *gtest.T) {
		p := gpool.X创建(time.Millisecond*200, nil, func(i interface{}) {})
		p.X入栈(1)
		time.Sleep(time.Millisecond * 100)
		p.X入栈(2)
		time.Sleep(time.Millisecond * 200)
	})

	gtest.C(t, func(t *gtest.T) {
		s := make([]int, 0)
		p := gpool.X创建(time.Millisecond*200, nil, func(i interface{}) {
			s = append(s, i.(int))
		})
		for i := 0; i < 5; i++ {
			p.X入栈(i)
			time.Sleep(time.Millisecond * 50)
		}
		val, err := p.X出栈()
		t.Assert(val, 2)
		t.AssertNil(err)
		t.Assert(p.X取数量(), 2)
	})
}
