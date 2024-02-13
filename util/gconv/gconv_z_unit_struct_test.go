// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 转换类_test

import (
	"testing"
	"time"
	
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/encoding/gjson"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gconv"
)

func Test_Struct_Basic1(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Uid      int
			Name     string
			Site_Url string
			NickName string
			Pass1    string `gconv:"password1"`
			Pass2    string `gconv:"password2"`
		}
		user := new(User)
		params1 := g.Map{
			"uid":       1,
			"Name":      "john",
			"siteurl":   "https://goframe.org",
			"nick_name": "johng",
			"PASS1":     "123",
			"PASS2":     "456",
		}
		if err := 转换类.Struct(params1, user); err != nil {
			t.Error(err)
		}
		t.Assert(user, &User{
			Uid:      1,
			Name:     "john",
			Site_Url: "https://goframe.org",
			NickName: "johng",
			Pass1:    "123",
			Pass2:    "456",
		})

		user = new(User)
		params2 := g.Map{
			"uid":       2,
			"name":      "smith",
			"site-url":  "https://goframe.org",
			"nick name": "johng",
			"password1": "111",
			"password2": "222",
		}
		if err := 转换类.Struct(params2, user); err != nil {
			t.Error(err)
		}
		t.Assert(user, &User{
			Uid:      2,
			Name:     "smith",
			Site_Url: "https://goframe.org",
			NickName: "johng",
			Pass1:    "111",
			Pass2:    "222",
		})
	})
}

func Test_Struct_Basic2(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Uid     int
			Name    string
			SiteUrl string
			Pass1   string
			Pass2   string
		}
		user := new(User)
		params := g.Map{
			"uid":      1,
			"Name":     "john",
			"site_url": "https://goframe.org",
			"PASS1":    "123",
			"PASS2":    "456",
		}
		if err := 转换类.Struct(params, user); err != nil {
			t.Error(err)
		}
		t.Assert(user, &User{
			Uid:     1,
			Name:    "john",
			SiteUrl: "https://goframe.org",
			Pass1:   "123",
			Pass2:   "456",
		})
	})
}

func Test_Struct_Attr_Pointer(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Uid  *int
			Name *string
		}
		user := new(User)
		params := g.Map{
			"uid":  "1",
			"Name": "john",
		}
		if err := 转换类.Struct(params, user); err != nil {
			t.Error(err)
		}
		t.Assert(user.Uid, 1)
		t.Assert(*user.Name, "john")
	})
}

func Test_Struct_Attr_Slice1(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Scores []int
		}
		scores := []interface{}{99, 100, 60, 140}
		user := new(User)
		if err := 转换类.Struct(g.Map{"Scores": scores}, user); err != nil {
			t.Error(err)
		} else {
			t.Assert(user, &User{
				Scores: []int{99, 100, 60, 140},
			})
		}
	})
}

// 目前还不支持这种类型的转换。
//func Test_Struct_Attr_Slice2(t *testing.T) {
//	gtest.C(t, func(t *gtest.T) { // 使用gtest框架进行测试
//		// 定义User结构体，其中Scores字段为二维整数切片
//		type User struct {
//			Scores [][]int
//		}
//		// 创建一个interface{}类型的scores变量，存储嵌套的整数切片
//		scores := []interface{}{[]interface{}{99, 100, 60, 140}}
//		// 初始化一个新的User结构体指针user
//		user := new(User)
//		// 尝试使用gconv.Struct将scores映射到user结构体中
//		if err := gconv.Struct(g.Map{"Scores": scores}, user); err != nil {
//			// 如果转换过程中出现错误，则输出错误信息
//			t.Error(err)
//		} else {
//			// 如果转换成功，则断言user的Scores字段与预期的二维整数切片相等
//			t.Assert(user, &User{
//				Scores: [][]int{{99, 100, 60, 140}},
//			})
//		}
//	})
//}

func Test_Struct_Attr_Struct(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type Score struct {
			Name   string
			Result int
		}
		type User struct {
			Scores Score
		}

		user := new(User)
		scores := map[string]interface{}{
			"Scores": map[string]interface{}{
				"Name":   "john",
				"Result": 100,
			},
		}

		// 嵌套struct转换
		if err := 转换类.Struct(scores, user); err != nil {
			t.Error(err)
		} else {
			t.Assert(user, &User{
				Scores: Score{
					Name:   "john",
					Result: 100,
				},
			})
		}
	})
}

