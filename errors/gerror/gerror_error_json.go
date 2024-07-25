// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gerror

// MarshalJSON实现了json.Marshal接口的MarshalJSON方法。
// 注意，这里不要使用指针作为接收者。 md5:b4f76062b07a5263
func (err Error) MarshalJSON() ([]byte, error) {
	return []byte(`"` + err.Error() + `"`), nil
}
