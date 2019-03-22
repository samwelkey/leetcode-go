func longestPalindrome(s string) string {
	var ss string
	longest := 0
	b := []byte(s)

	for i := 0; i < len(b); i++ {

		var deltaOdd,deltaEven int
		for (i-deltaOdd) >= 0 && (i+deltaOdd) < len(b) {
			if b[i-deltaOdd] == b[i+deltaOdd] {
				if (2*deltaOdd + 1) >= longest {
					longest = 2*deltaOdd + 1
					ss = string(b[i-deltaOdd : i+deltaOdd+1])
				}
				deltaOdd ++
			} else{
				break
			}
		}
        
        for (i - deltaEven) >= 0 && (i+1+deltaEven) < len(b){
            if b[i-deltaEven] == b[i+1+deltaEven] {
                if 2*deltaEven + 2 >= longest{
                    longest = 2*deltaEven + 2
                    ss = string(b[i-deltaEven:i+2+deltaEven])
                }
                deltaEven ++
            }else{
                break
            }
        }
        
        
	}
	return ss
}