func Test_Struct_Attr_Struct_Ptr(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type Score struct {
			Name   string
			Result int
		}
		type User struct {
			Scores *Score
		}

		user := new(User)
		scores := map[string]interface{}{
			"Scores": map[string]interface{}{
				"Name":   "john",
				"Result": 100,
			},
		}

		// 嵌套struct转换
		if err := 转换类.Struct(scores, user); err != nil {
			t.Error(err)
		} else {
			t.Assert(user.Scores, &Score{
				Name:   "john",
				Result: 100,
			})
		}
	})
}

func Test_Struct_Attr_Struct_Slice1(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type Score struct {
			Name   string
			Result int
		}
		type User struct {
			Scores []Score
		}

		user := new(User)
		scores := map[string]interface{}{
			"Scores": map[string]interface{}{
				"Name":   "john",
				"Result": 100,
			},
		}

		if err := 转换类.Struct(scores, user); err != nil {
			t.Error(err)
		} else {
			t.Assert(user.Scores, []Score{
				{
					Name:   "john",
					Result: 100,
				},
			})
		}
	})
}

func Test_Struct_Attr_Struct_Slice2(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type Score struct {
			Name   string
			Result int
		}
		type User struct {
			Scores []Score
		}

		user := new(User)
		scores := map[string]interface{}{
			"Scores": []interface{}{
				map[string]interface{}{
					"Name":   "john",
					"Result": 100,
				},
				map[string]interface{}{
					"Name":   "smith",
					"Result": 60,
				},
			},
		}

		if err := 转换类.Struct(scores, user); err != nil {
			t.Error(err)
		} else {
			t.Assert(user.Scores, []Score{
				{
					Name:   "john",
					Result: 100,
				},
				{
					Name:   "smith",
					Result: 60,
				},
			})
		}
	})
}

func Test_Struct_Attr_Struct_Slice_Ptr(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type Score struct {
			Name   string
			Result int
		}
		type User struct {
			Scores []*Score
		}

		user := new(User)
		scores := map[string]interface{}{
			"Scores": []interface{}{
				map[string]interface{}{
					"Name":   "john",
					"Result": 100,
				},
				map[string]interface{}{
					"Name":   "smith",
					"Result": 60,
				},
			},
		}

		// 嵌套struct转换，属性为slice类型，数值为slice map类型
		if err := 转换类.Struct(scores, user); err != nil {
			t.Error(err)
		} else {
			t.Assert(len(user.Scores), 2)
			t.Assert(user.Scores[0], &Score{
				Name:   "john",
				Result: 100,
			})
			t.Assert(user.Scores[1], &Score{
				Name:   "smith",
				Result: 60,
			})
		}
	})
}

func Test_Struct_Attr_CustomType1(t *testing.T) {
	type MyInt int
	type User struct {
		Id   MyInt
		Name string
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		user := new(User)
		err := 转换类.Struct(g.Map{"id": 1, "name": "john"}, user)
		t.AssertNil(err)
		t.Assert(user.Id, 1)
		t.Assert(user.Name, "john")
	})
}

func Test_Struct_Attr_CustomType2(t *testing.T) {
	type MyInt int
	type User struct {
		Id   []MyInt
		Name string
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		user := new(User)
		err := 转换类.Struct(g.Map{"id": g.Slice别名{1, 2}, "name": "john"}, user)
		t.AssertNil(err)
		t.Assert(user.Id, g.Slice别名{1, 2})
		t.Assert(user.Name, "john")
	})
}

// 来源：k8s.io/apimachinery@v0.22.0/pkg/apis/meta/v1/duration.go
// （由于您未提供具体的代码行，以下为一般性的翻译说明）
// 该注释表明了该代码片段所在的仓库、版本及路径信息：
// - k8s.io：Kubernetes的GitHub组织名
// - apimachinery：Kubernetes中处理通用API对象和元数据的包
// - v0.22.0：此代码对应的Kubernetes API machinery包的版本号
// - pkg/apis/meta/v1：在apimachinery包下的具体子目录，表示该代码与Kubernetes元数据API的v1版本相关
// - duration.go：此文件主要定义或处理与时间持续相关的功能
type MyDuration struct {
	time.Duration
}

// UnmarshalJSON 实现了 json.Unmarshaller 接口。
func (d *MyDuration) UnmarshalJSON(b []byte) error {
	var str string
	err := json.Unmarshal(b, &str)
	if err != nil {
		return err
	}

	pd, err := time.ParseDuration(str)
	if err != nil {
		return err
	}
	d.Duration = pd
	return nil
}

func Test_Struct_Attr_CustomType3(t *testing.T) {
	type Config struct {
		D MyDuration
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		config := new(Config)
		err := 转换类.Struct(g.Map{"d": "15s"}, config)
		t.AssertNil(err)
		t.Assert(config.D, "15s")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		config := new(Config)
		err := 转换类.Struct(g.Map{"d": `"15s"`}, config)
		t.AssertNil(err)
		t.Assert(config.D, "15s")
	})
}

