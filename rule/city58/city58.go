package rule

import (
	. "github.com/henrylee2cn/pholcus/app/spider"
	"github.com/henrylee2cn/pholcus/app/downloader/request"
	"strconv"
	"github.com/henrylee2cn/pholcus/common/goquery"
	"strings"
)

func init() {
	CitySpcz.Register()
}

var CitySpcz = &Spider{
	Name:         "business",
	Description:  "58同城商铺出租的信息",
	Pausetime:    5000,
	EnableCookie: true,
	RuleTree: &RuleTree{
		Root: func(ctx *Context) {
			ctx.AddQueue(&request.Request{
				Url:        "http://bj.58.com/shangpucz/pn1/",
				Rule:       "house_page",
				Temp:       map[string]interface{}{"p": 1},
				Reloadable: true,
			})
		},

		Trunk: map[string]*Rule{
			"house_page": {
				ParseFunc: func(ctx *Context) {
					var curr = ctx.GetTemp("p", 0).(int)
					//if c := ctx.GetDom().Find("div.content-side-left>div.pager>strong>span").Text(); c != strconv.Itoa(curr) {
					//	return
					//}

					//只抓取前五页
					if curr >= 5 {
						curr = 0
					}

					ctx.AddQueue(&request.Request{
						Url:        "http://bj.58.com/shangpucz/pn" + strconv.Itoa(curr+1) + "/",
						Rule:       "house_page",
						Temp:       map[string]interface{}{"p": curr + 1},
						Reloadable: true,
					})

					//用指定规则解析响应流
					ctx.Parse("house_list")
				},
			},

			"house_list": {
				ParseFunc: func(ctx *Context) {
					li := ctx.GetDom().Find("div.content-side-left>ul.house-list-wrap>li")
					li.Each(func(i int, s *goquery.Selection) {
						a := s.Find("div.pic>a")
						url, _ := a.Attr("href")
						if url != "" && strings.Contains(url, "x.shtml") {
							ctx.AddQueue(&request.Request{
								Url:      url,
								Rule:     "house_detail",
								Priority: 1,
							})
						}
					})
				},
			},

			"house_detail": {
				ItemFields: []string{
					"标题",
					"价格",
					"单位",
					"出租价格",
					"出租单位",
					"面积",
					"类型",
					"经营状态",
					"历史经营",
					"付款方式",
					"租约方式",
					"详细地址",
					"发布时间",
					"发布人",
					"联系电话",
					"详情",
				},

				ParseFunc: city58_parse,
			},
		},
	},
}
