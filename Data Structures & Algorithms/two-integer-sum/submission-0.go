func twoSum(nums []int, k int) []int {
    mp := make(map[int]int)
	for i:=0;i<len(nums);i++ {
		curr := nums[i]
		req := k - curr 
		if mp[req] != 0 {
			return []int{mp[req]-1, i}
		}
		mp[curr] = i+1
	}
	return []int{}
}
