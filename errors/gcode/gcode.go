// 版权所有 (c) GoFrame (https://goframe.org)，保留所有权利。
//
// 本源代码遵循MIT许可协议。若未随此文件分发MIT许可证的副本，
// 您可以从 https://github.com/gogf/gf 获取。 md5:c14c707c81272457

// 包gcode提供了通用的错误代码定义和常见的错误代码实现。 md5:cb91541987c67096
package gcode

// Code 是一个通用错误代码接口的定义。 md5:bc72f9cd69a9f042
type Code interface {
	// Code 返回当前错误代码的整数值。 md5:75b8de0b4b9fa0a7
	Code() int

	// Message返回当前错误代码的简要消息。 md5:e0440d2d9a5b929c
	Message() string

	// Detail返回当前错误代码的详细信息，主要用于作为错误代码的扩展字段。 md5:b363ac7e7695be15
	Detail() interface{}
}

//==============================================================================================================
// 公共错误码定义。
// 框架保留了内部错误码的使用范围：代码小于1000。
//============================================================================================================== md5:aeebc2e4a8ad2666

var (
	CodeNil                       = localCode{-1, "", nil}                             // 没有指定错误代码。 md5:f3402e31e47f29a9
	CodeOK                        = localCode{0, "OK", nil}                            // It is OK.
	CodeInternalError             = localCode{50, "Internal Error", nil}               // 发生了内部错误。 md5:68452eba157c4f37
	CodeValidationFailed          = localCode{51, "Validation Failed", nil}            // 数据验证失败。 md5:9bd6126b3a2cb386
	CodeDbOperationError          = localCode{52, "Database Operation Error", nil}     // 数据库操作错误。 md5:67c037697b9e335d
	CodeInvalidParameter          = localCode{53, "Invalid Parameter", nil}            // 给定的当前操作参数无效。 md5:ca885036e7406885
	CodeMissingParameter          = localCode{54, "Missing Parameter", nil}            // 当前操作缺少参数。 md5:1ce758fa97191ebc
	CodeInvalidOperation          = localCode{55, "Invalid Operation", nil}            // 这个函数不能这样使用。 md5:a02d2635b1d2a487
	CodeInvalidConfiguration      = localCode{56, "Invalid Configuration", nil}        // 当前操作的配置无效。 md5:babdd505987f15c5
	CodeMissingConfiguration      = localCode{57, "Missing Configuration", nil}        // 当前操作的配置缺失。 md5:8f05e88006bb7f7f
	CodeNotImplemented            = localCode{58, "Not Implemented", nil}              // 此操作尚未实现。 md5:5277696d372ccedc
	CodeNotSupported              = localCode{59, "Not Supported", nil}                // 此操作尚不支持。 md5:90cc232b1a9aa21e
	CodeOperationFailed           = localCode{60, "Operation Failed", nil}             // 我尝试了，但是我无法给你你想要的。 md5:a8cb7ffbfd6211e4
	CodeNotAuthorized             = localCode{61, "Not Authorized", nil}               // Not Authorized.
	CodeSecurityReason            = localCode{62, "Security Reason", nil}              // Security Reason.
	CodeServerBusy                = localCode{63, "Server Is Busy", nil}               // 服务器正忙，请稍后重试。 md5:474334c09e329e2d
	CodeUnknown                   = localCode{64, "Unknown Error", nil}                // Unknown error.
	CodeNotFound                  = localCode{65, "Not Found", nil}                    // 资源不存在。 md5:4e9493277f9141d8
	CodeInvalidRequest            = localCode{66, "Invalid Request", nil}              // Invalid request.
	CodeNecessaryPackageNotImport = localCode{67, "Necessary Package Not Import", nil} // 它需要必要的包导入。 md5:bd34126e0df110ff
	CodeInternalPanic             = localCode{68, "Internal Panic", nil}               // 内部发生了恐慌。 md5:f12430dbb6bb9ee9
	CodeBusinessValidationFailed  = localCode{300, "Business Validation Failed", nil}  // 业务验证失败。 md5:816812c09b9bed71
)

// New 创建并返回一个错误代码。
// 注意，它返回一个 Code 接口对象。 md5:a3d7ec3807589165
func New(code int, message string, detail interface{}) Code {
	return localCode{
		code:    code,
		message: message,
		detail:  detail,
	}
}

// WithCode 根据给定的`Code`创建并返回一个新的错误代码。
// 该错误代码的码和消息来自`code`，但详细信息来自`detail`。 md5:6f2355f302e9ea32
func WithCode(code Code, detail interface{}) Code {
	return localCode{
		code:    code.Code(),
		message: code.Message(),
		detail:  detail,
	}
}
