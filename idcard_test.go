package lbw

import "testing"

func TestService_IdCard(t *testing.T) {
	s := NewService("", "", "", "", nil)
	res, err := s.IdCard("", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(*res)
}
