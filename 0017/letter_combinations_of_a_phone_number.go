func letterCombinations(digits string) []string {
	m = make(map[byte][]byte)
    output = []string{}
	m['2'] = []byte{'a', 'b', 'c'}
	m['3'] = []byte{'d', 'e', 'f'}
	m['4'] = []byte{'g', 'h', 'i'}
	m['5'] = []byte{'j', 'k', 'l'}
	m['6'] = []byte{'m', 'n', 'o'}
	m['7'] = []byte{'p', 'q', 'r', 's'}
	m['8'] = []byte{'t', 'u', 'v'}
	m['9'] = []byte{'w', 'x', 'y', 'z'}
	b := []byte(digits)
    
    if len(digits) == 0 {
        return []string{}
    }
    
	backtrack("", b)

	return output

}

var output []string
var m map[byte][]byte

func backtrack(combination string, next_digits []byte) {
	if len(next_digits) == 0 {
		output = append(output, combination)
	} else {
		for _, c := range m[next_digits[0]] {
			backtrack(combination+fmt.Sprintf("%c", c), next_digits[1:])
		}
	}

}
