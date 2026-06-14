func longestCommonPrefix(strs []string) string {
	ssf := ""
	isf := 0
	for true {
		for i:=0;i<len(strs);i++ {
			curr := strs[i]
			if len(curr) <= isf || curr[isf] != strs[0][isf] {
				return ssf
			}
		}
		ssf += string(strs[0][isf])
		isf++
	}
	return ssf
}
