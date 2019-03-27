func generateParenthesis(n int) []string {
	ans := make([]string,0)
	backtrack1(&ans, "", 0, 0, n)
	return ans
}

func backtrack1(ans *[]string, cur string, open int, close int, max int) {
	if len(cur) == 2*max {
		*ans = append(*ans, cur)
		return
	}

	if open < max {
		backtrack1(ans, cur+"(", open+1, close, max)
	}

	if close < open {
		backtrack1(ans, cur+")", open, close+1, max)
	}
}
