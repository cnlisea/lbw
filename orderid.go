package lbw

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func (s *Service) OrderId() string {
	if s.OrderIdFunc != nil {
		return s.OrderIdFunc()
	}

	var (
		b strings.Builder
		t = time.Now()
	)
	b.WriteString(strconv.Itoa(t.Year()))       // year
	b.WriteString(strconv.Itoa(int(t.Month()))) // month
	b.WriteString(strconv.Itoa(t.Day()))        // day
	b.WriteString(strconv.Itoa(t.Hour()))       // hour
	b.WriteString(strconv.Itoa(t.Minute()))     // minute
	b.WriteString(strconv.Itoa(t.Second()))     // second
	b.WriteString(strconv.Itoa(t.Nanosecond())) // nanosecond
	omiNum := 32 - b.Len()
	if omiNum > 0 {
		var (
			temp = 1
			i    int
		)
		for i = 0; i < omiNum; i++ {
			temp = temp * 10
		}
		randomNum := strconv.Itoa(rand.Int() % temp)
		for i = 0; i < omiNum-len(randomNum); i++ {
			b.WriteString("0")
		}
		b.WriteString(randomNum) // random number
	}

	return b.String()
}
