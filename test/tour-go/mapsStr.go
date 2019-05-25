package main

import (
    "golang.org/x/tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    ret := make(map[string]int)
    //var ret map[string]int 之后还是要给ret赋值没有意义
    arr := strings.Fields(s)
    for _,x := range arr{
        ret[x]++
    }
    return ret
}

func main() {
    wc.Test(WordCount)
}
