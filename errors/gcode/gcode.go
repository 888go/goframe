// 版权所有 GoFrame gf 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一个。

// Package gcode 提供通用错误码定义及常见错误码实现。
package gcode

// Code 是通用错误码接口的定义。
type Code interface {
	// Code 返回当前错误代码的整数值。
	Code() int

	// Message 返回当前错误代码的简短消息。
	Message() string

// Detail 返回当前错误代码的详细信息，
// 主要设计为错误代码的扩展字段。
	Detail() interface{}
}

// ================================================================================================================
// 公共错误码定义。
// 框架内部预留了以下错误码：code < 1000。
// ================================================================================================================

var (
	CodeNil                       = localCode{-1, "", nil}                             // 未指定错误代码。
	CodeOK                        = localCode{0, "OK", nil}                            // It is OK.
	CodeInternalError             = localCode{50, "Internal Error", nil}               // 发生了内部错误。
	CodeValidationFailed          = localCode{51, "Validation Failed", nil}            // 数据验证失败。
	CodeDbOperationError          = localCode{52, "Database Operation Error", nil}     // 数据库操作错误。
	CodeInvalidParameter          = localCode{53, "Invalid Parameter", nil}            // 当前操作给定的参数无效。
	CodeMissingParameter          = localCode{54, "Missing Parameter", nil}            // 当前操作缺少参数。
	CodeInvalidOperation          = localCode{55, "Invalid Operation", nil}            // 该函数不能这样使用。
	CodeInvalidConfiguration      = localCode{56, "Invalid Configuration", nil}        // 当前操作的配置无效。
	CodeMissingConfiguration      = localCode{57, "Missing Configuration", nil}        // 当前操作缺少配置。
	CodeNotImplemented            = localCode{58, "Not Implemented", nil}              // 该操作尚未实现。
	CodeNotSupported              = localCode{59, "Not Supported", nil}                // 该操作尚未被支持。
	CodeOperationFailed           = localCode{60, "Operation Failed", nil}             // 我尝试了，但我无法给你你想要的东西。
	CodeNotAuthorized             = localCode{61, "Not Authorized", nil}               // Not Authorized.
	CodeSecurityReason            = localCode{62, "Security Reason", nil}              // Security Reason.
	CodeServerBusy                = localCode{63, "Server Is Busy", nil}               // 服务器繁忙，请稍后再试。
	CodeUnknown                   = localCode{64, "Unknown Error", nil}                // Unknown error.
	CodeNotFound                  = localCode{65, "Not Found", nil}                    // 资源不存在。
	CodeInvalidRequest            = localCode{66, "Invalid Request", nil}              // Invalid request.
	CodeNecessaryPackageNotImport = localCode{67, "Necessary Package Not Import", nil} // 它需要必要的包导入。
	CodeInternalPanic             = localCode{68, "Internal Panic", nil}               // 发生了内部 panic
	CodeBusinessValidationFailed  = localCode{300, "Business Validation Failed", nil}  // 业务验证失败。
)

// New 创建并返回一个错误代码。
// 注意，它返回的是 Code 接口对象。
func New(code int, message string, detail interface{}) Code {
	return localCode{
		code:    code,
		message: message,
		detail:  detail,
	}
}

// WithCode根据给定的Code创建并返回一个新的错误代码。
// 代码和消息来自给定的`code`，但详情来自给定的`detail`。
func WithCode(code Code, detail interface{}) Code {
	return localCode{
		code:    code.Code(),
		message: code.Message(),
		detail:  detail,
	}
}
