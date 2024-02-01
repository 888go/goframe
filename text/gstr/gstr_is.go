// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gstr


import (
	"github.com/888go/goframe/internal/utils"
	)
// IsNumeric测试给定的字符串s是否为数值型。
func IsNumeric(s string) bool {
	return utils.IsNumeric(s)
}
