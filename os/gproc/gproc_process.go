// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

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

	"github.com/gogf/gf/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/intlog"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/genv"
	"github.com/gogf/gf/v2/text/gstr"
)

// Process 是表示单个进程的结构体。 md5:f6524ce6eee4a18b
type Process struct {
	exec.Cmd
	Manager *Manager
	PPid    int
}

// NewProcess 创建并返回一个新的 Process。 md5:dbd46312fa39f087
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
				// 排除当前二进制文件的路径。 md5:a174ba38ac49d432
		start := 0
		if strings.EqualFold(path, args[0]) {
			start = 1
		}
		process.Args = append(process.Args, args[start:]...)
	}
	return process
}

// NewProcessCmd 创建并返回一个具有给定命令和可选环境变量数组的进程。 md5:01376a1e29c9935e
func NewProcessCmd(cmd string, environment ...[]string) *Process {
	return NewProcess(getShell(), append([]string{getShellOption()}, parseCommand(cmd)...), environment...)
}

// Start 以非阻塞方式开始执行进程。
// 如果成功，它将返回进程ID（pid），否则返回一个错误。
// md5:4607fc00f35e6338
func (p *Process) Start(ctx context.Context) (int, error) {
	if p.Process != nil {
		return p.Pid(), nil
	}
		// 为命令提供OpenTelemetry。 md5:46407dd5b38f692f
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

		// OpenTelemetry 传播。 md5:aecf3a0cccd13f96
	tracingEnv := tracingEnvFromCtx(ctx)
	if len(tracingEnv) > 0 {
		p.Env = append(p.Env, tracingEnv...)
	}
	p.Env = append(p.Env, fmt.Sprintf("%s=%d", envKeyPPid, p.PPid))
	p.Env = genv.Filter(p.Env)

		// 在 Windows 系统中，这可以工作，但在其他平台则无法工作. md5:9aac240ca7d717fe
	if runtime.GOOS == "windows" {
		joinProcessArgs(p)
	}

	if err := p.Cmd.Start(); err == nil {
		if p.Manager != nil {
			p.Manager.processes.Set(p.Process.Pid, p)
		}
		return p.Process.Pid, nil
	} else {
		return 0, err
	}
}

// Run以阻塞方式执行进程。 md5:aeab1ddf5fca3d31
func (p *Process) Run(ctx context.Context) error {
	if _, err := p.Start(ctx); err == nil {
		return p.Wait()
	} else {
		return err
	}
}

// Pid 获取并返回进程的PID。 md5:7f6e89391a9d1aac
func (p *Process) Pid() int {
	if p.Process != nil {
		return p.Process.Pid
	}
	return 0
}

// Send 向进程发送自定义数据。 md5:cb2381344fb13fd4
func (p *Process) Send(data []byte) error {
	if p.Process != nil {
		return Send(p.Process.Pid, data)
	}
	return gerror.NewCode(gcode.CodeInvalidParameter, "invalid process")
}

// Release 释放与进程 p 关联的任何资源，使其将来无法使用。
// 只有在不调用 Wait 的情况下才需要调用 Release。
// md5:f3540c25ba14f0ee
func (p *Process) Release() error {
	return p.Process.Release()
}

// Kill 立即导致 Process 终止。 md5:4bacb16ab3b9aebe
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
		// 它忽略这个错误，仅记录日志。 md5:578bff85a58d16e8
	_, err = p.Process.Wait()
	intlog.Errorf(context.TODO(), `%+v`, err)
	return nil
}

// Signal 向进程发送一个信号。
// 在Windows上发送Interrupt信号未实现。
// md5:c1afe56a9d236095
func (p *Process) Signal(sig os.Signal) error {
	return p.Process.Signal(sig)
}
