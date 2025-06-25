package address

import (
	"github.com/gate/gatepay-sdk-go/services/common"
	"github.com/shopspring/decimal"
)

type ChainsRequest struct {
	*common.BaseRequest
	Currency string `json:"currency"`
}

type ChainsResponse struct {
	Currency string           `json:"currency"`
	Chains   []*ChainNameItem `json:"chains"`
}

type ChainNameItem struct {
	Chain           string `json:"chain"`
	Currency        string `json:"currency"`       // 单纯的币种名称
	FullCurrType    string `json:"full_curr_type"` // 含链信息
	Symbol          string `json:"symbol"`
	ExplorerUrl     string `json:"explorer_url"`
	ShowChainNameEn string `json:"show_chain_name_en"`
}

type EnvRequest struct {
	TerminalType string `json:"terminalType" validator:"nonemptyString"`
	Scene        string `json:"scene" validator:"scene"`
}

type GoodsRequest struct {
	GoodsName   string `json:"goodsName"`
	GoodsDetail string `json:"goodsDetail"`
}

type CreateOrderRequest struct {
	common.BaseRequest
	MerchantTradeNo string          `json:"merchantTradeNo" validator:"nonemptyString"`
	Currency        string          `json:"currency" validator:"nonemptyString"`    // order_currency
	OrderAmount     decimal.Decimal `json:"orderAmount" validator:"boundedDecimal"` // order_amount default zero
	PayCurrency     string          `json:"payCurrency"`                            // pay_currency 非地址支付PayCurrency 在实际付款确定 ，地址支付在下单时候确定
	ActualCurrency  string          `json:"actualCurrency"`                         // merchant actual currency
	Env             EnvRequest      `json:"env"`
	Goods           GoodsRequest    `json:"goods"`
	OrderExpireTime int64           `json:"orderExpireTime"`
	ReturnURL       string          `json:"returnUrl"`
	CancelURL       string          `json:"cancelUrl"`
	MerchantUserId  int64           `json:"merchantUserId"` // userId in merchants' system
	Chain           string          `json:"chain"`
	FullCurrType    string          `json:"fullCurrType"`
	ChannelId       string          `json:"channelId"` // 客户渠道名称
}

type Chain struct {
	common.BaseRequest
	ChainType    string `json:"chain_type"`
	Address      string `json:"address"`
	FullCurrType string `json:"fullCurrType,omitempty"`
}

type AddrOrderResponse struct {
	PrepayID     string `json:"prepayId"`
	TerminalType string `json:"terminalType"`
	ExpireTime   int64  `json:"expireTime"`
	ChainInfo    Chain  `json:"chain"`
}

type QueryAddressOrderRequest struct {
	common.BaseRequest
	PrepayID        string `query:"prepayId" json:"prepayId"`
	MerchantTradeNo string `query:"merchantTradeNo" json:"merchantTradeNo"`
}

type QueryAddressOrderResp struct {
	PrepayID        string                `json:"prepayId"`
	MerchantID      int64                 `json:"merchantId"`
	MerchantTradeNo string                `json:"merchantTradeNo"`
	TransactionID   string                `json:"transactionId"`
	GoodsName       string                `json:"goodsName"`
	Currency        string                `json:"currency"`
	OrderAmount     string                `json:"orderAmount"`
	PayCurrency     string                `json:"payCurrency"`
	PayAmount       decimal.Decimal       `json:"payAmount"`
	Rate            decimal.Decimal       `json:"rate"`
	Status          string                `json:"status"`
	CreateTime      int64                 `json:"createTime"`
	ExpireTime      int64                 `json:"expireTime"`
	TransactTime    int64                 `json:"transactTime"`
	OrderName       string                `json:"order_name"`
	TransactionInfo *ChainTransactionInfo `json:"transaction_info"`
	ChannelId       string                `json:"channelId"` // 客户渠道名称
	Address         string                `json:"address"`   // 地址
	Chain           string                `json:"chain"`     // 网络
}

type ChainTransactionInfo struct {
	DoneAmount     decimal.Decimal `json:"done_amount"`
	ConfirmingList []*ConfirmItem  `json:"confirming_list"`
}

