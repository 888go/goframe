// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包gtoml提供了访问和转换TOML内容的功能。 md5:0bf787c2f078634a
package toml类

import (
	"bytes"

	"github.com/BurntSushi/toml"

	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/json"
)

func Encode(v interface{}) ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	if err := toml.NewEncoder(buffer).Encode(v); err != nil {
		err = gerror.Wrap(err, `toml.Encoder.Encode failed`)
		return nil, err
	}
	return buffer.Bytes(), nil
}

func Decode(v []byte) (interface{}, error) {
	var result interface{}
	if err := toml.Unmarshal(v, &result); err != nil {
		err = gerror.Wrap(err, `toml.Unmarshal failed`)
		return nil, err
	}
	return result, nil
}

func DecodeTo(v []byte, result interface{}) (err error) {
	err = toml.Unmarshal(v, result)
	if err != nil {
		err = gerror.Wrap(err, `toml.Unmarshal failed`)
	}
	return err
}

func ToJson(v []byte) ([]byte, error) {
	if r, err := Decode(v); err != nil {
		return nil, err
	} else {
		return json.Marshal(r)
	}
}
