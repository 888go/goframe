// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gsession
import (
	"context"
	"fmt"
	"os"
	"time"
	
	"github.com/888go/goframe/container/gmap"
	"github.com/888go/goframe/container/gset"
	"github.com/888go/goframe/crypto/gaes"
	"github.com/888go/goframe/encoding/gbinary"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/os/gtimer"
	)
// StorageFile实现了使用文件系统作为Session存储接口。
type StorageFile struct {
	StorageBase
	path          string        // 会话文件存储文件夹路径。
	ttl           time.Duration // Session TTL.
	cryptoKey     []byte        // 当启用加密功能时使用。
	cryptoEnabled bool          // 当启用加密功能时使用。
	updatingIdSet *gset.StrSet  // 待批量更新的会话ID集合。
}

const (
	DefaultStorageFileCryptoEnabled        = false
	DefaultStorageFileUpdateTTLInterval    = 10 * time.Second
	DefaultStorageFileClearExpiredInterval = time.Hour
)

var (
	DefaultStorageFilePath      = gfile.Temp("gsessions")
	DefaultStorageFileCryptoKey = []byte("Session storage file crypto key!")
)

// NewStorageFile 创建并返回一个用于存储session的文件存储对象。
func NewStorageFile(path string, ttl time.Duration) *StorageFile {
	var (
		ctx         = context.TODO()
		storagePath = DefaultStorageFilePath
	)
	if path != "" {
		storagePath, _ = gfile.Search(path)
		if storagePath == "" {
			panic(gerror.NewCodef(gcode.CodeInvalidParameter, `"%s" does not exist`, path))
		}
		if !gfile.IsWritable(storagePath) {
			panic(gerror.NewCodef(gcode.CodeInvalidParameter, `"%s" is not writable`, path))
		}
	}
	if storagePath != "" {
		if err := gfile.Mkdir(storagePath); err != nil {
			panic(gerror.Wrapf(err, `Mkdir "%s" failed in PWD "%s"`, path, gfile.Pwd()))
		}
	}
	s := &StorageFile{
		path:          storagePath,
		ttl:           ttl,
		cryptoKey:     DefaultStorageFileCryptoKey,
		cryptoEnabled: DefaultStorageFileCryptoEnabled,
		updatingIdSet: gset.NewStrSet(true),
	}

	gtimer.AddSingleton(ctx, DefaultStorageFileUpdateTTLInterval, s.timelyUpdateSessionTTL)
	gtimer.AddSingleton(ctx, DefaultStorageFileClearExpiredInterval, s.timelyClearExpiredSessionFile)
	return s
}

// timelyUpdateSessionTTL 批量及时更新会话的TTL（生存时间）
func (s *StorageFile) timelyUpdateSessionTTL(ctx context.Context) {
	var (
		sessionId string
		err       error
	)
	// 批量更新会话。
	for {
		if sessionId = s.updatingIdSet.Pop(); sessionId == "" {
			break
		}
		if err = s.updateSessionTTl(context.TODO(), sessionId); err != nil {
			intlog.Errorf(context.TODO(), `%+v`, err)
		}
	}
}

// 定时清理过期会话文件，及时删除所有已过期的文件。
func (s *StorageFile) timelyClearExpiredSessionFile(ctx context.Context) {
	files, err := gfile.ScanDirFile(s.path, "*.session", false)
	if err != nil {
		intlog.Errorf(ctx, `%+v`, err)
		return
	}
	for _, file := range files {
		if err = s.checkAndClearSessionFile(ctx, file); err != nil {
			intlog.Errorf(ctx, `%+v`, err)
		}
	}
}

// SetCryptoKey 设置会话存储的加密密钥。
// 当启用加密功能时，会使用此加密密钥。
func (s *StorageFile) SetCryptoKey(key []byte) {
	s.cryptoKey = key
}

// SetCryptoEnabled 用于启用/禁用会话存储的加密功能。
func (s *StorageFile) SetCryptoEnabled(enabled bool) {
	s.cryptoEnabled = enabled
}

// sessionFilePath根据给定的session id返回存储文件路径。
func (s *StorageFile) sessionFilePath(sessionId string) string {
	return gfile.Join(s.path, sessionId) + ".session"
}

// RemoveAll 从存储中删除所有键值对。
func (s *StorageFile) RemoveAll(ctx context.Context, sessionId string) error {
	return gfile.Remove(s.sessionFilePath(sessionId))
}

