// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gcfg

import (
	"context"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
)

// AdapterContent 使用内容实现 Adapter 接口。
// 配置内容支持的编码类型与 `gjson` 包中的编码类型相同。
// md5:81e39ba9e6de51fa
type AdapterContent struct {
	jsonVar *gvar.Var // 配置内容的修剪过的 JSON 对象，类型：*gjson.Json。 md5:379162a7a5ad528f
}

// NewAdapterContent 返回一个使用自定义内容的新配置管理对象。
// 参数 `content` 指定用于读取的默认配置内容。
// md5:efafcabf61d7087b
// ff:
// content:
func NewAdapterContent(content ...string) (*AdapterContent, error) {
	a := &AdapterContent{
		jsonVar: gvar.New(nil, true),
	}
	if len(content) > 0 {
		if err := a.SetContent(content[0]); err != nil {
			return nil, err
		}
	}
	return a, nil
}

// SetContent 为指定的`file`设置自定义配置内容。
// `file`是可选参数，默认值为DefaultConfigFile。
// md5:49ae38cf671e3b96
// ff:
// a:
// content:
func (a *AdapterContent) SetContent(content string) error {
	j, err := gjson.LoadContent(content, true)
	if err != nil {
		return gerror.Wrap(err, `load configuration content failed`)
	}
	a.jsonVar.Set(j)
	return nil
}

// 可用性检查并返回后端配置服务是否可用。
// 可选参数 `resource` 指定特定的配置资源。
// 
// 请注意，此函数不会返回错误，因为它只是简单地检查后端配置服务。
// md5:79f955eb2fcdd137
// ff:
// a:
// ctx:
// resource:
// ok:
func (a *AdapterContent) Available(ctx context.Context, resource ...string) (ok bool) {
	if a.jsonVar.IsNil() {
		return false
	}
	return true
}

// Get 通过当前资源中指定的`pattern`获取并返回值。
// `pattern`示例：
// "x.y.z" 用于map中的条目。
// "x.0.y" 用于切片中的条目。
// md5:39b9171603468968
// ff:
// a:
// ctx:
// pattern:
// value:
// err:
func (a *AdapterContent) Get(ctx context.Context, pattern string) (value interface{}, err error) {
	if a.jsonVar.IsNil() {
		return nil, nil
	}
	return a.jsonVar.Val().(*gjson.Json).Get(pattern).Val(), nil
}

// Data 获取并以映射的形式返回当前资源中的所有配置数据。
// 注意，如果配置数据量过大，此函数可能会占用大量内存。
// 如有需要，你可以根据实际情况实现这个函数。
// md5:19dfa88d9aa6ece5
// ff:
// a:
// ctx:
// data:
// err:
func (a *AdapterContent) Data(ctx context.Context) (data map[string]interface{}, err error) {
	if a.jsonVar.IsNil() {
		return nil, nil
	}
	return a.jsonVar.Val().(*gjson.Json).Var().Map(), nil
}
