package main

import (
    "fmt"
)

func isValid(s string) bool {
    if len(s)<2{
        return false
    }
    sum := 0
    for i:=0; i<len(s); i++{
        if s[i] == '('{
            sum += 1
            if s[i+1] == '}' || s[i+1] == ']'{
                return false
            }
        }else if s[i] == ')'{
            sum -= 1
        }
        if s[i] == '{'{
            sum += 1
            if s[i+1] == ')' || s[i+1] == ']'{
                return false
            }
        }else if s[i] == '}'{
            sum -= 1
        }
        if s[i] == '['{
            sum += 1
            if s[i+1] == '}' || s[i+1] == ')'{
                return false
            }
        }else if s[i] == ']'{
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
