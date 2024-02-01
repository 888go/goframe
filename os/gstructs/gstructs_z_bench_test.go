// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gstructs_test
import (
	"reflect"
	"testing"
	
	"github.com/888go/goframe/os/gstructs"
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
