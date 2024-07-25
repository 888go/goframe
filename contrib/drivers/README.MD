# // md5:34e63c84ac9f12dc# # 安装教程# // md5:d1c2b69639894015# # 易于复制
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
 md5:0bbd7dd7e96240ed# Supported Drivers

## // md5:993f339070bfdb48## # SQLite## // md5:d096a47263dc3bdf## // md5:b5c89ec6af200159## // md5:3b1d1e51b892a4d5## // md5:ae572708ba108995## // md5:9442ab8df9b797ce## // md5:b60f4ceb35b1e184## // md5:1d252b2894a4d8ce# // md5:a20a02aff96de6a1