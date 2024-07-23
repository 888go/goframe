// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gdb

import (
	"database/sql"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/internal/empty"
	"github.com/gogf/gf/v2/util/gconv"
)

// Json 将 `r` 转换为JSON格式的内容。 md5:60a0626b0a333d14
func (r Record) Json() string {
	content, _ := gjson.New(r.Map()).ToJsonString()
	return content
}

// Xml 将 `r` 转换为 XML 格式的内容。 md5:31a335fedb874d26
func (r Record) Xml(rootTag ...string) string {
	content, _ := gjson.New(r.Map()).ToXmlString(rootTag...)
	return content
}

// Map 将 `r` 转换为 map[string]interface{} 类型。 md5:5b4502a5f29602f9
func (r Record) Map() Map {
	m := make(map[string]interface{})
	for k, v := range r {
		m[k] = v.Val()
	}
	return m
}

// GMap将`r`转换为gmap。 md5:573ff0b484a9573f
func (r Record) GMap() *gmap.StrAnyMap {
	return gmap.NewStrAnyMapFrom(r.Map())
}

// Struct 将 `r` 转换为结构体。
// 注意参数 `pointer` 应为 *struct 或 **struct 类型。
//
// 注意，如果 `r` 为空，它将返回 sql.ErrNoRows。
// md5:9ad6d688dbdddb25
func (r Record) Struct(pointer interface{}) error {
	// 如果记录为空，它将返回错误。 md5:dc39009d7d477d46
	if r.IsEmpty() {
		if !empty.IsNil(pointer, true) {
			return sql.ErrNoRows
		}
		return nil
	}
	return gconv.StructTag(r, pointer, OrmTagForStruct)
}

// IsEmpty 检查 `r` 是否为空，然后返回结果。 md5:4ee28a47e769cceb
func (r Record) IsEmpty() bool {
	return len(r) == 0
}
