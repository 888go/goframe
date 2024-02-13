// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gins

import (
	"context"
	"fmt"
	
	"github.com/888go/goframe/database/gredis"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/consts"
	"github.com/888go/goframe/internal/instance"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/gutil"
)

// Redis 返回一个使用指定配置组名称的 Redis 客户端实例。
// 需要注意的是，如果在创建实例期间发生任何错误，它会引发 panic。
func Redis(name ...string) *redis类.Redis {
	var (
		err   error
		ctx   = context.Background()
		group = redis类.DefaultGroupName
	)
	if len(name) > 0 && name[0] != "" {
		group = name[0]
	}
	instanceKey := fmt.Sprintf("%s.%s", frameCoreComponentNameRedis, group)
	result := instance.GetOrSetFuncLock(instanceKey, func() interface{} {
		// 如果已经配置过，则返回redis实例。
		if _, ok := redis类.GetConfig(group); ok {
			return redis类.Instance(group)
		}
		if Config().X是否可用(ctx) {
			var (
				configMap   map[string]interface{}
				redisConfig *redis类.Config
				redisClient *redis类.Redis
			)
			if configMap, err = Config().X取Map(ctx); err != nil {
				intlog.Errorf(ctx, `retrieve config data map failed: %+v`, err)
			}
			if _, v := 工具类.MapPossibleItemByKey(configMap, consts.ConfigNodeNameRedis); v != nil {
				configMap = 转换类.X取Map(v)
			}
			if len(configMap) > 0 {
				if v, ok := configMap[group]; ok {
					if redisConfig, err = redis类.ConfigFromMap(转换类.X取Map(v)); err != nil {
						panic(err)
					}
				} else {
					intlog.Printf(ctx, `missing configuration for redis group "%s"`, group)
				}
			} else {
				intlog.Print(ctx, `missing configuration for redis: "redis" node not found`)
			}
			if redisClient, err = redis类.New(redisConfig); err != nil {
				panic(err)
			}
			return redisClient
		}
		panic(错误类.X创建错误码(
			错误码类.CodeMissingConfiguration,
			`no configuration found for creating redis client`,
		))
		return nil
	})
	if result != nil {
		return result.(*redis类.Redis)
	}
	return nil
}
