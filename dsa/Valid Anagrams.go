func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	stringMap := make(map[byte]int, 26)
	for i := 0; i < len(s); i++ {
		stringMap[s[i]]++
		stringMap[t[i]]--
	}
	for _, val := range stringMap {
		if val != 0 {
			return false
		}
	}
	return true
}