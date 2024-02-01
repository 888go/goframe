// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp
import (
	"bytes"
	"context"
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
	
	"github.com/888go/goframe/container/gtype"
	"github.com/888go/goframe/encoding/gjson"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/os/glog"
	"github.com/888go/goframe/os/gproc"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/os/gtimer"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
	)
const (
	// 在服务器启动后，允许在此间隔（以毫秒为单位）后执行管理命令。
	adminActionIntervalLimit = 2000
	adminActionNone          = 0
	adminActionRestarting    = 1
	adminActionShuttingDown  = 2
	adminActionReloadEnvKey  = "GF_SERVER_RELOAD"
	adminActionRestartEnvKey = "GF_SERVER_RESTART"
	adminGProcCommGroup      = "GF_GPROC_HTTP_SERVER"
)

var (
	// serverActionLocker 是针对服务器管理操作的锁。
	serverActionLocker sync.Mutex

	// serverActionLastTime 是上一次管理操作发生时的时间戳，单位为毫秒。
	serverActionLastTime = gtype.NewInt64(gtime.TimestampMilli())

	// serverProcessStatus 是当前进程运行操作的服务器状态。
	serverProcessStatus = gtype.NewInt()
)

// RestartAllServer 将优雅地重启进程中的所有服务器。
// 可选参数 `newExeFilePath` 指定了用于创建新进程的二进制文件。
func RestartAllServer(ctx context.Context, newExeFilePath string) error {
	if !gracefulEnabled {
		return gerror.NewCode(gcode.CodeInvalidOperation, "graceful reload feature is disabled")
	}
	serverActionLocker.Lock()
	defer serverActionLocker.Unlock()
	if err := checkProcessStatus(); err != nil {
		return err
	}
	if err := checkActionFrequency(); err != nil {
		return err
	}
	return restartWebServers(ctx, nil, newExeFilePath)
}

// ShutdownAllServer 将当前进程中的所有服务器优雅地关闭。
func ShutdownAllServer(ctx context.Context) error {
	serverActionLocker.Lock()
	defer serverActionLocker.Unlock()
	if err := checkProcessStatus(); err != nil {
		return err
	}
	if err := checkActionFrequency(); err != nil {
		return err
	}
	shutdownWebServersGracefully(ctx, nil)
	return nil
}

// checkProcessStatus 检查当前进程的服务器状态。
func checkProcessStatus() error {
	status := serverProcessStatus.Val()
	if status > 0 {
		switch status {
		case adminActionRestarting:
			return gerror.NewCode(gcode.CodeInvalidOperation, "server is restarting")

		case adminActionShuttingDown:
			return gerror.NewCode(gcode.CodeInvalidOperation, "server is shutting down")
		}
	}
	return nil
}

// checkActionFrequency 检查操作频率。
// 如果操作过于频繁，则返回错误。
func checkActionFrequency() error {
	interval := gtime.TimestampMilli() - serverActionLastTime.Val()
	if interval < adminActionIntervalLimit {
		return gerror.NewCodef(
			gcode.CodeInvalidOperation,
			"too frequent action, please retry in %d ms",
			adminActionIntervalLimit-interval,
		)
	}
	serverActionLastTime.Set(gtime.TimestampMilli())
	return nil
}

// forkReloadProcess 创建一个新的子进程，并将文件描述符(fd)复制到子进程中。
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
				for _, item := range gstr.SplitAndTrim(fdv, ",") {
					array := strings.Split(item, "#")
					fd := uintptr(gconv.Uint(array[1]))
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
	buffer, _ := gjson.Encode(sfm)
	p.Env = append(p.Env, adminActionReloadEnvKey+"="+string(buffer))
	if _, err := p.Start(ctx); err != nil {
		glog.Errorf(
			ctx,
			"%d: fork process failed, error:%s, %s",
			gproc.Pid(), err.Error(), string(buffer),
		)
		return err
	}
	return nil
}

// forkRestartProcess 创建一个新的服务进程。
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
		glog.Errorf(
			ctx,
			`%d: fork process failed, error:%s, are you running using "go run"?`,
			gproc.Pid(), err.Error(),
		)
		return err
	}
	return nil
}

// getServerFdMap 返回一个映射，其中包含了所有服务器名称到文件描述符的映射关系，以map形式返回。
func getServerFdMap() map[string]listenerFdMap {
	sfm := make(map[string]listenerFdMap)
	serverMapping.RLockFunc(func(m map[string]interface{}) {
		for k, v := range m {
			sfm[k] = v.(*Server).getListenerFdMap()
		}
	})
	return sfm
}

// bufferToServerFdMap 将二进制内容转换为文件描述符映射。
func bufferToServerFdMap(buffer []byte) map[string]listenerFdMap {
	sfm := make(map[string]listenerFdMap)
	if len(buffer) > 0 {
		j, _ := gjson.LoadContent(buffer)
		for k := range j.Var().Map() {
			m := make(map[string]string)
			for mapKey, mapValue := range j.Get(k).MapStrStr() {
				m[mapKey] = mapValue
			}
			sfm[k] = m
		}
	}
	return sfm
}

// 重启Web服务器 restartWebServers 函数会重启所有服务器。
func restartWebServers(ctx context.Context, signal os.Signal, newExeFilePath string) error {
	serverProcessStatus.Set(adminActionRestarting)
	if runtime.GOOS == "windows" {
		if signal != nil {
			// 由信号控制。
			forceCloseWebServers(ctx)
			if err := forkRestartProcess(ctx, newExeFilePath); err != nil {
				intlog.Errorf(ctx, `%+v`, err)
			}
			return nil
		}
// 由网页控制。
// 应确保响应已写入客户端，然后优雅地关闭所有服务器。
		gtimer.SetTimeout(ctx, time.Second, func(ctx context.Context) {
			forceCloseWebServers(ctx)
			if err := forkRestartProcess(ctx, newExeFilePath); err != nil {
				intlog.Errorf(ctx, `%+v`, err)
			}
		})
		return nil
	}
	if err := forkReloadProcess(ctx, newExeFilePath); err != nil {
		glog.Printf(ctx, "%d: server restarts failed", gproc.Pid())
		serverProcessStatus.Set(adminActionNone)
		return err
	} else {
		if signal != nil {
			glog.Printf(ctx, "%d: server restarting by signal: %s", gproc.Pid(), signal)
		} else {
			glog.Printf(ctx, "%d: server restarting by web admin", gproc.Pid())
		}
	}

	return nil
}

// shutdownWebServersGracefully 优雅地关闭所有服务器。
func shutdownWebServersGracefully(ctx context.Context, signal os.Signal) {
	serverProcessStatus.Set(adminActionShuttingDown)
	if signal != nil {
		glog.Printf(
			ctx,
			"%d: server gracefully shutting down by signal: %s",
			gproc.Pid(), signal.String(),
		)
	} else {
		glog.Printf(ctx, "%d: server gracefully shutting down by api", gproc.Pid())
	}
	serverMapping.RLockFunc(func(m map[string]interface{}) {
		for _, v := range m {
			server := v.(*Server)
			server.doServiceDeregister()
			for _, s := range server.servers {
				s.shutdown(ctx)
			}
		}
	})
}

// forceCloseWebServers 强制关闭所有服务器。
func forceCloseWebServers(ctx context.Context) {
	serverMapping.RLockFunc(func(m map[string]interface{}) {
		for _, v := range m {
			for _, s := range v.(*Server).servers {
				s.close(ctx)
			}
		}
	})
}

// handleProcessMessage 接收并处理来自进程的消息，
// 这通常用于优雅重载功能。
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
