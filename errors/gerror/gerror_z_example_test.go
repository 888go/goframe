// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 错误类_test

import (
	"errors"
	"fmt"

	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
)

func ExampleNewCode() {
	err := gerror.X创建错误码(gcode.New(10000, "", nil), "My Error")
	fmt.Println(err.Error())
	fmt.Println(gerror.X取错误码(err))

	// Output:
	// My Error
	// 10000
}

func ExampleNewCodef() {
	err := gerror.X创建错误码并格式化(gcode.New(10000, "", nil), "It's %s", "My Error")
	fmt.Println(err.Error())
	fmt.Println(gerror.X取错误码(err).Code())

	// Output:
	// It's My Error
	// 10000
}

func ExampleWrapCode() {
	err1 := errors.New("permission denied")
	err2 := gerror.X多层错误码(gcode.New(10000, "", nil), err1, "Custom Error")
	fmt.Println(err2.Error())
	fmt.Println(gerror.X取错误码(err2).Code())

	// Output:
	// Custom Error: permission denied
	// 10000
}

func ExampleWrapCodef() {
	err1 := errors.New("permission denied")
	err2 := gerror.X多层错误码并格式化(gcode.New(10000, "", nil), err1, "It's %s", "Custom Error")
	fmt.Println(err2.Error())
	fmt.Println(gerror.X取错误码(err2).Code())

	// Output:
	// It's Custom Error: permission denied
	// 10000
}

func ExampleEqual() {
	err1 := errors.New("permission denied")
	err2 := gerror.X创建("permission denied")
	err3 := gerror.X创建错误码(gcode.CodeNotAuthorized, "permission denied")
	fmt.Println(gerror.X是否相等(err1, err2))
	fmt.Println(gerror.X是否相等(err2, err3))

	// Output:
	// true
	// false
}

func ExampleIs() {
	err1 := errors.New("permission denied")
	err2 := gerror.X多层错误(err1, "operation failed")
	fmt.Println(gerror.X是否包含(err1, err1))
	fmt.Println(gerror.X是否包含(err2, err2))
	fmt.Println(gerror.X是否包含(err2, err1))
	fmt.Println(gerror.X是否包含(err1, err2))

	// Output:
	// false
	// true
	// true
	// false
}
