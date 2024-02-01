// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gerror_test
import (
	"errors"
	"fmt"
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	)

func ExampleNewCode() {
	err := gerror.NewCode(gcode.New(10000, "", nil), "My Error")
	fmt.Println(err.Error())
	fmt.Println(gerror.Code(err))

	// Output:
	// My Error
	// 10000
}

func ExampleNewCodef() {
	err := gerror.NewCodef(gcode.New(10000, "", nil), "It's %s", "My Error")
	fmt.Println(err.Error())
	fmt.Println(gerror.Code(err).Code())

	// Output:
	// It's My Error
	// 10000
}

func ExampleWrapCode() {
	err1 := errors.New("permission denied")
	err2 := gerror.WrapCode(gcode.New(10000, "", nil), err1, "Custom Error")
	fmt.Println(err2.Error())
	fmt.Println(gerror.Code(err2).Code())

	// Output:
	// Custom Error: permission denied
	// 10000
}

func ExampleWrapCodef() {
	err1 := errors.New("permission denied")
	err2 := gerror.WrapCodef(gcode.New(10000, "", nil), err1, "It's %s", "Custom Error")
	fmt.Println(err2.Error())
	fmt.Println(gerror.Code(err2).Code())

	// Output:
	// It's Custom Error: permission denied
	// 10000
}

func ExampleEqual() {
	err1 := errors.New("permission denied")
	err2 := gerror.New("permission denied")
	err3 := gerror.NewCode(gcode.CodeNotAuthorized, "permission denied")
	fmt.Println(gerror.Equal(err1, err2))
	fmt.Println(gerror.Equal(err2, err3))

	// Output:
	// true
	// false
}

func ExampleIs() {
	err1 := errors.New("permission denied")
	err2 := gerror.Wrap(err1, "operation failed")
	fmt.Println(gerror.Is(err1, err1))
	fmt.Println(gerror.Is(err2, err2))
	fmt.Println(gerror.Is(err2, err1))
	fmt.Println(gerror.Is(err1, err2))

	// Output:
	// false
	// true
	// true
	// false
}
