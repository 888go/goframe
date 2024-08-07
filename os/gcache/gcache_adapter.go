// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 缓存类

import (
	"context"
	"time"

	gvar "github.com/888go/goframe/container/gvar"
)

// Adapter是缓存功能的核心适配器。
// 
// 注意，实现者本身应确保这些函数的并发安全性。
// md5:cd91041442c2fdbf
type Adapter interface {
	// X设置值 使用 `key`-`value` 对设置缓存，该缓存在 `duration` 时间后过期。
	//
	// 如果 `duration` == 0，则不设置过期时间。
	// 如果 `duration` < 0 或给定的 `value` 为 nil，则删除 `data` 的键。
	// md5:3f5918d3cc5c36fd
	X设置值(ctx context.Context, key interface{}, value interface{}, duration time.Duration) error

	// X设置Map 批量设置缓存，使用 `data` 映射中的键值对，这些缓存在 `duration` 时间后过期。
	//
	// 如果 `duration` == 0，则不会过期。
	// 如果 `duration` < 0 或者给定的 `value` 为 nil，将删除 `data` 中的键。
	// md5:029757e42001dd48
	X设置Map(ctx context.Context, data map[interface{}]interface{}, duration time.Duration) error

	// X设置值并跳过已存在 如果缓存中不存在`key`，则设置过期时间为`duration`的`key`-`value`对。如果成功将`value`设置到缓存中，它会返回`true`，表示`key`在缓存中不存在；否则返回`false`。
	// 
	// 如果`duration`为0，缓存不会过期。
	// 如果`duration`小于0或给定的`value`为`nil`，它会删除`key`。
	// md5:a442e240e2ddb849
	X设置值并跳过已存在(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (ok bool, err error)

	// X设置值并跳过已存在_函数 如果缓存中不存在`key`，则使用函数`f`的结果设置`key`，并返回true。如果`key`已存在，则不做任何操作，返回false。
	//
	// 参数`value`可以是`func() interface{}`类型，但如果其结果为nil，则不执行任何操作。
	//
	// 如果`duration`为0，表示永不过期。如果`duration`小于0或给定的`value`为nil，则删除`key`。
	// md5:33f0e2bb534c4ac4
	X设置值并跳过已存在_函数(ctx context.Context, key interface{}, f Func, duration time.Duration) (ok bool, err error)

	// X设置值并跳过已存在_并发安全函数 如果`key`在缓存中不存在，则使用函数`f`的结果设置`key`并返回true。
	// 如果`key`已存在，则不做任何操作并返回false。
	//
	// 如果`duration`为0，则不设置过期时间。
	// 如果`duration`小于0或给定的`value`为nil，则删除`key`。
	//
	// 注意，它与函数`SetIfNotExistFunc`的不同之处在于，为了并发安全，函数`f`在写入互斥锁内部执行。
	// md5:906879fb08827346
	X设置值并跳过已存在_并发安全函数(ctx context.Context, key interface{}, f Func, duration time.Duration) (ok bool, err error)

	// X取值 获取并返回给定`key`关联的值。
	// 如果键不存在、其值为nil或已过期，它将返回nil。
	// 如果你想检查`key`是否在缓存中存在，最好使用Contains函数。
	// md5:a04abebd42f9db26
	X取值(ctx context.Context, key interface{}) (*gvar.Var, error)

	// X取值或设置值 从缓存中获取并返回`key`的值，如果`key`不存在，则设置`key-value`对，并返回`value`。缓存中的键值对在`duration`后过期。
	// 如果`duration`为0，则不会过期。
	// 如果`duration`小于0或给定的`value`为nil，它会删除`key`；但如果`value`是一个函数且函数结果为nil，则不做任何操作。
	// md5:a9525aacd8a5324e
	X取值或设置值(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (result *gvar.Var, err error)

	// X取值或设置值_函数 从缓存中获取并返回键`key`的值，如果键不存在，则使用函数`f`的结果设置键并返回该结果。键值对在`duration`后过期。
	// 
	// 如果`duration`为0，表示永不过期。
	// 如果`duration`小于0或给定的`value`为nil，它会删除键`key`。但如果`value`是一个函数并且函数结果为nil，它不会做任何操作。
	// md5:57a987bd75623802
	X取值或设置值_函数(ctx context.Context, key interface{}, f Func, duration time.Duration) (result *gvar.Var, err error)

	// X取值或设置值_并发安全函数 获取并返回`key`的值，如果`key`在缓存中不存在，则使用函数`f`的结果设置`key`并返回该结果。
	// 键值对将在`duration`时间后过期。
	//
	// 如果`duration`为0，则不设置过期时间。
	// 如果`duration`小于0或给定的`value`为nil，它将删除`key`，但若`value`是一个函数且函数结果为nil时，它不做任何操作。
	//
	// 需要注意的是，此函数与`GetOrSetFunc`的区别在于，函数`f`是在写入互斥锁内部执行的，以确保并发安全。
	// md5:b0a08f256bf6fcfc
	X取值或设置值_并发安全函数(ctx context.Context, key interface{}, f Func, duration time.Duration) (result *gvar.Var, err error)

		// X是否存在 检查并返回如果 `key` 在缓存中存在则为真，否则为假。 md5:4ff234995709b9ab
	X是否存在(ctx context.Context, key interface{}) (bool, error)

		// X取数量 返回缓存中的项目数量。 md5:2122f80de9340261
	X取数量(ctx context.Context) (size int, err error)

	// X取所有键值Map副本返回缓存中所有键值对的副本，以map类型。
	// 注意，此函数可能会导致大量内存使用。如果需要，您可以实现这个函数。
	// md5:96cf9c57d77ba2dd
	X取所有键值Map副本(ctx context.Context) (data map[interface{}]interface{}, err error)

		// X取所有键 返回缓存中所有键的切片。 md5:7ebd9dba01282dc2
	X取所有键(ctx context.Context) (keys []interface{}, err error)

		// X取所有值 返回缓存中所有的值作为切片。 md5:dc00b32eb8913e9b
	X取所有值(ctx context.Context) (values []interface{}, err error)

	// X更新值 更新`key`的值，不改变其过期时间，并返回旧的值。
	// 如果`key`在缓存中不存在，返回的值`exist`为false。
	//
	// 如果给定的`value`为nil，它会删除`key`。
	// 如果`key`不在缓存中，它不会做任何操作。
	// md5:28635aef7c0fc7a9
	X更新值(ctx context.Context, key interface{}, value interface{}) (oldValue *gvar.Var, exist bool, err error)

	// X更新过期时间 更新键`key`的过期时间，并返回旧的过期持续时间值。
	//
	// 如果`key`在缓存中不存在，它将返回-1并什么都不做。如果`duration`小于0，它会删除`key`。
	// md5:f1bb94e5134bebed
	X更新过期时间(ctx context.Context, key interface{}, duration time.Duration) (oldDuration time.Duration, err error)

	// X取过期时间 获取并返回缓存中 `key` 的过期时间。
	//
	// 注意：
	// 如果 `key` 没有设置过期时间，它将返回 0。
	// 如果 `key` 在缓存中不存在，它将返回 -1。
	// md5:6a059254c0534a31
	X取过期时间(ctx context.Context, key interface{}) (time.Duration, error)

	// X删除并带返回值 从缓存中删除一个或多个键，并返回其值。
	// 如果提供了多个键，它将返回最后一个被删除项的值。
	// md5:6e5f157befbc08c2
	X删除并带返回值(ctx context.Context, keys ...interface{}) (lastValue *gvar.Var, err error)

	// X清空 清空缓存中的所有数据。
	// 注意，此函数涉及敏感操作，应谨慎使用。
	// md5:8f66f62d0fce831a
	X清空(ctx context.Context) error

		// X关闭如果有必要，关闭缓存。 md5:f9a73a30e4b4b396
	X关闭(ctx context.Context) error
}
