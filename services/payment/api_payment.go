package payment

import (
	"context"
	"github.com/gate/gatepay-sdk-go/core"
	"github.com/gate/gatepay-sdk-go/services"
	nethttp "net/http"
	neturl "net/url"
)

type PayApiService services.Service

// /v1/pay/order/query
func (a *PayApiService) GetOrder(ctx context.Context, req OperateOrderRequest) (resp *QueryOrderResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := core.DefaultEndpoint + "/v1/pay/order/query"
	localVarPostBody = req

	//set用户设置的Header
	for k, v := range req.GetHeaders() {
		localVarHeaderParams.Set(k, v)
	}

	localVarHTTPContentTypes := core.ApplicationJSON
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	resp = &QueryOrderResponse{}
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}

	return resp, result, nil

}
