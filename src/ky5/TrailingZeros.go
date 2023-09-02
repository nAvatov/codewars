package main

import (
	"fmt"
	"math"
)

func FindTrailingZerosAmount() {
	var example int = 12
	fmt.Print(Zeros(example))
}

func Zeros(number int) int {
	var maxBound = Log(5, float64(number))
	numberOfTrailingZeros := 0

	for k := 1; k <= int(maxBound); k++ {
		numberOfTrailingZeros += number / int(math.Pow(5, float64(k)))
	}

	return numberOfTrailingZeros
}

func Log(base, x float64) float64 {
	return math.Log(x) / math.Log(base)
}
