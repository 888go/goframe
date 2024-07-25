// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gdb

import (
	"context"
	"sync"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gctx"
)

// internalCtxData 为内部使用目的，在 ctx 中存储数据。 md5:95073898cc1f4772
type internalCtxData struct {
	sync.Mutex
		// 当前操作中使用的配置节点。 md5:85f106587581bb38
	ConfigNode *ConfigNode
}

// column 用于内部目的，在ctx中存储列数据。 md5:12a8a80132bf8ae7
type internalColumnData struct {
	// 来自数据库服务器的响应结果中的第一列。
	// 此属性用于值/计数选择语句的目的，以避免可能修改结果列的HOOK处理器，这可能会混淆值/计数选择语句的逻辑。
	// md5:c678f20e25487136
	FirstResultColumn string
}

const (
	internalCtxDataKeyInCtx    gctx.StrKey = "InternalCtxData"
	internalColumnDataKeyInCtx gctx.StrKey = "InternalColumnData"

	// `ignoreResultKeyInCtx` 是为了一些不支持 `RowsAffected` 函数的数据库驱动（例如：`clickhouse`）设置的标记。`clickhouse` 不支持获取插入/更新的结果，但在执行 `RowsAffected` 时会返回错误。在这里，我们忽略对 `RowsAffected` 的调用，以避免触发错误，而不是在错误发生后忽略它们。
	// md5:4a7864c37326a119
	ignoreResultKeyInCtx gctx.StrKey = "IgnoreResult"
)

func (c *Core) injectInternalCtxData(ctx context.Context) context.Context {
		// 如果内部数据已经被注入，则不做任何操作。 md5:ae258e1c66cb106a
	if ctx.Value(internalCtxDataKeyInCtx) != nil {
		return ctx
	}
	return context.WithValue(ctx, internalCtxDataKeyInCtx, &internalCtxData{
		ConfigNode: c.config,
	})
}

func (c *Core) setConfigNodeToCtx(ctx context.Context, node *ConfigNode) error {
	value := ctx.Value(internalCtxDataKeyInCtx)
	if value == nil {
		return gerror.NewCode(gcode.CodeInternalError, `no internal data found in context`)
	}

	data := value.(*internalCtxData)
	data.Lock()
	defer data.Unlock()
	data.ConfigNode = node
	return nil
}

func (c *Core) getConfigNodeFromCtx(ctx context.Context) *ConfigNode {
	if value := ctx.Value(internalCtxDataKeyInCtx); value != nil {
		data := value.(*internalCtxData)
		data.Lock()
		defer data.Unlock()
		return data.ConfigNode
	}
	return nil
}

func (c *Core) injectInternalColumn(ctx context.Context) context.Context {
	return context.WithValue(ctx, internalColumnDataKeyInCtx, &internalColumnData{})
}

func (c *Core) getInternalColumnFromCtx(ctx context.Context) *internalColumnData {
	if v := ctx.Value(internalColumnDataKeyInCtx); v != nil {
		return v.(*internalColumnData)
	}
	return nil
}

func (c *Core) InjectIgnoreResult(ctx context.Context) context.Context {
	if ctx.Value(ignoreResultKeyInCtx) != nil {
		return ctx
	}
	return context.WithValue(ctx, ignoreResultKeyInCtx, true)
}

func (c *Core) GetIgnoreResultFromCtx(ctx context.Context) bool {
	return ctx.Value(ignoreResultKeyInCtx) != nil
}
