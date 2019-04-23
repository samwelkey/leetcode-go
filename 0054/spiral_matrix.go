func spiralOrder(matrix [][]int) []int {
    ans := make([]int,0)
    if len(matrix) == 0{
        return ans
    }
    if len(matrix[0]) == 0{
        return ans
    } 
    var rowStart,colStart = 0,0
    var rowEnd = len(matrix)-1
    var colEnd = len(matrix[0])-1
        
    for rowStart <= rowEnd && colStart <= colEnd{
        for i:= colStart;i<=colEnd;i++{
            ans = append(ans,matrix[rowStart][i])
        }
        rowStart++
        if rowStart > rowEnd {
            break
        }
        
        for i:= rowStart ; i<= rowEnd; i++{
            ans = append(ans,matrix[i][colEnd])
        }
        colEnd--
        if colStart > colEnd{
            break
        }
        
        for i:= colEnd;i>=colStart;i--{
            ans = append(ans,matrix[rowEnd][i])
        }
        
        rowEnd--
        if rowStart > rowEnd {
            break
        }
        
        for i:= rowEnd;i>=rowStart;i--{
            ans = append(ans,matrix[i][colStart])
            
        }
        colStart ++
        
    }
    
    return ans
    
    
    
}
