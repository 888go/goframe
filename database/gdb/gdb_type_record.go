// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gdb

import (
	"database/sql"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/internal/empty"
	"github.com/gogf/gf/v2/util/gconv"
)

// Json converts `r` to JSON format content.

// ff:取json
func (r Record) Json() string {
	content, _ := gjson.New(r.Map()).ToJsonString()
	return content
}

// Xml converts `r` to XML format content.

// ff:取xml
// rootTag:根标记
func (r Record) Xml(rootTag ...string) string {
	content, _ := gjson.New(r.Map()).ToXmlString(rootTag...)
	return content
}

// Map converts `r` to map[string]interface{}.

// ff:取Map
func (r Record) Map() Map {
	m := make(map[string]interface{})
	for k, v := range r {
		m[k] = v.Val()
	}
	return m
}

// GMap converts `r` to a gmap.

// ff:取Map类
func (r Record) GMap() *gmap.StrAnyMap {
	return gmap.NewStrAnyMapFrom(r.Map())
}

// Struct converts `r` to a struct.
// Note that the parameter `pointer` should be type of *struct/**struct.
//
// Note that it returns sql.ErrNoRows if `r` is empty.

// ff:取结构体指针
// pointer:结构体指针
func (r Record) Struct(pointer interface{}) error {
	// If the record is empty, it returns error.
	if r.IsEmpty() {
		if !empty.IsNil(pointer, true) {
			return sql.ErrNoRows
		}
		return nil
	}
	return gconv.StructTag(r, pointer, OrmTagForStruct)
}

// IsEmpty checks and returns whether `r` is empty.

// ff:是否为空
func (r Record) IsEmpty() bool {
	return len(r) == 0
}