// GetSession 通过给定的 session id 从存储中获取 session 数据，并以 *gmap.StrAnyMap 类型返回。
//
// 参数 `ttl` 指定了该 session 的生存时间（TTL），若生存时间已过，则返回 nil。
// 参数 `data` 是当前存储在内存中的旧 session 数据，如果禁用了内存存储，对于某些存储方式，此参数可能为 nil。
//
// 当每次 session 开始时，都会调用这个函数。
func (s *StorageFile) GetSession(ctx context.Context, sessionId string, ttl time.Duration) (sessionData *gmap.StrAnyMap, err error) {
	var (
		path    = s.sessionFilePath(sessionId)
		content = gfile.GetBytes(path)
	)
	// 如果会话文件已经存在，则仅更新TTL（生存时间）
	if len(content) > 8 {
		timestampMilli := gbinary.DecodeToInt64(content[:8])
		if timestampMilli+ttl.Nanoseconds()/1e6 < gtime.TimestampMilli() {
			return nil, nil
		}
		content = content[8:]
		// Decrypt with AES.
		if s.cryptoEnabled {
			content, err = gaes.Decrypt(content, DefaultStorageFileCryptoKey)
			if err != nil {
				return nil, err
			}
		}
		var m map[string]interface{}
		if err = json.UnmarshalUseNumber(content, &m); err != nil {
			return nil, err
		}
		if m == nil {
			return nil, nil
		}
		return gmap.NewStrAnyMapFrom(m, true), nil
	}
	return nil, nil
}

// SetSession 更新指定会话 ID 的数据映射。
// 在每次已标记为脏的、发生改变的会话关闭后，都会调用此函数。
// 此函数将内存中的所有会话数据映射复制到存储中。
func (s *StorageFile) SetSession(ctx context.Context, sessionId string, sessionData *gmap.StrAnyMap, ttl time.Duration) error {
	intlog.Printf(ctx, "StorageFile.SetSession: %s, %v, %v", sessionId, sessionData, ttl)
	path := s.sessionFilePath(sessionId)
	content, err := json.Marshal(sessionData)
	if err != nil {
		return err
	}
	// Encrypt with AES.
	if s.cryptoEnabled {
		content, err = gaes.Encrypt(content, DefaultStorageFileCryptoKey)
		if err != nil {
			return err
		}
	}
	file, err := gfile.OpenWithFlagPerm(
		path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm,
	)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err = file.Write(gbinary.EncodeInt64(gtime.TimestampMilli())); err != nil {
		err = gerror.Wrapf(err, `write data failed to file "%s"`, path)
		return err
	}
	if _, err = file.Write(content); err != nil {
		err = gerror.Wrapf(err, `write data failed to file "%s"`, path)
		return err
	}
	return nil
}

// UpdateTTL 更新指定会话ID的TTL（生存时间）。
// 此函数在非脏数据会话关闭后调用。
// 它只是将该会话ID添加到异步处理队列中。
func (s *StorageFile) UpdateTTL(ctx context.Context, sessionId string, ttl time.Duration) error {
	intlog.Printf(ctx, "StorageFile.UpdateTTL: %s, %v", sessionId, ttl)
	if ttl >= DefaultStorageFileUpdateTTLInterval {
		s.updatingIdSet.Add(sessionId)
	}
	return nil
}

// updateSessionTTL 更新指定会话ID的TTL（生存时间）
func (s *StorageFile) updateSessionTTl(ctx context.Context, sessionId string) error {
	intlog.Printf(ctx, "StorageFile.updateSession: %s", sessionId)
	path := s.sessionFilePath(sessionId)
	file, err := gfile.OpenWithFlag(path, os.O_WRONLY)
	if err != nil {
		return err
	}
	if _, err = file.WriteAt(gbinary.EncodeInt64(gtime.TimestampMilli()), 0); err != nil {
		err = gerror.Wrapf(err, `write data failed to file "%s"`, path)
		return err
	}
	return file.Close()
}

func (s *StorageFile) checkAndClearSessionFile(ctx context.Context, path string) (err error) {
	var (
		file                *os.File
		readBytesCount      int
		timestampMilliBytes = make([]byte, 8)
	)
	file, err = gfile.OpenWithFlag(path, os.O_RDONLY)
	if err != nil {
		return err
	}
	defer file.Close()
	// 读取会话文件更新的毫秒级时间戳。
	readBytesCount, err = file.Read(timestampMilliBytes)
	if err != nil {
		return
	}
	if readBytesCount != 8 {
		return gerror.Newf(`invalid read bytes count "%d", expect "8"`, readBytesCount)
	}
	// 移除过期的会话文件。
	var (
		ttlInMilliseconds     = s.ttl.Nanoseconds() / 1e6
		fileTimestampMilli    = gbinary.DecodeToInt64(timestampMilliBytes)
		currentTimestampMilli = gtime.TimestampMilli()
	)
	if fileTimestampMilli+ttlInMilliseconds < currentTimestampMilli {
		intlog.PrintFunc(ctx, func() string {
			return fmt.Sprintf(
				`clear expired session file "%s": updated datetime "%s", ttl "%s"`,
				path, gtime.NewFromTimeStamp(fileTimestampMilli), s.ttl,
			)
		})
		return gfile.Remove(path)
	}
	return nil
}
