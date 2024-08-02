// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

//go:build windows
// +build windows

package http类

import (
	"context"
	"os"

	gproc "github.com/888go/goframe/os/gproc"
)

// handleProcessSignal handles all signals from system in blocking way.
func handleProcessSignal() {
	var ctx = context.TODO()
	gproc.AddSigHandlerShutdown(func(sig os.Signal) {
		shutdownWebServersGracefully(ctx, sig)
	})

	gproc.Listen()
}
