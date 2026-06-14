func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	arr := make([]int, 26)
	n := len(s)
	for i:=0;i<n;i++ {
		cis := s[i]-'a'
		cit := t[i]-'a'
		arr[cis]++
		arr[cit]--
	}
	for i:=0;i<25;i++ {
		if arr[i] != 0 {
			return false
		}
	}
	return true
}
