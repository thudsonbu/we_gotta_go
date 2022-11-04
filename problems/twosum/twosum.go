package twosum

func TwoSum(nums []int, target int) []int {
	if nums == nil {
		return []int{0, 0}
	}

	numMap := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		numMap[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		c := target - nums[i]

		v, ok := numMap[c]

		if !ok {
			continue
		}

		if v == i {
			continue
		}

		return []int{v, i}
	}

	return []int{0, 0}
}
