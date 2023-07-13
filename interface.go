/**
 * @Title
 * @Author: liwei
 * @Description:  TODO
 * @File:  interface
 * @Version: 1.0.0
 * @Date: 2023/07/13 10:35
 * @Update liwei 2023/7/13 10:35
 */

package mycard

type WebSdk interface {
	MakeMyCardAuth() IsMyCardAuth
	MakeMyCardPay() MyCardPay
}

type IsMyCardAuth interface {
	MyCardCallbackRequestPreHashValue() string
}

type MyCardPay interface {
	//授权下单
	AuthGlobal(request *AuthGlobalRequest) (*MyCardAuthGlobalResponses, error)

	hashValue(request *AuthGlobalRequest) string

	//查询订单
	TradeQuery(authCode string, isSandBox int) (*MyCardTradeQueryResponses, error)

	//请款
	PayMentConfirm(authCode string, isSandBox int) (*MyCardPaymentConfirmResponses, error)
}
