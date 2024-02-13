// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 错误类_test

import (
	"errors"
	"fmt"
	"testing"
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/test/gtest"
)

func nilError() error {
	return nil
}

func Test_Nil(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(错误类.X创建(""), nil)
		t.Assert(错误类.X多层错误(nilError(), "test"), nil)
	})
}

func Test_New(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		err := 错误类.X创建("1")
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "1")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		err := 错误类.X创建并格式化("%d", 1)
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "1")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		err := 错误类.X创建并跳过堆栈(1, "1")
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "1")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		err := 错误类.X创建并跳过堆栈与格式化(1, "%d", 1)
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "1")
	})
}

func Test_Wrap(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		err := errors.New("1")
		err = 错误类.X多层错误(err, "2")
		err = 错误类.X多层错误(err, "3")
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "3: 2: 1")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		err := 错误类.X创建("1")
		err = 错误类.X多层错误(err, "2")
		err = 错误类.X多层错误(err, "3")
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "3: 2: 1")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		err := 错误类.X创建("1")
		err = 错误类.X多层错误(err, "")
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "1")
	})
}

func Test_Wrapf(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		err := errors.New("1")
		err = 错误类.X多层错误并格式化(err, "%d", 2)
		err = 错误类.X多层错误并格式化(err, "%d", 3)
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "3: 2: 1")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		err := 错误类.X创建("1")
		err = 错误类.X多层错误并格式化(err, "%d", 2)
		err = 错误类.X多层错误并格式化(err, "%d", 3)
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "3: 2: 1")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		err := 错误类.X创建("1")
		err = 错误类.X多层错误并格式化(err, "")
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "1")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(错误类.X多层错误并格式化(nil, ""), nil)
	})
}

func Test_WrapSkip(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(错误类.X多层错误并跳过堆栈(1, nil, "2"), nil)
		err := errors.New("1")
		err = 错误类.X多层错误并跳过堆栈(1, err, "2")
		err = 错误类.X多层错误并跳过堆栈(1, err, "3")
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "3: 2: 1")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		err := 错误类.X创建("1")
		err = 错误类.X多层错误并跳过堆栈(1, err, "2")
		err = 错误类.X多层错误并跳过堆栈(1, err, "3")
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "3: 2: 1")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		err := 错误类.X创建("1")
		err = 错误类.X多层错误并跳过堆栈(1, err, "")
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "1")
	})
}

func Test_WrapSkipf(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(错误类.X多层错误并跳过堆栈与格式化(1, nil, "2"), nil)
		err := errors.New("1")
		err = 错误类.X多层错误并跳过堆栈与格式化(1, err, "2")
		err = 错误类.X多层错误并跳过堆栈与格式化(1, err, "3")
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "3: 2: 1")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		err := 错误类.X创建("1")
		err = 错误类.X多层错误并跳过堆栈与格式化(1, err, "2")
		err = 错误类.X多层错误并跳过堆栈与格式化(1, err, "3")
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "3: 2: 1")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		err := 错误类.X创建("1")
		err = 错误类.X多层错误并跳过堆栈与格式化(1, err, "")
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "1")
	})
}

func Test_Cause(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(错误类.X取根错误(nil), nil)
		err := errors.New("1")
		t.Assert(错误类.X取根错误(err), err)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		err := errors.New("1")
		err = 错误类.X多层错误(err, "2")
		err = 错误类.X多层错误(err, "3")
		t.Assert(错误类.X取根错误(err), "1")
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		err := 错误类.X创建("1")
		t.Assert(错误类.X取根错误(err), "1")
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		err := 错误类.X创建("1")
		err = 错误类.X多层错误(err, "2")
		err = 错误类.X多层错误(err, "3")
		t.Assert(错误类.X取根错误(err), "1")
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(错误类.X取文本(nil), "")
		err := errors.New("1")
		t.Assert(错误类.X取文本(err), err)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		var e *错误类.Error = nil
		t.Assert(e.Cause(), nil)
	})
}