func Test_Struct_PrivateAttribute(t *testing.T) {
	type User struct {
		Id   int
		name string
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		user := new(User)
		err := 转换类.Struct(g.Map{"id": 1, "name": "john"}, user)
		t.AssertNil(err)
		t.Assert(user.Id, 1)
		t.Assert(user.name, "")
	})
}

func Test_StructEmbedded1(t *testing.T) {
	type Base struct {
		Age int
	}
	type User struct {
		Id   int
		Name string
		Base
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		user := new(User)
		params := g.Map{
			"id":   1,
			"name": "john",
			"age":  18,
		}
		err := 转换类.Struct(params, user)
		t.AssertNil(err)
		t.Assert(user.Id, params["id"])
		t.Assert(user.Name, params["name"])
		t.Assert(user.Age, 18)
	})
}

func Test_StructEmbedded2(t *testing.T) {
	type Ids struct {
		Id  int
		Uid int
	}
	type Base struct {
		Ids
		Time string
	}
	type User struct {
		Base
		Name string
	}
	params := g.Map{
		"id":   1,
		"uid":  10,
		"name": "john",
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		user := new(User)
		err := 转换类.Struct(params, user)
		t.AssertNil(err)
		t.Assert(user.Id, 1)
		t.Assert(user.Uid, 10)
		t.Assert(user.Name, "john")
	})
}

func Test_StructEmbedded3(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type Ids struct {
			Id  int `json:"id"`
			Uid int `json:"uid"`
		}
		type Base struct {
			Ids
			CreateTime string `json:"create_time"`
		}
		type User struct {
			Base
			Passport string `json:"passport"`
			Password string `json:"password"`
			Nickname string `json:"nickname"`
		}
		data := g.Map{
			"id":          100,
			"uid":         101,
			"passport":    "t1",
			"password":    "123456",
			"nickname":    "T1",
			"create_time": "2019",
		}
		user := new(User)
		err := 转换类.Struct(data, user)
		t.AssertNil(err)
		t.Assert(user.Id, 100)
		t.Assert(user.Uid, 101)
		t.Assert(user.Nickname, "T1")
		t.Assert(user.CreateTime, "2019")
	})
}

// 这是Go语言代码中的一行注释，其内容引用了GitHub上gogf/gf仓库下的第775号问题。 
// 翻译成中文：
// 引用了GitHub上gogf/gf项目中的第775号议题。
func Test_StructEmbedded4(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type Sub2 struct {
			SubName string
		}
		type sub1 struct {
			Sub2
			Name string
		}
		type Test struct {
			Sub sub1 `json:"sub"`
		}

		data := `{
    "sub": {
		"map":{"k":"v"},
        "Name": "name",
        "SubName": "subname"
    }}`

		expect := Test{
			Sub: sub1{
				Name: "name",
				Sub2: Sub2{
					SubName: "subname",
				},
			},
		}
		tx := new(Test)
		if err := 转换类.Struct(data, &tx); err != nil {
			panic(err)
		}
		t.Assert(tx, expect)
	})
}

func Test_StructEmbedded5(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type Base struct {
			Pass1 string `param:"password1"`
			Pass2 string `param:"password2"`
		}
		type UserWithBase1 struct {
			Id   int
			Name string
			Base
		}
		type UserWithBase2 struct {
			Id   int
			Name string
			Pass Base
		}

		data := g.Map{
			"id":        1,
			"name":      "john",
			"password1": "123",
			"password2": "456",
		}
		var err error
		user1 := new(UserWithBase1)
		user2 := new(UserWithBase2)
		err = 转换类.Struct(data, user1)
		t.AssertNil(err)
		t.Assert(user1, &UserWithBase1{1, "john", Base{"123", "456"}})

		err = 转换类.Struct(data, user2)
		t.AssertNil(err)
		t.Assert(user2, &UserWithBase2{1, "john", Base{"", ""}})

		var user3 *UserWithBase1
		err = 转换类.Struct(user1, &user3)
		t.AssertNil(err)
		t.Assert(user3, user1)
	})
}

