
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
		// Specially in Golang:
		// "v1.12.2-0.20200413154443-b17e3a6804fa" < "v1.12.2"
		// "v1.12.3-0.20200413154443-b17e3a6804fa" > "v1.12.2"
<原文结束>

# <翻译开始>
// 特别地在 Golang 中：
// "v1.12.2-0.20200413154443-b17e3a6804fa" 小于 "v1.12.2"
// "v1.12.3-0.20200413154443-b17e3a6804fa" 大于 "v1.12.2"
// 这段代码注释是关于 Golang 中版本字符串比较的特殊规则：
// 在 Golang 中，对于包含预发布版本号（如 "-0.20200413154443-b17e3a6804fa"）的版本字符串，在进行字符串比较时，主版本号、次版本号和补丁版本号部分会被优先比较。当这部分相同时，带有预发布版本号的版本会认为小于不带预发布版本号的版本。
// 因此，尽管 "v1.12.2-0.20200413154443-b17e3a6804fa" 的主要部分与 "v1.12.2" 相同，但由于其附加了预发布标识，所以在比较中它被认为小于 "v1.12.2"。
// 同样，"v1.12.3-0.20200413154443-b17e3a6804fa" 由于其主版本号部分高于 "v1.12.2"，所以即使它也有预发布版本号，依然会在比较中大于 "v1.12.2"。
# <翻译结束>


<原文开始>
		// Specially in Golang:
		// "v4.20.1+incompatible" < "v4.20.1"
<原文结束>

# <翻译开始>
// 特别在 Golang 中：
// "v4.20.1+incompatible" < "v4.20.1"
// （译注：这里表示在 Go 语言中，带后缀 "+incompatible" 的版本字符串比较时，会被视为小于不带此后缀的相同主版本号、次版本号和补丁号的版本。这是因为在 Go 语言模块系统中，“+incompatible”表示该版本并非遵循 Go 语义化版本规范，因此在排序时会排在符合规范的版本之后。）
# <翻译结束>


<原文开始>
// go test *.go -bench=".*"
<原文结束>

# <翻译开始>
// 运行go test命令，测试当前目录下所有.go文件，并执行所有benchmark测试
# <翻译结束>

