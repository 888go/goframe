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
	"sync"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/intlog"
	"github.com/gogf/gf/v2/net/gtcp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"
)

// MsgRequest是进程间通信的请求结构体。 md5:aa294ed7aef773f3
type MsgRequest struct {
	SenderPid   int    // Sender PID.
	ReceiverPid int    // Receiver PID.
	Group       string // Message group name.
	Data        []byte // Request data.
}

// MsgResponse 是进程通信中的响应结构体。 md5:a2e9e35f8a32b58e
type MsgResponse struct {
	Code    int    // 1: OK; Other: Error.
	Message string // Response message.
	Data    []byte // Response data.
}

const (
	defaultFolderNameForProcComm = "gf_pid_port_mapping" // 默认的保存pid到端口映射文件的文件夹名称。 md5:64d7a3cc62fc8b3c
	defaultGroupNameForProcComm  = ""                    // Default group name.
	defaultTcpPortForProcComm    = 10000                 // 用于接收者监听的起始端口号。 md5:57cde4f483b095cf
	maxLengthForProcMsgQueue     = 10000                 // 集群中每个消息队列的最大大小。 md5:64e3f3ac37111858
)

var (
	// commReceiveQueues 是一个用于存储接收到的数据的组名到队列的映射。
	// 该映射的值类型为*gqueue.Queue。
	// md5:adb11ba95544ea8c
	commReceiveQueues = gmap.NewStrAnyMap(true)

		// commPidFolderPath 指定了存储 PID 到端口映射文件的文件夹路径。 md5:bc9b0e25bfe8ea53
	commPidFolderPath string

		// commPidFolderPathOnce用于惰性计算`commPidFolderPath`，只有在必要时才进行。 md5:669e811a3607b61c
	commPidFolderPathOnce sync.Once
)

// getConnByPid 为指定的 pid 创建并返回一个 TCP 连接。 md5:19b60bfdf9f18aa2
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

// getPortByPid 根据指定的进程ID返回其监听的端口。
// 如果没有为指定的进程ID找到端口，则返回0。
// md5:1fc2deacfe985ab1
func getPortByPid(pid int) int {
	path := getCommFilePath(pid)
	if path == "" {
		return 0
	}
	return gconv.Int(gfile.GetContentsWithCache(path))
}

// getCommFilePath 返回给定pid的进程到端口映射文件路径。 md5:6b8e5776476edbb5
func getCommFilePath(pid int) string {
	path, err := getCommPidFolderPath()
	if err != nil {
		intlog.Errorf(context.TODO(), `%+v`, err)
		return ""
	}
	return gfile.Join(path, gconv.String(pid))
}

// getCommPidFolderPath 获取并返回用于存储进程映射文件的可用目录。 md5:d871e38ee1ac7054
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
