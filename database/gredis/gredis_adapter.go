// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gredis

import (
	"context"
	
	"github.com/888go/goframe/container/gvar"
)

// Adapter 是一个用于通用 Redis 操作的接口。
type Adapter interface {
	AdapterGroup

// 向服务器发送命令并返回接收到的回复。
// 在将结构体、切片或映射类型值提交到redis前，它使用json.Marshal进行序列化。
	Do(ctx context.Context, command string, args ...interface{}) (*gvar.Var, error)

// Conn 获取并返回一个用于连续操作的连接对象。
// 注意，如果你不再使用此连接，应手动调用 Close 函数。
	Conn(ctx context.Context) (conn Conn, err error)

	// Close 关闭当前的 Redis 客户端，关闭其连接池并释放所有相关的资源。
	Close(ctx context.Context) (err error)
}

// Conn 是一个通用 Redis 客户端连接的接口。
type Conn interface {
	ConnCommand

// 向服务器发送命令并返回接收到的回复。
// 在将结构体、切片或映射类型值提交到redis前，它使用json.Marshal进行序列化。
	Do(ctx context.Context, command string, args ...interface{}) (result *gvar.Var, err error)

	// Close将连接放回连接池。
	Close(ctx context.Context) (err error)
}

// AdapterGroup 是一个接口，用于管理针对 Redis 的组操作。
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

// ConnCommand 是一个接口，用于管理与特定连接相关的一些操作。
type ConnCommand interface {
// Subscribe 订阅函数，使客户端订阅指定的频道。
// 参考文档：https://redis.io/commands/subscribe/
	Subscribe(ctx context.Context, channel string, channels ...string) ([]*Subscription, error)

// PSubscribe 订阅客户端到给定的模式。
//
// 支持的glob风格模式：
// - h?llo 订阅hello, hallo和hxllo
// - h*llo 订阅hllo和heeeello
// - h[ae]llo 订阅hello和hallo，但不订阅hillo
//
// 如果你想精确匹配特殊字符，请使用\进行转义。
//
// 参考文档：https://redis.io/commands/psubscribe/
	PSubscribe(ctx context.Context, pattern string, patterns ...string) ([]*Subscription, error)

	// ReceiveMessage 从 Redis 服务器接收订阅的单条消息。
	ReceiveMessage(ctx context.Context) (*Message, error)

	// Receive 从Redis服务器接收单个回复作为gvar.Var。
	Receive(ctx context.Context) (result *gvar.Var, err error)
}
