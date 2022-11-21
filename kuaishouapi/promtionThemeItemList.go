package kuaishouapi

type PromtionThemeItemListRequest struct {
	ThemeId    int    `json:"themeId"`    //是 34143  专题ID
	SubThemeId int    `json:"subThemeId"` // 否 2435 子主题ID
	PCursor    string `json:"pcursor"`    // 否 1 游标
}

type PromtionThemeItemListResponse struct {
	Code     string         `json:"code"`      // 1 主返回码
	Msg      string         `json:"msg"`       // success 主返回信息
	SubCode  string         `json:"sub_code"`  //1 子返回码
	SubMsg   string         `json:"sub_msg"`   // SUCCESS 子返回信息
	Result   int            `json:"result"`    // 1 返回码
	ErrorMsg string         `json:"error_msg"` // SUCCESS 错误信息
	Data     []KuaishouItem `json:"data"`      //list  商品列表
	Pcursor  string         `json:"pcursor"`   // 1 游标
}
