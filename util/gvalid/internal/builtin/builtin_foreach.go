// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package builtin

// RuleForeach implements `foreach` rule:
// It tells the next validation using current value as an array and validates each of its element.
//
// Format: foreach
type RuleForeach struct{}

func init() {
	Register(RuleForeach{})
}


// ff:
func (r RuleForeach) Name() string {
	return "foreach"
}


// ff:
func (r RuleForeach) Message() string {
	return ""
}


// ff:
// in:
func (r RuleForeach) Run(in RunInput) error {
	return nil
}
