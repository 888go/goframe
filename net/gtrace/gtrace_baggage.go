// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gtrace

import (
	"context"
	
	"go.opentelemetry.io/otel/baggage"
	
	"github.com/888go/goframe/container/gmap"
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/util/gconv"
)

// Baggage 在所有追踪跨度中保存数据。
type Baggage struct {
	ctx context.Context
}

// NewBaggage 从给定的追踪上下文中创建并返回一个新的 Baggage 对象。
func NewBaggage(ctx context.Context) *Baggage {
	if ctx == nil {
		ctx = context.Background()
	}
	return &Baggage{
		ctx: ctx,
	}
}

// Ctx 返回 Baggage 持有的上下文。
func (b *Baggage) Ctx() context.Context {
	return b.ctx
}

// SetValue 是一个方便的函数，用于向 baggage 中添加一对键值对。
// 注意，它使用 attribute.Any 来设置键值对。
func (b *Baggage) SetValue(key string, value interface{}) context.Context {
	member, _ := baggage.NewMember(key, gconv.String(value))
	bag, _ := baggage.New(member)
	b.ctx = baggage.ContextWithBaggage(b.ctx, bag)
	return b.ctx
}

// SetMap 是一个方便的函数，用于向 baggage 添加映射键值对。
// 注意，它使用 attribute.Any 来设置键值对。
func (b *Baggage) SetMap(data map[string]interface{}) context.Context {
	members := make([]baggage.Member, 0)
	for k, v := range data {
		member, _ := baggage.NewMember(k, gconv.String(v))
		members = append(members, member)
	}
	bag, _ := baggage.New(members...)
	b.ctx = baggage.ContextWithBaggage(b.ctx, bag)
	return b.ctx
}

// GetMap 获取并以map形式返回 baggage 的值。
func (b *Baggage) GetMap() *gmap.StrAnyMap {
	m := gmap.NewStrAnyMap()
	members := baggage.FromContext(b.ctx).Members()
	for i := range members {
		m.Set(members[i].Key(), members[i].Value())
	}
	return m
}

// GetVar 从 baggage 中根据指定的键获取值，并返回一个指向该值的*gvar.Var指针。
func (b *Baggage) GetVar(key string) *gvar.Var {
	value := baggage.FromContext(b.ctx).Member(key).Value()
	return gvar.New(value)
}
