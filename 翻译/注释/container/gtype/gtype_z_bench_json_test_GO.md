
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
# <翻译结束>


<原文开始>
// go test *.go -bench=".+\_Json" -benchmem
<原文结束>

# <翻译开始>
// 运行Go语言测试，对所有.go文件进行测试，并且仅针对名称中包含"_Json"的基准测试（benchmark）进行执行，同时显示内存分配统计信息。
// 以下是逐行详细解释：
// ```go
// 使用go test命令来运行测试
// 测试的文件为当前目录下所有的.go文件（即`*.go`）
// `-bench=".+\_Json"` 参数表示仅执行那些基准测试函数名称中包含"_Json"的基准测试
// `-benchmem` 参数表示在输出的基准测试结果中，包含内存分配统计信息
# <翻译结束>

