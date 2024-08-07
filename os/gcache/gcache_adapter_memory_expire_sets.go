// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 缓存类

import (
	"sync"

	gset "github.com/888go/goframe/container/gset"
)

type adapterMemoryExpireSets struct {
	mu         sync.RWMutex        // expireSetMu 保证expireSets映射的并发安全性。 md5:1a74d61573e21f7e
	expireSets map[int64]*gset.Set // expireSets 是过期时间戳到其键集合的映射，用于快速索引和删除。 md5:d2c25eb345e1ea19
}

func newAdapterMemoryExpireSets() *adapterMemoryExpireSets {
	return &adapterMemoryExpireSets{
		expireSets: make(map[int64]*gset.Set),
	}
}

func (d *adapterMemoryExpireSets) X取值(key int64) (result *gset.Set) {
	d.mu.RLock()
	result = d.expireSets[key]
	d.mu.RUnlock()
	return
}

func (d *adapterMemoryExpireSets) GetOrNew(key int64) (result *gset.Set) {
	if result = d.X取值(key); result != nil {
		return
	}
	d.mu.Lock()
	if es, ok := d.expireSets[key]; ok {
		result = es
	} else {
		result = gset.X创建(true)
		d.expireSets[key] = result
	}
	d.mu.Unlock()
	return
}

func (d *adapterMemoryExpireSets) Delete(key int64) {
	d.mu.Lock()
	delete(d.expireSets, key)
	d.mu.Unlock()
}
