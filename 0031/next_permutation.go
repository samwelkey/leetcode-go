func nextPermutation(nums []int)  {
    var i = len(nums) -2
    for i >= 0 && nums[i+1] <= nums[i] {
        i--
    }
    
    if i >= 0 {
        var j = len(nums) - 1
        for j >= 0 && nums[j] <= nums[i] {
            j--
        }
        swap(nums,i,j)
    }
    reverse(nums,i+1)
}

func reverse(nums []int,start int){
    var i,j = start,len(nums)-1
    for i<j{
        swap(nums,i,j)
        i++
        j--
    }
}

func swap(nums []int,i,j int){
    temp := nums[i]
    nums[i] = nums[j]
    nums[j] = temp
}
