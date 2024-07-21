
<原文开始>
drivers

Powerful database drivers for package gdb.


<原文结束>

# <翻译开始>
// md5:34e63c84ac9f12dc
# <翻译结束>


<原文开始>
Installation

Let's take `mysql` for example.

```shell
go get -u github.com/gogf/gf/contrib/drivers/mysql/v2

<原文结束>

# <翻译开始>
# 安装教程

以 `mysql` 为例：

```shell
go get -u github.com/gogf/gf/contrib/drivers/mysql/v2
```

这是用于通过 Go 语言的包管理工具 `go get` 来安装 `gf` 框架的 MySQL 驱动。命令中的 `-u` 参数表示更新包到最新版本。

// md5:d1c2b69639894015
# <翻译结束>


<原文开始>
Easy to copy
go get -u github.com/gogf/gf/contrib/drivers/clickhouse/v2
go get -u github.com/gogf/gf/contrib/drivers/dm/v2
go get -u github.com/gogf/gf/contrib/drivers/mssql/v2
go get -u github.com/gogf/gf/contrib/drivers/oracle/v2
go get -u github.com/gogf/gf/contrib/drivers/pgsql/v2
go get -u github.com/gogf/gf/contrib/drivers/sqlite/v2
go get -u github.com/gogf/gf/contrib/drivers/sqlitecgo/v2
```

Choose and import the driver to your project:

```go
import _ "github.com/gogf/gf/contrib/drivers/mysql/v2"
```

Commonly imported at top of `main.go`:

```go
package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	// Other imported packages.
)

func main() {
	// Main logics.
}
```


<原文结束>

# <翻译开始>
# 易于复制
获取并更新以下GF贡献的数据库驱动：
```bash
go get -u github.com/gogf/gf/contrib/drivers/clickhouse/v2
go get -u github.com/gogf/gf/contrib/drivers/dm/v2
go get -u github.com/gogf/gf/contrib/drivers/mssql/v2
go get -u github.com/gogf/gf/contrib/drivers/oracle/v2
go get -u github.com/gogf/gf/contrib/drivers/pgsql/v2
go get -u github.com/gogf/gf/contrib/drivers/sqlite/v2
go get -u github.com/gogf/gf/contrib/drivers/sqlitecgo/v2
```
选择并导入你项目所需的驱动：
```go
import _ "github.com/gogf/gf/contrib/drivers/mysql/v2"
```
通常在`main.go`文件的顶部导入：
```go
package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	// 其他导入的包。
)

func main() {
	// 主要逻辑。
}
```

	// md5:0bbd7dd7e96240ed
# <翻译结束>


<原文开始>
MySQL/MariaDB/TiDB

```go
import _ "github.com/gogf/gf/contrib/drivers/mysql/v2"
```


<原文结束>

# <翻译开始>
// md5:993f339070bfdb48
# <翻译结束>


<原文开始>
SQLite

```go
import _ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
```

#
<原文结束>

# <翻译开始>
# SQLite

```go
导入_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
```

这段代码是Go语言的导入语句，它引入了`github.com/gogf/gf/contrib/drivers/sqlite/v2`包。`_`是一个空白标识符，用于仅导入包以便其在代码中可用但不使用其中的任何函数或类型。这个包提供SQLite（一个轻量级的关系型数据库）的驱动支持，用于GF（Golang Foundation）框架。

// md5:d096a47263dc3bdf
# <翻译结束>


<原文开始>
cgo version for 32-bit windows

```go
import _ "github.com/gogf/gf/contrib/drivers/sqlitecgo/v2"
```


<原文结束>

# <翻译开始>
// md5:b5c89ec6af200159
# <翻译结束>


<原文开始>
PostgreSQL

```go
import _ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
```

Note:

- It does not support `Replace` features.


<原文结束>

# <翻译开始>
# PostgreSQL

```go
导入 _ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
```

注意：

- 不支持`Replace`功能。

// md5:3b1d1e51b892a4d5
# <翻译结束>


<原文开始>
SQL Server

```go
import _ "github.com/gogf/gf/contrib/drivers/mssql/v2"
```

Note:

- It does not support `Replace` features.
- It does not support `LastInsertId`.
- It supports server version >= `SQL Server2005`
- It ONLY supports datetime2 and datetimeoffset types for auto handling created_at/updated_at/deleted_at columns, because datetime type does not support microseconds precision when column value is passed as string.


<原文结束>

# <翻译开始>
// md5:ae572708ba108995
# <翻译结束>


<原文开始>
Oracle

```go
import _ "github.com/gogf/gf/contrib/drivers/oracle/v2"
```

Note:

- It does not support `Replace` features.
- It does not support `LastInsertId`.


<原文结束>

# <翻译开始>
# Oracle

```go
import _ "github.com/gogf/gf/contrib/drivers/oracle/v2"
```

注意：

- 它不支持`Replace`特性。
- 它不支持`LastInsertId`。

// md5:9442ab8df9b797ce
# <翻译结束>


<原文开始>
ClickHouse

```go
import _ "github.com/gogf/gf/contrib/drivers/clickhouse/v2"
```

Note:

- It does not support `InsertIgnore/InsertGetId` features.
- It does not support `Save/Replace` features.
- It does not support `Transaction` feature.
- It does not support `RowsAffected` feature.


<原文结束>

# <翻译开始>
// md5:b60f4ceb35b1e184
# <翻译结束>


<原文开始>
DM

```go
import _ "github.com/gogf/gf/contrib/drivers/dm/v2"
```

Note:

- It does not support `Replace` features.


<原文结束>

# <翻译开始>
# 这是导入 DM 驱动的 Go 代码片段：

```go
import _ "github.com/gogf/gf/contrib/drivers/dm/v2"
```

注意：
- 它不支持 `Replace` 功能。

// md5:1d252b2894a4d8ce
# <翻译结束>


<原文开始>
Custom Drivers

It's quick and easy, please refer to current driver source.
It's quite appreciated if any PR for new drivers support into current repo.

<原文结束>

# <翻译开始>
// md5:a20a02aff96de6a1
# <翻译结束>

