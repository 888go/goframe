// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// Package gview implements a template engine based on text/template.
//
// Reserved template variable names:
package gview//bm:模板类

import (
	"context"

	"github.com/gogf/gf/v2"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/internal/intlog"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
)

// 用于模板引擎的视图对象。 md5:d5b31747d89262fc
type View struct {
	searchPaths  *garray.StrArray       // 在数组中搜索路径，为了性能原因，非并发安全。 md5:536357ec68a07213
	data         map[string]interface{} // 全局模板变量。 md5:a2dd7af9a0b6ae90
	funcMap      map[string]interface{} // 全局模板函数映射。 md5:b4a6607e9bbfd481
	fileCacheMap *gmap.StrAnyMap        // File cache map.
	config       Config                 // 为视图提供的额外配置。 md5:84cd8627170ef89d
}

type (
	Params  = map[string]interface{} // Params 是模板参数的类型。 md5:c3bc270bc0522288
	FuncMap = map[string]interface{} // FuncMap是自定义模板函数的类型。 md5:04d2535f72d33955
)

const (
	commandEnvKeyForPath = "gf.gview.path"
)

var (
	// Default view object.
	defaultViewObj *View
)

// checkAndInitDefaultView 检查并初始化默认视图对象。
// 默认视图对象仅会被初始化一次。
// md5:d74d29ccd894a2fa
func checkAndInitDefaultView() {
	if defaultViewObj == nil {
		defaultViewObj = New()
	}
}

// ParseContent 使用默认的视图对象直接解析模板内容，并返回解析后的内容。
// md5:8349c5832e5a90c1
// ff:
// ctx:
// content:
// params:
func ParseContent(ctx context.Context, content string, params ...Params) (string, error) {
	checkAndInitDefaultView()
	return defaultViewObj.ParseContent(ctx, content, params...)
}

// New 返回一个新的视图对象。
// 参数 `path` 指定加载模板文件的模板目录路径。
// md5:b96716da886c0dc3
// ff:
// path:
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
		// 从环境变量或命令行自定义的目录路径。 md5:8cfcbca968e23c5b
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
			// Dir 是工作目录的路径。 md5:0fba211853ea97a0
			if pwdPath := gfile.Pwd(); pwdPath != "" {
				if err := view.SetPath(pwdPath); err != nil {
					intlog.Errorf(context.TODO(), `%+v`, err)
				}
			}
			// Dir path of binary.
			if selfPath := gfile.SelfDir(); selfPath != "" && gfile.Exists(selfPath) {
				if err := view.AddPath(selfPath); err != nil {
					intlog.Errorf(context.TODO(), `%+v`, err)
				}
			}
			// 主包的目录路径。 md5:a4d2802779172abe
			if mainPath := gfile.MainPkgPath(); mainPath != "" && gfile.Exists(mainPath) {
				if err := view.AddPath(mainPath); err != nil {
					intlog.Errorf(context.TODO(), `%+v`, err)
				}
			}
		}
	}
	view.SetDelimiters("{{", "}}")
	// 默认内置变量。 md5:b0f8a83fbf9378e6
	view.data["GF"] = map[string]interface{}{
		"version": gf.VERSION,
	}
	// 默认内置函数。 md5:8ca9492d3b848286
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
