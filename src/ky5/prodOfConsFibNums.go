package ky5

import "fmt"

func RunProdOfConsFibNums() {
	fmt.Println(ProductFib(800))
}

/*

The Fibonacci numbers are the numbers in the following integer sequence (Fn):

0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, ...

such as

F(n) = F(n-1) + F(n-2) with F(0) = 0 and F(1) = 1.

productFib(714) # should return (21, 34, true),
                # since F(8) = 21, F(9) = 34 and 714 = 21 * 34

productFib(800) # should return (34, 55, false),
                # since F(8) = 21, F(9) = 34, F(10) = 55 and 21 * 34 < 800 < 34 * 55
*/

func ProductFib(prod uint64) [3]uint64 {
	var (
		consPair [2]uint64
		answer   uint64
	)
	answer, consPair = Fibonacci(0, 1, prod)
	return [3]uint64{consPair[0], consPair[1], answer}
}

func Fibonacci(f uint64, s uint64, some uint64) (uint64, [2]uint64) {
	if f*s == some {
		return 1, [2]uint64{f, s}
	} else if f*s > some {
		return 0, [2]uint64{f, s}
	}

	return Fibonacci(s, f+s, some)
}
