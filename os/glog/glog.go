// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package glog implements powerful and easy-to-use leveled logging functionality.
package glog

import (
	"context"

	"github.com/gogf/gf/v2/internal/command"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/util/gconv"
)

// ILogger is the API interface for logger.
type ILogger interface {
	Print(ctx context.Context, v ...interface{})//qm:输出  cz:Print(  
	Printf(ctx context.Context, format string, v ...interface{})//qm:输出并格式化  cz:Printf(  
	Debug(ctx context.Context, v ...interface{})//qm:输出DEBU  cz:Debug(  
	Debugf(ctx context.Context, format string, v ...interface{})//qm:输出并格式化DEBU  cz:Debugf(  
	Info(ctx context.Context, v ...interface{})//qm:输出INFO  cz:Info(  
	Infof(ctx context.Context, format string, v ...interface{})//qm:输出并格式化INFO  cz:Infof(  
	Notice(ctx context.Context, v ...interface{})//qm:输出NOTI  cz:Notice(  
	Noticef(ctx context.Context, format string, v ...interface{})//qm:输出并格式化NOTI  cz:Noticef(  
	Warning(ctx context.Context, v ...interface{})//qm:输出WARN  cz:Warning(  
	Warningf(ctx context.Context, format string, v ...interface{})//qm:输出并格式化WARN  cz:Warningf(  
	Error(ctx context.Context, v ...interface{})
	Errorf(ctx context.Context, format string, v ...interface{})//qm:输出并格式化ERR  cz:Errorf(  
	Critical(ctx context.Context, v ...interface{})//qm:输出CRIT  cz:Critical(  
	Criticalf(ctx context.Context, format string, v ...interface{})//qm:输出并格式化CRIT  cz:Criticalf(  
	Panic(ctx context.Context, v ...interface{})//qm:输出PANI  cz:Panic(  
	Panicf(ctx context.Context, format string, v ...interface{})//qm:输出并格式化PANI  cz:Panicf(  
	Fatal(ctx context.Context, v ...interface{})//qm:输出FATA  cz:Fatal(  
	Fatalf(ctx context.Context, format string, v ...interface{})//qm:输出并格式化FATA  cz:Fatalf(  
}

const (
	commandEnvKeyForDebug = "gf.glog.debug"
)

var (
	// Ensure Logger implements ILogger.
	_ ILogger = &Logger{}

	// Default logger object, for package method usage.
	defaultLogger = New()

	// Goroutine pool for async logging output.
	// It uses only one asynchronous worker to ensure log sequence.
	asyncPool = grpool.New(1)

	// defaultDebug enables debug level or not in default,
	// which can be configured using command option or system environment.
	defaultDebug = true
)

func init() {
	defaultDebug = gconv.Bool(command.GetOptWithEnv(commandEnvKeyForDebug, "true"))
	SetDebug(defaultDebug)
}

// DefaultLogger returns the default logger.
// ff:取默认日志类
func DefaultLogger() *Logger {
	return defaultLogger
}

// SetDefaultLogger sets the default logger for package glog.
// Note that there might be concurrent safety issue if calls this function
// in different goroutines.
// ff:设置默认日志类
// l:
func SetDefaultLogger(l *Logger) {
	defaultLogger = l
}
