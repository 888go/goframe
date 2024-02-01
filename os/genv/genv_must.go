// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package genv

// MustSet 表现如同 Set，但当发生任何错误时会触发panic（异常）。
func MustSet(key, value string) {
	if err := Set(key, value); err != nil {
		panic(err)
	}
}

// MustRemove 的行为与 Remove 相同，但是当发生任何错误时，它会触发 panic（异常）。
func MustRemove(key ...string) {
	if err := Remove(key...); err != nil {
		panic(err)
	}
}
