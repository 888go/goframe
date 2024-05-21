// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package builtin

import (
	"errors"

	"github.com/gogf/gf/v2/text/gregex"
)

// RulePostcode implements `postcode` rule:
// Postcode number.
//
// Format: postcode
type RulePostcode struct{}

func init() {
	Register(RulePostcode{})
}


// ff:
func (r RulePostcode) Name() string {
	return "postcode"
}


// ff:
func (r RulePostcode) Message() string {
	return "The {field} value `{value}` is not a valid postcode format"
}


// ff:
// in:
func (r RulePostcode) Run(in RunInput) error {
	ok := gregex.IsMatchString(
		`^\d{6}$`,
		in.Value.String(),
	)
	if ok {
		return nil
	}
	return errors.New(in.Message)
}
