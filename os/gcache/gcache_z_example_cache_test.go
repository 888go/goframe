// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gcache_test

import (
	"context"
	"fmt"
	"time"
	
	"github.com/888go/goframe/database/gredis"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/os/gcache"
	"github.com/888go/goframe/os/gctx"
)

func ExampleNew() {
// 创建一个缓存对象，
// 当然，你也可以很方便地直接使用gcache包的方法。
	c := gcache.New()

	// 设置缓存，不设置过期时间
	c.Set(ctx, "k1", "v1", 0)

	// Get cache
	v, _ := c.Get(ctx, "k1")
	fmt.Println(v)

	// Get cache size
	n, _ := c.Size(ctx)
	fmt.Println(n)

	// 指定的键名是否存在于缓存中
	b, _ := c.Contains(ctx, "k1")
	fmt.Println(b)

	// 删除并返回已删除的键值对
	fmt.Println(c.Remove(ctx, "k1"))

	// 关闭缓存对象，让垃圾回收器回收资源

	c.Close(ctx)

	// Output:
	// v1
	// 1
	// true
	// v1 <nil>
}

func ExampleCache_Set() {
	// Create a cache object,
	// Of course, you can also easily use the gcache package method directly
	c := gcache.New()

	// 设置缓存，不设置过期时间
	c.Set(ctx, "k1", g.Slice{1, 2, 3, 4, 5, 6, 7, 8, 9}, 0)

	// Get cache
	fmt.Println(c.Get(ctx, "k1"))

	// Output:
	// [1,2,3,4,5,6,7,8,9] <nil>
}

func ExampleCache_SetIfNotExist() {
	// Create a cache object,
	// Of course, you can also easily use the gcache package method directly
	c := gcache.New()

	// 当键名不存在时写入，并设置过期时间为1000毫秒
	k1, err := c.SetIfNotExist(ctx, "k1", "v1", 1000*time.Millisecond)
	fmt.Println(k1, err)

	// 当键名已存在时返回false
	k2, err := c.SetIfNotExist(ctx, "k1", "v2", 1000*time.Millisecond)
	fmt.Println(k2, err)

	// 打印当前键值对列表
	keys1, _ := c.Keys(ctx)
	fmt.Println(keys1)

	// 如果 `duration` 等于 0，则它不会过期。如果 `duration` 小于 0 或给定的 `value` 为 nil，则它会删除 `key`。
	c.SetIfNotExist(ctx, "k1", 0, -10000)

	// 等待1.5秒，直至K1: V1自动过期
	time.Sleep(1500 * time.Millisecond)

	// 再次打印当前键值对，会发现 K1: V1 已经过期
	keys2, _ := c.Keys(ctx)
	fmt.Println(keys2)

	// Output:
	// true <nil>
	// false <nil>
	// [k1]
	// [<nil>]
}

func ExampleCache_SetMap() {
	// Create a cache object,
	// Of course, you can also easily use the gcache package method directly
	c := gcache.New()

	// map[interface{}]interface{}：这是一个Go语言中的映射（map）类型，它的键和值都是接口类型（interface{}）。这意味着这个映射可以存储任意类型的键值对，因为interface{}可以表示任何类型。在实际使用中，这种类型的映射通常用于需要处理多种不同类型数据的场景，但需要注意，由于go的静态类型特性，在取值时需要进行类型断言转换。
	data := g.MapAnyAny{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
	}

// SetsBatch 函数通过 `data` 设置缓存中的键值对，该键值对在 `duration` 后过期。
// 如果 `duration` 等于 0，则不会过期。如果 `duration` 小于 0 或提供的 `value` 为 nil，则会删除 `data` 中的键。
	c.SetMap(ctx, data, 1000*time.Millisecond)

	// 获取指定键的值
	v1, _ := c.Get(ctx, "k1")
	v2, _ := c.Get(ctx, "k2")
	v3, _ := c.Get(ctx, "k3")

	fmt.Println(v1, v2, v3)

	// Output:
	// v1 v2 v3
}

func ExampleCache_Size() {
	// Create a cache object,
	// Of course, you can also easily use the gcache package method directly
	c := gcache.New()

	// 添加10个无过期时间的元素
	for i := 0; i < 10; i++ {
		c.Set(ctx, i, i, 0)
	}

	// Size 返回缓存中的项目数量。
	n, _ := c.Size(ctx)
	fmt.Println(n)

	// Output:
	// 10
}

