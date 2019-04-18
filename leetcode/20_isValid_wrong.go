package main

import (
    "fmt"
    "strings"
)

func Slash(r rune) rune{
    if r == '{' || r == '[' || r == '('{
        return '('
    }
    return ')'
}

func isValid(s string) bool {
    ss := strings.Map(Slash, s)
    sum := 0
    fmt.Println(ss)
    for _,x := range ss{
        if x == '('{
            sum += 1
        }else{
            sum -= 1
        }
        if sum < 0{
            return false
        }
    }
    return true
}

func main(){
    fmt.Println(isValid("(]"))
}
