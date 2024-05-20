// 版权所有 (c) GoFrame (https://goframe.org)，保留所有权利。
//
// 本源代码遵循MIT许可协议。若未随此文件分发MIT许可证的副本，
// 您可以从 https://github.com/gogf/gf 获取。
// md5:c14c707c81272457

package gcode

import "fmt"

// localCode 是仅为内部使用而实现接口 Code 的一个实现者。. md5:1ee44496d8d5c874
type localCode struct {
	code    int         // 错误代码，通常为一个整数。. md5:7a0e76ca5b652d15
	message string      // 此错误代码的简要消息。. md5:97fc96d82cf974fd
	detail  interface{} // 作为接口类型，它主要设计为错误代码的扩展字段。. md5:39b4a34ec7b24f89
}

// Code 返回当前错误代码的整数值。. md5:75b8de0b4b9fa0a7
func (c localCode) Code() int {
	return c.code
}

// Message返回当前错误代码的简要消息。. md5:e0440d2d9a5b929c
func (c localCode) Message() string {
	return c.message
}

// Detail 返回当前错误代码的详细信息，
// 主要设计为错误代码的扩展字段。
// md5:b97007ef3f91efa2
func (c localCode) Detail() interface{} {
	return c.detail
}

// String 返回当前错误代码的字符串表示。. md5:bef1417927806ff8
func (c localCode) String() string {
	if c.detail != nil {
		return fmt.Sprintf(`%d:%s %v`, c.code, c.message, c.detail)
	}
	if c.message != "" {
		return fmt.Sprintf(`%d:%s`, c.code, c.message)
	}
	return fmt.Sprintf(`%d`, c.code)
}
