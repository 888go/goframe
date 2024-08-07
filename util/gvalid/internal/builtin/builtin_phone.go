// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package builtin

import (
	"errors"

	gregex "github.com/888go/goframe/text/gregex"
)

// RulePhone 实现了 `phone` 规则：
//
//  1. 中国移动:
//     134, 135, 136, 137, 138, 139, 150, 151, 152, 157, 158, 159, 182, 183, 184, 187, 188,
//     178(4G), 147(网络)；
//     172
//
//  2. 中国联通:
//     130, 131, 132, 155, 156, 185, 186 ,176(4G), 145(网络), 175
//
//  3. 中国电信:
//     133, 153, 180, 181, 189, 177(4G)
//
//  4. 卫星电话:
//     1349
//
//  5. 虚拟运营商:
//     170, 173
//
//  6. 2018年后新增:
//     16x, 19x
//
// 格式: 手机号码
// md5:bbaf43c95c780522
type RulePhone struct{}

func init() {
	Register(RulePhone{})
}

func (r RulePhone) Name() string {
	return "phone"
}

func (r RulePhone) Message() string {
	return "The {field} value `{value}` is not a valid phone number"
}

func (r RulePhone) Run(in RunInput) error {
	ok := gregex.X是否匹配文本(
		`^13[\d]{9}$|^14[5,7]{1}\d{8}$|^15[^4]{1}\d{8}$|^16[\d]{9}$|^17[0,2,3,5,6,7,8]{1}\d{8}$|^18[\d]{9}$|^19[\d]{9}$`,
		in.Value.String(),
	)
	if ok {
		return nil
	}
	return errors.New(in.Message)
}
