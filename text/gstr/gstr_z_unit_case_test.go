// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 文本类_test

import (
	"testing"
	
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/text/gstr"
)

func Test_CaseCamel(t *testing.T) {
	cases := [][]string{
		{"test_case", "TestCase"},
		{"test", "Test"},
		{"TestCase", "TestCase"},
		{" test  case ", "TestCase"},
		{"userLogin_log.bak", "UserLoginLogBak"},
		{"", ""},
		{"many_many_words", "ManyManyWords"},
		{"AnyKind of_string", "AnyKindOfString"},
		{"odd-fix", "OddFix"},
		{"numbers2And55with000", "Numbers2And55With000"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := 文本类.X命名转换到首字母大写驼峰(in)
		if result != out {
			t.Error("'" + result + "' != '" + out + "'")
		}
	}
}

func Test_CaseCamelLower(t *testing.T) {
	cases := [][]string{
		{"foo-bar", "fooBar"},
		{"TestCase", "testCase"},
		{"", ""},
		{"AnyKind of_string", "anyKindOfString"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := 文本类.X命名转换到首字母小写驼峰(in)
		if result != out {
			t.Error("'" + result + "' != '" + out + "'")
		}
	}
}

func Test_CaseSnake(t *testing.T) {
	cases := [][]string{
		{"testCase", "test_case"},
		{"TestCase", "test_case"},
		{"Test Case", "test_case"},
		{" Test Case", "test_case"},
		{"Test Case ", "test_case"},
		{" Test Case ", "test_case"},
		{"test", "test"},
		{"test_case", "test_case"},
		{"Test", "test"},
		{"", ""},
		{"ManyManyWords", "many_many_words"},
		{"manyManyWords", "many_many_words"},
		{"AnyKind of_string", "any_kind_of_string"},
		{"numbers2and55with000", "numbers_2_and_55_with_000"},
		{"JSONData", "json_data"},
		{"userID", "user_id"},
		{"AAAbbb", "aa_abbb"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := 文本类.X命名转换到全小写蛇形(in)
		if result != out {
			t.Error("'" + in + "'('" + result + "' != '" + out + "')")
		}
	}
}

func Test_CaseDelimited(t *testing.T) {
	cases := [][]string{
		{"testCase", "test@case"},
		{"TestCase", "test@case"},
		{"Test Case", "test@case"},
		{" Test Case", "test@case"},
		{"Test Case ", "test@case"},
		{" Test Case ", "test@case"},
		{"test", "test"},
		{"test_case", "test@case"},
		{"Test", "test"},
		{"", ""},
		{"ManyManyWords", "many@many@words"},
		{"manyManyWords", "many@many@words"},
		{"AnyKind of_string", "any@kind@of@string"},
		{"numbers2and55with000", "numbers@2@and@55@with@000"},
		{"JSONData", "json@data"},
		{"userID", "user@id"},
		{"AAAbbb", "aa@abbb"},
		{"test-case", "test@case"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := 文本类.X命名转换按符号(in, '@')
		if result != out {
			t.Error("'" + in + "' ('" + result + "' != '" + out + "')")
		}
	}
}

func Test_CaseSnakeScreaming(t *testing.T) {
	cases := [][]string{
		{"testCase", "TEST_CASE"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := 文本类.X命名转换到大写蛇形(in)
		if result != out {
			t.Error("'" + result + "' != '" + out + "'")
		}
	}
}

func Test_CaseKebab(t *testing.T) {
	cases := [][]string{
		{"testCase", "test-case"},
		{"optimization1.0.0", "optimization-1-0-0"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := 文本类.X命名转换到小写短横线(in)
		if result != out {
			t.Error("'" + result + "' != '" + out + "'")
		}
	}
}

func Test_CaseKebabScreaming(t *testing.T) {
	cases := [][]string{
		{"testCase", "TEST-CASE"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := 文本类.X命名转换到大写驼峰短横线(in)
		if result != out {
			t.Error("'" + result + "' != '" + out + "'")
		}
	}
}

func Test_CaseDelimitedScreaming(t *testing.T) {
	cases := [][]string{
		{"testCase", "TEST.CASE"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := 文本类.X命名转换按符号与大小写(in, '.', true)
		if result != out {
			t.Error("'" + result + "' != '" + out + "'")
		}
	}
}

func Test_CaseSnakeFirstUpper(t *testing.T) {
	cases := [][]string{
		{"RGBCodeMd5", "rgb_code_md5"},
		{"testCase", "test_case"},
		{"Md5", "md5"},
		{"userID", "user_id"},
		{"RGB", "rgb"},
		{"RGBCode", "rgb_code"},
		{"_ID", "id"},
		{"User_ID", "user_id"},
		{"user_id", "user_id"},
		{"md5", "md5"},
		{"Numbers2And55With000", "numbers2_and55_with000"},
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		for _, item := range cases {
			t.Assert(文本类.X命名转换到全小写蛇形2(item[0]), item[1])
		}

		t.Assert(文本类.X命名转换到全小写蛇形2("RGBCodeMd5", "."), "rgb.code.md5")
	})

}
func Test_CaseTypeMatch(t *testing.T) {
	caseTypes := []文本类.CaseType{
		文本类.Camel,
		文本类.CamelLower,
		文本类.Snake,
		文本类.SnakeFirstUpper,
		文本类.SnakeScreaming,
		文本类.Kebab,
		文本类.KebabScreaming,
		文本类.Lower,
		"test", // invalid case type
	}
	testCaseTypes := []string{
		"camel",
		"camelLower",
		"snake",
		"snakeFirstUpper",
		"snakeScreaming",
		"kebab",
		"kebabScreaming",
		"lower",
		"test",
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		for i := 0; i < len(caseTypes); i++ {
			t.Assert(文本类.X命名方式判断(testCaseTypes[i]), caseTypes[i])
		}
	})
}

func Test_CaseConvert(t *testing.T) {
	caseTypes := []文本类.CaseType{
		文本类.Camel,
		文本类.CamelLower,
		文本类.Snake,
		文本类.SnakeFirstUpper,
		文本类.SnakeScreaming,
		文本类.Kebab,
		文本类.KebabScreaming,
		文本类.Lower,
		"test", // invalid case type
		"",     // invalid case type
	}
	testCaseTypes := []string{
		"AnyKindOfString",    // Camel
		"anyKindOfString",    // CamelLower
		"any_kind_of_string", // Snake
		"any_kind_of_string", // SnakeFirstUpper
		"ANY_KIND_OF_STRING", // SnakeScreaming
		"any-kind-of-string", // Kebab
		"ANY-KIND-OF-STRING", // KebabScreaming
		"any_kind_of_string", // Lower
		"any_kind_of_string", // invalid case type
		"any_kind_of_string", // invalid case type
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		for i := 0; i < len(caseTypes); i++ {
			t.Assert(文本类.X命名转换("any_kind_of_string", caseTypes[i]), testCaseTypes[i])
			t.Logf("test case: %s success", caseTypes[i])
		}
	})
}
