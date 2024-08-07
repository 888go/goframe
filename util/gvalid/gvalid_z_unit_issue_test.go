// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 效验类_test

import (
	"context"
	"testing"

	gvalid "github.com/888go/goframe/util/gvalid"
)

type Foo struct {
	Bar *Bar `p:"bar" v:"required-without:Baz"`
	Baz *Baz `p:"baz" v:"required-without:Bar"`
}
type Bar struct {
	BarKey string `p:"bar_key" v:"required"`
}
type Baz struct {
	BazKey string `p:"baz_key" v:"required"`
}

//github.com/gogf/gf/issues/2503. md5:f22d150cb76ec306
func Test_Issue2503(t *testing.T) {
	foo := &Foo{
		Bar: &Bar{BarKey: "value"},
	}
	err := gvalid.New().Data(foo).Run(context.Background())
	if err != nil {
		t.Fatal(err)
	}
}
