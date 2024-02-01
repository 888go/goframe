
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
// go test *.go -bench=".*" -benchmem
<原文结束>

# <翻译开始>
// 运行go test命令，测试所有.go文件，并执行基准测试（-bench=".*"），同时显示内存使用情况统计（-benchmem）
# <翻译结束>


<原文开始>
		// build
		//-n, --name       output binary name
		//-v, --version    output binary version
		//-a, --arch       output binary architecture, multiple arch separated with ','
		//-s, --system     output binary system, multiple os separated with ','
		//-o, --output     output binary path, used when building single binary file
		//-p, --path       output binary directory path, default is './bin'
		//-e, --extra      extra custom "go build" options
		//-m, --mod        like "-mod" option of "go build", use "-m none" to disable go module
		//-c, --cgo        enable or disable cgo feature, it's disabled in default
<原文结束>

# <翻译开始>
// 构建
//-n, --name       输出二进制文件名称
//-v, --version    输出二进制文件版本信息
//-a, --arch       输出二进制文件架构，多个架构使用','分隔
//-s, --system     输出二进制文件系统，多个操作系统使用','分隔
//-o, --output     输出二进制文件路径，仅在构建单个二进制文件时使用
//-p, --path       输出二进制目录路径，默认为'./bin'
//-e, --extra      额外的自定义 "go build" 选项
//-m, --mod        类似于 "go build" 命令中的 "-mod" 选项，使用 "-m none" 禁用 Go 模块功能
//-c, --cgo        启用或禁用 cgo 功能，默认是禁用状态
// 以上代码是对golang构建命令行参数的注释翻译，这些参数用于控制构建过程中的各种行为，如指定输出文件名、版本、架构、操作系统、输出路径等，并可对Go模块和cgo特性进行控制。
# <翻译结束>

