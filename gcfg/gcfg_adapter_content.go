// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gcfg

import (
	"context"
	
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
)

// AdapterContent 实现了 Adapter 接口，采用内容方式实现。
// 配置内容支持 `gjson` 包所支持的各种编码类型。
type AdapterContent struct {
	jsonVar *gvar.Var // 配置内容的解析后的JSON对象，类型为：*gjson.Json
}

// NewAdapterContent返回一个使用自定义内容的新的配置管理对象。
// 参数`content`指定了用于读取的默认配置内容。
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
// `file`参数不是必须的，默认值是DefaultConfigFile。
func (a *AdapterContent) SetContent(content string) error {
	j, err := gjson.LoadContent(content, true)
	if err != nil {
		return gerror.Wrap(err, `load configuration content failed`)
	}
	a.jsonVar.Set(j)
	return nil
}

// Available 检查并返回配置服务是否可用。
// 可选参数 `resource` 用于指定特定的配置资源。
//
// 注意，此函数不会返回错误，因为它只是简单地检查后端配置服务。
func (a *AdapterContent) Available(ctx context.Context, resource ...string) (ok bool) {
	if a.jsonVar.IsNil() {
		return false
	}
	return true
}

// Get 通过指定的 `pattern` 从当前资源中获取并返回值。
// pattern 格式如下：
// "x.y.z" 用于获取 map 中的项。
// "x.0.y" 用于获取 slice 中的项。
func (a *AdapterContent) Get(ctx context.Context, pattern string) (value interface{}, err error) {
	if a.jsonVar.IsNil() {
		return nil, nil
	}
	return a.jsonVar.Val().(*gjson.Json).Get(pattern).Val(), nil
}

// Data 函数从当前资源中获取并返回所有的配置数据，以 map 的形式。
// 注意：如果配置数据过大，该函数可能会导致大量内存使用。
// 如果有必要，你可以自行实现这个函数。
func (a *AdapterContent) Data(ctx context.Context) (data map[string]interface{}, err error) {
	if a.jsonVar.IsNil() {
		return nil, nil
	}
	return a.jsonVar.Val().(*gjson.Json).Var().Map(), nil
}
