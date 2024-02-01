// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gfile_test
import (
	"fmt"
	
	"github.com/888go/goframe/os/gfile"
	)

func ExampleHome() {
	// 用户的主目录
	homePath, _ := gfile.Home()
	fmt.Println(homePath)

	// May Output:
	// C:\Users\hailaz
}
