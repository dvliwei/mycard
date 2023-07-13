/**
 * @Title
 * @Author: liwei
 * @Description:  TODO
 * @File:  mycard_test
 * @Version: 1.0.0
 * @Date: 2023/07/13 11:13
 * @Update liwei 2023/7/13 11:13
 */

package mycard

import (
	"testing"
)

func Test_AuthGlobal(t *testing.T) {
	var request AuthGlobalRequest
	var res MyCardPayRes
	res.FacServiceId = "xxxxx"
	res.FacGameName = "xxxxx"
	res.FacGameId = "xxxxx"
	res.Key = "xxxxx"
	res.SandBox = "true"
	request.ProductName = "嘉年華禮包17"
	request.Currency = "TWD"
	request.Amount = "150"
	request.PaymentType = "COSTPOINT"
	request.FacTradeSeq = "MYCARD_202307130313063481513_44"
	request.FacReturnURL = "https://test-mall-ryqmx.acesdk.cn/result"
	request.CustomerId = "4500010000135"
	request.ServerId = "400001"
	request.ItemCode = ""
	result, err := res.AuthGlobal(&request)
	if err != nil {
		t.Error(err.Error())
	}
	t.Logf("result: %v", result)
}
