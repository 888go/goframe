// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package goai

import (
	"github.com/888go/goframe/container/gmap"
)

type Schemas struct {
	refs *map类.ListMap // map[string]SchemaRef 是一个Go语言中的映射类型，键为字符串(string)，值为SchemaRef类型。
// 这个映射可以用来存储一组键为字符串，值为SchemaRef的键值对数据，通常用于根据字符串标识符查找对应的SchemaRef结构体实例。
}

func createSchemas() Schemas {
	return Schemas{
		refs: map类.X创建链表mp(),
	}
}

func (s *Schemas) init() {
	if s.refs == nil {
		s.refs = map类.X创建链表mp()
	}
}

func (s *Schemas) Clone() Schemas {
	newSchemas := createSchemas()
	newSchemas.refs = s.refs.X取副本()
	return newSchemas
}

func (s *Schemas) Get(name string) *SchemaRef {
	s.init()
	value := s.refs.X取值(name)
	if value != nil {
		ref := value.(SchemaRef)
		return &ref
	}
	return nil
}

func (s *Schemas) X设置值(name string, ref SchemaRef) {
	s.init()
	s.refs.X设置值(name, ref)
}

func (s *Schemas) Removes(names []interface{}) {
	s.init()
	s.refs.X删除多个值(names)
}

func (s *Schemas) Map() map[string]SchemaRef {
	s.init()
	m := make(map[string]SchemaRef)
	s.refs.X遍历(func(key, value interface{}) bool {
		m[key.(string)] = value.(SchemaRef)
		return true
	})
	return m
}

func (s *Schemas) X遍历(f func(key string, ref SchemaRef) bool) {
	s.init()
	s.refs.X遍历(func(key, value interface{}) bool {
		return f(key.(string), value.(SchemaRef))
	})
}

func (s Schemas) MarshalJSON() ([]byte, error) {
	s.init()
	return s.refs.MarshalJSON()
}
