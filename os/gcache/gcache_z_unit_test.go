// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 使用`go test`命令，对所有`.go`文件进行测试，指定运行基准测试（Benchmark）中的所有模式（".*"），同时输出内存使用情况（-benchmem）。 md5:81db3d7bd1ed4da8

package 缓存类_test

import (
	"context"
	"math"
	"testing"
	"time"

	gset "github.com/888go/goframe/container/gset"
	"github.com/888go/goframe/frame/g"
	gcache "github.com/888go/goframe/os/gcache"
	grpool "github.com/888go/goframe/os/grpool"
	gtest "github.com/888go/goframe/test/gtest"
	guid "github.com/888go/goframe/util/guid"
)

var (
	ctx = context.Background()
)

func TestCache_GCache_Set(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertNil(gcache.X设置值(ctx, 1, 11, 0))
		defer gcache.X删除并带返回值(ctx, g.Slice别名{1, 2, 3}...)
		v, _ := gcache.X取值(ctx, 1)
		t.Assert(v, 11)
		b, _ := gcache.X是否存在(ctx, 1)
		t.Assert(b, true)
	})
}

func TestCache_Set(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		c := gcache.X创建()
		defer c.X关闭(ctx)
		t.Assert(c.X设置值(ctx, 1, 11, 0), nil)
		v, _ := c.X取值(ctx, 1)
		t.Assert(v, 11)
		b, _ := c.X是否存在(ctx, 1)
		t.Assert(b, true)
	})
}

func TestCache_Set_Expire(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		cache := gcache.X创建()
		t.Assert(cache.X设置值(ctx, 2, 22, 100*time.Millisecond), nil)
		v, _ := cache.X取值(ctx, 2)
		t.Assert(v, 22)
		time.Sleep(200 * time.Millisecond)
		v, _ = cache.X取值(ctx, 2)
		t.Assert(v, nil)
		time.Sleep(3 * time.Second)
		n, _ := cache.X取数量(ctx)
		t.Assert(n, 0)
		t.Assert(cache.X关闭(ctx), nil)
	})

	gtest.C(t, func(t *gtest.T) {
		cache := gcache.X创建()
		t.Assert(cache.X设置值(ctx, 1, 11, 100*time.Millisecond), nil)
		v, _ := cache.X取值(ctx, 1)
		t.Assert(v, 11)
		time.Sleep(200 * time.Millisecond)
		v, _ = cache.X取值(ctx, 1)
		t.Assert(v, nil)
	})
}

func TestCache_Update(t *testing.T) {
	// gcache
	gtest.C(t, func(t *gtest.T) {
		key := guid.X生成()
		t.AssertNil(gcache.X设置值(ctx, key, 11, 3*time.Second))
		expire1, _ := gcache.X取过期时间(ctx, key)
		oldValue, exist, err := gcache.X更新值(ctx, key, 12)
		t.AssertNil(err)
		t.Assert(oldValue, 11)
		t.Assert(exist, true)

		expire2, _ := gcache.X取过期时间(ctx, key)
		v, _ := gcache.X取值(ctx, key)
		t.Assert(v, 12)
		t.Assert(math.Ceil(expire1.Seconds()), math.Ceil(expire2.Seconds()))
	})
	// gcache.Cache
	gtest.C(t, func(t *gtest.T) {
		cache := gcache.X创建()
		t.AssertNil(cache.X设置值(ctx, 1, 11, 3*time.Second))

		oldValue, exist, err := cache.X更新值(ctx, 1, 12)
		t.AssertNil(err)
		t.Assert(oldValue, 11)
		t.Assert(exist, true)

		expire1, _ := cache.X取过期时间(ctx, 1)
		expire2, _ := cache.X取过期时间(ctx, 1)
		v, _ := cache.X取值(ctx, 1)
		t.Assert(v, 12)
		t.Assert(math.Ceil(expire1.Seconds()), math.Ceil(expire2.Seconds()))
	})
}

