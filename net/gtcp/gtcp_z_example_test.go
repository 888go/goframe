// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package tcp类_test

import (
	"fmt"
	
	"github.com/888go/goframe/net/gtcp"
)

func ExampleGetFreePort() {
	fmt.Println(tcp类.GetFreePort())

	// May Output:
	// 57429 <nil>
}

func ExampleGetFreePorts() {
	fmt.Println(tcp类.GetFreePorts(2))

	// May Output:
	// [57743 57744] <nil>
}
