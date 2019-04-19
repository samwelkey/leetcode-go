func rotate(matrix [][]int)  {
    if len(matrix) == 0 {
        return
    }
    for i:=0;i < len(matrix);i++{
        for j:=i;j<len(matrix);j++{
            if i == j {
                continue
            }
            
            temp := matrix[i][j]
            matrix[i][j] = matrix[j][i]
            matrix[j][i] = temp
            
        }
    }
    
    for i:=0;i < len(matrix)/2;i++{
        for j:=0;j<len(matrix);j++{
            temp:=matrix[j][i]
            matrix[j][i] = matrix[j][len(matrix)-1-i]
            matrix[j][len(matrix)-1-i] = temp
        }
    }
    return
}
