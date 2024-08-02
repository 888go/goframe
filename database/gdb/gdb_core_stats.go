// 版权所有 GoFrame 作者(https://goframe.org)。保留所有权利。
//
// 此源代码形式受 MIT 许可证的条款约束。
// 如果未随此文件一起分发 MIT 许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
// md5:a114f4bdd106ab31

package db类

import (
	"context"
	"database/sql"
)

type localStatsItem struct {
	node  *ConfigNode
	stats sql.DBStats
}

// Node 返回配置节点信息。 md5:868005c0df3fa483
func (item *localStatsItem) Node() ConfigNode {
	return *item.node
}

// Stats 返回当前节点的连接状态统计信息。 md5:b497e68c5fce778b
func (item *localStatsItem) Stats() sql.DBStats {
	return item.stats
}

// Stats获取并返回所有已建立节点的池统计信息。 md5:83a29e795d6705a8
func (c *Core) Stats(ctx context.Context) []StatsItem {
	var items = make([]StatsItem, 0)
	c.links.Iterator(func(k, v any) bool {
		var (
			node  = k.(ConfigNode)
			sqlDB = v.(*sql.DB)
		)
		items = append(items, &localStatsItem{
			node:  &node,
			stats: sqlDB.Stats(),
		})
		return true
	})
	return items
}
