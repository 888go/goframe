// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gsession

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/crypto/gaes"
	"github.com/gogf/gf/v2/encoding/gbinary"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/intlog"
	"github.com/gogf/gf/v2/internal/json"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/os/gtimer"
)

// StorageFile实现了使用文件系统作为会话存储的接口。 md5:bae13bc406aa3178
type StorageFile struct {
	StorageBase
	path          string        // 会话文件存储文件夹路径。 md5:a07352e7e4b2ee5d
	ttl           time.Duration // Session TTL.
	cryptoKey     []byte        // 在启用加密功能时使用。 md5:e2b00bed77b9f059
	cryptoEnabled bool          // 在启用加密功能时使用。 md5:e2b00bed77b9f059
	updatingIdSet *gset.StrSet  // 用于批量更新的会话ID集合。 md5:704d1ffa9a08b42d
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

// NewStorageFile 创建并返回一个用于会话的文件存储对象。 md5:047619bd552117d1
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

// timelyUpdateSessionTTL 批量及时更新会话的超时时间。 md5:8d440f6681b47013
func (s *StorageFile) timelyUpdateSessionTTL(ctx context.Context) {
	var (
		sessionId string
		err       error
	)
		// 批量更新会话。 md5:db1f90067d27cc66
	for {
		if sessionId = s.updatingIdSet.Pop(); sessionId == "" {
			break
		}
		if err = s.updateSessionTTl(context.TODO(), sessionId); err != nil {
			intlog.Errorf(context.TODO(), `%+v`, err)
		}
	}
}

// timelyClearExpiredSessionFile 及时删除所有过期的文件。 md5:5f02dbf03c17d4ca
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
// 当启用加密功能时，将使用此加密密钥。
// md5:dbc53d710307bd28
func (s *StorageFile) SetCryptoKey(key []byte) {
	s.cryptoKey = key
}

// SetCryptoEnabled 启用/禁用会话存储的加密功能。 md5:14228b4577da32ec
func (s *StorageFile) SetCryptoEnabled(enabled bool) {
	s.cryptoEnabled = enabled
}

// sessionFilePath 根据给定的会话ID返回存储文件的路径。 md5:9cec805dff8d12a7
func (s *StorageFile) sessionFilePath(sessionId string) string {
	return gfile.Join(s.path, sessionId) + ".session"
}

// RemoveAll 删除存储中的所有键值对。 md5:8b06607595d19a73
func (s *StorageFile) RemoveAll(ctx context.Context, sessionId string) error {
	return gfile.Remove(s.sessionFilePath(sessionId))
}

// GetSession 从存储中根据给定的会话ID获取会话数据，返回一个指向*gmap.StrAnyMap的指针。
//
// 参数`ttl`指定了此会话的有效期，如果超过有效期，则返回nil。参数`data`是当前存储在内存中的旧会话数据，对于某些存储方式，如果禁用了内存存储，它可能会为nil。
//
// 此函数在会话启动时会被调用。
// md5:01e56ce09d5fd934
func (s *StorageFile) GetSession(ctx context.Context, sessionId string, ttl time.Duration) (sessionData *gmap.StrAnyMap, err error) {
	var (
		path    = s.sessionFilePath(sessionId)
		content = gfile.GetBytes(path)
	)
		// 只有当会话文件已经存在时，它才会更新TTL。 md5:a9223056bbc67ae2
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

// SetSession 根据指定的会话ID更新数据映射。
// 当某个被标记为脏（即发生过修改）的会话关闭后，将调用此函数。
// 该操作会将所有会话数据从内存复制到存储中。
// md5:1caa26989d884fa4
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

// UpdateTTL 更新指定会话ID的生存时间（TTL）。
// 当一个未被修改（非脏）的会话关闭后，此函数会被调用。
// 它只是将会话ID添加到异步处理队列中。
// md5:cc5ac287cbbc0eab
func (s *StorageFile) UpdateTTL(ctx context.Context, sessionId string, ttl time.Duration) error {
	intlog.Printf(ctx, "StorageFile.UpdateTTL: %s, %v", sessionId, ttl)
	if ttl >= DefaultStorageFileUpdateTTLInterval {
		s.updatingIdSet.Add(sessionId)
	}
	return nil
}

// updateSessionTTL 更新指定会话ID的超时时间。 md5:1cff3164e4ca8226
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
		// 读取会话文件更新的时间戳（以毫秒为单位）。 md5:e3b93f5cb9dd863b
	readBytesCount, err = file.Read(timestampMilliBytes)
	if err != nil {
		return
	}
	if readBytesCount != 8 {
		return gerror.Newf(`invalid read bytes count "%d", expect "8"`, readBytesCount)
	}
		// 删除过期的会话文件。 md5:f3e7a080ff4d0135
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
