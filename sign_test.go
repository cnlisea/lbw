package lbw

import (
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestService_Sign(t *testing.T) {
	s := NewService("测试商户", "xyzxyzxyz", "efcefcefcefcefc", "01", nil)
	t.Log(s.Sign("nonononononononononono", "1555378976238"))
}

func TestNewService(test *testing.T) {
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
	test.Log(b.String())
}
