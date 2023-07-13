/**
 * @Title
 * @Author: liwei
 * @Description:  TODO
 * @File:  res
 * @Version: 1.0.0
 * @Date: 2023/07/13 10:46
 * @Update liwei 2023/7/13 10:46
 */

package mycard

import (
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"net/url"
	"strings"
	"time"
)

type MyCardPayRes struct {
	SandBox      string //沙盒标记
	FacServiceId string //mycard 分配的FacServiceId
	FacGameId    string //mycard 分配的FacGameId
	FacGameName  string //mycard FacGameName
	Key          string //mycard KEY
	TradeType    string //接入类型 1:Android SDK (手遊適用) 2：WEB
}

func (this *MyCardPayRes) AuthGlobal(request *AuthGlobalRequest) (*MyCardAuthGlobalResponses, error) {
	var result MyCardAuthGlobalResponses
	postUrl := ""
	if this.SandBox == "true" {
		postUrl = MyCardSandBoxHost + "/MyBillingPay/v1.2/AuthGlobal"
	} else {
		postUrl = MyCardHost + "/MyBillingPay/v1.2/AuthGlobal"
	}
	req := httplib.Post(postUrl)
	req.Param("FacServiceId", this.FacServiceId)
	req.Param("FacTradeSeq", request.FacTradeSeq)
	req.Param("FacGameId", this.FacGameId)
	req.Param("FacGameName", this.FacGameName)
	req.Param("TradeType", this.TradeType)
	req.Param("ServerId", request.ServerId)
	req.Param("CustomerId", request.CustomerId)
	req.Param("PaymentType", request.PaymentType)
	req.Param("ItemCode", request.ItemCode)
	req.Param("ProductName", request.ProductName)
	req.Param("Amount", request.Amount)
	req.Param("Currency", request.Currency)
	req.Param("SandBoxMode", this.SandBox)
	req.Param("FacReturnURL", request.FacReturnURL)
	preHashValue := this.hashValue(request)
	req.Param("Hash", preHashValue)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.SetTimeout(time.Second*30, time.Second*30)
	str, err := req.String()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(str), &result)
	if err != nil {
		return nil, err
	}
	fmt.Println(result)
	if result.ReturnCode != "1" {
		err := errors.New("mycard auth failed")
		return nil, err
	}
	result.Receipt = str
	return &result, nil
}

func (this *MyCardPayRes) hashValue(request *AuthGlobalRequest) string {
	value := request.FacServiceId + request.FacTradeSeq + request.FacGameId + request.FacGameName + this.TradeType +
		request.ServerId + request.CustomerId + request.PaymentType + request.ItemCode +
		request.ProductName + request.Amount + request.Currency + this.SandBox + request.FacReturnURL + this.Key
	value = url.QueryEscape(value)
	value = strings.ToLower(value)
	//sha256
	h := sha256.New()
	h.Write([]byte(value))
	value = hex.EncodeToString(h.Sum(nil))
	return value
}

func (this *MyCardPayRes) TradeQuery(authCode string, isSandBox int) (*MyCardTradeQueryResponses, error) {
	var result MyCardTradeQueryResponses
	getUrl := ""
	if isSandBox == 1 {
		getUrl = MyCardSandBoxHost + "/MyBillingPay/v1.2/TradeQuery"
	} else {
		getUrl = MyCardHost + "/MyBillingPay/v1.2/TradeQuery"
	}
	getUrl += "?AuthCode=" + authCode
	req := httplib.Get(getUrl)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.SetTimeout(time.Second*30, time.Second*30)
	str, err := req.String()
	if err != nil {
		err := errors.New(err.Error() + "查询地址:" + getUrl)
		return nil, err
	}
	err = json.Unmarshal([]byte(str), &result)
	if err != nil {
		err := errors.New(err.Error() + "查询地址:" + getUrl)
		return nil, err
	}
	receipt := fmt.Sprintf(`{"geturl":"%s","str":%s}`, getUrl, str)
	result.Receipt = receipt
	result.IsSandBox = isSandBox
	return &result, nil
}

func (this *MyCardPayRes) PayMentConfirm(authCode string, isSandBox int) (*MyCardPaymentConfirmResponses, error) {
	var result MyCardPaymentConfirmResponses
	getUrl := ""
	if isSandBox == 1 {
		getUrl = MyCardSandBoxHost + "/MyBillingPay/v1.2/PaymentConfirm"
	} else {
		getUrl = MyCardHost + "/MyBillingPay/v1.2/PaymentConfirm"
	}
	getUrl += "?AuthCode=" + authCode
	req := httplib.Get(getUrl)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.SetTimeout(time.Second*30, time.Second*30)
	str, err := req.String()
	if err != nil {
		err := errors.New(err.Error() + "请款地址:" + getUrl)
		return nil, err
	}
	err = json.Unmarshal([]byte(str), &result)
	if err != nil {
		err := errors.New(err.Error() + "请款地址:" + getUrl)
		return nil, err
	}

	result.Receipt = str
	result.IsSandBox = isSandBox
	return &result, nil
}
