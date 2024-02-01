// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gproc
import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/util/gutil"
	)
// SigHandler 定义了一个用于信号处理的函数类型。
type SigHandler func(sig os.Signal)

var (
// 使用内部变量来保证并发安全性
// 当发生多个Listen操作时。
	signalChan        = make(chan os.Signal, 1)
	signalHandlerMu   sync.Mutex
	signalHandlerMap  = make(map[os.Signal][]SigHandler)
	shutdownSignalMap = map[os.Signal]struct{}{
		syscall.SIGINT:  {},
		syscall.SIGQUIT: {},
		syscall.SIGKILL: {},
		syscall.SIGTERM: {},
		syscall.SIGABRT: {},
	}
)

func init() {
	for sig := range shutdownSignalMap {
		signalHandlerMap[sig] = make([]SigHandler, 0)
	}
}

// AddSigHandler 为一个或多个自定义信号添加自定义处理函数。
func AddSigHandler(handler SigHandler, signals ...os.Signal) {
	signalHandlerMu.Lock()
	defer signalHandlerMu.Unlock()
	for _, sig := range signals {
		signalHandlerMap[sig] = append(signalHandlerMap[sig], handler)
	}
}

// AddSigHandlerShutdown 添加自定义信号处理器以处理关闭信号：
// syscall.SIGINT（中断信号）
// syscall.SIGQUIT（退出并生成 core 文件信号）
// syscall.SIGKILL（强制终止信号，无法被捕获或忽略）
// syscall.SIGTERM（软件终止信号）
// syscall.SIGABRT（异常终止信号，如调用 abort 函数时触发）
func AddSigHandlerShutdown(handler ...SigHandler) {
	signalHandlerMu.Lock()
	defer signalHandlerMu.Unlock()
	for _, h := range handler {
		for sig := range shutdownSignalMap {
			signalHandlerMap[sig] = append(signalHandlerMap[sig], h)
		}
	}
}

// Listen阻塞并执行信号监听和处理。
func Listen() {
	var (
		signals = getHandlerSignals()
		ctx     = context.Background()
		wg      = sync.WaitGroup{}
		sig     os.Signal
	)
	signal.Notify(signalChan, signals...)
	for {
		sig = <-signalChan
		intlog.Printf(ctx, `signal received: %s`, sig.String())
		if handlers := getHandlersBySignal(sig); len(handlers) > 0 {
			for _, handler := range handlers {
				wg.Add(1)
				var (
					currentHandler = handler
					currentSig     = sig
				)
				gutil.TryCatch(ctx, func(ctx context.Context) {
					defer wg.Done()
					currentHandler(currentSig)
				}, func(ctx context.Context, exception error) {
					intlog.Errorf(ctx, `execute signal handler failed: %+v`, exception)
				})
			}
		}
		// 如果接收到的是关闭信号，则退出该信号监听。
		if _, ok := shutdownSignalMap[sig]; ok {
			intlog.Printf(
				ctx,
				`receive shutdown signal "%s", waiting all signal handler done`,
				sig.String(),
			)
			// 等待，直到信号处理器完成。
			wg.Wait()
			intlog.Print(ctx, `all signal handler done, exit process`)
			return
		}
	}
}

func getHandlerSignals() []os.Signal {
	signalHandlerMu.Lock()
	defer signalHandlerMu.Unlock()
	var signals = make([]os.Signal, 0)
	for s := range signalHandlerMap {
		signals = append(signals, s)
	}
	return signals
}

func getHandlersBySignal(sig os.Signal) []SigHandler {
	signalHandlerMu.Lock()
	defer signalHandlerMu.Unlock()
	return signalHandlerMap[sig]
}
