package main

import (
    "fmt"
    "errors"
)

type Set struct{
    s := make([]string, 10)
}

func (set *Set) AddItem(item string) (err error){
    set.s = append(set.s,item)
    return nil
}

func (set *Set) Contain(item string) (exist bool){
    for _, val in range set.s {
        if val==item{
            exist = True
            return exist
        }
    }
    return False
}

func (set *Set) Size() (size int){
    size = len(set.s)
    return size
}

func (set *Set) Delete(item string) (err error){
    position := 0
    for i, value in range set.s {
        if item==value{
            position = i
            break
        }
    }
    set.s = append(set.s[:position], set.s[position+1:]...)
    return nil
}
func (set *Set) Items() (sli []string, err error){
    return set.s, nil
}

func main(){

}
