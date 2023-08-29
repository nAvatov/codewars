package ky4

import (
	"fmt"
	"strconv"
)

func RunNextBiggestNumber() {
	fmt.Println(NextBigger(59884848459853))
}

/*

Create a function that takes a positive integer and returns the next bigger number that can be formed by rearranging its digits. For example:

12 ==> 21
513 ==> 531
2017 ==> 2071

If the digits can't be rearranged to form a bigger number, return -1

*/

func NextBigger(n int) int {
	if n < 11 || (n%11 == 0) {
		return -1
	}
	var (
		num     = []byte(strconv.Itoa(n))
		i       = len(num) - 1
		current int
	)
	for {
		if i == 0 {
			return -1
		}
		if num[i] > num[i-1] { // If next digit bigger then current
			if !FindClosest(num, i-1) { // Check for previos potential swap
				num[i], num[i-1] = num[i-1], num[i] // Otherwise, swap forward
			}
			Sort(num, i)
			current, _ = strconv.Atoi(string(num))
			return current
		} else {
			i--
			continue
		}
	}
}

func FindClosest(a []byte, currentElementIndex int) bool { // When we reached another digit, this function checks - Do we have a bigger and closest at the same time than/to current.
	var (
		relevantElementIndex int  = -1
		residual             byte = a[currentElementIndex+1] - a[currentElementIndex]
	)

	for i := currentElementIndex + 2; i < len(a); i++ {
		if a[i]-a[currentElementIndex] > 0 && a[i]-a[currentElementIndex] < residual {
			relevantElementIndex = i
			residual = a[relevantElementIndex] - a[currentElementIndex]
		}
	}
	if relevantElementIndex > -1 {
		a[currentElementIndex], a[relevantElementIndex] = a[relevantElementIndex], a[currentElementIndex]
		return true
	} else {
		return false
	}
}

func Sort(a []byte, startIndex int) {
	for i := startIndex; i < len(a)-1; i++ {
		for j := i; j < len(a); j++ {
			if a[j] < a[i] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
