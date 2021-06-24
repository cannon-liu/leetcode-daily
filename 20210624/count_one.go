package main

import "fmt"

func hammingWeight(num uint32) int {
	cnt := 0
	for i:=0; i<32;i++ {
		if 1<<i&num >0 {
			cnt ++
		}
	}
	return cnt
}

func hammingWeight2(num uint32) int {
	cnt := 0
	for ;num > 0;num = num & (num-1){
		cnt++
	}
	return cnt
}

func main() {
	sum := hammingWeight2(00000000000000000000000000001011)
	fmt.Println("sum is ",sum)
}
