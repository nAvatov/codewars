package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	result := Buddy(57345, 90061)
	fmt.Println(result)
	fmt.Println(time.Since(start).Seconds())
}

//57345, 90061
// Find all sums of each number between start and limit
// After that try to search second number
func Buddy(start, limit int) []int {
	if start < limit {
		buddyNumber := 0

		for number := start; number <= limit; number++ {
			buddyNumber = SummarizeProperDivisors(number) - 1
			//fmt.Println(buddyNumber)

			if buddyNumber > number && SummarizeProperDivisors(buddyNumber)-1 == number {
				return []int{number, buddyNumber}
			}

		}

		return nil
	}

	return nil
}

func SummarizeProperDivisors(num int) int {
	sum := 0

	for i := 1; i <= int(num/2); i++ {
		if num%i == 0 {
			sum += i
		}

		if i > 2 && sum < 2 { // Extract simple divisors
			return sum
		}
	}

	return sum
}
