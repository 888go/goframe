// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package genv

// MustSet 执行与 Set 相同的操作，但如果发生任何错误，它将引发恐慌。 md5:89753cb5f56f60cc
func MustSet(key, value string) {
	if err := Set(key, value); err != nil {
		panic(err)
	}
}

// MustRemove 的行为与 Remove 相同，但如果发生任何错误，它会直接 panic。 md5:ad4ac7324486398a
func MustRemove(key ...string) {
	if err := Remove(key...); err != nil {
		panic(err)
	}
}
