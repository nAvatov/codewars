package ky6

import (
	"fmt"
)

func RunAreTheTheSame() {
	var (
		a1 = []int{121, 144, 19, 161, 19, 144, 19, 11}
		a2 = []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}
	)
	fmt.Println(CompOptimized_02(a1, a2))
}

func CompOptimized(array1 []int, array2 []int) bool {
	if (array1 == nil || array2 == nil) || len(array1) != len(array2) {
		return false
	}
	var (
		flag                = false
		squared, nonSquared = Determine(array1, array2)
	)
	fmt.Println(squared, nonSquared)

	for _, s := range squared {
		for _, n := range nonSquared {
			if s == n*n {
				flag = true
				break
			}
		}
		if !flag { // If alg cant find any unsquared multiplication that equals squared value
			return false
		} else { // If alg found this equality, continue
			flag = false
		}
	}
	return true
}

func CompOptimized_02(array1 []int, array2 []int) bool {
	if (array1 == nil || array2 == nil) || len(array1) != len(array2) {
		return false
	}
	var (
		squared, nonSquared = Determine(array1, array2)
		nSq                 = make([]int, 0)
		sq                  = make([]int, 0)
		flag                = false
	)
	sq, nSq = squared, nonSquared

	for len(sq) != 0 {
		for j := 0; j < len(nSq); j++ {
			if sq[0] == nSq[j]*nSq[j] {
				flag = true
				sq = sq[1:]
				nSq = append(nSq[:j], nSq[j+1:]...)
				break
			}
		}

		if !flag {
			return flag
		} else {
			flag = false
		}
	}
	return true
}

func Determine(a []int, b []int) ([]int, []int) {
	maxA := 0
	maxB := 0
	for _, v := range a {
		if v < 0 {
			return b, a
		}
		if v > maxA {
			maxA = v
		}
	}

	for _, v := range b {
		if v < 0 {
			return a, b
		}
		if v > maxB {
			maxB = v
		}
	}

	if maxA > maxB {
		return a, b
	} else {
		return b, a
	}
}
