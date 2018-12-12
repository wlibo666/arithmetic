package main

import (
	"golangsrc/fmt"
	"strings"
)

func palindrome(line string) bool {
	tmpLine := strings.ToLower(line)
	var charArray []byte
	for _, c := range tmpLine {
		if (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') {
			charArray = append(charArray, byte(c))
		}
	}
	end := len(charArray) - 1
	fmt.Printf("chararray:%s\n", charArray)
	for i := 0; i < len(charArray)/2; i++ {
		if charArray[i] == charArray[end-i] {
			continue
		}
		return false
	}
	return true
}

func Palindrome() {
	line := "A man, a plan, a canal: Panama"
	fmt.Printf("line:%s, flag:%v\n", line, palindrome(line))

	line1 := "A1 man, a plan, a canal: Panama"
	fmt.Printf("line1:%s, flag:%v\n", line1, palindrome(line1))
}
