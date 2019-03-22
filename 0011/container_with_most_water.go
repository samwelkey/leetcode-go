func maxArea(height []int) int {
    max,i,j :=0.0,0, len(height)-1
    
    for i<j{
        max = math.Max(max,math.Min(float64(height[i]),float64(height[j])) * float64(j-i))
        if height[i] < height[j]{
            i ++
        }else{
            j--
        }
    }
    return int(max)
}