func Test_Format(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		err := errors.New("1")
		err = 错误类.X多层错误(err, "2")
		err = 错误类.X多层错误(err, "3")
		t.AssertNE(err, nil)
		t.Assert(fmt.Sprintf("%s", err), "3: 2: 1")
		t.Assert(fmt.Sprintf("%v", err), "3: 2: 1")
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		err := 错误类.X创建("1")
		err = 错误类.X多层错误(err, "2")
		err = 错误类.X多层错误(err, "3")
		t.AssertNE(err, nil)
		t.Assert(fmt.Sprintf("%s", err), "3: 2: 1")
		t.Assert(fmt.Sprintf("%v", err), "3: 2: 1")
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		err := 错误类.X创建("1")
		err = 错误类.X多层错误(err, "2")
		err = 错误类.X多层错误(err, "3")
		t.AssertNE(err, nil)
		t.Assert(fmt.Sprintf("%-s", err), "3")
		t.Assert(fmt.Sprintf("%-v", err), "3")
	})
}

func Test_Stack(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		err := errors.New("1")
		t.Assert(fmt.Sprintf("%+v", err), "1")
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		err := errors.New("1")
		err = 错误类.X多层错误(err, "2")
		err = 错误类.X多层错误(err, "3")
		t.AssertNE(err, nil)
		// 使用 %+v 格式化输出错误信息，会包含错误类型和详细堆栈信息
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		err := 错误类.X创建("1")
		t.AssertNE(fmt.Sprintf("%+v", err), "1")
		// 使用 %+v 格式化输出错误信息，会包含错误类型和详细堆栈信息
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		err := 错误类.X创建("1")
		err = 错误类.X多层错误(err, "2")
		err = 错误类.X多层错误(err, "3")
		t.AssertNE(err, nil)
		// 使用 %+v 格式化输出错误信息，会包含错误类型和详细堆栈信息
	})
}

func Test_Current(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(错误类.X取当前错误(nil), nil)
		err := errors.New("1")
		err = 错误类.X多层错误(err, "2")
		err = 错误类.X多层错误(err, "3")
		t.Assert(err.Error(), "3: 2: 1")
		t.Assert(错误类.X取当前错误(err).Error(), "3")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var e *错误类.Error = nil
		t.Assert(e.Current(), nil)
	})
}

func Test_Unwrap(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(错误类.X取下一层错误(nil), nil)
		err := errors.New("1")
		err = 错误类.X多层错误(err, "2")
		err = 错误类.X多层错误(err, "3")
		t.Assert(err.Error(), "3: 2: 1")

		err = 错误类.X取下一层错误(err)
		t.Assert(err.Error(), "2: 1")

		err = 错误类.X取下一层错误(err)
		t.Assert(err.Error(), "1")

		err = 错误类.X取下一层错误(err)
		t.AssertNil(err)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var e *错误类.Error = nil
		t.Assert(e.Unwrap(), nil)
	})
}

func Test_Code(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		err := errors.New("123")
		t.Assert(错误类.X取错误码(err), -1)
		t.Assert(err.Error(), "123")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		err := 错误类.X创建错误码(错误码类.CodeUnknown, "123")
		t.Assert(错误类.X取错误码(err), 错误码类.CodeUnknown)
		t.Assert(err.Error(), "123")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		err := 错误类.X创建错误码并格式化(错误码类.New(1, "", nil), "%s", "123")
		t.Assert(错误类.X取错误码(err).Code(), 1)
		t.Assert(err.Error(), "123")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		err := 错误类.X创建错误码并跳过堆栈(错误码类.New(1, "", nil), 0, "123")
		t.Assert(错误类.X取错误码(err).Code(), 1)
		t.Assert(err.Error(), "123")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		err := 错误类.X创建错误码并跳过堆栈与格式化(错误码类.New(1, "", nil), 0, "%s", "123")
		t.Assert(错误类.X取错误码(err).Code(), 1)
		t.Assert(err.Error(), "123")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(错误类.X多层错误码(错误码类.New(1, "", nil), nil, "3"), nil)
		err := errors.New("1")
		err = 错误类.X多层错误(err, "2")
		err = 错误类.X多层错误码(错误码类.New(1, "", nil), err, "3")
		t.Assert(错误类.X取错误码(err).Code(), 1)
		t.Assert(err.Error(), "3: 2: 1")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(错误类.X多层错误码并格式化(错误码类.New(1, "", nil), nil, "%s", "3"), nil)
		err := errors.New("1")
		err = 错误类.X多层错误(err, "2")
		err = 错误类.X多层错误码并格式化(错误码类.New(1, "", nil), err, "%s", "3")
		t.Assert(错误类.X取错误码(err).Code(), 1)
		t.Assert(err.Error(), "3: 2: 1")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(错误类.X多层错误码并跳过堆栈(错误码类.New(1, "", nil), 100, nil, "3"), nil)
		err := errors.New("1")
		err = 错误类.X多层错误(err, "2")
		err = 错误类.X多层错误码并跳过堆栈(错误码类.New(1, "", nil), 100, err, "3")
		t.Assert(错误类.X取错误码(err).Code(), 1)
		t.Assert(err.Error(), "3: 2: 1")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(错误类.X多层错误码并跳过堆栈与格式化(错误码类.New(1, "", nil), 100, nil, "%s", "3"), nil)
		err := errors.New("1")
		err = 错误类.X多层错误(err, "2")
		err = 错误类.X多层错误码并跳过堆栈与格式化(错误码类.New(1, "", nil), 100, err, "%s", "3")
		t.Assert(错误类.X取错误码(err).Code(), 1)
		t.Assert(err.Error(), "3: 2: 1")
	})
}

