// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gproc
import (
	"os"
	
	"github.com/888go/goframe/container/gmap"
	"github.com/888go/goframe/errors/gerror"
	)
// Manager 是一个进程管理器，用于维护多个进程。
type Manager struct {
	processes *gmap.IntAnyMap // 进程id到进程对象的映射。
}

// NewManager 创建并返回一个新的进程管理器。
func NewManager() *Manager {
	return &Manager{
		processes: gmap.NewIntAnyMap(true),
	}
}

// NewProcess 创建并返回一个 Process 对象。
func (m *Manager) NewProcess(path string, args []string, environment []string) *Process {
	p := NewProcess(path, args, environment)
	p.Manager = m
	return p
}

// GetProcess 获取并返回一个 Process 对象。
// 如果未找到给定 `pid` 的进程，则返回 nil。
func (m *Manager) GetProcess(pid int) *Process {
	if v := m.processes.Get(pid); v != nil {
		return v.(*Process)
	}
	return nil
}

// AddProcess 将一个进程添加到当前管理器中。
// 如果给定 `pid` 的进程不存在，则不做任何操作。
func (m *Manager) AddProcess(pid int) {
	if m.processes.Get(pid) == nil {
		if process, err := os.FindProcess(pid); err == nil {
			p := m.NewProcess("", nil, nil)
			p.Process = process
			m.processes.Set(pid, p)
		}
	}
}

// RemoveProcess 从当前管理器中移除一个进程。
func (m *Manager) RemoveProcess(pid int) {
	m.processes.Remove(pid)
}

// Processes 获取并返回当前管理器中的所有进程。
func (m *Manager) Processes() []*Process {
	processes := make([]*Process, 0)
	m.processes.RLockFunc(func(m map[int]interface{}) {
		for _, v := range m {
			processes = append(processes, v.(*Process))
		}
	})
	return processes
}

// Pids 获取并返回当前管理器中的所有进程ID数组。
func (m *Manager) Pids() []int {
	return m.processes.Keys()
}

// WaitAll等待直到所有进程退出。
func (m *Manager) WaitAll() {
	processes := m.Processes()
	if len(processes) > 0 {
		for _, p := range processes {
			_ = p.Wait()
		}
	}
}

// KillAll杀掉当前管理器中的所有进程。
func (m *Manager) KillAll() error {
	for _, p := range m.Processes() {
		if err := p.Kill(); err != nil {
			return err
		}
	}
	return nil
}

// SignalAll 向当前管理器中的所有进程发送信号 `sig`。
func (m *Manager) SignalAll(sig os.Signal) error {
	for _, p := range m.Processes() {
		if err := p.Signal(sig); err != nil {
			err = gerror.Wrapf(err, `send signal to process failed for pid "%d" with signal "%s"`, p.Process.Pid, sig)
			return err
		}
	}
	return nil
}

// Send 向当前管理器中的所有进程发送 data 字节。
func (m *Manager) Send(data []byte) {
	for _, p := range m.Processes() {
		_ = p.Send(data)
	}
}

// SendTo 向当前管理器中指定进程发送数据字节。
func (m *Manager) SendTo(pid int, data []byte) error {
	return Send(pid, data)
}

// Clear 清除当前管理器中的所有进程。
func (m *Manager) Clear() {
	m.processes.Clear()
}

// Size 返回当前管理器中进程的数量。
func (m *Manager) Size() int {
	return m.processes.Size()
}
