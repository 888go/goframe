# // md5:34e63c84ac9f12dc# # 安装教程

以 `mysql` 为例：

```shell
go get -u github.com/gogf/gf/contrib/drivers/mysql/v2
```

这是用于通过 Go 语言的包管理工具 `go get` 来安装 `gf` 框架的 MySQL 驱动。命令中的 `-u` 参数表示更新包到最新版本。

# 安装教程

以 `mysql` 为例：

```shell
go get -u github.com/gogf/gf/contrib/drivers/mysql/v2
```

这是用于通过 Go 语言的包管理工具 `go get` 来安装 `gf` 框架的 MySQL 驱动。命令中的 `-u` 参数表示更新包到最新版本。

// md5:d1c2b69639894015
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

# Supported Drivers

## // md5:993f339070bfdb48## # SQLite

```go
导入_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
```

这段代码是Go语言的导入语句，它引入了`github.com/gogf/gf/contrib/drivers/sqlite/v2`包。`_`是一个空白标识符，用于仅导入包以便其在代码中可用但不使用其中的任何函数或类型。这个包提供SQLite（一个轻量级的关系型数据库）的驱动支持，用于GF（Golang Foundation）框架。

# SQLite

```go
导入_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
```

这段代码是Go语言的导入语句，它引入了`github.com/gogf/gf/contrib/drivers/sqlite/v2`包。`_`是一个空白标识符，用于仅导入包以便其在代码中可用但不使用其中的任何函数或类型。这个包提供SQLite（一个轻量级的关系型数据库）的驱动支持，用于GF（Golang Foundation）框架。

// md5:d096a47263dc3bdf

```go
import _ "github.com/gogf/gf/contrib/drivers/sqlitecgo/v2"
```

## # PostgreSQL

```go
导入 _ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
```

注意：

- 不支持`Replace`功能。

# PostgreSQL

```go
导入 _ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
```

注意：

- 不支持`Replace`功能。

// md5:3b1d1e51b892a4d5

```go
import _ "github.com/gogf/gf/contrib/drivers/mssql/v2"
```

Note:

- It does not support `Replace` features.
- It does not support `LastInsertId`.
- It supports server version >= `SQL Server2005`
- It ONLY supports datetime2 and datetimeoffset types for auto handling created_at/updated_at/deleted_at columns, because datetime type does not support microseconds precision when column value is passed as string.

## # Oracle

```go
import _ "github.com/gogf/gf/contrib/drivers/oracle/v2"
```

注意：

- 它不支持`Replace`特性。
- 它不支持`LastInsertId`。

# Oracle

```go
import _ "github.com/gogf/gf/contrib/drivers/oracle/v2"
```

注意：

- 它不支持`Replace`特性。
- 它不支持`LastInsertId`。

// md5:9442ab8df9b797ce

```go
import _ "github.com/gogf/gf/contrib/drivers/clickhouse/v2"
```

Note:

- It does not support `InsertIgnore/InsertGetId` features.
- It does not support `Save/Replace` features.
- It does not support `Transaction` feature.
- It does not support `RowsAffected` feature.

## # 这是导入 DM 驱动的 Go 代码片段：

```go
import _ "github.com/gogf/gf/contrib/drivers/dm/v2"
```

注意：
- 它不支持 `Replace` 功能。

# 这是导入 DM 驱动的 Go 代码片段：

```go
import _ "github.com/gogf/gf/contrib/drivers/dm/v2"
```

注意：
- 它不支持 `Replace` 功能。

// md5:1d252b2894a4d8ce

It's quick and easy, please refer to current driver source.
It's quite appreciated if any PR for new drivers support into current repo.
