func hasDuplicate(nums []int) bool {
    var mp map[int]int = make(map[int]int)
	for i := 0; i < len(nums) ; i++ {
		ele := nums[i]
		if mp[ele] != 0 {
			return true
		}
		mp[ele]++
	}
	return false
}
