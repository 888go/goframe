// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package db类

import (
	"database/sql"

	gmap "github.com/888go/goframe/container/gmap"
	gjson "github.com/888go/goframe/encoding/gjson"
	"github.com/888go/goframe/internal/empty"
	gconv "github.com/888go/goframe/util/gconv"
)

// X取json 将 `r` 转换为JSON格式的内容。 md5:60a0626b0a333d14
func (r Record) X取json() string {
	content, _ := gjson.X创建(r.X取Map()).X取json文本()
	return content
}

// X取xml 将 `r` 转换为 XML 格式的内容。 md5:31a335fedb874d26
func (r Record) X取xml(根标记 ...string) string {
	content, _ := gjson.X创建(r.X取Map()).X取xml文本(根标记...)
	return content
}

// X取Map 将 `r` 转换为 map[string]interface{} 类型。 md5:5b4502a5f29602f9
func (r Record) X取Map() Map {
	m := make(map[string]interface{})
	for k, v := range r {
		m[k] = v.X取值()
	}
	return m
}

// X取Map类将`r`转换为gmap。 md5:573ff0b484a9573f
func (r Record) X取Map类() *gmap.StrAnyMap {
	return gmap.X创建AnyStr并从Map(r.X取Map())
}

// X取结构体指针 将 `r` 转换为结构体。
// 注意参数 `pointer` 应为 *struct 或 **struct 类型。
//
// 注意，如果 `r` 为空，它将返回 sql.ErrNoRows。
// md5:9ad6d688dbdddb25
func (r Record) X取结构体指针(结构体指针 interface{}) error {
		// 如果记录为空，它将返回错误。 md5:dc39009d7d477d46
	if r.X是否为空() {
		if !empty.IsNil(结构体指针, true) {
			return sql.ErrNoRows
		}
		return nil
	}
	return gconv.StructTag(r, 结构体指针, OrmTagForStruct)
}

// X是否为空 检查 `r` 是否为空，然后返回结果。 md5:4ee28a47e769cceb
func (r Record) X是否为空() bool {
	return len(r) == 0
}
