// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gsvc

import (
	"github.com/888go/goframe/container/gvar"
)

// Set 将键值对设置到元数据中。
func (m Metadata) X设置值(key string, value interface{}) {
	m[key] = value
}

// 设置将键值对设置到元数据中。
func (m Metadata) Sets(kvs map[string]interface{}) {
	for k, v := range kvs {
		m[k] = v
	}
}

// Get 函数用于获取并以 gvar 类型返回指定键的值。
func (m Metadata) Get(key string) *泛型类.Var {
	if v, ok := m[key]; ok {
		return 泛型类.X创建(v)
	}
	return nil
}

// IsEmpty 检查并返回当前 Metadata 是否为空。
func (m Metadata) IsEmpty() bool {
	return len(m) == 0
}
