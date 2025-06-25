package payment

import (
	"context"
	"github.com/gate/gatepay-sdk-go/core"
	"github.com/gate/gatepay-sdk-go/core/stringutillib"
	"log"
	"testing"
)

func TestChains(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("Mz6M_q4AkDnZCSoTDo03A6OtWzN5ut8_Uix3jyVjxAU=")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		return
	}

	ctx := context.Background()
	service := &PayApiService{Client: client}

	req := OperateOrderRequest{PrepayID: "370615104821190656", MerchantTradeNo: ""}
	req.AddHeader("X-GatePay-Certificate-ClientId", "mZ96D37oKk-HrWJc")
	resp, result, err := service.GetOrder(ctx, req)
	if err != nil {
		log.Printf("call GetAddressChains err:%s", err)
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}
