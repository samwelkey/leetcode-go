func intToRoman(num int) string {
    m := make(map[int]byte)
    initRoman(m)
    times := 1
    var ans []byte
    
    for num != 0 {
        var tmp []byte
        digit := num % 10
        if digit >= 0 && digit <= 3{
            for i:= 0;i<digit;i++{
                tmp = append(tmp,m[1*times])
            }
        }
        
        if digit == 4 {
            tmp = append(tmp,m[1*times],m[5*times])
        }
        
        if digit >= 5 && digit <= 8 {
            tmp = append(tmp,m[5*times])
            
            for i:= 0; i < digit-5;i ++ {
                tmp = append(tmp,m[1*times])
            }
        }
        
        if digit == 9 {
            tmp = append(tmp,m[1*times],m[10*times])
        }
        
        ans = append(tmp,ans...)
        num = num/10
        times *= 10
        
    }
    
    
    return string(ans[:])
}

func initRoman(m map[int]byte) {
    
    m[1] = 'I'
    m[5] = 'V'
    m[10] = 'X'
    m[50] = 'L'
    m[100] = 'C'
    m[500] = 'D'
    m[1000] = 'M'
    
}

