package main

import "fmt"

func maxPoints(points [][]int) int {
	n := len(points)
	if n <= 2 {
		return n
	}
	ans := 0
	for i,p := range points {
		if ans >= n -i || ans > n/2 {
			break
		}
		cnt := map[int]int{}
		for _,q := range points[i+1:] {
			x,y := p[0]-q[0],p[1]-q[1]
			if x == 0 {
				y = 1
			} else if y == 0 {
				x = 1
			} else {
				if y < 0 {
					// 因为最后算key时，并不是使用dy/dx,
					// dy + dx * num并不能友好的考虑同条线上，方向相反的情况，
					// dx dy都求反，相当于i指向j变成了j指向i,
					// 即i点上方的点指向i点，保证方向相同右上=>左下，左上=>右下
					x,y = -x,-y
				}
				g := gcd(abs(x),abs(y))
				x /= g
				y /= g
			}
			// 20001为(2*10^4),为什么有个2乘呢，因为点xy取值范围[-10^4,10^4],取计算亮点
			// dx的取值范围[-10^4-10^4,10^4-(-10^4)],dy取值范围为整数[0,20000]
			// dx为什么要乘这个取值范围呢？ 为了确保大于dy？空出前面20000，留出dy的位置，确保数值key不会重复
			cnt[y+x*20001]++
		}
		for _,c := range  cnt {
			ans = max(ans,c+1)
		}
	}

	return ans
}

func gcd(a,b int) int {
	for a != 0 {
		a,b = b%a,a
	}
	return b
}


func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}


func main() {

	points := make([][]int,0)
	var a []int
	a = []int{1,1}
	points = append(points,a)
	a = []int{2,2}
	points = append(points,a)
	a = []int{3,3}
	points = append(points,a)
	result := maxPoints(points)
	fmt.Println("result is ",result)
}