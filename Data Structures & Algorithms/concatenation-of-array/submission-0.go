func getConcatenation(nums []int) []int {
	var n int = len(nums)
    var res []int = make([]int, 2 * n)
	for i:=0;i<n;i++ {
		res[i] = nums[i]
		res[n+i] = nums[i]
	}
	return res
}
