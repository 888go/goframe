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
	"testing"

	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/json"
	gtest "github.com/888go/goframe/test/gtest"
)

func nilError() error {
	return nil
}

func Test_Nil(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gerror.X创建(""), nil)
		t.Assert(gerror.X多层错误(nilError(), "test"), nil)
	})
}

func Test_New(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		err := gerror.X创建("1")
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "1")
	})
	gtest.C(t, func(t *gtest.T) {
		err := gerror.X创建并格式化("%d", 1)
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "1")
	})
	gtest.C(t, func(t *gtest.T) {
		err := gerror.X创建并跳过堆栈(1, "1")
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "1")
	})
	gtest.C(t, func(t *gtest.T) {
		err := gerror.X创建并跳过堆栈与格式化(1, "%d", 1)
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "1")
	})
}

func Test_Wrap(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		err := errors.New("1")
		err = gerror.X多层错误(err, "2")
		err = gerror.X多层错误(err, "3")
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "3: 2: 1")
	})
	gtest.C(t, func(t *gtest.T) {
		err := gerror.X创建("1")
		err = gerror.X多层错误(err, "2")
		err = gerror.X多层错误(err, "3")
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "3: 2: 1")
	})
	gtest.C(t, func(t *gtest.T) {
		err := gerror.X创建("1")
		err = gerror.X多层错误(err, "")
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "1")
	})
}

func Test_Wrapf(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		err := errors.New("1")
		err = gerror.X多层错误并格式化(err, "%d", 2)
		err = gerror.X多层错误并格式化(err, "%d", 3)
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "3: 2: 1")
	})
	gtest.C(t, func(t *gtest.T) {
		err := gerror.X创建("1")
		err = gerror.X多层错误并格式化(err, "%d", 2)
		err = gerror.X多层错误并格式化(err, "%d", 3)
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "3: 2: 1")
	})
	gtest.C(t, func(t *gtest.T) {
		err := gerror.X创建("1")
		err = gerror.X多层错误并格式化(err, "")
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "1")
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gerror.X多层错误并格式化(nil, ""), nil)
	})
}

func Test_WrapSkip(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gerror.X多层错误并跳过堆栈(1, nil, "2"), nil)
		err := errors.New("1")
		err = gerror.X多层错误并跳过堆栈(1, err, "2")
		err = gerror.X多层错误并跳过堆栈(1, err, "3")
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "3: 2: 1")
	})
	gtest.C(t, func(t *gtest.T) {
		err := gerror.X创建("1")
		err = gerror.X多层错误并跳过堆栈(1, err, "2")
		err = gerror.X多层错误并跳过堆栈(1, err, "3")
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "3: 2: 1")
	})
	gtest.C(t, func(t *gtest.T) {
		err := gerror.X创建("1")
		err = gerror.X多层错误并跳过堆栈(1, err, "")
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "1")
	})
}

func Test_WrapSkipf(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gerror.X多层错误并跳过堆栈与格式化(1, nil, "2"), nil)
		err := errors.New("1")
		err = gerror.X多层错误并跳过堆栈与格式化(1, err, "2")
		err = gerror.X多层错误并跳过堆栈与格式化(1, err, "3")
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "3: 2: 1")
	})
	gtest.C(t, func(t *gtest.T) {
		err := gerror.X创建("1")
		err = gerror.X多层错误并跳过堆栈与格式化(1, err, "2")
		err = gerror.X多层错误并跳过堆栈与格式化(1, err, "3")
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "3: 2: 1")
	})
	gtest.C(t, func(t *gtest.T) {
		err := gerror.X创建("1")
		err = gerror.X多层错误并跳过堆栈与格式化(1, err, "")
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "1")
	})
}

func Test_Cause(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gerror.X取根错误(nil), nil)
		err := errors.New("1")
		t.Assert(gerror.X取根错误(err), err)
	})

	gtest.C(t, func(t *gtest.T) {
		err := errors.New("1")
		err = gerror.X多层错误(err, "2")
		err = gerror.X多层错误(err, "3")
		t.Assert(gerror.X取根错误(err), "1")
	})

	gtest.C(t, func(t *gtest.T) {
		err := gerror.X创建("1")
		t.Assert(gerror.X取根错误(err), "1")
	})

	gtest.C(t, func(t *gtest.T) {
		err := gerror.X创建("1")
		err = gerror.X多层错误(err, "2")
		err = gerror.X多层错误(err, "3")
		t.Assert(gerror.X取根错误(err), "1")
	})

	gtest.C(t, func(t *gtest.T) {
		t.Assert(gerror.X取文本(nil), "")
		err := errors.New("1")
		t.Assert(gerror.X取文本(err), err)
	})

	gtest.C(t, func(t *gtest.T) {
		var e *gerror.Error = nil
		t.Assert(e.Cause(), nil)
	})
}