func Test_Struct_Time(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			CreateTime time.Time
		}
		now := time.Now()
		user := new(User)
		转换类.Struct(g.Map{
			"create_time": now,
		}, user)
		t.Assert(user.CreateTime.UTC().String(), now.UTC().String())
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			CreateTime *time.Time
		}
		now := time.Now()
		user := new(User)
		转换类.Struct(g.Map{
			"create_time": &now,
		}, user)
		t.Assert(user.CreateTime.UTC().String(), now.UTC().String())
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			CreateTime *时间类.Time
		}
		now := time.Now()
		user := new(User)
		转换类.Struct(g.Map{
			"create_time": &now,
		}, user)
		t.Assert(user.CreateTime.Time.UTC().String(), now.UTC().String())
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			CreateTime 时间类.Time
		}
		now := time.Now()
		user := new(User)
		转换类.Struct(g.Map{
			"create_time": &now,
		}, user)
		t.Assert(user.CreateTime.Time.UTC().String(), now.UTC().String())
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			CreateTime 时间类.Time
		}
		now := time.Now()
		user := new(User)
		转换类.Struct(g.Map{
			"create_time": now,
		}, user)
		t.Assert(user.CreateTime.Time.UTC().String(), now.UTC().String())
	})
}

func Test_Struct_GTime(t *testing.T) {
	// 这是Go语言代码中的一行注释，其内容引用了GitHub上gogf/gf仓库下的第1387号议题。
// 翻译为：
// 参考GitHub上gogf/gf项目中的第1387个问题。
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Name       string
			CreateTime *时间类.Time
		}
		var user *User
		err := 转换类.Struct(`{"Name":"John","CreateTime":""}`, &user)
		t.AssertNil(err)
		t.AssertNE(user, nil)
		t.Assert(user.Name, `John`)
		t.Assert(user.CreateTime, nil)
	})
}

// 当给定指针时自动创建结构体。
func Test_Struct_Create(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Uid  int
			Name string
		}
		user := (*User)(nil)
		params := g.Map{
			"uid":  1,
			"Name": "john",
		}
		err := 转换类.Struct(params, &user)
		t.AssertNil(err)
		t.Assert(user.Uid, 1)
		t.Assert(user.Name, "john")
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Uid  int
			Name string
		}
		user := (*User)(nil)
		params := g.Map{
			"uid":  1,
			"Name": "john",
		}
		err := 转换类.Struct(params, user)
		t.AssertNE(err, nil)
		t.Assert(user, nil)
	})
}

func Test_Struct_Interface(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Uid  interface{}
			Name interface{}
		}
		user := (*User)(nil)
		params := g.Map{
			"uid":  1,
			"Name": nil,
		}
		err := 转换类.Struct(params, &user)
		t.AssertNil(err)
		t.Assert(user.Uid, 1)
		t.Assert(user.Name, nil)
	})
}

func Test_Struct_NilAttribute(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type Item struct {
			Title string `json:"title"`
			Key   string `json:"key"`
		}

		type M struct {
			Id    string                 `json:"id"`
			Me    map[string]interface{} `json:"me"`
			Txt   string                 `json:"txt"`
			Items []*Item                `json:"items"`
		}
		m := new(M)
		err := 转换类.Struct(g.Map{
			"id": "88888",
			"me": g.Map{
				"name": "mikey",
				"day":  "20009",
			},
			"txt":   "hello",
			"items": nil,
		}, m)
		t.AssertNil(err)
		t.AssertNE(m.Me, nil)
		t.Assert(m.Me["day"], "20009")
		t.Assert(m.Items, nil)
	})
}

