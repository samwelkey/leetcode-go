func firstMissingPositive(nums []int) int {
    sort.Ints(nums)
    min := 1
    for i:=0; i<len(nums);i++{
        if nums[i] > 0 {
            if nums[i] > min{
                break
            }
            
            if nums[i] == min{
                min ++
            }
        }
    }
    return min
}
