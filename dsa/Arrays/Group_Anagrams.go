package dsa

import "slices"

func groupAnagrams(strs []string) [][]string {
	sortedMap := make(map[string][]string, len(strs))
	for _, s := range strs {
		b := []byte(s)
		slices.Sort(b)
		key := string(b)
		sortedMap[key] = append(sortedMap[key], s)
	}
	result := make([][]string, 0, len(sortedMap))
	for _, grp := range sortedMap {
		result = append(result, grp)
	}
	return result
}
