package ky5

import (
	"fmt"
	"strconv"
)

func RunInt32ToIPV4() {
	fmt.Print(Convert(32))
}

func Convert(number32 int) string {
	multiplier := 0
	result := ""

	for multiplier >= 2 {
		result += strconv.Itoa(number32 % 2)
	}

	return result
}
