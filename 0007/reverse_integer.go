func reverse(x int) int {
    var num int
    for x != 0 {
        digit := x % 10
        num = num*10 + digit
        x = x / 10
    }
    if num >= math.MinInt32 && num <= math.MaxInt32 {
        return num
    }
    return 0
}
