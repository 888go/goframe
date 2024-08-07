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

	gdb "github.com/888go/goframe/database/gdb"
	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/consts"
	"github.com/888go/goframe/internal/instance"
	"github.com/888go/goframe/internal/intlog"
	gcfg "github.com/888go/goframe/os/gcfg"
	glog "github.com/888go/goframe/os/glog"
	gconv "github.com/888go/goframe/util/gconv"
	gutil "github.com/888go/goframe/util/gutil"
)

// Database 返回一个根据指定配置组名实例化的数据库ORM对象。
// 注意，如果在实例创建过程中发生任何错误，它将会直接 panic。
// md5:c8c0e8142b2f24af
func Database(name ...string) gdb.DB {
	var (
		ctx   = context.Background()
		group = gdb.DefaultGroupName
	)

	if len(name) > 0 && name[0] != "" {
		group = name[0]
	}
	instanceKey := fmt.Sprintf("%s.%s", frameCoreComponentNameDatabase, group)
	db := instance.GetOrSetFuncLock(instanceKey, func() interface{} {
				// 忽略返回的错误，以避免在不需要时出现“文件未找到”错误。 md5:47e693921809cd8c
		var (
			configMap     map[string]interface{}
			configNodeKey = consts.ConfigNodeNameDatabase
		)
				// 它首先搜索实例名称的配置。 md5:0b825658b318a2f7
		if configData, _ := Config().X取Map(ctx); len(configData) > 0 {
			if v, _ := gutil.MapPossibleItemByKey(configData, consts.ConfigNodeNameDatabase); v != "" {
				configNodeKey = v
			}
		}
		if v, _ := Config().X取值(ctx, configNodeKey); !v.X是否为空() {
			configMap = v.X取Map()
		}
				// 没有找到配置，它会格式化并引发 panic 错误。 md5:8716646cceaee999
		if len(configMap) == 0 && !gdb.X是否已配置数据库() {
						// 文件配置对象检查。 md5:fdae1c62b2593d55
			var err error
			if fileConfig, ok := Config().X取适配器().(*gcfg.AdapterFile); ok {
				if _, err = fileConfig.GetFilePath(); err != nil {
					panic(gerror.X多层错误码(gcode.CodeMissingConfiguration, err,
						`configuration not found, did you miss the configuration file or misspell the configuration file name`,
					))
				}
			}
						// 如果在Config对象或gdb配置中找不到任何内容，则引发恐慌。 md5:2c3aa642bbae15da
			if len(configMap) == 0 && !gdb.X是否已配置数据库() {
				panic(gerror.X创建错误码并格式化(
					gcode.CodeMissingConfiguration,
					`database initialization failed: configuration missing for database node "%s"`,
					consts.ConfigNodeNameDatabase,
				))
			}
		}

		if len(configMap) == 0 {
			configMap = make(map[string]interface{})
		}
				// 将 `m` 解析为映射切片，并将其添加到gdb包的全局配置中。 md5:8970d506724c2880
		for g, groupConfig := range configMap {
			cg := gdb.ConfigGroup{}
			switch value := groupConfig.(type) {
			case []interface{}:
				for _, v := range value {
					if node := parseDBConfigNode(v); node != nil {
						cg = append(cg, *node)
					}
				}
			case map[string]interface{}:
				if node := parseDBConfigNode(value); node != nil {
					cg = append(cg, *node)
				}
			}
			if len(cg) > 0 {
				if gdb.X取配置组配置(group) == nil {
					intlog.Printf(ctx, "add configuration for group: %s, %#v", g, cg)
					gdb.X设置组配置(g, cg)
				} else {
					intlog.Printf(ctx, "ignore configuration as it already exists for group: %s, %#v", g, cg)
					intlog.Printf(ctx, "%s, %#v", g, cg)
				}
			}
		}
		// 将 `m` 解析为单个节点配置，
		// 这是默认的组配置。
		// md5:8f62d1ad0b43783e
		if node := parseDBConfigNode(configMap); node != nil {
			cg := gdb.ConfigGroup{}
			if node.Link != "" || node.Host != "" {
				cg = append(cg, *node)
			}
			if len(cg) > 0 {
				if gdb.X取配置组配置(group) == nil {
					intlog.Printf(ctx, "add configuration for group: %s, %#v", gdb.DefaultGroupName, cg)
					gdb.X设置组配置(gdb.DefaultGroupName, cg)
				} else {
					intlog.Printf(
						ctx,
						"ignore configuration as it already exists for group: %s, %#v",
						gdb.DefaultGroupName, cg,
					)
					intlog.Printf(ctx, "%s, %#v", gdb.DefaultGroupName, cg)
				}
			}
		}

				// 使用给定的配置创建一个新的ORM对象。 md5:8114aaedeed4c350
		if db, err := gdb.X创建DB对象并按配置组(name...); err == nil {
						// 初始化ORM的日志记录器。 md5:5fbf0eb7ce9402d0
			var (
				loggerConfigMap map[string]interface{}
				loggerNodeName  = fmt.Sprintf("%s.%s", configNodeKey, consts.ConfigNodeNameLogger)
			)
			if v, _ := Config().X取值(ctx, loggerNodeName); !v.X是否为空() {
				loggerConfigMap = v.X取Map()
			}
			if len(loggerConfigMap) == 0 {
				if v, _ := Config().X取值(ctx, configNodeKey); !v.X是否为空() {
					loggerConfigMap = v.X取Map()
				}
			}
			if len(loggerConfigMap) > 0 {
				if logger, ok := db.X取日志记录器().(*glog.Logger); ok {
					if err = logger.X设置配置Map(loggerConfigMap); err != nil {
						panic(err)
					}
				}
			}
			return db
		} else {
						// 如果出现恐慌，通常是由于它没有找到给定组的配置。 md5:461786d647ecc99d
			panic(err)
		}
		return nil
	})
	if db != nil {
		return db.(gdb.DB)
	}
	return nil
}

func parseDBConfigNode(value interface{}) *gdb.ConfigNode {
	nodeMap, ok := value.(map[string]interface{})
	if !ok {
		return nil
	}
	var (
		node = &gdb.ConfigNode{}
		err  = gconv.Struct(nodeMap, node)
	)
	if err != nil {
		panic(err)
	}
		// 查找可能的 `Link` 配置内容。 md5:c3acedff678206f1
	if _, v := gutil.MapPossibleItemByKey(nodeMap, "Link"); v != nil {
		node.Link = gconv.String(v)
	}
	return node
}
