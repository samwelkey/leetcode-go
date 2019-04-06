func longestValidParentheses(s string) int {
	ss := []byte(s)
	var count int
	var longest int
	for i:=0; i<len(s); i++ {
		count = 0
		for j:=i; j<len(s); j++ {
			if ss[j] == '('{
				count ++
			}
			if ss[j] == ')'{
				count --
			}
			if count < 0 {
				break
			}
			if count == 0 && j-i+1 > longest{
				longest = j-i+1
			}
		}
	}
	return longest
}
