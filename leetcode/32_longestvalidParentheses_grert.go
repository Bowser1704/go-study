func longestValidParentheses(s string) int {
    length := len(s)
    resLeft := calculate(s, 0, 1, length, '(')
    resRight := calculate(s, length-1, -1, -1, ')')
    if resLeft < resRight{
        return resRight
    }
    return resLeft
}

func calculate(s string, i,flag,end int, cTem byte) int {
    // 定义cTem 为byte结构？
    var max, sum, currLen, validLen int
    for ; i != end; i += flag{
        if s[i] == cTem{
            sum += 1
        }else{
            sum -= 1
        }
        currLen += 1
        if sum < 0{
            // 中断连续了，前面的最长过去了，用max保存
            if max < validLen {
                max = validLen
            }
            sum = 0
            currLen = 0
            validLen = 0
        } else if sum == 0{
            // 一直为有效的
            validLen = currLen
        }
    }
    if max > validLen {
        return max
    }
    return validLen
}