func Test_Struct_Complex(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type ApplyReportDetail struct {
			ApplyScore        string `json:"apply_score"`
			ApplyCredibility  string `json:"apply_credibility"`
			QueryOrgCount     string `json:"apply_query_org_count"`
			QueryFinanceCount string `json:"apply_query_finance_count"`
			QueryCashCount    string `json:"apply_query_cash_count"`
			QuerySumCount     string `json:"apply_query_sum_count"`
			LatestQueryTime   string `json:"apply_latest_query_time"`
			LatestOneMonth    string `json:"apply_latest_one_month"`
			LatestThreeMonth  string `json:"apply_latest_three_month"`
			LatestSixMonth    string `json:"apply_latest_six_month"`
		}
		type BehaviorReportDetail struct {
			LoansScore         string `json:"behavior_report_detailloans_score"`
			LoansCredibility   string `json:"behavior_report_detailloans_credibility"`
			LoansCount         string `json:"behavior_report_detailloans_count"`
			LoansSettleCount   string `json:"behavior_report_detailloans_settle_count"`
			LoansOverdueCount  string `json:"behavior_report_detailloans_overdue_count"`
			LoansOrgCount      string `json:"behavior_report_detailloans_org_count"`
			ConsfinOrgCount    string `json:"behavior_report_detailconsfin_org_count"`
			LoansCashCount     string `json:"behavior_report_detailloans_cash_count"`
			LatestOneMonth     string `json:"behavior_report_detaillatest_one_month"`
			LatestThreeMonth   string `json:"behavior_report_detaillatest_three_month"`
			LatestSixMonth     string `json:"behavior_report_detaillatest_six_month"`
			HistorySucFee      string `json:"behavior_report_detailhistory_suc_fee"`
			HistoryFailFee     string `json:"behavior_report_detailhistory_fail_fee"`
			LatestOneMonthSuc  string `json:"behavior_report_detaillatest_one_month_suc"`
			LatestOneMonthFail string `json:"behavior_report_detaillatest_one_month_fail"`
			LoansLongTime      string `json:"behavior_report_detailloans_long_time"`
			LoansLatestTime    string `json:"behavior_report_detailloans_latest_time"`
		}
		type CurrentReportDetail struct {
			LoansCreditLimit    string `json:"current_report_detailloans_credit_limit"`
			LoansCredibility    string `json:"current_report_detailloans_credibility"`
			LoansOrgCount       string `json:"current_report_detailloans_org_count"`
			LoansProductCount   string `json:"current_report_detailloans_product_count"`
			LoansMaxLimit       string `json:"current_report_detailloans_max_limit"`
			LoansAvgLimit       string `json:"current_report_detailloans_avg_limit"`
			ConsfinCreditLimit  string `json:"current_report_detailconsfin_credit_limit"`
			ConsfinCredibility  string `json:"current_report_detailconsfin_credibility"`
			ConsfinOrgCount     string `json:"current_report_detailconsfin_org_count"`
			ConsfinProductCount string `json:"current_report_detailconsfin_product_count"`
			ConsfinMaxLimit     string `json:"current_report_detailconsfin_max_limit"`
			ConsfinAvgLimit     string `json:"current_report_detailconsfin_avg_limit"`
		}
		type ResultDetail struct {
			ApplyReportDetail    ApplyReportDetail    `json:"apply_report_detail"`
			BehaviorReportDetail BehaviorReportDetail `json:"behavior_report_detail"`
			CurrentReportDetail  CurrentReportDetail  `json:"current_report_detail"`
		}

		type Data struct {
			Code         string       `json:"code"`
			Desc         string       `json:"desc"`
			TransID      string       `json:"trans_id"`
			TradeNo      string       `json:"trade_no"`
			Fee          string       `json:"fee"`
			IDNo         string       `json:"id_no"`
			IDName       string       `json:"id_name"`
			Versions     string       `json:"versions"`
			ResultDetail ResultDetail `json:"result_detail"`
		}

		type XinYanModel struct {
			Success   bool        `json:"success"`
			Data      Data        `json:"data"`
			ErrorCode interface{} `json:"errorCode"`
			ErrorMsg  interface{} `json:"errorMsg"`
		}

		var data = `{
    "success": true,
    "data": {
        "code": "0",
        "desc": "查询成功",
        "trans_id": "14910304379231213",
        "trade_no": "201704011507240100057329",
        "fee": "Y",
        "id_no": "0783231bcc39f4957e99907e02ae401c",
        "id_name": "dd67a5943781369ddd7c594e231e9e70",
        "versions": "1.0.0",
        "result_detail":{
            "apply_report_detail": {
                "apply_score": "189",
                "apply_credibility": "84",
                "query_org_count": "7",
                "query_finance_count": "2",
                "query_cash_count": "2",
                "query_sum_count": "13",
                "latest_query_time": "2017-09-03",
                "latest_one_month": "1",
                "latest_three_month": "5",
                "latest_six_month": "12"
            },
            "behavior_report_detail": {
                "loans_score": "199",
                "loans_credibility": "90",
                "loans_count": "300",
                "loans_settle_count": "280",
                "loans_overdue_count": "20",
                "loans_org_count": "5",
                "consfin_org_count": "3",
                "loans_cash_count": "2",
                "latest_one_month": "3",
                "latest_three_month": "20",
                "latest_six_month": "23",
                "history_suc_fee": "30",
                "history_fail_fee": "25",
                "latest_one_month_suc": "5",
                "latest_one_month_fail": "20",
                "loans_long_time": "130",
                "loans_latest_time": "2017-09-16"
            },
            "current_report_detail": {
                "loans_credit_limit": "1400",
                "loans_credibility": "80",
                "loans_org_count": "7",
                "loans_product_count": "8",
                "loans_max_limit": "2000",
                "loans_avg_limit": "1000",
                "consfin_credit_limit": "1500",
                "consfin_credibility": "90",
                "consfin_org_count": "8",
                "consfin_product_count": "5",
                "consfin_max_limit": "5000",
                "consfin_avg_limit": "3000"
            }
        }
    },
    "errorCode": null,
    "errorMsg": null
}`
		m := make(g.Map)
		err := json.UnmarshalUseNumber([]byte(data), &m)
		t.AssertNil(err)

		model := new(XinYanModel)
		err = 转换类.Struct(m, model)
		t.AssertNil(err)
		t.Assert(model.ErrorCode, nil)
		t.Assert(model.ErrorMsg, nil)
		t.Assert(model.Success, true)
		t.Assert(model.Data.IDName, "dd67a5943781369ddd7c594e231e9e70")
		t.Assert(model.Data.TradeNo, "201704011507240100057329")
		t.Assert(model.Data.ResultDetail.ApplyReportDetail.ApplyScore, "189")
		t.Assert(model.Data.ResultDetail.BehaviorReportDetail.LoansSettleCount, "280")
		t.Assert(model.Data.ResultDetail.CurrentReportDetail.LoansProductCount, "8")
	})
}

