// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gdb

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
func (m *Model) Order(orderBy ...interface{}) *Model {
	if len(orderBy) == 0 {
		return m
	}
	model := m.getModel()
	if model.orderBy != "" {
		model.orderBy += ","
	}
	for _, v := range orderBy {
		switch v.(type) {
		case Raw, *Raw:
			model.orderBy += gconv.String(v)
			return model
		}
	}
	model.orderBy += model.db.GetCore().QuoteString(gstr.JoinAny(orderBy, " "))
	return model
}

// OrderAsc 为模型设置 "ORDER BY xxx ASC" 语句。
func (m *Model) OrderAsc(column string) *Model {
	if len(column) == 0 {
		return m
	}
	return m.Order(column + " ASC")
}

// OrderDesc 为模型设置 "ORDER BY xxx DESC" 语句。
func (m *Model) OrderDesc(column string) *Model {
	if len(column) == 0 {
		return m
	}
	return m.Order(column + " DESC")
}

// OrderRandom 为模型设置 "ORDER BY RANDOM()" 语句。
func (m *Model) OrderRandom() *Model {
	model := m.getModel()
	model.orderBy = "RAND()"
	return model
}

// Group 设置模型的 "GROUP BY" 语句。
func (m *Model) Group(groupBy ...string) *Model {
	if len(groupBy) == 0 {
		return m
	}
	model := m.getModel()
	if model.groupBy != "" {
		model.groupBy += ","
	}
	model.groupBy += model.db.GetCore().QuoteString(strings.Join(groupBy, ","))
	return model
}
