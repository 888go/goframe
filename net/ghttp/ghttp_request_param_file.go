// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package ghttp

import (
	"context"
	"io"
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/intlog"
	"github.com/gogf/gf/v2/internal/json"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
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

// Save 将单个上传的文件保存到指定的目录路径，并返回保存的文件名。
//
// 参数 `dirPath` 应为一个目录路径，否则会返回错误。
//
// 注意：如果目标位置已经存在同名文件，该函数将覆盖原有的文件。 md5:ffe3d8f90d14185a
func (f *UploadFile) Save(dirPath string, randomlyRename ...bool) (filename string, err error) {
	if f == nil {
		return "", gerror.NewCode(
			gcode.CodeMissingParameter,
			"file is empty, maybe you retrieve it from invalid field name or form enctype",
		)
	}
	if !gfile.Exists(dirPath) {
		if err = gfile.Mkdir(dirPath); err != nil {
			return
		}
	} else if !gfile.IsDir(dirPath) {
		return "", gerror.NewCode(gcode.CodeInvalidParameter, `parameter "dirPath" should be a directory path`)
	}

	file, err := f.Open()
	if err != nil {
		err = gerror.Wrapf(err, `UploadFile.Open failed`)
		return "", err
	}
	defer file.Close()

	name := gfile.Basename(f.Filename)
	if len(randomlyRename) > 0 && randomlyRename[0] {
		name = strings.ToLower(strconv.FormatInt(gtime.TimestampNano(), 36) + grand.S(6))
		name = name + gfile.Ext(f.Filename)
	}
	filePath := gfile.Join(dirPath, name)
	newFile, err := gfile.Create(filePath)
	if err != nil {
		return "", err
	}
	defer newFile.Close()
	intlog.Printf(f.ctx, `save upload file: %s`, filePath)
	if _, err = io.Copy(newFile, file); err != nil {
		err = gerror.Wrapf(err, `io.Copy failed from "%s" to "%s"`, f.Filename, filePath)
		return "", err
	}
	return gfile.Basename(filePath), nil
}

// Save 将所有上传的文件保存到指定的目录路径，并返回保存的文件名。
//
// 参数 `dirPath` 应该是一个目录路径，否则会返回错误。
//
// 参数 `randomlyRename` 指定是否为所有文件名随机重命名。 md5:de2b45ea5a89ccad
func (fs UploadFiles) Save(dirPath string, randomlyRename ...bool) (filenames []string, err error) {
	if len(fs) == 0 {
		return nil, gerror.NewCode(
			gcode.CodeMissingParameter,
			"file array is empty, maybe you retrieve it from invalid field name or form enctype",
		)
	}
	for _, f := range fs {
		if filename, err := f.Save(dirPath, randomlyRename...); err != nil {
			return filenames, err
		} else {
			filenames = append(filenames, filename)
		}
	}
	return
}

// GetUploadFile 通过指定的表单名称检索并返回上传中的文件。
// 此函数用于检索使用multipart/form-data内容类型上传的单个文件对象。
//
// 如果检索失败或没有给定名称的表单文件被上传，它将返回nil。
//
// 注意，`name` 是客户端multipart表单中文件字段的名称。 md5:a49268bd7e014ab6
func (r *Request) GetUploadFile(name string) *UploadFile {
	uploadFiles := r.GetUploadFiles(name)
	if len(uploadFiles) > 0 {
		return uploadFiles[0]
	}
	return nil
}

// GetUploadFiles 用于检索并返回具有指定表单名称的多个上传文件。
// 此函数用于获取多个上传文件对象，这些对象是使用多部分表单内容类型上传的。
//
// 如果检索失败或没有给定名称的表单文件被上传，则返回nil。
//
// 注意，`name` 是来自客户端的多部分表单中的文件字段名称。 md5:cbbf4db398137505
func (r *Request) GetUploadFiles(name string) UploadFiles {
	multipartFiles := r.GetMultipartFiles(name)
	if len(multipartFiles) > 0 {
		uploadFiles := make(UploadFiles, len(multipartFiles))
		for k, v := range multipartFiles {
			uploadFiles[k] = &UploadFile{
				ctx:        r.Context(),
				FileHeader: v,
			}
		}
		return uploadFiles
	}
	return nil
}
