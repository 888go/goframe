
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
// DriverDefault is the default driver for mysql database, which does nothing.
<原文结束>

# <翻译开始>
// DriverDefault 是 mysql 数据库的默认驱动，它实际上什么都不做。
# <翻译结束>


<原文开始>
// New creates and returns a database object for mysql.
// It implements the interface of gdb.Driver for extra database driver installation.
<原文结束>

# <翻译开始>
// New 创建并返回一个用于 mysql 的数据库对象。
// 它实现了 gdb.Driver 接口，以便进行额外的数据库驱动安装。
# <翻译结束>


<原文开始>
// Open creates and returns an underlying sql.DB object for mysql.
// Note that it converts time.Time argument to local timezone in default.
<原文结束>

# <翻译开始>
// Open 创建并返回一个用于 mysql 的底层 sql.DB 对象。
// 注意，它默认会将 time.Time 类型参数转换为本地时区。
# <翻译结束>


<原文开始>
// PingMaster pings the master node to check authentication or keeps the connection alive.
<原文结束>

# <翻译开始>
// PingMaster 用于向主节点发送心跳以检查身份验证或保持连接存活。
# <翻译结束>


<原文开始>
// PingSlave pings the slave node to check authentication or keeps the connection alive.
<原文结束>

# <翻译开始>
// PingSlave 向从节点发送ping请求，用于检查身份验证或保持连接活跃。
# <翻译结束>

