// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gredis
import (
	"context"
	
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/text/gstr"
	)
// Redis client.
type Redis struct {
	config *Config
	localAdapter
	localGroup
}

type (
	localGroup struct {
		localGroupGeneric
		localGroupHash
		localGroupList
		localGroupPubSub
		localGroupScript
		localGroupSet
		localGroupSortedSet
		localGroupString
	}
	localAdapter        = Adapter
	localGroupGeneric   = IGroupGeneric
	localGroupHash      = IGroupHash
	localGroupList      = IGroupList
	localGroupPubSub    = IGroupPubSub
	localGroupScript    = IGroupScript
	localGroupSet       = IGroupSet
	localGroupSortedSet = IGroupSortedSet
	localGroupString    = IGroupString
)

const (
	errorNilRedis = `the Redis object is nil`
)

var (
	errorNilAdapter = gstr.Trim(gstr.Replace(`
redis adapter is not set, missing configuration or adapter register? 
possible reference: https://github.com/gogf/gf/tree/master/contrib/nosql/redis
`, "\n", ""))
)

// initGroup 初始化redis的组对象
func (r *Redis) initGroup() *Redis {
	r.localGroup = localGroup{
		localGroupGeneric:   r.localAdapter.GroupGeneric(),
		localGroupHash:      r.localAdapter.GroupHash(),
		localGroupList:      r.localAdapter.GroupList(),
		localGroupPubSub:    r.localAdapter.GroupPubSub(),
		localGroupScript:    r.localAdapter.GroupScript(),
		localGroupSet:       r.localAdapter.GroupSet(),
		localGroupSortedSet: r.localAdapter.GroupSortedSet(),
		localGroupString:    r.localAdapter.GroupString(),
	}
	return r
}

// SetAdapter 更改当前Redis客户端的底层适配器，使用自定义适配器。
func (r *Redis) SetAdapter(adapter Adapter) {
	if r == nil {
		panic(gerror.NewCode(gcode.CodeInvalidParameter, errorNilRedis))
	}
	r.localAdapter = adapter
}

// GetAdapter 返回当前Redis客户端中设置的适配器。
func (r *Redis) GetAdapter() Adapter {
	if r == nil {
		return nil
	}
	return r.localAdapter
}

// Conn 获取并返回一个用于连续操作的连接对象。
// 注意，如果你不再使用此连接，应手动调用 Close 函数。
func (r *Redis) Conn(ctx context.Context) (Conn, error) {
	if r == nil {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, errorNilRedis)
	}
	if r.localAdapter == nil {
		return nil, gerror.NewCode(gcode.CodeNecessaryPackageNotImport, errorNilAdapter)
	}
	return r.localAdapter.Conn(ctx)
}

// 向服务器发送命令并返回接收到的回复。
// 在将结构体/切片/映射类型值提交给redis之前，它使用json.Marshal进行序列化。
func (r *Redis) Do(ctx context.Context, command string, args ...interface{}) (*gvar.Var, error) {
	if r == nil {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, errorNilRedis)
	}
	if r.localAdapter == nil {
		return nil, gerror.NewCodef(gcode.CodeMissingConfiguration, errorNilAdapter)
	}
	return r.localAdapter.Do(ctx, command, args...)
}

// MustConn 的行为与函数 Conn 相同，但如果内部发生任何错误，它会触发 panic。
func (r *Redis) MustConn(ctx context.Context) Conn {
	c, err := r.Conn(ctx)
	if err != nil {
		panic(err)
	}
	return c
}

// MustDo 执行与函数 Do 相同的操作，但如果内部出现任何错误，它会触发 panic（异常）。
func (r *Redis) MustDo(ctx context.Context, command string, args ...interface{}) *gvar.Var {
	v, err := r.Do(ctx, command, args...)
	if err != nil {
		panic(err)
	}
	return v
}

// Close 关闭当前的 Redis 客户端，关闭其连接池并释放所有相关资源。
func (r *Redis) Close(ctx context.Context) error {
	if r == nil || r.localAdapter == nil {
		return nil
	}
	return r.localAdapter.Close(ctx)
}
