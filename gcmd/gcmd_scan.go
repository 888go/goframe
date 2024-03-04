// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可证条款。
// 如果随此文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf 获取一份。
//

package gcmd

import (
	"bufio"
	"fmt"
	"os"
	
	"github.com/gogf/gf/v2/text/gstr"
)

// Scan 将 `info` 输出到标准输出（stdout），然后读取并返回用户输入，直到遇到换行符('\n')时停止。
func Scan(info ...interface{}) string {
	fmt.Print(info...)
	return readline()
}

// Scanf 函数将按照 `format` 格式打印 `info` 到标准输出（stdout），然后读取用户输入，直到遇到换行符 '\n' 为止，并返回所读取的用户输入内容。
func Scanf(format string, info ...interface{}) string {
	fmt.Printf(format, info...)
	return readline()
}

func readline() string {
	var s string
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	s = gstr.Trim(s)
	return s
}
