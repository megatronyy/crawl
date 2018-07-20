package rule

import (
	. "github.com/henrylee2cn/pholcus/app/spider"
	"github.com/henrylee2cn/pholcus/app/downloader/request"
	"strconv"
	"github.com/henrylee2cn/pholcus/common/goquery"
)

func init() {
	CitySpcz.Register()
}

var CitySpcz = &Spider{
	Name:         "58同城",
	Description:  "58同城商铺出租的信息",
	Pausetime:    3000,
	EnableCookie: false,
	RuleTree: &RuleTree{
		Root: func(ctx *Context) {
			ctx.AddQueue(&request.Request{
				Url:  "http://bj.58.com/shangpucz/pn1/",
				Rule: "请求列表",
				Temp: map[string]interface{}{"p": 1},
			})
		},

		Trunk: map[string]*Rule{
			"请求列表": {
				ParseFunc: func(ctx *Context) {
					var curr = ctx.GetTemp("p", 0).(int)
					if c := ctx.GetDom().Find(".content-side-left .pager strong span").Text(); c != strconv.Itoa(curr) {
						return
					}

					ctx.AddQueue(&request.Request{
						Url:  "http://bj.58.com/shangpucz/pn" + strconv.Itoa(curr+1) + "/",
						Rule: "请求列表",
						Temp: map[string]interface{}{"p": curr + 1},
					})

					//用指定规则解析响应流
					ctx.Parse("获取列表")
				},
			},

			"获取列表": {
				ParseFunc: func(ctx *Context) {
					ctx.GetDom().Find(".content-side-left .house-list-wrap li").Each(func(i int, s *goquery.Selection) {
						url, _ := s.Find(".pc a").Attr("href")
						ctx.AddQueue(&request.Request{
							Url:      url,
							Rule:     "保存结果",
							Priority: 1,
						})
					})
				},
			},

			"保存结果": {
				ItemFields: []string{
					"标题",
					"价格",
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
				},

				ParseFunc: func(ctx *Context) {
					query := ctx.GetDom()

					var 标题, 价格, 单位, 出租价格, 出租单位, 面积 string

					//, 类型, 经营状态, 历史经营, 付款方式, 租约方式, 详细地址, 发布时间, 发布人, 联系电话

					标题 = query.Find(".house-title .f20").Text()
					价格 = query.Find(".house_basic_title_money_num").Text()
					单位 = query.Find(".house_basic_title_money_unit").Text()
					出租价格 = query.Find(".house_basic_title_money_num_chuzu").Text()
					出租单位 = query.Find(".house_basic_title_money_unit_chuzu").Text()
					面积 = query.Find(".house-basic-right .house_basic_title_content").Eq(1).
						Find(".house_basic_title_content_item2").Text()

					ctx.Output(map[int]interface{}{
						0: 标题,
						1: 价格,
						2: 单位,
						3: 出租价格,
						4: 出租单位,
						5: 面积,
					})
				},
			},
		},
	},
}