func ExampleCache_Update() {
	// Create a cache object,
	// Of course, you can also easily use the gcache package method directly
	c := gcache.New()

// SetsBatch 函数通过 `data` 设置缓存中的键值对，该键值对在 `duration` 后过期。
// 如果 `duration` 等于 0，则不会过期。如果 `duration` 小于 0 或提供的 `value` 为 nil，则会删除 `data` 中的键。
	c.SetMap(ctx, g.MapAnyAny{"k1": "v1", "k2": "v2", "k3": "v3"}, 0)

	// Print the current key value pair
	k1, _ := c.Get(ctx, "k1")
	fmt.Println(k1)
	k2, _ := c.Get(ctx, "k2")
	fmt.Println(k2)
	k3, _ := c.Get(ctx, "k3")
	fmt.Println(k3)

	// Update 更新键 `key` 的值，但不会改变其过期时间，并返回旧的值。
	re, exist, _ := c.Update(ctx, "k1", "v11")
	fmt.Println(re, exist)

// 若`key`在缓存中不存在，返回的值`exist`为false。
// 若`key`在缓存中不存在，则此操作不做任何处理。
	re1, exist1, _ := c.Update(ctx, "k4", "v44")
	fmt.Println(re1, exist1)

	kup1, _ := c.Get(ctx, "k1")
	fmt.Println(kup1)
	kup2, _ := c.Get(ctx, "k2")
	fmt.Println(kup2)
	kup3, _ := c.Get(ctx, "k3")
	fmt.Println(kup3)

	// Output:
	// v1
	// v2
	// v3
	// v1 true
	//  false
	// v11
	// v2
	// v3
}

func ExampleCache_UpdateExpire() {
	// Create a cache object,
	// Of course, you can also easily use the gcache package method directly
	c := gcache.New()

	c.Set(ctx, "k1", "v1", 1000*time.Millisecond)
	expire, _ := c.GetExpire(ctx, "k1")
	fmt.Println(expire)

// UpdateExpire 更新键 `key` 的过期时间，并返回旧的过期持续时间值。
// 如果 `key` 不存在于缓存中，则返回 -1 并不做任何操作。
	c.UpdateExpire(ctx, "k1", 500*time.Millisecond)

	expire1, _ := c.GetExpire(ctx, "k1")
	fmt.Println(expire1)

	// May Output:
	// 1s
	// 500ms
}

func ExampleCache_Values() {
	// Create a cache object,
	// Of course, you can also easily use the gcache package method directly
	c := gcache.New()

	// Write value
	c.Set(ctx, "k1", g.Map{"k1": "v1", "k2": "v2"}, 0)
// c.Set(ctx, "k2", "Here is Value2", 0)
// 在给定的上下文ctx中，将键为"k2"的值设置为"Here is Value2"，并设置过期时间为0（表示永不过期）
// c.Set(ctx, "k3", 111, 0)
// 在给定的上下文ctx中，将键为"k3"的值设置为整数111，并设置过期时间为0（表示永不过期）
// 在上述代码中，`c` 应该是一个具有缓存功能的对象，`Set` 方法用于设置缓存项，参数包括操作的上下文、键名和对应的值以及可选的过期时间。这里设置的过期时间是0，通常意味着缓存项永不过期。

	// Values 返回缓存中的所有值作为一个切片。
	data, _ := c.Values(ctx)
	fmt.Println(data)

	// May Output:
	// [map[k1:v1 k2:v2]]
}

func ExampleCache_Close() {
	// Create a cache object,
	// Of course, you can also easily use the gcache package method directly
	c := gcache.New()

	// Set Cache
	c.Set(ctx, "k1", "v", 0)
	data, _ := c.Get(ctx, "k1")
	fmt.Println(data)

	// Close 在必要时关闭缓存。
	c.Close(ctx)

	data1, _ := c.Get(ctx, "k1")

	fmt.Println(data1)

	// Output:
	// v
	// v
}

func ExampleCache_Contains() {
	// Create a cache object,
	// Of course, you can also easily use the gcache package method directly
	c := gcache.New()

	// Set Cache
	c.Set(ctx, "k", "v", 0)

// Contains 返回 true 如果 `key` 存在于缓存中，否则返回 false。
	data, _ := c.Contains(ctx, "k")
	fmt.Println(data)

	// return false
	data1, _ := c.Contains(ctx, "k1")
	fmt.Println(data1)

	// Output:
	// true
	// false
}

