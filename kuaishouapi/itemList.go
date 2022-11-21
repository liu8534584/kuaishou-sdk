package kuaishouapi

type ItemListRequest struct {
	RangeList   []ItemListRequestRange `json:"rangeList,omitempty"`   //List<RangeList> 否 [] 范围筛选列表
	SortType    string                 `json:"sortType,omitempty"`    // 否  DEFAULT_SORT 排序规则：DEFAULT_SORT-默认，RATE_ASC-佣金比率升序，RATE_DESC-佣金比率降序，PRICE_ASC-价格升序，PRICE_DESC-价格降序，VOLUME_ASC-30天销量升序，VOLUME_DESC-30天销量降序
	PageIndex   string                 `json:"pageIndex,omitempty"`   // 是 string 分页游标，第一次请求不需传参，后续请求使用上次请求返回的值
	ChannelId   []int                  `json:"channelId,omitempty"`   // 否  [123,456] 频道id列表
	PageSize    int                    `json:"pageSize,omitempty"`    // 是  20  分页大小，最大不超过200
	ExpressType int                    `json:"expressType,omitempty"` //否  1  0或者不传：无筛选，1：不包邮，2:包邮
	PlanType    int                    `json:"planType,omitempty"`    //是  1 分销计划类型。1-普通计划，2-商品定向4-店铺定向，5-专属计划，6-普通招商，7-专属招商
	Keyword     string                 `json:"keyword,omitempty"`     // 否 关键词 搜索关键词
	ItemLevel   string                 `json:"itemLevel,omitempty"`   // 否L1 商品分层等级。L1-新品上线,L2-达标分销商品,L3-优质好物,L4-行业爆款,L5-行业尖货
	SellerId    uint64                 `json:"sellerId,omitempty"`    //否 123456 卖家ID
}

type ItemListRequestRange struct {
	RangeId   string `json:"rangeId,omitempty"`   // 是 PRICE 筛选条件id：PRICE-商品价格，PROMOTION_RATE-佣金比率，THIRTY_DAYS_VOLUME-30天销量
	RangeFrom int64  `json:"rangeFrom,omitempty"` // 是 0 区间起始值
	RangeTo   int64  `json:"rangeTo,omitempty"`   // 是 1000 区间结束值
}

type ItemListResponse struct {
	Code     string               `json:"code"`      // 1 主返回码
	Msg      string               `json:"msg"`       // success 主返回信息
	SubCode  string               `json:"sub_code"`  //1 子返回码
	SubMsg   string               `json:"sub_msg"`   // SUCCESS 子返回信息
	Result   int                  `json:"result"`    // 1 返回码
	ErrorMsg string               `json:"error_msg"` // SUCCESS 错误信息
	Data     ItemListResponseData `json:"data"`      //返回数据
}

type ItemListResponseData struct {
	PageIndex string         `json:"pageIndex"` //1 分页游标，用于下一次请求获取选品列表
	ItemList  []KuaishouItem `json:"itemList"`  //list  商品列表
}
