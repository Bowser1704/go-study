//参考https://tundrazone.com/?p=191
func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    var newNums [][]int
    numsLen := len(nums)

    //for i,_ :=  range nums[:numsLen-2]{  Line 7: panic: runtime error: slice bounds out of range
    for i:=0; i<numsLen; i++{
        if i > 0 && nums[i] == nums[i-1]{
            continue
            //去重
        }
        target := -1 * nums[i]
        j, k := i+1, numsLen-1
        // k 定位最后一个元素
        for j < k{
            if nums[j]+nums[k] == target{
                x := []int{nums[i], nums[j], nums[k]}
                newNums = append(newNums, x)
                j += 1
                for j<k && nums[j] == nums[j-1]{
                    j +=1
                    //去重
                }
            }else if nums[j] + nums[k] < target {
                j += 1
            }else {
                k -= 1
            }
        }
    }
    return newNums
}
