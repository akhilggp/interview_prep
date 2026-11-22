package dsa

func topKFrequent(nums []int, k int) []int {

	uniqueCount := make(map[int]int)

	for _, num := range nums {
		uniqueCount[num]++
	}
	// maxK := len(uniqueCount)
	bucketList := make([][]int, len(nums)+1)

	for key, value := range uniqueCount {
		bucketList[value] = append(bucketList[value], key)
	}
	result := make([]int, 0, k)
	for i := len(bucketList) - 1; i > 0 && len(result) < k; i-- {
		if bucketList[i] != nil {
			result = append(result, bucketList[i]...)
		}
	}
	return result[:k]
}
