// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package goai

// Config 为 OpenApiV3 的实现提供了额外的配置功能。
type Config struct {
	ReadContentTypes        []string    // ReadContentTypes 指定了在未配置 MIME 类型时默认用于消费的 MIME 类型。
	WriteContentTypes       []string    // WriteContentTypes 指定在未配置 MIME 类型时生成内容时的默认 MIME 类型。
	CommonRequest           interface{} // 通用请求结构，适用于所有路径。
	CommonRequestDataField  string      // 该注释描述了一个通用请求字段名称，它将在特定业务请求结构中被替换。例如：`Data`, `Request`。 
// 翻译为：
// 这是一个通用请求字段名，通常会在特定的业务请求结构中用相应的内容进行替换。例如：`Data`、`Request`。
	CommonResponse          interface{} // 通用响应结构，适用于所有路径。
	CommonResponseDataField string      // 常用的响应字段名称，将会被替换为特定业务响应结构。例如：`Data`、`Response`。
	IgnorePkgPath           bool        // 忽略包名以获取模式名称。
}

// fillWithDefaultValue 如果未配置，则使用默认值填充 `oai` 的配置对象。
func (oai *OpenApiV3) fillWithDefaultValue() {
	if oai.OpenAPI == "" {
		oai.OpenAPI = `3.0.0`
	}
	if len(oai.Config.ReadContentTypes) == 0 {
		oai.Config.ReadContentTypes = defaultReadContentTypes
	}
	if len(oai.Config.WriteContentTypes) == 0 {
		oai.Config.WriteContentTypes = defaultWriteContentTypes
	}
}