func TestCache_UpdateExpire(t *testing.T) {
	// gcache
	gtest.C(t, func(t *gtest.T) {
		key := guid.X生成()
		t.AssertNil(gcache.X设置值(ctx, key, 11, 3*time.Second))
		defer gcache.X删除并带返回值(ctx, key)
		oldExpire, _ := gcache.X取过期时间(ctx, key)
		newExpire := 10 * time.Second
		oldExpire2, err := gcache.X更新过期时间(ctx, key, newExpire)
		t.AssertNil(err)
		t.AssertIN(oldExpire2, g.Slice别名{oldExpire, `2.999s`})

		e, _ := gcache.X取过期时间(ctx, key)
		t.AssertNE(e, oldExpire)
		e, _ = gcache.X取过期时间(ctx, key)
		t.Assert(math.Ceil(e.Seconds()), 10)
	})
	// gcache.Cache
	gtest.C(t, func(t *gtest.T) {
		cache := gcache.X创建()
		t.AssertNil(cache.X设置值(ctx, 1, 11, 3*time.Second))
		oldExpire, _ := cache.X取过期时间(ctx, 1)
		newExpire := 10 * time.Second
		oldExpire2, err := cache.X更新过期时间(ctx, 1, newExpire)
		t.AssertNil(err)
		t.AssertIN(oldExpire2, g.Slice别名{oldExpire, `2.999s`})

		e, _ := cache.X取过期时间(ctx, 1)
		t.AssertNE(e, oldExpire)

		e, _ = cache.X取过期时间(ctx, 1)
		t.Assert(math.Ceil(e.Seconds()), 10)
	})
}

func TestCache_Keys_Values(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		c := gcache.X创建()
		for i := 0; i < 10; i++ {
			t.Assert(c.X设置值(ctx, i, i*10, 0), nil)
		}
		var (
			keys, _   = c.X取所有键(ctx)
			values, _ = c.X取所有值(ctx)
		)
		t.Assert(len(keys), 10)
		t.Assert(len(values), 10)
		t.AssertIN(0, keys)
		t.AssertIN(90, values)
	})
}

func TestCache_LRU(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		cache := gcache.X创建(2)
		for i := 0; i < 10; i++ {
			t.AssertNil(cache.X设置值(ctx, i, i, 0))
		}
		n, _ := cache.X取数量(ctx)
		t.Assert(n, 10)
		v, _ := cache.X取值(ctx, 6)
		t.Assert(v, 6)
		time.Sleep(4 * time.Second)
		g.X日志类().X输出并格式化DEBU(ctx, `items after lru: %+v`, cache.X取所有键值Map副本PANI(ctx))
		n, _ = cache.X取数量(ctx)
		t.Assert(n, 2)
		v, _ = cache.X取值(ctx, 6)
		t.Assert(v, 6)
		v, _ = cache.X取值(ctx, 1)
		t.Assert(v, nil)
		t.Assert(cache.X关闭(ctx), nil)
	})
}

func TestCache_LRU_expire(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		cache := gcache.X创建(2)
		t.Assert(cache.X设置值(ctx, 1, nil, 1000), nil)
		n, _ := cache.X取数量(ctx)
		t.Assert(n, 1)
		v, _ := cache.X取值(ctx, 1)

		t.Assert(v, nil)
	})
}

func TestCache_SetIfNotExist(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		cache := gcache.X创建()
		ok, err := cache.X设置值并跳过已存在(ctx, 1, 11, 0)
		t.AssertNil(err)
		t.Assert(ok, true)

		v, _ := cache.X取值(ctx, 1)
		t.Assert(v, 11)

		ok, err = cache.X设置值并跳过已存在(ctx, 1, 22, 0)
		t.AssertNil(err)
		t.Assert(ok, false)

		v, _ = cache.X取值(ctx, 1)
		t.Assert(v, 11)

		ok, err = cache.X设置值并跳过已存在(ctx, 2, 22, 0)
		t.AssertNil(err)
		t.Assert(ok, true)

		v, _ = cache.X取值(ctx, 2)
		t.Assert(v, 22)

		gcache.X删除并带返回值(ctx, g.Slice别名{1, 2, 3}...)
		ok, err = gcache.X设置值并跳过已存在(ctx, 1, 11, 0)
		t.AssertNil(err)
		t.Assert(ok, true)

		v, _ = gcache.X取值(ctx, 1)
		t.Assert(v, 11)

		ok, err = gcache.X设置值并跳过已存在(ctx, 1, 22, 0)
		t.AssertNil(err)
		t.Assert(ok, false)

		v, _ = gcache.X取值(ctx, 1)
		t.Assert(v, 11)
	})
}

