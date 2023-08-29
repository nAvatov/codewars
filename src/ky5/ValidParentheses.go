package ky5

import (
	"fmt"
	"strings"
)

func RunValidParentheses() {
	var (
		answer = false
	)
	answer = ValidBraces("(())((()())())")
	fmt.Println(answer)
}

func ValidBraces(s string) bool { // Approve that string isn't valid
	if len(s) == 0 {
		return true
	}
	if len(s)%2 != 0 || len(s) > 100 {
		return false
	} else {
		var (
			openBracket   = "("
			closedBracket = ")"
			slicedString  = strings.Split(s, "")
			flag          = true
			valid         = true
			i             = 0
		)

		if slicedString[0] == closedBracket && slicedString[len(slicedString)-1] == openBracket {
			return false
		} else {
			for flag {
				if len(slicedString) > 2 && i < len(slicedString)-2 && slicedString[i+1] != closedBracket {
					i++
				} else if slicedString[i] == openBracket && slicedString[i+1] == closedBracket {
					slicedString = append(slicedString[:i], slicedString[i+2:]...)
					i = 0
				} else {
					return false
				}
				if len(slicedString) == 0 {
					valid = true
					flag = false
					break
				} else if len(slicedString) == 1 {
					flag = false
					valid = false
					break
				}
			}
		}
		return valid
	}
}
