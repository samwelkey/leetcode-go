func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var tempList []int

	btrack(&res, tempList, candidates, target, 0)
	return res
}

func btrack(res *[][]int, tmp []int, candidates []int, remain, start int) {
	if remain < 0 {
		return
	} else if remain == 0 {
		t := make([]int,len(tmp))
		copy(t,tmp)
		*res = append(*res, t)
	} else {
		for i := start; i < len(candidates); i++ {
			tmp = append(tmp, candidates[i])
			btrack(res, tmp, candidates, remain-candidates[i], i)
			tmp = tmp[:len(tmp)-1]
		}
	}

}
