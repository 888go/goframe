// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gredis

import (
	"context"
	"fmt"
)

// IGroupPubSub 管理 redis 的发布/订阅操作。
// 实现了 redis.GroupPubSub 接口。 md5:9987807bf281dfb2
type IGroupPubSub interface {
	Publish(ctx context.Context, channel string, message interface{}) (int64, error)
	Subscribe(ctx context.Context, channel string, channels ...string) (Conn, []*Subscription, error)
	PSubscribe(ctx context.Context, pattern string, patterns ...string) (Conn, []*Subscription, error)
}

// 作为另一客户端发出PUBLISH命令的结果接收到的消息。 md5:1b54d1f2ea66a492
type Message struct {
	Channel      string
	Pattern      string
	Payload      string
	PayloadSlice []string
}

// 成功订阅频道后接收到的订阅信息。 md5:18e121df658bbde2
type Subscription struct {
	Kind    string // 可以是 "subscribe"、"unsubscribe"、"psubscribe" 或 "punsubscribe"。 md5:1344d1b60899d10a
	Channel string // 我们已订阅的频道名称。 md5:aed766d24e08c59a
	Count   int    // 我们当前订阅的频道数量。 md5:3af00e20830c4421
}

// String 将当前对象转换为可读的字符串。 md5:b831e7b38fabfe56
func (m *Subscription) String() string {
	return fmt.Sprintf("%s: %s", m.Kind, m.Channel)
}
