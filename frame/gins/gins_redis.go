// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gins

import (
	"context"
	"fmt"

	gredis "github.com/888go/goframe/database/gredis"
	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/consts"
	"github.com/888go/goframe/internal/instance"
	"github.com/888go/goframe/internal/intlog"
	gconv "github.com/888go/goframe/util/gconv"
	gutil "github.com/888go/goframe/util/gutil"
)

// Redis 返回一个根据指定配置组名初始化的 Redis 客户端实例。
// 注意，如果在创建实例过程中发生任何错误，它将会直接 panic。
// md5:d7ef478cb4351b62
func Redis(name ...string) *gredis.Redis {
	var (
		err   error
		ctx   = context.Background()
		group = gredis.DefaultGroupName
	)
	if len(name) > 0 && name[0] != "" {
		group = name[0]
	}
	instanceKey := fmt.Sprintf("%s.%s", frameCoreComponentNameRedis, group)
	result := instance.GetOrSetFuncLock(instanceKey, func() interface{} {
				// 如果已经配置，它将返回redis实例。 md5:c43cd5e840daf9f4
		if _, ok := gredis.GetConfig(group); ok {
			return gredis.Instance(group)
		}
		if Config().Available(ctx) {
			var (
				configMap   map[string]interface{}
				redisConfig *gredis.Config
				redisClient *gredis.Redis
			)
			if configMap, err = Config().Data(ctx); err != nil {
				intlog.Errorf(ctx, `retrieve config data map failed: %+v`, err)
			}
			if _, v := gutil.MapPossibleItemByKey(configMap, consts.ConfigNodeNameRedis); v != nil {
				configMap = gconv.Map(v)
			}
			if len(configMap) > 0 {
				if v, ok := configMap[group]; ok {
					if redisConfig, err = gredis.ConfigFromMap(gconv.Map(v)); err != nil {
						panic(err)
					}
				} else {
					intlog.Printf(ctx, `missing configuration for redis group "%s"`, group)
				}
			} else {
				intlog.Print(ctx, `missing configuration for redis: "redis" node not found`)
			}
			if redisClient, err = gredis.New(redisConfig); err != nil {
				panic(err)
			}
			return redisClient
		}
		panic(gerror.NewCode(
			gcode.CodeMissingConfiguration,
			`no configuration found for creating redis client`,
		))
		return nil
	})
	if result != nil {
		return result.(*gredis.Redis)
	}
	return nil
}
