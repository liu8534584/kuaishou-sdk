package kuaishou

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
)

// Sign 计算签名
func (c *KuaishouClient) sign(accessToken, method string, timestamp int64, version int, param interface{}) (paramJson, sign string) {

	paramByte, _ := json.Marshal(param)
	paramJson = string(paramByte)

	// 按给定规则拼接参数
	paramPattern := fmt.Sprintf("access_token=%v&appkey=%v&method=%v&param=%v&signMethod=%v&timestamp=%d&version=%d&signSecret=%v",
		accessToken, c.appKey, method, paramJson, "MD5", timestamp, version, c.signSecret)

	// HMAC_SHA256
	sign = Md5(paramPattern)
	return
}

// MD5
func Md5(s string) string {
	h := md5.New()
	_, _ = io.WriteString(h, s)
	return hex.EncodeToString(h.Sum(nil))
}
