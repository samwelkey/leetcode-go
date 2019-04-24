canJump(nums []int) bool {
    var curFurthest, curEnd float64
	for i := 0; i < len(nums)-1; i++ {
		curFurthest = math.Max(curFurthest, float64(i+nums[i]))
		if i == int(curEnd) {
			curEnd = curFurthest
		}
	}
    if int(curEnd) >= len(nums)-1{
        return true
    }
    
    return false
}
