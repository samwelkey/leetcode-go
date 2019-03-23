func threeSum(nums []int) [][]int {
	m := make(map[[3]int]byte)
	var ans [][]int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					sl := []int{nums[i], nums[j], nums[k]}
					sort.Ints(sl)
					arrays := [3]int{sl[0],sl[1],sl[2]}

					_, exist := m[arrays]
					if exist {
						break
					}
					m[arrays] = 1
					ans = append(ans, sl)
				}
			}
		}
	}
	return ans
}
