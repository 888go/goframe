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
func Database(name ...string) db类.DB {
	var (
		ctx   = context.Background()
		group = db类.DefaultGroupName
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
		if configData, _ := Config().X取Map(ctx); len(configData) > 0 {
			if v, _ := 工具类.MapPossibleItemByKey(configData, consts.ConfigNodeNameDatabase); v != "" {
				configNodeKey = v
			}
		}
		if v, _ := Config().X取值(ctx, configNodeKey); !v.X是否为空() {
			configMap = v.X取Map()
		}
		// 未找到配置，它将格式化并引发错误。
		if len(configMap) == 0 && !db类.X是否已配置数据库() {
			// 文件配置对象检查。
			var err error
			if fileConfig, ok := Config().X取适配器().(*配置类.AdapterFile); ok {
				if _, err = fileConfig.GetFilePath(); err != nil {
					panic(错误类.X多层错误码(错误码类.CodeMissingConfiguration, err,
						`configuration not found, did you miss the configuration file or misspell the configuration file name`,
					))
				}
			}
			// 如果在Config对象或gdb配置中未找到任何内容，则引发panic。
			if len(configMap) == 0 && !db类.X是否已配置数据库() {
				panic(错误类.X创建错误码并格式化(
					错误码类.CodeMissingConfiguration,
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
			cg := db类.X配置组{}
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
				if db类.X取配置组配置(group) == nil {
					intlog.Printf(ctx, "add configuration for group: %s, %#v", g, cg)
					db类.X设置组配置(g, cg)
				} else {
					intlog.Printf(ctx, "ignore configuration as it already exists for group: %s, %#v", g, cg)
					intlog.Printf(ctx, "%s, %#v", g, cg)
				}
			}
		}
// 将`m`解析为单节点配置，
// 这是默认的组配置。
		if node := parseDBConfigNode(configMap); node != nil {
			cg := db类.X配置组{}
			if node.X自定义链接信息 != "" || node.X地址 != "" {
				cg = append(cg, *node)
			}
			if len(cg) > 0 {
				if db类.X取配置组配置(group) == nil {
					intlog.Printf(ctx, "add configuration for group: %s, %#v", db类.DefaultGroupName, cg)
					db类.X设置组配置(db类.DefaultGroupName, cg)
				} else {
					intlog.Printf(
						ctx,
						"ignore configuration as it already exists for group: %s, %#v",
						db类.DefaultGroupName, cg,
					)
					intlog.Printf(ctx, "%s, %#v", db类.DefaultGroupName, cg)
				}
			}
		}

		// 使用给定的配置创建一个新的ORM对象。
		if db, err := db类.X创建DB对象并按配置组(name...); err == nil {
			// 初始化ORM的logger（日志器）。
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
				if logger, ok := db.X取日志记录器().(*日志类.Logger); ok {
					if err = logger.X设置配置Map(loggerConfigMap); err != nil {
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
		return db.(db类.DB)
	}
	return nil
}

func parseDBConfigNode(value interface{}) *db类.X配置项 {
	nodeMap, ok := value.(map[string]interface{})
	if !ok {
		return nil
	}
	var (
		node = &db类.X配置项{}
		err  = 转换类.Struct(nodeMap, node)
	)
	if err != nil {
		panic(err)
	}
	// 查找可能的`Link`配置内容。
	if _, v := 工具类.MapPossibleItemByKey(nodeMap, "Link"); v != nil {
		node.X自定义链接信息 = 转换类.String(v)
	}
	return node
}
