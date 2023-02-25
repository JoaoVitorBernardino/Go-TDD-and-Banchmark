package main

import "fmt"

func main() {
	fmt.Println(isPalindrome("reviver"))
	fmt.Println(isPalindrome2("reviver"))

}

func isPalindrome(text string) bool {
	if len(text) == 0 {
		return false
	}

	return invertText(text) == text
}

func isPalindrome2(text string) bool {
	if len(text) == 0 {
		return false
	}

	for i, j := 0, len(text)-1; i < j; i, j = i+1, j-1 {
		if text[i] != text[j] {
			return false
		}
	}

	return true
}

func invertText(text string) string {
	var inverted string

	for i := len(text) - 1; i >= 0; i-- {
		inverted += string(text[i])
	}

	return inverted
}
