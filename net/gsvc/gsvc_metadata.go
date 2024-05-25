// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gsvc

import (
	"github.com/gogf/gf/v2/container/gvar"
)

// Set 将键值对设置到元数据中。. md5:3d5704977db787ca
func (m Metadata) Set(key string, value interface{}) {
	m[key] = value
}

// Sets 将键值对设置到元数据中。. md5:5e97d454777313f3
func (m Metadata) Sets(kvs map[string]interface{}) {
	for k, v := range kvs {
		m[k] = v
	}
}

// Get 通过指定的键获取并返回作为gvar的值。. md5:000450be6d0167a3
func (m Metadata) Get(key string) *gvar.Var {
	if v, ok := m[key]; ok {
		return gvar.New(v)
	}
	return nil
}

// IsEmpty 检查当前Metadata是否为空，并返回结果。. md5:0ebeb5f8ed40404a
func (m Metadata) IsEmpty() bool {
	return len(m) == 0
}
