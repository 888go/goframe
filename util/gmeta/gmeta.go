// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gmeta 提供了为结构体嵌入元数据的功能。
package 元数据类

import (
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/os/gstructs"
)

// Meta 用于作为结构体的嵌入属性，以启用元数据功能。
type Meta struct{}

const (
	metaAttributeName = "Meta"       // metaAttributeName 是结构体中元数据的属性名称。
	metaTypeName      = "元数据类.Meta" // metaTypeName 用于类型字符串的比较。
)

// Data 从`object`获取并返回所有元数据。
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

// Get 通过 `key` 从 `object` 中获取并返回指定的元数据。
func Get(object interface{}, key string) *泛型类.Var {
	v, ok := Data(object)[key]
	if !ok {
		return nil
	}
	return 泛型类.X创建(v)
}
