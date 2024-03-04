// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gjson

import (
	"testing"
	
	"github.com/gogf/gf/v2/test/gtest"
)

func Test_checkDataType(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		data := []byte(`
bb           = """
                   dig := dig;                         END;"""
`)
		t.Assert(checkDataType(data), "toml")
	})

	gtest.C(t, func(t *gtest.T) {
		data := []byte(`
# 模板引擎目录
viewpath = "/home/www/templates/"
# MySQL数据库配置
[redis]
dd = 11
[redis]
    disk  = "127.0.0.1:6379,0"
    cache = "127.0.0.1:6379,1"
`)
		t.Assert(checkDataType(data), "toml")
	})

	gtest.C(t, func(t *gtest.T) {
		data := []byte(`
"gf.gvalid.rule.required"             = "The :attribute field is required"
"gf.gvalid.rule.required-if"          = "The :attribute field is required"
"gf.gvalid.rule.required-unless"      = "The :attribute field is required"
"gf.gvalid.rule.required-with"        = "The :attribute field is required"
"gf.gvalid.rule.required-with-all"    = "The :attribute field is required"
"gf.gvalid.rule.required-without"     = "The :attribute field is required"
"gf.gvalid.rule.required-without-all" = "The :attribute field is required"
"gf.gvalid.rule.date"                 = "The :attribute value is not a valid date"
"gf.gvalid.rule.date-format"          = "The :attribute value does not match the format :format"
"gf.gvalid.rule.email"                = "The :attribute value must be a valid email address"
"gf.gvalid.rule.phone"                = "The :attribute value must be a valid phone number"
"gf.gvalid.rule.telephone"            = "The :attribute value must be a valid telephone number"
"gf.gvalid.rule.passport"             = "The :attribute value is not a valid passport format"
"gf.gvalid.rule.password"             = "The :attribute value is not a valid passport format"
"gf.gvalid.rule.password2"            = "The :attribute value is not a valid passport format"
"gf.gvalid.rule.password3"            = "The :attribute value is not a valid passport format"
"gf.gvalid.rule.postcode"             = "The :attribute value is not a valid passport format"
"gf.gvalid.rule.resident-id"          = "The :attribute value is not a valid resident id number"
"gf.gvalid.rule.bank-card"            = "The :attribute value must be a valid bank card number"
"gf.gvalid.rule.qq"                   = "The :attribute value must be a valid QQ number"
"gf.gvalid.rule.ip"                   = "The :attribute value must be a valid IP address"
"gf.gvalid.rule.ipv4"                 = "The :attribute value must be a valid IPv4 address"
"gf.gvalid.rule.ipv6"                 = "The :attribute value must be a valid IPv6 address"
"gf.gvalid.rule.mac"                  = "The :attribute value must be a valid MAC address"
"gf.gvalid.rule.url"                  = "The :attribute value must be a valid URL address"
"gf.gvalid.rule.domain"               = "The :attribute value must be a valid domain format"
"gf.gvalid.rule.length"               = "The :attribute value length must be between :min and :max"
"gf.gvalid.rule.min-length"           = "The :attribute value length must be equal or greater than :min"
"gf.gvalid.rule.max-length"           = "The :attribute value length must be equal or lesser than :max"
"gf.gvalid.rule.between"              = "The :attribute value must be between :min and :max"
"gf.gvalid.rule.min"                  = "The :attribute value must be equal or greater than :min"
"gf.gvalid.rule.max"                  = "The :attribute value must be equal or lesser than :max"
"gf.gvalid.rule.json"                 = "The :attribute value must be a valid JSON string"
"gf.gvalid.rule.xml"                  = "The :attribute value must be a valid XML string"
"gf.gvalid.rule.array"                = "The :attribute value must be an array"
"gf.gvalid.rule.integer"              = "The :attribute value must be an integer"
"gf.gvalid.rule.float"                = "The :attribute value must be a float"
"gf.gvalid.rule.boolean"              = "The :attribute value field must be true or false"
"gf.gvalid.rule.same"                 = "The :attribute value must be the same as field :field"
"gf.gvalid.rule.different"            = "The :attribute value must be different from field :field"
"gf.gvalid.rule.in"                   = "The :attribute value is not in acceptable range"
"gf.gvalid.rule.not-in"               = "The :attribute value is not in acceptable range"
"gf.gvalid.rule.regex"                = "The :attribute value is invalid"
`)
// 判断data是否匹配正则表达式：以任意空白符（包括空格、制表符、换行符、回车符）开始，后面跟着至少一个由字母、数字或破折号组成的单词，其后是任意数量的空白符，然后是一个冒号和任意数量的空白符，最后是一个由双引号包裹的任意非空字符串。
// fmt.Println(gregex.IsMatch(`^[\s\t\n\r]*[\w\-]+\s*:\s*".+"`, data))
// 判断data是否匹配正则表达式：以任意空白符开始，后面跟着至少一个由字母、数字或破折号组成的单词，其后是任意数量的空白符，然后是一个冒号和任意数量的空白符，最后是一个由任意数量的字母或数字组成的非空字符串。
// fmt.Println(gregex.IsMatch(`^[\s\t\n\r]*[\w\-]+\s*:\s*\w+`, data))
// 判断data是否匹配正则表达式：包含至少一个空白符（包括空格、制表符、换行符、回车符），后面跟着至少一个由字母、数字或破折号组成的单词，其后是任意数量的空白符，然后是一个冒号和任意数量的空白符，最后是一个由双引号包裹的任意非空字符串。
// fmt.Println(gregex.IsMatch(`[\s\t\n\r]+[\w\-]+\s*:\s*".+"`, data))
// 判断data是否匹配正则表达式：包含至少一个换行符或回车符，后面跟着至少一个由字母、数字、破折号或空白符组成的单词，其后是任意数量的空白符，然后是一个冒号和任意数量的空白符，最后是一个由任意数量的字母或数字组成的非空字符串。
// fmt.Println(gregex.IsMatch(`[\n\r]+[\w\-\s\t]+\s*:\s*\w+`, data))
// 判断将data转换为字符串类型后，是否匹配正则表达式：包含至少一个换行符或回车符，后面跟着至少一个由字母、数字、破折号或空白符组成的单词，其后是任意数量的空白符，然后是一个冒号和任意数量的空白符，最后是一个由任意数量的字母或数字组成的非空字符串。
// fmt.Println(gregex.MatchString(`[\n\r]+[\w\-\s\t]+\s*:\s*\w+`, string(data)))
		t.Assert(checkDataType(data), "toml")
	})

	gtest.C(t, func(t *gtest.T) {
		data := []byte(`
[default]
db.engine         = mysql
db.max.idle.conns = 5
db.max.open.conns = 100
allow_ips         =
api.key           =
api.secret        =
enable_tls        = false
concurrency.queue = 500
auth_secret       = 63358e6f3daf0e5775ec3fb4d2516b01d41530bf30960aa76972f6ce7e08552f
ca_file           =
cert_file         =
key_file          =
host_port         = 8088
log_path          = /Users/zhaosuji/go/src/git.medlinker.com/foundations/gocron/log
#k8s-api地址(只提供内网访问)
k8s-inner-api = http://127.0.0.1:8081/kube/add
conf_dir = ./config
app_conf = ./config/app.ini
`)
		t.Assert(checkDataType(data), "ini")
	})

	gtest.C(t, func(t *gtest.T) {
		data := []byte(`
# API Server
[server]
    address = ":8199"

# Jenkins
[jenkins]
    url          = "https://jenkins-swimlane.com"
    nodeJsStaticBuildCmdTpl = """
npm i --registry=https://registry.npm.taobao.org
wget http:// 从consul.infra的8500端口获取v1版本的kv存储中，键为'app_{{.SwimlaneName}}/{{.RepoName}}/.env.qa'的值，并以原始数据（raw）格式获取，然后将结果输出并重定向保存到当前目录下的.env.qa文件
npm run build:qa
"""
`)
		t.Assert(checkDataType(data), "toml")
	})
}
