func generateMatrix(n int) [][]int {
    var rowStart,colStart = 0,0
    var rowEnd,colEnd = n-1,n-1
    
    var v = 1
    
    ans := make([][]int,n,n)
    for i:=0;i<len(ans);i++{
        ans[i] = make([]int,n,n)
    }
    
    for rowStart <= rowEnd && colStart <= colEnd{
        for i:= colStart;i<=colEnd;i++{
            ans[rowStart][i] = v
            v++
        }
        
        rowStart ++
        
        for i:= rowStart;i<=rowEnd;i++{
            ans[i][colEnd] = v
            v++
        }
        
        colEnd--
        
        for i:=colEnd;i>=colStart;i--{
            ans[rowEnd][i] = v
            v++
        }
        
        rowEnd --
        
        for i:= rowEnd;i>=rowStart;i--{
            ans[i][colStart] = v
            v++
        }
        
        colStart ++
    }
    
    return ans
}
