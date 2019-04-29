func lengthOfLastWord(s string) int {
    if s == ""{
        return 0
    }
    trimed := strings.Trim(s, " ")
    ss := strings.Split(trimed," ")
    return len(ss[len(ss)-1])
}
