func twoSum(nums []int, target int) []int {
    indices := map[int]int{}

	for i, num := range nums {
		diff := target - num
		if v, ok := indices[diff]; ok {
			return []int{v, i}
		}
		indices[num] = i
	}

	return []int{}
}
