// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package db类

import (
	"strings"
	
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
)

// Order 为模型设置 "ORDER BY" 语句。
//
// 示例：
// Order("id desc") // 按 id 倒序排序
// Order("id", "desc") // 等同于 Order("id desc")
// Order("id desc,name asc") // 先按 id 倒序，再按 name 正序排序
// Order("id desc").Order("name asc") // 分别对 id 和 name 进行倒序和正序排序
// Order(gdb.Raw("field(id, 3,1,2)")) // 使用原生表达式进行排序，如MySQL中的 field 函数指定排序字段的顺序
func (m *Model) X排序(字段名与排序方式 ...interface{}) *Model {
	if len(字段名与排序方式) == 0 {
		return m
	}
	model := m.getModel()
	if model.orderBy != "" {
		model.orderBy += ","
	}
	for _, v := range 字段名与排序方式 {
		switch v.(type) {
		case X原生sql, *X原生sql:
			model.orderBy += 转换类.String(v)
			return model
		}
	}
	model.orderBy += model.db.X取Core对象().X底层QuoteString(文本类.X连接Any(字段名与排序方式, " "))
	return model
}

// OrderAsc 为模型设置 "ORDER BY xxx ASC" 语句。
func (m *Model) X排序ASC(字段名称 string) *Model {
	if len(字段名称) == 0 {
		return m
	}
	return m.X排序(字段名称 + " ASC")
}

// OrderDesc 为模型设置 "ORDER BY xxx DESC" 语句。
func (m *Model) X排序Desc(字段名称 string) *Model {
	if len(字段名称) == 0 {
		return m
	}
	return m.X排序(字段名称 + " DESC")
}

// OrderRandom 为模型设置 "ORDER BY RANDOM()" 语句。
func (m *Model) X排序随机() *Model {
	model := m.getModel()
	model.orderBy = "RAND()"
	return model
}

// Group 设置模型的 "GROUP BY" 语句。
func (m *Model) X排序分组(分组名称 ...string) *Model {
	if len(分组名称) == 0 {
		return m
	}
	model := m.getModel()
	if model.groupBy != "" {
		model.groupBy += ","
	}
	model.groupBy += model.db.X取Core对象().X底层QuoteString(strings.Join(分组名称, ","))
	return model
}