func Test_Struct_CatchPanic(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type Score struct {
			Name   string
			Result int
		}
		type User struct {
			Score Score
		}

		user := new(User)
		scores := map[string]interface{}{
			"Score": 1,
		}
		err := 转换类.Struct(scores, user)
		t.AssertNE(err, nil)
	})
}

type T struct {
	Name string
}

func (t *T) Test() string {
	return t.Name
}

type TestInterface interface {
	Test() string
}

type TestStruct struct {
	TestInterface
}

func Test_Struct_Embedded(t *testing.T) {
	// 实现了接口属性。
	单元测试类.C(t, func(t *单元测试类.T) {
		v1 := TestStruct{
			TestInterface: &T{"john"},
		}
		v2 := g.Map{}
		err := 转换类.Struct(v2, &v1)
		t.AssertNil(err)
		t.Assert(v1.Test(), "john")
	})
	// 实现了接口属性。
	单元测试类.C(t, func(t *单元测试类.T) {
		v1 := TestStruct{
			TestInterface: &T{"john"},
		}
		v2 := g.Map{
			"name": "test",
		}
		err := 转换类.Struct(v2, &v1)
		t.AssertNil(err)
		t.Assert(v1.Test(), "test")
	})
	// 未实现接口属性。
	单元测试类.C(t, func(t *单元测试类.T) {
		v1 := TestStruct{}
		v2 := g.Map{
			"name": "test",
		}
		err := 转换类.Struct(v2, &v1)
		t.AssertNil(err)
		t.Assert(v1.TestInterface, nil)
	})
}

func Test_Struct_Slice(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Scores []int
		}
		user := new(User)
		array := g.Slice别名{1, 2, 3}
		err := 转换类.Struct(g.Map{"scores": array}, user)
		t.AssertNil(err)
		t.Assert(user.Scores, array)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Scores []int32
		}
		user := new(User)
		array := g.Slice别名{1, 2, 3}
		err := 转换类.Struct(g.Map{"scores": array}, user)
		t.AssertNil(err)
		t.Assert(user.Scores, array)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Scores []int64
		}
		user := new(User)
		array := g.Slice别名{1, 2, 3}
		err := 转换类.Struct(g.Map{"scores": array}, user)
		t.AssertNil(err)
		t.Assert(user.Scores, array)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Scores []uint
		}
		user := new(User)
		array := g.Slice别名{1, 2, 3}
		err := 转换类.Struct(g.Map{"scores": array}, user)
		t.AssertNil(err)
		t.Assert(user.Scores, array)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Scores []uint32
		}
		user := new(User)
		array := g.Slice别名{1, 2, 3}
		err := 转换类.Struct(g.Map{"scores": array}, user)
		t.AssertNil(err)
		t.Assert(user.Scores, array)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Scores []uint64
		}
		user := new(User)
		array := g.Slice别名{1, 2, 3}
		err := 转换类.Struct(g.Map{"scores": array}, user)
		t.AssertNil(err)
		t.Assert(user.Scores, array)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Scores []float32
		}
		user := new(User)
		array := g.Slice别名{1, 2, 3}
		err := 转换类.Struct(g.Map{"scores": array}, user)
		t.AssertNil(err)
		t.Assert(user.Scores, array)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Scores []float64
		}
		user := new(User)
		array := g.Slice别名{1, 2, 3}
		err := 转换类.Struct(g.Map{"scores": array}, user)
		t.AssertNil(err)
		t.Assert(user.Scores, array)
	})
}

