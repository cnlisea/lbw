package lbw

import (
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"strings"
	"time"
)

type Bankcard2C struct {
	Status  bool   // 认证状态
	ErrMsg  string // 错误信息
	OrderId string // 订单号
}

func (s *Service) Bankcard2C(realName string, bankcard string) (*Bankcard2C, error) {
	var (
		reqParam = map[string]string{
			"bizType":    "010101",
			"merchantId": s.MchId,
			"pan":        bankcard,
			"chName":     realName,
			"mchInfo":    s.MchInfo,
			"riskInfo":   s.RiskInfo,
		}
		ts      string
		orderId string
		client  = &http.Client{
			Timeout: time.Second * 3, // settings 3 second http request timeout
		}
		req     *http.Request
		res     *http.Response
		resBody struct {
			RespCode        string `json:"respCode"`
			RespDesc        string `json:"respDesc"`
			MerchantTraceNo string `json:"merchantTraceNo"`
		}
		err error
	)
	for {
		ts = s.Ts()
		orderId = s.OrderId()
		reqParam["merchantTraceNo"] = orderId
		reqParam["merchantRequestTime"] = ts
		reqParam["sign"] = s.Sign(orderId, ts)
		req, err = http.NewRequest(http.MethodPost, "https://99num.com/dsp-v1/api/validate/bank", strings.NewReader(ValueEncode(reqParam)))
		if err != nil {
			return nil, err
		}

		res, err = client.Do(req)
		if err != nil {
			return nil, err
		}

		err = jsoniter.NewDecoder(res.Body).Decode(&resBody)
		res.Body.Close()
		if err != nil {
			return nil, err
		}

		if resBody.RespCode != "S0000020" {
			break
		}
	}
	return &Bankcard2C{
		Status:  resBody.RespCode == "00000000",
		ErrMsg:  resBody.RespDesc,
		OrderId: resBody.MerchantTraceNo,
	}, nil
}
