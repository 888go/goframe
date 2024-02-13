// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package html类_test

import (
	"testing"
	
	"github.com/888go/goframe/encoding/ghtml"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/test/gtest"
)

func Test_StripTags(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		src := `<p>Test paragraph.</p><!-- Comment -->  <a href="#fragment">Other text</a>`
		dst := `Test paragraph.  Other text`
		t.Assert(html类.X删除HTML标记(src), dst)
	})
}

func Test_Entities(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		src := `A 'quote' "is" <b>bold</b>`
		dst := `A &#39;quote&#39; &#34;is&#34; &lt;b&gt;bold&lt;/b&gt;`
		t.Assert(html类.X编码(src), dst)
		t.Assert(html类.X解码(dst), src)
	})
}

func Test_SpecialChars(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		src := `A 'quote' "is" <b>bold</b>`
		dst := `A &#39;quote&#39; &#34;is&#34; &lt;b&gt;bold&lt;/b&gt;`
		t.Assert(html类.X编码特殊字符(src), dst)
		t.Assert(html类.X解码特殊字符(dst), src)
	})
}

func Test_SpecialCharsMapOrStruct_Map(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a := g.Map{
			"Title":   "<h1>T</h1>",
			"Content": "<div>C</div>",
		}
		err := html类.X编码Map(a)
		t.AssertNil(err)
		t.Assert(a["Title"], `&lt;h1&gt;T&lt;/h1&gt;`)
		t.Assert(a["Content"], `&lt;div&gt;C&lt;/div&gt;`)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		a := g.MapStrStr{
			"Title":   "<h1>T</h1>",
			"Content": "<div>C</div>",
		}
		err := html类.X编码Map(a)
		t.AssertNil(err)
		t.Assert(a["Title"], `&lt;h1&gt;T&lt;/h1&gt;`)
		t.Assert(a["Content"], `&lt;div&gt;C&lt;/div&gt;`)
	})
}

func Test_SpecialCharsMapOrStruct_Struct(t *testing.T) {
	type A struct {
		Title   string
		Content string
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		a := &A{
			Title:   "<h1>T</h1>",
			Content: "<div>C</div>",
		}
		err := html类.X编码Map(a)
		t.AssertNil(err)
		t.Assert(a.Title, `&lt;h1&gt;T&lt;/h1&gt;`)
		t.Assert(a.Content, `&lt;div&gt;C&lt;/div&gt;`)
	})
}
