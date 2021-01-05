package http

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/kaylyu/tcb-openapi-go-sdk/context"
	"github.com/kaylyu/tcb-openapi-go-sdk/util"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
)

type signIn struct {
	host        string    // 域名
	uri         string    // URL 参数，不包含域名
	method      string    // 请求方法
	contentType string    // 请求方法
	queryString string    // GET: 为 ? 后面字符串; POST: 该参数为空
	payload     []byte    // GET: 固定为空字符串; POST: HTTP 请求正文 payload
	now         time.Time // 请求时间
	secretId    string    // secretId
	secretKey   string    // SecretKey
	debug       bool      // 显示签名过程
}

func TestSign() (s string, authorization string, err error) {
	return NewHttpClient(&context.Context{
		Logger: &logrus.Logger{
			Out:          os.Stdout,
			Formatter:    &util.CustomerFormatter{Prefix: "TestSign"},
			Level:        logrus.DebugLevel,
			ExitFunc:     os.Exit,
			ReportCaller: true,
		},
	}, nil).signature(&signIn{
		host:        "api.tcloudbase.com",
		uri:         "//api.tcloudbase.com/",
		method:      "POST",
		contentType: "application/json; charset=utf-8",
		queryString: "",
		payload:     []byte(""),
		now:         time.Unix(1600227242, 0),
		secretId:    "xx",
		secretKey:   "yy",
		debug:       true,
	})
}

// 签名
/*
https://docs.cloudbase.net/api-reference/openapi/introduction.html#ru-he-huo-qu-ping-zheng

支持的 HTTP 请求方法：GET, POST, PUT, PATCH, DELETE

HTTP HEADER
Content-Type:application/json
X-CloudBase-Authorization: your authorization
X-CloudBase-SessionToken: your sessiontoken
X-CloudBase-TimeStamp: 1551113065
*/
func (c *Client) signature(in *signIn) (signature string, authorization string, err error) {
	if in.debug {
		c.ctx.Logger.Debugln("++++++++++ 生成签名 ++++++++++")
	}
	in.now = in.now.UTC()
	// 1 构造待签字符串
	// 参与签名的头部信息，至少包含 host 和 content-type 两个头部，也可加入定义的头部参与签名以提高自身请求的唯一性和安全性。
	// 拼接规则:头部key和value统一转成小 写，并去掉首尾空格，按照 key:value\n 格式拼接
	canonicalHeaders := fmt.Sprintf("content-type:%s\nhost:%s\n", in.contentType, in.host)
	// 参与签名的头部key信息，说明此次请求有哪些头部参与了签名，和 CanonicalHeaders 包含的头部内容是一一对应的
	var keys []string
	if in.contentType != "" {
		keys = append(keys, "content-type")
	}
	if in.host != "" {
		keys = append(keys, "host")
	}
	if len(keys) == 0 {
		canonicalHeaders = ""
	}
	signedHeaders := strings.Join(keys, ";")

	// 2.CanonicalRequest
	// Lowercase(HexEncode(Hash.SHA256(RequestPayload)))
	// 对HTTP请求整个正文 payload 做 SHA256 哈希，然后十六进制编码，最后编码串转换成小写字母。
	sha := sha256.New()
	sha.Write(in.payload)
	payloadHash := sha.Sum(nil)
	payloadHashStr := strings.ToLower(hex.EncodeToString(payloadHash))
	canonicalRequest := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s",
		in.method,
		in.uri,
		"", //in.queryString,
		canonicalHeaders,
		signedHeaders,
		payloadHashStr)
	if in.debug {
		c.ctx.Logger.Debugln("canonicalRequest:\n", canonicalRequest)
	}
	// 待签名信息哈希值
	// HashedCanonicalRequest = Lowercase(HexEncode(Hash.SHA256(CanonicalRequest)));
	// 结果形如:ebcf3607af71c291311ede9aa5a7ac04d09c12212cdb701fa9390a9ba14adbeb
	sha = sha256.New()
	sha.Write([]byte(canonicalRequest))
	canonicalRequestHash := strings.ToLower(hex.EncodeToString(sha.Sum(nil)))
	if in.debug {
		c.ctx.Logger.Debugln("canonicalRequestHash：", canonicalRequestHash)
	}

	// 3.StringToSign
	algorithm := "TC3-HMAC-SHA256"
	credDate := in.now.Format("2006-01-02")
	credentialScope := credDate + "/tcb/tc3_request" // 如4.1章节示例请求中，取值为 2019-07-28/tcb/tc3_request
	// 签名原串
	/*
		StringToSign = Algorithm + '\n' 签名算法，固定值
			+ RequestTimestamp + '\n' 请求时间戳 UTC， 即请求头部的X-TDEA-Timestamp取值，形如 1564352003
			+ CredentialScope + '\n' 凭证范围，包含日期、所请求的服务和终止字符串('tc3_request')，即请 求头部Authorization中的Credential包含内容。。。
			+ HashedCanonicalRequest; 即4.3.2章节中所计算得到的请求串哈希值
	*/
	stringToSign := algorithm + "\n" + fmt.Sprintf("%d", in.now.Unix()) + "\n" + credentialScope + "\n" + canonicalRequestHash
	if in.debug {
		c.ctx.Logger.Debugln("stringToSign：\n", stringToSign)
	}

	// 4.Signature
	// 计算派生签名密钥
	secretSigning := getSignatureKey(in.secretKey, credDate)
	//# 计算签名摘要
	signature = hex.EncodeToString(sign(secretSigning, []byte(stringToSign)))

	// 5.拼装
	authorization = fmt.Sprintf("TC3-HMAC-SHA256 Credential=%s/%s/tcb/tc3_request,SignedHeaders=%s,Signature=%s",
		in.secretId, in.now.Format("2006-01-02"), signedHeaders, signature)
	if in.debug {
		c.ctx.Logger.Debugln("authorization：", authorization)
		c.ctx.Logger.Debugln("++++++++++ 生成签名 ++++++++++")
	}
	return
}

// 计算签名
func getSignatureKey(key, creDate string) (secretSigning []byte) {
	secretDate := sign([]byte("TC3"+key), []byte(creDate))
	//c.ctx.Logger.Debugln("secretDate:", secretDate)

	secretService := sign(secretDate, []byte("tcb"))
	//c.ctx.Logger.Debugln("secretService:", secretService)

	secretSigning = sign(secretService, []byte("tc3_request"))
	//c.ctx.Logger.Debugln("secretSigning:", secretSigning)
	return
}

//
func sign(key, msg []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(msg)
	return mac.Sum(nil)
}
