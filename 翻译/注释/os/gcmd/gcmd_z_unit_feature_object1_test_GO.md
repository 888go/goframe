
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// go test *.go -bench=".*" -benchmem
<原文结束>

# <翻译开始>
// 使用`go test`命令，对所有`.go`文件进行测试，指定运行基准测试（Benchmark）中的所有模式（".*"），同时输出内存使用情况（-benchmem）。 md5:81db3d7bd1ed4da8
# <翻译结束>


<原文开始>
// test name tag and orphan tag true
<原文结束>

# <翻译开始>
		// 测试名称标签和孤儿标签为真. md5:5b0679661bf5c22b
# <翻译结束>


<原文开始>
		//os.Args = []string{"root", "test", "a", "b", "c", "-h"}
		//value, err := cmd.RunWithValueError(ctx)
		//t.AssertNil(err)
<原文结束>

# <翻译开始>
		// os.Args 是一个字符串切片，内容为 ["root", "test", "a", "b", "c", "-h"]
		// 使用 cmd.RunWithValueError 函数运行命令并获取值和错误
		// 使用 t.AssertNil 函数断言错误应为 nil（即无错误）
		// md5:9a669f8340465dad
# <翻译结束>

