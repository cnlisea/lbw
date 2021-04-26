package lbw

import (
	"strings"
)

func ValueEncode(params map[string]string) string {
	array := make([]string, 0, len(params))
	for key, value := range params {
		if key == "" || value == "" {
			continue
		}
		array = append(array, key+"="+value)
	}
	return strings.Join(array, "&")
}
