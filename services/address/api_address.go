package address

import (
	"context"
	"github.com/gate/gatepay-sdk-go/core"
	"github.com/gate/gatepay-sdk-go/services"
	"github.com/gate/gatepay-sdk-go/services/common"
	nethttp "net/http"
	neturl "net/url"
)

type AddressApiService services.Service

func (a *AddressApiService) GetAddressChains(ctx context.Context, req ChainsRequest) (resp *ChainsResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := core.DefaultEndpoint + "/v1/pay/address/chains"
	localVarQueryParams = neturl.Values{}
	localVarQueryParams.Add("currency", req.Currency)

	localVarHTTPContentTypes := core.ApplicationJSON
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	resp = &ChainsResponse{}
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}

	return resp, result, nil
}

// /v1/pay/address/query
func (a *AddressApiService) QueryAddressOrder(ctx context.Context, req QueryAddressOrderRequest) (resp *QueryAddressOrderResp, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := core.DefaultEndpoint + "/v1/pay/address/query"
	localVarQueryParams = neturl.Values{}
	localVarQueryParams.Add("merchantTradeNo", req.MerchantTradeNo)
	localVarQueryParams.Add("prepayId", req.PrepayID)

	//set用户设置的Header
	for k, v := range req.GetHeaders() {
		localVarHeaderParams.Set(k, v)
	}

	localVarHTTPContentTypes := core.ApplicationJSON
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	resp = &QueryAddressOrderResp{}
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil

}

// /v1/pay/address/currencies
func (a *AddressApiService) GetAddressCurrencies(ctx context.Context, req common.BaseRequest) (resp *SupportedCurrenciesRes, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := core.DefaultEndpoint + "/v1/pay/address/currencies"
	localVarQueryParams = neturl.Values{}
	localVarHTTPContentTypes := core.ApplicationJSON

	//set用户设置的Header
	for k, v := range req.GetHeaders() {
		localVarHeaderParams.Set(k, v)
	}

	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	resp = &SupportedCurrenciesRes{}
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// /v1/pay/address/supportedconvertcurrencies
func (a *AddressApiService) GetAddressSupportedConvertCurrencies(ctx context.Context, req SupportedConvertCurrenciesReq) (resp *SupportedConvertCurrenciesRes, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := core.DefaultEndpoint + "/v1/pay/address/supportedconvertcurrencies"
	localVarQueryParams = neturl.Values{}
	localVarQueryParams.Add("currency", req.Currency)
	localVarHTTPContentTypes := core.ApplicationJSON

	//set用户设置的Header
	for k, v := range req.GetHeaders() {
		localVarHeaderParams.Set(k, v)
	}

	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	resp = &SupportedConvertCurrenciesRes{}
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// /v1/pay/address/create
func (a *AddressApiService) CreateAddress(ctx context.Context, req CreateOrderRequest) (resp *AddrOrderResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	//set用户设置的Header
	for k, v := range req.GetHeaders() {
		localVarHeaderParams.Set(k, v)
	}

	localVarPath := core.DefaultEndpoint + "/v1/pay/address/create"
	localVarPostBody = req
	localVarHTTPContentTypes := core.ApplicationJSON

	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	resp = &AddrOrderResponse{}
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// /v1/pay/address/refund
func (a *AddressApiService) RefundAddress(ctx context.Context, req CreateAddressRefundRequest) (resp *CreateAddressRefundResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	//set用户设置的Header
	for k, v := range req.GetHeaders() {
		localVarHeaderParams.Set(k, v)
	}

	localVarPath := core.DefaultEndpoint + "/v1/pay/address/refund"
	localVarPostBody = req
	localVarHTTPContentTypes := core.ApplicationJSON

	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	resp = &CreateAddressRefundResponse{}
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// /v1/pay/address/refundconvert
func (a *AddressApiService) RefundAddressConvert(ctx context.Context, req CreateAddressRefundConvertRequest) (resp *CreateAddressRefundConvertResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	//set用户设置的Header
	for k, v := range req.GetHeaders() {
		localVarHeaderParams.Set(k, v)
	}

	localVarPath := core.DefaultEndpoint + "/v1/pay/address/refundconvert"
	localVarPostBody = req
	localVarHTTPContentTypes := core.ApplicationJSON
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	resp = &CreateAddressRefundConvertResponse{}
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// /v1/pay/address/transactiondetail
func (a *AddressApiService) QueryAddressTransactionDetail(ctx context.Context, req TransactionDetailReq) (resp *TransactionDetailResp, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := core.DefaultEndpoint + "/v1/pay/address/transactiondetail"
	localVarQueryParams = neturl.Values{}
	localVarQueryParams.Add("prepayID", req.PrepayID)
	localVarHTTPContentTypes := core.ApplicationJSON

	//set用户设置的Header
	for k, v := range req.GetHeaders() {
		localVarHeaderParams.Set(k, v)
	}

	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	resp = &TransactionDetailResp{}
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
