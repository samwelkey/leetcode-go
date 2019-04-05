func divide(dividend int, divisor int) int {
	var flag = 1
	var out = 0
	if dividend < 0{
		flag = -flag
		dividend = -dividend
	}
	if divisor < 0{
		flag = -flag
		divisor = -divisor
	}
	for dividend >= divisor{
		dividend -= divisor
		out ++
	}
	if flag == 1 || out == 0{
		if out > math.MaxInt32 {
			return math.MaxInt32
		}
		return out
	}
	return -out
}
