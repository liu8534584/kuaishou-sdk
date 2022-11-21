package kuaishouapi

type LinkCreateRequest struct {
	LinkType      int    `json:"linkType"`      // 是 100 链接类型，100-直播间链接，101-商品链接
	LinkCarrierId string `json:"linkCarrierId"` // 是 2153705985 携带id（达人快手id或商品id）
	Comments      string `json:"comments"`      // 是 备注 备注（透传字段，拉取订单接口返回）
	CpsPid        string `json:"cpsPid"`        // 是 ks_2153646980_102_UsPG8hEvtaE 快赚客推广位cps pid
	GenPoster     bool   `json:"genPoster"`     // 否 false 是否返回海报信息（默认值false，不返回海报信息）
	CustomContent string `json:"customContent"` //否 "" 自定义文案
}

type LinkCreateResponse struct {
	Code     string                 `json:"code"`      // 1 主返回码
	Msg      string                 `json:"msg"`       // success 主返回信息
	SubCode  string                 `json:"sub_code"`  //1 子返回码
	SubMsg   string                 `json:"sub_msg"`   // SUCCESS 子返回信息
	Result   int                    `json:"result"`    // 1 返回码
	ErrorMsg string                 `json:"error_msg"` // SUCCESS 错误信息
	Data     LinkCreateResponseData `json:"data"`      //list  商品列表
}

type LinkCreateResponseData struct {
	LinkCode       string `json:"linkCode"`       // ZHD01的直播很精彩，快来围观！http://v-short.staging.kuaishou.com/8poTHS 复制此消息，打开【快手】直接观看！  口令
	LinkUrl        string `json:"linkUrl"`        // http://v-short.staging.kuaishou.com/8poTHS  转链链接
	KwaiUrl        string `json:"kwaiUrl"`        // kwai://~ kwai url
	CpsPid         string `json:"cpsPid"`         // ks_2153646980_102_UsPG8hEvtaE  快赚客推广位cps pid
	NebulaKwaiUrl  string `json:"nebulaKwaiUrl"`  //ksnebula://~  极速版kwai url
	ShortContent   string `json:"shortContent"`   // coco自动化商品\n【在售价】1.00 元\n-----------------\n【立即下单】点击链接立即下单：https://ksfx83zns.8oxye9xqef.com/f/X-4TKRW6iY40f1mc_A 短链文案
	LongContent    string `json:"longContent"`    // coco自动化商品\n【在售价】1.00 元\n-----------------\n【立即下单】点击链接立即下单：https://ksfx83zns.8oxye9xqef.com/merchant/shop/detail?id=2146600649592&cc=share_copylink&shareMode=app&shareMethod=token&kpn=KUAISHOU&subBiz=KS_FENXIAO&shareToken=X-4TKRW6iY40f1mc_A&fid=2337072173  长链文案
	CommandContent string `json:"commandContent"` //##X-4TKRW6iY40f1mc_A## 复制此消息 ，打开【快手】直接购买！\n----------------\ncoco自动化商品\n【在售价】1.00 元 口令文案
	ShareToken     string `json:"shareToken"`     // X-4TKRW6iY40f1mc_A 分享token
	PosterContent  struct {
		MinaCode           string `json:"minaCode"`           // https://ali-ec.static.yximgs.com/bs2/upload-distribute-kwaimoney/3f9bb8ece6054920a34c9af12fa15428.png  小程序码
		ItemImg            string `json:"itemImg"`            // https://ali-ec.static.yximgs.com/ufile/adsocial/7b5a2132-4ce5-4ee5-94fb-79b2c6f7f770.jpg  商品图片
		ItemTitle          string `json:"itemTitle"`          // coco自动化商品 商品标题
		ItemPrice          int64  `json:"itemPrice"`          // 100 商品价格（单位分）
		PromoterHeadImgUrl string `json:"promoterHeadImgUrl"` //http://  达人头像
		PromoterNickname   string `json:"promoterNickname"`   // 昵称 达人昵称
	} `json:"posterContent"` // 海报信息（目前仅限商品推广）
	MinaJumpContent struct {
		PageUrl string `json:"pageUrl"` // /package-live/pages/live/live?principalId=2172034724&kpn=KUAISHOU&subBiz=KS_ZHIBOJIANFENXIAO&shareToken=XXtukLIrzBOv63&fid=2167544025  小程序页面路径
		AppId   string `json:"appId"`   // wxddc app id
	} `json:"minaJumpContent"` //跳小程序物料
	MinaShortLink string `json:"minaShortLink"` // #小程序://快手短视频/商品详情页/pbf8NeioceC82An  小程序短链
}