func TestCache_SetIfNotExistFunc(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		cache := gcache.X创建()
		exist, err := cache.X设置值并跳过已存在_函数(ctx, 1, func(ctx context.Context) (value interface{}, err error) {
			return 11, nil
		}, 0)
		t.AssertNil(err)
		t.Assert(exist, true)

		v, _ := cache.X取值(ctx, 1)
		t.Assert(v, 11)

		exist, err = cache.X设置值并跳过已存在_函数(ctx, 1, func(ctx context.Context) (value interface{}, err error) {
			return 22, nil
		}, 0)
		t.AssertNil(err)
		t.Assert(exist, false)

		v, _ = cache.X取值(ctx, 1)
		t.Assert(v, 11)
	})
	gtest.C(t, func(t *gtest.T) {
		gcache.X删除并带返回值(ctx, g.Slice别名{1, 2, 3}...)

		ok, err := gcache.X设置值并跳过已存在_函数(ctx, 1, func(ctx context.Context) (value interface{}, err error) {
			return 11, nil
		}, 0)
		t.AssertNil(err)
		t.Assert(ok, true)

		v, _ := gcache.X取值(ctx, 1)
		t.Assert(v, 11)

		ok, err = gcache.X设置值并跳过已存在_函数(ctx, 1, func(ctx context.Context) (value interface{}, err error) {
			return 22, nil
		}, 0)
		t.AssertNil(err)
		t.Assert(ok, false)

		v, _ = gcache.X取值(ctx, 1)
		t.Assert(v, 11)
	})
}

func TestCache_SetIfNotExistFuncLock(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		cache := gcache.X创建()
		exist, err := cache.X设置值并跳过已存在_并发安全函数(ctx, 1, func(ctx context.Context) (value interface{}, err error) {
			return 11, nil
		}, 0)
		t.AssertNil(err)
		t.Assert(exist, true)

		v, _ := cache.X取值(ctx, 1)
		t.Assert(v, 11)

		exist, err = cache.X设置值并跳过已存在_并发安全函数(ctx, 1, func(ctx context.Context) (value interface{}, err error) {
			return 22, nil
		}, 0)
		t.AssertNil(err)
		t.Assert(exist, false)

		v, _ = cache.X取值(ctx, 1)
		t.Assert(v, 11)
	})
	gtest.C(t, func(t *gtest.T) {
		gcache.X删除并带返回值(ctx, g.Slice别名{1, 2, 3}...)

		exist, err := gcache.X设置值并跳过已存在_并发安全函数(ctx, 1, func(ctx context.Context) (value interface{}, err error) {
			return 11, nil
		}, 0)
		t.AssertNil(err)
		t.Assert(exist, true)

		v, _ := gcache.X取值(ctx, 1)
		t.Assert(v, 11)

		exist, err = gcache.X设置值并跳过已存在_并发安全函数(ctx, 1, func(ctx context.Context) (value interface{}, err error) {
			return 22, nil
		}, 0)
		t.AssertNil(err)
		t.Assert(exist, false)

		v, _ = gcache.X取值(ctx, 1)
		t.Assert(v, 11)
	})
}

func TestCache_SetMap(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		cache := gcache.X创建()
		t.AssertNil(cache.X设置Map(ctx, g.MapAnyAny{1: 11, 2: 22}, 0))
		v, _ := cache.X取值(ctx, 1)
		t.Assert(v, 11)

		gcache.X删除并带返回值(ctx, g.Slice别名{1, 2, 3}...)
		t.AssertNil(gcache.X设置Map(ctx, g.MapAnyAny{1: 11, 2: 22}, 0))
		v, _ = cache.X取值(ctx, 1)
		t.Assert(v, 11)
	})
}

func TestCache_GetOrSet(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		cache := gcache.X创建()
		value, err := cache.X取值或设置值(ctx, 1, 11, 0)
		t.AssertNil(err)
		t.Assert(value, 11)

		v, _ := cache.X取值(ctx, 1)
		t.Assert(v, 11)
		value, err = cache.X取值或设置值(ctx, 1, 111, 0)
		t.AssertNil(err)
		t.Assert(value, 11)

		v, _ = cache.X取值(ctx, 1)
		t.Assert(v, 11)
	})

	gtest.C(t, func(t *gtest.T) {
		gcache.X删除并带返回值(ctx, g.Slice别名{1, 2, 3}...)
		value, err := gcache.X取值或设置值(ctx, 1, 11, 0)
		t.AssertNil(err)
		t.Assert(value, 11)

		v, err := gcache.X取值(ctx, 1)
		t.AssertNil(err)
		t.Assert(v, 11)

		value, err = gcache.X取值或设置值(ctx, 1, 111, 0)
		t.AssertNil(err)
		t.Assert(value, 11)

		v, err = gcache.X取值(ctx, 1)
		t.AssertNil(err)
		t.Assert(v, 11)
	})
}

