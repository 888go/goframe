// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	gtype "github.com/888go/goframe/container/gtype"
	gjson "github.com/888go/goframe/encoding/gjson"
	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	glog "github.com/888go/goframe/os/glog"
	gproc "github.com/888go/goframe/os/gproc"
	gtime "github.com/888go/goframe/os/gtime"
	gtimer "github.com/888go/goframe/os/gtimer"
	gstr "github.com/888go/goframe/text/gstr"
	gconv "github.com/888go/goframe/util/gconv"
)

const (
		// 允许在服务器启动后经过此毫秒间隔执行管理命令。 md5:0a7e1d2b4fe2af39
	adminActionIntervalLimit = 2000
	adminActionNone          = 0
	adminActionRestarting    = 1
	adminActionShuttingDown  = 2
	adminActionReloadEnvKey  = "GF_SERVER_RELOAD"
	adminActionRestartEnvKey = "GF_SERVER_RESTART"
	adminGProcCommGroup      = "GF_GPROC_HTTP_SERVER"
)

var (
		// serverActionLocker 是用于服务器管理操作的锁。 md5:3de592f90d7f4ae4
	serverActionLocker sync.Mutex

		// serverActionLastTime 是上一次管理操作的时间戳（以毫秒为单位）。 md5:df445bcc172577e2
	serverActionLastTime = gtype.NewInt64(gtime.X取时间戳毫秒())

		// serverProcessStatus是当前进程运行时的服务器状态。 md5:3d55829242522190
	serverProcessStatus = gtype.NewInt()
)

// X平滑重启所有服务 优雅地重启进程中的所有服务器。
// 可选参数 `newExeFilePath` 指定了用于创建进程的新二进制文件路径。
// md5:cd148e150eddefe2
func X平滑重启所有服务(上下文 context.Context, 新可执行文件路径 string) error {
	if !gracefulEnabled {
		return gerror.X创建错误码(gcode.CodeInvalidOperation, "graceful reload feature is disabled")
	}
	serverActionLocker.Lock()
	defer serverActionLocker.Unlock()
	if err := checkProcessStatus(); err != nil {
		return err
	}
	if err := checkActionFrequency(); err != nil {
		return err
	}
	return restartWebServers(上下文, nil, 新可执行文件路径)
}

// X平滑关闭所有服务 优雅地关闭当前进程中的所有服务器。 md5:1eb1bf001c79c66c
func X平滑关闭所有服务(上下文 context.Context) error {
	serverActionLocker.Lock()
	defer serverActionLocker.Unlock()
	if err := checkProcessStatus(); err != nil {
		return err
	}
	if err := checkActionFrequency(); err != nil {
		return err
	}
	shutdownWebServersGracefully(上下文, nil)
	return nil
}

// checkProcessStatus 检查当前进程的服务器状态。 md5:f49e9c4fdac4de86
func checkProcessStatus() error {
	status := serverProcessStatus.X取值()
	if status > 0 {
		switch status {
		case adminActionRestarting:
			return gerror.X创建错误码(gcode.CodeInvalidOperation, "server is restarting")

		case adminActionShuttingDown:
			return gerror.X创建错误码(gcode.CodeInvalidOperation, "server is shutting down")
		}
	}
	return nil
}

// checkActionFrequency 检查操作频率。
// 如果频率过高，它会返回错误。
// md5:b5db2b4c0ba2cdf7
func checkActionFrequency() error {
	interval := gtime.X取时间戳毫秒() - serverActionLastTime.X取值()
	if interval < adminActionIntervalLimit {
		return gerror.X创建错误码并格式化(
			gcode.CodeInvalidOperation,
			"too frequent action, please retry in %d ms",
			adminActionIntervalLimit-interval,
		)
	}
	serverActionLastTime.X设置值(gtime.X取时间戳毫秒())
	return nil
}

// forkReloadProcess 创建一个新的子进程，并将文件描述符复制到子进程中。 md5:5de49cf62f76603e
func forkReloadProcess(ctx context.Context, newExeFilePath ...string) error {
	var (
		path = os.Args[0]
	)
	if len(newExeFilePath) > 0 && newExeFilePath[0] != "" {
		path = newExeFilePath[0]
	}
	var (
		p   = gproc.NewProcess(path, os.Args, os.Environ())
		sfm = getServerFdMap()
	)
	for name, m := range sfm {
		for fdk, fdv := range m {
			if len(fdv) > 0 {
				s := ""
				for _, item := range gstr.X分割并忽略空值(fdv, ",") {
					array := strings.Split(item, "#")
					fd := uintptr(gconv.X取正整数(array[1]))
					if fd > 0 {
						s += fmt.Sprintf("%s#%d,", array[0], 3+len(p.ExtraFiles))
						p.ExtraFiles = append(p.ExtraFiles, os.NewFile(fd, ""))
					} else {
						s += fmt.Sprintf("%s#%d,", array[0], 0)
					}
				}
				sfm[name][fdk] = strings.TrimRight(s, ",")
			}
		}
	}
	buffer, _ := gjson.X变量到json字节集(sfm)
	p.Env = append(p.Env, adminActionReloadEnvKey+"="+string(buffer))
	if _, err := p.Start(ctx); err != nil {
		glog.X输出并格式化ERR(
			ctx,
			"%d: fork process failed, error:%s, %s",
			gproc.Pid(), err.Error(), string(buffer),
		)
		return err
	}
	return nil
}

