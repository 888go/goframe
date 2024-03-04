
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with gm file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可协议条款。如果随gm文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf获取一个。
# <翻译结束>


<原文开始>
// TreeMap based on red-black tree, alias of RedBlackTree.
<原文结束>

# <翻译开始>
// TreeMap基于红黑树实现，是RedBlackTree的别名。
# <翻译结束>


<原文开始>
// NewTreeMap instantiates a tree map with the custom comparator.
// The parameter `safe` is used to specify whether using tree in concurrent-safety,
// which is false in default.
<原文结束>

# <翻译开始>
// NewTreeMap 创建一个使用自定义比较器的树形映射。
// 参数`safe`用于指定是否在并发安全的情况下使用树，其默认值为false。
# <翻译结束>


<原文开始>
// NewTreeMapFrom instantiates a tree map with the custom comparator and `data` map.
// Note that, the param `data` map will be set as the underlying data map(no deep copy),
// there might be some concurrent-safe issues when changing the map outside.
// The parameter `safe` is used to specify whether using tree in concurrent-safety,
// which is false in default.
<原文结束>

# <翻译开始>
// NewTreeMapFrom 通过自定义比较器和`data`映射实例化一个新的树形映射。
// 注意，参数`data`映射将被直接设置为底层数据映射（非深度复制），
// 因此在外部修改该映射时可能存在并发安全问题。
// 参数`safe`用于指定是否使用线程安全的树形结构，默认情况下为false。
# <翻译结束>

