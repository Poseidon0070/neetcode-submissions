func containsNearbyDuplicate(nums []int, k int) bool {
	mp := make(map[int][]int)
	for i:=0;i<len(nums);i++ {
		curr := nums[i]
		mp[curr] = append(mp[curr], i)
	}
	for _, arr := range mp {
		if len(arr) < 2 {
			continue
		}
		for i:=1;i<len(arr);i++ {
			if arr[i]-arr[i-1] <= k {
				return true
			}
		}
	}
	return false
}
