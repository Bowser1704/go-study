func longestCommonPrefix(strs []string) string {
    if cap(strs) == 0 {
        return ""
    }
    var res string
    for i := 0; i < len(strs[0]); i++{
        j := 1
        for ; j < len(strs); j++{
            if i >= len(strs[j]) || strs[0][i] != strs[j][i] {
                //这里是>= 要仔细思考一下
                return res
            }
        }
        if j == len(strs){
            res += string(strs[0][i])
                //这里要做强制类型转换否则报错invalid operation: res += strs[0][i] (mismatched types string and byte) (solution.go)
        }
    }
    return res
}
