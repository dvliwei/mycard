/**
 * @Title
 * @Author: liwei
 * @Description:  TODO
 * @File:  types
 * @Version: 1.0.0
 * @Date: 2023/07/13 10:39
 * @Update liwei 2023/7/13 10:39
 */

package mycard

var MyCardHost = "https://b2b.mycard520.com.tw"
var MyCardSandBoxHost = "https://testb2b.mycard520.com.tw"

type AuthGlobalRequest struct {
	FacServiceId string `json:"FacServiceId"`
	FacTradeSeq  string `json:"FacTradeSeq"`
	FacGameId    string `json:"FacGameId"`
	FacGameName  string `json:"FacGameName"`
	TradeType    string `json:"TradeType"`
	ServerId     string `json:"ServerId"`
	CustomerId   string `json:"CustomerId"`
	PaymentType  string `json:"PaymentType"`
	ItemCode     string `json:"ItemCode"`
	ProductName  string `json:"ProductName"`
	Amount       string `json:"Amount"`
	Currency     string `json:"Currency"`
	SandBoxMode  string `json:"SandBoxMode"`
	FacReturnURL string `json:"FacReturnURL"`
	Hash         string `json:"Hash"`
}

type MyCardAuthGlobalResponses struct {
	InGameSaveType string `json:"InGameSaveType"`
	ReturnCode     string `json:"ReturnCode"`
	ReturnMsg      string `json:"ReturnMsg"`
	AuthCode       string `json:"AuthCode"`
	TradeSeq       string `json:"TradeSeq"`
	TransactionUrl string `json:"TransactionUrl"`
	//自定义补充
	Receipt string `json:"Receipt,omitempty"`
}

type MyCardTradeQueryResponses struct {
	ReturnCode    string `json:"ReturnCode"`    //查詢結果代碼 1 為成功 其他則為失敗  注意: ReturnCode 為 1 並不代表交易成功， 正確交易結果請參考 PayResult
	ReturnMsg     string `json:"ReturnMsg"`     //ReturnCode 訊息描述
	PayResult     string `json:"PayResult"`     // 交易結果代碼  交易成功為 3;交易失敗為 0
	FacTradeSeq   string `json:"FacTradeSeq"`   //廠商交易序號
	PaymentType   string `json:"PaymentType"`   //付費方式
	Amount        string `json:"Amount"`        //金額
	Currency      string `json:"Currency"`      //幣別
	MyCardTradeNo string `json:"MyCardTradeNo"` // aymentType = INGAME 時，傳 MyCard 卡   PaymentType = COSTPOINT 時，傳會員扣 點交易序號，格式為 CGM 開頭+數字 其餘 PaymentType 為 Billing 小額付款交 易，傳 Billing 交易序號  若 BILLING 為卡片儲值時，此欄位傳 MyCard 卡片號碼    string `json:"MyCardType"`
	PromoCode     string `json:"PromoCode"`     //通路代碼 PaymentType = INGAME 時才有值
	SerialId      string `json:"SerialId"`      //活動代碼

	//自定义补充
	Receipt   string `json:"Receipt,omitempty"`
	IsSandBox int    `json:"IsSandBox,omitempty"`
}

type MyCardPaymentConfirmResponses struct {
	ReturnCode  string `json:"ReturnCode"`
	ReturnMsg   string `json:"ReturnMsg"`
	FacTradeSeq string `json:"FacTradeSeq"`
	TradeSeq    string `json:"TradeSeq"`
	SerialId    string `json:"SerialId"`
	//自定义补充
	Receipt   string `json:"Receipt,omitempty"`
	IsSandBox int    `json:"IsSandBox,omitempty"`
}