func Test_Format(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		err := errors.New("1")
		err = gerror.X多层错误(err, "2")
		err = gerror.X多层错误(err, "3")
		t.AssertNE(err, nil)
		t.Assert(fmt.Sprintf("%s", err), "3: 2: 1")
		t.Assert(fmt.Sprintf("%v", err), "3: 2: 1")
	})

	gtest.C(t, func(t *gtest.T) {
		err := gerror.X创建("1")
		err = gerror.X多层错误(err, "2")
		err = gerror.X多层错误(err, "3")
		t.AssertNE(err, nil)
		t.Assert(fmt.Sprintf("%s", err), "3: 2: 1")
		t.Assert(fmt.Sprintf("%v", err), "3: 2: 1")
	})

	gtest.C(t, func(t *gtest.T) {
		err := gerror.X创建("1")
		err = gerror.X多层错误(err, "2")
		err = gerror.X多层错误(err, "3")
		t.AssertNE(err, nil)
		t.Assert(fmt.Sprintf("%-s", err), "3")
		t.Assert(fmt.Sprintf("%-v", err), "3")
	})
}

func Test_Stack(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		err := errors.New("1")
		t.Assert(fmt.Sprintf("%+v", err), "1")
	})

	gtest.C(t, func(t *gtest.T) {
		err := errors.New("1")
		err = gerror.X多层错误(err, "2")
		err = gerror.X多层错误(err, "3")
		t.AssertNE(err, nil)
		// fmt.Printf("%+v", err)
	})

	gtest.C(t, func(t *gtest.T) {
		err := gerror.X创建("1")
		t.AssertNE(fmt.Sprintf("%+v", err), "1")
		// fmt.Printf("%+v", err)
	})

	gtest.C(t, func(t *gtest.T) {
		err := gerror.X创建("1")
		err = gerror.X多层错误(err, "2")
		err = gerror.X多层错误(err, "3")
		t.AssertNE(err, nil)
		// fmt.Printf("%+v", err)
	})
}

func Test_Current(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gerror.X取当前错误(nil), nil)
		err := errors.New("1")
		err = gerror.X多层错误(err, "2")
		err = gerror.X多层错误(err, "3")
		t.Assert(err.Error(), "3: 2: 1")
		t.Assert(gerror.X取当前错误(err).Error(), "3")
	})
	gtest.C(t, func(t *gtest.T) {
		var e *gerror.Error = nil
		t.Assert(e.Current(), nil)
	})
}

func Test_Unwrap(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gerror.X取下一层错误(nil), nil)
		err := errors.New("1")
		err = gerror.X多层错误(err, "2")
		err = gerror.X多层错误(err, "3")
		t.Assert(err.Error(), "3: 2: 1")

		err = gerror.X取下一层错误(err)
		t.Assert(err.Error(), "2: 1")

		err = gerror.X取下一层错误(err)
		t.Assert(err.Error(), "1")

		err = gerror.X取下一层错误(err)
		t.AssertNil(err)
	})
	gtest.C(t, func(t *gtest.T) {
		var e *gerror.Error = nil
		t.Assert(e.Unwrap(), nil)
	})
}

func Test_Code(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		err := errors.New("123")
		t.Assert(gerror.X取错误码(err), -1)
		t.Assert(err.Error(), "123")
	})
	gtest.C(t, func(t *gtest.T) {
		err := gerror.X创建错误码(gcode.CodeUnknown, "123")
		t.Assert(gerror.X取错误码(err), gcode.CodeUnknown)
		t.Assert(err.Error(), "123")
	})
	gtest.C(t, func(t *gtest.T) {
		err := gerror.X创建错误码并格式化(gcode.New(1, "", nil), "%s", "123")
		t.Assert(gerror.X取错误码(err).Code(), 1)
		t.Assert(err.Error(), "123")
	})
	gtest.C(t, func(t *gtest.T) {
		err := gerror.X创建错误码并跳过堆栈(gcode.New(1, "", nil), 0, "123")
		t.Assert(gerror.X取错误码(err).Code(), 1)
		t.Assert(err.Error(), "123")
	})
	gtest.C(t, func(t *gtest.T) {
		err := gerror.X创建错误码并跳过堆栈与格式化(gcode.New(1, "", nil), 0, "%s", "123")
		t.Assert(gerror.X取错误码(err).Code(), 1)
		t.Assert(err.Error(), "123")
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gerror.X多层错误码(gcode.New(1, "", nil), nil, "3"), nil)
		err := errors.New("1")
		err = gerror.X多层错误(err, "2")
		err = gerror.X多层错误码(gcode.New(1, "", nil), err, "3")
		t.Assert(gerror.X取错误码(err).Code(), 1)
		t.Assert(err.Error(), "3: 2: 1")
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gerror.X多层错误码并格式化(gcode.New(1, "", nil), nil, "%s", "3"), nil)
		err := errors.New("1")
		err = gerror.X多层错误(err, "2")
		err = gerror.X多层错误码并格式化(gcode.New(1, "", nil), err, "%s", "3")
		t.Assert(gerror.X取错误码(err).Code(), 1)
		t.Assert(err.Error(), "3: 2: 1")
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gerror.X多层错误码并跳过堆栈(gcode.New(1, "", nil), 100, nil, "3"), nil)
		err := errors.New("1")
		err = gerror.X多层错误(err, "2")
		err = gerror.X多层错误码并跳过堆栈(gcode.New(1, "", nil), 100, err, "3")
		t.Assert(gerror.X取错误码(err).Code(), 1)
		t.Assert(err.Error(), "3: 2: 1")
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gerror.X多层错误码并跳过堆栈与格式化(gcode.New(1, "", nil), 100, nil, "%s", "3"), nil)
		err := errors.New("1")
		err = gerror.X多层错误(err, "2")
		err = gerror.X多层错误码并跳过堆栈与格式化(gcode.New(1, "", nil), 100, err, "%s", "3")
		t.Assert(gerror.X取错误码(err).Code(), 1)
		t.Assert(err.Error(), "3: 2: 1")
	})
}

