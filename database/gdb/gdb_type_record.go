// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package db类

import (
	"database/sql"
	
	"github.com/888go/goframe/container/gmap"
	"github.com/888go/goframe/encoding/gjson"
	"github.com/888go/goframe/internal/empty"
	"github.com/888go/goframe/util/gconv"
)

// Json 将 `r` 转换为 JSON 格式的内容。
func (r Record) X取json() string {
	content, _ := json类.X创建(r.X取Map()).X取json文本()
	return content
}

// Xml 将`r`转换为XML格式的内容。
func (r Record) X取xml(根标记 ...string) string {
	content, _ := json类.X创建(r.X取Map()).X取xml文本(根标记...)
	return content
}

// Map 将 `r` 转换为 map[string]interface{} 类型。
func (r Record) X取Map() Map {
	m := make(map[string]interface{})
	for k, v := range r {
		m[k] = v.X取值()
	}
	return m
}

// GMap 将 `r` 转换为一个 gmap。
func (r Record) X取Map类() *map类.StrAnyMap {
	return map类.X创建AnyStr并从Map(r.X取Map())
}

// Struct 将 `r` 转换为结构体。
// 注意，参数 `pointer` 应为 *struct 或 **struct 类型。
//
// 注意，如果 `r` 为空，则返回 sql.ErrNoRows 错误。
func (r Record) X取结构体指针(结构体指针 interface{}) error {
	// 如果记录为空，则返回错误。
	if r.X是否为空() {
		if !empty.X是否为Nil(结构体指针, true) {
			return sql.ErrNoRows
		}
		return nil
	}
	return 转换类.StructTag(r, 结构体指针, OrmTagForStruct)
}

// IsEmpty 检查并返回 `r` 是否为空。
func (r Record) X是否为空() bool {
	return len(r) == 0
}
