// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package glog

import (
	"context"
	"io"
)

// SetConfig set configurations for the defaultLogger.

// ff:设置配置项
// config:配置项
func SetConfig(config Config) error {
	return defaultLogger.SetConfig(config)
}

// SetConfigWithMap set configurations with map for the defaultLogger.

// ff:设置配置Map
// m:
func SetConfigWithMap(m map[string]interface{}) error {
	return defaultLogger.SetConfigWithMap(m)
}

// SetPath sets the directory path for file logging.

// ff:设置文件路径
// path:文件路径
func SetPath(path string) error {
	return defaultLogger.SetPath(path)
}

// GetPath returns the logging directory path for file logging.
// It returns empty string if no directory path set.

// ff:取文件路径
func GetPath() string {
	return defaultLogger.GetPath()
}

// SetFile sets the file name `pattern` for file logging.
// Datetime pattern can be used in `pattern`, eg: access-{Ymd}.log.
// The default file name pattern is: Y-m-d.log, eg: 2018-01-01.log

// ff:设置文件名格式
// pattern:文件名格式
func SetFile(pattern string) {
	defaultLogger.SetFile(pattern)
}

// SetLevel sets the default logging level.

// ff:设置级别
// level:级别
func SetLevel(level int) {
	defaultLogger.SetLevel(level)
}

// GetLevel returns the default logging level value.

// ff:取级别
func GetLevel() int {
	return defaultLogger.GetLevel()
}

// SetWriter sets the customized logging `writer` for logging.
// The `writer` object should implements the io.Writer interface.
// Developer can use customized logging `writer` to redirect logging output to another service,
// eg: kafka, mysql, mongodb, etc.

// ff:设置Writer
// writer:
func SetWriter(writer io.Writer) {
	defaultLogger.SetWriter(writer)
}

// GetWriter returns the customized writer object, which implements the io.Writer interface.
// It returns nil if no customized writer set.

// ff:取Writer
func GetWriter() io.Writer {
	return defaultLogger.GetWriter()
}

// SetDebug enables/disables the debug level for default defaultLogger.
// The debug level is enabled in default.

// ff:设置debug
// debug:开启
func SetDebug(debug bool) {
	defaultLogger.SetDebug(debug)
}

// SetAsync enables/disables async logging output feature for default defaultLogger.

// ff:设置异步输出
// enabled:开启
func SetAsync(enabled bool) {
	defaultLogger.SetAsync(enabled)
}

// SetStdoutPrint sets whether ouptput the logging contents to stdout, which is true in default.

// ff:设置是否同时输出到终端
// enabled:开启
func SetStdoutPrint(enabled bool) {
	defaultLogger.SetStdoutPrint(enabled)
}

// SetHeaderPrint sets whether output header of the logging contents, which is true in default.

// ff:设置是否输出头信息
// enabled:开启
func SetHeaderPrint(enabled bool) {
	defaultLogger.SetHeaderPrint(enabled)
}

// SetPrefix sets prefix string for every logging content.
// Prefix is part of header, which means if header output is shut, no prefix will be output.

// ff:设置前缀
// prefix:前缀
func SetPrefix(prefix string) {
	defaultLogger.SetPrefix(prefix)
}

// SetFlags sets extra flags for logging output features.

// ff:设置额外标识
// flags:标识
func SetFlags(flags int) {
	defaultLogger.SetFlags(flags)
}

// GetFlags returns the flags of defaultLogger.

// ff:取标识
func GetFlags() int {
	return defaultLogger.GetFlags()
}

// SetCtxKeys sets the context keys for defaultLogger. The keys is used for retrieving values
// from context and printing them to logging content.
//
// Note that multiple calls of this function will overwrite the previous set context keys.

// ff:设置上下文名称
// keys:名称
func SetCtxKeys(keys ...interface{}) {
	defaultLogger.SetCtxKeys(keys...)
}

// GetCtxKeys retrieves and returns the context keys for logging.

// ff:取上下文名称
func GetCtxKeys() []interface{} {
	return defaultLogger.GetCtxKeys()
}

// PrintStack prints the caller stack,
// the optional parameter `skip` specify the skipped stack offset from the end point.

// ff:输出堆栈信息
// skip:偏移量
// ctx:上下文
func PrintStack(ctx context.Context, skip ...int) {
	defaultLogger.PrintStack(ctx, skip...)
}

// GetStack returns the caller stack content,
// the optional parameter `skip` specify the skipped stack offset from the end point.

// ff:取堆栈信息
// skip:偏移量
func GetStack(skip ...int) string {
	return defaultLogger.GetStack(skip...)
}

// SetStack enables/disables the stack feature in failure logging outputs.

// ff:设置堆栈跟踪
// enabled:开启
func SetStack(enabled bool) {
	defaultLogger.SetStack(enabled)
}

// SetLevelStr sets the logging level by level string.

// ff:设置文本级别
// levelStr:级别
func SetLevelStr(levelStr string) error {
	return defaultLogger.SetLevelStr(levelStr)
}

// SetLevelPrefix sets the prefix string for specified level.

// ff:设置级别前缀
// prefix:前缀
// level:级别
func SetLevelPrefix(level int, prefix string) {
	defaultLogger.SetLevelPrefix(level, prefix)
}

// SetLevelPrefixes sets the level to prefix string mapping for the defaultLogger.

// ff:设置级别前缀Map
// prefixes:前缀Map
func SetLevelPrefixes(prefixes map[int]string) {
	defaultLogger.SetLevelPrefixes(prefixes)
}

// GetLevelPrefix returns the prefix string for specified level.

// ff:取级别前缀
// level:级别
func GetLevelPrefix(level int) string {
	return defaultLogger.GetLevelPrefix(level)
}

// SetHandlers sets the logging handlers for default defaultLogger.

// ff:设置中间件
// handlers:处理函数
func SetHandlers(handlers ...Handler) {
	defaultLogger.SetHandlers(handlers...)
}

// SetWriterColorEnable sets the file logging with color

// ff:设置文件是否输出颜色
// enabled:开启
func SetWriterColorEnable(enabled bool) {
	defaultLogger.SetWriterColorEnable(enabled)
}
