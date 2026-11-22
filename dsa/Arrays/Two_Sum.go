package dsa

func twoSum(nums []int, target int) []int {
	setMap := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if j, ok := setMap[target-nums[i]]; ok {
			return []int{i, j}
		}
		setMap[nums[i]] = i
	}
	return nil
}
