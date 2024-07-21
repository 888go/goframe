// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gutil

import (
	"github.com/gogf/gf/v2/internal/reflection"
)

type (
	OriginValueAndKindOutput = reflection.OriginValueAndKindOutput
	OriginTypeAndKindOutput  = reflection.OriginTypeAndKindOutput
)

// OriginValueAndKind 获取并返回原始的反射值和类型。 md5:e2cdf5d39aa2b981
// ff:
// value:
// out:
func OriginValueAndKind(value interface{}) (out OriginValueAndKindOutput) {
	return reflection.OriginValueAndKind(value)
}

// OriginTypeAndKind 获取并返回原始反射类型和种类。 md5:ee62836f1445e493
// ff:
// value:
// out:
func OriginTypeAndKind(value interface{}) (out OriginTypeAndKindOutput) {
	return reflection.OriginTypeAndKind(value)
}
