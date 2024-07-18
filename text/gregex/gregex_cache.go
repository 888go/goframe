// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gregex

import (
	"regexp"
	"sync"

	"github.com/gogf/gf/v2/errors/gerror"
)

var (
	regexMu = sync.RWMutex{}

// 正则表达式对象的缓存。
// 注意：
// 1. 使用 sync.RWMutex 确保并发安全性。
// 2. 这个映射表中没有过期逻辑。
// md5:645e245ad93c001d
	regexMap = make(map[string]*regexp.Regexp)
)

// getRegexp 使用给定的 `pattern` 返回一个 *regexp.Regexp 对象。
// 它使用缓存来提升正则表达式模式编译的性能，
// 即，对于相同的正则表达式模式，它会返回同一个 *regexp.Regexp 对象。
// 
// 它是多线程安全的，适用于多个goroutine。
// md5:7c16df93c3eeb2b1
func getRegexp(pattern string) (regex *regexp.Regexp, err error) {
	// 使用读取锁获取正则表达式对象。 md5:8d1b5f1036b66cce
	regexMu.RLock()
	regex = regexMap[pattern]
	regexMu.RUnlock()
	if regex != nil {
		return
	}
// 如果该模式不存在于缓存中，
// 则编译该模式并创建一个。
// md5:16abd6a4a92df88a
	if regex, err = regexp.Compile(pattern); err != nil {
		err = gerror.Wrapf(err, `regexp.Compile failed for pattern "%s"`, pattern)
		return
	}
	// 使用写入锁缓存结果对象。 md5:4d4db9dbdc7391d7
	regexMu.Lock()
	regexMap[pattern] = regex
	regexMu.Unlock()
	return
}
