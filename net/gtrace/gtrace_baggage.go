// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gtrace

import (
	"context"

	"go.opentelemetry.io/otel/baggage"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/util/gconv"
)

// Baggage在所有跟踪span中持有数据。. md5:0ad27152ec042f81
type Baggage struct {
	ctx context.Context
}

// NewBaggage 从给定的追踪上下文中创建并返回一个新的Baggage对象。. md5:6c3e8093bd06a60a
func NewBaggage(ctx context.Context) *Baggage {
	if ctx == nil {
		ctx = context.Background()
	}
	return &Baggage{
		ctx: ctx,
	}
}

// Ctx 返回Baggage持有的上下文。. md5:37268f528c617799
func (b *Baggage) Ctx() context.Context {
	return b.ctx
}

// SetValue 是一个方便的函数，用于向 baggage 中添加一个键值对。
// 注意，它使用 attribute.Any 设置键值对。
// md5:830faae9a81721ce
func (b *Baggage) SetValue(key string, value interface{}) context.Context {
	member, _ := baggage.NewMember(key, gconv.String(value))
	bag, _ := baggage.New(member)
	b.ctx = baggage.ContextWithBaggage(b.ctx, bag)
	return b.ctx
}

// SetMap 是一个方便的函数，用于将映射键值对添加到行李中。
// 注意，它使用 attribute.Any 设置键值对。
// md5:a18951801562457c
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

// GetMap 获取并以映射形式返回baggage值。. md5:d6024d765655a29e
func (b *Baggage) GetMap() *gmap.StrAnyMap {
	m := gmap.NewStrAnyMap()
	members := baggage.FromContext(b.ctx).Members()
	for i := range members {
		m.Set(members[i].Key(), members[i].Value())
	}
	return m
}

// GetVar 从行李中获取指定键的值，并返回一个 *gvar.Var。. md5:6cda7fcfb8ff1c6e
func (b *Baggage) GetVar(key string) *gvar.Var {
	value := baggage.FromContext(b.ctx).Member(key).Value()
	return gvar.New(value)
}
