// 版权所有 GoFrame gf 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一个。

package gcode

import (
	"fmt"
)

// localCode 是一个仅用于内部使用的 Code 接口的实现者。
type localCode struct {
	code    int         // 错误代码，通常是一个整数。
	message string      // 此错误代码的简短消息。
	detail  interface{} // 作为接口类型，它主要设计为错误码的扩展字段。
}

// Code 返回当前错误代码的整数值。
func (c localCode) Code() int {
	return c.code
}

// Message 返回当前错误代码的简短消息。
func (c localCode) Message() string {
	return c.message
}

// Detail 返回当前错误代码的详细信息，
// 主要设计为错误代码的扩展字段。
func (c localCode) Detail() interface{} {
	return c.detail
}

// String 将当前错误代码以字符串形式返回。
func (c localCode) String() string {
	if c.detail != nil {
		return fmt.Sprintf(`%d:%s %v`, c.code, c.message, c.detail)
	}
	if c.message != "" {
		return fmt.Sprintf(`%d:%s`, c.code, c.message)
	}
	return fmt.Sprintf(`%d`, c.code)
}
