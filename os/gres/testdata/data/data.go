package data

import "coding.net/gogit/go/goframe/os/gres"

func init() {
	if err := gres.Add("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"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
