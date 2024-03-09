// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 文件类_test

import (
	"fmt"
	
	"github.com/888go/goframe/gfile"
)

func ExampleHome() {
	// 用户的主目录
	homePath, _ := 文件类.X取用户目录()
	fmt.Println(homePath)

	// May Output:
	// C:\Users\hailaz
}
