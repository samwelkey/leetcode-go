func lengthOfLongestSubstring(s string) int {
	var ans, window, i, j int
	b := []byte(s)
	m := make(map[byte]uint8)
	for i < len(s) && j < len(s) {
		_, exist := m[b[j]]
		if !exist {
			m[b[j]] = 1
			j++
			window = j - i
			if ans < window {
				ans = window
			}
		} else {
			delete(m, b[i])
			i++
		}
	}
	return ans
}
