// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gdb

import (
	"database/sql"
	
	"github.com/888go/goframe/container/gmap"
	"github.com/888go/goframe/encoding/gjson"
	"github.com/888go/goframe/internal/empty"
	"github.com/888go/goframe/util/gconv"
)

// Json 将 `r` 转换为 JSON 格式的内容。
func (r Record) Json() string {
	content, _ := gjson.New(r.Map()).ToJsonString()
	return content
}

// Xml 将`r`转换为XML格式的内容。
func (r Record) Xml(rootTag ...string) string {
	content, _ := gjson.New(r.Map()).ToXmlString(rootTag...)
	return content
}

// Map 将 `r` 转换为 map[string]interface{} 类型。
func (r Record) Map() Map {
	m := make(map[string]interface{})
	for k, v := range r {
		m[k] = v.Val()
	}
	return m
}

// GMap 将 `r` 转换为一个 gmap。
func (r Record) GMap() *gmap.StrAnyMap {
	return gmap.NewStrAnyMapFrom(r.Map())
}

// Struct 将 `r` 转换为结构体。
// 注意，参数 `pointer` 应为 *struct 或 **struct 类型。
//
// 注意，如果 `r` 为空，则返回 sql.ErrNoRows 错误。
func (r Record) Struct(pointer interface{}) error {
	// 如果记录为空，则返回错误。
	if r.IsEmpty() {
		if !empty.IsNil(pointer, true) {
			return sql.ErrNoRows
		}
		return nil
	}
	return gconv.StructTag(r, pointer, OrmTagForStruct)
}

// IsEmpty 检查并返回 `r` 是否为空。
func (r Record) IsEmpty() bool {
	return len(r) == 0
}
