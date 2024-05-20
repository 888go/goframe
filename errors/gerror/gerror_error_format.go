// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gerror

import (
	"fmt"
	"io"
)

// Format 根据fmt.Formatter接口格式化框架。
// 
// %v, %s：打印所有错误字符串；
// %-v, %-s：打印当前级别的错误字符串；
// %+s：打印完整的堆栈错误列表；
// %+v：打印错误字符串和完整的堆栈错误列表。
// md5:68a2fa33dd7a1faa
func (err *Error) Format(s fmt.State, verb rune) {
	switch verb {
	case 's', 'v':
		switch {
		case s.Flag('-'):
			if err.text != "" {
				_, _ = io.WriteString(s, err.text)
			} else {
				_, _ = io.WriteString(s, err.Error())
			}
		case s.Flag('+'):
			if verb == 's' {
				_, _ = io.WriteString(s, err.Stack())
			} else {
				_, _ = io.WriteString(s, err.Error()+"\n"+err.Stack())
			}
		default:
			_, _ = io.WriteString(s, err.Error())
		}
	}
}
