// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试所有.go文件，并执行基准测试（-bench=".*"），同时显示内存使用情况统计（-benchmem）

package 缓存类_test

import (
	"context"
	"math"
	"testing"
	"time"
	
	"github.com/888go/goframe/container/gset"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/os/gcache"
	"github.com/888go/goframe/os/grpool"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/guid"
)

var (
	ctx = context.Background()
)

func TestCache_GCache_Set(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertNil(缓存类.X设置值(ctx, 1, 11, 0))
		defer 缓存类.X删除并带返回值(ctx, g.Slice别名{1, 2, 3}...)
		v, _ := 缓存类.X取值(ctx, 1)
		t.Assert(v, 11)
		b, _ := 缓存类.X是否存在(ctx, 1)
		t.Assert(b, true)
	})
}

func TestCache_Set(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		c := 缓存类.X创建()
		defer c.X关闭(ctx)
		t.Assert(c.X设置值(ctx, 1, 11, 0), nil)
		v, _ := c.X取值(ctx, 1)
		t.Assert(v, 11)
		b, _ := c.X是否存在(ctx, 1)
		t.Assert(b, true)
	})
}

func TestCache_Set_Expire(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		cache := 缓存类.X创建()
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

	单元测试类.C(t, func(t *单元测试类.T) {
		cache := 缓存类.X创建()
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
	单元测试类.C(t, func(t *单元测试类.T) {
		key := uid类.X生成()
		t.AssertNil(缓存类.X设置值(ctx, key, 11, 3*time.Second))
		expire1, _ := 缓存类.X取过期时间(ctx, key)
		oldValue, exist, err := 缓存类.X更新值(ctx, key, 12)
		t.AssertNil(err)
		t.Assert(oldValue, 11)
		t.Assert(exist, true)

		expire2, _ := 缓存类.X取过期时间(ctx, key)
		v, _ := 缓存类.X取值(ctx, key)
		t.Assert(v, 12)
		t.Assert(math.Ceil(expire1.Seconds()), math.Ceil(expire2.Seconds()))
	})
	// gcache.Cache
	单元测试类.C(t, func(t *单元测试类.T) {
		cache := 缓存类.X创建()
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
	单元测试类.C(t, func(t *单元测试类.T) {
		key := uid类.X生成()
		t.AssertNil(缓存类.X设置值(ctx, key, 11, 3*time.Second))
		defer 缓存类.X删除并带返回值(ctx, key)
		oldExpire, _ := 缓存类.X取过期时间(ctx, key)
		newExpire := 10 * time.Second
		oldExpire2, err := 缓存类.X更新过期时间(ctx, key, newExpire)
		t.AssertNil(err)
		t.AssertIN(oldExpire2, g.Slice别名{oldExpire, `2.999s`})

		e, _ := 缓存类.X取过期时间(ctx, key)
		t.AssertNE(e, oldExpire)
		e, _ = 缓存类.X取过期时间(ctx, key)
		t.Assert(math.Ceil(e.Seconds()), 10)
	})
	// gcache.Cache
	单元测试类.C(t, func(t *单元测试类.T) {
		cache := 缓存类.X创建()
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
	单元测试类.C(t, func(t *单元测试类.T) {
		c := 缓存类.X创建()
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
	单元测试类.C(t, func(t *单元测试类.T) {
		cache := 缓存类.X创建(2)
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
	单元测试类.C(t, func(t *单元测试类.T) {
		cache := 缓存类.X创建(2)
		t.Assert(cache.X设置值(ctx, 1, nil, 1000), nil)
		n, _ := cache.X取数量(ctx)
		t.Assert(n, 1)
		v, _ := cache.X取值(ctx, 1)

		t.Assert(v, nil)
	})
}

func TestCache_SetIfNotExist(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		cache := 缓存类.X创建()
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

		缓存类.X删除并带返回值(ctx, g.Slice别名{1, 2, 3}...)
		ok, err = 缓存类.X设置值并跳过已存在(ctx, 1, 11, 0)
		t.AssertNil(err)
		t.Assert(ok, true)

		v, _ = 缓存类.X取值(ctx, 1)
		t.Assert(v, 11)

		ok, err = 缓存类.X设置值并跳过已存在(ctx, 1, 22, 0)
		t.AssertNil(err)
		t.Assert(ok, false)

		v, _ = 缓存类.X取值(ctx, 1)
		t.Assert(v, 11)
	})
}

func TestCache_SetIfNotExistFunc(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		cache := 缓存类.X创建()
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
	单元测试类.C(t, func(t *单元测试类.T) {
		缓存类.X删除并带返回值(ctx, g.Slice别名{1, 2, 3}...)

		ok, err := 缓存类.X设置值并跳过已存在_函数(ctx, 1, func(ctx context.Context) (value interface{}, err error) {
			return 11, nil
		}, 0)
		t.AssertNil(err)
		t.Assert(ok, true)

		v, _ := 缓存类.X取值(ctx, 1)
		t.Assert(v, 11)

		ok, err = 缓存类.X设置值并跳过已存在_函数(ctx, 1, func(ctx context.Context) (value interface{}, err error) {
			return 22, nil
		}, 0)
		t.AssertNil(err)
		t.Assert(ok, false)

		v, _ = 缓存类.X取值(ctx, 1)
		t.Assert(v, 11)
	})
}

func TestCache_SetIfNotExistFuncLock(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		cache := 缓存类.X创建()
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
	单元测试类.C(t, func(t *单元测试类.T) {
		缓存类.X删除并带返回值(ctx, g.Slice别名{1, 2, 3}...)

		exist, err := 缓存类.X设置值并跳过已存在_并发安全函数(ctx, 1, func(ctx context.Context) (value interface{}, err error) {
			return 11, nil
		}, 0)
		t.AssertNil(err)
		t.Assert(exist, true)

		v, _ := 缓存类.X取值(ctx, 1)
		t.Assert(v, 11)

		exist, err = 缓存类.X设置值并跳过已存在_并发安全函数(ctx, 1, func(ctx context.Context) (value interface{}, err error) {
			return 22, nil
		}, 0)
		t.AssertNil(err)
		t.Assert(exist, false)

		v, _ = 缓存类.X取值(ctx, 1)
		t.Assert(v, 11)
	})
}

func TestCache_SetMap(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		cache := 缓存类.X创建()
		t.AssertNil(cache.X设置Map(ctx, g.MapAnyAny{1: 11, 2: 22}, 0))
		v, _ := cache.X取值(ctx, 1)
		t.Assert(v, 11)

		缓存类.X删除并带返回值(ctx, g.Slice别名{1, 2, 3}...)
		t.AssertNil(缓存类.X设置Map(ctx, g.MapAnyAny{1: 11, 2: 22}, 0))
		v, _ = cache.X取值(ctx, 1)
		t.Assert(v, 11)
	})
}

func TestCache_GetOrSet(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		cache := 缓存类.X创建()
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

	单元测试类.C(t, func(t *单元测试类.T) {
		缓存类.X删除并带返回值(ctx, g.Slice别名{1, 2, 3}...)
		value, err := 缓存类.X取值或设置值(ctx, 1, 11, 0)
		t.AssertNil(err)
		t.Assert(value, 11)

		v, err := 缓存类.X取值(ctx, 1)
		t.AssertNil(err)
		t.Assert(v, 11)

		value, err = 缓存类.X取值或设置值(ctx, 1, 111, 0)
		t.AssertNil(err)
		t.Assert(value, 11)

		v, err = 缓存类.X取值(ctx, 1)
		t.AssertNil(err)
		t.Assert(v, 11)
	})
}

func TestCache_GetOrSetFunc(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		cache := 缓存类.X创建()
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

		缓存类.X删除并带返回值(ctx, g.Slice别名{1, 2, 3}...)

		缓存类.X取值或设置值_函数(ctx, 1, func(ctx context.Context) (value interface{}, err error) {
			return 11, nil
		}, 0)
		v, _ = cache.X取值(ctx, 1)
		t.Assert(v, 11)

		缓存类.X取值或设置值_函数(ctx, 1, func(ctx context.Context) (value interface{}, err error) {
			return 111, nil
		}, 0)
		v, _ = cache.X取值(ctx, 1)
		t.Assert(v, 11)
	})
}

func TestCache_GetOrSetFuncLock(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		cache := 缓存类.X创建()
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

		缓存类.X删除并带返回值(ctx, g.Slice别名{1, 2, 3}...)
		缓存类.X取值或设置值_并发安全函数(ctx, 1, func(ctx context.Context) (value interface{}, err error) {
			return 11, nil
		}, 0)
		v, _ = cache.X取值(ctx, 1)
		t.Assert(v, 11)

		缓存类.X取值或设置值_并发安全函数(ctx, 1, func(ctx context.Context) (value interface{}, err error) {
			return 111, nil
		}, 0)
		v, _ = cache.X取值(ctx, 1)
		t.Assert(v, 11)
	})
}

func TestCache_Clear(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		cache := 缓存类.X创建()
		cache.X设置Map(ctx, g.MapAnyAny{1: 11, 2: 22}, 0)
		cache.X清空(ctx)
		n, _ := cache.X取数量(ctx)
		t.Assert(n, 0)
	})
}

