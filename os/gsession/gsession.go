// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// gsession包实现了会话的管理器和存储功能。
package session类

import (
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/util/guid"
)

var (
	// ErrorDisabled 用于标记某些接口函数未被使用。
	ErrorDisabled = 错误类.NewWithOption(错误类.Option{
		Text: "this feature is disabled in this storage",
		Code: 错误码类.CodeNotSupported,
	})
)

// NewSessionId 创建并返回一个新的、唯一的会话ID字符串，
// 这个字符串长度为32字节。
func NewSessionId() string {
	return uid类.X生成()
}
