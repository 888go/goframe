// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gproc

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	
	"github.com/888go/goframe"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/net/gtrace"
	"github.com/888go/goframe/os/genv"
	"github.com/888go/goframe/text/gstr"
)

// Process是用于单个进程的结构体。
type Process struct {
	exec.Cmd
	Manager *Manager
	PPid    int
}

// NewProcess 创建并返回一个新的 Process。
func NewProcess(path string, args []string, environment ...[]string) *Process {
	env := os.Environ()
	if len(environment) > 0 {
		env = append(env, environment[0]...)
	}
	process := &Process{
		Manager: nil,
		PPid:    os.Getpid(),
		Cmd: exec.Cmd{
			Args:       []string{path},
			Path:       path,
			Stdin:      os.Stdin,
			Stdout:     os.Stdout,
			Stderr:     os.Stderr,
			Env:        env,
			ExtraFiles: make([]*os.File, 0),
		},
	}
	process.Dir, _ = os.Getwd()
	if len(args) > 0 {
		// 排除当前二进制文件路径。
		start := 0
		if strings.EqualFold(path, args[0]) {
			start = 1
		}
		process.Args = append(process.Args, args[start:]...)
	}
	return process
}

// NewProcessCmd根据给定的命令和可选的环境变量数组创建并返回一个进程。
func NewProcessCmd(cmd string, environment ...[]string) *Process {
	return NewProcess(getShell(), append([]string{getShellOption()}, parseCommand(cmd)...), environment...)
}

// Start以非阻塞方式启动进程执行。
// 如果成功，返回pid；否则返回错误。
func (p *Process) Start(ctx context.Context) (int, error) {
	if p.Process != nil {
		return p.Pid(), nil
	}
	// OpenTelemetry 用于命令。
	var (
		span trace.Span
		tr   = otel.GetTracerProvider().Tracer(
			tracingInstrumentName,
			trace.WithInstrumentationVersion(gf.VERSION),
		)
	)
	ctx, span = tr.Start(
		otel.GetTextMapPropagator().Extract(
			ctx,
			propagation.MapCarrier(genv.Map()),
		),
		gstr.Join(os.Args, " "),
		trace.WithSpanKind(trace.SpanKindInternal),
	)
	defer span.End()
	span.SetAttributes(gtrace.CommonLabels()...)

	// OpenTelemetry 传播
	tracingEnv := tracingEnvFromCtx(ctx)
	if len(tracingEnv) > 0 {
		p.Env = append(p.Env, tracingEnv...)
	}
	p.Env = append(p.Env, fmt.Sprintf("%s=%d", envKeyPPid, p.PPid))
	p.Env = genv.Filter(p.Env)

	if err := p.Cmd.Start(); err == nil {
		if p.Manager != nil {
			p.Manager.processes.Set(p.Process.Pid, p)
		}
		return p.Process.Pid, nil
	} else {
		return 0, err
	}
}

// Run以阻塞方式执行进程。
func (p *Process) Run(ctx context.Context) error {
	if _, err := p.Start(ctx); err == nil {
		return p.Wait()
	} else {
		return err
	}
}

// Pid 获取并返回当前进程的PID（进程标识符）
func (p *Process) Pid() int {
	if p.Process != nil {
		return p.Process.Pid
	}
	return 0
}

// Send 向进程发送自定义数据。
func (p *Process) Send(data []byte) error {
	if p.Process != nil {
		return Send(p.Process.Pid, data)
	}
	return gerror.NewCode(gcode.CodeInvalidParameter, "invalid process")
}

// Release 会释放与进程p关联的任何资源，
// 使其在未来无法使用。
// 只有在不调用Wait的情况下，才需要调用Release。
func (p *Process) Release() error {
	return p.Process.Release()
}

// Kill 导致 Process 立即退出。
func (p *Process) Kill() (err error) {
	err = p.Process.Kill()
	if err != nil {
		err = gerror.Wrapf(err, `kill process failed for pid "%d"`, p.Process.Pid)
		return err
	}
	if p.Manager != nil {
		p.Manager.processes.Remove(p.Pid())
	}
	if runtime.GOOS != "windows" {
		if err = p.Process.Release(); err != nil {
			intlog.Errorf(context.TODO(), `%+v`, err)
		}
	}
	// 它忽略这个错误，仅将其记录到日志中。
	_, err = p.Process.Wait()
	intlog.Errorf(context.TODO(), `%+v`, err)
	return nil
}

// Signal 向 Process 发送一个信号。
// 在 Windows 系统上发送 Interrupt 信号尚未实现。
func (p *Process) Signal(sig os.Signal) error {
	return p.Process.Signal(sig)
}
