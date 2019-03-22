func convert(s string, numRows int) string {
    if numRows <=1 {
        return s
    }
    var res []byte
    b := []byte(s)
    
    size := 2* numRows -2
    for i := 0;i<numRows; i ++{
        for j:= i;j <len(b);j+=size{
            res = append(res,b[j])
            tmp := j+size-2*i
            if i!=0 && i != numRows-1 && tmp < len(b){
                res = append(res,b[tmp])
            }
        }
    }
    return string(res)
}


