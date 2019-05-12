package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}
func search(nums []int, target int) int {
	result := -1

	left, right := 0, len(nums) - 1
	middle := (left + right) / 2
	
	if len(nums) == 0 || (target < nums[left] && target> nums[right]){
		return result
	}
	if nums[left] == target{
		return left;
	}
	if nums[right] == target{
		return right
	}

	for left <= right {
		if nums[middle] == target{
			result = middle
			break
		}

		if nums[middle] >= nums[left]{
			if nums[left] <= target && target < nums[middle]{
				right = middle - 1
			}else{
				left = middle + 1
			}
		}else{
			if nums[right]>=target && target>nums[middle]{
				left = middle + 1
			}else{
				right = middle -1 
			}
		}
		middle = (left + right) / 2
	}	
	return result
}
