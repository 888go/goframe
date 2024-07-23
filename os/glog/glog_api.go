// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package glog//bm:日志类

import "context"

// Print prints `v` with newline using fmt.Sprintln.
// The parameter `v` can be multiple variables.
// ff:输出
// ctx:上下文
// v:值
func Print(ctx context.Context, v ...interface{}) {
	defaultLogger.Print(ctx, v...)
}

// Printf prints `v` with format `format` using fmt.Sprintf.
// The parameter `v` can be multiple variables.
// ff:输出并格式化
// ctx:上下文
// format:格式
// v:值
func Printf(ctx context.Context, format string, v ...interface{}) {
	defaultLogger.Printf(ctx, format, v...)
}

// Fatal prints the logging content with [FATA] header and newline, then exit the current process.
// ff:输出FATA
// ctx:上下文
// v:值
func Fatal(ctx context.Context, v ...interface{}) {
	defaultLogger.Fatal(ctx, v...)
}

// Fatalf prints the logging content with [FATA] header, custom format and newline, then exit the current process.
// ff:输出并格式化FATA
// ctx:上下文
// format:格式
// v:值
func Fatalf(ctx context.Context, format string, v ...interface{}) {
	defaultLogger.Fatalf(ctx, format, v...)
}

// Panic prints the logging content with [PANI] header and newline, then panics.
// ff:输出PANI
// ctx:上下文
// v:值
func Panic(ctx context.Context, v ...interface{}) {
	defaultLogger.Panic(ctx, v...)
}

// Panicf prints the logging content with [PANI] header, custom format and newline, then panics.
// ff:输出并格式化PANI
// ctx:上下文
// format:格式
// v:值
func Panicf(ctx context.Context, format string, v ...interface{}) {
	defaultLogger.Panicf(ctx, format, v...)
}

// Info prints the logging content with [INFO] header and newline.
// ff:输出INFO
// ctx:上下文
// v:值
func Info(ctx context.Context, v ...interface{}) {
	defaultLogger.Info(ctx, v...)
}

// Infof prints the logging content with [INFO] header, custom format and newline.
// ff:输出并格式化INFO
// ctx:上下文
// format:格式
// v:值
func Infof(ctx context.Context, format string, v ...interface{}) {
	defaultLogger.Infof(ctx, format, v...)
}

// Debug prints the logging content with [DEBU] header and newline.
// ff:输出DEBU
// ctx:上下文
// v:值
func Debug(ctx context.Context, v ...interface{}) {
	defaultLogger.Debug(ctx, v...)
}

// Debugf prints the logging content with [DEBU] header, custom format and newline.
// ff:输出并格式化DEBU
// ctx:上下文
// format:格式
// v:值
func Debugf(ctx context.Context, format string, v ...interface{}) {
	defaultLogger.Debugf(ctx, format, v...)
}

// Notice prints the logging content with [NOTI] header and newline.
// It also prints caller stack info if stack feature is enabled.
// ff:输出NOTI
// ctx:上下文
// v:值
func Notice(ctx context.Context, v ...interface{}) {
	defaultLogger.Notice(ctx, v...)
}

// Noticef prints the logging content with [NOTI] header, custom format and newline.
// It also prints caller stack info if stack feature is enabled.
// ff:输出并格式化NOTI
// ctx:上下文
// format:格式
// v:值
func Noticef(ctx context.Context, format string, v ...interface{}) {
	defaultLogger.Noticef(ctx, format, v...)
}

// Warning prints the logging content with [WARN] header and newline.
// It also prints caller stack info if stack feature is enabled.
// ff:输出WARN
// ctx:上下文
// v:值
func Warning(ctx context.Context, v ...interface{}) {
	defaultLogger.Warning(ctx, v...)
}

// Warningf prints the logging content with [WARN] header, custom format and newline.
// It also prints caller stack info if stack feature is enabled.
// ff:输出并格式化WARN
// ctx:上下文
// format:格式
// v:值
func Warningf(ctx context.Context, format string, v ...interface{}) {
	defaultLogger.Warningf(ctx, format, v...)
}

// Error prints the logging content with [ERRO] header and newline.
// It also prints caller stack info if stack feature is enabled.
// ff:
// ctx:上下文
// v:值
func Error(ctx context.Context, v ...interface{}) {
	defaultLogger.Error(ctx, v...)
}

// Errorf prints the logging content with [ERRO] header, custom format and newline.
// It also prints caller stack info if stack feature is enabled.
// ff:输出并格式化ERR
// ctx:上下文
// format:格式
// v:值
func Errorf(ctx context.Context, format string, v ...interface{}) {
	defaultLogger.Errorf(ctx, format, v...)
}

// Critical prints the logging content with [CRIT] header and newline.
// It also prints caller stack info if stack feature is enabled.
// ff:输出CRIT
// ctx:上下文
// v:值
func Critical(ctx context.Context, v ...interface{}) {
	defaultLogger.Critical(ctx, v...)
}

// Criticalf prints the logging content with [CRIT] header, custom format and newline.
// It also prints caller stack info if stack feature is enabled.
// ff:输出并格式化CRIT
// ctx:上下文
// format:格式
// v:值
func Criticalf(ctx context.Context, format string, v ...interface{}) {
	defaultLogger.Criticalf(ctx, format, v...)
}