func Test_Struct_To_Struct(t *testing.T) {
	var TestA struct {
		Id   int       `p:"id"`
		Date time.Time `p:"date"`
	}

	var TestB struct {
		Id   int       `p:"id"`
		Date time.Time `p:"date"`
	}
	TestB.Id = 666
	TestB.Date = time.Now()

	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(转换类.Struct(TestB, &TestA), nil)
		t.Assert(TestA.Id, TestB.Id)
		t.Assert(TestA.Date, TestB.Date)
	})
}

func Test_Struct_WithJson(t *testing.T) {
	type A struct {
		Name string
	}
	type B struct {
		A
		Score int
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		b1 := &B{}
		b1.Name = "john"
		b1.Score = 100
		b, _ := json.Marshal(b1)
		b2 := &B{}
		err := 转换类.Struct(b, b2)
		t.AssertNil(err)
		t.Assert(b2, b1)
	})
}

func Test_Struct_AttrStructHasTheSameTag(t *testing.T) {
	type Product struct {
		Id              int       `json:"id"`
		UpdatedAt       time.Time `json:"-" `
		UpdatedAtFormat string    `json:"updated_at" `
	}

	type Order struct {
		Id        int       `json:"id"`
		UpdatedAt time.Time `json:"updated_at"`
		Product   Product   `json:"products"`
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		data := g.Map{
			"id":         1,
			"updated_at": time.Now(),
		}
		order := new(Order)
		err := 转换类.Struct(data, order)
		t.AssertNil(err)
		t.Assert(order.Id, data["id"])
		t.Assert(order.UpdatedAt, data["updated_at"])
		t.Assert(order.Product.Id, 0)
		t.Assert(order.Product.UpdatedAt.IsZero(), true)
		t.Assert(order.Product.UpdatedAtFormat, "")
	})
}

func Test_Struct_DirectReflectSet(t *testing.T) {
	type A struct {
		Id   int
		Name string
	}

	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			a = &A{
				Id:   1,
				Name: "john",
			}
			b *A
		)
		err := 转换类.Struct(a, &b)
		t.AssertNil(err)
		t.AssertEQ(a, b)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			a = A{
				Id:   1,
				Name: "john",
			}
			b A
		)
		err := 转换类.Struct(a, &b)
		t.AssertNil(err)
		t.AssertEQ(a, b)
	})
}

func Test_Struct_NilEmbeddedStructAttribute(t *testing.T) {
	type A struct {
		Name string
	}
	type B struct {
		*A
		Id int
	}

	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			b *B
		)
		err := 转换类.Struct(g.Map{
			"id":   1,
			"name": nil,
		}, &b)
		t.AssertNil(err)
		t.Assert(b.Id, 1)
		t.Assert(b.Name, "")
	})
}

func Test_Struct_JsonParam(t *testing.T) {
	type A struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
	// struct
	单元测试类.C(t, func(t *单元测试类.T) {
		var a = A{}
		err := 转换类.Struct([]byte(`{"id":1,"name":"john"}`), &a)
		t.AssertNil(err)
		t.Assert(a.Id, 1)
		t.Assert(a.Name, "john")
	})
	// *struct
	单元测试类.C(t, func(t *单元测试类.T) {
		var a = &A{}
		err := 转换类.Struct([]byte(`{"id":1,"name":"john"}`), a)
		t.AssertNil(err)
		t.Assert(a.Id, 1)
		t.Assert(a.Name, "john")
	})
	// *struct nil
	单元测试类.C(t, func(t *单元测试类.T) {
		var a *A
		err := 转换类.Struct([]byte(`{"id":1,"name":"john"}`), &a)
		t.AssertNil(err)
		t.Assert(a.Id, 1)
		t.Assert(a.Name, "john")
	})
}

func Test_Struct_GVarAttribute(t *testing.T) {
	type A struct {
		Id     int    `json:"id"`
		Name   string `json:"name"`
		Status bool   `json:"status"`
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			a    = A{}
			data = g.Map{
				"id":     100,
				"name":   "john",
				"status": 泛型类.X创建(false),
			}
		)
		err := 转换类.Struct(data, &a)
		t.AssertNil(err)
		t.Assert(a.Id, data["id"])
		t.Assert(a.Name, data["name"])
		t.Assert(a.Status, data["status"])
	})

}

func Test_Struct_MapAttribute(t *testing.T) {
	type NodeStatus struct {
		ID int
	}
	type Nodes map[string]NodeStatus
	type Output struct {
		Nodes Nodes
	}

	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			out  = Output{}
			data = g.Map{
				"nodes": g.Map{
					"name": g.Map{
						"id": 10000,
					},
				},
			}
		)
		err := 转换类.Struct(data, &out)
		t.AssertNil(err)
	})
}

func Test_Struct_Empty_MapStringString(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type S struct {
			Id int
		}
		var (
			s   = &S{}
			err = 转换类.Struct(map[string]string{}, s)
		)
		t.AssertNil(err)
	})
}

