func myAtoi(str string) int {
	b := []byte(str)
	atHead := true
	var bb []byte
	for i := 0; i < len(b); i++ {
		if b[i] == ' ' {
			if atHead {
				continue
			} else {
				break
			}
		}

        if b[i] >= '0' && b[i] <= '9' || (b[i] == '-'||b[i] == '+') && atHead {
			atHead = false
			bb = append(bb, b[i])

			continue
		}
		break
	}

	if len(bb) == 0 {
		return 0
	}

	res, _ := strconv.Atoi(string(bb))

	if res > math.MaxInt32 {
		return math.MaxInt32
	}

	if res < math.MinInt32 {
		return math.MinInt32
	}

	return res
}
