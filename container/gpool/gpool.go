// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 包gpool提供了一个对象可重用的并发安全池。
package gpool

import (
	"context"
	"time"
	
	"github.com/888go/goframe/container/glist"
	"github.com/888go/goframe/container/gtype"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/os/gtimer"
)

// Pool 是一个对象可重用池。
type Pool struct {
	list    *glist.List                 // 可用/空闲项目列表。
	closed  *gtype.Bool                 // 是否池已关闭。
	TTL     time.Duration               // Time To Live for pool items. // 池中项目（对象）的生存时间。
	NewFunc func() (interface{}, error) // 回调函数，用于创建池中的项目。
// ExpireFunc 是用于过期项销毁的函数。
// 当池中的项需要执行额外的销毁操作时，需要定义此函数。
// 例如：net.Conn（网络连接），os.File（操作系统文件）等。
	ExpireFunc func(interface{})
}

// Pool item.
type poolItem struct {
	value    interface{} // Item value.
	expireAt int64       // 过期时间戳（毫秒）
}

// NewFunc 创建对象的构造函数
type NewFunc func() (interface{}, error)

// ExpireFunc 对象的销毁函数。
type ExpireFunc func(interface{})

// New 创建并返回一个新的对象池。
// 为了确保执行效率，一旦设置过期时间则不可再修改。
//
// 注意过期逻辑：
// ttl = 0 : 不过期；
// ttl < 0 : 使用后立即过期；
// ttl > 0 : 超时后过期；
func New(ttl time.Duration, newFunc NewFunc, expireFunc ...ExpireFunc) *Pool {
	r := &Pool{
		list:    glist.New(true),
		closed:  gtype.NewBool(),
		TTL:     ttl,
		NewFunc: newFunc,
	}
	if len(expireFunc) > 0 {
		r.ExpireFunc = expireFunc[0]
	}
	gtimer.AddSingleton(context.Background(), time.Second, r.checkExpireItems)
	return r
}

// Put 将一个项目放入池中。
func (p *Pool) Put(value interface{}) error {
	if p.closed.Val() {
		return gerror.NewCode(gcode.CodeInvalidOperation, "pool is closed")
	}
	item := &poolItem{
		value: value,
	}
	if p.TTL == 0 {
		item.expireAt = 0
	} else {
// 对于 Go 语言版本小于 1.13 的情况，time.Duration 类型没有内置的 Milliseconds 方法。
// 因此我们需要通过其纳秒值来计算毫秒值。
		item.expireAt = gtime.TimestampMilli() + p.TTL.Nanoseconds()/1000000
	}
	p.list.PushBack(item)
	return nil
}

// MustPut 将一个项目放入池中，如果发生任何错误，它会引发panic。
func (p *Pool) MustPut(value interface{}) {
	if err := p.Put(value); err != nil {
		panic(err)
	}
}

// Clear 清空 pool，这意味着它会从 pool 中移除所有项目。
func (p *Pool) Clear() {
	if p.ExpireFunc != nil {
		for {
			if r := p.list.PopFront(); r != nil {
				p.ExpireFunc(r.(*poolItem).value)
			} else {
				break
			}
		}
	} else {
		p.list.RemoveAll()
	}
}

// Get 从池中获取并返回一个项目。如果池为空且定义了 NewFunc，
// 则会通过 NewFunc 创建并返回一个新的项目。
func (p *Pool) Get() (interface{}, error) {
	for !p.closed.Val() {
		if r := p.list.PopFront(); r != nil {
			f := r.(*poolItem)
			if f.expireAt == 0 || f.expireAt > gtime.TimestampMilli() {
				return f.value, nil
			} else if p.ExpireFunc != nil {
				// TODO: 将过期函数调用异步移出 `Get` 操作。
				p.ExpireFunc(f.value)
			}
		} else {
			break
		}
	}
	if p.NewFunc != nil {
		return p.NewFunc()
	}
	return nil, gerror.NewCode(gcode.CodeInvalidOperation, "pool is empty")
}

// Size 返回 pool 中可用项目的数量。
func (p *Pool) Size() int {
	return p.list.Len()
}

// Close 关闭连接池。如果 `p` 拥有 ExpireFunc，
// 则在关闭前会自动使用该函数关闭所有项目。
// 通常情况下，你不需要手动调用这个函数。
func (p *Pool) Close() {
	p.closed.Set(true)
}

// checkExpire 每隔一秒从池中移除已过期的项目。
func (p *Pool) checkExpireItems(ctx context.Context) {
	if p.closed.Val() {
// 如果p拥有ExpireFunc，
// 那么它必须使用这个函数关闭所有项。
		if p.ExpireFunc != nil {
			for {
				if r := p.list.PopFront(); r != nil {
					p.ExpireFunc(r.(*poolItem).value)
				} else {
					break
				}
			}
		}
		gtimer.Exit()
	}
	// 所有项目永不过期。
	if p.TTL == 0 {
		return
	}
	// 最近一项过期时间戳（以毫秒为单位）
	var latestExpire int64 = -1
// 获取当前时间戳（毫秒级），通过与此时间戳进行比较来决定缓存项是否过期。这种方法并非对每一个缓存项的过期判断都精确，但具有高性能的特点。
	var timestampMilli = gtime.TimestampMilli()
	for {
		if latestExpire > timestampMilli {
			break
		}
		if r := p.list.PopFront(); r != nil {
			item := r.(*poolItem)
			latestExpire = item.expireAt
			// TODO: 改进池的自动过期机制。
			if item.expireAt > timestampMilli {
				p.list.PushFront(item)
				break
			}
			if p.ExpireFunc != nil {
				p.ExpireFunc(item.value)
			}
		} else {
			break
		}
	}
}
