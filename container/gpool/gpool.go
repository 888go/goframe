// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包gpool提供了对象可重用的并发安全池。 md5:d111530cd572ede7
package gpool

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/container/glist"
	"github.com/gogf/gf/v2/container/gtype"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/os/gtimer"
)

// Pool是一个对象可重用池。 md5:08c256ba80594199
type Pool struct {
	list    *glist.List                 // 可用/闲置项目列表。 md5:f93c8d617cafe97f
	closed  *gtype.Bool                 // 是否关闭了连接池。 md5:73ea5526318af92f
	TTL     time.Duration               // 对象池中项目的生存时间。 md5:d9c944077d869281
	NewFunc func() (interface{}, error) // 创建池项的回调函数。 md5:f37bfc92a2188739
	// ExpireFunc 是用于过期项目销毁的函数。
	// 当池中的项目需要执行额外销毁操作时，需要定义这个函数。
	// 例如：net.Conn、os.File 等。
	// md5:f09911de2780aeaa
	ExpireFunc func(interface{})
}

// Pool item.
type poolItem struct {
	value    interface{} // Item value.
	expireAt int64       // 过期时间戳，单位为毫秒。 md5:d7096ed51593fa59
}

// NewFunc 对象的创建函数。 md5:245f622ac151f3ff
type NewFunc func() (interface{}, error)

// ExpireFunc 对象的销毁函数。 md5:cd0e3912eae30a98
type ExpireFunc func(interface{})

// New 创建并返回一个新的对象池。
// 为了确保执行效率，一旦设置，过期时间将不能修改。
// 
// 注意过期逻辑：
// ttl = 0：未过期；
// ttl < 0：使用后立即过期；
// ttl > 0：超时过期。
// md5:9f724382dd2313e7
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

// Put 将一个项目放入池中。 md5:d7b57780f7e8f1cc
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
		// 对于Golang版本小于1.13的，time.Duration没有Milliseconds方法。
		// 因此我们需要使用其纳秒值来计算毫秒数。
		// md5:87b516a9573fac98
		item.expireAt = gtime.TimestampMilli() + p.TTL.Nanoseconds()/1000000
	}
	p.list.PushBack(item)
	return nil
}

// MustPut 将一个项目放入池中，如果发生任何错误，它将引发恐慌。 md5:10206f4587a99039
func (p *Pool) MustPut(value interface{}) {
	if err := p.Put(value); err != nil {
		panic(err)
	}
}

// Clear 清空池子，这意味着它将从池中移除所有项目。 md5:c141b6e6c215bc68
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

// Get 从池中选取并返回一个项目。如果池是空的并且已定义了NewFunc，
// 则会使用NewFunc创建并返回一个项目。
// md5:7782b49d380b807b
func (p *Pool) Get() (interface{}, error) {
	for !p.closed.Val() {
		if r := p.list.PopFront(); r != nil {
			f := r.(*poolItem)
			if f.expireAt == 0 || f.expireAt > gtime.TimestampMilli() {
				return f.value, nil
			} else if p.ExpireFunc != nil {
				// 待办事项：将过期功能调用异步移出`Get`操作。 md5:13d59efb4d92da03
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

// Size 返回池中可用项目的数量。 md5:2b8a683e177e1586
func (p *Pool) Size() int {
	return p.list.Len()
}

// Close 关闭池。如果 `p` 有 ExpireFunc，那么在关闭之前，它会自动使用这个函数关闭所有项目。通常情况下，你不需要手动调用这个函数。
// md5:368c18d44115f9cc
func (p *Pool) Close() {
	p.closed.Set(true)
}

// checkExpire 每秒从池中移除过期的项目。 md5:1177ab8b3e8a341e
func (p *Pool) checkExpireItems(ctx context.Context) {
	if p.closed.Val() {
		// 如果p具有ExpireFunc，
		// 则必须使用此函数关闭所有项。
		// md5:8ec38193671c9632
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
	// 所有项目都不会过期。 md5:9c3b9311c20c9c20
	if p.TTL == 0 {
		return
	}
	// 最新的项目过期时间（以毫秒为单位）。 md5:46946a9b5c1228ca
	var latestExpire int64 = -1
	// 获取当前时间戳（毫秒），使用这个时间戳来判断项目是否过期。
	// 对于每个项目过期的判断不是非常精确，但性能较高。
	// md5:5dc686eec927131e
	var timestampMilli = gtime.TimestampMilli()
	for {
		if latestExpire > timestampMilli {
			break
		}
		if r := p.list.PopFront(); r != nil {
			item := r.(*poolItem)
			latestExpire = item.expireAt
			// TODO 优化池的自动过期机制。 md5:b4e2c483478d7737
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
