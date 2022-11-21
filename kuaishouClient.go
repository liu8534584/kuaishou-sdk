package kuaishou

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/liu8534584/kuaishou-sdk/kuaishouapi"
	"github.com/parnurzeal/gorequest"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type KuaishouClient struct {
	appKey     string
	signSecret string
	gatewayUrl string
}

func NewKuaishouClient(appKey, signSecret string) *KuaishouClient {
	return &KuaishouClient{
		appKey:     appKey,
		signSecret: signSecret,
		gatewayUrl: "https://openapi.kwaixiaodian.com/",
	}
}

func (c *KuaishouClient) execute(method, accessToken string, params interface{}, version int, isPost bool) (string, error) {

	request := gorequest.New()
	request.Timeout(3 * time.Second)
	var sysParams = make(map[string]string)
	sysParams["appkey"] = c.appKey
	timestamp := time.Now().Unix() * 1000
	sysParams["timestamp"] = strconv.FormatInt(timestamp, 10)
	sysParams["access_token"] = accessToken
	sysParams["version"] = strconv.Itoa(version)
	paramJson, sign := c.sign(accessToken, method, timestamp, version, params)
	sysParams["sign"] = sign
	sysParams["method"] = method
	sysParams["signMethod"] = "MD5"

	apiUrl := c.gatewayUrl + strings.ReplaceAll(method, ".", "/")
	var errs []error
	var body string
	if isPost {
		sysParams["param"] = paramJson
		_, body, errs = request.Send(sysParams).Post(apiUrl).End()

	} else {
		sysParams["param"] = url.QueryEscape(paramJson)
		_, body, errs = request.Query(sysParams).Get(apiUrl).End()
	}
	if len(errs) > 0 {
		return body, errs[0]
	}
	return body, nil
}

func (c *KuaishouClient) ItemList(ctx context.Context, request kuaishouapi.ItemListRequest, accessToken string) (kuaishouapi.ItemListResponseData, error) {

	var res kuaishouapi.ItemListResponse
	httpExecute, httpErr := c.execute(kuaishouapi.MethodItemList, accessToken, request, 1, false)
	if httpErr != nil {
		return res.Data, httpErr
	}

	jsonErr := json.Unmarshal([]byte(httpExecute), &res)
	if jsonErr != nil {
		return res.Data, jsonErr
	}

	if res.Result != 1 {
		return res.Data, NewApiError(res.Result, fmt.Sprintf("message:%s,code:%s,subMessage:%s,subCode:%s,err:%v", res.Msg, res.Code, res.SubMsg, res.SubCode, res))
	}

	return res.Data, nil
}

func (c *KuaishouClient) ItemDetail(ctx context.Context, request kuaishouapi.ItemDetailRequest, accessToken string) ([]kuaishouapi.KuaishouItem, error) {

	var res kuaishouapi.ItemDetailResponse
	httpExecute, httpErr := c.execute(kuaishouapi.MethodItemDetail, accessToken, request, 1, false)
	if httpErr != nil {
		return res.Data, httpErr
	}

	jsonErr := json.Unmarshal([]byte(httpExecute), &res)
	if jsonErr != nil {
		return res.Data, jsonErr
	}

	if res.Result != 1 {
		return res.Data, NewApiError(res.Result, fmt.Sprintf("message:%s,code:%s,subMessage:%s,subCode:%s,err:%v", res.Msg, res.Code, res.SubMsg, res.SubCode, res))
	}

	return res.Data, nil
}

func (c *KuaishouClient) LinkCreate(ctx context.Context, request kuaishouapi.LinkCreateRequest, accessToken string) (kuaishouapi.LinkCreateResponseData, error) {

	var res kuaishouapi.LinkCreateResponse
	httpExecute, httpErr := c.execute(kuaishouapi.MethodLinkCreate, accessToken, request, 1, true)
	if httpErr != nil {
		return res.Data, httpErr
	}

	jsonErr := json.Unmarshal([]byte(httpExecute), &res)
	if jsonErr != nil {
		return res.Data, jsonErr
	}

	if res.Result != 1 {
		return res.Data, NewApiError(res.Result, fmt.Sprintf("message:%s,code:%s,subMessage:%s,subCode:%s,err:%v", res.Msg, res.Code, res.SubMsg, res.SubCode, res))
	}

	return res.Data, nil
}

func (c *KuaishouClient) ChannelList(ctx context.Context, accessToken string) ([]kuaishouapi.ChannelListResponseChannel, error) {

	var res kuaishouapi.ChannelListResponse
	httpExecute, httpErr := c.execute(kuaishouapi.MethodChannelList, accessToken, map[string]string{}, 1, false)
	if httpErr != nil {
		return res.Data, httpErr
	}

	jsonErr := json.Unmarshal([]byte(httpExecute), &res)
	if jsonErr != nil {
		return res.Data, jsonErr
	}

	if res.Result != 1 {
		return res.Data, NewApiError(res.Result, fmt.Sprintf("message:%s,code:%s,subMessage:%s,subCode:%s,err:%v", res.Msg, res.Code, res.SubMsg, res.SubCode, res))
	}

	return res.Data, nil
}

func (c *KuaishouClient) PromtionThemeItemList(ctx context.Context, request kuaishouapi.PromtionThemeItemListRequest, accessToken string) ([]kuaishouapi.KuaishouItem, string, error) {

	var res kuaishouapi.PromtionThemeItemListResponse
	httpExecute, httpErr := c.execute(kuaishouapi.MethodPromtionThemeItemList, accessToken, request, 1, false)
	if httpErr != nil {
		return res.Data, res.Pcursor, httpErr
	}

	jsonErr := json.Unmarshal([]byte(httpExecute), &res)
	if jsonErr != nil {
		return res.Data, res.Pcursor, jsonErr
	}

	if res.Result != 1 {
		return res.Data, res.Pcursor, NewApiError(res.Result, fmt.Sprintf("message:%s,code:%s,subMessage:%s,subCode:%s,err:%v", res.Msg, res.Code, res.SubMsg, res.SubCode, res))
	}

	return res.Data, res.Pcursor, nil
}
