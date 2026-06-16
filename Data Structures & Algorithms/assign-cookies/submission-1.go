func findContentChildren(g []int, s []int) int {
	sort.Slice(g, func(i, j int) bool {
		return g[i] < g[j]
	})
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	// fmt.Println(g)
	// fmt.Println(s)
	i, j, res := 0, 0, 0;
	for i < len(g) && j < len(s) {
		if g[i] > s[j] {
			j++
		} else {
			i++ 
			j++ 
			res++
		}
	}
	return res

}
