func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    //避免重复输出元素
    var newNums [][]int

    for i,a := range nums{
        for j,b := range nums[i+1:]{
            for _,c := range nums[i+j+1:]{
                if a+b+c == 0{
                    x := []int{a,b,c}
                    newNums = append(newNums,x)
                }
            }
        }
    }
    return newNums
}
