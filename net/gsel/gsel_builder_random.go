// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gsel

type builderRandom struct{}


// ff:
func NewBuilderRandom() Builder {
	return &builderRandom{}
}


// ff:
func (*builderRandom) Name() string {
	return "BalancerRandom"
}


// ff:
func (*builderRandom) Build() Selector {
	return NewSelectorRandom()
}
