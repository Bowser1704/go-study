func longestValidParentheses(s string) int {
    dp, res := make([]int, len(s)), 0
    for i,_ := range s{
        if s[i]==')'{
            if i>0 && s[i-1]=='('{
                dp[i] = dp[i-1] + 2
            }
            if i>0 && s[i-1]==')'{
                if i-dp[i-1]>0 && s[i-dp[i-1]-1]=='('{
                    if i-dp[i-1]-1 > 0{
                        dp[i] = dp[i-1] + dp[i-dp[i-1]+2] + 2
                    }else{
                        dp[i] = dp[i-1] + 2
                    }
                }
            }
        }
        if dp[i] > res{
            res = dp[i]
        }
    }
    return res
}
