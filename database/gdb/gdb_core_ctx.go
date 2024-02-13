// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package db类

import (
	"context"
	
	"github.com/888go/goframe/os/gctx"
)

// internalCtxData 用于存储内部使用目的的数据到 ctx。
type internalCtxData struct {
	// Operation DB.
	DB DB

	// 当前操作中使用的配置节点。
	ConfigNode *X配置项

// 此处为从数据库服务器获取的结果响应中的第一列。
// 该属性用于值/计数选择语句的目的，
// 其目的是避免可能修改结果列的HOOK处理器，
// 这样做是为了防止混淆值/计数选择语句逻辑。
	FirstResultColumn string
}

const (
	internalCtxDataKeyInCtx 上下文类.StrKey = "InternalCtxData"

// `ignoreResultKeyInCtx` 是一个标志，用于某些不支持 `RowsAffected` 函数的数据库驱动程序，
// 例如：`clickhouse`。ClickHouse 在执行插入/更新操作时不支持获取结果，
// 而是在调用 `RowsAffected` 时返回错误。在这里忽略对 `RowsAffected` 的调用，
// 以避免触发错误，而不是在错误触发后忽略它们。
	ignoreResultKeyInCtx 上下文类.StrKey = "IgnoreResult"
)

func (c *Core) 底层_InjectInternalCtxData(ctx context.Context) context.Context {
	// 如果内部数据已经注入，则不做任何操作。
	if ctx.Value(internalCtxDataKeyInCtx) != nil {
		return ctx
	}
	return context.WithValue(ctx, internalCtxDataKeyInCtx, &internalCtxData{
		DB:         c.db,
		ConfigNode: c.config,
	})
}

func (c *Core) 底层_GetInternalCtxDataFromCtx(ctx context.Context) *internalCtxData {
	if v := ctx.Value(internalCtxDataKeyInCtx); v != nil {
		return v.(*internalCtxData)
	}
	return nil
}

func (c *Core) 底层_InjectIgnoreResult(ctx context.Context) context.Context {
	if ctx.Value(ignoreResultKeyInCtx) != nil {
		return ctx
	}
	return context.WithValue(ctx, ignoreResultKeyInCtx, true)
}

func (c *Core) 底层_GetIgnoreResultFromCtx(ctx context.Context) bool {
	return ctx.Value(ignoreResultKeyInCtx) != nil
}
