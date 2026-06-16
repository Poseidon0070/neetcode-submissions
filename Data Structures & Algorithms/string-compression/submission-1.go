func compress(chars []byte) int {
	ans := ""
	i := 0
	ln := 0
	var prev byte = '#'

	for i < len(chars) {
		if chars[i] != prev {
			ans += string(chars[i])
			ln = 1
			prev = chars[i]
			i++
		} else {
			for i < len(chars) && chars[i] == prev {
				i++
				ln++
			}
			ans += strconv.Itoa(ln)
			ln = 0
		}
	}

	fmt.Println(ans)
	copy(chars, ans)
	return len(ans)
}