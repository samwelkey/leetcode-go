func getPermutation(n int, k int) string {
	nn := 1
	temp := n
	if n > 3{
		for temp-1 > 1{
			nn *= (temp-1)
			temp -= 1
		}
		nn += 1
	}else{
		nn = n
	}

	if n == 0{
		return ""
	}

	if nn == 1{
		return "1"
	}


	s := make([]int,0)
	for i:=0;i<n;i++{
		s = append(s,i+1)
	}
	str := ""

	for len(s) > 0 {

		index := (k-1)/(nn-1)
		str = str+strconv.Itoa(s[index])
		s = append(s[:index],s[index+1:]...)
		k = k-index*(nn-1)
		n -= 1
		temp := n
		if n > 3{
			nn = 1
			for temp-1 > 1{
				nn *= (temp-1)
				temp -= 1
			}
			nn += 1
		}else{
			nn = n
		}

		if nn == 1{
			str = str+strconv.Itoa(s[0])
			break
		}

	}
	return str
}
