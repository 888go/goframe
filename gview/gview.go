// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gview 实现了一个基于 text/template 的模板引擎。
//
// 预留的模板变量名称：
// I18nLanguage: 将此变量赋值以在每一页上定义 i18n 语言。
// 这段 Go 代码注释翻译成中文后为：
// ```go
// 包 gview 实现了一个基于 text/template 标准库的模板引擎功能。
//
// 已保留的模板变量名称：
// I18nLanguage：将该变量进行赋值，以便在每个页面上定义国际化的（i18n）语言设置。
package 模板类

import (
	"context"

	"github.com/888go/goframe/gview/internal/intlog"
	"github.com/gogf/gf/v2"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
)

// 模板引擎的视图对象。
type View struct {
	searchPaths  *garray.StrArray       // 为了性能考虑，以下代码在切片中搜索路径，但并不保证并发安全。
	data         map[string]interface{} // 全局模板变量。
	funcMap      map[string]interface{} // 全局模板函数映射。
	fileCacheMap *gmap.StrAnyMap        // File cache map.
	config       Config                 // 额外的视图配置
}

type (
	Params  = map[string]interface{} // Params 是模板参数的类型。
	FuncMap = map[string]interface{} // FuncMap 是用于自定义模板函数的类型。
)

const (
	commandEnvKeyForPath = "gf.gview.path"
)

var (
	// 默认视图对象。
	defaultViewObj *View
)

// checkAndInitDefaultView 检查并初始化默认视图对象。
// 默认视图对象仅会被初始化一次。
func checkAndInitDefaultView() {
	if defaultViewObj == nil {
		defaultViewObj = New()
	}
}

// ParseContent 使用默认视图对象直接解析模板内容，
// 并返回已解析的内容。
func ParseContent(ctx context.Context, content string, params ...Params) (string, error) {
	checkAndInitDefaultView()
	return defaultViewObj.ParseContent(ctx, content, params...)
}

// New返回一个新的视图对象。
// 参数`path`用于指定加载模板文件的模板目录路径。
func New(path ...string) *View {
	var (
		ctx = context.TODO()
	)
	view := &View{
		searchPaths:  garray.NewStrArray(),
		data:         make(map[string]interface{}),
		funcMap:      make(map[string]interface{}),
		fileCacheMap: gmap.NewStrAnyMap(true),
		config:       DefaultConfig(),
	}
	if len(path) > 0 && len(path[0]) > 0 {
		if err := view.SetPath(path[0]); err != nil {
			intlog.Errorf(context.TODO(), `%+v`, err)
		}
	} else {
		// 从环境变量/命令行自定义目录路径。
		if envPath := gcmd.GetOptWithEnv(commandEnvKeyForPath).String(); envPath != "" {
			if gfile.Exists(envPath) {
				if err := view.SetPath(envPath); err != nil {
					intlog.Errorf(context.TODO(), `%+v`, err)
				}
			} else {
				if errorPrint() {
					glog.Errorf(ctx, "Template directory path does not exist: %s", envPath)
				}
			}
		} else {
			// Dir：工作目录的路径。
			if err := view.SetPath(gfile.Pwd()); err != nil {
				intlog.Errorf(context.TODO(), `%+v`, err)
			}
			// Dir 二进制文件的路径。
			if selfPath := gfile.SelfDir(); selfPath != "" && gfile.Exists(selfPath) {
				if err := view.AddPath(selfPath); err != nil {
					intlog.Errorf(context.TODO(), `%+v`, err)
				}
			}
			// Dir：主包的路径。
			if mainPath := gfile.MainPkgPath(); mainPath != "" && gfile.Exists(mainPath) {
				if err := view.AddPath(mainPath); err != nil {
					intlog.Errorf(context.TODO(), `%+v`, err)
				}
			}
		}
	}
	view.SetDelimiters("{{", "}}")
	// 默认内置变量
	view.data["GF"] = map[string]interface{}{
		"version": gf.VERSION,
	}
	// 默认内置函数
	view.BindFuncMap(FuncMap{
		"eq":         view.buildInFuncEq,
		"ne":         view.buildInFuncNe,
		"lt":         view.buildInFuncLt,
		"le":         view.buildInFuncLe,
		"gt":         view.buildInFuncGt,
		"ge":         view.buildInFuncGe,
		"text":       view.buildInFuncText,
		"html":       view.buildInFuncHtmlEncode,
		"htmlencode": view.buildInFuncHtmlEncode,
		"htmldecode": view.buildInFuncHtmlDecode,
		"encode":     view.buildInFuncHtmlEncode,
		"decode":     view.buildInFuncHtmlDecode,
		"url":        view.buildInFuncUrlEncode,
		"urlencode":  view.buildInFuncUrlEncode,
		"urldecode":  view.buildInFuncUrlDecode,
		"date":       view.buildInFuncDate,
		"substr":     view.buildInFuncSubStr,
		"strlimit":   view.buildInFuncStrLimit,
		"concat":     view.buildInFuncConcat,
		"replace":    view.buildInFuncReplace,
		"compare":    view.buildInFuncCompare,
		"hidestr":    view.buildInFuncHideStr,
		"highlight":  view.buildInFuncHighlight,
		"toupper":    view.buildInFuncToUpper,
		"tolower":    view.buildInFuncToLower,
		"nl2br":      view.buildInFuncNl2Br,
		"include":    view.buildInFuncInclude,
		"dump":       view.buildInFuncDump,
		"map":        view.buildInFuncMap,
		"maps":       view.buildInFuncMaps,
		"json":       view.buildInFuncJson,
		"xml":        view.buildInFuncXml,
		"ini":        view.buildInFuncIni,
		"yaml":       view.buildInFuncYaml,
		"yamli":      view.buildInFuncYamlIndent,
		"toml":       view.buildInFuncToml,
		"plus":       view.buildInFuncPlus,
		"minus":      view.buildInFuncMinus,
		"times":      view.buildInFuncTimes,
		"divide":     view.buildInFuncDivide,
	})
	return view
}
