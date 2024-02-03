// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gvalid

import (
	"testing"
	
	"github.com/888go/goframe/test/gtest"
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
