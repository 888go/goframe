// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gsel

type builderLeastConnection struct{}

// ff:
func NewBuilderLeastConnection() Builder {
	return &builderLeastConnection{}
}

// ff:
func (*builderLeastConnection) Name() string {
	return "BalancerLeastConnection"
}

// ff:
func (*builderLeastConnection) Build() Selector {
	return NewSelectorLeastConnection()
}