func TestCache_SetConcurrency(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		cache := 缓存类.X创建()
		pool := 协程类.New(4)
		go func() {
			for {
				pool.Add(ctx, func(ctx context.Context) {
					cache.X设置值并跳过已存在(ctx, 1, 11, 10)
				})
			}
		}()
		select {
		case <-time.After(2 * time.Second):
			// t.Log("第一部分结束")
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
			// t.Log("第二部分结束")
		}
	})
}

func TestCache_Basic(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		{
			cache := 缓存类.X创建()
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
			t.Assert(集合类.X创建并按值(g.Slice别名{1, 2}).X是否相等(集合类.X创建并按值(keys)), true)
			keyStrs, _ := cache.X取所有键文本(ctx)
			t.Assert(集合类.X创建并按值(g.Slice别名{"1", "2"}).X是否相等(集合类.X创建并按值(keyStrs)), true)
			values, _ := cache.X取所有值(ctx)
			t.Assert(集合类.X创建并按值(g.Slice别名{11, 22}).X是否相等(集合类.X创建并按值(values)), true)
			removeData1, _ := cache.X删除并带返回值(ctx, 1)
			t.Assert(removeData1, 11)
			n, _ = cache.X取数量(ctx)
			t.Assert(n, 1)

			cache.X删除并带返回值(ctx, 2)
			n, _ = cache.X取数量(ctx)
			t.Assert(n, 0)
		}

		缓存类.X删除并带返回值(ctx, g.Slice别名{1, 2, 3}...)
		{
			缓存类.X设置Map(ctx, g.MapAnyAny{1: 11, 2: 22}, 0)
			b, _ := 缓存类.X是否存在(ctx, 1)
			t.Assert(b, true)
			v, _ := 缓存类.X取值(ctx, 1)
			t.Assert(v, 11)
			data, _ := 缓存类.X取所有键值Map副本(ctx)
			t.Assert(data[1], 11)
			t.Assert(data[2], 22)
			t.Assert(data[3], nil)
			n, _ := 缓存类.X取数量(ctx)
			t.Assert(n, 2)
			keys, _ := 缓存类.X取所有键(ctx)
			t.Assert(集合类.X创建并按值(g.Slice别名{1, 2}).X是否相等(集合类.X创建并按值(keys)), true)
			keyStrs, _ := 缓存类.X取所有键文本(ctx)
			t.Assert(集合类.X创建并按值(g.Slice别名{"1", "2"}).X是否相等(集合类.X创建并按值(keyStrs)), true)
			values, _ := 缓存类.X取所有值(ctx)
			t.Assert(集合类.X创建并按值(g.Slice别名{11, 22}).X是否相等(集合类.X创建并按值(values)), true)
			removeData1, _ := 缓存类.X删除并带返回值(ctx, 1)
			t.Assert(removeData1, 11)
			n, _ = 缓存类.X取数量(ctx)
			t.Assert(n, 1)
			缓存类.X删除并带返回值(ctx, 2)
			n, _ = 缓存类.X取数量(ctx)
			t.Assert(n, 0)
		}
	})
}

