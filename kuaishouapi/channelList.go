package kuaishouapi

type ChannelListResponse struct {
	Code     string                       `json:"code"`      // 1 主返回码
	Msg      string                       `json:"msg"`       // success 主返回信息
	SubCode  string                       `json:"sub_code"`  //1 子返回码
	SubMsg   string                       `json:"sub_msg"`   // SUCCESS 子返回信息
	Result   int                          `json:"result"`    // 1 返回码
	ErrorMsg string                       `json:"error_msg"` // SUCCESS 错误信息
	Data     []ChannelListResponseChannel `json:"data"`      //list  商品列表
}

type ChannelListResponseChannel struct {
	ChannelId   int64  // 1001  频道ID
	ChannelName string //频道  频道名称
}
