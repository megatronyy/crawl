package rule

import (
	"strings"
	"github.com/henrylee2cn/pholcus/common/goquery"
)

/**获取标题**/
func gettitle(query *goquery.Document) string {
	标题 := query.Find("div.main-wrap>div.house-title>h1").Text()
	return trim(标题)
}

/**获取价格**/
func getprice(query *goquery.Document) string {
	价格 := query.Find("div.house-basic-right>p.house_basic_title_money>span.house_basic_title_money_num").Text()
	return trim(价格)
}

/**获取价格单位**/
func getuint(query *goquery.Document) string {
	单位 := query.Find("div.house-basic-right>p.house_basic_title_money>span.house_basic_title_money_unit").Text()
	return trim(单位)
}

/**获取出租价格**/
func getczprice(query *goquery.Document) string {
	出租价格 := query.Find("div.house-basic-right>p.house_basic_title_money>span.house_basic_title_money_num_chuzu").Text()
	return trim(出租价格)
}

/**获取出租单位**/
func getczuint(query *goquery.Document) string {
	出租单位 := query.Find("div.house-basic-right>p.house_basic_title_money>span.house_basic_title_money_unit_chuzu").Text()
	return trim(出租单位)
}

/**获取发布人**/
func publisher(query *goquery.Document) string {
	发布人, _ := query.Find("div.house-basic-right>div.house_basic_jingji>p.nav>a>img").Attr("alt")
	return trim(发布人)
}

/**获取发布时间**/
func publishtime(query *goquery.Document) string {
	发布时间 := trim(query.Find("div.main-wrap>div.house-title>p.house-update-info>span").Eq(0).Text())
	发布时间 = strings.Replace(发布时间, "更新于", "", -1)
	发布时间 = strings.Replace(发布时间, "创建于", "", -1)
	return trim(发布时间)
}

/**获取发布人电话及归属地**/
func publisherphone(query *goquery.Document) (string, string) {
	联系电话 := query.Find("#houseChatEntry>div.house-chat-phone>p.phone-num").Text()
	电话归属 := query.Find("#houseChatEntry>div.house-chat-phone>p.phone-belong>span").Eq(1).Text()
	return trim(联系电话), trim(电话归属)
}

/**获取详细描述**/
func getdetails(query *goquery.Document) string {
	详情, _ := query.Find("div#generalSound>div.general-item-wrap").Html()
	return 详情
}
/**获取其它**/
func getother(query *goquery.Document) (string, string, string, string, string, string, string) {
	rs := query.Find("div.house-basic-right>ul.house_basic_title_content>li")
	面积 := trim(rs.Eq(0).Find(".house_basic_title_content_item2").Text())
	类型 := trim(rs.Eq(0).Find("a.house_basic_title_content_item3").Text()) + trim(rs.Eq(1).Find("span.house_basic_title_content_item3").Text())
	经营状态 := trim(rs.Eq(1).Find("span.house_basic_title_content_item3").Text())
	历史经营 := trim(rs.Eq(2).Find("span.house_basic_title_content_item3").Text())
	付款方式 := trim(rs.Eq(3).Find("span.house_basic_title_content_item3").Text())
	租约方式 := trim(rs.Eq(4).Find("span.house_basic_title_content_item3").Text())
	详细地址 := getaddress(rs.Eq(5))

	return 面积, 类型, 经营状态, 历史经营, 付款方式, 租约方式, 详细地址
}
/**获取详细地址**/
func getaddress(s *goquery.Selection) string {
	var a string
	s.Find(".house_basic_title_content_item3").Each(func(i int, selection *goquery.Selection) {
		if i==0{
			a += trim(selection.Text()) + " "
		}else{
			a += trim(selection.Text())
		}

	})
	return a
}

/**去掉空格换行符**/
func trim(str string) string {
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, "\t", "", -1)
	return str
}
