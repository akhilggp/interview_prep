package dsa

func productExceptSelf(nums []int) []int {
	product := 1 // product of the non-zero numbers
	zeroCount := 0

	for _, num := range nums {
		if num == 0 {
			zeroCount++
			continue
		}
		product *= num
	}

	result := make([]int, 0, len(nums))
	for _, num := range nums {
		switch {
		case zeroCount > 1:
			result = append(result, 0)
		case zeroCount == 1 && num == 0:
			result = append(result, product)
		case zeroCount == 1:
			result = append(result, 0)
		default:
			result = append(result, product/num)
		}
	}
	return result
}
