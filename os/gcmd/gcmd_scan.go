// 版权所有 GoFrame 作者(https://goframe.org)。保留所有权利。
//
// 此源代码形式受 MIT 许可证的条款约束。
// 如果未随此文件一起分发 MIT 许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
// md5:a114f4bdd106ab31

package gcmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gogf/gf/v2/text/gstr"
)

// Scan 将 `info` 打印到标准输出，读取并返回用户输入，直到遇到 '\n'。 md5:ddd0cd56978ea021
func Scan(info ...interface{}) string {
	fmt.Print(info...)
	return readline()
}

// Scanf 将 `info` 按照 `format` 格式打印到标准输出，然后读取并返回用户输入，直到遇到换行符'\n'停止。 md5:8aa27cd5ac6f9224
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
