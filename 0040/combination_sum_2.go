func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	var tempList []int

	sort.Ints(candidates)

	btrack(&res, tempList, candidates, target, 0)

	m := make(map[string]struct{})
	var newRes [][]int
	for _, r := range res {
		a, _ := json.Marshal(r)
		if _, ok := m[string(a)]; !ok {
			newRes = append(newRes, r)
			m[string(a)] = struct{}{}
		}
	}

	return newRes
}

func btrack(res *[][]int, tmp []int, candidates []int, remain, start int) {
	if remain < 0 {
		return
	} else if remain == 0 {
		t := make([]int, len(tmp))
		copy(t, tmp)
		*res = append(*res, t)
	} else {
		for i := start; i < len(candidates); i++ {
			tmp = append(tmp, candidates[i])
			btrack(res, tmp, candidates, remain-candidates[i], i+1)
			tmp = tmp[:len(tmp)-1]
		}
	}

}

