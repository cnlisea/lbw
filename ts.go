package lbw

import (
	"strconv"
	"time"
)

func (s *Service) Ts() string {
	return strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
}
