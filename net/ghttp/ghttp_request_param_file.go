// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp
import (
	"context"
	"io"
	"mime/multipart"
	"strconv"
	"strings"
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/util/grand"
	)
// UploadFile 通过提供更多的便捷功能，对使用multipart方式上传文件进行了封装。
type UploadFile struct {
	*multipart.FileHeader `json:"-"`
	ctx                   context.Context
}

// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
func (f UploadFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.FileHeader)
}

// UploadFiles 是 *UploadFile 类型的数组。
type UploadFiles []*UploadFile

// Save 保存单个上传的文件到指定目录路径，并返回已保存的文件名。
//
// 参数`dirPath`应为一个目录路径，否则将返回错误。
//
// 注意：如果目标位置已经存在同名文件，则会直接覆盖该文件。
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

// Save 保存所有上传的文件到指定的目录路径，并返回已保存的文件名。
//
// 参数`dirPath`应为一个目录路径，否则将返回错误。
//
// 参数`randomlyRename`指定了是否随机重命名所有的文件名。
// 以下是详细的中文注释：
// ```go
// Save 函数的作用是将所有正在上传的文件保存到指定的目录路径下，并返回这些文件在保存后的文件名列表。
//
// 参数 `dirPath` 表示目标保存目录的路径，如果该路径不是一个有效的目录，则函数会返回错误信息。
//
// 参数 `randomlyRename` 是一个布尔值，用来指定是否对所有上传的文件进行随机重命名操作。如果设为 true，则在保存文件时将会生成随机文件名；否则，文件将以原有文件名进行保存。
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

// GetUploadFile 通过指定表单名称获取并返回上传文件。
// 此函数用于检索单个上传文件对象，该对象是通过多部分表单内容类型上传的。
//
// 如果检索失败或没有提交给定名称的表单文件，则返回 nil。
//
// 注意，`name` 是客户端多部分表单中文件字段的名称。
func (r *Request) GetUploadFile(name string) *UploadFile {
	uploadFiles := r.GetUploadFiles(name)
	if len(uploadFiles) > 0 {
		return uploadFiles[0]
	}
	return nil
}

// GetUploadFiles 根据指定表单名称获取并返回多个上传文件。
// 此函数用于检索多个使用 multipart/form-data 类型上传的文件对象。
//
// 如果检索失败或没有接收到客户端通过给定名称上传的表单文件，将返回 nil。
//
// 注意，`name` 是客户端 multipart 表单中文件字段的名称。
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
