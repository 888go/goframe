		// 版权归GoFrame作者(https:		//goframe.org)所有。保留所有权利。
		//
		// 本源代码形式受MIT许可证条款约束。
		// 如果未随本文件一同分发MIT许可证副本，
		// 您可以在https:		//github.com/gogf/gf处获取。
		// md5:a9832f33b234e3f3

package gstructs_test

import (
	"reflect"
	"testing"

	"github.com/gogf/gf/v2/os/gstructs"
)

type User struct {
	Id   int
	Name string `params:"name"`
	Pass string `my-tag1:"pass1" my-tag2:"pass2" params:"pass"`
}

var (
	user           = new(User)
	userNilPointer *User
)

func Benchmark_ReflectTypeOf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reflect.TypeOf(user).String()
	}
}

func Benchmark_TagFields(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gstructs.TagFields(user, []string{"params", "my-tag1"})
	}
}

func Benchmark_TagFields_NilPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gstructs.TagFields(&userNilPointer, []string{"params", "my-tag1"})
	}
}
