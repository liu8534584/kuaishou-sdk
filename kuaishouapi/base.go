package kuaishouapi

const MethodItemList = "open.distribution.cps.kwaimoney.selection.item.list"
const MethodItemDetail = "open.distribution.cps.kwaimoney.selection.item.detail"
const MethodLinkCreate = "open.distribution.cps.kwaimoney.link.create"
const MethodChannelList = "open.distribution.cps.kwaimoney.selection.channel.list"
const MethodPromtionThemeItemList = "open.distribution.cps.promtion.theme.item.list"

type KuaishouItem struct {
	ChannelIds       []int    `json:"channelIds"`       //[123,456] 频道id列表
	GoodsId          int64    `json:"goodsId"`          // 1001 商品id
	GoodsPrice       int64    `json:"goodsPrice"`       // 100 商品价格（分）
	ZkGoodsPrice     int64    `json:"zkGoodsPrice"`     //98  折扣价
	GoodsTitle       string   `json:"goodsTitle"`       //标题 商品标题
	GoodsDesc        string   `json:"goodsDesc"`        // 9 商品描述
	MallId           int64    `json:"mallId"`           // 1001 卖家id
	MallName         string   `json:"mallName"`         // 店铺名称 店铺名称
	MallType         int      `json:"mallType"`         // 1 店铺类型（列表接口下无数据）：0-未知，1-旗舰店，2-专门店，3-专营店，4-卖场旗舰店，5-普通企业店，6-个人店，7-个体工商户
	GoodsImageUrl    string   `json:"goodsImageUrl"`    // https://xx.xx/xx.jpg  商品图片链接
	CategoryId       int64    `json:"categoryId"`       // 1001 类目id
	PromotionRate    int      `json:"promotionRate"`    // 100 佣金率（千分比）
	SalesTip         int64    `json:"salesTip"`         // 100 30销量
	GoodsGalleryUrls []string `json:"goodsGalleryUrls"` // [] 轮播图列表，仅在获取详情接口时返回
	PromotionAmount  int64    `json:"promotionAmount"`  //1  推广预估佣金（分）
	ItemDescUrls     []string `json:"itemDescUrls"`     // [] 商品详情图，仅在获取详情接口时返回
	ExpressId        int64    `json:"expressId"`        // 2  运费模版id，-1时为全国包邮模版，0时表示状态不正确。其他值可以使用开放平台获取运费模版接口获取具体运费信息
}
