/*
1. iterate through each string create a freq array and then craete a bitmask out of the frequency created
2. create map of string and array of string that pushes each string with key of its bitmask
3. at the end group and return each group based on mask
*/

func groupAnagrams(strs []string) [][]string {
	n := len(strs)
	maskMap := make(map[string][]string)
	for i:=0;i<n;i++ {
		fa := make([]int, 26)
		curr := strs[i]

		for j:=0;j<len(curr);j++ {
			fa[curr[j]-'a']++
		}

		mask := ""
		for j:=0;j<26;j++ {
			mask += string(fa[j])
		}

		maskMap[mask] = append(maskMap[mask], curr)
	}
	res := [][]string{}
	for _, val:= range maskMap {
		res = append(res, val)
	}

	return res
}
