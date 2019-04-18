func permuteUnique(nums []int) [][]int {
    var out [][]int
    sort.Ints(nums)
	addOne([]int{},nums,&out)
	return out
}

func addOne(now,remain []int,out *[][]int){
	if len(remain) == 0 {
		*out = append(*out,now)
		return
	}
    m := make(map[int]struct{})
	for i:=0;i<len(remain);i++{
        
        
		raw,_ := json.Marshal(remain)
        rawNow,_:= json.Marshal(now)
		var r,n []int
		json.Unmarshal(raw,&r)
        json.Unmarshal(rawNow,&n)
        
        if _,ok := m[r[i]] ;ok{
            continue
        }
        
        m[r[i]] = struct{}{}
        
        
		nextNow := append(n,r[i])
		nextRemain := append(r[:i],r[i+1:]...)
		addOne(nextNow,nextRemain,out)
	}
}
