package address

import (
	"context"
	"github.com/gate/gatepay-sdk-go/core"
	"github.com/gate/gatepay-sdk-go/core/stringutillib"
	"github.com/gate/gatepay-sdk-go/services/common"
	"github.com/shopspring/decimal"
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
	service := &AddressApiService{Client: client}

	req := ChainsRequest{Currency: "USDT"}
	resp, result, err := service.GetAddressChains(ctx, req)
	if err != nil {
		log.Printf("call GetAddressChains err:%s", err)
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func TestCreateAddressOrder(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("Mz6M_q4AkDnZCSoTDo03A6OtWzN5ut8_Uix3jyVjxAU=")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &AddressApiService{Client: client}

	req := CreateOrderRequest{
		MerchantTradeNo: "j4058308409230424822343104",
		OrderExpireTime: 1750834613000,
		OrderAmount:     decimal.NewFromInt(1),
		Currency:        "USDT",
		CancelURL:       "https://www.baidu.com",
		ChannelId:       "",

		Env: EnvRequest{
			Scene:        "APP",
			TerminalType: "IOS",
		},
		Goods: GoodsRequest{
			GoodsDetail: "goods",
			GoodsName:   "goods",
		},
		Chain:          "ETH",
		FullCurrType:   "USDT_ETH",
		MerchantUserId: 6790011,
	}

	req.AddHeader("X-GatePay-Certificate-ClientId", "mZ96D37oKk-HrWJc")
	resp, result, err := service.CreateAddress(ctx, req)
	if err != nil {
		log.Printf("call CreateAddress err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func TestQueryAddressOrder(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("Mz6M_q4AkDnZCSoTDo03A6OtWzN5ut8_Uix3jyVjxAU=")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		return
	}

	ctx := context.Background()
	service := &AddressApiService{Client: client}
	req := QueryAddressOrderRequest{MerchantTradeNo: "j4058308409230424822343104", PrepayID: "372425901225357312"}
	req.AddHeader("X-GatePay-Certificate-ClientId", "mZ96D37oKk-HrWJc")

	resp, result, err := service.QueryAddressOrder(ctx, req)
	if err != nil {
		log.Printf("call QueryAddressOrder err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func TestGetAddressCurrencies(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("Mz6M_q4AkDnZCSoTDo03A6OtWzN5ut8_Uix3jyVjxAU=")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		return
	}

	ctx := context.Background()
	service := &AddressApiService{Client: client}
	//接口没有参数，可以传common.BaseRequest,这样可以在BaseRequest传入client_id用于签名
	req := common.BaseRequest{}
	req.AddHeader("X-GatePay-Certificate-ClientId", "mZ96D37oKk-HrWJc")

	resp, result, err := service.GetAddressCurrencies(ctx, req)
	if err != nil {
		log.Printf("call QueryAddressOrder err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func TestSupportedConvertCurrencies(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("Mz6M_q4AkDnZCSoTDo03A6OtWzN5ut8_Uix3jyVjxAU=")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		return
	}

	ctx := context.Background()
	service := &AddressApiService{Client: client}

	req := SupportedConvertCurrenciesReq{Currency: "ETH"}
	req.AddHeader("X-GatePay-Certificate-ClientId", "mZ96D37oKk-HrWJc")
	resp, result, err := service.GetAddressSupportedConvertCurrencies(ctx, req)
	if err != nil {
		log.Printf("call QueryAddressOrder err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func TestQueryAddressTransactionDetail(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("Mz6M_q4AkDnZCSoTDo03A6OtWzN5ut8_Uix3jyVjxAU=")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		return
	}

	ctx := context.Background()
	service := &AddressApiService{Client: client}
	req := TransactionDetailReq{PrepayID: "370615104821190656"}

	req.AddHeader("X-GatePay-Certificate-ClientId", "mZ96D37oKk-HrWJc")
	resp, result, err := service.QueryAddressTransactionDetail(ctx, req)
	if err != nil {
		log.Printf("call QueryAddressOrder err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}
