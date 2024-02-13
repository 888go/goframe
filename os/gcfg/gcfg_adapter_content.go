// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 配置类

import (
	"context"
	
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/encoding/gjson"
	"github.com/888go/goframe/errors/gerror"
)

// AdapterContent 实现了 Adapter 接口，采用内容方式实现。
// 配置内容支持 `gjson` 包所支持的各种编码类型。
type AdapterContent struct {
	jsonVar *泛型类.Var // 配置内容的解析后的JSON对象，类型为：*gjson.Json
}

// NewAdapterContent返回一个使用自定义内容的新的配置管理对象。
// 参数`content`指定了用于读取的默认配置内容。
func NewAdapterContent(content ...string) (*AdapterContent, error) {
	a := &AdapterContent{
		jsonVar: 泛型类.X创建(nil, true),
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
	j, err := json类.X加载并自动识别格式(content, true)
	if err != nil {
		return 错误类.X多层错误(err, `load configuration content failed`)
	}
	a.jsonVar.X设置值(j)
	return nil
}

// Available 检查并返回配置服务是否可用。
// 可选参数 `resource` 用于指定特定的配置资源。
//
// 注意，此函数不会返回错误，因为它只是简单地检查后端配置服务。
func (a *AdapterContent) Available(ctx context.Context, resource ...string) (ok bool) {
	if a.jsonVar.X是否为Nil() {
		return false
	}
	return true
}

// Get 通过指定的 `pattern` 从当前资源中获取并返回值。
// pattern 格式如下：
// "x.y.z" 用于获取 map 中的项。
// "x.0.y" 用于获取 slice 中的项。
func (a *AdapterContent) Get(ctx context.Context, pattern string) (value interface{}, err error) {
	if a.jsonVar.X是否为Nil() {
		return nil, nil
	}
	return a.jsonVar.X取值().(*json类.Json).X取值(pattern).X取值(), nil
}

// Data 函数从当前资源中获取并返回所有的配置数据，以 map 的形式。
// 注意：如果配置数据过大，该函数可能会导致大量内存使用。
// 如果有必要，你可以自行实现这个函数。
func (a *AdapterContent) Data(ctx context.Context) (data map[string]interface{}, err error) {
	if a.jsonVar.X是否为Nil() {
		return nil, nil
	}
	return a.jsonVar.X取值().(*json类.Json).X取泛型类().X取Map(), nil
}
