// 版权所有 (c) GoFrame (https://goframe.org)，保留所有权利。
//
// 本源代码遵循MIT许可协议。若未随此文件分发MIT许可证的副本，
// 您可以从 https://github.com/gogf/gf 获取。 md5:c14c707c81272457

package gmetric

import (
	"bytes"
	"fmt"
	"os"

	"github.com/gogf/gf/v2/internal/json"
	"github.com/gogf/gf/v2/os/gfile"
)

// Attributes 是一个 Attribute 类型的切片。 md5:5e8ed6feb0b054bf
type Attributes []Attribute

// Attribute是Metric的键值对项。 md5:d998f30eea2094b2
type Attribute interface {
	Key() string // 这个属性的键。 md5:0ddb614f69d6d447
	Value() any  // 此属性的值。 md5:855798b766242495
}

// AttributeKey 是属性键。 md5:9221fc74fb7697dc
type AttributeKey string

// Option包含执行度量操作的选项。 md5:1a7865b57252c62c
type Option struct {
	// Attributes保存动态的键值对元数据。 md5:837d5c5300f22ee1
	Attributes Attributes
}

// localAttribute 实现了 Attribute 接口。 md5:f861eb05ab05e971
type localAttribute struct {
	key   string
	value any
}

var (
	hostname    string
	processPath string
)

func init() {
	hostname, _ = os.Hostname()
	processPath = gfile.SelfPath()
}

// CommonAttributes 返回乐器常用的属性。 md5:a3240e3fe755b09a
func CommonAttributes() Attributes {
	return Attributes{
		NewAttribute(`os.host.name`, hostname),
		NewAttribute(`process.path`, processPath),
	}
}

// NewAttribute通过给定的`key`和`value`创建并返回一个Attribute。 md5:dc10aeeb3e9da0df
func NewAttribute(key string, value any) Attribute {
	return &localAttribute{
		key:   key,
		value: value,
	}
}

// Key 返回属性的键。 md5:27fe16d7c522fc43
func (l *localAttribute) Key() string {
	return l.key
}

// Value 返回属性的值。 md5:7b8a05d03d68be89
func (l *localAttribute) Value() any {
	return l.value
}

// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。 md5:43c3b36e60a18f9a
func (l *localAttribute) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"%s":%#v}`, l.key, l.value)), nil
}

// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。 md5:43c3b36e60a18f9a
func (attrs Attributes) String() string {
	bs, _ := attrs.MarshalJSON()
	return string(bs)
}

// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。 md5:43c3b36e60a18f9a
func (attrs Attributes) MarshalJSON() ([]byte, error) {
	var (
		bs     []byte
		err    error
		buffer = bytes.NewBuffer(nil)
	)
	buffer.WriteByte('[')
	for _, attr := range attrs {
		bs, err = json.Marshal(attr)
		if err != nil {
			return nil, err
		}
		if buffer.Len() > 1 {
			buffer.WriteByte(',')
		}
		buffer.Write(bs)
	}
	buffer.WriteByte(']')
	return buffer.Bytes(), nil
}
