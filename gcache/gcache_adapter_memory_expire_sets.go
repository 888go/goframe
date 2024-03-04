// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gcache

import (
	"sync"
	
	"github.com/gogf/gf/v2/container/gset"
)

type adapterMemoryExpireSets struct {
	mu         sync.RWMutex        // expireSetMu 用于确保 expireSets 这个映射（map）在并发环境下的安全性。
	expireSets map[int64]*gset.Set // expireSets 是一个映射表，用于存储即将过期的时间戳及其对应的键集合。这个映射表用于快速索引和删除操作。
}

func newAdapterMemoryExpireSets() *adapterMemoryExpireSets {
	return &adapterMemoryExpireSets{
		expireSets: make(map[int64]*gset.Set),
	}
}

func (d *adapterMemoryExpireSets) Get(key int64) (result *gset.Set) {
	d.mu.RLock()
	result = d.expireSets[key]
	d.mu.RUnlock()
	return
}

func (d *adapterMemoryExpireSets) GetOrNew(key int64) (result *gset.Set) {
	if result = d.Get(key); result != nil {
		return
	}
	d.mu.Lock()
	if es, ok := d.expireSets[key]; ok {
		result = es
	} else {
		result = gset.New(true)
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
