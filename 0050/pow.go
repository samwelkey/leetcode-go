func myPow(x float64, n int) float64 {
    if n == 0 || x == 1{
        return 1.0
    }
    
    if x == -1 {
        if n%2 == 0 {
            return 1.0
        }
        return -1.0
    }
    
    var base float64
    if n >0{
        base = x
    }else{
        base = 1/x
        n = -n
    }
    var ans float64 = 1
    for i:=0;i<n;i++{
        ans = ans*base
    }
    return ans
    
}