func TestError_Error(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var e *错误类.Error = nil
		t.Assert(e.Error(), nil)
	})
}

func TestError_Code(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var e *错误类.Error = nil
		t.Assert(e.Code(), 错误码类.CodeNil)
	})
}

func Test_SetCode(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		err := 错误类.X创建("123")
		t.Assert(错误类.X取错误码(err), -1)
		t.Assert(err.Error(), "123")

		err.(*错误类.Error).SetCode(错误码类.CodeValidationFailed)
		t.Assert(错误类.X取错误码(err), 错误码类.CodeValidationFailed)
		t.Assert(err.Error(), "123")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var err *错误类.Error = nil
		err.SetCode(错误码类.CodeValidationFailed)
	})
}

func Test_Json(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		err := 错误类.X多层错误(错误类.X创建("1"), "2")
		b, e := json.Marshal(err)
		t.Assert(e, nil)
		t.Assert(string(b), `"2: 1"`)
	})
}

func Test_HasStack(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		err1 := errors.New("1")
		err2 := 错误类.X创建("1")
		t.Assert(错误类.X判断是否带堆栈(err1), false)
		t.Assert(错误类.X判断是否带堆栈(err2), true)
	})
}

func Test_Equal(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		err1 := errors.New("1")
		err2 := errors.New("1")
		err3 := 错误类.X创建("1")
		err4 := 错误类.X创建("4")
		t.Assert(错误类.X是否相等(err1, err2), false)
		t.Assert(错误类.X是否相等(err1, err3), true)
		t.Assert(错误类.X是否相等(err2, err3), true)
		t.Assert(错误类.X是否相等(err3, err4), false)
		t.Assert(错误类.X是否相等(err1, err4), false)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var e = new(错误类.Error)
		t.Assert(e.Equal(e), true)
	})
}

func Test_Is(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		err1 := errors.New("1")
		err2 := 错误类.X多层错误(err1, "2")
		err2 = 错误类.X多层错误(err2, "3")
		t.Assert(错误类.X是否包含(err2, err1), true)
	})
}

func Test_HashError(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		err1 := errors.New("1")
		err2 := 错误类.X多层错误(err1, "2")
		err2 = 错误类.X多层错误(err2, "3")
		t.Assert(错误类.HasError别名(err2, err1), true)
	})
}

func Test_HashCode(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(错误类.X是否包含错误码(nil, 错误码类.CodeNotAuthorized), false)
		err1 := errors.New("1")
		err2 := 错误类.X多层错误码(错误码类.CodeNotAuthorized, err1, "2")
		err3 := 错误类.X多层错误(err2, "3")
		err4 := 错误类.X多层错误(err3, "4")
		t.Assert(错误类.X是否包含错误码(err1, 错误码类.CodeNotAuthorized), false)
		t.Assert(错误类.X是否包含错误码(err2, 错误码类.CodeNotAuthorized), true)
		t.Assert(错误类.X是否包含错误码(err3, 错误码类.CodeNotAuthorized), true)
		t.Assert(错误类.X是否包含错误码(err4, 错误码类.CodeNotAuthorized), true)
	})
}

func Test_NewOption(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertNE(错误类.NewWithOption(错误类.Option{
			Error: errors.New("NewOptionError"),
			Stack: true,
			Text:  "Text",
			Code:  错误码类.CodeNotAuthorized,
		}), 错误类.X创建("NewOptionError"))
	})
}
