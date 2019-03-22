func isPalindrome(x int) bool {
    origin := x
    
    if x == 0 {
        return true
    }
    
    s := strconv.Itoa(x)
    b:= []byte(s)
    
    if b[0] == '+' || b[0] == '-'{
        return false
    }
    
    var res int
    
    for x != 0 {
        digit := x % 10
        res = res * 10 + digit
        x = x / 10
    }
    
    if res == origin {
        return true
    }
    
    return false
    
}
