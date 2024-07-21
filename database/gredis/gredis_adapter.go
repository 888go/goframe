// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gredis

import (
	"context"

	"github.com/gogf/gf/v2/container/gvar"
)

// Adapter是通用Redis操作的接口。 md5:9c96b73f93ac5323
type Adapter interface {
	AdapterGroup
	AdapterOperation
}

// AdapterGroup 是一个接口，用于管理 Redis 的组操作。 md5:f603a1b02c295995
type AdapterGroup interface {
	GroupGeneric() IGroupGeneric
	GroupHash() IGroupHash
	GroupList() IGroupList
	GroupPubSub() IGroupPubSub
	GroupScript() IGroupScript
	GroupSet() IGroupSet
	GroupSortedSet() IGroupSortedSet
	GroupString() IGroupString
}

// AdapterOperation 是 Redis 的核心操作函数。
// 这些函数可以被自定义实现轻松覆盖。
// md5:6a3c39d3c764e39e
type AdapterOperation interface {
	// 发送一个命令到服务器并返回接收到的回复。
	// 在将结构体/切片/映射类型的值提交到redis之前，它使用json.Marshal进行编码。
	// md5:5a464ca35e177113
	Do(ctx context.Context, command string, args ...interface{}) (*gvar.Var, error)

	// Conn 获取并返回一个用于连续操作的连接对象。
	// 请注意，如果您不再使用此连接，请手动调用 Close 函数。
	// md5:adf083088afcd372
	Conn(ctx context.Context) (conn Conn, err error)

	// Close 方法关闭当前Redis客户端，关闭其连接池并释放所有相关资源。 md5:bfd91d0269572038
	Close(ctx context.Context) (err error)
}

// Conn 是一个通用 Redis 客户端连接的接口。 md5:75bf8588ab4ad4e1
type Conn interface {
	ConnCommand

	// 发送一个命令到服务器并返回接收到的回复。
	// 在将结构体/切片/映射类型的值提交到redis之前，它使用json.Marshal进行编码。
	// md5:5a464ca35e177113
	Do(ctx context.Context, command string, args ...interface{}) (result *gvar.Var, err error)

	// Close 将连接放回连接池。 md5:7cc2158c987fb9c1
	Close(ctx context.Context) (err error)
}

// ConnCommand是一个接口，用于管理与特定连接绑定的一些操作。 md5:25fa514417ce2230
type ConnCommand interface {
	// Subscribe 将客户端订阅到指定的频道。
	// 参考链接：https:	//redis.io/commands/subscribe/
	// md5:a7414ed1d330bfc7
	Subscribe(ctx context.Context, channel string, channels ...string) ([]*Subscription, error)

	// PSubscribe 将客户端订阅给定的模式。
	//
	// 支持的glob风格模式：
	// - h?llo 订阅hello, hallo和hxllo
	// - h*llo 订阅hllo和heeeello
	// - h[ae]llo 订阅hello和hallo，但不订阅hillo
	//
	// 如果需要匹配特殊字符本身，请使用\进行转义。
	//
	// https:	//redis.io/commands/psubscribe/
	// md5:0bfeb7ebd0d003a7
	PSubscribe(ctx context.Context, pattern string, patterns ...string) ([]*Subscription, error)

	// ReceiveMessage 从Redis服务器接收一个订阅的消息。 md5:dbf6509713a7b2b3
	ReceiveMessage(ctx context.Context) (*Message, error)

	// Receive 从Redis服务器接收一个作为gvar.Var的单个回复。 md5:c4dad7138865cef4
	Receive(ctx context.Context) (result *gvar.Var, err error)
}
