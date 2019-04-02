func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}
	wordLen := len(words[0])
	ss := []byte(s)

	out := make([]int, 0)

	var flag bool
	m := make(map[string]int)
	for _, w := range words {
		m[w] ++
	}
	for i := 0; i < len(ss); i++ {

		var mm map[string]int

		deepCopy(&mm, m)

		index := i
		if i+wordLen > len(ss) {
			continue
		}

		if mm[string(ss[i:i+wordLen])] == 0 {
			continue
		}

		flag = true
		for j := 0; j < len(words); j++ {
			if i+(1+j)*wordLen > len(ss) {
				flag = false
				break
			}
			count, ok := mm[string(ss[i+j*wordLen:i+(1+j)*wordLen])]
			if !ok {
				flag = false
				break
			}

			if count == 0 {
				flag = false
				break
			}

			mm[string(ss[i+j*wordLen:i+(1+j)*wordLen])]--

		}
		if flag {
			out = append(out, index)
		}

	}

	return out
}

func deepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}
