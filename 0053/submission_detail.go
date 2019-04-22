func maxSubArray(nums []int) int {
    if len(nums) == 0{
        return 0
    }
    dp := make(map[int]int)
    dp[0] = nums[0]
    max := dp[0]
    
    for i:= 1;i<len(nums);i++{
        if dp[i-1] > 0 {
            dp[i] = dp[i-1] + nums[i]
        }else{
            dp[i] = nums[i]
        }
        
        if dp[i] > max{
            max = dp[i]
        }
    }
    
    return max
}
