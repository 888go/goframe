// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gregex

import (
	"regexp"
	"sync"
	
	"github.com/gogf/gf/v2/errors/gerror"
)

var (
	regexMu = sync.RWMutex{}

// 正则表达式对象缓存.
// 注意:
// 1. 使用 sync.RWMutex 确保了并发安全性.
// 2. 该映射没有设置过期逻辑.
// 以下是详细翻译：
// ```go
// 定义一个正则表达式对象的缓存.
// 需要注意以下两点：
// 1. 该缓存使用了 sync.RWMutex，这意味着在多线程或并发环境下，对缓存的操作是线程安全的。
// 2. 这个映射（map）并未实现任何关于缓存项过期的逻辑，也就是说，一旦有项被添加到缓存中，它将不会自动移除或失效。
	regexMap = make(map[string]*regexp.Regexp)
)

// getRegexp 根据给定的 `pattern` 返回 *regexp.Regexp 对象。
// 它使用缓存来提升正则表达式模式编译的性能，
// 这意味着，对于相同的正则表达式模式，它将返回同一个 *regexp.Regexp 对象。
//
// 对于多个goroutine，它是线程安全的。
func getRegexp(pattern string) (regex *regexp.Regexp, err error) {
	// 使用读取锁获取正则表达式对象。
	regexMu.RLock()
	regex = regexMap[pattern]
	regexMu.RUnlock()
	if regex != nil {
		return
	}
// 如果模式不在缓存中，
// 则编译该模式并创建一个。
	if regex, err = regexp.Compile(pattern); err != nil {
		err = gerror.Wrapf(err, `regexp.Compile failed for pattern "%s"`, pattern)
		return
	}
	// 使用写入锁缓存结果对象。
	regexMu.Lock()
	regexMap[pattern] = regex
	regexMu.Unlock()
	return
}
