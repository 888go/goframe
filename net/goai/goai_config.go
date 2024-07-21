// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package goai

// Config 为 OpenApiV3 实现提供了额外的配置功能。
// 备注: 此配置结构不做名称翻译, 防止通过map载入配置时, 会直接将文本名称转换成配置项名称, 导致找不到原名的配置项. (2024-07-21)
// md5:094e64f2ea1fb1e8
type Config struct {
	ReadContentTypes        []string    // ReadContentTypes 指定了在未配置 MIME 类型时的默认消费 MIME 类型。 md5:a693e149658e6922
	WriteContentTypes       []string    // WriteContentTypes 指定如果没有配置 MIME 类型，则用于生成的默认 MIME 类型。 md5:08e1c2c9661386cb
	CommonRequest           interface{} // 所有路径的常见请求结构。 md5:abc50dd97166429e
	CommonRequestDataField  string      // 通用请求字段名，将被替换为特定业务请求结构。例如：`Data`、`Request`。 md5:91f6acb0ea17d306
	CommonResponse          interface{} // 所有路径的通用响应结构。 md5:1b61004412fdcd06
	CommonResponseDataField string      // 要替换为特定业务响应结构的常见响应字段名称，例如：`Data`、`Response.`。 md5:bb674debcf674e63
	IgnorePkgPath           bool        // 忽略包名，使用schema名称。 md5:d3c2c4cc40dd7ee2
}

// fillWithDefaultValue 如果未配置，默认情况下使用默认值填充 `oai` 的配置对象。 md5:d191b7feabf4c49d
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
