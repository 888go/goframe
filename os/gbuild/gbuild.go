// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gbuild 管理来自 "gf build" 的内置变量。
package gbuild

import (
	"context"
	"runtime"
	
	"github.com/888go/goframe"
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/encoding/gbase64"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/internal/json"
)

// BuildInfo 用于维护当前二进制文件的构建信息。
type BuildInfo struct {
	GoFrame string                 // 使用的GoFrame版本。
	Golang  string                 // 使用的Golang构建版本。
	Git     string                 // 此处注释翻译为：// 根据所用git仓库，构建提交ID及日期时间。
	Time    string                 // Built datetime.
	Version string                 // Built version.
	Data    map[string]interface{} // 所有自定义构建的数据键值对。
}

const (
	gfVersion    = `gfVersion`
	goVersion    = `goVersion`
	BuiltGit     = `builtGit`
	BuiltTime    = `builtTime`
	BuiltVersion = `builtVersion`
)

var (
	builtInVarStr = ""                       // Raw变量是一个Base64编码的字符串，该字符串由go build标志注入。
	builtInVarMap = map[string]interface{}{} // 二进制自定义变量映射已解码。
)

func init() {
	// `builtInVarStr`是由go构建标志注入的。
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

// Info 返回二进制文件的基本构建信息，以 map 的形式。
// 注意，此函数需配合 gf-cli 工具的 "gf build" 命令使用，
// 该命令会自动向二进制文件中注入必要的信息。
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

// Get 函数用于获取并返回指定名称的内置二进制变量。
func Get(name string, def ...interface{}) *gvar.Var {
	if v, ok := builtInVarMap[name]; ok {
		return gvar.New(v)
	}
	if len(def) > 0 {
		return gvar.New(def[0])
	}
	return nil
}

// Data 返回自定义内置变量作为映射（map）。
func Data() map[string]interface{} {
	return builtInVarMap
}
