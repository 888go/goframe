// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gcache_test

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
)

func ExampleNew() {
	// 创建一个缓存对象，
	// 当然，你也可以直接轻松地使用gcache包提供的方法。 md5:ffacc88538a18ac1
	c := gcache.New()

	// 设置不带过期时间的缓存. md5:10e925a877b589df
	c.Set(ctx, "k1", "v1", 0)

	// Get cache
	v, _ := c.Get(ctx, "k1")
	fmt.Println(v)

	// Get cache size
	n, _ := c.Size(ctx)
	fmt.Println(n)

	// 指定的键名是否存在于缓存中. md5:9f18fb652f082e8d
	b, _ := c.Contains(ctx, "k1")
	fmt.Println(b)

	// 删除并返回被删除的键值对. md5:01f96ef0f0eae0e3
	fmt.Println(c.Remove(ctx, "k1"))

	// 关闭缓存对象，允许GC回收资源. md5:e6035a9d9a9583ce

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

	// 设置不带过期时间的缓存. md5:10e925a877b589df
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

	// 如果键名不存在，则写入，并设置过期时间为1000毫秒. md5:41f1844c720c4e5f
	k1, err := c.SetIfNotExist(ctx, "k1", "v1", 1000*time.Millisecond)
	fmt.Println(k1, err)

	// 当键名已存在时返回false. md5:db29c3756eb62f45
	k2, err := c.SetIfNotExist(ctx, "k1", "v2", 1000*time.Millisecond)
	fmt.Println(k2, err)

	// 打印当前的键值对列表. md5:9b92acd8c3138f30
	keys1, _ := c.Keys(ctx)
	fmt.Println(keys1)

	// 如果`duration`等于0，它不会过期。如果`duration`小于0或给定的`value`为nil，它将删除`key`。 md5:a0794bc140ecff80
	c.SetIfNotExist(ctx, "k1", 0, -10000)

	// 等待1.5秒，让K1: V1自动过期. md5:58d8a37819ac1a31
	time.Sleep(1500 * time.Millisecond)

	// 再次打印当前的键值对，发现K1: V1 已过期. md5:50e60452bc46b7a6
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

	// 一个映射，其键和值都是接口类型. md5:5e21ae79c908b4df
	data := g.MapAnyAny{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
	}

	// 使用键值对`data`设置批处理缓存，过期时间为`duration`。
	// 如果`duration`为0，则不会过期。如果`duration`小于0或给定的`value`为nil，将删除`data`中的键。 md5:b2f121999e39c24d
	c.SetMap(ctx, data, 1000*time.Millisecond)

	// 获取指定的键值. md5:78fc35a6610c5179
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

	// 添加10个不带过期的元素. md5:cbefc995139f6ed9
	for i := 0; i < 10; i++ {
		c.Set(ctx, i, i, 0)
	}

	// Size 返回缓存中的项目数量。 md5:2122f80de9340261
	n, _ := c.Size(ctx)
	fmt.Println(n)

	// Output:
	// 10
}

func ExampleCache_Update() {
	// Create a cache object,
	// Of course, you can also easily use the gcache package method directly
	c := gcache.New()

	// 使用键值对`data`设置批处理缓存，过期时间为`duration`。
	// 如果`duration`为0，则不会过期。如果`duration`小于0或给定的`value`为nil，将删除`data`中的键。 md5:b2f121999e39c24d
	c.SetMap(ctx, g.MapAnyAny{"k1": "v1", "k2": "v2", "k3": "v3"}, 0)

	// Print the current key value pair
	k1, _ := c.Get(ctx, "k1")
	fmt.Println(k1)
	k2, _ := c.Get(ctx, "k2")
	fmt.Println(k2)
	k3, _ := c.Get(ctx, "k3")
	fmt.Println(k3)

	// Update 更新 `key` 对应的值，而不改变其过期时间，并返回旧值。 md5:1e7dc1ae84b2f449
	re, exist, _ := c.Update(ctx, "k1", "v11")
	fmt.Println(re, exist)

	// 返回的值`exist`如果`key`在缓存中不存在，则为false。
	// 如果`key`不在缓存中，它将不执行任何操作。 md5:1ecdac9a4397de58
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

	// UpdateExpire 更新键`key`的过期时间，并返回旧的过期持续时间值。如果缓存中不存在`key`，则返回-1并什么都不做。 md5:8a59f61ead71c844
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
	// 将键为 "k2" 的值设置为 "Here is Value2"，过期时间为 0（默认不设置过期时间）
	// 将键为 "k3" 的值设置为整数 111，过期时间为 0（默认不设置过期时间） md5:b8ee9984ed9da43d

	// Values 返回缓存中所有的值作为切片。 md5:dc00b32eb8913e9b
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

	// Close如果有必要，关闭缓存。 md5:f9a73a30e4b4b396
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

	// Contains 如果`key`存在于缓存中，则返回true，否则返回false。 md5:370eab2c5835bfb1
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

	// Get 获取并返回给定`key`关联的值。
	// 如果键不存在、其值为nil或已过期，它将返回nil。 md5:2999106994454771
	data, _ := c.Get(ctx, "k1")
	fmt.Println(data)

	// Output:
	// v1
}

