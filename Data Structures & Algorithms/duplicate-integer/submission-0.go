func hasDuplicate(nums []int) bool {
    set := map[int]bool{}
    for i := range nums {
        num := nums[i]
        if _, ok := set[num]; ok {
            return true
        }
        set[num] = true
    }
    return false
}
