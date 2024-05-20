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
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"
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

// initGroup 初始化Redis的组对象。. md5:f3c3ba5dbd6196a2
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

// SetAdapter 将当前 Redis 客户端的底层适配器替换为自定义适配器。. md5:f503f97750dcd95a
func (r *Redis) SetAdapter(adapter Adapter) {
	if r == nil {
		panic(gerror.NewCode(gcode.CodeInvalidParameter, errorNilRedis))
	}
	r.localAdapter = adapter
}

// GetAdapter 返回当前 Redis 客户端设置的适配器。. md5:c46228b935b43204
func (r *Redis) GetAdapter() Adapter {
	if r == nil {
		return nil
	}
	return r.localAdapter
}

// Conn 获取并返回一个用于连续操作的连接对象。
// 请注意，如果您不再使用此连接，请手动调用 Close 函数。
// md5:b0379f4ab8131447
func (r *Redis) Conn(ctx context.Context) (Conn, error) {
	if r == nil {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, errorNilRedis)
	}
	if r.localAdapter == nil {
		return nil, gerror.NewCode(gcode.CodeNecessaryPackageNotImport, errorNilAdapter)
	}
	return r.localAdapter.Conn(ctx)
}

// Do 向服务器发送命令并返回接收到的回复。
// 它在将结构体、切片或映射类型值提交到Redis之前，使用json.Marshal进行序列化。
// md5:bbe59d4e1ff07fa3
func (r *Redis) Do(ctx context.Context, command string, args ...interface{}) (*gvar.Var, error) {
	if r == nil {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, errorNilRedis)
	}
	if r.localAdapter == nil {
		return nil, gerror.NewCodef(gcode.CodeMissingConfiguration, errorNilAdapter)
	}
	return r.localAdapter.Do(ctx, command, args...)
}

// MustConn 表现如同 Conn 函数，但是如果内部发生任何错误，它将引发 panic。. md5:555eb0f8f348b94c
func (r *Redis) MustConn(ctx context.Context) Conn {
	c, err := r.Conn(ctx)
	if err != nil {
		panic(err)
	}
	return c
}

// MustDo 执行与 Do 相同的操作，但如果内部出现任何错误，它将引发 panic。. md5:0d30101f0e9e6a4e
func (r *Redis) MustDo(ctx context.Context, command string, args ...interface{}) *gvar.Var {
	v, err := r.Do(ctx, command, args...)
	if err != nil {
		panic(err)
	}
	return v
}

// Close 方法关闭当前Redis客户端，关闭其连接池并释放所有相关资源。. md5:bfd91d0269572038
func (r *Redis) Close(ctx context.Context) error {
	if r == nil || r.localAdapter == nil {
		return nil
	}
	return r.localAdapter.Close(ctx)
}
