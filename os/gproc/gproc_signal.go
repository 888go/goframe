// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gproc

import (
	"context"
	"github.com/gogf/gf/v2/internal/intlog"
	"github.com/gogf/gf/v2/util/gutil"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// SigHandler 定义了一个用于处理信号的函数类型。 md5:d7319108f37510cd
type SigHandler func(sig os.Signal)

var (
	// 使用内部变量来保证当发生多个监听（Listen）时的并发安全。
	// md5:7c2a9e90bd5be8da
	listenOnce        = sync.Once{}
	waitChan          = make(chan struct{})
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

// AddSigHandler 为自定义的一个或多个信号添加自定义的信号处理器。 md5:996226c8d75ebdf5
func AddSigHandler(handler SigHandler, signals ...os.Signal) {
	signalHandlerMu.Lock()
	defer signalHandlerMu.Unlock()
	for _, sig := range signals {
		signalHandlerMap[sig] = append(signalHandlerMap[sig], handler)
	}
	notifySignals()
}

// AddSigHandlerShutdown 为关闭信号添加自定义信号处理器：
// - syscall.SIGINT（中断信号，通常由Ctrl+C触发）
// - syscall.SIGQUIT（退出信号，通常通过Ctrl+\触发）
// - syscall.SIGKILL（杀死信号，不可被捕获或忽略，用于强制终止进程）
// - syscall.SIGTERM（终止信号，用来请求程序正常退出）
// - syscall.SIGABRT（异常终止信号，通常由调用abort函数产生，用于指示严重错误）
// md5:6fd417c58f499e80
func AddSigHandlerShutdown(handler ...SigHandler) {
	signalHandlerMu.Lock()
	defer signalHandlerMu.Unlock()
	for _, h := range handler {
		for sig := range shutdownSignalMap {
			signalHandlerMap[sig] = append(signalHandlerMap[sig], h)
		}
	}
	notifySignals()
}

// Listen 会阻塞并进行信号监听和处理。 md5:2425bc5d9026c36f
func Listen() {
	listenOnce.Do(func() {
		go listen()
	})

	<-waitChan
}

func listen() {
	defer close(waitChan)

	var (
		ctx = context.Background()
		wg  = sync.WaitGroup{}
		sig os.Signal
	)
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
				// 如果它是退出信号，它将退出此信号监听。 md5:9b1cb86f40c5e361
		if _, ok := shutdownSignalMap[sig]; ok {
			intlog.Printf(
				ctx,
				`receive shutdown signal "%s", waiting all signal handler done`,
				sig.String(),
			)
						// 等待信号处理器完成。 md5:4d1ee13d17a0a193
			wg.Wait()
			intlog.Print(ctx, `all signal handler done, exit process`)
			return
		}
	}
}

func notifySignals() {
	var signals = make([]os.Signal, 0)
	for s := range signalHandlerMap {
		signals = append(signals, s)
	}
	signal.Notify(signalChan, signals...)
}

func getHandlersBySignal(sig os.Signal) []SigHandler {
	signalHandlerMu.Lock()
	defer signalHandlerMu.Unlock()
	return signalHandlerMap[sig]
}
