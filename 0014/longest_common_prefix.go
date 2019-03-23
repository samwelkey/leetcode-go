func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var first byte
	var newStrs []string
	for _, str := range strs {
		if str == "" {
			return ""
		}

		if first == 0 {
			first = []byte(str)[0]
		}

		if first != []byte(str)[0] {
			return ""
		}

		newStrs = append(newStrs, string([]byte(str)[1:]))
	}
	return string(first) + longestCommonPrefix(newStrs)
}
