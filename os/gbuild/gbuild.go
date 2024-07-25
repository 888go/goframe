// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包gbuild管理"gf build"中的内置变量。 md5:c98a1c81088c9728
package gbuild

import (
	"context"
	"runtime"

	"github.com/gogf/gf/v2"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/internal/intlog"
	"github.com/gogf/gf/v2/internal/json"
)

// BuildInfo 保存了当前二进制文件的构建信息。 md5:fccdfbf6d632a3f0
type BuildInfo struct {
	GoFrame string                 // 使用的GoFrame版本。 md5:f48f34673b81dcdc
	Golang  string                 // 使用的Golang版本构建。 md5:3a88eeca4a698cae
	Git     string                 // 使用git仓库构建的，包含提交ID和日期时间。 md5:31599fedeb28a501
	Time    string                 // Built datetime.
	Version string                 // Built version.
	Data    map[string]interface{} // 所有自定义构建的数据键值对。 md5:fe64f79fd979385c
}

const (
	gfVersion    = `gfVersion`
	goVersion    = `goVersion`
	BuiltGit     = `builtGit`
	BuiltTime    = `builtTime`
	BuiltVersion = `builtVersion`
)

var (
	builtInVarStr = ""                       // Raw 变量是一个由 go build 标志注入的 base64 字符串。 md5:a031cbbc33deb61d
	builtInVarMap = map[string]interface{}{} // 解码自定义变量映射到二进制。 md5:db1598eabdc63d93
)

func init() {
		// `builtInVarStr` 是由 Go 构建标志注入的。 md5:3fb4d6289f310035
	if builtInVarStr != "" {
		err := json.UnmarshalUseNumber(gbase64.MustDecodeString(builtInVarStr), &builtInVarMap)
		if err != nil {
			intlog.Errorf(context.TODO(), `%+v`, err)
		}
		builtInVarMap[gfVersion] = gf.VERSION
		builtInVarMap[goVersion] = runtime.Version()
		intlog.Printf(context.TODO(), "build variables: %+v", builtInVarMap)
	} else {
		intlog.Print(context.TODO(), "no build variables")
	}
}

// Info 返回二进制文件的基本构建信息，以映射的形式呈现。
// 注意，它应当与 gf-cli 工具的 "gf build" 命令配合使用，
// 该命令会自动将必需的信息注入到二进制文件中。
// md5:d327a3b92b2b2006
func Info() BuildInfo {
	return BuildInfo{
		GoFrame: Get(gfVersion).String(),
		Golang:  Get(goVersion).String(),
		Git:     Get(BuiltGit).String(),
		Time:    Get(BuiltTime).String(),
		Version: Get(BuiltVersion).String(),
		Data:    Data(),
	}
}

// Get 获取并返回具有给定名称的内置二进制变量。 md5:3b52dd5dc655096c
func Get(name string, def ...interface{}) *gvar.Var {
	if v, ok := builtInVarMap[name]; ok {
		return gvar.New(v)
	}
	if len(def) > 0 {
		return gvar.New(def[0])
	}
	return nil
}

// Data返回自定义的内置变量作为映射。 md5:d1a22399bad5a8b6
func Data() map[string]interface{} {
	return builtInVarMap
}
