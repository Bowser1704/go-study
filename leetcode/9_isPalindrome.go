package main

import "fmt"
//40ms
func isPalindrome(x int) bool {
    if x<0 {
        return false
    }
    div := 1
    for x/div >=10 {
        div *= 10
    }
    for x>0 {
        if x/div != x%10 {
            return false
        }
        x = (x%div)/10
        div /= 100
    }
    return true
}
//32ms 更复杂的算法，
func isPalindrome2(x int) bool {
    if x<0 {
        return false
    }
    str := strconv.Itoa(x)
    length := len(str)
    for i:=0; i<=length/2-1; i++ {
        if str[i] != str[length-i-1] {
            return false
        }
    }
    return true
}
func main(){
    fmt.Println("Hello, World!")
}

