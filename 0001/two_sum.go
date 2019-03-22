func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for k,v := range nums{
        m[v] = k
    }
    
    for i := 0;i < len(nums);i++{
        complement := target - nums[i]
        index,exsit := m[complement]
        if exsit && index > i{
            r := []int{i,index}
            return r
        }
    }
    return nil
}
