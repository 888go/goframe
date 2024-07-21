// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gstructs

// Signature 返回此类型的一个唯一字符串。 md5:cba9a7124f69dbaa
// ff:
// t:
func (t Type) Signature() string {
	return t.PkgPath() + "/" + t.String()
}

// FieldKeys 返回当前结构体/映射的键。 md5:6361c05b8d9fe2e0
// ff:
// t:
func (t Type) FieldKeys() []string {
	keys := make([]string, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		keys[i] = t.Field(i).Name
	}
	return keys
}
