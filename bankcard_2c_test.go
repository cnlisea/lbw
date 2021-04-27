package lbw

import "testing"

func TestService_Bankcard2C(t *testing.T) {
	s := NewService("", "", "", "", nil)
	res, err := s.Bankcard2C("", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(*res)
}