func ExampleCache_GetExpire() {
	// Create a cache object,
	// Of course, you can also easily use the gcache package method directly
	c := gcache.New()

	// 设置不带过期时间的缓存. md5:10e925a877b589df
	c.Set(ctx, "k", "v", 10000*time.Millisecond)

	// GetExpire 从缓存中检索并返回`key`的过期时间。如果`key`不过期，它将返回0。如果`key`在缓存中不存在，它将返回-1。 md5:a60a46e9632013e1
	expire, _ := c.GetExpire(ctx, "k")
	fmt.Println(expire)

	// May Output:
	// 10s
}

func ExampleCache_GetOrSet() {
	// Create a cache object,
	// Of course, you can also easily use the gcache package method directly
	c := gcache.New()

	// GetOrSet 从缓存中检索并返回`key`对应的值，如果`key`在缓存中不存在，则设置`key-value`对，并返回`value`。 md5:f1f24272b9b4a43c
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

	// GetOrSetFunc 方法尝试获取并返回`key`对应的值，如果`key`在缓存中不存在，则使用函数`f`的结果设置`key`的值，
	// 并返回该函数的结果。 md5:6a21b8ed72969e95
	c.GetOrSetFunc(ctx, "k1", func(ctx context.Context) (value interface{}, err error) {
		return "v1", nil
	}, 10000*time.Millisecond)
	v, _ := c.Get(ctx, "k1")
	fmt.Println(v)

	// 如果函数返回nil，不执行任何操作. md5:a7c277c9df4048d6
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

	// 修改锁定注意，为了并发安全，函数 `f` 应该在写入锁的保护下执行。 md5:a86de4ea66d58271
	c.GetOrSetFuncLock(ctx, "k1", func(ctx context.Context) (value interface{}, err error) {
		return "v1", nil
	}, 0)
	v, _ := c.Get(ctx, "k1")
	fmt.Println(v)

	// Modification failed
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

	// 打印当前的键值对列表. md5:9b92acd8c3138f30
	keys1, _ := c.Keys(ctx)
	fmt.Println(keys1)

	// Output:
	// [k1]
}

func ExampleCache_KeyStrings() {
	c := gcache.New()

	c.SetMap(ctx, g.MapAnyAny{"k1": "v1", "k2": "v2"}, 0)

	// KeyStrings返回缓存中的所有键作为字符串切片。 md5:3b0126221389825e
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

	// Remove 从缓存中删除一个或多个键，并返回其值。
	// 如果提供了多个键，它将返回最后一个被删除项的值。 md5:6e5f157befbc08c2
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

	// Remove 从缓存中删除一个或多个键，并返回其值。
	// 如果提供了多个键，它将返回最后一个被删除项的值。 md5:6e5f157befbc08c2
	c.Removes(ctx, g.Slice{"k1", "k2", "k3"})

	data, _ := c.Data(ctx)
	fmt.Println(data)

	// Output:
	// map[k4:v4]
}

func ExampleCache_Clear() {
	// Create a cache object,
	// Of course, you can also easily use the gcache package method directly
	c := gcache.New()

	c.SetMap(ctx, g.MapAnyAny{"k1": "v1", "k2": "v2", "k3": "v3", "k4": "v4"}, 0)

	// 清空缓存的所有数据。 md5:13010db2c416938b
	c.Clear(ctx)

	data, _ := c.Data(ctx)
	fmt.Println(data)

	// Output:
	// map[]
}

