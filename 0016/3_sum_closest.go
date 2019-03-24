func threeSumClosest(nums []int, target int) int {
    abs := math.MaxInt32
    var result int
    for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
                tmp := math.Abs(float64(nums[i]+nums[j]+nums[k]-target))
                if  tmp < float64(abs) {
                    result = nums[i]+nums[j]+nums[k]
                    abs = int(tmp)
                } 
			}
		}
	}
    return result
}
