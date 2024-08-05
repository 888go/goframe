// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gmetric

import (
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

func truncateGlobalAttributes() {
	globalAttributesMu.Lock()
	defer globalAttributesMu.Unlock()
	globalAttributes = make([]globalAttributeItem, 0)
}

func Test_GlobalAttributes(t *testing.T) {
	defer truncateGlobalAttributes()
	gtest.C(t, func(t *gtest.T) {
		SetGlobalAttributes(Attributes{
			NewAttribute("global", "gl"),
		}, SetGlobalAttributesOption{
			Instrument:        "",
			InstrumentVersion: "",
			InstrumentPattern: "",
		})

		SetGlobalAttributes(Attributes{
			NewAttribute("a", 1),
		}, SetGlobalAttributesOption{
			Instrument:        "ins_a",
			InstrumentVersion: "v1.0",
			InstrumentPattern: "",
		})
		SetGlobalAttributes(Attributes{
			NewAttribute("b", 2),
		}, SetGlobalAttributesOption{
			Instrument:        "ins_bb",
			InstrumentVersion: "v1.1",
			InstrumentPattern: "",
		})
		SetGlobalAttributes(Attributes{
			NewAttribute("c", 3),
		}, SetGlobalAttributesOption{
			Instrument:        "ins_bb",
			InstrumentVersion: "v1.1",
			InstrumentPattern: "",
		})
		SetGlobalAttributes(Attributes{
			NewAttribute("d", 4),
		}, SetGlobalAttributesOption{
			Instrument:        "",
			InstrumentVersion: "v1.0",
			InstrumentPattern: "ins.+",
		})
		SetGlobalAttributes(Attributes{
			NewAttribute("e", 5),
		}, SetGlobalAttributesOption{
			Instrument:        "",
			InstrumentVersion: "v1.0",
			InstrumentPattern: "ins_b.+",
		})
		SetGlobalAttributes(Attributes{
			NewAttribute("f", 6),
		}, SetGlobalAttributesOption{
			Instrument:        "",
			InstrumentVersion: "v1.1",
			InstrumentPattern: "ins_b.+",
		})

		t.Assert(GetGlobalAttributes(GetGlobalAttributesOption{
			Instrument:        "",
			InstrumentVersion: "",
		}), Attributes{
			NewAttribute("global", "gl"),
		})
		t.Assert(GetGlobalAttributes(GetGlobalAttributesOption{
			Instrument:        "ins_a",
			InstrumentVersion: "",
		}), Attributes{})
		t.Assert(GetGlobalAttributes(GetGlobalAttributesOption{
			Instrument:        "ins_a",
			InstrumentVersion: "1.1",
		}), Attributes{})
		t.Assert(GetGlobalAttributes(GetGlobalAttributesOption{
			Instrument:        "ins_bb",
			InstrumentVersion: "v1.0",
		}), Attributes{
			NewAttribute("d", 4),
			NewAttribute("e", 5),
		})
		t.Assert(GetGlobalAttributes(GetGlobalAttributesOption{
			Instrument:        "ins_bb",
			InstrumentVersion: "v1.1",
		}), Attributes{
			NewAttribute("b", 2),
			NewAttribute("c", 3),
			NewAttribute("f", 6),
		})
		t.Assert(GetGlobalAttributes(GetGlobalAttributesOption{
			Instrument:        "ins_cc",
			InstrumentVersion: "v1.1",
		}), Attributes{})
	})
}
