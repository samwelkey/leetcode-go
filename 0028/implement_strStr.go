func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	h := []byte(haystack)
	n := []byte(needle)
	var ok bool
	for i := 0; i < len(haystack); i++ {
		if h[i] == n[0] {
			if i+len(needle) <= len(haystack) {
				ok = true
				for j := 0; j < len(needle); j++ {
					if h[i+j] != n[j] {
						ok = false
						break
					}
				}
				if ok {
					return i
				}

			} else {
				return -1
			}

		}

	}
	return -1
}
