func fourSum(nums []int, target int) [][]int {
	var result [][]int
	m := make(map[[4]int]int)
	
	if len(nums) == 0{
		return [][]int{}
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for l := j + 1; l < len(nums); l++ {
				for k := l + 1; k < len(nums); k++ {
					if nums[i]+nums[j]+nums[l]+nums[k] == target {
						tmp := []int{nums[i], nums[j], nums[l], nums[k]}
						sort.Ints(tmp)
						arrays := [4]int{tmp[0], tmp[1], tmp[2], tmp[3]}
						_, exist := m[arrays]
						if exist {
							continue
						}

						m[arrays] = 1

						result = append(result, tmp)

					}
				}
			}
		}
	}
	return result
}