// 这是Go语言代码中的一行注释，其内容引用了GitHub上gogf/gf项目的一个问题编号1563。
// 中文翻译：
// 参考GitHub上gogf/gf项目的第1563号问题。
func Test_Struct_Issue1563(t *testing.T) {
	type User struct {
		Pass1 string `c:"password1"`
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		for i := 0; i < 100; i++ {
			user := new(User)
			params2 := g.Map{
				"password1": "111",
				// "PASS1":     "222",
// 这是一行Go语言中的字符串字面量，用作键值对的形式出现，通常用于初始化映射（map）或结构体等数据结构。
// 注释翻译为中文：
// `"PASS1"`:   `"222"`,
// 这里定义了一个键为"PASS1"，值为"222"的键值对。
				"Pass1": "333",
			}
			if err := 转换类.Struct(params2, user); err == nil {
				t.Assert(user.Pass1, `111`)
			}
		}
	})
}

// 这是GitHub上gogf/gf仓库的第1597号问题
func Test_Struct_Issue1597(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type S struct {
			A int
			B json.RawMessage
		}

		jsonByte := []byte(`{
		"a":1, 
		"b":{
			"c": 3
		}
	}`)
		data, err := json类.X解码到json(jsonByte)
		t.AssertNil(err)
		s := &S{}
		err = data.X取结构体指针(s)
		t.AssertNil(err)
		t.Assert(s.B, `{"c":3}`)
	})
}

// 这是GitHub上gogf/gf仓库的第2980个issue的链接
func Test_Struct_Issue2980(t *testing.T) {
	type Post struct {
		CreatedAt *时间类.Time `json:"createdAt" `
	}

	type PostWithUser struct {
		Post
		UserName string `json:"UserName"`
	}

	单元测试类.C(t, func(t *单元测试类.T) {
		date := 时间类.X创建("2023-09-22 12:00:00").X取UTC时区()
		params := g.Map{
			"CreatedAt": 时间类.X创建("2023-09-22 12:00:00").X取UTC时区(),
			"UserName":  "Galileo",
		}
		postWithUser := new(PostWithUser)
		err := 转换类.Scan(params, postWithUser)
		t.AssertNil(err)
		t.Assert(date.Location(), postWithUser.CreatedAt.Location())
		t.Assert(date.Unix(), postWithUser.CreatedAt.Unix())
	})
}

func Test_Scan_WithDoubleSliceAttribute(t *testing.T) {
	inputData := [][]string{
		{"aa", "bb", "cc"},
		{"11", "22", "33"},
	}
	data := struct {
		Data [][]string
	}{
		Data: inputData,
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		jv := json类.X创建(json类.X变量到json文本PANI(data))
		err := jv.X取结构体指针(&data)
		t.AssertNil(err)
		t.Assert(data.Data, inputData)
	})

}

func Test_Struct_WithCustomType(t *testing.T) {
	type PayMode int

	type Req1 struct {
		PayMode PayMode
	}
	type Req2 struct {
		PayMode *PayMode
	}
	var (
		params = 转换类.X取Map(`{"PayMode": 1000}`)
		req1   *Req1
		req2   *Req2
		err1   error
		err2   error
	)
	err1 = 转换类.Struct(params, &req1)
	err2 = 转换类.Struct(params, &req2)
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertNil(err1)
		t.Assert(req1.PayMode, 1000)

		t.AssertNil(err2)
		t.AssertNE(req2.PayMode, nil)
		t.Assert(*req2.PayMode, 1000)
	})
}

func Test_Struct_EmptyStruct(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var err error

		type StructA struct {
		}

		type StructB struct {
		}

		var s1 StructA
		var s2 *StructB

		err = 转换类.Scan(s1, &s2)
		t.AssertNil(err)

		err = 转换类.Scan(&s1, &s2)
		t.AssertNil(err)

		type StructC struct {
			Val int `json:"val,omitempty"`
		}

		type StructD struct {
			Val int
		}

		var s3 StructC
		var s4 *StructD

		err = 转换类.Scan(s3, &s4)
		t.AssertNil(err)
		t.Assert(s4.Val, 0)

		err = 转换类.Scan(&s3, &s4)
		t.AssertNil(err)
		t.Assert(s4.Val, 0)

		s3.Val = 123
		err = 转换类.Scan(s3, &s4)
		t.AssertNil(err)
		t.Assert(s4.Val, 123)

		err = 转换类.Scan(&s3, &s4)
		t.AssertNil(err)
		t.Assert(s4.Val, 123)

	})
}
