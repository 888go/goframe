// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var content string
	var output string
	flag.StringVar(&content, "c", "", "写入内容")
	flag.StringVar(&output, "o", "", "写入路径")
	flag.Parse()
	fmt.Println(os.Args)
	fmt.Println(content)
	fmt.Println(output)
	if output != "" {
		file, err := os.Create(output)
		if err != nil {
			panic("create file fail: " + err.Error())
		}
		defer file.Close()
		file.WriteString(content)
	}
}
