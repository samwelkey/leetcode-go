func groupAnagrams(strs []string) [][]string {
    
    if len(strs) == 0{
        return [][]string{}
    }
    var newStrs []string
    for i:= 0;i < len(strs);i++{
        b := []byte(strs[i])
        sort.Slice(b, func(i, j int) bool {
                if b[i] < b[j]{
                        return true
                }
                return false
            })
        newStrs = append(newStrs,string(b))
    }

    m := make(map[string][]int)
    for k,v := range newStrs{
        if _,ok := m[v];!ok{
            m[v] = []int{k}
            continue
        }

        m[v] = append(m[v],k)
    }
    var ans [][]string
    for _,v := range m{
        var temp []string
        for _,vv :=range v{
            temp = append(temp,strs[vv])
        }
        ans = append(ans,temp)
    }

    return ans
}
