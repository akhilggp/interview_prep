package dsa

import "strings"

type Solution struct{}

func (s *Solution) Encode(strs []string) string {

	encodeChar := rune(256)
	var result string
	for _, str := range strs {
		result = result + str + string(encodeChar)
	}
	return result
}

func (s *Solution) Decode(encoded string) []string {
	encodeChar := rune(256)
	result := strings.Split(encoded, string(encodeChar))
	return result[:len(result)-1]
}
