// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gfpool
import (
	"os"
	"time"
	
	"github.com/888go/goframe/container/gpool"
	"github.com/888go/goframe/container/gtype"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/os/gfsnotify"
	)
// New根据给定的文件路径、标志和打开权限创建并返回一个文件指针池。
//
// 注意过期逻辑：
// ttl = 0 : 不过期；
// ttl < 0 : 使用后立即过期；
// ttl > 0 : 超时后过期；
// 默认情况下，它不会过期。
func New(path string, flag int, perm os.FileMode, ttl ...time.Duration) *Pool {
	var fpTTL time.Duration
	if len(ttl) > 0 {
		fpTTL = ttl[0]
	}
	p := &Pool{
		id:   gtype.NewInt(),
		ttl:  fpTTL,
		init: gtype.NewBool(),
	}
	p.pool = newFilePool(p, path, flag, perm, fpTTL)
	return p
}

// newFilePool 根据给定的文件路径、标志和打开权限创建并返回一个文件指针池。
func newFilePool(p *Pool, path string, flag int, perm os.FileMode, ttl time.Duration) *gpool.Pool {
	pool := gpool.New(ttl, func() (interface{}, error) {
		file, err := os.OpenFile(path, flag, perm)
		if err != nil {
			err = gerror.Wrapf(err, `os.OpenFile failed for file "%s", flag "%d", perm "%s"`, path, flag, perm)
			return nil, err
		}
		return &File{
			File: file,
			pid:  p.id.Val(),
			pool: p,
			flag: flag,
			perm: perm,
			path: path,
		}, nil
	}, func(i interface{}) {
		_ = i.(*File).File.Close()
	})
	return pool
}

// File 从文件指针池中获取文件项并返回，如果文件指针池为空，则创建一个新的文件项。
// 注意：当文件项不再使用时，应关闭它。当其被关闭时，并非真正关闭底层的文件指针，而是将其放回文件指针池中。
func (p *Pool) File() (*File, error) {
	if v, err := p.pool.Get(); err != nil {
		return nil, err
	} else {
		f := v.(*File)
		f.stat, err = os.Stat(f.path)
		if f.flag&os.O_CREATE > 0 {
			if os.IsNotExist(err) {
				if f.File, err = os.OpenFile(f.path, f.flag, f.perm); err != nil {
					return nil, err
				} else {
					// 获取新创建文件的状态。
					if f.stat, err = f.File.Stat(); err != nil {
						return nil, err
					}
				}
			}
		}
		if f.flag&os.O_TRUNC > 0 {
			if f.stat.Size() > 0 {
				if err = f.Truncate(0); err != nil {
					return nil, err
				}
			}
		}
		if f.flag&os.O_APPEND > 0 {
			if _, err = f.Seek(0, 2); err != nil {
				return nil, err
			}
		} else {
			if _, err = f.Seek(0, 0); err != nil {
				return nil, err
			}
		}
		// 为了性能优化，首先使用 !p.init.Val() 进行检查。
		if !p.init.Val() && p.init.Cas(false, true) {
			_, _ = gfsnotify.Add(f.path, func(event *gfsnotify.Event) {
				// 如果文件被删除或重命名，通过增加pool id重新创建pool。
				if event.IsRemove() || event.IsRename() {
					// 它会丢弃旧的连接池。
					p.id.Add(1)
					// 清除池中留存的池项。
					p.pool.Clear()
// 它利用另一个添加操作来丢弃两个添加之间的文件项。
// 每当池ID发生变化时，将会重新创建该池。
					p.id.Add(1)
				}
			}, false)
		}
		return f, nil
	}
}

// Close 关闭当前文件指针池。
func (p *Pool) Close() {
	p.pool.Close()
}
