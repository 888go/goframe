// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gproc

import (
	"context"
	"fmt"
	"net"
	
	"github.com/888go/goframe/container/gqueue"
	"github.com/888go/goframe/container/gtype"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/net/gtcp"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/glog"
	"github.com/888go/goframe/util/gconv"
)

var (
	// tcpListened 标记接收端监听服务是否已启动。
	tcpListened = gtype.NewBool()
)

// 接收区块并通过本地TCP监听从其他进程接收消息。
// 注意，只有当调用此函数时，才会启用TCP监听服务。
func Receive(group ...string) *MsgRequest {
	// 使用原子操作以确保只有一个接收goroutine在监听。
	if tcpListened.Cas(false, true) {
		go receiveTcpListening()
	}
	var groupName string
	if len(group) > 0 {
		groupName = group[0]
	} else {
		groupName = defaultGroupNameForProcComm
	}
	queue := commReceiveQueues.GetOrSetFuncLock(groupName, func() interface{} {
		return gqueue.New(maxLengthForProcMsgQueue)
	}).(*gqueue.Queue)

	// 阻塞接收。
	if v := queue.Pop(); v != nil {
		return v.(*MsgRequest)
	}
	return nil
}

// receiveTcpListening 在本地扫描可用端口并开始监听。
func receiveTcpListening() {
	var (
		listen  *net.TCPListener
		conn    net.Conn
		port    = gtcp.MustGetFreePort()
		address = fmt.Sprintf("127.0.0.1:%d", port)
	)
	tcpAddress, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		panic(gerror.Wrap(err, `net.ResolveTCPAddr failed`))
	}
	listen, err = net.ListenTCP("tcp", tcpAddress)
	if err != nil {
		panic(gerror.Wrapf(err, `net.ListenTCP failed for address "%s"`, address))
	}
	// 将端口保存到pid文件中。
	if err = gfile.PutContents(getCommFilePath(Pid()), gconv.String(port)); err != nil {
		panic(err)
	}
	// Start listening.
	for {
		if conn, err = listen.Accept(); err != nil {
			glog.Error(context.TODO(), err)
		} else if conn != nil {
			go receiveTcpHandler(gtcp.NewConnByNetConn(conn))
		}
	}
}

// receiveTcpHandler 是用于接收数据的连接处理器。
func receiveTcpHandler(conn *gtcp.Conn) {
	var (
		ctx      = context.TODO()
		result   []byte
		response MsgResponse
	)
	for {
		response.Code = 0
		response.Message = ""
		response.Data = nil
		buffer, err := conn.RecvPkg()
		if len(buffer) > 0 {
			// Package decoding.
			msg := new(MsgRequest)
			if err = json.UnmarshalUseNumber(buffer, msg); err != nil {
				continue
			}
			if msg.ReceiverPid != Pid() {
				// Not mine package.
				response.Message = fmt.Sprintf(
					"receiver pid not match, target: %d, current: %d",
					msg.ReceiverPid, Pid(),
				)
			} else if v := commReceiveQueues.Get(msg.Group); v == nil {
				// Group check.
				response.Message = fmt.Sprintf("group [%s] does not exist", msg.Group)
			} else {
				// 将元素推送到缓冲队列中。
				response.Code = 1
				v.(*gqueue.Queue).Push(msg)
			}
		} else {
			// Empty package.
			response.Message = "empty package"
		}
		if err == nil {
			result, err = json.Marshal(response)
			if err != nil {
				glog.Error(ctx, err)
			}
			if err = conn.SendPkg(result); err != nil {
				glog.Error(ctx, err)
			}
		} else {
			// 如果发生任何错误，仅关闭连接即可。
			if err = conn.Close(); err != nil {
				glog.Error(ctx, err)
			}
			break
		}
	}
}
