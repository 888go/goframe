package data

import (
	gres "github.com/888go/goframe/os/gres"
)

func init() {
	if err := gres.Add("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"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
