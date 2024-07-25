// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gvalid

import (
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

func Test_parseSequenceTag(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := "name@required|length:2,20|password3|same:password1#||密码强度不足|两次密码不一致"
		field, rule, msg := ParseTagValue(s)
		t.Assert(field, "name")
		t.Assert(rule, "required|length:2,20|password3|same:password1")
		t.Assert(msg, "||密码强度不足|两次密码不一致")
	})
	gtest.C(t, func(t *gtest.T) {
		s := "required|length:2,20|password3|same:password1#||密码强度不足|两次密码不一致"
		field, rule, msg := ParseTagValue(s)
		t.Assert(field, "")
		t.Assert(rule, "required|length:2,20|password3|same:password1")
		t.Assert(msg, "||密码强度不足|两次密码不一致")
	})
	gtest.C(t, func(t *gtest.T) {
		s := "required|length:2,20|password3|same:password1"
		field, rule, msg := ParseTagValue(s)
		t.Assert(field, "")
		t.Assert(rule, "required|length:2,20|password3|same:password1")
		t.Assert(msg, "")
	})
	gtest.C(t, func(t *gtest.T) {
		s := "required"
		field, rule, msg := ParseTagValue(s)
		t.Assert(field, "")
		t.Assert(rule, "required")
		t.Assert(msg, "")
	})
}

func Test_GetTags(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(structTagPriority, GetTags())
	})
}