func ExampleCache_Data() {
	// Create a cache object,
	// Of course, you can also easily use the gcache package method directly
	c := gcache.New()

	c.SetMap(ctx, g.MapAnyAny{"k1": "v1"}, 0)

	data, _ := c.Data(ctx)
	fmt.Println(data)

	// Set Cache
	c.Set(ctx, "k5", "v5", 0)
	data1, _ := c.Get(ctx, "k1")
	fmt.Println(data1)

	// Output:
	// map[k1:v1]
	// v1
}

func ExampleCache_Get() {
	// Create a cache object,
	// Of course, you can also easily use the gcache package method directly
	c := gcache.New()

	// Set Cache Object
	c.Set(ctx, "k1", "v1", 0)

// Get 方法用于获取并返回给定`key`关联的值。
// 如果该键不存在，其值为 nil 或已过期，则返回 nil。
	data, _ := c.Get(ctx, "k1")
	fmt.Println(data)

	// Output:
	// v1
}

func ExampleCache_GetExpire() {
	// Create a cache object,
	// Of course, you can also easily use the gcache package method directly
	c := gcache.New()

	// 设置缓存，不设置过期时间
	c.Set(ctx, "k", "v", 10000*time.Millisecond)

// GetExpire 从缓存中获取并返回`key`的过期时间。
// 如果`key`永不过期，则返回0。如果`key`在缓存中不存在，则返回-1。
	expire, _ := c.GetExpire(ctx, "k")
	fmt.Println(expire)

	// May Output:
	// 10s
}

func ExampleCache_GetOrSet() {
	// Create a cache object,
	// Of course, you can also easily use the gcache package method directly
	c := gcache.New()

// GetOrSet 获取并返回 `key` 对应的值，如果 `key` 不存在于缓存中，则设置 `key`-`value` 键值对并返回 `value`。
	data, _ := c.GetOrSet(ctx, "k", "v", 10000*time.Millisecond)
	fmt.Println(data)

	data1, _ := c.Get(ctx, "k")
	fmt.Println(data1)

	// Output:
	// v
	// v

}

func ExampleCache_GetOrSetFunc() {
	// Create a cache object,
	// Of course, you can also easily use the gcache package method directly
	c := gcache.New()

// GetOrSetFunc 函数用于获取并返回 `key` 对应的值，如果 `key` 不存在于缓存中，则使用函数 `f` 的结果设置 `key` 并返回其执行结果。
	c.GetOrSetFunc(ctx, "k1", func(ctx context.Context) (value interface{}, err error) {
		return "v1", nil
	}, 10000*time.Millisecond)
	v, _ := c.Get(ctx, "k1")
	fmt.Println(v)

	// 如果函数返回nil，则不执行任何操作
	c.GetOrSetFunc(ctx, "k2", func(ctx context.Context) (value interface{}, err error) {
		return nil, nil
	}, 10000*time.Millisecond)
	v1, _ := c.Get(ctx, "k2")
	fmt.Println(v1)

	// Output:
	// v1
}

func ExampleCache_GetOrSetFuncLock() {
	// Create a cache object,
	// Of course, you can also easily use the gcache package method directly
	c := gcache.New()

	// 修改锁定：请注意，为了保证并发安全，函数`f`应当在写入互斥锁保护下执行。
	c.GetOrSetFuncLock(ctx, "k1", func(ctx context.Context) (value interface{}, err error) {
		return "v1", nil
	}, 0)
	v, _ := c.Get(ctx, "k1")
	fmt.Println(v)

	// 修改失败
	c.GetOrSetFuncLock(ctx, "k1", func(ctx context.Context) (value interface{}, err error) {
		return "update v1", nil
	}, 0)
	v, _ = c.Get(ctx, "k1")
	fmt.Println(v)

	c.Remove(ctx, g.Slice{"k1"}...)

	// Output:
	// v1
	// v1
}

func ExampleCache_Keys() {
	// Create a cache object,
	// Of course, you can also easily use the gcache package method directly
	c := gcache.New()

	c.SetMap(ctx, g.MapAnyAny{"k1": "v1"}, 0)

	// 打印当前键值对列表
	keys1, _ := c.Keys(ctx)
	fmt.Println(keys1)

	// Output:
	// [k1]
}

func ExampleCache_KeyStrings() {
	c := gcache.New()

	c.SetMap(ctx, g.MapAnyAny{"k1": "v1", "k2": "v2"}, 0)

	// KeyStrings 返回缓存中的所有键，以字符串切片形式。
	keys, _ := c.KeyStrings(ctx)
	fmt.Println(keys)

	// May Output:
	// [k1 k2]
}

