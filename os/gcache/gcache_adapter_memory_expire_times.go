// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 缓存类

import (
	"sync"
)

type adapterMemoryExpireTimes struct {
	mu          sync.RWMutex          // expireTimeMu 确保 expireTimes 映射表在并发环境下的安全性。
	expireTimes map[interface{}]int64 // expireTimes 是一个过期键与其时间戳的映射，用于快速索引和删除。
}

func newAdapterMemoryExpireTimes() *adapterMemoryExpireTimes {
	return &adapterMemoryExpireTimes{
		expireTimes: make(map[interface{}]int64),
	}
}

func (d *adapterMemoryExpireTimes) X取值(key interface{}) (value int64) {
	d.mu.RLock()
	value = d.expireTimes[key]
	d.mu.RUnlock()
	return
}

func (d *adapterMemoryExpireTimes) X设置值(key interface{}, value int64) {
	d.mu.Lock()
	d.expireTimes[key] = value
	d.mu.Unlock()
}

func (d *adapterMemoryExpireTimes) Delete(key interface{}) {
	d.mu.Lock()
	delete(d.expireTimes, key)
	d.mu.Unlock()
}
