func jump(nums []int) int {
	var count, curFurthest, curEnd float64
	for i := 0; i < len(nums)-1; i++ {
		curFurthest = math.Max(curFurthest, float64(i+nums[i]))
		if i == int(curEnd) {
			count++
			curEnd = curFurthest
		}
	}
	return int(count)
}
