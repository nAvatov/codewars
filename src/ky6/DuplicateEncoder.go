package ky6

import (
	"fmt"
	"strings"
)

func RunDuplicateEncoder() {
	fmt.Println(DuplicateEncode("Success"))
}

/*
The goal of this exercise is to convert a string to a new string where each character in
the new string is "(" if that character appears only once in the original string, or ")" if
that character appears more than once in the original string. Ignore capitalization when
determining if a character is a duplicate.
"din"      =>  "((("
"recede"   =>  "()()()"
"Success"  =>  ")())())"
"(( @"     =>  "))(("
*/

func DuplicateEncode(word string) string {
	word = strings.ToLower(word)
	//ncodedWord := make([]string, 0)
	encodedWord := strings.Split(word, "")
	for i := range encodedWord {
		encodedWord[i] = "("
	}
	for i := 0; i < len(word)-1; i++ {
		for j := i + 1; j < len(word); j++ {
			if word[i] == word[j] {
				encodedWord[i] = ")"
				encodedWord[j] = ")"
			}
		}
	}
	return strings.Join(encodedWord[:], "")
}
