// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gfile_test

import (
	"fmt"

	"github.com/gogf/gf/v2/os/gfile"
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
