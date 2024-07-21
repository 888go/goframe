// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gdb

import (
	"strings"

	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// Order sets the "ORDER BY" statement for the model.
//
// Order("id desc")
// Order("id", "desc").
// Order("id desc,name asc")
// Order("id desc").Order("name asc")
// Order(gdb.Raw("field(id, 3,1,2)")).
// ff:排序
// m:
// orderBy:字段名与排序方式
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

// OrderAsc 设置模型的“ORDER BY xxx ASC”语句。 md5:be417beb10c6b9c1
// ff:排序ASC
// m:
// column:字段名称
func (m *Model) OrderAsc(column string) *Model {
	if len(column) == 0 {
		return m
	}
	return m.Order(column + " ASC")
}

// OrderDesc 为模型设置 "ORDER BY xxx DESC" 语句。 md5:ae573bad83990472
// ff:排序Desc
// m:
// column:字段名称
func (m *Model) OrderDesc(column string) *Model {
	if len(column) == 0 {
		return m
	}
	return m.Order(column + " DESC")
}

// OrderRandom 为模型设置 "ORDER BY RANDOM()" 语句。 md5:e0a71e8f00c3d926
// ff:排序随机
// m:
func (m *Model) OrderRandom() *Model {
	model := m.getModel()
	model.orderBy = "RAND()"
	return model
}

// Group 设置模型的“GROUP BY”语句。 md5:51d1d1d81a2ab77a
// ff:排序分组
// m:
// groupBy:分组名称
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
