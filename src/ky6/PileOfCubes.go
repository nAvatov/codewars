package ky6

import (
	"fmt"
	"math"
)

func RunPileOfCubes() {
	var m int = 4183059834009
	// var n int = 25
	fmt.Println(FindNb(m)) //4183059834009
	//fmt.Println(Test(2485))
	// fmt.Println(math.Cbrt(float64(Test(n))))
}

func FindNb(m int) int { // (n^3 + (n-1)^3 + (n-2)^3 + ... + 1^3 = m) => find n
	var (
		n     int
		predM int = 0
	)
	n = int(math.Cbrt(float64(m)))
	fmt.Println(n)
	for i := n; i >= 1; i-- {
		for j := 0; j <= i; j++ {
			predM += (i - j) * (i - j) * (i - j)
			if predM > m {
				break
			}
		}
		fmt.Println(predM)
		if predM == m {
			return i
		} else if predM < m {
			return -1
		}
		predM = 0
	}
	return -1
}

func Test(n int) int {
	var m int = 0
	for i := 0; i <= n; i++ {
		m += (n - i) * (n - i) * (n - i)
	}
	return m
}