func TestCache_GetOrSetFunc(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		cache := gcache.X创建()
		cache.X取值或设置值_函数(ctx, 1, func(ctx context.Context) (value interface{}, err error) {
			return 11, nil
		}, 0)
		v, _ := cache.X取值(ctx, 1)
		t.Assert(v, 11)

		cache.X取值或设置值_函数(ctx, 1, func(ctx context.Context) (value interface{}, err error) {
			return 111, nil
		}, 0)
		v, _ = cache.X取值(ctx, 1)
		t.Assert(v, 11)

		gcache.X删除并带返回值(ctx, g.Slice别名{1, 2, 3}...)

		gcache.X取值或设置值_函数(ctx, 1, func(ctx context.Context) (value interface{}, err error) {
			return 11, nil
		}, 0)
		v, _ = cache.X取值(ctx, 1)
		t.Assert(v, 11)

		gcache.X取值或设置值_函数(ctx, 1, func(ctx context.Context) (value interface{}, err error) {
			return 111, nil
		}, 0)
		v, _ = cache.X取值(ctx, 1)
		t.Assert(v, 11)
	})
}

func TestCache_GetOrSetFuncLock(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		cache := gcache.X创建()
		cache.X取值或设置值_并发安全函数(ctx, 1, func(ctx context.Context) (value interface{}, err error) {
			return 11, nil
		}, 0)
		v, _ := cache.X取值(ctx, 1)
		t.Assert(v, 11)

		cache.X取值或设置值_并发安全函数(ctx, 1, func(ctx context.Context) (value interface{}, err error) {
			return 111, nil
		}, 0)
		v, _ = cache.X取值(ctx, 1)
		t.Assert(v, 11)

		gcache.X删除并带返回值(ctx, g.Slice别名{1, 2, 3}...)
		gcache.X取值或设置值_并发安全函数(ctx, 1, func(ctx context.Context) (value interface{}, err error) {
			return 11, nil
		}, 0)
		v, _ = cache.X取值(ctx, 1)
		t.Assert(v, 11)

		gcache.X取值或设置值_并发安全函数(ctx, 1, func(ctx context.Context) (value interface{}, err error) {
			return 111, nil
		}, 0)
		v, _ = cache.X取值(ctx, 1)
		t.Assert(v, 11)
	})
}

func TestCache_Clear(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		cache := gcache.X创建()
		cache.X设置Map(ctx, g.MapAnyAny{1: 11, 2: 22}, 0)
		cache.X清空(ctx)
		n, _ := cache.X取数量(ctx)
		t.Assert(n, 0)
	})
}

func TestCache_SetConcurrency(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		cache := gcache.X创建()
		pool := grpool.New(4)
		go func() {
			for {
				pool.Add(ctx, func(ctx context.Context) {
					cache.X设置值并跳过已存在(ctx, 1, 11, 10)
				})
			}
		}()
		select {
		case <-time.After(2 * time.Second):
									// t.Log("第一部分结束"). md5:a52c91ecebe84cfc
		}

		go func() {
			for {
				pool.Add(ctx, func(ctx context.Context) {
					cache.X设置值并跳过已存在(ctx, 1, nil, 10)
				})
			}
		}()
		select {
		case <-time.After(2 * time.Second):
									// t.Log("第二部分结束"). md5:9a0317b9c13ab35a
		}
	})
}

func TestCache_Basic(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		{
			cache := gcache.X创建()
			cache.X设置Map(ctx, g.MapAnyAny{1: 11, 2: 22}, 0)
			b, _ := cache.X是否存在(ctx, 1)
			t.Assert(b, true)
			v, _ := cache.X取值(ctx, 1)
			t.Assert(v, 11)
			data, _ := cache.X取所有键值Map副本(ctx)
			t.Assert(data[1], 11)
			t.Assert(data[2], 22)
			t.Assert(data[3], nil)
			n, _ := cache.X取数量(ctx)
			t.Assert(n, 2)
			keys, _ := cache.X取所有键(ctx)
			t.Assert(gset.X创建并按值(g.Slice别名{1, 2}).X是否相等(gset.X创建并按值(keys)), true)
			keyStrs, _ := cache.X取所有键文本(ctx)
			t.Assert(gset.X创建并按值(g.Slice别名{"1", "2"}).X是否相等(gset.X创建并按值(keyStrs)), true)
			values, _ := cache.X取所有值(ctx)
			t.Assert(gset.X创建并按值(g.Slice别名{11, 22}).X是否相等(gset.X创建并按值(values)), true)
			removeData1, _ := cache.X删除并带返回值(ctx, 1)
			t.Assert(removeData1, 11)
			n, _ = cache.X取数量(ctx)
			t.Assert(n, 1)

			cache.X删除并带返回值(ctx, 2)
			n, _ = cache.X取数量(ctx)
			t.Assert(n, 0)
		}

		gcache.X删除并带返回值(ctx, g.Slice别名{1, 2, 3}...)
		{
			gcache.X设置Map(ctx, g.MapAnyAny{1: 11, 2: 22}, 0)
			b, _ := gcache.X是否存在(ctx, 1)
			t.Assert(b, true)
			v, _ := gcache.X取值(ctx, 1)
			t.Assert(v, 11)
			data, _ := gcache.X取所有键值Map副本(ctx)
			t.Assert(data[1], 11)
			t.Assert(data[2], 22)
			t.Assert(data[3], nil)
			n, _ := gcache.X取数量(ctx)
			t.Assert(n, 2)
			keys, _ := gcache.X取所有键(ctx)
			t.Assert(gset.X创建并按值(g.Slice别名{1, 2}).X是否相等(gset.X创建并按值(keys)), true)
			keyStrs, _ := gcache.X取所有键文本(ctx)
			t.Assert(gset.X创建并按值(g.Slice别名{"1", "2"}).X是否相等(gset.X创建并按值(keyStrs)), true)
			values, _ := gcache.X取所有值(ctx)
			t.Assert(gset.X创建并按值(g.Slice别名{11, 22}).X是否相等(gset.X创建并按值(values)), true)
			removeData1, _ := gcache.X删除并带返回值(ctx, 1)
			t.Assert(removeData1, 11)
			n, _ = gcache.X取数量(ctx)
			t.Assert(n, 1)
			gcache.X删除并带返回值(ctx, 2)
			n, _ = gcache.X取数量(ctx)
			t.Assert(n, 0)
		}
	})
}

