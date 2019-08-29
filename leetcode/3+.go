package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
    len := len(s)
    var shift, max int
    for i:=0; i<len; i++ {
        for j:=shift; j<i; j++ {
            //fmt.Println(i,j)
            if(s[j] == s[i]) {
                shift = j+1
                break   //重点    
            }
        }
        if i-shift+1 > max {
            max = i-shift+1
        }
        fmt.Println(max)
    }
    return max
}

func lengthOfLongestSubstring2(s string) int {
    length := len(s)
    var start, end, ans int
    end = -1
    array:=[256]int{}
    for start < length {
        fmt.Println(start, end, array['a'])
        if end+1<length && array[s[end+1]] == 0 {
            array[s[end+1]]++   //顺序很重要
            end++
        }else {
            array[s[start]]--   //顺序很重要
            start++
        }
        ans = max(end-start+1, ans)
    }
    return ans
}
func max(a, b int) int {
    if a>b {
        return a
    }else {
        return b
    }
}
func main(){
    //fmt.Println(lengthOfLongestSubstring2("hijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789hijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789hijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789hijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789hijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789hijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"))
    fmt.Println(lengthOfLongestSubstring2("abcabcaa"))
}

