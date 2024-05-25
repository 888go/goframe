// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gproc

import (
	"os"
	"syscall"
	"testing"
	"time"

	"github.com/gogf/gf/v2/test/gtest"
)

func Test_Signal(t *testing.T) {
	go Listen()

	// non shutdown signal
	gtest.C(t, func(t *gtest.T) {
		sigRec := make(chan os.Signal, 1)
		AddSigHandler(func(sig os.Signal) {
			sigRec <- sig
		}, syscall.SIGUSR1, syscall.SIGUSR2)

		sendSignal(syscall.SIGUSR1)
		select {
		case s := <-sigRec:
			t.AssertEQ(s, syscall.SIGUSR1)
			t.AssertEQ(false, isWaitChClosed())
		case <-time.After(time.Second):
			t.Error("signal SIGUSR1 handler timeout")
		}

		sendSignal(syscall.SIGUSR2)
		select {
		case s := <-sigRec:
			t.AssertEQ(s, syscall.SIGUSR2)
			t.AssertEQ(false, isWaitChClosed())
		case <-time.After(time.Second):
			t.Error("signal SIGUSR2 handler timeout")
		}

		sendSignal(syscall.SIGHUP)
		select {
		case <-sigRec:
			t.Error("signal SIGHUP should not be listen")
		case <-time.After(time.Millisecond * 100):
		}

		// multiple listen
		go Listen()
		go Listen()
		sendSignal(syscall.SIGUSR1)
		cnt := 0
		timeout := time.After(time.Second)
		for {
			select {
			case <-sigRec:
				cnt++
			case <-timeout:
				if cnt == 0 {
					t.Error("signal SIGUSR2 handler timeout")
				}
				if cnt != 1 {
					t.Error("multi Listen() repetitive execution")
				}
				return
			}
		}
	})

	// test shutdown signal
	gtest.C(t, func(t *gtest.T) {
		sigRec := make(chan os.Signal, 1)
		AddSigHandlerShutdown(func(sig os.Signal) {
			sigRec <- sig
		})

		sendSignal(syscall.SIGTERM)
		// wait the listen done
		time.Sleep(time.Second)

		select {
		case s := <-sigRec:
			t.AssertEQ(s, syscall.SIGTERM)
			t.AssertEQ(true, isWaitChClosed())
		case <-time.After(time.Second):
			t.Error("signal SIGUSR2 handler timeout")
		}
	})
}

func sendSignal(sig os.Signal) {
	signalChan <- sig
}

func isWaitChClosed() bool {
	select {
	case _, ok := <-waitChan:
		return !ok
	default:
		return false
	}
}
