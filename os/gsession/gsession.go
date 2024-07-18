// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// gsession 包实现了会话的管理与存储功能。 md5:743dbbcf3d74735d
package gsession//bm:session类

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/guid"
)

var (
	// ErrorDisabled 用于标记某些接口函数未被使用。 md5:2a81adf17310353b
	ErrorDisabled = gerror.NewWithOption(gerror.Option{
		Text: "this feature is disabled in this storage",
		Code: gcode.CodeNotSupported,
	})
)

// NewSessionId 创建并返回一个新的、唯一的会话ID字符串，该字符串长度为32字节。
// md5:c16016979687f8e5
// ff:
func NewSessionId() string {
	return guid.S()
}
