package ky6

import (
	"fmt"
	"strings"
)

func RunWeirdCase() {
	fmt.Println(toWeirdCase("ABC"))
}

/*
Write a function toWeirdCase (weirdcase in Ruby) that accepts a string, and returns the same string
with all even indexed characters in each word upper cased, and all odd indexed characters in each word
lower cased. The indexing just explained is zero based, so the zero-ith index is even, therefore that
character should be upper cased and you need to start over for each word.

The passed in string will only consist of alphabetical characters and spaces(' ').
Spaces will only be present if there are multiple words. Words will be separated by a single space(' ').
*/

func toWeirdCase(str string) string {
	// Your code here and happy coding!
	var (
		weirdString string
		even        = true
	)

	weirdString = strings.Map(
		func(r rune) rune {
			if even && r >= 97 && r <= 122 {
				r -= 32
			} else if !even && r >= 65 && r <= 90 {
				r += 32
			}
			if r == 32 {
				even = true
			} else {
				even = !even
			}
			return r
		},

		str)

	return weirdString
}
