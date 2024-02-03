// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gproc

import (
	"io"
	
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/net/gtcp"
)

// Send 将数据发送给指定的进程ID（pid）所对应的进程。
func Send(pid int, data []byte, group ...string) error {
	msg := MsgRequest{
		SenderPid:   Pid(),
		ReceiverPid: pid,
		Group:       defaultGroupNameForProcComm,
		Data:        data,
	}
	if len(group) > 0 {
		msg.Group = group[0]
	}
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	var conn *gtcp.PoolConn
	conn, err = getConnByPid(pid)
	if err != nil {
		return err
	}
	defer conn.Close()
	// Do the sending.
	var result []byte
	result, err = conn.SendRecvPkg(msgBytes, gtcp.PkgOption{
		Retry: gtcp.Retry{
			Count: 3,
		},
	})
	if len(result) > 0 {
		response := new(MsgResponse)
		if err = json.UnmarshalUseNumber(result, response); err == nil {
			if response.Code != 1 {
				err = gerror.New(response.Message)
			}
		}
	}
	// EOF 并不是一个真正的错误。
	if err == io.EOF {
		err = nil
	}
	return err
}
