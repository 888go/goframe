// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包 gsel 提供选择器的定义和实现。 md5:009a7633736cc0b4
package gsel

import (
	"context"

	"github.com/gogf/gf/v2/net/gsvc"
)

// Builder 在运行时创建并返回选择器。 md5:ffd073ea24966ab1
type Builder interface {
	Name() string
	Build() Selector
}

// 用于服务均衡器的选择器。 md5:6331f98beb651e36
type Selector interface {
	// Pick 选择并返回服务。 md5:2719544a395c30f1
	Pick(ctx context.Context) (node Node, done DoneFunc, err error)

	// Update 更新服务到Selector中。 md5:8ca4ca3b3f8963c1
	Update(ctx context.Context, nodes Nodes) error
}

// Node 是节点接口。 md5:333117a374d3b386
type Node interface {
	Service() gsvc.Service
	Address() string
}

// Nodes 包含多个 Node。 md5:96478357456f6f66
type Nodes []Node

// DoneFunc 是RPC调用完成时的回调函数。 md5:571d85b6ab17bac9
type DoneFunc func(ctx context.Context, di DoneInfo)

// DoneInfo包含done的额外信息。 md5:fe59ad436c6d2598
type DoneInfo struct {
	// Err 是RPC完成时伴随的错误。它可能为nil（即没有错误）。 md5:6f92d17caab8ccb5
	Err error

	// Trailer 包含了如果存在的话，RPC调用的追踪元数据。 md5:291f630111ba0776
	Trailer DoneInfoMD

	// BytesSent 表示是否已将任何字节发送到服务器。 md5:d9e7a45a59ef93ae
	BytesSent bool

	// BytesReceived表示是否已从服务器接收任何字节。 md5:76160cbbde4b2264
	BytesReceived bool

// ServerLoad 是从服务器接收到的负载。它通常作为尾部元数据的一部分发送。
//
// 目前唯一支持的类型是 *orca_v1.LoadReport。
// md5:631ea6798c4714cb
	ServerLoad interface{}
}

// DoneInfoMD 是一个从元数据键映射到值数组的映射。
// 用户应使用以下两个便利函数 New 和 Pairs 来生成 MD。
// md5:46c19d6d882119ab
type DoneInfoMD interface {
	// Len 返回md中的项目数量。 md5:dd1e6c01d830a7fd
	Len() int

// Get 获取指定键的值。
//
// 在对 md 进行搜索之前，k 会被转换为小写。
// md5:cd83da607a1a524c
	Get(k string) []string

// Set 使用一系列值设置给定键的值。
//
// 在存储到 md 之前，k 将被转换为小写。
// md5:723075d936727645
	Set(key string, values ...string)

// Append 将值添加到键 k，但不会覆盖已经存储在该键上的内容。
//
// 在存储到 md 中之前，k 会被转换为小写。
// md5:8b748588b95754ce
	Append(k string, values ...string)

	// Delete 删除给定键 k 的值，该键在从 md 中删除之前会被转换为小写。
	// md5:fa165ee7e187c245
	Delete(k string)
}

// String 方法将节点格式化并返回为字符串。 md5:8aee26fc061ca942
// ff:
// ns:
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