type ConfirmItem struct {
	Amount  decimal.Decimal `json:"amount"`
	Confirm int             `json:"confirm"`
}

type SupportedCurrenciesRes struct {
	Currencies []string `json:"currencies"`
}

type SupportedConvertCurrenciesReq struct {
	common.BaseRequest
	Currency string `query:"currency" json:"currency"`
}

type SupportedConvertCurrenciesRes struct {
	Currencies []string `json:"currencies"`
}

type CreateAddressRefundRequest struct {
	common.BaseRequest
	RefundRequestID string          `json:"refundRequestId" validator:"nonemptyString"`
	PrepayID        string          `json:"prepayId" validator:"nonemptyString"`
	RefundAmount    decimal.Decimal `json:"refundAmount"  validator:"boundedDecimal"`
	RefundReason    string          `json:"refundReason"`
	ReceiverId      int64           `json:"receiverId"`
}

type CreateAddressRefundResponse struct {
	RefundRequestID string `json:"refundRequestId"`
	PrepayID        string `json:"prepayId"`
	OrderAmount     string `json:"orderAmount"`
	RefundAmount    string `json:"refundAmount"`
}

type CreateAddressRefundConvertRequest struct {
	common.BaseRequest
	RefundRequestID     string          `json:"refundRequestId"`
	PrepayID            string          `json:"prepayId"`
	RefundOrderCurrency string          `json:"refundOrderCurrency"`
	RefundOrderAmount   decimal.Decimal `json:"refundOrderAmount"`
	RefundPayCurrency   string          `json:"refundPayCurrency"`
	RefundPayAmount     decimal.Decimal `json:"refundPayAmount"`
	RefundReason        string          `json:"refundReason"`
	ReceiverId          int64           `json:"receiverId"`
}

type CreateAddressRefundConvertResponse struct {
	RefundRequestID   string          `json:"refundRequestId"`
	PrepayID          string          `json:"prepayId"`
	OrderCurrency     string          `json:"orderCurrency"`
	OrderAmount       decimal.Decimal `json:"orderAmount"`
	RefundOrderAmount decimal.Decimal `json:"refundOrderAmount"`
	PayCurrency       string          `json:"payCurrency"`
	PayAmount         decimal.Decimal `json:"payAmount"`
	RefundPayAmount   decimal.Decimal `json:"refundPayAmount"`
}

type TransactionDetailReq struct {
	common.BaseRequest
	PrepayID string `json:"prepayID"`
}

type TransactionDetailResp struct {
	PrepayID          string             `json:"prepayId"`
	MerchantID        int64              `json:"merchantId"`
	MerchantTradeNo   string             `json:"merchantTradeNo"`
	TransactionID     string             `json:"transactionId"`
	GoodsName         string             `json:"goodsName"`
	Currency          string             `json:"currency"`
	OrderAmount       string             `json:"orderAmount"`
	PayCurrency       string             `json:"payCurrency"`
	PayAmount         string             `json:"payAmount"`
	Status            string             `json:"status"`
	UtcCreateTime     string             `json:"utcCreateTime"`
	UtcExpireTime     string             `json:"utcExpireTime"`
	UtcUpdateTime     string             `json:"utcUpdateTime"`
	TransactTime      int64              `json:"transactTime"`
	OrderName         string             `json:"order_name"`
	TransactionDetail *TransactionDetail `json:"transactionDetail"`
	ChannelId         string             `json:"channelId"`
}

type TransactionDetail struct {
	InTerm    *TxDetail `json:"inTerm"`
	OutOfTerm *TxDetail `json:"outOfTerm"`
}

type TxDetail struct {
	Done *TxDetailStateItem `json:"done"`
	Wait *TxDetailStateItem `json:"wait"`
}

type TxDetailStateItem struct {
	Amount decimal.Decimal `json:"amount"`
	TxList []*TxItem       `json:"txList"`
}

type TxItem struct {
	Chain         string          `json:"chain"`
	Address       string          `json:"address"`
	FullCurrType  string          `json:"fullCurrType"`
	Amount        decimal.Decimal `json:"amount"`
	TxId          string          `json:"txId"`
	UtcCreateTime string          `json:"utcCreateTime"`
	UtcUpdateTime string          `json:"utcUpdateTime"`
}
