func containsDuplicate(nums []int) bool {

	duplicateMap := make(map[int]struct{})

	for _, num := range nums {
		if _, ok := duplicateMap[num]; ok {
			return true
		}
		duplicateMap[num] = struct{}{}
	}
	return false
	// duplicateSet := make(map[int]struct{},len(nums))

	// for _,num :=range nums {
	//     duplicateSet[num] = struct{}{}
	// }

	// if len(nums) != len(duplicateSet) {
	//     return true
	// }
	// return false
}