func ExampleCache_Remove() {
	// Create a cache object,
	// Of course, you can also easily use the gcache package method directly
	c := gcache.New()

	c.SetMap(ctx, g.MapAnyAny{"k1": "v1", "k2": "v2"}, 0)

// Remove 从缓存中删除一个或多个键，并返回其对应的值。
// 如果提供了多个键，它将返回最后被删除项的值。
	remove, _ := c.Remove(ctx, "k1")
	fmt.Println(remove)

	data, _ := c.Data(ctx)
	fmt.Println(data)

	// Output:
	// v1
	// map[k2:v2]
}

func ExampleCache_Removes() {
	// Create a cache object,
	// Of course, you can also easily use the gcache package method directly
	c := gcache.New()

	c.SetMap(ctx, g.MapAnyAny{"k1": "v1", "k2": "v2", "k3": "v3", "k4": "v4"}, 0)

// Remove 从缓存中删除一个或多个键，并返回其对应的值。
// 如果提供了多个键，它将返回最后被删除项的值。
	c.Removes(ctx, g.Slice{"k1", "k2", "k3"})

	data, _ := c.Data(ctx)
	fmt.Println(data)

	// Output:
	// map[k4:v4]
}

func ExampleCache_MustGet() {
// 拦截 panic 异常信息
// err 为空，因此不执行 panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	// Create a cache object,
	// Of course, you can also easily use the gcache package method directly
	c := gcache.New()

	// Set Cache Object
	c.Set(ctx, "k1", "v1", 0)

	// MustGet 行为类似于 Get，但是当出现任何错误时，它会触发panic（异常）。
	k2 := c.MustGet(ctx, "k2")
	fmt.Println(k2)

	k1 := c.MustGet(ctx, "k1")
	fmt.Println(k1)

	// Output:
	// v1
}

func ExampleCache_MustGetOrSet() {
	// Create a cache object,
	// Of course, you can also easily use the gcache package method directly
	c := gcache.New()

	// MustGetOrSet 行为类似于 GetOrSet，但当发生任何错误时，它会触发panic（异常）。
	k1 := c.MustGetOrSet(ctx, "k1", "v1", 0)
	fmt.Println(k1)

	k2 := c.MustGetOrSet(ctx, "k1", "v2", 0)
	fmt.Println(k2)

	// Output:
	// v1
	// v1
}

func ExampleCache_MustGetOrSetFunc() {
	// Create a cache object,
	// Of course, you can also easily use the gcache package method directly
	c := gcache.New()

	// MustGetOrSetFunc 行为类似于 GetOrSetFunc，但当出现任何错误时，它会触发panic（异常）。
	c.MustGetOrSetFunc(ctx, "k1", func(ctx context.Context) (value interface{}, err error) {
		return "v1", nil
	}, 10000*time.Millisecond)
	v := c.MustGet(ctx, "k1")
	fmt.Println(v)

	c.MustGetOrSetFunc(ctx, "k2", func(ctx context.Context) (value interface{}, err error) {
		return nil, nil
	}, 10000*time.Millisecond)
	v1 := c.MustGet(ctx, "k2")
	fmt.Println(v1)

	// Output:
	// v1
	//
}

func ExampleCache_MustGetOrSetFuncLock() {
	// Create a cache object,
	// Of course, you can also easily use the gcache package method directly
	c := gcache.New()

	// MustGetOrSetFuncLock 行为类似于 GetOrSetFuncLock，但是当发生任何错误时，它会引发panic（恐慌）。
	c.MustGetOrSetFuncLock(ctx, "k1", func(ctx context.Context) (value interface{}, err error) {
		return "v1", nil
	}, 0)
	v := c.MustGet(ctx, "k1")
	fmt.Println(v)

	// 修改失败
	c.MustGetOrSetFuncLock(ctx, "k1", func(ctx context.Context) (value interface{}, err error) {
		return "update v1", nil
	}, 0)
	v = c.MustGet(ctx, "k1")
	fmt.Println(v)

	// Output:
	// v1
	// v1
}

func ExampleCache_MustContains() {

	// Create a cache object,
	// Of course, you can also easily use the gcache package method directly
	c := gcache.New()

	// Set Cache
	c.Set(ctx, "k", "v", 0)

// MustContains 返回一个布尔值，如果 `key` 存在于缓存中则返回 true，否则返回 false。
	data := c.MustContains(ctx, "k")
	fmt.Println(data)

	// return false
	data1 := c.MustContains(ctx, "k1")
	fmt.Println(data1)

	// Output:
	// true
	// false
}

