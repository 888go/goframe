
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
// goRootForFilter is used for stack filtering purpose.
<原文结束>

# <翻译开始>
	// goRootForFilter 用于栈过滤目的。 md5:538cfd57e5493ca3
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
// MainPkgPath 返回包含入口函数main的package main的绝对文件路径。
//
// 它仅在开发环境中可用。
//
// 注意1：仅对源代码开发环境有效，
// 即仅对生成此可执行文件的系统有效。
//
// 注意2：首次调用此方法时，如果处于异步goroutine中，
// 方法可能无法获取到main包的路径。
// md5:7fb1d2fdcb626f85
# <翻译结束>


<原文开始>
// It is only for source development environments.
<原文结束>

# <翻译开始>
	// 仅供源代码开发环境使用。 md5:56e807aeb00eee19
# <翻译结束>


<原文开始>
			// Check if it is called in package initialization function,
			// in which it here cannot retrieve main package path,
			// it so just returns that can make next check.
<原文结束>

# <翻译开始>
			// 检查它是否在包初始化函数中被调用，
			// 在这种情况下，无法获取主包路径，
			// 所以直接返回，以便进行下一步检查。
			// md5:e583ee52c2832f4d
# <翻译结束>


<原文开始>
	// If it still cannot find the path of the package main,
	// it recursively searches the directory and its parents directory of the last go file.
	// It's usually necessary for uint testing cases of business project.
<原文结束>

# <翻译开始>
	// 如果仍然无法找到main包的路径，它会递归地搜索最后一个go文件的目录及其父目录。这对于商业项目中的整数测试用例通常是必要的。
	// md5:5bee1ce703ae05d8
# <翻译结束>