// forkRestartProcess 创建一个新的服务器进程。 md5:f786ce6758d0d9ed
func forkRestartProcess(ctx context.Context, newExeFilePath ...string) error {
	var (
		path = os.Args[0]
	)
	if len(newExeFilePath) > 0 && newExeFilePath[0] != "" {
		path = newExeFilePath[0]
	}
	if err := os.Unsetenv(adminActionReloadEnvKey); err != nil {
		intlog.Errorf(ctx, `%+v`, err)
	}
	env := os.Environ()
	env = append(env, adminActionRestartEnvKey+"=1")
	p := gproc.NewProcess(path, os.Args, env)
	if _, err := p.Start(ctx); err != nil {
		glog.X输出并格式化ERR(
			ctx,
			`%d: fork process failed, error:%s, are you running using "go run"?`,
			gproc.Pid(), err.Error(),
		)
		return err
	}
	return nil
}

// getServerFdMap 返回所有服务器名称到文件描述符映射的map。 md5:dd5b6c5b0372c1b6
func getServerFdMap() map[string]listenerFdMap {
	sfm := make(map[string]listenerFdMap)
	serverMapping.X遍历读锁定(func(m map[string]interface{}) {
		for k, v := range m {
			sfm[k] = v.(*X服务).getListenerFdMap()
		}
	})
	return sfm
}

// bufferToServerFdMap 将二进制内容转换为fd映射。 md5:f02ae7f98f43f216
func bufferToServerFdMap(buffer []byte) map[string]listenerFdMap {
	sfm := make(map[string]listenerFdMap)
	if len(buffer) > 0 {
		j, _ := gjson.X加载并自动识别格式(buffer)
		for k := range j.X取泛型类().X取Map() {
			m := make(map[string]string)
			for mapKey, mapValue := range j.X取值(k).X取文本Map() {
				m[mapKey] = mapValue
			}
			sfm[k] = m
		}
	}
	return sfm
}

// restartWebServers 重启所有服务器。 md5:cad06ab5309d1725
func restartWebServers(ctx context.Context, signal os.Signal, newExeFilePath string) error {
	serverProcessStatus.X设置值(adminActionRestarting)
	if runtime.GOOS == "windows" {
		if signal != nil {
			// Controlled by signal.
			forceCloseWebServers(ctx)
			if err := forkRestartProcess(ctx, newExeFilePath); err != nil {
				intlog.Errorf(ctx, `%+v`, err)
			}
			return nil
		}
		// 由网页控制。
		// 它应该确保响应已写入客户端，然后优雅地关闭所有服务器。
		// md5:a5b2bfe1eb0f3681
		gtimer.SetTimeout别名(ctx, time.Second, func(ctx context.Context) {
			forceCloseWebServers(ctx)
			if err := forkRestartProcess(ctx, newExeFilePath); err != nil {
				intlog.Errorf(ctx, `%+v`, err)
			}
		})
		return nil
	}
	if err := forkReloadProcess(ctx, newExeFilePath); err != nil {
		glog.X输出并格式化(ctx, "%d: server restarts failed", gproc.Pid())
		serverProcessStatus.X设置值(adminActionNone)
		return err
	} else {
		if signal != nil {
			glog.X输出并格式化(ctx, "%d: server restarting by signal: %s", gproc.Pid(), signal)
		} else {
			glog.X输出并格式化(ctx, "%d: server restarting by web admin", gproc.Pid())
		}
	}

	return nil
}

// shutdownWebServersGracefully 延长关闭所有服务器。 md5:032a0ea9c1919f82
func shutdownWebServersGracefully(ctx context.Context, signal os.Signal) {
	serverProcessStatus.X设置值(adminActionShuttingDown)
	if signal != nil {
		glog.X输出并格式化(
			ctx,
			"%d: server gracefully shutting down by signal: %s",
			gproc.Pid(), signal.String(),
		)
	} else {
		glog.X输出并格式化(ctx, "%d: server gracefully shutting down by api", gproc.Pid())
	}
	serverMapping.X遍历读锁定(func(m map[string]interface{}) {
		for _, v := range m {
			server := v.(*X服务)
			server.doServiceDeregister()
			for _, s := range server.servers {
				s.shutdown(ctx)
			}
		}
	})
}

// forceCloseWebServers 强制关闭所有服务器。 md5:e7c5bd88a9acbd9e
func forceCloseWebServers(ctx context.Context) {
	serverMapping.X遍历读锁定(func(m map[string]interface{}) {
		for _, v := range m {
			for _, s := range v.(*X服务).servers {
				s.close(ctx)
			}
		}
	})
}

// handleProcessMessage 接收并处理来自进程的消息，
// 这通常用于实现优雅重启功能。
// md5:80ebd3c82cd48199
func handleProcessMessage() {
	var (
		ctx = context.TODO()
	)
	for {
		if msg := gproc.Receive(adminGProcCommGroup); msg != nil {
			if bytes.EqualFold(msg.Data, []byte("exit")) {
				intlog.Printf(ctx, "%d: process message: exit", gproc.Pid())
				shutdownWebServersGracefully(ctx, nil)
				allShutdownChan <- struct{}{}
				intlog.Printf(ctx, "%d: process message: exit done", gproc.Pid())
				return
			}
		}
	}
}
