// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 错误类_test

import (
	"errors"
	"fmt"
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
)

func ExampleNewCode() {
	err := 错误类.X创建错误码(错误码类.New(10000, "", nil), "My Error")
	fmt.Println(err.Error())
	fmt.Println(错误类.X取错误码(err))

	// Output:
	// My Error
	// 10000
}

func ExampleNewCodef() {
	err := 错误类.X创建错误码并格式化(错误码类.New(10000, "", nil), "It's %s", "My Error")
	fmt.Println(err.Error())
	fmt.Println(错误类.X取错误码(err).Code())

	// Output:
	// It's My Error
	// 10000
}

func ExampleWrapCode() {
	err1 := errors.New("permission denied")
	err2 := 错误类.X多层错误码(错误码类.New(10000, "", nil), err1, "Custom Error")
	fmt.Println(err2.Error())
	fmt.Println(错误类.X取错误码(err2).Code())

	// Output:
	// Custom Error: permission denied
	// 10000
}

func ExampleWrapCodef() {
	err1 := errors.New("permission denied")
	err2 := 错误类.X多层错误码并格式化(错误码类.New(10000, "", nil), err1, "It's %s", "Custom Error")
	fmt.Println(err2.Error())
	fmt.Println(错误类.X取错误码(err2).Code())

	// Output:
	// It's Custom Error: permission denied
	// 10000
}

func ExampleEqual() {
	err1 := errors.New("permission denied")
	err2 := 错误类.X创建("permission denied")
	err3 := 错误类.X创建错误码(错误码类.CodeNotAuthorized, "permission denied")
	fmt.Println(错误类.X是否相等(err1, err2))
	fmt.Println(错误类.X是否相等(err2, err3))

	// Output:
	// true
	// false
}

func ExampleIs() {
	err1 := errors.New("permission denied")
	err2 := 错误类.X多层错误(err1, "operation failed")
	fmt.Println(错误类.X是否包含(err1, err1))
	fmt.Println(错误类.X是否包含(err2, err2))
	fmt.Println(错误类.X是否包含(err2, err1))
	fmt.Println(错误类.X是否包含(err1, err2))

	// Output:
	// false
	// true
	// true
	// false
}
