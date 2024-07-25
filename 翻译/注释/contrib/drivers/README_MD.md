
<原文开始>
drivers

Powerful database drivers for package gdb.


<原文结束>

# <翻译开始>
// md5:34e63c84ac9f12dc# # 安装教程
# <翻译结束>


<原文开始>
Installation

Let's take `mysql` for example.

```shell
go get -u github.com/gogf/gf/contrib/drivers/mysql/v2

<原文结束>

# <翻译开始>
// md5:d1c2b69639894015# # 易于复制
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
	// md5:0bbd7dd7e96240ed# Supported Drivers
# <翻译结束>


<原文开始>
MySQL/MariaDB/TiDB

```go
import _ "github.com/gogf/gf/contrib/drivers/mysql/v2"
```


<原文结束>

# <翻译开始>
// md5:993f339070bfdb48## # SQLite## 
# <翻译结束>


<原文开始>
SQLite

```go
import _ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
```

#
<原文结束>

# <翻译开始>
// md5:d096a47263dc3bdf## 
# <翻译结束>


<原文开始>
cgo version for 32-bit windows

```go
import _ "github.com/gogf/gf/contrib/drivers/sqlitecgo/v2"
```


<原文结束>

# <翻译开始>
// md5:b5c89ec6af200159## 
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
// md5:3b1d1e51b892a4d5## 
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
// md5:ae572708ba108995## 
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
// md5:9442ab8df9b797ce## 
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
// md5:b60f4ceb35b1e184## 
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
// md5:1d252b2894a4d8ce# 
# <翻译结束>


<原文开始>
Custom Drivers

It's quick and easy, please refer to current driver source.
It's quite appreciated if any PR for new drivers support into current repo.

<原文结束>

# <翻译开始>
// md5:a20a02aff96de6a1
# <翻译结束>