func TestError_Error(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var e *gerror.Error = nil
		t.Assert(e.Error(), nil)
	})
}

func TestError_Code(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var e *gerror.Error = nil
		t.Assert(e.Code(), gcode.CodeNil)
	})
}

func Test_SetCode(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		err := gerror.X创建("123")
		t.Assert(gerror.X取错误码(err), -1)
		t.Assert(err.Error(), "123")

		err.(*gerror.Error).SetCode(gcode.CodeValidationFailed)
		t.Assert(gerror.X取错误码(err), gcode.CodeValidationFailed)
		t.Assert(err.Error(), "123")
	})
	gtest.C(t, func(t *gtest.T) {
		var err *gerror.Error = nil
		err.SetCode(gcode.CodeValidationFailed)
	})
}

func Test_Json(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		err := gerror.X多层错误(gerror.X创建("1"), "2")
		b, e := json.Marshal(err)
		t.Assert(e, nil)
		t.Assert(string(b), `"2: 1"`)
	})
}

func Test_HasStack(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		err1 := errors.New("1")
		err2 := gerror.X创建("1")
		t.Assert(gerror.X判断是否带堆栈(err1), false)
		t.Assert(gerror.X判断是否带堆栈(err2), true)
	})
}

func Test_Equal(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		err1 := errors.New("1")
		err2 := errors.New("1")
		err3 := gerror.X创建("1")
		err4 := gerror.X创建("4")
		t.Assert(gerror.X是否相等(err1, err2), false)
		t.Assert(gerror.X是否相等(err1, err3), true)
		t.Assert(gerror.X是否相等(err2, err3), true)
		t.Assert(gerror.X是否相等(err3, err4), false)
		t.Assert(gerror.X是否相等(err1, err4), false)
	})
	gtest.C(t, func(t *gtest.T) {
		var e = new(gerror.Error)
		t.Assert(e.Equal(e), true)
	})
}

func Test_Is(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		err1 := errors.New("1")
		err2 := gerror.X多层错误(err1, "2")
		err2 = gerror.X多层错误(err2, "3")
		t.Assert(gerror.X是否包含(err2, err1), true)
	})
}

func Test_HashError(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		err1 := errors.New("1")
		err2 := gerror.X多层错误(err1, "2")
		err2 = gerror.X多层错误(err2, "3")
		t.Assert(gerror.HasError别名(err2, err1), true)
	})
}

func Test_HashCode(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gerror.X是否包含错误码(nil, gcode.CodeNotAuthorized), false)
		err1 := errors.New("1")
		err2 := gerror.X多层错误码(gcode.CodeNotAuthorized, err1, "2")
		err3 := gerror.X多层错误(err2, "3")
		err4 := gerror.X多层错误(err3, "4")
		err5 := gerror.X多层错误码(gcode.CodeInvalidParameter, err4, "5")
		t.Assert(gerror.X是否包含错误码(err1, gcode.CodeNotAuthorized), false)
		t.Assert(gerror.X是否包含错误码(err2, gcode.CodeNotAuthorized), true)
		t.Assert(gerror.X是否包含错误码(err3, gcode.CodeNotAuthorized), true)
		t.Assert(gerror.X是否包含错误码(err4, gcode.CodeNotAuthorized), true)
		t.Assert(gerror.X是否包含错误码(err5, gcode.CodeNotAuthorized), true)
		t.Assert(gerror.X是否包含错误码(err5, gcode.CodeInvalidParameter), true)
		t.Assert(gerror.X是否包含错误码(err5, gcode.CodeInternalError), false)
	})
}

func Test_NewOption(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertNE(gerror.NewWithOption(gerror.Option{
			Error: errors.New("NewOptionError"),
			Stack: true,
			Text:  "Text",
			Code:  gcode.CodeNotAuthorized,
		}), gerror.X创建("NewOptionError"))
	})
}
