// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包gmeta为struct提供了嵌入式元数据功能。. md5:31c7b2b4ae02619a
package gmeta

import (
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gstructs"
)

// Meta 作为结构体的嵌入属性，用于启用元数据功能。. md5:12d03f1f1bd9c041
type Meta struct{}

const (
	metaAttributeName = "Meta"       // metaAttributeName 是结构体中元数据的属性名称。. md5:08a529ac77f54a13
	metaTypeName      = "gmeta.Meta" // metaTypeName 用于类型字符串比较。. md5:00017d3b5ff6e69c
)

// Data 从`object`中检索并返回所有元数据。. md5:d0b2cb45c581d982
func Data(object interface{}) map[string]string {
	reflectType, err := gstructs.StructType(object)
	if err != nil {
		return nil
	}
	if field, ok := reflectType.FieldByName(metaAttributeName); ok {
		if field.Type.String() == metaTypeName {
			return gstructs.ParseTag(string(field.Tag))
		}
	}
	return map[string]string{}
}

// Get通过`key`从`object`中检索并返回指定的元数据。. md5:cb5e4223dc341860
func Get(object interface{}, key string) *gvar.Var {
	v, ok := Data(object)[key]
	if !ok {
		return nil
	}
	return gvar.New(v)
}
