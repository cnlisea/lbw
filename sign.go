package lbw

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func (s *Service) Sign(orderId string, ts string) string {
	hash := md5.New()
	hash.Write([]byte(strings.Join([]string{
		s.MchId,
		orderId,
		ts,
		s.MchSecurity,
	}, "&")))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
