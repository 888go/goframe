// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gins

import (
	"context"
	"fmt"
	
	"github.com/888go/goframe/database/gdb"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/consts"
	"github.com/888go/goframe/internal/instance"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/os/gcfg"
	"github.com/888go/goframe/os/glog"
	"github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/gutil"
)

// Database 返回指定配置组名称的数据库 ORM 对象实例。
// 注意，如果实例创建过程中发生任何错误，它会触发 panic。
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
		// 它忽略返回的错误，以防止在不必要的时候出现文件未找到的错误。
		var (
			configMap     map[string]interface{}
			configNodeKey = consts.ConfigNodeNameDatabase
		)
		// 首先，它会搜索实例名称的配置。
		if configData, _ := Config().Data(ctx); len(configData) > 0 {
			if v, _ := gutil.MapPossibleItemByKey(configData, consts.ConfigNodeNameDatabase); v != "" {
				configNodeKey = v
			}
		}
		if v, _ := Config().Get(ctx, configNodeKey); !v.IsEmpty() {
			configMap = v.Map()
		}
		// 未找到配置，它将格式化并引发错误。
		if len(configMap) == 0 && !gdb.IsConfigured() {
			// 文件配置对象检查。
			var err error
			if fileConfig, ok := Config().GetAdapter().(*gcfg.AdapterFile); ok {
				if _, err = fileConfig.GetFilePath(); err != nil {
					panic(gerror.WrapCode(gcode.CodeMissingConfiguration, err,
						`configuration not found, did you miss the configuration file or misspell the configuration file name`,
					))
				}
			}
			// 如果在Config对象或gdb配置中未找到任何内容，则引发panic。
			if len(configMap) == 0 && !gdb.IsConfigured() {
				panic(gerror.NewCodef(
					gcode.CodeMissingConfiguration,
					`database initialization failed: configuration missing for database node "%s"`,
					consts.ConfigNodeNameDatabase,
				))
			}
		}

		if len(configMap) == 0 {
			configMap = make(map[string]interface{})
		}
		// 将`m`解析为映射切片并将其添加到gdb包的全局配置中。
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
				if gdb.GetConfig(group) == nil {
					intlog.Printf(ctx, "add configuration for group: %s, %#v", g, cg)
					gdb.SetConfigGroup(g, cg)
				} else {
					intlog.Printf(ctx, "ignore configuration as it already exists for group: %s, %#v", g, cg)
					intlog.Printf(ctx, "%s, %#v", g, cg)
				}
			}
		}
// 将`m`解析为单节点配置，
// 这是默认的组配置。
		if node := parseDBConfigNode(configMap); node != nil {
			cg := gdb.ConfigGroup{}
			if node.Link != "" || node.Host != "" {
				cg = append(cg, *node)
			}
			if len(cg) > 0 {
				if gdb.GetConfig(group) == nil {
					intlog.Printf(ctx, "add configuration for group: %s, %#v", gdb.DefaultGroupName, cg)
					gdb.SetConfigGroup(gdb.DefaultGroupName, cg)
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

		// 使用给定的配置创建一个新的ORM对象。
		if db, err := gdb.NewByGroup(name...); err == nil {
			// 初始化ORM的logger（日志器）。
			var (
				loggerConfigMap map[string]interface{}
				loggerNodeName  = fmt.Sprintf("%s.%s", configNodeKey, consts.ConfigNodeNameLogger)
			)
			if v, _ := Config().Get(ctx, loggerNodeName); !v.IsEmpty() {
				loggerConfigMap = v.Map()
			}
			if len(loggerConfigMap) == 0 {
				if v, _ := Config().Get(ctx, configNodeKey); !v.IsEmpty() {
					loggerConfigMap = v.Map()
				}
			}
			if len(loggerConfigMap) > 0 {
				if logger, ok := db.GetLogger().(*glog.Logger); ok {
					if err = logger.SetConfigWithMap(loggerConfigMap); err != nil {
						panic(err)
					}
				}
			}
			return db
		} else {
			// 如果出现 panic，通常是因为在给定的组中没有找到其配置。
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
	// 查找可能的`Link`配置内容。
	if _, v := gutil.MapPossibleItemByKey(nodeMap, "Link"); v != nil {
		node.Link = gconv.String(v)
	}
	return node
}