func ExampleCache_MustGet() {
	// 拦截恐慌异常信息
	// err 为空，因此不执行恐慌操作 md5:aa899aa9abc889f7
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

	// MustGet 的行为就像 Get 一样，但如果发生任何错误，它会引发 panic。 md5:9004545d221e9637
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

	// MustGetOrSet 的行为类似于 GetOrSet，但是如果发生任何错误，它会直接 panic。 md5:684c6b06451a2f6f
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

	// MustGetOrSetFunc 行为类似于 GetOrSetFunc，但如果发生任何错误，则会引发 panic。 md5:07fd1ef2dbfce0b4
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

	// MustGetOrSetFuncLock 行为与 GetOrSetFuncLock 类似，但如果发生任何错误，它将引发恐慌。 md5:7f84f54a71da5305
	c.MustGetOrSetFuncLock(ctx, "k1", func(ctx context.Context) (value interface{}, err error) {
		return "v1", nil
	}, 0)
	v := c.MustGet(ctx, "k1")
	fmt.Println(v)

	// Modification failed
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

	// MustContains 返回true如果`key`在缓存中存在，否则返回false。
	// 返回true md5:226a8dda1fb50b87
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

	// 设置不带过期时间的缓存. md5:10e925a877b589df
	c.Set(ctx, "k", "v", 10000*time.Millisecond)

	// MustGetExpire 的行为类似于 GetExpire，但如果发生任何错误，它会直接 panic。 md5:c97fa5941bbc47a3
	expire := c.MustGetExpire(ctx, "k")
	fmt.Println(expire)

	// May Output:
	// 10s
}

func ExampleCache_MustSize() {
	// Create a cache object,
	// Of course, you can also easily use the gcache package method directly
	c := gcache.New()

	// 添加10个不带过期的元素. md5:cbefc995139f6ed9
	for i := 0; i < 10; i++ {
		c.Set(ctx, i, i, 0)
	}

	// Size 返回缓存中的项目数量。 md5:2122f80de9340261
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

	// MustKeys 行为与 Keys 类似，但如果发生任何错误，它将引发 panic。 md5:7f7801d0cd170166
	keys1 := c.MustKeys(ctx)
	fmt.Println(keys1)

	// May Output:
	// [k1 k2]

}

func ExampleCache_MustKeyStrings() {
	c := gcache.New()

	c.SetMap(ctx, g.MapAnyAny{"k1": "v1", "k2": "v2"}, 0)

	// MustKeyStrings 返回缓存中的所有键作为字符串切片。
	// MustKeyStrings 的行为类似于 KeyStrings，但如果发生任何错误，它将引发恐慌。 md5:e647a507f2385601
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

	// MustValues 返回缓存中的所有值作为切片。 md5:8b4269e238366b51
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
	// 创建Redis客户端对象。 md5:f412dd032940c79e
	redis, err := gredis.New(redisConfig)
	if err != nil {
		panic(err)
	}
	// 创建Redis缓存适配器，并将其设置到缓存对象中。 md5:fe080f47e7881b0a
	cache.SetAdapter(gcache.NewAdapterRedis(redis))

	// 使用缓存对象进行设置和获取。 md5:27779d48bc4565ef
	err = cache.Set(ctx, cacheKey, cacheValue, time.Second)
	if err != nil {
		panic(err)
	}
	fmt.Println(cache.MustGet(ctx, cacheKey).String())

	// 使用redis客户端获取。 md5:413fbbedeb694205
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

	// 使用缓存对象进行设置和获取。 md5:27779d48bc4565ef
	err = cache.Set(ctx, cacheKey, cacheValue, time.Second)
	if err != nil {
		panic(err)
	}
	fmt.Println(cache.MustGet(ctx, cacheKey).String())

	// 使用redis客户端获取。 md5:413fbbedeb694205
	v, err := cache.GetAdapter().(*gcache.AdapterRedis).Get(ctx, cacheKey)
	fmt.Println(err)
	fmt.Println(v.String())

	// May Output:
	// value
	// <nil>
	// value
}
