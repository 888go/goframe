// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gtag 提供了为结构体存储标签内容的功能。
//
// 注意：该包中的函数调用不保证并发安全，
// 这意味着你不能在运行时调用它们，而只能在启动过程中调用。
// 这段注释翻译成中文后为：
// ```go
// 包gtag 提供了用于为结构体提供标签内容存储的服务。
//
// 需要注意的是，本包中提供的函数并非线程安全的，
// 即这些函数不能在程序运行时并发调用，只能在初始化或启动阶段调用。
package gtag

const (
	Default           = "default"      // 结构体字段的默认值标签，用于从HTTP请求接收参数。
	DefaultShort      = "d"            // Default的简称。
	Param             = "param"        // 将特定参数转换为指定结构体字段的参数名称。
	ParamShort        = "p"            // Param的简称。
	Valid             = "valid"        // 结构体字段的验证规则标签
	ValidShort        = "v"            // Valid的简称
	NoValidation      = "nv"           // 对指定的结构体/字段不进行验证。
	ORM               = "orm"          // ORM标签用于ORM特性，根据不同的场景执行不同的功能。
	Arg               = "arg"          // Arg标签用于结构体，通常用于命令行参数选项。
	Brief             = "brief"        // 结构体的简短标签，通常被视为概述或摘要。
	Root              = "root"         // Root标签用于结构体，通常用于嵌套命令的管理。
	Additional        = "additional"   // 结构体的附加标签，通常用于命令的额外描述。
	AdditionalShort   = "ad"           // Additional的简称。
	Path              = `path`         // HTTP请求的路由路径。
	Method            = `method`       // Route 方法用于处理 HTTP 请求。
	Domain            = `domain`       // 为HTTP请求路由域名。
	Mime              = `mime`         // HTTP请求/响应的MIME类型。
	Consumes          = `consumes`     // HTTP请求的MIME类型。
	Summary           = `summary`      // Summary 用于结构体，通常用于OpenAPI中的请求结构体。
	SummaryShort      = `sm`           // Summary的简称。
	SummaryShort2     = `sum`          // Summary的简称。
	Description       = `description`  // 结构体描述，通常用于OpenAPI中的请求结构体。
	DescriptionShort  = `dc`           // 描述的简称。
	DescriptionShort2 = `des`          // 描述的简称。
	Example           = `example`      // 示例：用于结构体，通常用于OpenAPI中的请求结构体。
	ExampleShort      = `eg`           // Example 的简称。
	Examples          = `examples`     // 示例代码展示结构体的用法，通常用于OpenAPI中的请求结构体。
	ExamplesShort     = `egs`          // 示例的简称。
	ExternalDocs      = `externalDocs` // 结构体外部文档，始终用于OpenAPI中的请求结构体。
	ExternalDocsShort = `ed`           // ExternalDocs 的简称。
	GConv             = "gconv"        // GConv 定义了为特定结构体字段指定的转换目标名称。
	GConvShort        = "c"            // GConv 定义了为特定结构体字段指定的转换目标名称。
	Json              = "json"         // Json标签受到stdlib的支持。
	Security          = "security"     // Security 定义了身份验证方案。详情请参阅 https://swagger.io/docs/specification/authentication/
	In                = "in"           // Swagger 根据参数位置区分以下几种参数类型。详情请参阅 https://swagger.io/docs/specification/describing-parameters/
// 在Swagger中，根据参数所在的位置，对参数进行如下类型的区别定义
)
