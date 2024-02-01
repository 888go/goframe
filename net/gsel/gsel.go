// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gsel 提供选择器定义及其实现。
package gsel
import (
	"context"
	
	"github.com/888go/goframe/net/gsvc"
	)
// Builder在运行时创建并返回选择器。
type Builder interface {
	Name() string
	Build() Selector
}

// 服务负载均衡器选择器。
type Selector interface {
	// Pick 选择并返回服务。
	Pick(ctx context.Context) (node Node, done DoneFunc, err error)

	// Update 将服务更新到Selector中。
	Update(ctx context.Context, nodes Nodes) error
}

// Node 是节点接口。
type Node interface {
	Service() gsvc.Service
	Address() string
}

// Nodes 包含多个 Node。
type Nodes []Node

// DoneFunc 是当 RPC 调用完成时的回调函数。
type DoneFunc func(ctx context.Context, di DoneInfo)

// DoneInfo 包含 done 的附加信息。
type DoneInfo struct {
	// Err 是 RPC 结束时的错误，可能为 nil。
	Err error

	// Trailer 包含了RPC的尾部元数据（如果存在的话）。
	Trailer DoneInfoMD

	// BytesSent 表示是否已向服务器发送了任何字节。
	BytesSent bool

	// BytesReceived 表示是否已从服务器接收到任何字节。
	BytesReceived bool

// ServerLoad表示从服务器接收到的负载信息。通常，它作为
// 附属元数据的一部分进行发送。
//
// 当前唯一支持的数据类型是*orca_v1.LoadReport。
	ServerLoad interface{}
}

// DoneInfoMD 是一个从元数据键到值数组的映射。
// 用户应使用以下两个便捷函数 New 和 Pairs 来生成 MD。
type DoneInfoMD interface {
	// Len 返回 md 中的项目数量。
	Len() int

// Get 方法用于获取给定键对应的值。
//
// 在md中搜索之前，会将k转换为小写。
	Get(k string) []string

// Set 用于设置给定键的值为一个切片类型的值。
//
// 在存储到 md 之前，k 会被转换为小写。
	Set(key string, values ...string)

// Append 将值添加到键 k，但不会覆盖该键已存储的内容。
//
// 在存储到 md 之前，会将 k 转换为小写。
	Append(k string, values ...string)

// Delete 删除给定键 k 的值，在从 md 中移除前，会将键 k 转换为小写。
	Delete(k string)
}

// String 方法格式化并以字符串形式返回 Nodes。
func (ns Nodes) String() string {
	var s string
	for _, node := range ns {
		if s != "" {
			s += ","
		}
		s += node.Address()
	}
	return s
}
