permute(nums []int) [][]int {
	var out [][]int
	addOne([]int{},nums,&out)
	return out
}

func addOne(now,remain []int,out *[][]int){
	if len(remain) == 0 {
		*out = append(*out,now)
		return
	}

	for i:=0;i<len(remain);i++{
		raw,_ := json.Marshal(remain)
		var r []int
		json.Unmarshal(raw,&r)
		nextNow := append(now,r[i])
		nextRemain := append(r[:i],r[i+1:]...)
		addOne(nextNow,nextRemain,out)
	}
}
