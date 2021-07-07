package main

import "log"

func twoSum(nums []int, target int) []int {
	numLen := len(nums)
	data := make([]int,0)
	for i:=0;i<numLen;i++ {
		for j:=i+1;j<numLen;j++ {
			if nums[i]+nums[j] == target {
				data = append(data,i,j)
			}
		}
	}
	return data
}




func main() {
	nums := []int{2,7,11,15}
	target := 9
	result := twoSum(nums,target)
	log.Println("result is ",result)
}
