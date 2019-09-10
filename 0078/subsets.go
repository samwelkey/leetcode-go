func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	res := make([][]int, 0)
	backtracking(nums, []int{}, &res)
	return res
}

func backtracking(nums, path []int, result *[][]int) {
	*result = append(*result, path)
	var dup *Dup
	once := sync.Once{}
	for i := 0; i < len(nums); i++ {
		if dup != nil && nums[i] == dup.dup {
			continue
		}

		once.Do(func() {
			dup = &Dup{}
		})
		dup.dup = nums[i]
		newPath := append([]int{}, path...)
		backtracking(nums[i+1:], append(newPath, nums[i]), result)
	}
}

type Dup struct {
	dup int
}
