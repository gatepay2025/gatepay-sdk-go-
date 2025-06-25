package payment

import "github.com/gate/gatepay-sdk-go/services/common"

type OperateOrderRequest struct {
	common.BaseRequest
	PrepayID        string `json:"prepayId"`
	MerchantTradeNo string `json:"merchantTradeNo"`
}

type QueryOrderResponse struct {
	PrepayID        string `json:"prepayId"`
	MerchantID      int64  `json:"merchantId"`
	MerchantTradeNo string `json:"merchantTradeNo"`
	TransactionID   string `json:"transactionId"`
	GoodsName       string `json:"goodsName"`
	Currency        string `json:"currency"`
	OrderAmount     string `json:"orderAmount"`
	Status          string `json:"status"`
	CreateTime      int64  `json:"createTime"`
	ExpireTime      int64  `json:"expireTime"`
	TransactTime    int64  `json:"transactTime"`
	OrderName       string `json:"order_name"`
	PayCurrency     string `json:"pay_currency"`
	PayAmount       string `json:"pay_amount"`
	ExpectCurrency  string `json:"expectCurrency,omitempty"`
	ActualCurrency  string `json:"actualCurrency,omitempty"`
	ActualAmount    string `json:"actualAmount,omitempty"`
	Rate            string `json:"rate"`
	PayChan         string `json:"channel_type"`
	PayAccount      string `json:"pay_account"`
	AppName         string `json:"appName"`
	AppLogo         string `json:"appLogo"`
	InUsdt          string `json:"inUsdt"`
	ChannelId       string `json:"channelId"`
}
