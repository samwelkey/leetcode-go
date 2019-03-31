func removeElement(nums []int, val int) int {
    if len(nums) == 0{
        return 0
    }
    
    var i = 0
    for j:=0;j < len(nums) ;j++{
        if nums[j]  != val{
            nums[i] = nums[j]
            i ++
        }
    }
    nums = nums[:i]
    return i
}
