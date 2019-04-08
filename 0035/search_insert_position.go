func searchInsert(nums []int, target int) int {
	var i = 0
	for i <= len(nums) {
        if i == len(nums) {
            break
        }
		if nums[i] < target {
			i++
			continue
		}else{
			break
		}
	}
	return i
}
