// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gredis

import (
	"context"

	"github.com/gogf/gf/v2/container/gvar"
)

// IGroupSortedSet 管理 redis 排序集合操作。
// 实现了 redis.GroupSortedSet。
// md5:85c86f571889c1f2
type IGroupSortedSet interface {
	ZAdd(ctx context.Context, key string, option *ZAddOption, member ZAddMember, members ...ZAddMember) (*gvar.Var, error)
	ZScore(ctx context.Context, key string, member interface{}) (float64, error)
	ZIncrBy(ctx context.Context, key string, increment float64, member interface{}) (float64, error)
	ZCard(ctx context.Context, key string) (int64, error)
	ZCount(ctx context.Context, key string, min, max string) (int64, error)
	ZRange(ctx context.Context, key string, start, stop int64, option ...ZRangeOption) (gvar.Vars, error)
	ZRevRange(ctx context.Context, key string, start, stop int64, option ...ZRevRangeOption) (*gvar.Var, error)
	ZRank(ctx context.Context, key string, member interface{}) (int64, error)
	ZRevRank(ctx context.Context, key string, member interface{}) (int64, error)
	ZRem(ctx context.Context, key string, member interface{}, members ...interface{}) (int64, error)
	ZRemRangeByRank(ctx context.Context, key string, start, stop int64) (int64, error)
	ZRemRangeByScore(ctx context.Context, key string, min, max string) (int64, error)
	ZRemRangeByLex(ctx context.Context, key string, min, max string) (int64, error)
	ZLexCount(ctx context.Context, key, min, max string) (int64, error)
}

// ZAddOption 为 ZAdd 函数提供选项。 md5:b3e234b14d4a1ca8
type ZAddOption struct {
	XX bool // Only update elements that already exist. Don't add new elements.
	NX bool // Only add new elements. Don't update already existing elements.
// 只有当新分数小于当前分数时，才更新已存在的元素。
// 此标志不会阻止添加新元素。
// md5:df3556a5d410e3c9
	LT bool

// 只有当新分数大于当前分数时，才更新现有元素。此标志不会阻止添加新元素。
// md5:4866b5e44d3c1bec
	GT bool

// 将返回值从新添加元素的数量修改为更改的总元素数量（CH代表已更改）。 
// 已更改的元素包括新添加的元素和分数已被更新的现有元素。 
// 因此，命令行中指定的与过去相同的分数的元素不计入总数。 
// 注意：通常情况下，ZAdd的返回值只计算新添加的元素数量。
// md5:f80865660e63c42c
	CH bool

	// 当指定了此选项时，ZAdd 命令的行为类似于 ZIncrBy。在这种模式下，只能指定一个分数-元素对。 md5:bb002fb3eec4eb13
	INCR bool
}

// ZAddMember 是集合（set）中的元素结构体。 md5:eb7d172c444324d7
type ZAddMember struct {
	Score  float64
	Member interface{}
}

// ZRangeOption为ZRange函数提供额外选项。 md5:61532d16fe5a1260
type ZRangeOption struct {
	ByScore bool
	ByLex   bool
// 可选的REV参数会反转顺序，因此元素按照分数从高到低排序，分数相同时则通过反向字典序进行解析。
// md5:a1c79d75cedbff1b
	Rev   bool
	Limit *ZRangeOptionLimit
	// 可选的 WithScores 参数会在命令回复中补充返回元素的分数。 md5:26db0341550d511b
	WithScores bool
}

// ZRangeOptionLimit 为 ZRange 函数提供 LIMIT 参数。
// 可选的 LIMIT 参数可用于从匹配的元素中获取子范围（类似于 SQL 中的 SELECT LIMIT 偏移量, 数量）。当 `Count` 为负数时，从 `Offset` 开始返回所有元素。
// md5:a910bb82b51914ef
type ZRangeOptionLimit struct {
	Offset *int
	Count  *int
}

// ZRevRangeOption为ZRevRange函数提供了选项。 md5:cd0b627793d48f50
type ZRevRangeOption struct {
	WithScores bool
}
