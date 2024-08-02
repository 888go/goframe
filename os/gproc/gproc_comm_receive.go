// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 进程类

import (
	"context"
	"fmt"
	"net"

	gqueue "github.com/888go/goframe/container/gqueue"
	gtype "github.com/888go/goframe/container/gtype"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/json"
	gtcp "github.com/888go/goframe/net/gtcp"
	gfile "github.com/888go/goframe/os/gfile"
	glog "github.com/888go/goframe/os/glog"
	gconv "github.com/888go/goframe/util/gconv"
)

var (
		// tcpListened 标记接收监听服务是否已启动。 md5:3e322f46ef5e3873
	tcpListened = gtype.NewBool()
)

// Receive 函数通过本地TCP监听来阻塞并接收来自其他进程的消息。
// 注意，只有当调用此函数时，才会启用TCP监听服务。
// md5:dbf5481b7dcc4222
func Receive(group ...string) *MsgRequest {
		// 使用原子操作来保证只有一个接收者goroutine在监听。 md5:3ddf24c1b343c721
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

	// Blocking receiving.
	if v := queue.Pop(); v != nil {
		return v.(*MsgRequest)
	}
	return nil
}

// receiveTcpListening 在本地扫描可用端口并开始监听。 md5:797539e564e3129a
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
		// 将端口保存到pid文件中。 md5:d6040683bfc292d5
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

// receiveTcpHandler 是用于接收数据的连接处理器。 md5:41801f5c109ca561
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
				// Push to buffer queue.
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
						// 如果发生任何错误，只需关闭连接。 md5:23e98bee4057f221
			if err = conn.Close(); err != nil {
				glog.Error(ctx, err)
			}
			break
		}
	}
}
