// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类

import (
	"context"
	"io"
	"mime/multipart"
	"strconv"
	"strings"

	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/internal/json"
	gfile "github.com/888go/goframe/os/gfile"
	gtime "github.com/888go/goframe/os/gtime"
	grand "github.com/888go/goframe/util/grand"
)

// UploadFile 包装了multipart上传文件，提供了更多和更方便的功能。 md5:a7173285d087c4aa
type UploadFile struct {
	*multipart.FileHeader `json:"-"`
	ctx                   context.Context
}

// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。 md5:43c3b36e60a18f9a
func (f UploadFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.FileHeader)
}

// UploadFiles 是 *UploadFile 的数组类型。 md5:94b6aef81609f12b
type UploadFiles []*UploadFile

// X保存 将单个上传的文件保存到指定的目录路径，并返回保存的文件名。
//
// 参数 `dirPath` 应为一个目录路径，否则会返回错误。
//
// 注意：如果目标位置已经存在同名文件，该函数将覆盖原有的文件。
// md5:ffe3d8f90d14185a
func (f *UploadFile) X保存(目录路径 string, 随机重命名 ...bool) (文件名 string, 错误 error) {
	if f == nil {
		return "", gerror.X创建错误码(
			gcode.CodeMissingParameter,
			"file is empty, maybe you retrieve it from invalid field name or form enctype",
		)
	}
	if !gfile.X是否存在(目录路径) {
		if 错误 = gfile.X创建目录(目录路径); 错误 != nil {
			return
		}
	} else if !gfile.X是否存在目录(目录路径) {
		return "", gerror.X创建错误码(gcode.CodeInvalidParameter, `parameter "dirPath" should be a directory path`)
	}

	file, 错误 := f.Open()
	if 错误 != nil {
		错误 = gerror.X多层错误并格式化(错误, `UploadFile.Open failed`)
		return "", 错误
	}
	defer file.Close()

	name := gfile.X路径取文件名(f.Filename)
	if len(随机重命名) > 0 && 随机重命名[0] {
		name = strings.ToLower(strconv.FormatInt(gtime.X取时间戳纳秒(), 36) + grand.X文本(6))
		name = name + gfile.X路径取扩展名(f.Filename)
	}
	filePath := gfile.X路径生成(目录路径, name)
	newFile, 错误 := gfile.X创建文件与目录(filePath)
	if 错误 != nil {
		return "", 错误
	}
	defer newFile.Close()
	intlog.Printf(f.ctx, `save upload file: %s`, filePath)
	if _, 错误 = io.Copy(newFile, file); 错误 != nil {
		错误 = gerror.X多层错误并格式化(错误, `io.Copy failed from "%s" to "%s"`, f.Filename, filePath)
		return "", 错误
	}
	return gfile.X路径取文件名(filePath), nil
}

// X保存 将所有上传的文件保存到指定的目录路径，并返回保存的文件名。
//
// 参数 `dirPath` 应该是一个目录路径，否则会返回错误。
//
// 参数 `randomlyRename` 指定是否为所有文件名随机重命名。
// md5:de2b45ea5a89ccad
func (fs UploadFiles) X保存(目录路径 string, 随机重命名 ...bool) (文件名切片 []string, 错误 error) {
	if len(fs) == 0 {
		return nil, gerror.X创建错误码(
			gcode.CodeMissingParameter,
			"file array is empty, maybe you retrieve it from invalid field name or form enctype",
		)
	}
	for _, f := range fs {
		if filename, err := f.X保存(目录路径, 随机重命名...); err != nil {
			return 文件名切片, err
		} else {
			文件名切片 = append(文件名切片, filename)
		}
	}
	return
}

// X取上传文件对象 通过指定的表单名称检索并返回上传中的文件。
// 此函数用于检索使用multipart/form-data内容类型上传的单个文件对象。
//
// 如果检索失败或没有给定名称的表单文件被上传，它将返回nil。
//
// 注意，`name` 是客户端multipart表单中文件字段的名称。
// md5:a49268bd7e014ab6
func (r *Request) X取上传文件对象(名称 string) *UploadFile {
	uploadFiles := r.X取上传文件切片对象(名称)
	if len(uploadFiles) > 0 {
		return uploadFiles[0]
	}
	return nil
}

// X取上传文件切片对象 用于检索并返回具有指定表单名称的多个上传文件。
// 此函数用于获取多个上传文件对象，这些对象是使用多部分表单内容类型上传的。
//
// 如果检索失败或没有给定名称的表单文件被上传，则返回nil。
//
// 注意，`name` 是来自客户端的多部分表单中的文件字段名称。
// md5:cbbf4db398137505
func (r *Request) X取上传文件切片对象(名称 string) UploadFiles {
	multipartFiles := r.X取multipart表单文件切片对象(名称)
	if len(multipartFiles) > 0 {
		uploadFiles := make(UploadFiles, len(multipartFiles))
		for k, v := range multipartFiles {
			uploadFiles[k] = &UploadFile{
				ctx:        r.Context别名(),
				FileHeader: v,
			}
		}
		return uploadFiles
	}
	return nil
}
