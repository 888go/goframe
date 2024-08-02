// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package dm

import (
	"database/sql"
	"fmt"

	"net/url"
	"strings"

	gdb "github.com/888go/goframe/database/gdb"
	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
)

// Open 创建并返回一个底层的 sql.DB 对象，用于 pgsql。 md5:a5d566f750df5890
func (d *Driver) Open(config *gdb.ConfigNode) (db *sql.DB, err error) {
	var (
		source               string
		underlyingDriverName = "dm"
	)
	if config.Name == "" {
		return nil, fmt.Errorf(
			`dm.Open failed for driver "%s" without DB Name`, underlyingDriverName,
		)
	}
	// DM8 的数据源名称:
	// dm:	//用户名:密码@IP地址:端口号/数据库名
	// dm:	//用户名:密码@DW/数据库名?DW=(192.168.1.1:5236,192.168.1.2:5236)
	// 
	// 这里的注释解释了如何连接到DM8数据库。`dm:	//` 是驱动的URL格式，`userName` 和 `password` 是用于身份验证的用户名和密码，`ip:port` 是数据库服务器的地址和端口，`dbname` 是要连接的数据库名。在第二个示例中，`DW` 可能是分布式仓库的标识，后面跟着一个列表 `(192.168.1.1:5236,192.168.1.2:5236)`，表示数据库集群的多个节点。
	// md5:15e6de4613ebd611
	var domain string
	if config.Port != "" {
		domain = fmt.Sprintf("%s:%s", config.Host, config.Port)
	} else {
		domain = config.Host
	}
	source = fmt.Sprintf(
		"dm://%s:%s@%s/%s?charset=%s&schema=%s",
		config.User, config.Pass, domain, config.Name, config.Charset, config.Name,
	)
	// 时区设置演示：
	// &loc=亚洲/上海
	// md5:dcdf1db7d830174b
	if config.Timezone != "" {
		if strings.Contains(config.Timezone, "/") {
			config.Timezone = url.QueryEscape(config.Timezone)
		}
		source = fmt.Sprintf("%s&loc%s", source, config.Timezone)
	}
	if config.Extra != "" {
		source = fmt.Sprintf("%s&%s", source, config.Extra)
	}

	if db, err = sql.Open(underlyingDriverName, source); err != nil {
		err = gerror.WrapCodef(
			gcode.CodeDbOperationError, err,
			`dm.Open failed for driver "%s" by source "%s"`, underlyingDriverName, source,
		)
		return nil, err
	}
	return
}
