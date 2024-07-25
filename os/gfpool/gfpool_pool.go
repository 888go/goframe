// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gfpool

import (
	"os"
	"time"

	"github.com/gogf/gf/v2/container/gpool"
	"github.com/gogf/gf/v2/container/gtype"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gfsnotify"
)

// New 创建并返回一个具有给定文件路径、标志和打开权限的文件指针池。
//
// 注意过期逻辑：
// ttl = 0：不过期；
// ttl < 0：使用后立即过期；
// ttl > 0：超时过期；
// 默认情况下，它不过期。 md5:521d6eb77a70a063
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

// newFilePool 创建并返回一个具有给定文件路径、标志和打开权限的文件指针池。 md5:62fda6ff96d41c0f
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

// File 从文件指针池中获取文件项并返回。如果文件指针池为空，它将创建一个。
// 注意，当文件不再会被使用时，应当关闭它。当它被“关闭”时，并不是真正关闭底层的文件指针，而是将其放回文件指针池中。 md5:b6c4b6a3ade746fc
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
					// 获取新创建的文件的状态。 md5:dbe21999357cbc52
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
		// 为了提高性能，它首先使用！p.init.Val()进行检查。 md5:bd8c9ebe349c994a
		if !p.init.Val() && p.init.Cas(false, true) {
			_, _ = gfsnotify.Add(f.path, func(event *gfsnotify.Event) {
				// 如果文件被删除或重命名，通过增加池ID来重新创建池。 md5:e825bec9648178de
				if event.IsRemove() || event.IsRename() {
					// It drops the old pool.
					p.id.Add(1)
					// 清除池中残留的项目。 md5:630859bb0da3cfb4
					p.pool.Clear()
					// 它使用另一个添加操作来移除两个添加操作之间的文件项。
					// 每当池ID改变时，池将被重新创建。 md5:d5f8fd9aa698b70a
					p.id.Add(1)
				}
			}, false)
		}
		return f, nil
	}
}

// Close关闭当前文件指针池。 md5:01a922bcbbea5a0f
func (p *Pool) Close() {
	p.pool.Close()
}
