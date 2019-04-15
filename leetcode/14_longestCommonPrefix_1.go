func longestCommonPrefix(strs []string) string {
    if cap(strs) == 0 {
        return ""
    }
    prefix := strs[0]
    for str := range strs {
        for !strings.HasPrefix(str, prefix){
            prefix = prefix.
        }
    }
}
