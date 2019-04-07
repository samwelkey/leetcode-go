func searchRange(nums []int, target int) []int {
	res := make([]int, 0)
	count := -1
	for i := 0; i < len(nums); i ++ {
		if nums[i] == target {
			res = append(res, i)
			j := i
			for j < len(nums) {
				if nums[j] == target {
					count ++
					j ++
				} else {
					break
				}

			}
			res = append(res, i+count)
			break

		}
	}

	if len(res) == 0 {
		return []int{-1, -1}
	}

	return res
}
