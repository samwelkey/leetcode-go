func romanToInt(s string) int {
    m := make(map[byte]int)
    initRoman(m)
    count := 0
    b := []byte(s)
    for i:=0;i < len(b) ;i++{
        if i+1 < len(b) && m[b[i]] < m[b[i+1]]{
            count -= m[b[i]]
        }else{
            count += m[b[i]]
        }
        
        
    }
    
    return count
    
}

func initRoman(m map[byte]int) {
    m['I'] = 1
    m['V'] = 5
    m['X'] = 10
    m['L'] = 50
    m['C'] = 100
    m['D'] = 500
    m['M'] = 1000
}
