// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package redis类

import (
	"context"
	"fmt"
)

// IGroupPubSub 管理 Redis 发布/订阅操作。
// 实现请参见 redis.GroupPubSub。
type IGroupPubSub interface {
	Publish(ctx context.Context, channel string, message interface{}) (int64, error)
	Subscribe(ctx context.Context, channel string, channels ...string) (Conn, []*Subscription, error)
	PSubscribe(ctx context.Context, pattern string, patterns ...string) (Conn, []*Subscription, error)
}

// 作为另一个客户端发出的 PUBLISH 命令的结果接收到的消息。
type Message struct {
	Channel      string
	Pattern      string
	Payload      string
	PayloadSlice []string
}

// 订阅成功后接收到的订阅信息（针对通道）
type Subscription struct {
	Kind    string // 可以是 "subscribe", "unsubscribe", "psubscribe" 或 "punsubscribe"。
	Channel string // 我们已订阅的通道名称。
	Count   int    // 当前我们已订阅的频道数量。
}

// String 将当前对象转换为可读的字符串。
func (m *Subscription) String() string {
	return fmt.Sprintf("%s: %s", m.Kind, m.Channel)
}
