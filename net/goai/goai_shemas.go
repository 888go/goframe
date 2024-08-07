// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package goai

import (
	gmap "github.com/888go/goframe/container/gmap"
)

type Schemas struct {
	refs *gmap.ListMap // map[string]SchemaRef
}

func createSchemas() Schemas {
	return Schemas{
		refs: gmap.X创建链表mp(),
	}
}

func (s *Schemas) init() {
	if s.refs == nil {
		s.refs = gmap.X创建链表mp()
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
