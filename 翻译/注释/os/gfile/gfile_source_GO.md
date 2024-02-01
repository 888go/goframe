
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
// goRootForFilter is used for stack filtering purpose.
<原文结束>

# <翻译开始>
// goRootForFilter 用于堆栈过滤的目的。
# <翻译结束>


<原文开始>
// MainPkgPath returns absolute file path of package main,
// which contains the entrance function main.
//
// It's only available in develop environment.
//
// Note1: Only valid for source development environments,
// IE only valid for systems that generate this executable.
//
// Note2: When the method is called for the first time, if it is in an asynchronous goroutine,
// the method may not get the main package path.
<原文结束>

# <翻译开始>
// MainPkgPath 返回包含入口函数 main 的 main 包的绝对文件路径。
//
// 该功能仅在开发环境中可用。
//
// 注意1：仅对源代码开发环境有效，即仅对生成此可执行文件的系统有效。
//
// 注意2：当首次调用该方法时，如果处于异步 goroutine 中，则该方法可能无法获取到 main 包的路径。
# <翻译结束>


<原文开始>
// It is only for source development environments.
<原文结束>

# <翻译开始>
// 这仅适用于源代码开发环境。
# <翻译结束>


<原文开始>
			// Check if it is called in package initialization function,
			// in which it here cannot retrieve main package path,
			// it so just returns that can make next check.
<原文结束>

# <翻译开始>
// 检查是否在包初始化函数中被调用，
// 在这种情况下，此处无法获取主包路径，
// 因此仅返回一个值以便进行后续检查。
# <翻译结束>


<原文开始>
	// If it still cannot find the path of the package main,
	// it recursively searches the directory and its parents directory of the last go file.
	// It's usually necessary for uint testing cases of business project.
<原文结束>

# <翻译开始>
// 如果仍然无法找到包main的路径，
// 它会递归地搜索最后一个go文件所在的目录及其父目录。
// 这通常对于业务项目进行单元测试的情况是必要的。
# <翻译结束>

