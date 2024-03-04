// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gfile_test

import (
	"fmt"
	
	"github.com/888go/goframe/gfile"
)

func ExampleMTime() {
	t := gfile.MTime(gfile.Temp())
	fmt.Println(t)

	// May Output:
	// 2021-11-02 15:18:43.901141 +0800 CST
}

func ExampleMTimestamp() {
	t := gfile.MTimestamp(gfile.Temp())
	fmt.Println(t)

	// May Output:
	// 1635838398
}

func ExampleMTimestampMilli() {
	t := gfile.MTimestampMilli(gfile.Temp())
	fmt.Println(t)

	// May Output:
	// 1635838529330
}
