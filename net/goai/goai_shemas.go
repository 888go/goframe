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
	refs *gmap.ListMap // map[string]SchemaRef 是一个Go语言中的映射类型，键为字符串(string)，值为SchemaRef类型。
// 这个映射可以用来存储一组键为字符串，值为SchemaRef的键值对数据，通常用于根据字符串标识符查找对应的SchemaRef结构体实例。
}

func createSchemas() Schemas {
	return Schemas{
		refs: gmap.NewListMap(),
	}
}

func (s *Schemas) init() {
	if s.refs == nil {
		s.refs = gmap.NewListMap()
	}
}

func (s *Schemas) Clone() Schemas {
	newSchemas := createSchemas()
	newSchemas.refs = s.refs.Clone()
	return newSchemas
}

func (s *Schemas) Get(name string) *SchemaRef {
	s.init()
	value := s.refs.Get(name)
	if value != nil {
		ref := value.(SchemaRef)
		return &ref
	}
	return nil
}

func (s *Schemas) Set(name string, ref SchemaRef) {
	s.init()
	s.refs.Set(name, ref)
}

func (s *Schemas) Removes(names []interface{}) {
	s.init()
	s.refs.Removes(names)
}

func (s *Schemas) Map() map[string]SchemaRef {
	s.init()
	m := make(map[string]SchemaRef)
	s.refs.Iterator(func(key, value interface{}) bool {
		m[key.(string)] = value.(SchemaRef)
		return true
	})
	return m
}

func (s *Schemas) Iterator(f func(key string, ref SchemaRef) bool) {
	s.init()
	s.refs.Iterator(func(key, value interface{}) bool {
		return f(key.(string), value.(SchemaRef))
	})
}

func (s Schemas) MarshalJSON() ([]byte, error) {
	s.init()
	return s.refs.MarshalJSON()
}
