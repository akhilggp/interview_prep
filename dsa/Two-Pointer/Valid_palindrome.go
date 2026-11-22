package twopointer

import "fmt"

func isPalindrome(s string) bool {
	compString := cleanPalindrome(s)
	fmt.Println(compString)
	i, j := 0, len(compString)-1
	for i < j {
		if compString[i] != compString[j] {
			return false
		}
		i++
		j--
	}
	return true
}
func isAlNum(c byte) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}

func cleanPalindrome(s string) []byte {
	clean := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if !isAlNum(c) {
			continue
		}
		if c >= 'A' && c <= 'Z' {
			c += 32
		}
		clean = append(clean, c)
	}
	return clean
}
