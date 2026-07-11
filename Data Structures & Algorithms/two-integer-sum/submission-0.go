func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
		l := nums[i]
		for j := i + 1; j < len(nums); j++ {
			r := nums[j]
			if l + r == target {
				return []int{min(i, j), max(i, j)}
			}
		}
	}

	return nil
}
