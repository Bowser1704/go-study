package main

import (
    s"strings"
    "fmt"
    "errors"
)

func Find(stra, strb string) (position int, err error){
    position = s.Index(stra,strb)
    if position < 0{
        return -1, errors.New("not found")
    }
    return position, nil
}

func main(){
    //var stra, strb string
    stra := "xxmuxixx"
    strb := "muxi"
    re, err := Find(stra,strb)
    fmt.Println(re,err)
}
