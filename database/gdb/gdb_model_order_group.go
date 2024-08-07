// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package db类

import (
	"strings"

	gstr "github.com/888go/goframe/text/gstr"
	gconv "github.com/888go/goframe/util/gconv"
)

// X排序 设置模型的 "ORDER BY" 语句。
//
// 示例：
// X排序("id desc")
// X排序("id", "desc").
// X排序("id desc,name asc")
// X排序("id desc").X排序("name asc")
// X排序(gdb.Raw("field(id, 3,1,2)"))
// md5:41ff2d0293c241c6
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
		case Raw, *Raw:
			model.orderBy += gconv.String(v)
			return model
		}
	}
	model.orderBy += model.db.X取Core对象().X底层QuoteString(gstr.X连接Any(字段名与排序方式, " "))
	return model
}

// X排序ASC 设置模型的“ORDER BY xxx ASC”语句。 md5:be417beb10c6b9c1
func (m *Model) X排序ASC(字段名称 string) *Model {
	if len(字段名称) == 0 {
		return m
	}
	return m.X排序(字段名称 + " ASC")
}

// X排序Desc 为模型设置 "ORDER BY xxx DESC" 语句。 md5:ae573bad83990472
func (m *Model) X排序Desc(字段名称 string) *Model {
	if len(字段名称) == 0 {
		return m
	}
	return m.X排序(字段名称 + " DESC")
}

// X排序随机 为模型设置 "ORDER BY RANDOM()" 语句。 md5:e0a71e8f00c3d926
func (m *Model) X排序随机() *Model {
	model := m.getModel()
	model.orderBy = "RAND()"
	return model
}

// X排序分组 设置模型的“GROUP BY”语句。 md5:51d1d1d81a2ab77a
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
