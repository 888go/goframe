// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gproc

import (
	"os"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/errors/gerror"
)

// Manager 是一个管理多个进程的进程管理器。 md5:608ec304d3cca78a
type Manager struct {
	processes *gmap.IntAnyMap // 进程ID到进程对象的映射。 md5:e1aabd18695c16fa
}

// NewManager 创建并返回一个新的进程管理器。 md5:bfef06576c70f94f
func NewManager() *Manager {
	return &Manager{
		processes: gmap.NewIntAnyMap(true),
	}
}

// NewProcess 创建并返回一个进程对象。 md5:41e1fd6b109e05e7
func (m *Manager) NewProcess(path string, args []string, environment []string) *Process {
	p := NewProcess(path, args, environment)
	p.Manager = m
	return p
}

// GetProcess 获取并返回一个Process对象。
// 如果找不到具有给定`pid`的进程，它将返回nil。
// md5:d5b11d4d0e9fa1a3
func (m *Manager) GetProcess(pid int) *Process {
	if v := m.processes.Get(pid); v != nil {
		return v.(*Process)
	}
	return nil
}

// AddProcess 向当前管理器添加一个进程。
// 如果给定的 `pid` 对应的进程不存在，它不会做任何操作。
// md5:c51d5832fb1ce691
func (m *Manager) AddProcess(pid int) {
	if m.processes.Get(pid) == nil {
		if process, err := os.FindProcess(pid); err == nil {
			p := m.NewProcess("", nil, nil)
			p.Process = process
			m.processes.Set(pid, p)
		}
	}
}

// RemoveProcess 从当前管理器中移除一个进程。 md5:0076407de3a7d26a
func (m *Manager) RemoveProcess(pid int) {
	m.processes.Remove(pid)
}

// Processes 获取并返回当前管理器中的所有进程。 md5:30ac76e5c68d45de
func (m *Manager) Processes() []*Process {
	processes := make([]*Process, 0)
	m.processes.RLockFunc(func(m map[int]interface{}) {
		for _, v := range m {
			processes = append(processes, v.(*Process))
		}
	})
	return processes
}

// Pids 获取并返回当前管理器中的所有进程ID数组。 md5:a5ef21ec52c87400
func (m *Manager) Pids() []int {
	return m.processes.Keys()
}

// WaitAll等待直到所有进程退出。 md5:1d27f65463fe8c00
func (m *Manager) WaitAll() {
	processes := m.Processes()
	if len(processes) > 0 {
		for _, p := range processes {
			_ = p.Wait()
		}
	}
}

// KillAll 在当前管理器中杀死所有进程。 md5:337f683854b75187
func (m *Manager) KillAll() error {
	for _, p := range m.Processes() {
		if err := p.Kill(); err != nil {
			return err
		}
	}
	return nil
}

// SignalAll 向当前管理器中的所有进程发送信号 `sig`。 md5:64ce0027dcad8808
func (m *Manager) SignalAll(sig os.Signal) error {
	for _, p := range m.Processes() {
		if err := p.Signal(sig); err != nil {
			err = gerror.Wrapf(err, `send signal to process failed for pid "%d" with signal "%s"`, p.Process.Pid, sig)
			return err
		}
	}
	return nil
}

// Send 将数据字节发送到当前管理器中的所有进程。 md5:05d5ed3b0a5c7e3e
func (m *Manager) Send(data []byte) {
	for _, p := range m.Processes() {
		_ = p.Send(data)
	}
}

// SendTo 向当前管理器中的指定进程发送数据字节。 md5:b477f09d2f5cca5f
func (m *Manager) SendTo(pid int, data []byte) error {
	return Send(pid, data)
}

// Clear 会清除当前管理器中的所有进程。 md5:26053a86c2f65b33
func (m *Manager) Clear() {
	m.processes.Clear()
}

// Size 返回当前管理器中进程的数量。 md5:ffaeaa3ed9b66ed1
func (m *Manager) Size() int {
	return m.processes.Size()
}