func ExampleCache_MustGetExpire() {
	// Create a cache object,
	// Of course, you can also easily use the gcache package method directly
	c := gcache.New()

	// 设置缓存，不设置过期时间
	c.Set(ctx, "k", "v", 10000*time.Millisecond)

	// MustGetExpire 行为类似于 GetExpire，但是当发生任何错误时它会触发panic（异常）。
	expire := c.MustGetExpire(ctx, "k")
	fmt.Println(expire)

	// May Output:
	// 10s
}

func ExampleCache_MustSize() {
	// Create a cache object,
	// Of course, you can also easily use the gcache package method directly
	c := gcache.New()

	// 添加10个无过期时间的元素
	for i := 0; i < 10; i++ {
		c.Set(ctx, i, i, 0)
	}

	// Size 返回缓存中的项目数量。
	n := c.MustSize(ctx)
	fmt.Println(n)

	// Output:
	// 10
}

func ExampleCache_MustData() {
	// Create a cache object,
	// Of course, you can also easily use the gcache package method directly
	c := gcache.New()

	c.SetMap(ctx, g.MapAnyAny{"k1": "v1", "k2": "v2"}, 0)

	data := c.MustData(ctx)
	fmt.Println(data)

	// May Output:
	// map[k1:v1 k2:v2]
}

func ExampleCache_MustKeys() {
	// Create a cache object,
	// Of course, you can also easily use the gcache package method directly
	c := gcache.New()

	c.SetMap(ctx, g.MapAnyAny{"k1": "v1", "k2": "v2"}, 0)

	// MustKeys 行为类似于 Keys，但如果出现任何错误，它会引发 panic。
	keys1 := c.MustKeys(ctx)
	fmt.Println(keys1)

	// May Output:
	// [k1 k2]

}

func ExampleCache_MustKeyStrings() {
	c := gcache.New()

	c.SetMap(ctx, g.MapAnyAny{"k1": "v1", "k2": "v2"}, 0)

// MustKeyStrings 返回缓存中的所有键作为字符串切片。
// MustKeyStrings 类似于 KeyStrings，但在出现任何错误时会触发 panic。
	keys := c.MustKeyStrings(ctx)
	fmt.Println(keys)

	// May Output:
	// [k1 k2]
}

func ExampleCache_MustValues() {
	// Create a cache object,
	// Of course, you can also easily use the gcache package method directly
	c := gcache.New()

	// Write value
	c.Set(ctx, "k1", "v1", 0)

	// MustValues 返回缓存中所有值的切片。
	data := c.MustValues(ctx)
	fmt.Println(data)

	// Output:
	// [v1]
}

func ExampleCache_SetAdapter() {
	var (
		err         error
		ctx         = gctx.New()
		cache       = gcache.New()
		redisConfig = &gredis.Config{
			Address: "127.0.0.1:6379",
			Db:      9,
		}
		cacheKey   = `key`
		cacheValue = `value`
	)
	// 创建Redis客户端对象。
	redis, err := gredis.New(redisConfig)
	if err != nil {
		panic(err)
	}
	// 创建Redis缓存适配器并将它设置为缓存对象。
	cache.SetAdapter(gcache.NewAdapterRedis(redis))

	// 使用缓存对象进行设置和获取操作。
	err = cache.Set(ctx, cacheKey, cacheValue, time.Second)
	if err != nil {
		panic(err)
	}
	fmt.Println(cache.MustGet(ctx, cacheKey).String())

	// 使用redis客户端获取。
	fmt.Println(redis.MustDo(ctx, "GET", cacheKey).String())

	// May Output:
	// value
	// value
}

func ExampleCache_GetAdapter() {
	var (
		err         error
		ctx         = gctx.New()
		cache       = gcache.New()
		redisConfig = &gredis.Config{
			Address: "127.0.0.1:6379",
			Db:      10,
		}
		cacheKey   = `key`
		cacheValue = `value`
	)
	redis, err := gredis.New(redisConfig)
	if err != nil {
		panic(err)
	}
	cache.SetAdapter(gcache.NewAdapterRedis(redis))

	// 使用缓存对象进行设置和获取操作。
	err = cache.Set(ctx, cacheKey, cacheValue, time.Second)
	if err != nil {
		panic(err)
	}
	fmt.Println(cache.MustGet(ctx, cacheKey).String())

	// 使用redis客户端获取。
	v, err := cache.GetAdapter().(*gcache.AdapterRedis).Get(ctx, cacheKey)
	fmt.Println(err)
	fmt.Println(v.String())

	// May Output:
	// value
	// <nil>
	// value
}
