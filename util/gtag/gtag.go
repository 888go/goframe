// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

// 包gtag提供结构体标签内容存储的功能。
//
// 注意，此包中的函数不保证并发安全，这意味着你不能在运行时调用它们，而应该在启动过程中调用。 md5:d4777cc34c9b1efa
package gtag

const (
	Default           = "default"      // 结构体字段的默认值标签，用于从HTTP请求中接收参数。 md5:471b058e0363f634
	DefaultShort      = "d"            // Short name of Default.
	Param             = "param"        // 参数名称，用于将特定参数转换为指定的结构体字段。 md5:33028ec3fe66a79b
	ParamShort        = "p"            // Short name of Param.
	Valid             = "valid"        // 用于结构体字段的验证规则标签。 md5:c9767807e4fe1067
	ValidShort        = "v"            // Short name of Valid.
	NoValidation      = "nv"           // 对指定的结构体/字段不进行验证。 md5:93e44f590e6cd034
	ORM               = "orm"          // ORM标签用于ORM功能，根据不同的场景执行不同的功能。 md5:e901afff806207ee
	Arg               = "arg"          // Arg 标签用于结构体，通常用于命令行参数选项。 md5:819c478470e2eba6
	Brief             = "brief"        // 对结构体的简要标签，通常被视为摘要。 md5:8cc2e142b471c9b7
	Root              = "root"         // 结构体的根标签，通常用于嵌套命令的管理。 md5:626e49be7a599621
	Additional        = "additional"   // 结构体的附加标签，通常用于命令的额外描述。 md5:b76bf06e005ec042
	AdditionalShort   = "ad"           // Additional 的简短名称。 md5:2f155740ef9ae268
	Path              = `path`         // HTTP请求的路由路径。 md5:f9eb0f9346f6006b
	Method            = `method`       // 用于HTTP请求的路由方法。 md5:26ee97992f417756
	Domain            = `domain`       // 为HTTP请求路由域名。 md5:993e9bc0c52ceff9
	Mime              = `mime`         // HTTP请求/响应的MIME类型。 md5:03b1bf9c4c1ec6e7
	Consumes          = `consumes`     // HTTP请求的MIME类型。 md5:8cbc13892b01d8ed
	Summary           = `summary`      // 结构体的摘要，通常用于请求结构体中的OpenAPI。 md5:a1c99f363a83aea5
	SummaryShort      = `sm`           // Short name of Summary.
	SummaryShort2     = `sum`          // Short name of Summary.
	Description       = `description`  // 结构体的描述，通常用于请求结构体的OpenAPI中。 md5:90f620e59b2029fd
	DescriptionShort  = `dc`           // Description的简短名称。 md5:338cad0bd7b5b50b
	DescriptionShort2 = `des`          // Description的简短名称。 md5:338cad0bd7b5b50b
	Example           = `example`      // 用于结构体的示例，通常用于OpenAPI请求结构体中。 md5:d16a29f9cbc58ce1
	ExampleShort      = `eg`           // Short name of Example.
	Examples          = `examples`     // 用于结构体的示例，通常用于请求结构体中的OpenAPI。 md5:d49c8328492a0c33
	ExamplesShort     = `egs`          // 示例的简短名称。 md5:4eb934008e0e6ef7
	ExternalDocs      = `externalDocs` // 对结构体的外部文档说明，通常用于OpenAPI的请求结构体。 md5:92492f2c20afc6ea
	ExternalDocsShort = `ed`           // 外部文档的简短名称。 md5:5fb4e7b4dfa5434e
	GConv             = "gconv"        // GConv 定义了指定结构字段的转换目标名称。 md5:2be73c0f237cbc0f
	GConvShort        = "c"            // GConv 定义了指定结构字段的转换目标名称。 md5:2be73c0f237cbc0f
	Json              = "json"         // JSON标签受到stdlib的支持。 md5:2d11fa95e0d1d656
	Security          = "security"     //swagger.io/docs/specification/authentication/. md5:28d2c5914a05dfdb
	In                = "in"           //swagger.io/docs/specification/describing-parameters/. md5:13bb3747232af609
)

// StructTagPriority 定义了Map*/Struct*函数的默认优先级标签。
// 注意，`gconv/param` 标签由旧版本的包使用。强烈建议未来改用简短的标签 `c/p`。 md5:0d809ae278226ae3
var StructTagPriority = []string{
	GConv, Param, GConvShort, ParamShort, Json,
}
