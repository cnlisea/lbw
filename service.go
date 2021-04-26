package lbw

type Service struct {
	MchInfo     string        // merchant info
	MchId       string        // merchant id
	MchSecurity string        // merchant security
	RiskInfo    string        // risk info
	OrderIdFunc func() string // order id gen
}

func NewService(mchName string, mchId string, mchSecurity string, riskInfo string, orderIdFunc func() string) *Service {
	return &Service{
		MchInfo:     mchName,
		MchId:       mchId,
		MchSecurity: mchSecurity,
		RiskInfo:    riskInfo,
		OrderIdFunc: orderIdFunc,
	}
}
