package rule

import (
	. "github.com/henrylee2cn/pholcus/app/spider"
)

func city58_cz_parse(ctx *Context) {
	query := ctx.GetDom()

	var 标题, 价格, 单位, 出租价格, 出租单位, 面积, 类型, 经营状态, 历史经营, 付款方式, 租约方式, 详细地址, 发布时间, 发布人, 联系电话, 电话归属, 详情 string

	标题 = gettitle(query)
	价格 = getprice(query)
	单位 = getuint(query)
	出租价格 = getczprice(query)
	出租单位 = getczuint(query)
	面积, 类型, 经营状态, 历史经营, 付款方式, 租约方式, 详细地址 = getother(query)
	发布时间 = publishtime(query)
	发布人 = publisher(query)
	联系电话, 电话归属 = publisherphone(query)
	详情 = getdetails(query)

	if 标题 != ""{
		ctx.Output(map[int]interface{}{
			0:  标题,
			1:  价格,
			2:  单位,
			3:  出租价格,
			4:  出租单位,
			5:  面积,
			6:  类型,
			7:  经营状态,
			8:  历史经营,
			9:  付款方式,
			10: 租约方式,
			11: 详细地址,
			12: 发布时间,
			13: 发布人,
			14: 联系电话,
			15: 电话归属,
			16: 详情,
		})
	}
}
