// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gproc

import (
	"context"
	"fmt"
	"sync"
	
	"github.com/888go/goframe/container/gmap"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/net/gtcp"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/util/gconv"
)

// MsgRequest 是用于进程间通信的请求结构体。
type MsgRequest struct {
	SenderPid   int    // Sender PID.
	ReceiverPid int    // Receiver PID.
	Group       string // 消息组名称。
	Data        []byte // Request data.
}

// MsgResponse 是用于进程间通信的响应结构体。
type MsgResponse struct {
	Code    int    // 1: 表示成功；其它值：表示错误
	Message string // Response message.
	Data    []byte // Response data.
}

const (
	defaultFolderNameForProcComm = "gf_pid_port_mapping" // 默认的文件夹名称，用于存储pid到端口映射的文件。
	defaultGroupNameForProcComm  = ""                    // 默认分组名称。
	defaultTcpPortForProcComm    = 10000                 // 接收器监听的起始端口号。
	maxLengthForProcMsgQueue     = 10000                 // 每个组的消息队列的最大尺寸。
)

var (
// commReceiveQueues 是用于存储接收到数据的组名到队列映射。
// 该映射的值类型为 *gqueue.Queue。
	commReceiveQueues = gmap.NewStrAnyMap(true)

	// commPidFolderPath 指定存储进程ID到端口映射文件的文件夹路径。
	commPidFolderPath string

	// commPidFolderPathOnce 用于延迟计算，只有在必要时才计算 `commPidFolderPath`。
	commPidFolderPathOnce sync.Once
)

// getConnByPid 为指定的pid创建并返回一个TCP连接。
func getConnByPid(pid int) (*gtcp.PoolConn, error) {
	port := getPortByPid(pid)
	if port > 0 {
		if conn, err := gtcp.NewPoolConn(fmt.Sprintf("127.0.0.1:%d", port)); err == nil {
			return conn, nil
		} else {
			return nil, err
		}
	}
	return nil, gerror.Newf(`could not find port for pid "%d"`, pid)
}

// getPortByPid 根据指定的进程id返回其监听的端口号。
// 如果指定的pid没有找到对应的端口，则返回0。
func getPortByPid(pid int) int {
	path := getCommFilePath(pid)
	if path == "" {
		return 0
	}
	return gconv.Int(gfile.GetContentsWithCache(path))
}

// getCommFilePath 根据给定的pid返回其对应的端口映射文件路径。
func getCommFilePath(pid int) string {
	path, err := getCommPidFolderPath()
	if err != nil {
		intlog.Errorf(context.TODO(), `%+v`, err)
		return ""
	}
	return gfile.Join(path, gconv.String(pid))
}

// getCommPidFolderPath 获取并返回可用于存储pid映射文件的可用目录。
func getCommPidFolderPath() (folderPath string, err error) {
	commPidFolderPathOnce.Do(func() {
		availablePaths := []string{
			"/var/tmp",
			"/var/run",
		}
		if path, _ := gfile.Home(".config"); path != "" {
			availablePaths = append(availablePaths, path)
		}
		availablePaths = append(availablePaths, gfile.Temp())
		for _, availablePath := range availablePaths {
			checkPath := gfile.Join(availablePath, defaultFolderNameForProcComm)
			if !gfile.Exists(checkPath) && gfile.Mkdir(checkPath) != nil {
				continue
			}
			if gfile.IsWritable(checkPath) {
				commPidFolderPath = checkPath
				break
			}
		}
		if commPidFolderPath == "" {
			err = gerror.Newf(
				`cannot find available folder for storing pid to port mapping files in paths: %+v`,
				availablePaths,
			)
		}
	})
	folderPath = commPidFolderPath
	return
}
