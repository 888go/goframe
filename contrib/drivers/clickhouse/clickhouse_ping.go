// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package clickhouse

import (
	"database/sql"
	"fmt"

	"github.com/ClickHouse/clickhouse-go/v2"
)

// PingMaster 向主节点发送请求以检查身份验证或保持连接活动。 md5:47a7df55cbee8583
func (d *Driver) PingMaster() error {
	conn, err := d.Master()
	if err != nil {
		return err
	}
	return d.ping(conn)
}

// PingSlave 调用ping命令检查从节点的认证或者维持连接。 md5:62272b38d874eda6
func (d *Driver) PingSlave() error {
	conn, err := d.Slave()
	if err != nil {
		return err
	}
	return d.ping(conn)
}

// ping 返回Clickhouse特有的错误。 md5:1cf9d2a1e2ac7219
func (d *Driver) ping(conn *sql.DB) error {
	err := conn.Ping()
	if exception, ok := err.(*clickhouse.Exception); ok {
		return fmt.Errorf("[%d]%s", exception.Code, exception.Message)
	}
	return err
}
