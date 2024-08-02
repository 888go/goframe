// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

//go:build !windows
// +build !windows

package http类

import (
	"context"
	"os"
	"syscall"

	"github.com/888go/goframe/internal/intlog"
	glog "github.com/888go/goframe/os/glog"
	gproc "github.com/888go/goframe/os/gproc"
)

// handleProcessSignal handles all signals from system in blocking way.
func handleProcessSignal() {
	var ctx = context.TODO()
	gproc.AddSigHandlerShutdown(func(sig os.Signal) {
		shutdownWebServersGracefully(ctx, sig)
	})
	gproc.AddSigHandler(func(sig os.Signal) {
		// If the graceful restart feature is not enabled,
		// it does nothing except printing a warning log.
		if !gracefulEnabled {
			glog.Warning(ctx, "graceful reload feature is disabled")
			return
		}
		if err := restartWebServers(ctx, sig, ""); err != nil {
			intlog.Errorf(ctx, `%+v`, err)
		}
	}, syscall.SIGUSR1)

	gproc.Listen()
}