func TestCache_Removes(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		cache := 缓存类.X创建()
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

	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertNil(缓存类.X设置值(ctx, 1, 11, 0))
		t.AssertNil(缓存类.X设置值(ctx, 2, 22, 0))
		t.AssertNil(缓存类.X设置值(ctx, 3, 33, 0))
		t.AssertNil(缓存类.X删除(ctx, g.Slice别名{2, 3}))

		ok, err := 缓存类.X是否存在(ctx, 1)
		t.AssertNil(err)
		t.Assert(ok, true)

		ok, err = 缓存类.X是否存在(ctx, 2)
		t.AssertNil(err)
		t.Assert(ok, false)
	})
}

func TestCache_Basic_Must(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		defer 缓存类.X删除并带返回值(ctx, g.Slice别名{1, 2, 3, 4}...)

		t.AssertNil(缓存类.X设置值(ctx, 1, 11, 0))
		v := 缓存类.X取值PANI(ctx, 1)
		t.Assert(v, 11)
		缓存类.X取值或设置值PANI(ctx, 2, 22, 0)
		v = 缓存类.X取值PANI(ctx, 2)
		t.Assert(v, 22)

		缓存类.X取值或设置值_函数PANI(ctx, 3, func(ctx context.Context) (value interface{}, err error) {
			return 33, nil
		}, 0)
		v = 缓存类.X取值PANI(ctx, 3)
		t.Assert(v, 33)

		缓存类.X取值或设置值_并发安全函数(ctx, 4, func(ctx context.Context) (value interface{}, err error) {
			return 44, nil
		}, 0)
		v = 缓存类.X取值PANI(ctx, 4)
		t.Assert(v, 44)

		t.Assert(缓存类.X是否存在PANI(ctx, 1), true)

		t.AssertNil(缓存类.X设置值(ctx, 1, 11, 3*time.Second))
		expire := 缓存类.X取过期时间PANI(ctx, 1)
		t.AssertGE(expire, 0)

		n := 缓存类.X取数量PANI(ctx)
		t.Assert(n, 4)

		data := 缓存类.X取所有键值Map副本PANI(ctx)
		t.Assert(len(data), 4)

		keys := 缓存类.X取所有键PANI(ctx)
		t.Assert(len(keys), 4)

		keyStrings := 缓存类.X取所有键文本PANI(ctx)
		t.Assert(len(keyStrings), 4)

		values := 缓存类.X取所有值PANI(ctx)
		t.Assert(len(values), 4)
	})
}

func TestCache_NewWithAdapter(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		cache := 缓存类.X创建并按适配器(缓存类.X创建内存适配器())
		t.AssertNE(cache, nil)
	})
}
