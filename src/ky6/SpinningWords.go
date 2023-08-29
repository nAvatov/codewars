package ky6

import (
	"fmt"
	"strings"
)

func RunSpinningWords() {
	fmt.Println(SpinWords("Hey fellow warriors"))
}

/*

Write a function that takes in a string of one or more words, and returns the same string, but with all five or more letter words reversed (like the name of this kata).

Strings passed in will consist of only letters and spaces.
Spaces will be included only when more than one word is present.
Examples:

spinWords("Hey fellow warriors") => "Hey wollef sroirraw"
spinWords("This is a test") => "This is a test"
spinWords("This is another test") => "This is rehtona test"

*/

func SpinWords(str string) string {
	var (
		words = make([]string, 0)
	)

	words = strings.Split(str, " ")

	for i, z := range words {
		if len(z) >= 5 {
			words[i] = string(ReverseString(z))
		}
	}
	return BuildString(words)
}

func ReverseString(s string) []byte {
	reversedString := make([]byte, 0)
	for i := range s {
		reversedString = append(reversedString, s[len(s)-i-1])
	}
	return reversedString
}

func BuildString(a []string) string {
	var buildedString string
	for i, word := range a {
		if i != len(a)-1 {
			buildedString += word + " "
		} else {
			buildedString += word
		}
	}
	return buildedString
}
