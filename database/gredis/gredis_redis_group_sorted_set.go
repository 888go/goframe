// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gredis
import (
	"context"
	
	"github.com/888go/goframe/container/gvar"
	)
// IGroupSortedSet 管理 Redis 有序集合操作。
// 实现参考 redis.GroupSortedSet。
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

// ZAddOption 提供了函数 ZAdd 的选项。
type ZAddOption struct {
	XX bool // Only update elements that already exist. Don't add new elements.
	NX bool // Only add new elements. Don't update already existing elements.
// 只有当新分数小于当前分数时，才更新已存在的元素。
// 但请注意，此标志不会阻止添加新的元素。
	LT bool

// 如果新分数大于当前分数，则仅更新现有元素。
// 此标志不会阻止添加新元素。
	GT bool

// 将返回值由新增元素的数量修改为更改过的总元素数量（CH 是“changed”的缩写）。
// 更改过的元素包括新添加的元素以及已存在但分数被更新的元素。
// 因此，命令行中指定且其分数与过去相同的元素不会被计算在内。
// 注意：通常情况下，ZAdd 的返回值仅计算新增元素的数量。
	CH bool

	// 当指定了此选项时，ZAdd 表现得如同 ZIncrBy。在这种模式下，只能指定一个分数-元素对。
	INCR bool
}

// ZAddMember 是集合中元素的结构体。
type ZAddMember struct {
	Score  float64
	Member interface{}
}

// ZRangeOption 为 ZRange 函数提供了额外的选项。
type ZRangeOption struct {
	ByScore bool
	ByLex   bool
// 可选参数 REV 用于反转排序顺序，因此元素按从高到低的分数进行排序，
// 当分数相同时，采用反字典序进行排序结果的确定。
	Rev   bool
	Limit *ZRangeOptionLimit
	// 可选的 WithScores 参数会用返回元素的分数来补充命令的回复。
	WithScores bool
}

// ZRangeOptionLimit 为 ZRange 函数提供 LIMIT 参数。
// 可选的 LIMIT 参数可用于从匹配元素中获取一个子范围（类似于 SQL 中的 SELECT LIMIT offset, count）。
// 如果 `Count` 为负数，则返回从 `Offset` 开始的所有元素。
type ZRangeOptionLimit struct {
	Offset *int
	Count  *int
}

// ZRevRangeOption 提供了 ZRevRange 函数的选项。
type ZRevRangeOption struct {
	WithScores bool
}
