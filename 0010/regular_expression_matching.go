func isMatch(s string, p string) bool {
	sb := []byte(s)
	pb := []byte(p)

	if len(p) == 0 {
		return len(s) == 0
	}
	firstMatch := !(len(sb) == 0) && (pb[0] == sb[0] || pb[0] == '.')

	if len(pb) >= 2 && pb[1] == '*' {
		return isMatch(s, string(pb[2:])) || (firstMatch && isMatch(string(sb[1:]), p))
	} else {
		return firstMatch && isMatch(string(sb[1:]), string(pb[1:]))
	}

}
