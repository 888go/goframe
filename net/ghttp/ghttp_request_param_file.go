// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类

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
type X上传文件 struct {
	*multipart.FileHeader `json:"-"`
	ctx                   context.Context
}

// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
func (f X上传文件) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.FileHeader)
}

// UploadFiles 是 *UploadFile 类型的数组。
type X上传文件数组 []*X上传文件

// Save 保存单个上传的文件到指定目录路径，并返回已保存的文件名。
//
// 参数`dirPath`应为一个目录路径，否则将返回错误。
//
// 注意：如果目标位置已经存在同名文件，则会直接覆盖该文件。
func (f *X上传文件) X保存(目录路径 string, 随机重命名 ...bool) (文件名 string, 错误 error) {
	if f == nil {
		return "", 错误类.X创建错误码(
			错误码类.CodeMissingParameter,
			"file is empty, maybe you retrieve it from invalid field name or form enctype",
		)
	}
	if !文件类.X是否存在(目录路径) {
		if 错误 = 文件类.X创建目录(目录路径); 错误 != nil {
			return
		}
	} else if !文件类.X是否存在目录(目录路径) {
		return "", 错误类.X创建错误码(错误码类.CodeInvalidParameter, `parameter "dirPath" should be a directory path`)
	}

	file, 错误 := f.Open()
	if 错误 != nil {
		错误 = 错误类.X多层错误并格式化(错误, `UploadFile.Open failed`)
		return "", 错误
	}
	defer file.Close()

	name := 文件类.X路径取文件名(f.Filename)
	if len(随机重命名) > 0 && 随机重命名[0] {
		name = strings.ToLower(strconv.FormatInt(时间类.X取时间戳纳秒(), 36) + 随机类.X文本(6))
		name = name + 文件类.X路径取扩展名(f.Filename)
	}
	filePath := 文件类.X路径生成(目录路径, name)
	newFile, 错误 := 文件类.X创建文件与目录(filePath)
	if 错误 != nil {
		return "", 错误
	}
	defer newFile.Close()
	intlog.Printf(f.ctx, `save upload file: %s`, filePath)
	if _, 错误 = io.Copy(newFile, file); 错误 != nil {
		错误 = 错误类.X多层错误并格式化(错误, `io.Copy failed from "%s" to "%s"`, f.Filename, filePath)
		return "", 错误
	}
	return 文件类.X路径取文件名(filePath), nil
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
func (fs X上传文件数组) X保存(目录路径 string, 随机重命名 ...bool) (文件名数组 []string, 错误 error) {
	if len(fs) == 0 {
		return nil, 错误类.X创建错误码(
			错误码类.CodeMissingParameter,
			"file array is empty, maybe you retrieve it from invalid field name or form enctype",
		)
	}
	for _, f := range fs {
		if filename, err := f.X保存(目录路径, 随机重命名...); err != nil {
			return 文件名数组, err
		} else {
			文件名数组 = append(文件名数组, filename)
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
func (r *X请求) X取上传文件对象(名称 string) *X上传文件 {
	uploadFiles := r.X取上传文件数组对象(名称)
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
func (r *X请求) X取上传文件数组对象(名称 string) X上传文件数组 {
	multipartFiles := r.X取multipart表单文件数组对象(名称)
	if len(multipartFiles) > 0 {
		uploadFiles := make(X上传文件数组, len(multipartFiles))
		for k, v := range multipartFiles {
			uploadFiles[k] = &X上传文件{
				ctx:        r.Context别名(),
				FileHeader: v,
			}
		}
		return uploadFiles
	}
	return nil
}
