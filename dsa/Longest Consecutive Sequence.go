package dsa

import "slices"

func longestConsecutive(nums []int) int {
	slices.Sort(nums)
	best, count := 1, 1
	if len(nums) == 0 {
		return 0
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		if nums[i] == nums[i-1]+1 {
			count++
		} else {
			count = 1
		}
		best = max(best, count)
	}

	return best
}

// O(n) solution using a set to store the numbers and then checking for the start of a sequence
// func longestConsecutive(nums []int) int {
//     set := make(map[int]struct{}, len(nums))
//     for _, n := range nums {
//         set[n] = struct{}{}
//     }
//     best := 0
//     for n := range set {
//         if _, ok := set[n-1]; ok {
//             continue // not the start of a sequence
//         }
//         length := 1
//         for {
//             if _, ok := set[n+length]; !ok {
//                 break
//             }
//             length++
//         }
//         best = max(best, length)
//     }
//     return best
// }