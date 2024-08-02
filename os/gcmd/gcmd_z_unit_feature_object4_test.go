// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package cmd类_test

import (
	"context"
	"os"
	"testing"

	"github.com/888go/goframe/frame/g"
	gcmd "github.com/888go/goframe/os/gcmd"
	gctx "github.com/888go/goframe/os/gctx"
	gtest "github.com/888go/goframe/test/gtest"
)

type TestNoNameTagCase struct {
	g.Meta `name:"root"`
}

type TestNoNameTagCaseRootInput struct {
	Name string
}

type TestNoNameTagCaseRootOutput struct {
	Content string
}

func (c *TestNoNameTagCase) TEST(ctx context.Context, in TestNoNameTagCaseRootInput) (out *TestNoNameTagCaseRootOutput, err error) {
	out = &TestNoNameTagCaseRootOutput{
		Content: in.Name,
	}
	return
}

func Test_Command_NoNameTagCase(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var ctx = gctx.New()
		cmd, err := gcmd.NewFromObject(TestNoNameTagCase{})
		t.AssertNil(err)

		os.Args = []string{"root", "TEST", "-name=john"}
		value, err := cmd.RunWithValueError(ctx)
		t.AssertNil(err)
		t.Assert(value, `{"Content":"john"}`)
	})
}
