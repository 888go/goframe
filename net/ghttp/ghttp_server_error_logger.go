// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package ghttp

import (
	"bytes"
	"context"

	"github.com/gogf/gf/v2/os/glog"
)

// errorLogger is the error logging logger for underlying net/http.Server.
type errorLogger struct {
	logger *glog.Logger
}

// Write implements the io.Writer interface.
// ff:
// l:
// p:
// n:
// err:
func (l *errorLogger) Write(p []byte) (n int, err error) {
	l.logger.Skip(1).Error(context.TODO(), string(bytes.TrimRight(p, "\r\n")))
	return len(p), nil
}
