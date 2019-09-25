package main

import "fmt"

//16ms
func permute(nums []int) [][]int {
    if len(nums) == 1{
        return [][]int{nums}
    }
    var res [][]int

    for _,s := range permute(nums[1:]) {
        length := len(s)
        for i:=0; i<length; i++ {
            temp := make([]int, length+1)
            copy(temp[:i], s[:i])
            copy(temp[i+1:], s[i:])
            temp[i] = nums[0]
            res = append(res, temp)
        }
        //Don't forget the last number
        res = append(res, append(s, nums[0]))
    }
    return res
}

//4ms 
func permute1(nums []int) [][]int {
    return subNumberSlice(nums)
}

func subNumberSlice(nums []int) [][]int {
    if len(nums) == 0 {
        return nil
    }
    if len(nums) == 1 {
        return [][]int{{nums[0]}}
    }
    if len(nums) == 2 {
        return [][]int{{nums[0], nums[1]}, {nums[1], nums[0]}}
    }

    result := [][]int{}
    for index, value := range nums {
        var numsCopy = make([]int, len(nums))
        copy(numsCopy, nums)
        numsSubOne := append(numsCopy[:index],numsCopy[index+1:]...)
        valueSlice := []int{value}
        newSubSlice := subNumberSlice(numsSubOne)
        for _, newValue := range newSubSlice {
            result = append(result, append(valueSlice, newValue...))
        }
    }
    return result
}

//main
func main(){
    fmt.Println("Hello, World!")
}