func TestCache_Removes(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		cache := gcache.X创建()
		t.AssertNil(cache.X设置值(ctx, 1, 11, 0))
		t.AssertNil(cache.X设置值(ctx, 2, 22, 0))
		t.AssertNil(cache.X设置值(ctx, 3, 33, 0))
		t.AssertNil(cache.X删除(ctx, g.Slice别名{2, 3}))

		ok, err := cache.X是否存在(ctx, 1)
		t.AssertNil(err)
		t.Assert(ok, true)

		ok, err = cache.X是否存在(ctx, 2)
		t.AssertNil(err)
		t.Assert(ok, false)
	})

	gtest.C(t, func(t *gtest.T) {
		t.AssertNil(gcache.X设置值(ctx, 1, 11, 0))
		t.AssertNil(gcache.X设置值(ctx, 2, 22, 0))
		t.AssertNil(gcache.X设置值(ctx, 3, 33, 0))
		t.AssertNil(gcache.X删除(ctx, g.Slice别名{2, 3}))

		ok, err := gcache.X是否存在(ctx, 1)
		t.AssertNil(err)
		t.Assert(ok, true)

		ok, err = gcache.X是否存在(ctx, 2)
		t.AssertNil(err)
		t.Assert(ok, false)
	})
}

func TestCache_Basic_Must(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		defer gcache.X删除并带返回值(ctx, g.Slice别名{1, 2, 3, 4}...)

		t.AssertNil(gcache.X设置值(ctx, 1, 11, 0))
		v := gcache.X取值PANI(ctx, 1)
		t.Assert(v, 11)
		gcache.X取值或设置值PANI(ctx, 2, 22, 0)
		v = gcache.X取值PANI(ctx, 2)
		t.Assert(v, 22)

		gcache.X取值或设置值_函数PANI(ctx, 3, func(ctx context.Context) (value interface{}, err error) {
			return 33, nil
		}, 0)
		v = gcache.X取值PANI(ctx, 3)
		t.Assert(v, 33)

		gcache.X取值或设置值_并发安全函数(ctx, 4, func(ctx context.Context) (value interface{}, err error) {
			return 44, nil
		}, 0)
		v = gcache.X取值PANI(ctx, 4)
		t.Assert(v, 44)

		t.Assert(gcache.X是否存在PANI(ctx, 1), true)

		t.AssertNil(gcache.X设置值(ctx, 1, 11, 3*time.Second))
		expire := gcache.X取过期时间PANI(ctx, 1)
		t.AssertGE(expire, 0)

		n := gcache.X取数量PANI(ctx)
		t.Assert(n, 4)

		data := gcache.X取所有键值Map副本PANI(ctx)
		t.Assert(len(data), 4)

		keys := gcache.X取所有键PANI(ctx)
		t.Assert(len(keys), 4)

		keyStrings := gcache.X取所有键文本PANI(ctx)
		t.Assert(len(keyStrings), 4)

		values := gcache.X取所有值PANI(ctx)
		t.Assert(len(values), 4)
	})
}

func TestCache_NewWithAdapter(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		cache := gcache.X创建并按适配器(gcache.X创建内存适配器())
		t.AssertNE(cache, nil)
	})
}
